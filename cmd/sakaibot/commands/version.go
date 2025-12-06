package commands

import (
	"fmt"
	"runtime"

	"github.com/capamir/sakaibot-go/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	// These can be set during build with -ldflags
	version   = "0.1.0"
	buildDate = "dev"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "📦 Show version information",
	Long:  `Display the current version of SakaiBot along with build information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		utils.Info("SakaiBot Go 🤖")
		fmt.Println()
		
		fmt.Printf("📦 Version:     %s\n", version)
		fmt.Printf("🗓️  Build Date:  %s\n", buildDate)
		fmt.Printf("⚙️  Go Version:  %s\n", runtime.Version())
		fmt.Printf("💻 OS/Arch:     %s/%s\n", runtime.GOOS, runtime.GOARCH)
		
		fmt.Println()
	},
}

func init() {
	// Register this command with root
	rootCmd.AddCommand(versionCmd)
}
