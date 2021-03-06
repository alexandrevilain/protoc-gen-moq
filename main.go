package main

import (
	"flag"
	"path"
	"strings"

	"github.com/alexandrevilain/protoc-gen-moq/internal/forked/github.com/matryer/moq/template"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	contextPackage = protogen.GoImportPath("context")
	grpcPackage    = protogen.GoImportPath("google.golang.org/grpc")

	filePrefix *string
)

func main() {
	var flags flag.FlagSet
	filePrefix = flags.String("file_prefix", "", "Add prefix to the generated file (eg. /mock)")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			err := generateFile(gen, f, *filePrefix)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func rewriteFileNameWithPrefix(filepath, prefix string) string {
	return path.Join(path.Dir(filepath), prefix+path.Base(filepath))
}

func generateFile(gen *protogen.Plugin, file *protogen.File, filePrefix string) error {
	if len(file.Services) == 0 {
		return nil // skip file generation when no service found
	}

	filepath := file.GeneratedFilenamePrefix + "_moq.pb.go"
	if filePrefix != "" {
		filepath = rewriteFileNameWithPrefix(filepath, filePrefix)
	}

	importPath := file.GoImportPath
	if strings.HasSuffix(filePrefix, "/") {
		basePath := strings.Trim(importPath.String(), "\"")
		importPath = protogen.GoImportPath(path.Join(basePath, filePrefix))
	}

	g := gen.NewGeneratedFile(filepath, importPath)

	t, err := template.New()
	if err != nil {
		return err
	}

	mocks := make([]template.MockData, len(file.Services))
	for i, service := range file.Services {
		clientName := service.GoName + "Client"

		methods := make([]template.MethodData, len(service.Methods))
		for i, method := range service.Methods {
			// Generate params list
			params := []template.ParamData{
				{
					Name:     "ctx",
					Type:     g.QualifiedGoIdent(contextPackage.Ident("Context")),
					Variadic: false,
				},
			}

			if !method.Desc.IsStreamingClient() {
				params = append(params, template.ParamData{
					Name:     "in",
					Type:     g.QualifiedGoIdent(method.Input.GoIdent),
					Pointer:  true,
					Variadic: false,
				})
			}

			params = append(params, template.ParamData{
				Name:     "opts",
				Type:     g.QualifiedGoIdent(grpcPackage.Ident("CallOption")),
				Variadic: true,
			})

			// Generate return list
			var result template.ParamData

			if !method.Desc.IsStreamingClient() && !method.Desc.IsStreamingServer() {
				result = template.ParamData{
					Name:    "out",
					Type:    g.QualifiedGoIdent(method.Output.GoIdent),
					Pointer: true,
				}
			} else {
				result = template.ParamData{
					Name:    "out",
					Type:    method.Parent.GoName + "_" + method.GoName + "Client",
					Pointer: false,
				}
			}

			methods[i] = template.MethodData{
				Name:   method.GoName,
				Params: params,
				Returns: []template.ParamData{
					result,
					{
						Name: "err",
						Type: "error",
					},
				},
			}
		}

		mocks[i] = template.MockData{
			InterfaceName: g.QualifiedGoIdent(file.GoImportPath.Ident(clientName)),
			MockName:      clientName + "Mock",
			Methods:       methods,
		}
	}

	data := template.Data{
		PkgName:  string(file.GoPackageName),
		Mocks:    mocks,
		StubImpl: false,
		SyncPkg:  g.QualifiedGoIdent(protogen.GoImportPath("sync").Ident("")),
	}

	return t.Execute(g, data)
}
