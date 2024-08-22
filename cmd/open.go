package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"runtime"
)

var filename string

const configFile string = ".nanocmd"

// saveCmd represents the save command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Efface la console, si on est sur mac (darwin)
		if runtime.GOOS == "darwin" {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			// Si on est sur Windows
		} else if runtime.GOOS == "windows" {
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
		fBytes, err := os.ReadFile(configFile)
		fStr := string(fBytes)
		if err != nil {
			_, err1 := os.Create(configFile)
			if err1 != nil {
				return
			}
		} else if filename == "" && fStr != "" {
			filename = fStr
		} else if filename != "" && fStr == "" {
			err2 := os.WriteFile(configFile, []byte(filename), 0666)
			if err2 != nil {
				return
			}
		} else if filename != fStr {
			fmt.Printf("Error: You should close %s before openning %s", filename, fStr)
		}
	},
}

func init() {
	nanocmdCmd.AddCommand(openCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	openCmd.Flags().StringVarP(&filename, "filename", "f", "", "Filename you want to edit")
}
