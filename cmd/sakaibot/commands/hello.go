package commands

import (
	"github.com/capamir/sakaibot-go/pkg/utils"
	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "👋 Say hello from SakaiBot",
	Long:  `A simple hello command to test that the CLI is working properly.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the name flag value
		name, _ := cmd.Flags().GetString("name")
		
		if name != "" {
			utils.Success("Hello, " + name + "! 👋")
		} else {
			utils.Success("Hello from SakaiBot Go! 🤖")
		}
		
		utils.Info("CLI framework is working perfectly!")
	},
}

func init() {
	// Register this command with root
	rootCmd.AddCommand(helloCmd)
	
	// Add a --name flag (optional)
	helloCmd.Flags().StringP("name", "n", "", "Your name to greet")
}
