package generator

import (
	"fmt"
	"github.com/go-leo/leo/v3/cmd/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strconv"
	"sync"
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
	filename := file.GeneratedFilenamePrefix + "_http.leo.pb.go"
	g := f.Plugin.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-grpc. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	for _, service := range f.Services {
		if err := f.GenerateNewServer(service, g); err != nil {
			return err
		}
	}
	for _, service := range f.Services {
		if err := f.GenerateClient(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := f.GenerateNewClient(service, g); err != nil {
			return err
		}
	}
	return nil
}

func (f *Generator) GenerateNewServer(service *internal.Service, generatedFile *protogen.GeneratedFile) error {
	generatedFile.P("func New", service.HTTPServerName(), "(")
	generatedFile.P("endpoints interface {")
	for _, endpoint := range service.Endpoints {
		generatedFile.P(endpoint.Name(), "() ", internal.EndpointPackage.Ident("Endpoint"))
	}
	generatedFile.P("},")
	generatedFile.P("mdw []", internal.EndpointPackage.Ident("Middleware"), ",")
	generatedFile.P("opts ...", internal.HttpTransportPackage.Ident("ServerOption"), ",")
	generatedFile.P(") ", internal.HttpPackage.Ident("Handler"), " {")
	generatedFile.P("r := ", internal.MuxPackage.Ident("NewRouter"), "()")
	for _, endpoint := range service.Endpoints {
		httpRule := endpoint.HttpRule()
		// 调整路径，来适应 github.com/gorilla/mux 路由规则
		path, _, _, _ := httpRule.RegularizePath(httpRule.Path())
		generatedFile.P("r.Name(", strconv.Quote(endpoint.FullName()), ").")
		generatedFile.P("Methods(", strconv.Quote(httpRule.Method()), ").")
		generatedFile.P("Path(", strconv.Quote(path), ").")
		generatedFile.P("Handler(", internal.HttpTransportPackage.Ident("NewServer"), "(")
		generatedFile.P(internal.EndpointxPackage.Ident("Chain"), "(endpoints.", endpoint.Name(), "(), mdw...), ")
		if err := f.PrintDecodeRequestFunc(generatedFile, endpoint); err != nil {
			return err
		}
		if err := f.PrintEncodeResponseFunc(generatedFile, endpoint, httpRule); err != nil {
			return err
		}

		generatedFile.P("},")
		generatedFile.P("opts...,")
		generatedFile.P("))")
	}
	generatedFile.P("return r")
	generatedFile.P("}")
	generatedFile.P()
	return nil
}

func (f *Generator) PrintDecodeRequestFunc(
	generatedFile *protogen.GeneratedFile, endpoint *internal.Endpoint,
) error {
	generatedFile.P("func(ctx ", internal.ContextPackage.Ident("Context"), ", r *", internal.HttpPackage.Ident("Request"), ") (any, error) {")
	generatedFile.P("req := &", endpoint.InputGoIdent(), "{}")

	bodyMessage, bodyField, namedPathFields, pathFields, queryFields, err := endpoint.ParseParameters()
	if err != nil {
		return err
	}

	if bodyMessage != nil {
		switch bodyMessage.Desc.FullName() {
		case "google.api.HttpBody":
			f.PrintApiFromBody(generatedFile, nil)
		case "google.rpc.HttpRequest":
			f.PrintRpcBody(generatedFile, nil)
		default:
			f.printStarBody(generatedFile)
		}
	} else if bodyField != nil {
		if err := f.printFieldBody(generatedFile, bodyField); err != nil {
			return err
		}
	}

	var pathOnce sync.Once
	for i, namedPathField := range namedPathFields {
		pathOnce.Do(func() {
			generatedFile.P("vars := ", internal.MuxPackage.Ident("Vars"), "(r)")
		})
		fullFieldName := internal.FullFieldName(namedPathFields[:i+1])
		if i < len(namedPathFields)-1 {
			generatedFile.P("if req.", fullFieldName, " == nil {")
			generatedFile.P("req.", fullFieldName, " = &", namedPathField.Message.GoIdent, "{}")
			generatedFile.P("}")
		} else {
			httpRule := endpoint.HttpRule()
			_, _, namedPathTemplate, namedPathParameters := httpRule.RegularizePath(httpRule.Path())
			left := []any{"req.", fullFieldName, " = "}
			right := []any{internal.FmtPackage.Ident("Sprintf"), "(", strconv.Quote(namedPathTemplate)}
			for _, namedPathParameter := range namedPathParameters {
				right = append(right, ", vars[", strconv.Quote(namedPathParameter), "]")
			}
			right = append(right, ")")
			if err := f.printAssign(generatedFile, namedPathField, left, right, false); err != nil {
				return err
			}
		}
	}

	for _, pathField := range pathFields {
		pathOnce.Do(func() {
			generatedFile.P("vars := ", internal.MuxPackage.Ident("Vars"), "(r)")
		})
		left := []any{"req.", pathField.GoName, " = "}
		right := []any{"vars[", strconv.Quote(string(pathField.Desc.Name())), "]"}
		if err := f.printAssign(generatedFile, pathField, left, right, false); err != nil {
			return err
		}
	}

	var queryOnce sync.Once
	for _, field := range queryFields {
		queryOnce.Do(func() {
			generatedFile.P("queries := r.URL.Query()")
		})
		fieldName := string(field.Desc.Name())
		if field.Message != nil && field.Message.Desc.FullName() == "google.protobuf.FieldMask" {
			if bodyField != nil {
				generatedFile.P("mask, err := ", internal.FieldmaskpbPackage.Ident("New"), "(req.", bodyField.GoName, ", queries[", strconv.Quote(fieldName), "]...)")
			} else if bodyMessage != nil {
				generatedFile.P("mask, err := ", internal.FieldmaskpbPackage.Ident("New"), "(req", ", queries[", strconv.Quote(fieldName), "]...)")
			}
			generatedFile.P("if err != nil {")
			generatedFile.P("return nil, err")
			generatedFile.P("}")
			generatedFile.P("req.UpdateMask = mask")
			continue
		}
		left := []any{"req.", field.GoName, " = "}
		right := []any{"queries.Get(", strconv.Quote(fieldName), ")"}
		if field.Desc.IsList() {
			right = []any{"queries[", strconv.Quote(fieldName), "]"}
		}
		if err := f.printAssign(generatedFile, field, left, right, field.Desc.IsList()); err != nil {
			return err
		}
	}

	generatedFile.P("return req, nil")
	generatedFile.P("},")
	return nil
}

func (f *Generator) PrintApiFromBody(generatedFile *protogen.GeneratedFile, field *protogen.Field) {
	prefix := "req."
	if field != nil {
		prefix = prefix + field.GoName + "."
	}
	generatedFile.P(prefix, "ContentType = r.Header.Get(", strconv.Quote("Content-Type"), ")")
	generatedFile.P("body, err := ", internal.IOPackage.Ident("ReadAll"), "(r.Body)")
	generatedFile.P("if err != nil {")
	generatedFile.P("return nil, err")
	generatedFile.P("}")
	generatedFile.P(prefix, "Data = body")
}

func (f *Generator) PrintRpcBody(generatedFile *protogen.GeneratedFile, field *protogen.Field) {
	prefix := "req."
	if field != nil {
		prefix = prefix + field.GoName + "."
	}
	generatedFile.P(prefix, "Method = r.Method")
	generatedFile.P(prefix, "Uri = r.RequestURI")
	generatedFile.P(prefix, "Headers = make([]*", internal.RpcHttpPackage.Ident("HttpHeader"), ", 0, len(r.Header))")
	generatedFile.P("for key, values := range r.Header {")
	generatedFile.P("for _, value := range values {")
	generatedFile.P(prefix, "Headers = append(", prefix, "Headers, &", internal.RpcHttpPackage.Ident("HttpHeader"), "{Key: key, Value: value})")
	generatedFile.P("}")
	generatedFile.P("}")
	generatedFile.P("body, err := ", internal.IOPackage.Ident("ReadAll"), "(r.Body)")
	generatedFile.P("if err != nil {")
	generatedFile.P("return nil, err")
	generatedFile.P("}")
	generatedFile.P(prefix, "Body = body")
}

func (f *Generator) printFieldBody(generatedFile *protogen.GeneratedFile, field *protogen.Field) error {
	message := field.Message
	switch {
	case message != nil && message.Desc.FullName() == "google.api.HttpBody":
		f.PrintApiFromBody(generatedFile, field)
	case message != nil && message.Desc.FullName() == "google.rpc.HttpRequest":
		f.PrintRpcBody(generatedFile, field)
	default:
		generatedFile.P("body, err := ", internal.IOPackage.Ident("ReadAll"), "(r.Body)")
		generatedFile.P("if err != nil {")
		generatedFile.P("return nil, err")
		generatedFile.P("}")
		left := []any{"req.", field.GoName, " = "}
		right := []any{"string(body)"}
		if err := f.printAssign(generatedFile, field, left, right, false); err != nil {
			return err
		}
	}
	return nil
}

func (f *Generator) printStarBody(generatedFile *protogen.GeneratedFile) {
	generatedFile.P("body, err := ", internal.IOPackage.Ident("ReadAll"), "(r.Body)")
	generatedFile.P("if err != nil {")
	generatedFile.P("return nil, err")
	generatedFile.P("}")
	generatedFile.P("if err := ", internal.ProtoJsonPackage.Ident("Unmarshal"), "(body, req); err != nil {")
	generatedFile.P("return nil, err")
	generatedFile.P("}")
}

func (f *Generator) printAssign(generatedFile *protogen.GeneratedFile, field *protogen.Field, left []any, right []any, isList bool) error {
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		// bool
		if isList {
			right = append([]any{"if v, err := ", internal.ConvxPackage.Ident("ParseBoolSlice"), "("}, right...)
		} else {
			right = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseBool"), "("}, right...)
		}
		right = append(right, "); err != nil {")
		generatedFile.P(right...)
		generatedFile.P("return nil, err")
		generatedFile.P("} else {")
		if field.Desc.HasOptionalKeyword() {
			generatedFile.P(append(left, internal.ProtoPackage.Ident("Bool"), "(v)")...)
		} else {
			generatedFile.P(append(left, "v")...)
		}
		generatedFile.P("}")
	case protoreflect.EnumKind:
		generatedFile.P("// enum")

	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		// int32
		if isList {
			right = append([]any{"if v, err := ", internal.ConvxPackage.Ident("ParseIntSlice[int32]"), "("}, right...)
		} else {
			right = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseInt"), "("}, right...)
		}
		right = append(right, ", 10, 32); err != nil {")
		generatedFile.P(right...)
		generatedFile.P("return nil, err")
		generatedFile.P("} else {")
		if field.Desc.HasOptionalKeyword() {
			generatedFile.P(append(left, internal.ProtoPackage.Ident("Int32"), "(int32(v))")...)
		} else if isList {
			generatedFile.P(append(left, "v")...)
		} else {
			generatedFile.P(append(left, "int32(v)")...)
		}
		generatedFile.P("}")
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		// uint32
		if isList {
			right = append([]any{"if v, err := ", internal.ConvxPackage.Ident("ParseUintSlice[uint32]"), "("}, right...)
		} else {
			right = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseUint"), "("}, right...)
		}
		right = append(right, ", 10, 32); err != nil {")
		generatedFile.P(right...)
		generatedFile.P("return nil, err")
		generatedFile.P("} else {")
		if field.Desc.HasOptionalKeyword() {
			generatedFile.P(append(left, internal.ProtoPackage.Ident("Uint32"), "(uint32(v))")...)
		} else if isList {
			generatedFile.P(append(left, "v")...)
		} else {
			generatedFile.P(append(left, "uint32(v)")...)
		}
		generatedFile.P("}")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		// int64
		if isList {
			right = append([]any{"if v, err := ", internal.ConvxPackage.Ident("ParseIntSlice[int64]"), "("}, right...)
		} else {
			right = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseInt"), "("}, right...)
		}
		right = append(right, ", 10, 64); err != nil {")
		generatedFile.P(right...)
		generatedFile.P("return nil, err")
		generatedFile.P("} else {")
		if field.Desc.HasOptionalKeyword() {
			generatedFile.P(append(left, internal.ProtoPackage.Ident("Int64"), "(v)")...)
		} else {
			generatedFile.P(append(left, "v")...)
		}
		generatedFile.P("}")
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		// uint64
		if isList {
			right = append([]any{"if v, err := ", internal.ConvxPackage.Ident("ParseUintSlice[uint64]"), "("}, right...)
		} else {
			right = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseUint"), "("}, right...)
		}
		right = append(right, ", 10, 64); err != nil {")
		generatedFile.P(right...)
		generatedFile.P("return nil, err")
		generatedFile.P("} else {")
		if field.Desc.HasOptionalKeyword() {
			generatedFile.P(append(left, internal.ProtoPackage.Ident("Uint64"), "(v)")...)
		} else {
			generatedFile.P(append(left, "v")...)
		}
		generatedFile.P("}")
	case protoreflect.FloatKind:
		// float32
		if isList {
			right = append([]any{"if v, err := ", internal.ConvxPackage.Ident("ParseFloatSlice[float32]"), "("}, right...)
		} else {
			right = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseFloat"), "("}, right...)
		}
		right = append(right, ", 32); err != nil {")
		generatedFile.P(right...)
		generatedFile.P("return nil, err")
		generatedFile.P("} else {")
		if field.Desc.HasOptionalKeyword() {
			generatedFile.P(append(left, internal.ProtoPackage.Ident("Float32"), "(float32(v))")...)
		} else if isList {
			generatedFile.P(append(left, "v")...)
		} else {
			generatedFile.P(append(left, "float32(v)")...)
		}
		generatedFile.P("}")
	case protoreflect.DoubleKind:
		// float64
		if isList {
			right = append([]any{"if v, err := ", internal.ConvxPackage.Ident("ParseFloatSlice[float64]"), "("}, right...)
		} else {
			right = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseFloat"), "("}, right...)
		}
		right = append(right, ", 32); err != nil {")
		generatedFile.P(right...)
		generatedFile.P("return nil, err")
		generatedFile.P("} else {")
		if field.Desc.HasOptionalKeyword() {
			generatedFile.P(append(left, internal.ProtoPackage.Ident("Float64"), "(v)")...)
		} else {
			generatedFile.P(append(left, "v")...)
		}
		generatedFile.P("}")
	case protoreflect.StringKind:
		// string
		if field.Desc.HasOptionalKeyword() {
			a := []any{internal.ProtoPackage.Ident("String"), "("}
			right = append(a, right...)
			right = append(right, ")")
			generatedFile.P(append(left, right...)...)
		} else {
			generatedFile.P(append(left, right...)...)
		}
	case protoreflect.BytesKind:
		// []byte
		if isList {
			right = append([]any{internal.ConvxPackage.Ident("ParseBytesSlice"), "("}, right...)
			right = append(right, ")")
			generatedFile.P(append(left, right...)...)
		} else {
			right = append([]any{"[]byte("}, right...)
			right = append(right, ")")
			generatedFile.P(append(left, right...)...)
		}
	case protoreflect.MessageKind:
		switch field.Message.Desc.FullName() {
		case "google.protobuf.DoubleValue":
			right = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseFloat"), "("}, right...)
			right = append(right, ", 64); err != nil {")
			generatedFile.P(right...)
			generatedFile.P("return nil, err")
			generatedFile.P("} else {")
			generatedFile.P(append(left, internal.WrapperspbPackage.Ident("Double"), "(v)")...)
			generatedFile.P("}")
		case "google.protobuf.FloatValue":
			right = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseFloat"), "("}, right...)
			right = append(right, ", 32); err != nil {")
			generatedFile.P(right...)
			generatedFile.P("return nil, err")
			generatedFile.P("} else {")
			generatedFile.P(append(left, internal.WrapperspbPackage.Ident("Float"), "(float32(v))")...)
			generatedFile.P("}")
		case "google.protobuf.Int64Value":
			right = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseInt"), "("}, right...)
			right = append(right, ", 10, 64); err != nil {")
			generatedFile.P(right...)
			generatedFile.P("return nil, err")
			generatedFile.P("} else {")
			generatedFile.P(append(left, internal.WrapperspbPackage.Ident("Int64"), "(v)")...)
			generatedFile.P("}")
		case "google.protobuf.UInt64Value":
			right = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseUint"), "("}, right...)
			right = append(right, ", 10, 64); err != nil {")
			generatedFile.P(right...)
			generatedFile.P("return nil, err")
			generatedFile.P("} else {")
			generatedFile.P(append(left, internal.WrapperspbPackage.Ident("UInt64"), "(v)")...)
			generatedFile.P("}")
		case "google.protobuf.Int32Value":
			right = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseInt"), "("}, right...)
			right = append(right, ", 10, 32); err != nil {")
			generatedFile.P(right...)
			generatedFile.P("return nil, err")
			generatedFile.P("} else {")
			generatedFile.P(append(left, internal.WrapperspbPackage.Ident("Int32"), "(int32(v))")...)
			generatedFile.P("}")
		case "google.protobuf.UInt32Value":
			right = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseUint"), "("}, right...)
			right = append(right, ", 10, 32); err != nil {")
			generatedFile.P(right...)
			generatedFile.P("return nil, err")
			generatedFile.P("} else {")
			generatedFile.P(append(left, internal.WrapperspbPackage.Ident("UInt32"), "(uint32(v))")...)
			generatedFile.P("}")
		case "google.protobuf.BoolValue":
			right = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseBool"), "("}, right...)
			right = append(right, "); err != nil {")
			generatedFile.P(right...)
			generatedFile.P("return nil, err")
			generatedFile.P("} else {")
			generatedFile.P(append(left, internal.WrapperspbPackage.Ident("Bool"), "(v)")...)
			generatedFile.P("}")
		case "google.protobuf.StringValue":
			a := []any{internal.WrapperspbPackage.Ident("String"), "("}
			right = append(a, right...)
			right = append(right, ")")
			generatedFile.P(append(left, right...)...)
		case "google.protobuf.BytesValue":
			a := []any{internal.WrapperspbPackage.Ident("Bytes"), "([]byte("}
			right = append(a, right...)
			right = append(right, "))")
			generatedFile.P(append(left, right...)...)
		default:
			generatedFile.P("if err := ", internal.ProtoJsonPackage.Ident("Unmarshal"), "(body, req.", field.GoName, "); err != nil {")
			generatedFile.P("return nil, err")
			generatedFile.P("}")
		}
	case protoreflect.GroupKind:
		generatedFile.P("// group")

	default:
		return fmt.Errorf("unsupported field type: %+v", internal.FullMessageTypeName(field.Desc.Message()))
	}
	return nil
}

func (f *Generator) PrintEncodeResponseFunc(generatedFile *protogen.GeneratedFile, endpoint *internal.Endpoint, httpRule *internal.HttpRule) error {
	generatedFile.P("func(ctx ", internal.ContextPackage.Ident("Context"), ", w ", internal.HttpPackage.Ident("ResponseWriter"), ", obj any) error {")
	generatedFile.P("resp := obj.(*", endpoint.Output().GoIdent, ")")
	generatedFile.P("_ = resp")
	bodyParameter := httpRule.ResponseBody()
	switch bodyParameter {
	case "":
		if err := f.PrintResponse(generatedFile, endpoint.Output(), "resp"); err != nil {
			return err
		}
	default:
		field := internal.FindField(bodyParameter, endpoint.Output())
		if field == nil {
			return errNotFoundField(endpoint, []string{bodyParameter})
		}
		if err := f.PrintResponse(generatedFile, field.Message, "resp."+field.GoName); err != nil {
			return err
		}
	}
	generatedFile.P("return nil")
	return nil
}

func (f *Generator) PrintResponse(generatedFile *protogen.GeneratedFile, message *protogen.Message, prefix string) error {
	switch message.Desc.FullName() {
	case "google.api.HttpBody":
		generatedFile.P("w.WriteHeader(", internal.HttpPackage.Ident("StatusOK"), ")")
		generatedFile.P("w.Header().Set(", strconv.Quote("Content-Type"), ", ", prefix, ".GetContentType())")
		generatedFile.P()
		generatedFile.P("if ", "_, err := w.Write(", prefix, ".GetData())", "; err != nil {")
		generatedFile.P("return err")
		generatedFile.P("}")
	case "google.rpc.HttpResponse":
		generatedFile.P("w.WriteHeader(int(", prefix, ".GetStatus()))")
		generatedFile.P("for _, header := range ", prefix, ".GetHeaders() {")
		generatedFile.P("w.Header().Add(header.Key, header.Value)")
		generatedFile.P("}")
		generatedFile.P("if ", "_, err := w.Write(", prefix, ".GetBody())", "; err != nil {")
		generatedFile.P("return err")
		generatedFile.P("}")
	default:
		generatedFile.P("w.WriteHeader(", internal.HttpPackage.Ident("StatusOK"), ")")
		generatedFile.P("data, err := ", internal.ProtoJsonPackage.Ident("Marshal"), "(", prefix, ")")
		generatedFile.P("if err != nil {")
		generatedFile.P("return err")
		generatedFile.P("}")
		generatedFile.P("if _, err := w.Write(data); err != nil {")
		generatedFile.P("return err")
		generatedFile.P("}")
	}
	return nil
}

func (f *Generator) GenerateClient(service *internal.Service, generatedFile *protogen.GeneratedFile) error {
	generatedFile.P("type ", service.UnexportedHTTPClientName(), " struct {")
	for _, endpoint := range service.Endpoints {
		generatedFile.P(endpoint.UnexportedName(), " ", internal.EndpointPackage.Ident("Endpoint"))
	}
	generatedFile.P("}")
	generatedFile.P()
	for _, endpoint := range service.Endpoints {
		generatedFile.P("func (c *", service.UnexportedHTTPClientName(), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error){")
		generatedFile.P("rep, err := c.", endpoint.UnexportedName(), "(ctx, request)")
		generatedFile.P("if err != nil {")
		generatedFile.P("return nil, err")
		generatedFile.P("}")
		generatedFile.P("return rep.(*", endpoint.OutputGoIdent(), "), nil")
		generatedFile.P("}")
		generatedFile.P()
	}
	return nil
}

func (f *Generator) GenerateNewClient(service *internal.Service, generatedFile *protogen.GeneratedFile) error {
	generatedFile.P("func New", service.HTTPClientName(), "(")
	generatedFile.P("instance string,")
	generatedFile.P("mdw []", internal.EndpointPackage.Ident("Middleware"), ",")
	generatedFile.P("opts ...", internal.HttpTransportPackage.Ident("ClientOption"), ",")
	generatedFile.P(") interface {")
	for _, endpoint := range service.Endpoints {
		generatedFile.P(endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error)")
	}
	generatedFile.P("} {")
	generatedFile.P("r := ", internal.MuxPackage.Ident("NewRouter"), "()")
	for _, endpoint := range service.Endpoints {
		httpRule := endpoint.HttpRule()
		// 调整路径，来适应 github.com/gorilla/mux 路由规则
		path, _, _, _ := httpRule.RegularizePath(httpRule.Path())
		generatedFile.P("r.Name(", strconv.Quote(endpoint.FullName()), ").")
		generatedFile.P("Methods(", strconv.Quote(httpRule.Method()), ").")
		generatedFile.P("Path(", strconv.Quote(path), ")")
	}
	generatedFile.P("return &", service.UnexportedHTTPClientName(), "{")
	for _, endpoint := range service.Endpoints {
		generatedFile.P(endpoint.UnexportedName(), ":    ", internal.EndpointxPackage.Ident("Chain"), "(")
		generatedFile.P(internal.HttpTransportPackage.Ident("NewExplicitClient"), "(")
		if err := f.PrintEncodeRequestFunc(generatedFile, endpoint); err != nil {
			return err
		}
		if err := f.PrintDecodeResponseFunc(generatedFile); err != nil {
			return err
		}
		generatedFile.P("opts...,")
		generatedFile.P(").Endpoint(),")
		generatedFile.P("mdw...),")
	}
	generatedFile.P("}")
	generatedFile.P("}")
	generatedFile.P()
	return nil
}

func (f *Generator) PrintEncodeRequestFunc(generatedFile *protogen.GeneratedFile, endpoint *internal.Endpoint) error {
	httpRule := endpoint.HttpRule()
	generatedFile.P("func(ctx context.Context, obj interface{}) (*http1.Request, error) {")
	generatedFile.P("req := obj.(*", endpoint.InputGoIdent(), ")")
	generatedFile.P("var method = ", strconv.Quote(httpRule.Method()))
	path, _, _, _ := httpRule.RegularizePath(httpRule.Path())
	generatedFile.P("var url = ", strconv.Quote(path))
	generatedFile.P("var body ", internal.IOPackage.Ident("Reader"))
	_ = httpRule
	bodyMessage, bodyField, namedPathFields, pathFields, queryFields, err := endpoint.ParseParameters()
	if err != nil {
		return err
	}

	if bodyMessage != nil {
		message := bodyMessage
		srcValue := []any{"req"}
		f.PrintMessageBody(generatedFile, message, srcValue, false)
	} else if bodyField != nil {
		field := bodyField
		srcValue := []any{"req.", field.GoName}
		isList := field.Desc.IsList()
		isMap := field.Desc.IsMap()
		isOptional := field.Desc.HasOptionalKeyword()
		switch field.Desc.Kind() {
		case protoreflect.BoolKind:
			// bool
			if isList {
				f.PrintListFieldBody(generatedFile, srcValue)
			} else if isOptional {
				format := []any{internal.StrconvPackage.Ident("FormatBool"), "(*", "req.", field.GoName, ")"}
				f.PrintOptionalFieldBody(generatedFile, internal.StringsPackage, field, format)
			} else {
				format := append(append([]any{internal.StrconvPackage.Ident("FormatBool"), "("}, srcValue...), []any{")"}...)
				f.PrintSampleFieldBody(generatedFile, internal.StringsPackage, format)
			}

		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
			// int32
			if isList {
				f.PrintListFieldBody(generatedFile, srcValue)
			} else if isOptional {
				format := []any{internal.StrconvPackage.Ident("FormatInt"), "(int64(*", "req.", field.GoName, "), 10)"}
				f.PrintOptionalFieldBody(generatedFile, internal.StringsPackage, field, format)
			} else {
				format := append(append([]any{internal.StrconvPackage.Ident("FormatInt"), "(int64("}, srcValue...), []any{"), 10)"}...)
				f.PrintSampleFieldBody(generatedFile, internal.StringsPackage, format)
			}

		case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
			// uint32
			if isList {
				f.PrintListFieldBody(generatedFile, srcValue)
			} else if isOptional {
				format := []any{internal.StrconvPackage.Ident("FormatUint"), "(uint64(*", "req.", field.GoName, "), 10)"}
				f.PrintOptionalFieldBody(generatedFile, internal.StringsPackage, field, format)
			} else {
				format := append(append([]any{internal.StrconvPackage.Ident("FormatUint"), "(uint64("}, srcValue...), []any{"), 10)"}...)
				f.PrintSampleFieldBody(generatedFile, internal.StringsPackage, format)
			}

		case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
			// int64
			if isList {
				f.PrintListFieldBody(generatedFile, srcValue)
			} else if isOptional {
				format := []any{internal.StrconvPackage.Ident("FormatInt"), "(*", "req.", field.GoName, ", 10)"}
				f.PrintOptionalFieldBody(generatedFile, internal.StringsPackage, field, format)
			} else {
				format := append(append([]any{internal.StrconvPackage.Ident("FormatInt"), "("}, srcValue...), []any{", 10)"}...)
				f.PrintSampleFieldBody(generatedFile, internal.StringsPackage, format)
			}

		case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
			// uint64
			if isList {
				f.PrintListFieldBody(generatedFile, srcValue)
			} else if isOptional {
				format := []any{internal.StrconvPackage.Ident("FormatUint"), "(*", "req.", field.GoName, ", 10)"}
				f.PrintOptionalFieldBody(generatedFile, internal.StringsPackage, field, format)
			} else {
				format := append(append([]any{internal.StrconvPackage.Ident("FormatUint"), "("}, srcValue...), []any{", 10)"}...)
				f.PrintSampleFieldBody(generatedFile, internal.StringsPackage, format)
			}

		case protoreflect.FloatKind:
			// float32
			if isList {
				f.PrintListFieldBody(generatedFile, srcValue)
			} else if isOptional {
				format := []any{internal.StrconvPackage.Ident("FormatFloat"), "(float64(*", "req.", field.GoName, "), 'f', -1, 32)"}
				f.PrintOptionalFieldBody(generatedFile, internal.StringsPackage, field, format)
			} else {
				format := append(append([]any{internal.StrconvPackage.Ident("FormatFloat"), "(float64("}, srcValue...), []any{"), 'f', -1, 32)"}...)
				f.PrintSampleFieldBody(generatedFile, internal.StringsPackage, format)
			}

		case protoreflect.DoubleKind:
			// float64
			if isList {
				f.PrintListFieldBody(generatedFile, srcValue)
			} else if isOptional {
				format := []any{internal.StrconvPackage.Ident("FormatFloat"), "(*", "req.", field.GoName, ", 'f', -1, 64)"}
				f.PrintOptionalFieldBody(generatedFile, internal.StringsPackage, field, format)
			} else {
				format := append(append([]any{internal.StrconvPackage.Ident("FormatFloat"), "("}, srcValue...), []any{", 'f', -1, 64)"}...)
				f.PrintSampleFieldBody(generatedFile, internal.StringsPackage, format)
			}

		case protoreflect.StringKind:
			// string
			if isList {
				f.PrintListFieldBody(generatedFile, srcValue)
			} else if isOptional {
				format := []any{"*", "req.", field.GoName}
				f.PrintOptionalFieldBody(generatedFile, internal.StringsPackage, field, format)
			} else {
				format := append(append([]any{}, srcValue...), []any{}...)
				f.PrintSampleFieldBody(generatedFile, internal.StringsPackage, format)
			}

		case protoreflect.BytesKind:
			// []byte
			if isList {
				f.PrintListFieldBody(generatedFile, srcValue)
			} else if isOptional {
				format := []any{"req.", field.GoName}
				f.PrintSampleFieldBody(generatedFile, internal.BytesPackage, format)
			} else {
				format := append(append([]any{}, srcValue...), []any{}...)
				f.PrintSampleFieldBody(generatedFile, internal.BytesPackage, format)
			}

		case protoreflect.EnumKind:
			// enum
			if isList {
				f.PrintListFieldBody(generatedFile, srcValue)
			} else if isOptional {
				format := []any{internal.StrconvPackage.Ident("FormatInt"), "(int64(*", "req.", field.GoName, "), 10)"}
				f.PrintOptionalFieldBody(generatedFile, internal.StringsPackage, field, format)
			} else {
				format := append(append([]any{internal.StrconvPackage.Ident("FormatInt"), "(int64("}, srcValue...), []any{"), 10)"}...)
				f.PrintSampleFieldBody(generatedFile, internal.StringsPackage, format)
			}

		case protoreflect.MessageKind:
			message := field.Message
			f.PrintMessageBody(generatedFile, message, srcValue, isMap)
		case protoreflect.GroupKind:
			generatedFile.P("// group")
		default:
			return fmt.Errorf("unsupported field type: %+v", internal.FullMessageTypeName(field.Desc.Message()))
		}
	}
	_ = namedPathFields
	_ = pathFields
	_ = queryFields
	//
	//var pathOnce sync.Once
	//for i, namedPathField := range namedPathFields {
	//	pathOnce.Do(func() {
	//		generatedFile.P("vars := ", internal.MuxPackage.Ident("Vars"), "(r)")
	//	})
	//	fullFieldName := internal.FullFieldName(namedPathFields[:i+1])
	//	if i < len(namedPathFields)-1 {
	//		generatedFile.P("if req.", fullFieldName, " == nil {")
	//		generatedFile.P("req.", fullFieldName, " = &", namedPathField.Message.GoIdent, "{}")
	//		generatedFile.P("}")
	//	} else {
	//		_, _, namedPathTemplate, namedPathParameters := httpRule.RegularizePath(httpRule.Path())
	//		left := []any{"req.", fullFieldName, " = "}
	//		right := []any{internal.FmtPackage.Ident("Sprintf"), "(", strconv.Quote(namedPathTemplate)}
	//		for _, namedPathParameter := range namedPathParameters {
	//			right = append(right, ", vars[", strconv.Quote(namedPathParameter), "]")
	//		}
	//		right = append(right, ")")
	//		if err := f.printAssign(generatedFile, namedPathField, left, right, false); err != nil {
	//			return err
	//		}
	//	}
	//}
	//
	//for _, pathField := range pathFields {
	//	pathOnce.Do(func() {
	//		generatedFile.P("vars := ", internal.MuxPackage.Ident("Vars"), "(r)")
	//	})
	//	left := []any{"req.", pathField.GoName, " = "}
	//	right := []any{"vars[", strconv.Quote(string(pathField.Desc.Name())), "]"}
	//	if err := f.printAssign(generatedFile, pathField, left, right, false); err != nil {
	//		return err
	//	}
	//}
	//
	//var queryOnce sync.Once
	//for _, field := range queryFields {
	//	queryOnce.Do(func() {
	//		generatedFile.P("queries := r.URL.Query()")
	//	})
	//	fieldName := string(field.Desc.Name())
	//	if field.Message != nil && field.Message.Desc.FullName() == "google.protobuf.FieldMask" {
	//		if bodyField != nil {
	//			generatedFile.P("mask, err := ", internal.FieldmaskpbPackage.Ident("New"), "(req.", bodyField.GoName, ", queries[", strconv.Quote(fieldName), "]...)")
	//		} else if bodyMessage != nil {
	//			generatedFile.P("mask, err := ", internal.FieldmaskpbPackage.Ident("New"), "(req", ", queries[", strconv.Quote(fieldName), "]...)")
	//		}
	//		generatedFile.P("if err != nil {")
	//		generatedFile.P("return nil, err")
	//		generatedFile.P("}")
	//		generatedFile.P("req.UpdateMask = mask")
	//		continue
	//	}
	//	left := []any{"req.", field.GoName, " = "}
	//	right := []any{"queries.Get(", strconv.Quote(fieldName), ")"}
	//	if field.Desc.IsList() {
	//		right = []any{"queries[", strconv.Quote(fieldName), "]"}
	//	}
	//	if err := f.printAssign(generatedFile, field, left, right, field.Desc.IsList()); err != nil {
	//		return err
	//	}
	//}

	generatedFile.P("r, err := http1.NewRequestWithContext(ctx, method, url, body)")
	generatedFile.P("if err != nil {")
	generatedFile.P("return nil, err")
	generatedFile.P("}")

	generatedFile.P("return r, nil")
	generatedFile.P("},")
	return nil
}

func (f *Generator) PrintSampleFieldBody(generatedFile *protogen.GeneratedFile, readerPkg protogen.GoImportPath, format []any) {
	generatedFile.P(append(append([]any{"body = ", readerPkg.Ident("NewReader"), "("}, format...), ")")...)
}

func (f *Generator) PrintOptionalFieldBody(generatedFile *protogen.GeneratedFile, readerPkg protogen.GoImportPath, field *protogen.Field, format []any) {
	generatedFile.P("if ", "req.", field.GoName, " != nil {")
	generatedFile.P(append(append([]any{"body = ", readerPkg.Ident("NewReader"), "("}, format...), ")")...)
	generatedFile.P("}")
}

func (f *Generator) PrintMessageBody(generatedFile *protogen.GeneratedFile, message *protogen.Message, srcValue []any, isMap bool) {
	switch message.Desc.FullName() {
	case "google.protobuf.DoubleValue":
		format := append(append([]any{internal.StrconvPackage.Ident("FormatFloat"), "("}, srcValue...), []any{".Value", ", 'f', -1, 64)"}...)
		f.PrintWrapFieldBody(generatedFile, internal.StringsPackage, srcValue, format)
	case "google.protobuf.FloatValue":
		format := append(append([]any{internal.StrconvPackage.Ident("FormatFloat"), "(float64("}, srcValue...), []any{".Value", "), 'f', -1, 32)"}...)
		f.PrintWrapFieldBody(generatedFile, internal.StringsPackage, srcValue, format)
	case "google.protobuf.Int64Value":
		format := append(append([]any{internal.StrconvPackage.Ident("FormatInt"), "("}, srcValue...), []any{".Value", ", 10)"}...)
		f.PrintWrapFieldBody(generatedFile, internal.StringsPackage, srcValue, format)
	case "google.protobuf.UInt64Value":
		format := append(append([]any{internal.StrconvPackage.Ident("FormatUint"), "("}, srcValue...), []any{".Value", ", 10)"}...)
		f.PrintWrapFieldBody(generatedFile, internal.StringsPackage, srcValue, format)
	case "google.protobuf.Int32Value":
		format := append(append([]any{internal.StrconvPackage.Ident("FormatInt"), "(int64("}, srcValue...), []any{".Value", "), 10)"}...)
		f.PrintWrapFieldBody(generatedFile, internal.StringsPackage, srcValue, format)
	case "google.protobuf.UInt32Value":
		format := append(append([]any{internal.StrconvPackage.Ident("FormatUint"), "(uint64("}, srcValue...), []any{".Value", "), 10)"}...)
		f.PrintWrapFieldBody(generatedFile, internal.StringsPackage, srcValue, format)
	case "google.protobuf.BoolValue":
		format := append(append([]any{internal.StrconvPackage.Ident("FormatBool"), "("}, srcValue...), []any{".Value", ")"}...)
		f.PrintWrapFieldBody(generatedFile, internal.StringsPackage, srcValue, format)
	case "google.protobuf.StringValue":
		format := append(append([]any{}, srcValue...), []any{".Value"}...)
		f.PrintWrapFieldBody(generatedFile, internal.StringsPackage, srcValue, format)
	case "google.protobuf.BytesValue":
		format := append(append([]any{}, srcValue...), []any{".Value"}...)
		f.PrintWrapFieldBody(generatedFile, internal.BytesPackage, srcValue, format)
	case "google.api.HttpBody":
		format := append(append([]any{}, srcValue...), []any{".Data"}...)
		f.PrintWrapFieldBody(generatedFile, internal.BytesPackage, srcValue, format)
	case "google.rpc.HttpRequest":
		format := append(append([]any{}, srcValue...), []any{".Body"}...)
		f.PrintWrapFieldBody(generatedFile, internal.BytesPackage, srcValue, format)
	default:
		if isMap {
			f.PrintMessageFieldBody(generatedFile, internal.JsonPackage, srcValue)
		} else {
			f.PrintMessageFieldBody(generatedFile, internal.ProtoJsonPackage, srcValue)
		}
	}
}

func (f *Generator) PrintWrapFieldBody(generatedFile *protogen.GeneratedFile, readerPkg protogen.GoImportPath, srcValue, format []any) {
	generatedFile.P(append(append([]any{"if "}, srcValue...), " != nil {")...)
	generatedFile.P(append(append([]any{"body = ", readerPkg.Ident("NewReader"), "("}, format...), ")")...)
	generatedFile.P("}")
}

func (f *Generator) PrintMessageFieldBody(generatedFile *protogen.GeneratedFile, marshalPkg protogen.GoImportPath, srcValue []any) {
	generatedFile.P(append(append([]any{"data, err := ", marshalPkg.Ident("Marshal"), "("}, srcValue...), []any{")"}...)...)
	generatedFile.P("if err != nil {")
	generatedFile.P("return nil, err")
	generatedFile.P("}")
	generatedFile.P("body = ", internal.BytesPackage.Ident("NewBuffer"), "(data)")
}

func (f *Generator) PrintListFieldBody(generatedFile *protogen.GeneratedFile, srcValue []any) {
	generatedFile.P(append(append([]any{"if "}, srcValue...), []any{" != nil {"}...)...)
	generatedFile.P(append(append([]any{"if err := ", internal.JsonPackage.Ident("NewDecoder"), "(body).Decode("}, srcValue...), []any{"); err != nil {"}...)...)
	generatedFile.P("return nil, err")
	generatedFile.P("}")
	generatedFile.P("}")
}

func (f *Generator) PrintDecodeResponseFunc(generatedFile *protogen.GeneratedFile) error {
	generatedFile.P("func(ctx context.Context, r *http1.Response) (interface{}, error) {")
	generatedFile.P("return nil, nil")
	generatedFile.P("},")

	return nil
}

func (f *Generator) PrintApiToBody(generatedFile *protogen.GeneratedFile, endpoint *internal.Endpoint, field *protogen.Field) {
	prefix := "req."
	if field != nil {
		prefix = prefix + field.GoName + "."
	}
	generatedFile.P(prefix, "ContentType = r.Header.Get(", strconv.Quote("Content-Type"), ")")
	generatedFile.P("body, err := ", internal.IOPackage.Ident("ReadAll"), "(r.Body)")
	generatedFile.P("if err != nil {")
	generatedFile.P("return nil, err")
	generatedFile.P("}")
	generatedFile.P(prefix, "Data = body")

	rule := endpoint.HttpRule()

	generatedFile.P("r, err := ", internal.HttpPackage.Ident("NewRequest"), "(", strconv.Quote(rule.Method()), ", ", strconv.Quote("/v1/users"), ", ", internal.BytesPackage.Ident("NewReader"), "(req.GetData()))")
	generatedFile.P("if err != nil {")
	generatedFile.P("return nil, err")
	generatedFile.P("}")
	generatedFile.P("r.Header.Set(", strconv.Quote("Content-Type"), ", ", prefix, "GetContentType())")
	generatedFile.P("return r, nil")
}
