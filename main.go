package main

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
    "github.com/AbdulSayyed/go-lib-examp/internal/modc"
)

func main() {
    // Create a new Cobra root command
    var rootCmd = &cobra.Command{
        Use:   "greet",
        Short: "Greets from all modules",
        Long:  `This command greets from Module A, Module B, and Module C.`,
        Run: func(cmd *cobra.Command, args []string) {
            // This is where we call the function from modc
            fmt.Println(modc.SayHelloC())
        },
    }

    // Execute the Cobra command
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

