package cmd

import (
    "fmt"

    "github.com/common-nighthawk/go-figure"
    "github.com/spf13/cobra"

    "github.com/autobots/touchbase/constants"
    log "github.com/autobots/touchbase/logger"
    "github.com/autobots/touchbase/types"
    "github.com/autobots/touchbase/validations"
)

const (
    fontName = "big"
)

var (
    debugMode bool
)

func generateBanner(phrase string) string {
    return figure.NewFigure(phrase, fontName, true).String()
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "touchbase",
    Short: "Connect with people and share your profile!",
    Long: fmt.Sprintf(`%sTouchbase helps to connect with people, share your profile with short 
description about yourself and your resume/portfolio!`, generateBanner(constants.AppName)),
}

// Execute adds all child commands to the root command and sets flags
// appropriately. This is called by main.main(). It only needs to happen once
// to the rootCmd.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        getLogger().Fatal("Error in executing command", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
    }
}

func init() {
    cobra.OnInitialize(initializeLogging)

    // Add sub commands
    rootCmd.AddCommand(configCmd)
    rootCmd.AddCommand(reachOutCmd)

    // Define your flags
    rootCmd.PersistentFlags().BoolVarP(&debugMode, "debug", "X", false, "Enable debug mode (default false)")

    // Init validations
    validations.InitValidator()
}

func initializeLogging() {
    initLogging(constants.ConsoleFormat, debugMode)
}
