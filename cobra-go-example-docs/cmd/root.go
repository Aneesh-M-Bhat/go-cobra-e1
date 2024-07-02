package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go.
				  Complete documentation is available at https://gohugo.io/documentation/`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var (
	author string
	Region string
	u      string
	pw     string
	ofJson bool
	ofYaml bool
)

func init() {
	rootCmd.PersistentFlags().StringVar(&author, "author", "YOUR NAME", "Author name for copyright attribution")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	// rootCmd.Flags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
	// rootCmd.MarkFlagRequired("region")

	rootCmd.PersistentFlags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
	rootCmd.MarkPersistentFlagRequired("region")

	rootCmd.Flags().StringVarP(&u, "username", "u", "", "Username (required if password is set)")
	rootCmd.Flags().StringVarP(&pw, "password", "p", "", "Password (required if username is set)")
	rootCmd.MarkFlagsRequiredTogether("username", "password")

	// rootCmd.Flags().BoolVar(&ofJson, "json", false, "Output in JSON")
	// rootCmd.Flags().BoolVar(&ofYaml, "yaml", false, "Output in YAML")
	// rootCmd.MarkFlagsMutuallyExclusive("json", "yaml")

	rootCmd.Flags().BoolVar(&ofJson, "json", false, "Output in JSON")
	rootCmd.Flags().BoolVar(&ofYaml, "yaml", false, "Output in YAML")
	rootCmd.MarkFlagsOneRequired("json", "yaml")
	rootCmd.MarkFlagsMutuallyExclusive("json", "yaml")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
