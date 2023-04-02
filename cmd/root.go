package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "rpg",
		Short: "reverse proxy service.",
		Long:  `reverse proxy service.`,
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		logrus.Errorln(err)
	}
}

func init() {
	rootCmd.AddCommand(run)
	run.PersistentFlags().StringVarP(&Target, "ip", "i", "0.0.0.0", "target ip")
	err := run.MarkFlagRequired("ip")
	if err != nil {
		logrus.Errorln(err)
	}
	LocalPort = run.Flags().IntSliceP("from", "f", []int{8080}, "--from 80,443")
	run.Flags().IntVarP(&TargetPort, "to", "t", 0, "--to 443")
}
