package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"strconv"
)

var content, addText, insertMsg, delMsg string


// saveCmd represents the save command
var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fByte, err := os.ReadFile(configFile)
		if err != nil {
			fmt.Println("Error: Can't access to the filename !")
			return
		}
		filename = string(fByte)
		file, err := os.Open(filename)

		if addText != "" {
			file.Write([]byte(addText))
		}
		if insertMsg != "" {
			strLineNo, text, _ := strings.Cut(insertMsg, " ")
			scanner := bufio.NewScanner(file)
			lineNo, err := strconv.Atoi(strLineNo)
			if err != nil {
				fmt.Println("Error: Can't convert string to int !")
				return
			}
			countLine := 0
			for scanner.Scan() {
				countLine++
				if countLine == lineNo {
					arr := scanner.Bytes()
					arrBytes := []byte(text)
					for i := 0; i<len(arrBytes); i++ {
						arr = append(arr, arrBytes[i])
					}
					file.WriteAt(arr, int64(16))
				}
			}
		}
		if delMsg != "" {

		}
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
	writeCmd.Flags().StringVarP(&addText,"append", "a", "textappendbottom", "append text at the bottom of the file")
	writeCmd.Flags().StringVarP(&insertMsg,"insert", "i", "8 sometext", "insert text at a certain line")
	writeCmd.Flags().StringVarP(&delMsg,"delete", "d", "5", "delete text at a certain line or multiple (5 - 8)")
}
