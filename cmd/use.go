/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/code-innovator-zyx/gvm/internal/consts"
	"github.com/code-innovator-zyx/gvm/pkg"
	"github.com/spf13/cobra"
	"os"
	"regexp"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use [version]",
	Short: "Switch to a specific Go version",
	Long: `Switch the current Go environment to the specified version.

Example:
  gvm use go1.21
This will activate Go 1.21 for the current shell session.`,
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]
		if version == "" {
			modVersion, err := getModVersion()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			version = modVersion
		}
		localVersions, _ := pkg.NewManager(false).List(consts.All)
		for _, localVersion := range localVersions {
			if localVersion.String() != version {
				continue
			}
			if err := pkg.SwitchVersion(localVersion); err != nil {
				fmt.Println(err.Error())
			}
			return
		}
		// TODO: Replace this with your actual logic
		// For example:
		// 1. Verify that the version exists in your gvm versions directory
		// 2. Update a "current" symlink (e.g., ~/.g/go/current -> ~/.g/go/versions/go1.21)
		// 3. Adjust PATH so the chosen version takes effect

		//fmt.Printf("Switched to Go version %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// useCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// useCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var goModReg = regexp.MustCompile(`(?m)^go\s+(\d+\.\d+(?:\.\d+)?(?:beta\d+|rc\d+)?)\s*(?:$|//.*)`)

func getModVersion() (string, error) {
	// Uses go.mod if available and version is omitted
	goModData, err := os.ReadFile("go.mod")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", errors.New("no go.mod file found")
		}
		return "", err
	}
	match := goModReg.FindStringSubmatch(string(goModData))
	if len(match) > 1 {
		return match[1], nil
	}
	return "", errors.New("no version found in go.mod")
}
