package main

import (
	"github.com/spf13/cobra"
)

func main() {}

func add() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			return cmd.Help()
		},
	}
}

func rm() *cobra.Command {
	return &cobra.Command{
		Use: "rm",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			return cmd.Help()
		},
	}
}

func verify() *cobra.Command {
	return &cobra.Command{
		Use: "verify",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
}
