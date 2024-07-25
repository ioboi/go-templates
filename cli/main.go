package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type exitCode int

const (
	exitOK    exitCode = 0
	exitError exitCode = 1
)

func mainRun() exitCode {
	root := newCmdRoot()

	ctx := context.Background()
	if err := root.ExecuteContext(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitError
	}

	return exitOK
}

func main() {
	code := mainRun()
	os.Exit(int(code))
}

func newCmdRoot() *cobra.Command {
	var (
		file string = "file"
	)

	cmd := &cobra.Command{
		Use: "cli [-f file]",
		RunE: func(cmd *cobra.Command, args []string) error {
			file, err := os.ReadFile(file)
			if err != nil {
				return err
			}

			fmt.Print(file)

			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&file, "file", "f", file, "A file to read")

	return cmd
}
