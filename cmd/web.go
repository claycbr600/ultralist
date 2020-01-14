package cmd

import (
	"github.com/spf13/cobra"
	"github.com/claycbr600/ultralist/ultralist"
)

var (
	webCmdDesc     = "Open your list on ultralist.io"
	webCmdLongDesc = "\nIf your list is synced with ultralist.io, use this command to open your list with your web browser."
)

var webCmd = &cobra.Command{
	Use:   "web",
	Long:  webCmdLongDesc,
	Short: webCmdDesc,
	Run: func(cmd *cobra.Command, args []string) {
		ultralist.NewApp().OpenWeb()
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
}
