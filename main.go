package main

import (
	"github.com/bluemanos/telemetry-udp-server/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmd.ServerCmd)
	rootCmd.Execute()
}
