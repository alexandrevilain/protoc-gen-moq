package template

import (
	"fmt"
	"strings"
)

// Data is the template data used to render the Moq template.
type Data struct {
	PkgName  string
	Mocks    []MockData
	StubImpl bool
	SyncPkg  string
}

// MockData is the data used to generate a mock for some interface.
type MockData struct {
	InterfaceName string
	MockName      string
	Methods       []MethodData
}

// MethodData is the data which represents a method on some interface.
type MethodData struct {
	Name    string
	Params  []ParamData
	Returns []ParamData
}

// ArgList is the string representation of method parameters, ex:
// 's string, n int, foo bar.Baz'.
func (m MethodData) ArgList() string {
	params := make([]string, len(m.Params))
	for i, p := range m.Params {
		params[i] = p.MethodArg()
	}
	return strings.Join(params, ", ")
}

// ArgCallList is the string representation of method call parameters,
// ex: 's, n, foo'. In case of a last variadic parameter, it will be of
// the format 's, n, foos...'
func (m MethodData) ArgCallList() string {
	params := make([]string, len(m.Params))
	for i, p := range m.Params {
		params[i] = p.CallName()
	}
	return strings.Join(params, ", ")
}

// ReturnArgTypeList is the string representation of method return
// types, ex: 'bar.Baz', '(string, error)'.
func (m MethodData) ReturnArgTypeList() string {
	params := make([]string, len(m.Returns))
	for i, p := range m.Returns {
		params[i] = p.TypeString()
	}
	if len(m.Returns) > 1 {
		return fmt.Sprintf("(%s)", strings.Join(params, ", "))
	}
	return strings.Join(params, ", ")
}

// ReturnArgNameList is the string representation of values being
// returned from the method, ex: 'foo', 's, err'.
func (m MethodData) ReturnArgNameList() string {
	params := make([]string, len(m.Returns))
	for i, p := range m.Returns {
		params[i] = p.Name
	}
	return strings.Join(params, ", ")
}

// ParamData is the data which represents a parameter to some method of
// an interface.
type ParamData struct {
	Name     string
	Type     string
	Pointer  bool
	Variadic bool
}

// MethodArg is the representation of the parameter in the function
// signature, ex: 'name a.Type'.
func (p ParamData) MethodArg() string {
	if p.Variadic {
		return fmt.Sprintf("%s ...%s", p.Name, p.TypeString()[2:])
	}
	return fmt.Sprintf("%s %s", p.Name, p.TypeString())
}

// CallName returns the string representation of the parameter to be
// used for a method call. For a variadic paramter, it will be of the
// format 'foos...'.
func (p ParamData) CallName() string {
	if p.Variadic {
		return p.Name + "..."
	}
	return p.Name
}

// TypeString returns the variable type with the package qualifier in the
// format 'pkg.Type'.
func (p ParamData) TypeString() string {
	result := ""
	if p.Variadic {
		result += "[]"
	}
	if p.Pointer {
		result += "*"
	}
	result += p.Type
	return result
}
