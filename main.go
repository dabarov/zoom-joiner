package main

import (
    "fmt"
    "os"
    "github.com/dabarov/zoom-joiner/commands"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "zoom-joiner",
    Short: "A CLI to save and join Zoom meetings by alias",
    Long:  `A simple CLI application to save Zoom meeting data and join them using an alias.`,
}

func main() {
    rootCmd.AddCommand(commands.AddCmd)
    rootCmd.AddCommand(commands.JoinCmd)

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

