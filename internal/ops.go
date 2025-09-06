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
		panic(fmt.Errorf("no numérico: %T", v))
	}
}

func toInt(v any) int {
	switch x := v.(type) {
	case int:
		return x
	case float64:
		return int(x)
	default:
		panic(fmt.Errorf("no entero: %T", v))
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
			panic(fmt.Errorf("índice fuera de rango"))
		}
		return c[i]
	case string:
		i := toInt(index)
		if i < 0 || i >= len(c) {
			panic(fmt.Errorf("índice fuera de rango"))
		}
		return string(c[i])
	default:
		panic(fmt.Errorf("subscript no soportado para %T", container))
	}
}

func applyStoreSubscr(container, index, value any) {
	switch c := container.(type) {
	case []any:
		i := toInt(index)
		if i < 0 || i >= len(c) {
			panic(fmt.Errorf("índice fuera de rango"))
		}
		c[i] = value
	default:
		panic(fmt.Errorf("STORE_SUBSCR no soportado para %T", container))
	}
}

// -------------------- Data / variables --------------------

func OpLoadConst(st *Stack[any], arg string) {
	val := parseConst(arg) // parseConst existe en vm.go
	st.Push(val)
}

func OpLoadFast(st *Stack[any], locals *VarList, name string) {
	// Evitamos GetVar por bug; usamos searchVar porque estamos en el mismo paquete
	v := locals.searchVar(name) // :contentReference[oaicite:4]{index=4}
	if v == nil {
		panic(fmt.Errorf("LOAD_FAST: variable '%s' no definida", name))
	}
	st.Push(v.Value)
}

func OpStoreFast(st *Stack[any], locals *VarList, name string) {
	val, err := st.Pop()
	if err != nil {
		panic(err)
	}
	// Si existe, set; si no, la creamos (más práctico que fallar)
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
		panic(fmt.Errorf("LOAD_GLOBAL: '%s' no definido", name))
	}
	st.Push(v.Value)
}

// -------------------- Aritmética / lógica --------------------

func OpBinary(st *Stack[any], kind string) {
	b, err := st.Pop()
	if err != nil {
		panic(err)
	}
	a, err := st.Pop()
	if err != nil {
		panic(err)
	}

	switch kind {
	case "BINARY_ADD":
		// suma numérica y concatenación string
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
			panic("división por cero")
		}
		st.Push(asFloat(a) / den)
	case "BINARY_MODULO":
		st.Push(float64(int(asFloat(a)) % int(asFloat(b))))

	default:
		panic("op binaria no soportada: " + kind)
	}
}

func OpLogical(st *Stack[any], kind string) {
	b, err := st.Pop()
	if err != nil {
		panic(err)
	}
	a, err := st.Pop()
	if err != nil {
		panic(err)
	}
	switch kind {
	case "BINARY_AND":
		st.Push(truthy(a) && truthy(b))
	case "BINARY_OR":
		st.Push(truthy(a) || truthy(b))
	default:
		panic("op lógica no soportada: " + kind)
	}
}

func OpCompare(st *Stack[any], op string) {
	b, err := st.Pop()
	if err != nil {
		panic(err)
	}
	a, err := st.Pop()
	if err != nil {
		panic(err)
	}
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
		panic("COMPARE_OP desconocido: " + op)
	}
}

// -------------------- Listas / subscript --------------------

func OpBuildList(st *Stack[any], arg string) {
	n := toInt(parseConst(arg))
	if st.Size() < n {
		panic("BUILD_LIST: elementos insuficientes")
	}
	out := make([]any, n)
	for i := n - 1; i >= 0; i-- {
		v, _ := st.Pop()
		out[i] = v
	}
	st.Push(out)
}

func OpBinarySubscr(st *Stack[any]) {
	idx, _ := st.Pop()
	cont, _ := st.Pop()
	st.Push(applySubscr(cont, idx))
}

func OpStoreSubscr(st *Stack[any]) {
	val, _ := st.Pop()
	idx, _ := st.Pop()
	cont, _ := st.Pop()
	applyStoreSubscr(cont, idx, val)
	st.Push(cont)
}

// -------------------- Control de flujo --------------------
// Devuelven (nuevoIP, jumped). Si jumped==false, ignorar retorno.

func OpJumpAbsolute(arg string) (int, bool) {
	return toInt(parseConst(arg)), true
}

func OpJumpIfTrue(st *Stack[any], arg string) (int, bool) {
	target := toInt(parseConst(arg))
	cond, _ := st.Pop()
	if truthy(cond) {
		return target, true
	}
	return -1, false
}

func OpJumpIfFalse(st *Stack[any], arg string) (int, bool) {
	target := toInt(parseConst(arg))
	cond, _ := st.Pop()
	if !truthy(cond) {
		return target, true
	}
	return -1, false
}

// -------------------- Call (placeholder) --------------------
// Depende de cómo representes funciones en "globals" (punteros a función, direcciones de código, etc.)
func OpCallFunction(_ *Stack[any], _ *VarList, _ string) {
	// para cuando llamen el print
	if name == "Print" {
		arg, err := st.Pop()
		if err != nil {
			panic(err)
		}
		fn, err := globals.GetVar("Print")
		if err != nil {
			panic(err)
		}
		if f, ok := fn.(func(any) any); ok {
			result := f(arg)
			st.Push(result)
			return
		}
		panic("Nombre " + name + " no es función")
	}
	// TODO: implementar cuando definas convención (nombre+arity, o salto a dirección).
	panic("CALL_FUNCTION: definir convención antes de implementar")
}
