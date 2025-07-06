package cmd /* Copyright © 2022 Ken Pepple <ken@pepple.io> */

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var debug bool
var pushoverMessage bool
var WAP []string
var UNAME, CERTFILE, INFORMURL, WAPLIST, PUSHOVER_APP_TOKEN, PUSHOVER_USER_KEY string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "unifi-adopt",
	Short: "Monitor your Ubiqiti WAPs to ensure that stay connected to your Unifi server",
	Long: `Monitor your Ubiqiti WAPs to ensure that stay connected to your Unifi server.
This tool will read a config file (defaults to ~/.unifi-adopt) and query each WAP configured
by SSHing to it and set the inform-url to your configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		if debug {
			printConfigs()
		}
		// TODO fixme
		w, err := parseWAP(WAPLIST)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = os.Open(CERTFILE)
		if err != nil {
			fmt.Fprintf(os.Stderr, "SSH certificates unusable: %v", err)
		}
		for x := range w {
			connected, err := checkConnected(w[x], UNAME, CERTFILE)
			if err != nil {
				fmt.Println("ssh connection error: ", err)
				if pushoverMessage {
					if checkPushoverKeys(PUSHOVER_USER_KEY, PUSHOVER_APP_TOKEN) {
						sendPush(fmt.Sprintf("%s cannot be contacted", w[x]), PUSHOVER_USER_KEY, PUSHOVER_APP_TOKEN)
					}
				}
				// continue
			}
			if !connected {
				fmt.Printf("%s is not connected: %s", w[x], err)
				informed, err2 := setInform(w[x], UNAME, CERTFILE, INFORMURL)
				if err2 != nil || !informed {
					fmt.Printf("%s cannot be configured: %v", w[x], err2)
				} else {
					fmt.Printf("%s set to informed URL\n", w[x])
				}
				if pushoverMessage {
					if checkPushoverKeys(PUSHOVER_USER_KEY, PUSHOVER_APP_TOKEN) {
						sendPush(fmt.Sprintf("%s inform URL set", w[x]), PUSHOVER_USER_KEY, PUSHOVER_APP_TOKEN)
					}
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

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "$HOME/.unifi-adopt", "config file")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "debug to see all network calls")
	rootCmd.PersistentFlags().BoolVarP(&pushoverMessage, "pushover", "p", false, "send pushover messages on all actions")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if debug {
		fmt.Printf("\nUsing config file: %s\n", cfgFile)
	}
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		// viper.SetConfigType("yaml")
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		// Search config in home directory with name ".unifi-adopt" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		// viper.SetConfigType("yaml")
		// viper.SetConfigType("env")
		viper.SetConfigName(".unifi-adopt")
	}

	// viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if ok {
			fmt.Fprintln(os.Stderr, "No configuration file found:", viper.ConfigFileUsed())
		}
		_, ok = err.(viper.ConfigMarshalError)
		if ok {
			fmt.Fprintln(os.Stderr, "config marshall error:", viper.ConfigFileUsed())
		}
		_, ok = err.(viper.ConfigParseError)
		if ok {
			fmt.Fprintln(os.Stderr, "config parsing error:", viper.ConfigFileUsed())
		}
		if debug {
			fmt.Println(err)
			viper.Debug()
		}
		// os.Exit(1)
	}
	WAPLIST = viper.GetString("WAPLIST")
	UNAME = viper.GetString("USERNAME")
	INFORMURL = viper.GetString("INFORMURL")
	CERTFILE = viper.GetString("CERTFILE")
	PUSHOVER_APP_TOKEN = viper.GetString("PUSHOVER_APP_TOKEN")
	PUSHOVER_USER_KEY = viper.GetString("PUSHOVER_USER_KEY")
	if pushoverMessage && !checkPushoverKeys(PUSHOVER_USER_KEY, PUSHOVER_APP_TOKEN) {
		fmt.Fprintln(os.Stderr, "No usable pushover keys found. Please create them in ~/.unifi-adopt config file.")
		os.Exit(1)
	}
	if debug {
		viper.Debug()
		fmt.Println("end config setting")
	}
}

func printConfigs() {
	fmt.Printf("WAPS: %v\nUSERNAME: %v\nINFORMURL: %v\nCERTFILE: %v\n", WAPLIST, UNAME, INFORMURL, CERTFILE)
}
