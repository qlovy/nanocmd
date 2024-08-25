package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"strconv"
)

var addText, insertMsg, delMsg string


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
		file, err := os.OpenFile(filename, os.O_RDWR, 0666)
		// S'il n'existe pas on le crée
		if err != nil {
			file, err = os.Create(filename)
			if err != nil {
				fmt.Printf("Error: Can't create %s !\n", filename)
				return
			}
		}
		// Si le paramètre "append" a été utilisé
		if addText != "" {
			addText = "\n" + addText	// Ajoute une nouvelle ligne
			fileInfo, err := file.Stat()
			if err != nil {
				fmt.Println("Error: Can't acces stat of the file !")
				return
			}
			len_file := fileInfo.Size()	// La longueur du fichier en bytes
			// Ecrit le message à la fin du fichier
			file.WriteAt([]byte(addText), int64(len_file))
		// Si le paramètre "insert" a été utilisé
		}else if insertMsg != "" {
			// Coupe le message en deux, la ligne et le texte
			strLineNo, text, _ := strings.Cut(insertMsg, " ")
			scanner := bufio.NewScanner(file)
			lineNo, err := strconv.Atoi(strLineNo)	// Convertit de chaîne de carctère en nombre 
			// S'il y a une erreure
			if err != nil {
				fmt.Println("Error: Can't convert string to int !")
				return
			}
			countLine := 0
			content := ""
			// Scanne le fichier ligne par ligne
			for scanner.Scan() {
				countLine++
				content += scanner.Text()
				if countLine == lineNo {
					content += text
				}
				content += "\n"
			}
			err = os.WriteFile(filename, []byte(content), 0666)
			if err != nil {
				fmt.Printf("Error: Can't write in %s !\n", filename)
				return
			}
		// Si le paramètre "delete" a été utilisé
		}else if delMsg != "" {
			//i := strings.Index(delMsg, "-")
			before, after, _ := strings.Cut(delMsg, "-")
			linesNo := []int{}
			if before != "" && after != "" {
				n, err := strconv.Atoi(before)
				if err != nil {
					fmt.Println("Error: Can't convert non-integer in integer")
					return
				}
				linesNo = append(linesNo, n)
				n, err = strconv.Atoi(after)
				if err != nil {
					fmt.Println("Error: Can't convert non-integer in integer")
					return
				}
				linesNo = append(linesNo, n)
			}else{
				n, err := strconv.Atoi(before)
				if err != nil {
					fmt.Println("Error: Can't convert non-integer in integer")
					return
				}
				linesNo = append(linesNo, n)
			}
			fmt.Println(linesNo)
			scanner := bufio.NewScanner(file)
			countLine := 0
			content := ""
			for scanner.Scan() {
				countLine++
				if len(linesNo) == 2 {
					if countLine < linesNo[0] || countLine > linesNo[1] {
						content += scanner.Text() + "\n"
					}
				}else{
					if countLine != linesNo[0] {
						content += scanner.Text() + "\n"
					}
				}
			}
			err := os.WriteFile(filename, []byte(content), 0666)
			if err != nil {
				fmt.Printf("Error: Can't write in %s !\n", filename)
				return
			}
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
	writeCmd.Flags().StringVarP(&addText,"append", "a", "", "append text at the bottom of the file")
	writeCmd.Flags().StringVarP(&insertMsg,"insert", "i", "", "insert text at a certain line (\"line text\")")
	writeCmd.Flags().StringVarP(&delMsg,"delete", "d", "", "delete text at a certain line or multiple (5 - 8)")
}
