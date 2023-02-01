package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitter/internal/config"
	"gopkg.in/yaml.v3"

	"github.com/cwdot/go-stdlib/wood"
)

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configPrintCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configuration management; prints location with no arguments",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		printConfLocation()
	},
}

var configPrintCmd = &cobra.Command{
	Use:   "print",
	Short: "Print current repo configuration",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.DefaultConfigFile()
		if err != nil {
			wood.Fatal(err)
		}
		jsonStr, err := yaml.Marshal(conf)
		fmt.Println(string(jsonStr))
	},
}

func printConfLocation() {
	conf, err := config.DefaultConfigFile()
	if err != nil {
		wood.Fatal(err)
	}
	fmt.Printf("Config: %v\n", conf.Location)
}
