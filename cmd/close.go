package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	//"os"
	//"os/exec"
)

// saveCmd represents the save command
var closeCmd = &cobra.Command{
	Use:   "close",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		os.WriteFile(configFile, []byte(""), 0666)
		fmt.Println("The current has been closed")
	},
}

func init() {
	nanocmdCmd.AddCommand(closeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

}
