package cmd

import (
	"fddeploy/upload"

	"github.com/spf13/cobra"
)

var (
	targetName    string
	uploadMessage string

	uploadCmd = &cobra.Command{
		Use:   "upload",
		Short: "upload frontend static files to server",
		Run: func(cmd *cobra.Command, args []string) {
			upload.Run(targetName, uploadMessage)
		},
	}
)

func init() {
	uploadCmd.PersistentFlags().StringVarP(&targetName, "target", "t", "APP", "需部署的应用名")
	uploadCmd.PersistentFlags().StringVarP(&uploadMessage, "message", "m", "", "更新描述")

	rootCmd.AddCommand(uploadCmd)
}
