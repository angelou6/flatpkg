package cmd

import (
	"flatpkg/exit"
	"flatpkg/flathub"
	"strings"

	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:     "uninstall",
	Aliases: []string{"remove"},
	Short:   "uninstall packages",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			exit.ErrorExit("Package to uninstall needs to be specified.")
		}

		yes, _ := cmd.Flags().GetBool("yes")
		deleteData, _ := cmd.Flags().GetBool("delete-data")
		err := flathub.Remove(strings.Join(args, " "), yes, deleteData)
		if err != nil {
			exit.ErrorExit(err.Error())
		}
	},
}

func init() {
	uninstallCmd.Flags().BoolP("yes", "y", false, "skips user interactions")
	uninstallCmd.Flags().BoolP("delete-data", "d", false, "Deletes application data")
	rootCmd.AddCommand(uninstallCmd)
}
