package stringer

import(
	"fmt"
	"os"
	"github.com/spf13/cobra"
)
var rootCmd=&cobra.Command{
	Use: "stringer",
	Short: "stringer - a simple CLI to transform and inspect strings",
	Long : "stringer is a super fancy cli tool ", 

	Run: func(cmd *cobra.Command, args []string) {
},
}
func Execute() {
	if err:=rootCmd.Execute(); err!= nil {
	fmt.Fprintf(os.Stderr, "whoops, there was an erro while executing your CLI '%s'",err)
	os.Exit(1)
	}
}
