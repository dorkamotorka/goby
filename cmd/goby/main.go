package main

import (
	"log"
	"os"

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
	cmd := &cobra.Command{
		Use:   "goby",
		Short: "Generate eBPF Golang project.",
	}

	cmd.AddCommand(initCmd())
	return cmd
}

func initCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init <path>",
		Short: "Initialize an eBPF Golang project at the specified path.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			if err := run(path); err != nil {
				log.Fatalf("failed to initialize project: %+v", err)
			}
		},
	}
	return cmd
}

func run(path string) error {
	// Create output directory if it doesn't exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}

	err := generator.GenerateGoMain(path)
	if err != nil {
		return err
	}

	err = generator.GenerateeBPFProgram(path)
	if err != nil {
		return err
	}

	err = generator.DumpBTF(path)
	if err != nil {
		return err
	}

	err = generator.DumpMake(path)
	if err != nil {
		return err
	}

	return nil
}
