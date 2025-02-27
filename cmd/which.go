package cmd

import (
	"github.com/11ujjawal/gonix/internal/which"
	"github.com/spf13/cobra"
)

// whichCmd represents the which command
var whichCmd = &cobra.Command{
	Use:   "which program",
	Short: "locate a program file in user's path",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return which.Execute(args[0])
	},
}

func init() {
	rootCmd.AddCommand(whichCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// whichCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// whichCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
