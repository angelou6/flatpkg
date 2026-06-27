package cmd

import (
	c "flatpkg/colorize"
	"flatpkg/flathub"
	"flatpkg/utils"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func searchCompletion(
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
		completions = append(completions, cobra.CompletionWithDesc(hit.Name, hit.Id))
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search a package",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			utils.ErrorExit("Package to search needs to be specified.")
		}

		query := strings.Join(args, " ")

		res, err := flathub.Search(query)
		if err != nil {
			utils.ErrorExit(err.Error())
		}

		for _, hit := range res.Hits {
			fmt.Printf("%s (%s)\n", hit.Name, hit.Id)
			c.Grey.Println(" ", hit.Summary)
		}

	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.ValidArgsFunction = searchCompletion
}
