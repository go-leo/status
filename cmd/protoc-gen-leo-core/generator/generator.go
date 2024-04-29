package generator

import (
	"github.com/go-leo/leo/v3/cmd/internal"
	"google.golang.org/protobuf/compiler/protogen"
)

type Generator struct {
	Plugin   *protogen.Plugin
	File     *protogen.File
	Services []*internal.Service
}

func NewGenerator(plugin *protogen.Plugin, file *protogen.File) (*Generator, error) {
	services, err := internal.NewServices(file)
	if err != nil {
		return nil, err
	}
	return &Generator{Plugin: plugin, File: file, Services: services}, nil
}

func (f *Generator) Generate() error {
	return f.GenerateFile()
}

func (f *Generator) GenerateFile() error {
	file := f.File
	filename := file.GeneratedFilenamePrefix + "_core.leo.pb.go"
	g := f.Plugin.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-grpc. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	for _, service := range f.Services {
		if err := f.GenerateEndpoints(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := f.GenerateNewEndpoints(service, g); err != nil {
			return err
		}
	}
	return nil
}

func (f *Generator) GenerateEndpoints(service *internal.Service, generatedFile *protogen.GeneratedFile) error {
	generatedFile.P("type ", service.UnexportedEndpointsName(), " struct {")
	generatedFile.P()
	generatedFile.P("svc interface {")
	for _, endpoint := range service.Endpoints {
		generatedFile.P(endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error)")
	}
	generatedFile.P("}")
	generatedFile.P("}")
	generatedFile.P()

	for _, endpoint := range service.Endpoints {
		generatedFile.P("func (e *", service.UnexportedEndpointsName(), ") ", endpoint.Name(), "() ", internal.EndpointPackage.Ident("Endpoint"), "{")
		generatedFile.P("return func(ctx ", internal.ContextPackage.Ident("Context"), ", request any) (any, error) {")
		generatedFile.P("return e.svc.", endpoint.Name(), "(ctx, request.(*", endpoint.InputGoIdent(), "))")
		generatedFile.P("}")
		generatedFile.P("}")
		generatedFile.P()
	}
	return nil
}

func (f *Generator) GenerateNewEndpoints(service *internal.Service, generatedFile *protogen.GeneratedFile) error {
	generatedFile.P("func New", service.EndpointsName(), "(")
	generatedFile.P("svc interface {")
	for _, endpoint := range service.Endpoints {
		generatedFile.P(endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error)")
	}
	generatedFile.P("},")
	generatedFile.P(") interface {")
	for _, endpoint := range service.Endpoints {
		generatedFile.P(endpoint.Name(), "() ", internal.EndpointPackage.Ident("Endpoint"))
	}
	generatedFile.P("} {")
	generatedFile.P("return &", service.UnexportedEndpointsName(), "{svc: svc}")
	generatedFile.P("}")
	generatedFile.P()
	return nil
}
