package cmd

import (
	c "flatpkg/colorize"
	"flatpkg/flathub"
	"flatpkg/utils"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

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
			fmt.Println(hit.Name)
			c.Grey.Println(" ", hit.Summary)
		}

	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.ValidArgsFunction = utils.FlathubCompletions
}
