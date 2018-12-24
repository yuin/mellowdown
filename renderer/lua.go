package renderer

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
	lua "github.com/yuin/gopher-lua"
	"go.uber.org/multierr"
	luar "layeh.com/gopher-luar"
)

func toString(l *lua.LState) int {
	arg := l.CheckAny(1)
	if ud, ok := arg.(*lua.LUserData); ok {
		if s, ok := ud.Value.([]byte); ok {
			l.Push(lua.LString(string(s)))
		}
	} else {
		l.Push(lua.LString(arg.String()))
	}
	return 1
}

type LuaRenderer struct {
	scripts   string
	l         *lua.LState
	renderers []*lua.LTable
}

func NewLuaRenderer() Renderer {
	return &LuaRenderer{}
}

func (r *LuaRenderer) Name() string {
	return "lua"
}

func (r *LuaRenderer) SetOutputDirectory(path string) {
	var e error
	for _, lr := range r.renderers {
		if err := r.call(lr, "set_output_directory", luar.New(r.l, path)); err != nil {
			e = appendError(e, err)
		}
	}
	if e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
	}
}

func (r *LuaRenderer) SetFile(path string) {
	var e error
	for _, lr := range r.renderers {
		if err := r.call(lr, "set_file", luar.New(r.l, path)); err != nil {
			e = appendError(e, err)
		}
	}
	if e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
	}
}

func (r *LuaRenderer) AddOption() {
	flag.StringVar(&r.scripts, "lua", "", "comma separeted lua renderers")
}

func (r *LuaRenderer) InitOption() {
	l := lua.NewState()
	l.SetGlobal("tostring", l.NewFunction(toString))
	if len(r.scripts) == 0 {
		return
	}
	for _, script := range strings.Split(r.scripts, ",") {
		if err := l.DoFile(script); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to load a lua renderer: %s", err.Error())
		} else {
			r.renderers = append(r.renderers, l.Get(-1).(*lua.LTable))
			l.Pop(1)
		}
	}
	r.l = l

	var e error
	for _, lr := range r.renderers {
		if err := r.call(lr, "init_option"); err != nil {
			e = appendError(e, err)
		}
	}
	if e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
	}
}

func (r *LuaRenderer) NewDocument() {
	var e error
	for _, lr := range r.renderers {
		if err := r.call(lr, "new_document"); err != nil {
			e = appendError(e, err)
		}
	}
	if e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
	}
}

func (r *LuaRenderer) call(lr *lua.LTable, method string, args ...lua.LValue) error {
	f := lr.RawGet(lua.LString(method))
	if f == lua.LNil {
		return nil
	}
	if fn, ok := f.(*lua.LFunction); ok {
		if err := r.l.CallByParam(lua.P{
			Fn:      fn,
			NRet:    1,
			Protect: true,
		}, append([]lua.LValue{lr}, args...)...); err != nil {
			return err
		}
		return nil
	}
	return errors.Errorf("Failed to call method: %s", method)
}

func (r *LuaRenderer) accept(lr *lua.LTable, n Node) bool {
	if err := r.call(lr, "accept", luar.New(r.l, n)); err != nil {
		return false
	}
	ret := r.l.Get(-1)
	r.l.Pop(1)
	return lua.LVAsBool(ret)
}

func (r *LuaRenderer) Accept(n Node) bool {
	for _, lr := range r.renderers {
		if r.accept(lr, n) {
			return true
		}
	}
	return false
}

func (r *LuaRenderer) RenderHeader(w io.Writer) error {
	var e error
	for _, lr := range r.renderers {
		if err := r.call(lr, "render_header", luar.New(r.l, w)); err != nil {
			e = appendError(e, err)
			continue
		}
		if s, ok := getResult(r.l).(lua.LString); ok {
			e = appendError(e, errors.New(string(s)))
		}
	}
	return e
}

func (r *LuaRenderer) Render(w io.Writer, node Node) error {
	var e error
	for _, lr := range r.renderers {
		if !r.accept(lr, node) {
			continue
		}
		if err := r.call(lr, "render", luar.New(r.l, w), luar.New(r.l, node)); err != nil {
			e = appendError(e, err)
			continue
		}
		if s, ok := getResult(r.l).(lua.LString); ok {
			if e == nil {
				e = errors.New(string(s))
			} else {
				e = multierr.Append(e, errors.New(string(s)))
			}
		}
	}
	return e
}

func (r *LuaRenderer) RenderFooter(w io.Writer) error {
	var e error
	for _, lr := range r.renderers {
		if err := r.call(lr, "render_footer", luar.New(r.l, w)); err != nil {
			e = appendError(e, err)
			continue
		}
		if s, ok := getResult(r.l).(lua.LString); ok {
			e = appendError(e, errors.New(string(s)))
		}
	}
	return e
}

func getResult(l *lua.LState) lua.LValue {
	ret := l.Get(-1)
	l.Pop(1)
	return ret
}

func appendError(e, other error) error {
	if e == nil {
		return other
	}
	return multierr.Append(e, other)
}
