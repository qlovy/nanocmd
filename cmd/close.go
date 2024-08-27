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
	Short: "To close the file correctly.",
	Long:  `Close the file, so you can open an other one.`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error: Can't open the file !")
			return
		}
		err = file.Close()
		if err != nil {
			fmt.Println("Error: Can't close the file !")
			return
		}
		fmt.Println("The current file has been closed")
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
