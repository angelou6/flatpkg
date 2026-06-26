package utils

import (
	"flatpkg/flathub"

	"github.com/spf13/cobra"
)

func FlathubCompletions(
	cmd *cobra.Command,
	args []string,
	toComplete string,
) ([]string, cobra.ShellCompDirective) {
	if len(args) > 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	res, err := flathub.Search(toComplete)
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}

	completions := make([]string, 0, len(res.Hits))
	for _, hit := range res.Hits {
		completions = append(completions, hit.Name)
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

func InstalledCompletion(
	cmd *cobra.Command,
	args []string,
	toComplete string,
) ([]string, cobra.ShellCompDirective) {
	if len(args) > 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	installed, err := flathub.GetInstalledNames()
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}

	completions := make([]string, 0, len(installed))
	for _, pkg := range installed {
		completions = append(completions, pkg)
	}

	return completions, cobra.ShellCompDirectiveNoFileComp

}
