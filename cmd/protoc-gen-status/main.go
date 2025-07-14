package main

import (
	"flag"
	"fmt"

	"github.com/go-leo/status/cmd/protoc-gen-status/gen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-leo %v\n", "v0.0.1")
		return
	}

	var flags flag.FlagSet
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

		// 错误状态码生成
		statusGenerator := gen.NewGenerator(plugin, file)
		statusGenerator.Generate()

	}
	return nil
}
