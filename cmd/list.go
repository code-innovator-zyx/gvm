/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/code-innovator-zyx/gvm/internal/consts"
	"github.com/code-innovator-zyx/gvm/internal/prettyout"
	"github.com/code-innovator-zyx/gvm/internal/version"
	"github.com/code-innovator-zyx/gvm/pkg"
	"github.com/spf13/cobra"
	"slices"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Go versions",
	Long: `Display all Go versions.
Example:
  gvm list
    Show all Go versions installed locally.
  gvm list -r
    Show all available Go versions remotely.`,
	Aliases: []string{"l", "ls"},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		remote, _ := cmd.Flags().GetBool("remote")
		vk := consts.All
		if remote {
			kind, _ := cmd.Flags().GetString("type")
			vk, err = consts.ParseVersionKind(kind)
			if err != nil {
				return err
			}
		}

		versions, err := pkg.NewVManager(remote, pkg.WithLocal()).List(vk)
		if err != nil {
			return err
		}
		slices.SortFunc(versions, func(a, b *version.Version) int {
			return a.Compare(b)
		})
		versions = slices.Compact(versions)
		for _, v := range versions {
			if v.Installed {
				if v.CurrentUsed {
					prettyout.PrettyInfo(cmd.OutOrStderr(), "* %s\n", v.String())
					continue
				}
				prettyout.PrettyInfo(cmd.OutOrStderr(), " %s\n", v.String())
				continue
			}
			cmd.Printf(" %s\n", v.String())
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("remote", "r", false, "List remote Go versions")
	listCmd.Flags().StringP("type", "t", string(consts.All), "Version type (default all): stable | unstable | archived ")
}
