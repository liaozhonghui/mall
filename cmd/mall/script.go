package mall

import (
	"fmt"

	"github.com/spf13/cobra"
)

var scriptCmd = &cobra.Command{
	Use:   "script",
	Short: "Run database scripts",
	Long:  `Run database scripts for mall application`,
	Run:   runScripts,
}

func init() {
	rootCmd.AddCommand(scriptCmd)
	scriptCmd.Flags().StringVarP(&config, "config", "c", "", "config file path")
}

func runScripts(cmd *cobra.Command, args []string) {
	fmt.Printf("Run Script")

	fmt.Printf("Script executed successfully")
}
