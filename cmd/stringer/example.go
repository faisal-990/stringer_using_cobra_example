package stringer

import(
	"fmt"
	"github.com/spf13/cobra"
)

var example=&cobra.Command {
	Use:"example",
	Short:"this is used to illustrate the example",
	Aliases:[]string{"show"},
	Run:func(cmd* cobra.Command,args []string){
	fmt.Println("go run main.go [command] [arguments]")
	},
}
func init() {
rootCmd.AddCommand(example)
}
