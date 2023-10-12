package main

import (
	"fmt"
	cmd2 "github.com/lfcifuentes/go-repository-pattern/app/cmd"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{Use: "repository_pattern_example"}

	rootCmd.AddCommand(
		cmd2.Run,
		cmd2.Migrate,
		cmd2.AppSecret,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
