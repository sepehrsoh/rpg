package cmd

import (
	"github.com/spf13/cobra"
	"rpg/proxy"
)

var (
	Target     string
	LocalPort  int
	TargetPort int
	run        = &cobra.Command{
		Use: "run [src]",
		Run: func(cmd *cobra.Command, args []string) {
			reverseArgs := proxy.NewReversArgs(LocalPort, TargetPort, Target)
			proxy.ReverseProxy(reverseArgs)
		},
	}
)
