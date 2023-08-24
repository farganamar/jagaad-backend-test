package main

import (
	"fmt"
	fetchcmd "jagaad-backend-test/cmd/fetch"
	readcmd "jagaad-backend-test/cmd/read"
	"jagaad-backend-test/config"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Long: "This is Jagaat Backend Test"}

	rootCmd.AddCommand(fetchcmd.Command())
	rootCmd.AddCommand(readcmd.Command())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	config.Init()
}
