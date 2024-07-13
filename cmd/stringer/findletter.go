package stringer

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var finditCmd = &cobra.Command{
	Use:     "findchar [string] [char]",
	Short:   "This command helps to find if a character is present in a string or not",
	Long:    "This command checks if the specified character is present in the given string.",
	Aliases: []string{"dhundo", "find"},
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		str := args[0]
		char := args[1]

		if len(char) != 1 {
			fmt.Println("The second argument must be a single character.")
			return
		}

		found := strings.ContainsRune(str, []rune(char)[0])
		if found {
			fmt.Printf("Character '%s' is present in the string '%s'.\n", char, str)
		} else {
			fmt.Printf("Character '%s' is not present in the string '%s'.\n", char, str)
		}
	},
}

func init() {
	rootCmd.AddCommand(finditCmd)
}
