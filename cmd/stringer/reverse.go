//this code will help to reverse the string given as argument by the user 
package stringer

import(
	"fmt"
	"github.com/faisal/stringer/pkg/stringer"
	"github.com/spf13/cobra"
)
var reverseCmd=&cobra.Command {
	Use: "reverse",
	Aliases : []string{"rev","ultakaro"},
	Short: "this function reverse the given string",
	Long: "this is a long description of the function containing use cases with example",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	res:=stringer.Reverse(args[0])
	fmt.Println(res)
	},
}
	func init() {
		rootCmd.AddCommand(reverseCmd)
}
