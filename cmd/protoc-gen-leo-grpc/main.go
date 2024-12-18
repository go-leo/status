package main

import (
	"flag"
	"github.com/go-leo/leo/v3/cmd/protoc-gen-leo-grpc/generator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var flags flag.FlagSet

func main() {
	options := &protogen.Options{ParamFunc: flags.Set}
	options.Run(func(plugin *protogen.Plugin) error {
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		return generate(plugin)
	})
}

func generate(plugin *protogen.Plugin) error {
	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}
		if len(file.Services) <= 0 {
			continue
		}

		gen, err := generator.NewGenerator(plugin, file)
		if err != nil {
			return err
		}
		if err := gen.Generate(); err != nil {
			return err
		}
	}
	return nil
}
