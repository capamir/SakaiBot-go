package commands

import (
	"os"

	"github.com/capamir/sakaibot-go/pkg/utils"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sakaibot",
	Short: "🤖 SakaiBot - Advanced Telegram Userbot with AI",
	Long: `SakaiBot is an advanced Telegram userbot with AI capabilities.
Built with Go for performance and reliability.

Use 'sakaibot [command] --help' for more information about a command.`,
	// Don't show usage on every error
	SilenceUsage: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		utils.Error(err.Error())
		os.Exit(1)
	}
}

func init() {
	// Global flags can be added here if needed
	// rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}
