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
	run.PersistentFlags().StringVarP(&Target, "ip", "i", "0.0.0.0", "target ip")
	run.MarkFlagRequired("ip")
	LocalPort = run.Flags().IntSliceP("from", "f", []int{8080}, "--from 80,443")
	run.Flags().IntVarP(&TargetPort, "to", "t", 8080, "--to 443")
}
