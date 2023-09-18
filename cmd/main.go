package main

import (
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

func main() {
	css := api.Build(api.BuildOptions{
		EntryPoints:      []string{"css/daisy.css"},
		Bundle:           true,
		Outfile:          "assets/bundle.css",
		Write:            true,
		MinifyWhitespace: true,
		MinifySyntax:     true,
		TreeShaking:      api.TreeShakingTrue,
	})

	if len(css.Errors) > 0 {
		os.Exit(1)
	}

	js := api.Build(api.BuildOptions{
		EntryPoints:      []string{"css/tailwind.js"},
		Bundle:           true,
		Outfile:          "assets/bundle.js",
		Write:            true,
		MinifyWhitespace: true,
		MinifySyntax:     true,
		TreeShaking:      api.TreeShakingTrue,
	})

	if len(js.Errors) > 0 {
		os.Exit(1)
	}
}
