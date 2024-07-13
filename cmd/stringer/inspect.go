package stringer
import(
	"fmt"
	"github.com/spf13/cobra"
	"github.com/faisal/stringer/pkg/stringer"
)
var inspectCmd=&cobra.Command {
	Use:"inspect",
	Short:"this command is used to find plurals in the passed string",
	Long:"I cannot think of a long description at this point",
	Aliases:[]string{"ins","findit"},
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		i:=args[0];
		res,kind :=stringer.Inspect(i,false)
		plurals:= "s"
		if res == 1 {
		plurals=""
		}
		fmt.Printf("%s has a %d %s %s.\n",i,res,kind,plurals)
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}
