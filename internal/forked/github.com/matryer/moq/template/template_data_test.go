package template_test

import (
	"testing"

	"github.com/alexandrevilain/protoc-gen-moq/internal/forked/github.com/matryer/moq/template"
	"github.com/stretchr/testify/assert"
)

func TestParamData(t *testing.T) {
	testCases := []struct {
		desc	string
		paramData template.ParamData
		expectedCallName string
		expectedTypeString string
		expectedMethodArg string
	}{
		{
			desc: "Simple variable type",
			paramData: template.ParamData{
				Name: "ctx",
				Type: "context.Context",
				Pointer: false,
				Variadic: false,
			},
			expectedCallName: "ctx",
			expectedTypeString: "context.Context",
			expectedMethodArg: "ctx context.Context",
		},
		{
			desc: "Variadic variable",
			paramData: template.ParamData{
				Name: "opts",
				Type: "grpc.CallOption",
				Pointer: false,
				Variadic: true,
			},
			expectedMethodArg: "opts ...grpc.CallOption",
			expectedTypeString: "[]grpc.CallOption",
			expectedCallName: "opts...",
		},
		{
			desc: "Pointer variable",
			paramData: template.ParamData{
				Name: "opts",
				Type: "grpc.CallOption",
				Pointer: true,
				Variadic: false,
			},
			expectedMethodArg: "opts *grpc.CallOption",
			expectedTypeString: "*grpc.CallOption",
			expectedCallName: "opts",
		},
		{
			desc: "Variadic pointer variable",
			paramData: template.ParamData{
				Name: "opts",
				Type: "grpc.CallOption",
				Pointer: true,
				Variadic: true,
			},
			expectedMethodArg: "opts ...*grpc.CallOption",
			expectedTypeString: "[]*grpc.CallOption",
			expectedCallName: "opts...",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(tt *testing.T) {
			assert.Equal(tt, tC.expectedCallName, tC.paramData.CallName())
			assert.Equal(tt, tC.expectedTypeString, tC.paramData.TypeString())
			assert.Equal(tt, tC.expectedMethodArg, tC.paramData.MethodArg())
		})
	}
}

func TestMethodData(t *testing.T) {
	testCases := []struct {
		desc	string
		methodData template.MethodData
		expectedArgList string
		expectedArgCallList string
		expectedReturnArgNameList string
		expectedReturnArgTypeList string
	}{
		{
			desc: "Single param with no return values",
			methodData: template.MethodData{
				Name: "DoSomething",
				Params: []template.ParamData{
					{
						Name: "ctx",
						Type: "context.Context",
						Pointer: false,
						Variadic: false,
					},
				},
				Returns: nil,
			},
			expectedArgList: "ctx context.Context",
			expectedArgCallList: "ctx",
			expectedReturnArgNameList: "",
			expectedReturnArgTypeList: "",
			
		},
		{
			desc: "Single param with error returned",
			methodData: template.MethodData{
				Name: "DoSomething",
				Params: []template.ParamData{
					{
						Name: "ctx",
						Type: "context.Context",
						Pointer: false,
						Variadic: false,
					},
				},
				Returns: []template.ParamData{
					{
						Name: "err",
						Type: "error",
					},
				},
			},
			expectedArgList: "ctx context.Context",
			expectedArgCallList: "ctx",
			expectedReturnArgNameList: "err",
			expectedReturnArgTypeList: "error",
		},
		{
			desc: "Multiple params with one pointer with error returned",
			methodData: template.MethodData{
				Name: "DoSomething",
				Params: []template.ParamData{
					{
						Name: "ctx",
						Type: "context.Context",
						Pointer: false,
						Variadic: false,
					},
					{
						Name: "in",
						Type: "hellov1.SayHelloRequest",
						Pointer: true,
						Variadic: false,
					},
				},
				Returns: []template.ParamData{
					{
						Name: "err",
						Type: "error",
					},
				},
			},
			expectedArgList: "ctx context.Context, in *hellov1.SayHelloRequest",
			expectedArgCallList: "ctx, in",
			expectedReturnArgNameList: "err",
			expectedReturnArgTypeList: "error",
		},
		{
			desc: "Multiple params with one pointer with pointer and error returned",
			methodData: template.MethodData{
				Name: "DoSomething",
				Params: []template.ParamData{
					{
						Name: "ctx",
						Type: "context.Context",
						Pointer: false,
						Variadic: false,
					},
					{
						Name: "in",
						Type: "hellov1.SayHelloRequest",
						Pointer: true,
						Variadic: false,
					},
				},
				Returns: []template.ParamData{
					{
						Name: "resp",
						Type: "hellov1.SayHelloResponse",
						Pointer: true,
						Variadic: false,
					},
					{
						Name: "err",
						Type: "error",
					},
				},
			},
			expectedArgList: "ctx context.Context, in *hellov1.SayHelloRequest",
			expectedArgCallList: "ctx, in",
			expectedReturnArgNameList: "resp, err",
			expectedReturnArgTypeList: "(*hellov1.SayHelloResponse, error)",
		},
		{
			desc: "Multiple params with one variadic with pointer and error returned",
			methodData: template.MethodData{
				Name: "DoSomething",
				Params: []template.ParamData{
					{
						Name: "ctx",
						Type: "context.Context",
						Pointer: false,
						Variadic: false,
					},
					{
						Name: "opts",
						Type: "grpc.CallOption",
						Pointer: false,
						Variadic: true,
					},
				},
				Returns: []template.ParamData{
					{
						Name: "resp",
						Type: "hellov1.SayHelloResponse",
						Pointer: true,
						Variadic: false,
					},
					{
						Name: "err",
						Type: "error",
					},
				},
			},
			expectedArgList: "ctx context.Context, opts ...grpc.CallOption",
			expectedArgCallList: "ctx, opts...",
			expectedReturnArgNameList: "resp, err",
			expectedReturnArgTypeList: "(*hellov1.SayHelloResponse, error)",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(tt *testing.T) {
			assert.Equal(tt, tC.expectedArgList, tC.methodData.ArgList())
			assert.Equal(tt, tC.expectedArgCallList, tC.methodData.ArgCallList())
			assert.Equal(tt, tC.expectedReturnArgNameList, tC.methodData.ReturnArgNameList())
			assert.Equal(tt, tC.expectedReturnArgTypeList, tC.methodData.ReturnArgTypeList())
		})
	}
}