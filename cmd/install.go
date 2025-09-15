/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/code-innovator-zyx/gvm/internal/consts"
	"github.com/code-innovator-zyx/gvm/internal/version"
	"github.com/code-innovator-zyx/gvm/pkg"
	"github.com/mholt/archiver/v3"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "A brief description of your command",
	Args:  cobra.MinimumNArgs(1),
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		installVersion := args[0]
		if pkg.LocalInstalled(installVersion) {
			cmd.Printf("%s has already been installed\n", installVersion)
			return
		}
		versions, err := pkg.NewManager(true).List(consts.All)
		if err != nil {
			cmd.Println(err.Error())
			return
		}
		v, err := version.NewFinder(versions).Find(installVersion)
		if err != nil {
			cmd.Println(err.Error())
			return
		}
		artifact, err := v.FindArtifact()
		if nil != err {
			cmd.Println(err.Error())
			return
		}
		defer artifact.Clean()

		_, err = artifact.Download()
		if nil != err {
			cmd.Println(err.Error())
			return
		}
		unArchiveFile := filepath.Join(filepath.Dir(artifact.SaveFile()), "go")
		err = archiver.Unarchive(artifact.SaveFile(), consts.VERSION_DIR)
		if nil != err {
			cmd.Println(err.Error())
			return
		}
		err = os.Rename(unArchiveFile, fmt.Sprintf("go%s", installVersion))
		if err != nil {
			cmd.Println(err.Error())
			return
		}
		cmd.Println("successfully installed " + installVersion)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
