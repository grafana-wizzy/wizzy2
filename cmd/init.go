package cmd

import (
	"github.com/grafana-wizzy/wizzy2/log"
	"github.com/spf13/cobra"

	"github.com/grafana-wizzy/wizzy2/config"
)

// versionCmd represents the version command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize wizzy configuration",
	Long:  `All software has versions. This is generated code example`,
	Run: func(cmd *cobra.Command, args []string) {
		err := config.InitCmd()
		if err != nil {
			log.Error(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
