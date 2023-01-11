package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "rpg",
		Short: "reverse proxy service.",
		Long:  `reverse proxy service.`,
	}
)

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(run)
	run.PersistentFlags().StringVarP(&Target, "target", "t", "", "target host")
	run.MarkFlagRequired("target")
	run.Flags().IntVarP(&LocalPort, "port", "p", 8080, "local port")
	run.Flags().IntVarP(&TargetPort, "targetPort", "i", 8080, "target port")
}
