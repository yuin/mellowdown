package renderer

import (
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/pkg/errors"
	lua "github.com/yuin/gopher-lua"
	"github.com/yuin/mellowdown/util"
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

func newLState() *lua.LState {
	l := lua.NewState()
	cfg := luar.GetConfig(l)
	cfg.FieldNames = func(s reflect.Type, f reflect.StructField) []string {
		return []string{util.ToSnakeCase(f.Name)}
	}
	cfg.MethodNames = func(t reflect.Type, m reflect.Method) []string {
		return []string{util.ToSnakeCase(m.Name)}
	}
	l.SetGlobal("tostring", l.NewFunction(toString))
	l.SetGlobal("any", lua.LString("*"))
	l.SetGlobal("node_fenced_code", lua.LNumber(NodeFencedCode))
	l.SetGlobal("node_function", lua.LNumber(NodeFunction))
	l.SetGlobal("read_meta", luar.New(l, readMeta))
	return l
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
	l := newLState()
	if err := l.DoFile(script); err != nil {
		return nil, err
	} else {
		r.renderer = getResult(l).(*lua.LTable)
	}
	r.l = l
	return r, nil
}

func (r *LuaRenderer) Name() string {
	if err := r.call("name", 1); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	if s, ok := getResult(r.l).(lua.LString); ok {
		return string(s)
	}
	return "lua"
}

func (r *LuaRenderer) AddOption(o Option) {
}

func (r *LuaRenderer) InitOption(o Option) {
}

func (r *LuaRenderer) NewDocument(c RenderingContext) {
	if err := r.call("new_document", 1, luar.New(r.l, c)); err != nil {
		panic(err)
	}
}

func (r *LuaRenderer) call(method string, nret int, args ...lua.LValue) error {
	f := r.renderer.RawGet(lua.LString(method))
	if f == lua.LNil {
		return nil
	}
	if fn, ok := f.(*lua.LFunction); ok {
		if err := r.l.CallByParam(lua.P{
			Fn:      fn,
			NRet:    nret,
			Protect: true,
		}, append([]lua.LValue{r.renderer}, args...)...); err != nil {
			return err
		}
		return nil
	}
	return errors.Errorf("Failed to call method: %s", method)
}

func (r *LuaRenderer) Acceptable() (NodeType, string) {
	if err := r.call("acceptable", 2); err != nil {
		panic(err)
	}
	name := string(getResult(r.l).(lua.LString))
	t := NodeType(int(getResult(r.l).(lua.LNumber)))
	return t, name
}

func (r *LuaRenderer) RenderHeader(w io.Writer, c RenderingContext) error {
	if err := r.call("header", 1, luar.New(r.l, w), luar.New(r.l, c)); err != nil {
		return err
	}

	if s, ok := getResult(r.l).(lua.LString); ok {
		return errors.New(string(s))
	}
	return nil
}

func (r *LuaRenderer) Render(w io.Writer, node Node, c RenderingContext) error {
	if err := r.call("render", 1, luar.New(r.l, w), luar.New(r.l, node), luar.New(r.l, c)); err != nil {
		return err
	}
	if s, ok := getResult(r.l).(lua.LString); ok {
		return errors.New(string(s))
	}
	return nil
}

func (r *LuaRenderer) RenderFooter(w io.Writer, c RenderingContext) error {
	if err := r.call("footer", 1, luar.New(r.l, w), luar.New(r.l, c)); err != nil {
		return err
	}
	if s, ok := getResult(r.l).(lua.LString); ok {
		return errors.New(string(s))
	}
	return nil
}
