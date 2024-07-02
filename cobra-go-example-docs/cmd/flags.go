//	var cmd = &cobra.Command{
//		Short: "hello",
//		Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
//		Run: func(cmd *cobra.Command, args []string) {
//			fmt.Println("Hello, World!")
//		},
//	}
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(flagsCmd)
}

func validationLogic(test string) bool {
	fmt.Println("test data is", test)
	return test == "exit"
}

var flagsCmd = &cobra.Command{
	Use:   "flags",
	Short: "flags",
	Args: func(cmd *cobra.Command, args []string) error {
		// Optionally run one of the validators provided by cobra
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		// Run the custom validation logic
		if validationLogic(args[0]) {
			return nil
		}
		return fmt.Errorf("invalid color specified: %s", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!")
	},
}
