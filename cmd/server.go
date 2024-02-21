package cmd

import (
	"fmt"
	"github.com/bluemanos/telemetry-udp-server/cmd/app/forzam8"
	"github.com/spf13/cobra"
)

var (
	ServerCmd = &cobra.Command{
		Use:   "server",
		Short: "Simple UDP server for Forza Motorsport 8 telemetry data",
		Run:   run,
	}

	flagFrequency    int
	flagAddress      string
	flagInfinityLoop bool
)

func init() {
	ServerCmd.Flags().IntVarP(&flagFrequency, "frequency", "f", 1, "frequency (in Hz) of sending data")
	ServerCmd.Flags().StringVarP(&flagAddress, "address", "a", "localhost:20777", "address of the client")
	ServerCmd.Flags().BoolVarP(&flagInfinityLoop, "infinity", "i", false, "Infinity loop?")
	ServerCmd.MarkFlagRequired("address")
}

func run(cmd *cobra.Command, args []string) {
	valid, errors := validateFlags()
	if !valid {
		for _, err := range errors {
			fmt.Println(err)
		}
		return
	}

	for {
		forzam8.Run(flagAddress, flagFrequency)
		if !flagInfinityLoop {
			break
		}
	}
}

func validateFlags() (bool, []error) {
	valid := true
	var errors []error
	if flagFrequency < 1 {
		valid = false
		errors = append(errors, fmt.Errorf("frequency must be greater than 0"))
	}

	return valid, errors
}
