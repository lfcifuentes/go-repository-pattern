package main

import (
	"fmt"
	"github.com/lfcifuentes/go-repository-pattern/cmd"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{Use: "repository_pattern_example"}

	rootCmd.AddCommand(
		cmd.Run,
		cmd.Migrate,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
