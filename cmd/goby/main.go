package main

import (
	"log"
	"github.com/spf13/cobra"

	"github.com/dorkamotorka/goby/internal/generator"
)

func main() {
	err := rootCmd().Execute()
	if err != nil {
		log.Fatalf("failed to run command: %+v", err)
	}
}

func rootCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "goby",
		Short: "Generate eBPF Golang project.",
		Run: func(_ *cobra.Command, _ []string) {
			err := run()
			if err != nil {
				//nolint:revive
				log.Fatalf("run goby: %+v", err)
			}
		},
	}

	return cmd
}

func run() error {
	generator.GenerateGoMain()
	generator.GenerateeBPFProgram()
	generator.DumpBTF()

	return nil
}
