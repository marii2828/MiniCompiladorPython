package internal

import "fmt"

// -------------------- Helpers genéricos --------------------

func asFloat(v any) float64 {
	switch x := v.(type) {
	case int:
		return float64(x)
	case float64:
		return x
	default:
		panic(fmt.Errorf("Not a number: %T", v))
	}
}

func toInt(v any) int {
	switch x := v.(type) {
	case int:
		return x
	case float64:
		return int(x)
	default:
		panic(fmt.Errorf("Not an integer: %T", v))
	}
}

func truthy(v any) bool {
	switch t := v.(type) {
	case bool:
		return t
	case int:
		return t != 0
	case float64:
		return t != 0
	case string:
		return t != ""
	case []any:
		return len(t) != 0
	default:
		return v != nil
	}
}

func applySubscr(container, index any) any {
	switch c := container.(type) {
	case []any:
		i := toInt(index)
		if i < 0 || i >= len(c) {
			panic(fmt.Errorf("Index out of range (len=%d, i=%d)", len(c), i))
		}
		return c[i]
	case string:
		i := toInt(index)
		if i < 0 || i >= len(c) {
			panic(fmt.Errorf("Index out of range (len=%d, i=%d)", len(c), i))
		}
		return string(c[i])
	default:
		panic(fmt.Errorf("subscript not supported for %T", container))
	}
}

func applyStoreSubscr(container, index, value any) {
	switch c := container.(type) {
	case []any:
		i := toInt(index)
		if i < 0 || i >= len(c) {
			panic(fmt.Errorf("STORE_SUBSCR: index out of range (len=%d, i=%d)", len(c), i))
		}
		c[i] = value
	default:
		panic(fmt.Errorf("STORE_SUBSCR not supported for %T", container))
	}
}

// -------------------- Data / variables --------------------

func OpLoadConst(st *Stack[any], arg string) {
	val := parseConst(arg) // parseConst existe en vm.go
	st.Push(val)
}

func OpLoadFast(st *Stack[any], locals *VarList, name string) {
	v := locals.searchVar(name)
	if v == nil {
		panic(fmt.Errorf("LOAD_FAST: variable '%s' not defined", name))
	}
	st.Push(v.Value)
}

func OpStoreFast(st *Stack[any], locals *VarList, name string) {
	if err := st.Ensure(1); err != nil {
		panic(fmt.Errorf("STORE_FAST: %w", err))
	}
	val, _ := st.Pop()
	if v := locals.searchVar(name); v != nil {
		v.Value = val
		return
	}
	if err := locals.AddVar(name, val); err != nil {
		panic(err)
	}
}

func OpLoadGlobal(st *Stack[any], globals *VarList, name string) {
	v := globals.searchVar(name)
	if v == nil {
		panic(fmt.Errorf("LOAD_GLOBAL: '%s' not defined", name))
	}
	st.Push(v.Value)
}

// -------------------- Aritmética / lógica --------------------

func OpBinary(st *Stack[any], kind string) {
	if err := st.Ensure(2); err != nil {
		panic(fmt.Errorf("%s: %w", kind, err))
	}
	b, _ := st.Pop()
	a, _ := st.Pop()

	switch kind {
	case "BINARY_ADD":
		if sa, ok := a.(string); ok {
			st.Push(sa + fmt.Sprintf("%v", b))
			return
		}
		st.Push(asFloat(a) + asFloat(b))
	case "BINARY_SUBSTRACT":
		st.Push(asFloat(a) - asFloat(b))
	case "BINARY_MULTIPLY":
		st.Push(asFloat(a) * asFloat(b))
	case "BINARY_DIVIDE":
		den := asFloat(b)
		if den == 0 {
			panic("division by zero")
		}
		st.Push(asFloat(a) / den)
	case "BINARY_MODULO":
		st.Push(float64(int(asFloat(a)) % int(asFloat(b))))
	default:
		panic("opcode not supported: " + kind)
	}
}

func OpLogical(st *Stack[any], kind string) {
	if err := st.Ensure(2); err != nil {
		panic(fmt.Errorf("%s: %w", kind, err))
	}
	b, _ := st.Pop()
	a, _ := st.Pop()
	switch kind {
	case "BINARY_AND":
		st.Push(truthy(a) && truthy(b))
	case "BINARY_OR":
		st.Push(truthy(a) || truthy(b))
	default:
		panic("logical op not supported: " + kind)
	}
}

func OpCompare(st *Stack[any], op string) {
	if err := st.Ensure(2); err != nil {
		panic(fmt.Errorf("COMPARE_OP: %w", err))
	}
	b, _ := st.Pop()
	a, _ := st.Pop()
	switch op {
	case "==":
		st.Push(fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b))
	case "!=":
		st.Push(fmt.Sprintf("%v", a) != fmt.Sprintf("%v", b))
	case "<":
		st.Push(asFloat(a) < asFloat(b))
	case ">":
		st.Push(asFloat(a) > asFloat(b))
	case "<=":
		st.Push(asFloat(a) <= asFloat(b))
	case ">=":
		st.Push(asFloat(a) >= asFloat(b))
	default:
		panic("COMPARE_OP not supported: " + op)
	}
}

// -------------------- Listas / subscript --------------------

func OpBuildList(st *Stack[any], arg string) {
	n := toInt(parseConst(arg))
	if err := st.Ensure(n); err != nil {
		panic(fmt.Errorf("BUILD_LIST: need %d elements, have %d", n, st.Size()))
	}
	items, _ := st.PopN(n)
	st.Push(items)
}

func OpBinarySubscr(st *Stack[any]) {
	if err := st.Ensure(2); err != nil {
		panic(fmt.Errorf("BINARY_SUBSCR: %w", err))
	}
	idx, _ := st.Pop()
	cont, _ := st.Pop()
	st.Push(applySubscr(cont, idx))
}

func OpStoreSubscr(st *Stack[any]) {
	if err := st.Ensure(3); err != nil {
		panic(fmt.Errorf("STORE_SUBSCR: %w", err))
	}
	val, _ := st.Pop()
	idx, _ := st.Pop()
	cont, _ := st.Pop()
	applyStoreSubscr(cont, idx, val)
	st.Push(cont)
}

// -------------------- Control de flujo --------------------

func OpJumpAbsolute(arg string) (int, bool) {
	return toInt(parseConst(arg)), true
}

func OpJumpIfTrue(st *Stack[any], arg string) (int, bool) {
	if err := st.Ensure(1); err != nil {
		panic(fmt.Errorf("JUMP_IF_TRUE: %w", err))
	}
	target := toInt(parseConst(arg))
	cond, _ := st.Pop()
	if truthy(cond) {
		return target, true
	}
	return -1, false
}

func OpJumpIfFalse(st *Stack[any], arg string) (int, bool) {
	if err := st.Ensure(1); err != nil {
		panic(fmt.Errorf("JUMP_IF_FALSE: %w", err))
	}
	target := toInt(parseConst(arg))
	cond, _ := st.Pop()
	if !truthy(cond) {
		return target, true
	}
	return -1, false
}

// -------------------- Call --------------------

func OpCallFunction(st *Stack[any], globals *VarList, arg string) {
	var nargs int
	if arg != "" {
		nargs = toInt(parseConst(arg)) // usa argumento textual si existe
	} else {
		if err := st.Ensure(1); err != nil {
			panic(fmt.Errorf("CALL_FUNCTION: %w (missing nargs)", err))
		}
		nargsAny, _ := st.Pop()
		nargs = toInt(nargsAny)
	}

	if err := st.Ensure(nargs + 1); err != nil {
		panic(fmt.Errorf("CALL_FUNCTION: need %d items (args+fn), have %d", nargs+1, st.Size()))
	}

	args, _ := st.PopN(nargs)
	fnAny, _ := st.Pop()

	switch fn := fnAny.(type) {
	case func(...interface{}):
		fn(args...)
	case func(...interface{}) error:
		_ = fn(args...)
	case func(...interface{}) (int, error):
		_, _ = fn(args...)
	case func(...interface{}) int:
		_ = fn(args...)
	default:
		panic(fmt.Errorf("CALL_FUNCTION: not a function, got %T", fnAny))
	}
}
