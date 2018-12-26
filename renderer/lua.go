package renderer

import (
	"fmt"
	"io"
	"os"

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

type LuaRenderer struct {
	script   string
	l        *lua.LState
	renderer *lua.LTable
}

func NewLuaRenderer(script string) (Renderer, error) {
	r := &LuaRenderer{
		script:   script,
		l:        nil,
		renderer: nil,
	}
	l := lua.NewState()
	l.SetGlobal("tostring", l.NewFunction(toString))
	if err := l.DoFile(script); err != nil {
		return nil, err
	} else {
		r.renderer = getResult(l).(*lua.LTable)
	}
	r.l = l
	return r, nil
}

func (r *LuaRenderer) Name() string {
	if err := r.call("name"); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	if s, ok := getResult(r.l).(lua.LString); ok {
		return string(s)
	}
	return "lua"
}

func (r *LuaRenderer) SetOutputDirectory(path string) {
	if err := r.call("set_output_directory", luar.New(r.l, path)); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

func (r *LuaRenderer) SetFile(path string) {
	if err := r.call("set_file", luar.New(r.l, path)); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

func (r *LuaRenderer) AddOption() {
}

func (r *LuaRenderer) InitOption() {
}

func (r *LuaRenderer) NewDocument() {
	if err := r.call("new_document"); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

func (r *LuaRenderer) call(method string, args ...lua.LValue) error {
	f := r.renderer.RawGet(lua.LString(method))
	if f == lua.LNil {
		return nil
	}
	if fn, ok := f.(*lua.LFunction); ok {
		if err := r.l.CallByParam(lua.P{
			Fn:      fn,
			NRet:    1,
			Protect: true,
		}, append([]lua.LValue{r.renderer}, args...)...); err != nil {
			return err
		}
		return nil
	}
	return errors.Errorf("Failed to call method: %s", method)
}

func (r *LuaRenderer) Accept(n Node) bool {
	if err := r.call("accept", luar.New(r.l, n)); err != nil {
		return false
	}
	return lua.LVAsBool(getResult(r.l))
}

func (r *LuaRenderer) RenderHeader(w io.Writer) error {
	if err := r.call("header", luar.New(r.l, w)); err != nil {
		return err
	}
	if s, ok := getResult(r.l).(lua.LString); ok {
		return errors.New(string(s))
	}
	return nil
}

func (r *LuaRenderer) Render(w io.Writer, node Node) error {
	if err := r.call("render", luar.New(r.l, w), luar.New(r.l, node)); err != nil {
		return err
	}
	if s, ok := getResult(r.l).(lua.LString); ok {
		return errors.New(string(s))
	}
	return nil
}

func (r *LuaRenderer) RenderFooter(w io.Writer) error {
	if err := r.call("footer", luar.New(r.l, w)); err != nil {
		return err
	}
	if s, ok := getResult(r.l).(lua.LString); ok {
		return errors.New(string(s))
	}
	return nil
}
