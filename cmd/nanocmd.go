package cmd

import (
	//"os"
	//"os/exec"
	//"runtime"
	"github.com/spf13/cobra"
	//"strconv"
	//"bufio"
)

const configFile string = ".nanocmd"

// nanocmdCmd represents the nanocmd command
var nanocmdCmd = &cobra.Command{
	Use:   "nanocmd",
	Short: "Call the editor if a file name is given.",
	Long: `Note:
  You have to call nanocmd with a file name. If you don't you can't use the editor.`,
	Run: func(cmd *cobra.Command, args []string) {
		/*const configFile = ".nanocmd"
		filename := ""
		// S'il n'y a pas de nom de fichier
		f_byte, err := os.ReadFile(configFile)
		if err == nil {
			filename = string(f_byte)
		}else if err == nil && len(args) > 0{
			fmt.Println("Error: You need to close the file to launch this command")
			return
		}else if err != nil && len(args) > 0{
			filename = args[0]
		}else{
			fmt.Println("Error: You need a filename to launch the command")
			return
		}

		// Efface la console, si on est sur mac (darwin)
		if runtime.GOOS == "darwin" {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
		file, err := os.Open(filename)
		if err != nil {
			os.Create(configFile)
			err := os.WriteFile(configFile, []byte(filename), 0666)
			if err != nil {
				fmt.Println(err)
			}
		}

		// Affichage du nom du fichier dans l'interface
		filenameLen := len(filename)
		nbOfChar := 80 - filenameLen
		for i := 0; i < nbOfChar / 2; i++ {
			fmt.Print("-")
		}
		fmt.Print(filename)

		for i := 0; i < 80 - nbOfChar/2 - filenameLen; i++ {
			fmt.Print("-")
		}

		fmt.Println()
		scanner := bufio.NewScanner(file)
		lineNo := 0
		for scanner.Scan() {
			lineNo++
			fmt.Printf("%4d. ", lineNo)
			fmt.Println(scanner.Text())
		}

		fmt.Println("--------------------------------------------------------------------------------")
		*/
	},
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
