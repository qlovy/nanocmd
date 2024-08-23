package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	//"os"
	//"os"
	//"os/exec"
)

var content string

// saveCmd represents the save command
var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("writting i guess")
	},
}

func init() {
	nanocmdCmd.AddCommand(writeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	writeCmd.Flags().StringP("append", "a", "", "append text at the bottom of the file")
	writeCmd.Flags().IntP("insert", "i", 0, "insert text at a certain line")
	writeCmd.Flags().IntP("delete", "d", 0, "delete text at a certain line")
}
