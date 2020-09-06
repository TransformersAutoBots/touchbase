package cmd

import (
    "fmt"
    "github.com/common-nighthawk/go-figure"
    "github.com/spf13/cobra"
    "os"
)

const (
    fontName = "big"
    appName  = "touchbase"
)

func generateBanner(phrase string) string {
    return figure.NewFigure(phrase, fontName, true).String()
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "touchbase",
    Short: "Connect with people and send your profile!",
    Long:  fmt.Sprintf("%sTouchbase will help you connect with people, share your profile with short description about yourself and attach your resume/portfolio!", generateBanner(appName)),
}

// Execute adds all child commands to the root command and sets flags
// appropriately. This is called by main.main(). It only needs to happen once
// to the rootCmd.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
