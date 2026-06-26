package cmd

import (
	"flatpkg/flathub"
	"flatpkg/utils"
	"strings"

	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:     "uninstall",
	Aliases: []string{"remove"},
	Short:   "uninstall packages",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			utils.ErrorExit("Package to uninstall needs to be specified.")
		}

		yes, _ := cmd.Flags().GetBool("yes")
		deleteData, _ := cmd.Flags().GetBool("delete-data")
		err := flathub.Remove(strings.Join(args, " "), yes, deleteData)
		if err != nil {
			utils.ErrorExit(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
	uninstallCmd.Flags().BoolP("yes", "y", false, "skips user interactions")
	uninstallCmd.Flags().BoolP("delete-data", "d", false, "Deletes application data")
	uninstallCmd.ValidArgsFunction = utils.InstalledCompletion
}
