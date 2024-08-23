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
		fByte, err := os.ReadFile(configFile)
		if err != nil {
			return
		}
		fmt.Println(string(fByte))
		/*
			err = os.WriteFile(fStr, []byte(content), 0666)
			if err != nil {
				fmt.Println("Error: Can't close the file !")
				return
			}

		*/
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
