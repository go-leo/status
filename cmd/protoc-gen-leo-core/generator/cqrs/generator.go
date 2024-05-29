package cqrs

import (
	"github.com/go-leo/gox/stringx"
	"github.com/go-leo/leo/v3/cmd/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"os"
	"path"
	"strconv"
)

type Generator struct {
	Plugin   *protogen.Plugin
	File     *protogen.File
	Services []*internal.Service
}

func NewGenerator(plugin *protogen.Plugin, file *protogen.File) (*Generator, error) {
	services, err := internal.NewCQRSServices(file)
	if err != nil {
		return nil, err
	}
	return &Generator{Plugin: plugin, File: file, Services: services}, nil
}

func (f *Generator) Generate(g *protogen.GeneratedFile) error {
	for _, service := range f.Services {
		if err := f.GenerateEndpoints(service); err != nil {
			return err
		}
	}
	for _, service := range f.Services {
		if err := f.GenerateAssembler(service, g); err != nil {
			return err
		}
	}
	for _, service := range f.Services {
		if err := f.GenerateCQRSService(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := f.GenerateBus(service, g); err != nil {
			return err
		}
	}
	return nil
}

func (f *Generator) GenerateEndpoints(service *internal.Service) error {
	for _, endpoint := range service.Endpoints {
		if err := f.GenerateEndpoint(service, endpoint); err != nil {
			return err
		}
	}
	return nil
}

func (f *Generator) GenerateEndpoint(service *internal.Service, endpoint *internal.Endpoint) error {
	if endpoint.IsStreaming() {
		return nil
	}
	switch {
	case endpoint.IsCommand():
		return f.GenerateCommand(service, endpoint)
	case endpoint.IsQuery():
		return f.GenerateQuery(service, endpoint)
	}
	return nil
}

func (f *Generator) GenerateCommand(service *internal.Service, endpoint *internal.Endpoint) error {
	filename := path.Join(service.Command.RelPath(), stringx.SnackCase(endpoint.Name())+".go")
	_, err := os.Stat(filename)
	if err == nil {
		return nil
	}
	if !os.IsNotExist(err) {
		return err
	}
	generatedFile := f.Plugin.NewGeneratedFile(filename, f.File.GoImportPath)
	generatedFile.P("package ", service.Command.Name())
	generatedFile.P("type ", endpoint.ArgsName(), " struct {")
	generatedFile.P("}")
	generatedFile.P()
	generatedFile.P("type ", endpoint.Name(), " ", internal.CqrsPackage.Ident("CommandHandler"), "[*", endpoint.ArgsName(), "]")
	generatedFile.P()
	generatedFile.P("func New", endpoint.Name(), "() ", endpoint.Name(), " {")
	generatedFile.P("return &", endpoint.UnexportedName(), "{}")
	generatedFile.P("}")
	generatedFile.P()
	generatedFile.P("type ", endpoint.UnexportedName(), " struct {")
	generatedFile.P("}")
	generatedFile.P()
	generatedFile.P("func (h *", endpoint.UnexportedName(), ") Handle(ctx ", internal.ContextPackage.Ident("Context"), ", args *", endpoint.ArgsName(), ") (", internal.MetadataxPackage.Ident("Metadata"), ", error) {")
	generatedFile.P(internal.Comments("TODO implement me"))
	generatedFile.P("panic(", strconv.Quote("implement me"), ")")
	generatedFile.P("}")
	return nil
}

func (f *Generator) GenerateQuery(service *internal.Service, endpoint *internal.Endpoint) error {
	filename := path.Join(service.Query.RelPath(), stringx.SnackCase(endpoint.Name())+".go")
	_, err := os.Stat(filename)
	if err == nil {
		return nil
	}
	if !os.IsNotExist(err) {
		return err
	}
	generatedFile := f.Plugin.NewGeneratedFile(filename, f.File.GoImportPath)
	generatedFile.P("package ", service.Query.Name())
	generatedFile.P("type ", endpoint.ArgsName(), " struct {")
	generatedFile.P("}")
	generatedFile.P()
	generatedFile.P("type ", endpoint.ResName(), " struct {")
	generatedFile.P("}")
	generatedFile.P()
	generatedFile.P("type ", endpoint.Name(), " ", internal.CqrsPackage.Ident("QueryHandler"), "[*", endpoint.ArgsName(), ", *", endpoint.ResName(), "]")
	generatedFile.P()
	generatedFile.P("func New", endpoint.Name(), "() ", endpoint.Name(), " {")
	generatedFile.P("return &", endpoint.UnexportedName(), "{}")
	generatedFile.P("}")
	generatedFile.P()
	generatedFile.P("type ", endpoint.UnexportedName(), " struct {")
	generatedFile.P("}")
	generatedFile.P()
	generatedFile.P("func (h *", endpoint.UnexportedName(), ") Handle(ctx ", internal.ContextPackage.Ident("Context"), ", args *", endpoint.ArgsName(), ") (*", endpoint.ResName(), ", error) {")
	generatedFile.P(internal.Comments("TODO implement me"))
	generatedFile.P("panic(", strconv.Quote("implement me"), ")")
	generatedFile.P("}")
	return nil
}

func (f *Generator) GenerateAssembler(service *internal.Service, generatedFile *protogen.GeneratedFile) error {
	generatedFile.P(internal.Comments(service.AssemblerName() + " responsible for completing the transformation between domain model objects and DTOs"))
	generatedFile.P("type ", service.AssemblerName(), " interface {")
	for _, endpoint := range service.Endpoints {
		if endpoint.IsStreaming() {
			continue
		}
		switch {
		case endpoint.IsCommand():
			generatedFile.P()
			generatedFile.P(internal.Comments("From" + endpoint.RequestName() + " convert request to command arguments"))
			generatedFile.P("From", endpoint.RequestName(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", protogen.GoImportPath(service.Command.FullName()).Ident(endpoint.ArgsName()), ", ", internal.ContextPackage.Ident("Context"), ", error)")
			generatedFile.P()
			generatedFile.P(internal.Comments("To" + endpoint.ResponseName() + " convert query result to response"))
			generatedFile.P("To", endpoint.ResponseName(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ", metadata ", internal.MetadataxPackage.Ident("Metadata"), ") (*", endpoint.OutputGoIdent(), ", error)")
		case endpoint.IsQuery():
			generatedFile.P()
			generatedFile.P(internal.Comments("From" + endpoint.RequestName() + " convert request to query arguments"))
			generatedFile.P("From", endpoint.RequestName(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", protogen.GoImportPath(service.Query.FullName()).Ident(endpoint.ArgsName()), ", ", internal.ContextPackage.Ident("Context"), ", error)")
			generatedFile.P()
			generatedFile.P(internal.Comments("To" + endpoint.ResponseName() + " convert query result to response"))
			generatedFile.P("To", endpoint.ResponseName(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ", res *", protogen.GoImportPath(service.Query.FullName()).Ident(endpoint.ResName()), ") (*", endpoint.OutputGoIdent(), ", error)")
		}
	}
	generatedFile.P("}")
	return nil
}

func (f *Generator) GenerateCQRSService(service *internal.Service, generatedFile *protogen.GeneratedFile) error {
	generatedFile.P(internal.Comments(service.UnexportedCQRSName() + " implement the " + service.ServiceName() + " with CQRS pattern"))
	generatedFile.P("type ", service.UnexportedCQRSName(), " struct {")
	generatedFile.P("bus       ", internal.CqrsPackage.Ident("Bus"))
	generatedFile.P("assembler ", service.AssemblerName())
	generatedFile.P("}")
	generatedFile.P()

	for _, endpoint := range service.Endpoints {
		switch {
		case endpoint.IsCommand():
			generatedFile.P("func (svc *", service.UnexportedCQRSName(), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error){")
			generatedFile.P("args, ctx, err := svc.assembler.From", endpoint.Name(), "Request(ctx, request)")
			generatedFile.P("if err != nil {")
			generatedFile.P("return nil, err")
			generatedFile.P("}")
			generatedFile.P("metadata, err := svc.bus.Exec(ctx, args)")
			generatedFile.P("if err != nil {")
			generatedFile.P("return nil, err")
			generatedFile.P("}")
			generatedFile.P("return svc.assembler.To", endpoint.Name(), "Response(ctx, request, metadata)")
			generatedFile.P("}")
			generatedFile.P()
		case endpoint.IsQuery():
			generatedFile.P("func (svc *", service.UnexportedCQRSName(), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error){")
			generatedFile.P("args, ctx, err := svc.assembler.From", endpoint.Name(), "Request(ctx, request)")
			generatedFile.P("if err != nil {")
			generatedFile.P("return nil, err")
			generatedFile.P("}")
			generatedFile.P("res, err := svc.bus.Query(ctx, args)")
			generatedFile.P("if err != nil {")
			generatedFile.P("return nil, err")
			generatedFile.P("}")
			generatedFile.P("return svc.assembler.To", endpoint.Name(), "Response(ctx, request, res.(*", protogen.GoImportPath(service.Query.FullName()).Ident(endpoint.Name()+"Res"), "))")
			generatedFile.P("}")
			generatedFile.P()
		}
	}

	generatedFile.P("func New", service.CQRSName(), "(bus ", internal.CqrsPackage.Ident("Bus"), ", assembler ", service.AssemblerName(), ") ", service.ServiceName(), " {")
	generatedFile.P("return &", service.UnexportedCQRSName(), "{bus: bus, assembler: assembler}")
	generatedFile.P("}")
	generatedFile.P()
	return nil
}

func (f *Generator) GenerateBus(service *internal.Service, generatedFile *protogen.GeneratedFile) error {
	generatedFile.P("func New", service.BusName(), "(")
	for _, endpoint := range service.Endpoints {
		if endpoint.IsStreaming() {
			continue
		}
		switch {
		case endpoint.IsCommand():
			importPath := protogen.GoImportPath(service.Command.FullName())
			generatedFile.P(endpoint.UnexportedName(), " ", importPath.Ident(endpoint.Name()), ",")
		case endpoint.IsQuery():
			importPath := protogen.GoImportPath(service.Query.FullName())
			generatedFile.P(endpoint.UnexportedName(), " ", importPath.Ident(endpoint.Name()), ",")
		}
	}
	generatedFile.P(") (", internal.CqrsPackage.Ident("Bus"), ", error) {")
	generatedFile.P("bus := ", internal.CqrsPackage.Ident("NewBus"), "()")
	for _, endpoint := range service.Endpoints {
		if endpoint.IsStreaming() {
			continue
		}
		switch {
		case endpoint.IsCommand():
			generatedFile.P("if err := bus.RegisterCommand(", endpoint.UnexportedName(), "); err != nil {")
			generatedFile.P("return nil, err")
			generatedFile.P("}")
		case endpoint.IsQuery():
			generatedFile.P("if err := bus.RegisterQuery(", endpoint.UnexportedName(), "); err != nil {")
			generatedFile.P("return nil, err")
			generatedFile.P("}")
		}
	}
	generatedFile.P("return bus, nil")
	generatedFile.P("}")
	return nil
}
