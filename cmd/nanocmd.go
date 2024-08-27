package cmd

import (
	"github.com/spf13/cobra"
)

const configFile string = ".nanocmd"

// nanocmdCmd represents the nanocmd command
var nanocmdCmd = &cobra.Command{
	Use:   "nanocmd",
	Short: "Works with his subcommands.",
	Long: `  
	  nanocmd is only useable with his subcommands. -h to see it.`,
}

func init() {
	rootCmd.AddCommand(nanocmdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nanocmdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nanocmdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
