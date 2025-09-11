package internal

import "fmt"

type Variable struct {
	Name  string
	Value any
}

type VarList struct {
	Vars []Variable
}

var LocalVarListVM *VarList = NewVarList()

var GlobalVarListVM *VarList = NewVarList()

func NewVarList() *VarList {
	return &VarList{Vars: []Variable{}}
}

func PrintVars(vl *VarList) {
	fmt.Println("\nCurrent Variables:")
	for _, v := range vl.Vars {
		fmt.Printf("Name: %s, Value: %v\n", v.Name, v.Value)
	}
}

func (vl *VarList) searchVar(name string) *Variable {
	for i := range vl.Vars {
		if vl.Vars[i].Name == name {
			return &vl.Vars[i]
		}
	}
	return nil
}

func (vl *VarList) SetVar(name string, value any) error {
	//serch if the variable already exists
	variable := vl.searchVar(name)
	if variable != nil {
		variable.Value = value
		return nil
	}
	return fmt.Errorf("Variable not found")
}

func (vl *VarList) GetVar(name string) (any, error) {
	variable := vl.searchVar(name)
	if variable != nil {
		return variable.Value, nil
	}
	return nil, fmt.Errorf("Variable not found")
}

func (vl *VarList) AddVar(name string, value any) error {
	//serch if the variable already exists
	variable := vl.searchVar(name)
	if variable != nil {
		return fmt.Errorf("Variable already exists")
	}
	vl.Vars = append(vl.Vars, Variable{Name: name, Value: value})
	return nil
}

func (vl *VarList) DeleteVar(name string) error {
	for i := range vl.Vars {
		if vl.Vars[i].Name == name {
			vl.Vars = append(vl.Vars[:i], vl.Vars[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Variable not found")
}

func init() {
	GlobalVarListVM.AddVar("print", fmt.Println)
}
