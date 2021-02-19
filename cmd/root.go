/*
 * Copyright 2021 Teo Mrnjavac <teo.mrnjavac@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

// Package cmd contains all the entry points for command line
// subcommands, following library convention.
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/teo/zcascade/app"
	"github.com/teo/zcascade/common/logger"
)

var log = logger.New(logrus.StandardLogger(), "zcascade")

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   app.NAME,
	Short: app.PRETTY_SHORTNAME,
	Long: fmt.Sprintf(`%s is a command line program for fast, large scale ZFS dataset replication.

The following options are always available with any %s command.
For more information on the available commands, see the individual documentation for each command.`, app.PRETTY_SHORTNAME, app.PRETTY_SHORTNAME),
}

func GetRootCmd() *cobra.Command { // Used for docs generator
	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.WithField("error", err).Fatal("cannot run command")
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	viper.Set("version", app.VERSION)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("optional configuration file for %s (default $HOME/.config/%s/settings.yaml)", app.NAME, app.NAME))
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "show verbose output for debug purposes")

	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetDefault("log.level", "info")
	viper.SetDefault("verbose", false)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.WithField("error", err).Error("cannot find configuration file")
			os.Exit(1)
		}

		viper.AddConfigPath(path.Join(home, ".config/" + app.NAME))
		viper.SetConfigName("settings")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logLevel, err := logrus.ParseLevel(viper.GetString("log.level"))
		if err == nil {
			logrus.SetLevel(logLevel)
		}
		log.WithField("file", viper.ConfigFileUsed()).
			Debug("configuration loaded")
	}

	if viper.GetBool("verbose") {
		viper.Set("log.level", "debug")
		logrus.SetLevel(logrus.DebugLevel)
	}
}
