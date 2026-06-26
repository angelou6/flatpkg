package cmd

import (
	c "flatpkg/colorize"
	"flatpkg/exit"
	"flatpkg/flathub"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search a package",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			exit.ErrorExit("Package to search needs to be specified.")
		}

		query := strings.Join(args, " ")

		res, err := flathub.Search(query)
		if err != nil {
			exit.ErrorExit(err.Error())
		}

		for _, hit := range res.Hits {
			fmt.Println(hit.Name)
			c.Grey.Println(" ", hit.Summary)
		}

	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
