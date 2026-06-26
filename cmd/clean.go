package cmd

import (
	c "flatpkg/colorize"
	"flatpkg/exit"
	"flatpkg/flathub"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/spf13/cobra"
)

var (
	home, _    = os.UserHomeDir()
	flatpakDir = filepath.Join(home, ".var/app")
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean all uninstalled application data",
	Run: func(cmd *cobra.Command, args []string) {
		ids, err := flathub.GetInstalledIds()
		if err != nil {
			exit.ErrorExit(err.Error())
		}

		dirs, err := os.ReadDir(flatpakDir)
		if err != nil {
			exit.ErrorExit(err.Error())
		}

		for _, dir := range dirs {
			if !slices.Contains(ids, dir.Name()) {
				name := dir.Name()
				err := os.RemoveAll(filepath.Join(flatpakDir, name))
				if err != nil {
					exit.ErrorExit(fmt.Sprintf("Could not delete data for %s", name))
				}

				c.Green.Print("[Deleted] ")
				fmt.Println(dir.Name())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
