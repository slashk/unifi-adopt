package cmd /* Copyright © 2022 Ken Pepple <kpepple@weedmaps.com> */

import (
	"fmt"

	"github.com/spf13/cobra"
)

// variables for version
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the current version",
	Long:  `Prints the current version`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(PrintVersion(version, commit, date))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func PrintVersion(v, c, d string) string {
	return fmt.Sprintf("%v, commit %v, built at %v", v, c, d)
}
