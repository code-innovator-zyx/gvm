package main

import (
	"github.com/code-innovator-zyx/gvm/internal/tui/pipeline"
	"os"
)

func main() {
	os.Setenv("http_proxy", "127.0.0.1:7890")
	os.Setenv("https_proxy", "127.0.0.1:7890")
	pipeline.NewtProgram().Run()
}
