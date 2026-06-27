package cmd

import (
	"flatpkg/flathub"
	"flatpkg/utils"

	"github.com/spf13/cobra"
)

func installCompletion(
	cmd *cobra.Command,
	args []string,
	toComplete string,
) ([]string, cobra.ShellCompDirective) {
	res, err := flathub.Search(toComplete)
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}

	completions := []cobra.Completion{}
	for _, hit := range res.Hits {
		completions = append(completions, cobra.CompletionWithDesc(hit.Id, hit.Name))
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install packages",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			utils.ErrorExit("Package to install needs to be specified.")
		}

		yes, _ := cmd.Flags().GetBool("yes")
		var packages []string
		for _, arg := range args {
			res, err := flathub.Search(arg)
			if err != nil {
				utils.ErrorExit(err.Error())
			}
			packages = append(packages, res.Hits[0].Id)
		}

		err := flathub.Install(packages, yes)
		if err != nil {
			utils.ErrorExit(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().BoolP("yes", "y", false, "skips user interactions")
	installCmd.ValidArgsFunction = installCompletion
}
