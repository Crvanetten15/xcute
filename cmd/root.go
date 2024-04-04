package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings" // Import the strings package

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "xcute",
	Short: "List files in the current directory and allow user to select one",
	Long: `xcute is a CLI application that lists all files in the current directory 
and allows the user to select one of them.`,
	Run: func(cmd *cobra.Command, args []string) {
		files, err := ioutil.ReadDir(".")
		if err != nil {
			fmt.Println("Error reading directory:", err)
			os.Exit(1)
		}

		fmt.Println("Select a file:")
		for i, file := range files {
			fmt.Printf("%d: %s\n", i+1, file.Name())
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter number: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		// Trim the input to remove leading and trailing whitespace
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input:", err)
			os.Exit(1)
		}

		if choice < 1 || choice > len(files) {
			fmt.Println("Choice out of range")
			os.Exit(1)
		}

		selectedFile := files[choice-1]
		fmt.Println("You selected:", selectedFile.Name())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %s\n", err)
		os.Exit(1)
	}
}

func init() {
	// Initialization code here
}
