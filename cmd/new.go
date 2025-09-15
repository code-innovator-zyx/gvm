/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new Go project with the current active version",
	Long: `Create a new Go project initialized with the Go version currently set by gvm.

Example:
  gvm new myapp
This will create a folder 'myapp', initialize a Go module,
and set it up using the active Go version.`,
	Args: cobra.ExactArgs(1), // require project name
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		// TODO: Determine current Go version (e.g., from ~/.g/go/current symlink)
		// currentVersion := detectCurrentGoVersion()

		// 1. Create project folder
		if err := os.Mkdir(projectName, 0755); err != nil {
			cmd.Printf("Error creating project folder: %v\n", err)
			return
		}

		// 2. Initialize go.mod
		cmdStr := exec.Command("go", "mod", "init", projectName)
		cmdStr.Dir = projectName
		if output, err := cmdStr.CombinedOutput(); err != nil {
			cmd.Printf("Error running go mod init: %v\nOutput: %s\n", err, string(output))
			return
		}

		// 3. Optionally: create a main.go
		mainGo := `package main

import "fmt"

func main() {
	fmt.Println("Hello from ` + projectName + `!")
}`
		if err := os.WriteFile(filepath.Join(projectName, "main.go"), []byte(mainGo), 0644); err != nil {
			cmd.Printf("Error writing main.go: %v\n", err)
			return
		}

		cmd.Printf("✅ New Go project '%s' created successfully!\n", projectName)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
