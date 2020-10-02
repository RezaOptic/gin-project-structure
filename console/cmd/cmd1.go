package cmd

import "github.com/spf13/cobra"

// rootCmd represents the base command when called without any subcommands
var cmd1 = &cobra.Command{
	Use:   "cmd1",
	Short: "A sample console command",
	Long: `A longer description fo sample console command`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	rootCmd.AddCommand(cmd1)
}
