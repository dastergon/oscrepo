package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/go-ini/ini"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var cfgFile string
var config *ini.File
var CfgUsername string
var CfgPassword string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "oscrepo",
	Short: "A tiny utility to get .repo URLs without openning the browser",
	Long: `oscrepo is a CLI tool that contacts the openSUSE API (api.opensuse.org)
	and retrieves the available projects and the respect .repo URLs from each repository`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Persistent Flags
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.oscrc)")
	RootCmd.PersistentFlags().StringP("username", "u", "", "openSUSE Connect Account Username")
	RootCmd.PersistentFlags().StringP("password", "p", "", "Plaintext openSUSE Connect Account Password")
	RootCmd.PersistentFlags().Int32P("entry", "e", -1, "Repository entry")
}

// initConfig reads in config file if set.
func initConfig() {
	if cfgFile == "" {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		cfgFile = filepath.Join(home, ".oscrc")
	}
	// Read config.
	config, err := ini.Load(cfgFile)
	if err != nil {
		log.Fatalln(err)
	}

	apiSection, err := config.GetSection("https://api.opensuse.org")
	if err != nil {
		log.Fatalln(err)
	}

	CfgUsername = apiSection.Key("user").String()
	CfgPassword = apiSection.Key("pass").String()
}
