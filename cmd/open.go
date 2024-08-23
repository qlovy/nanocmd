package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"runtime"
)

var filename string

// openCmd represent the Open command
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
		// Lis le fichier .nanocmd
		fBytes, err := os.ReadFile(configFile)
		fStr := string(fBytes) // Transforme en chaîne de carctère
		// S'il y a une erreur (donc le fichier n'existe pas)
		if err != nil {
			// Création du fichier
			_, err = os.Create(configFile)
			// S'il y a une erreur, on quitte
			if err != nil {

				fmt.Println("Error: Can't create the config file !")
				return

			}
			// S'il n'y a pas de nom de fichier donné mais qu'il est actuellement ouvert (le .nanocmd contient quelque chose)
		} else if filename == "" && fStr != "" {
			filename = fStr
			// Si un nom de fichier est donné et qu'il n'y a pas de fichier ouvert
		} else if filename != "" && fStr == "" {
			// Store the name of the file
			err = os.WriteFile(configFile, []byte(filename), 0666)
			if err != nil {

				fmt.Println("Error: Can't write in the config file !")
				return

			}
			// Si le nom de fichier donné n'est pas le même que celui qui est ouvert
		} else if filename != fStr {
			fmt.Printf("Error: You should close %s before openning %s", filename, fStr)
			return
		}

		file, err := os.Open(filename)

		// Affichage du nom du fichier dans l'interface
		filenameLen := len(filename)
		nbOfChar := 80 - filenameLen
		for i := 0; i < nbOfChar/2; i++ {
			fmt.Print("-")
		}
		fmt.Print(filename)

		for i := 0; i < 80-nbOfChar/2-filenameLen; i++ {
			fmt.Print("-")
		}

		fmt.Println()
		scanner := bufio.NewScanner(file)
		lineNo := 0
		for scanner.Scan() {
			lineNo++
			fmt.Printf("%4d ", lineNo)
			fmt.Println(scanner.Text())
		}

		fmt.Println("--------------------------------------------------------------------------------")
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
