package cmd /* Copyright © 2022 Ken Pepple <kpepple@weedmaps.com> */

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var debug bool
var WAP []string
var USERNAME, CERTFILE, INFORMURL, WAPLIST string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "unifi-adopt",
	Short: "Monitor your Ubiqiti WAPs to ensure that stay connected to your Unifi server",
	Long: `Monitor your Ubiqiti WAPs to ensure that stay connected to your Unifi server.
This tool will read a config file (defaults to ~/.unifi-adopt) and query each WAP configured 
by SSHing to it and checking it's config. If it is not connected, it will set the inform-url to your configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		if debug {
			printConfigs()
		}
		// TODO fixme
		w, _ := parseWAP(WAPLIST)
		for x := range w {
			connected, err := checkConnected(w[x], USERNAME, CERTFILE)
			if err != nil {
				// figure out error protocol
				fmt.Println(err)
			}
			if !connected {
				if debug {
					fmt.Printf("%s is not connected: %s", w[x], err)
				}
				informed, err2 := setInform(w[x], USERNAME, CERTFILE, INFORMURL)
				if err2 != nil || !informed {
					fmt.Printf("%s cannot be configured: %v", w[x], err2)
				} else {
					fmt.Printf("%s set to informed URL", w[x])
				}
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.unifi-adopt.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "debug to see all network calls")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".unifi-adopt" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("env")
		viper.SetConfigName(".unifi-adopt")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if debug {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
	}
	WAPLIST = viper.GetString("WAPLIST")
	USERNAME = viper.GetString("USERNAME")
	INFORMURL = viper.GetString("INFORMURL")
	CERTFILE = viper.GetString("CERTFILE")
	if debug {
		viper.Debug()
	}
}

func printConfigs() {
	fmt.Printf("WAPS: %v\nUSERNAME: %v\nINFORMURL: %v\nCERTFILE: %v\n", WAPLIST, USERNAME, INFORMURL, CERTFILE)
}
