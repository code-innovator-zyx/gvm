/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/code-innovator-zyx/gvm/cmd"
	_ "github.com/code-innovator-zyx/gvm/internal/tui/list"
	_ "github.com/code-innovator-zyx/gvm/internal/tui/progress"
	_ "github.com/code-innovator-zyx/gvm/internal/tui/spinner"
	_ "github.com/code-innovator-zyx/gvm/pkg"
)

func main() {
	cmd.Execute()
}
