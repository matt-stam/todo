/*
Copyright © 2019 Matt Stam mattstam@live.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
  "fmt"
  "os"
  "github.com/spf13/cobra"
  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"

)


var cfgFile string

var rootCmd = &cobra.Command{
	Use: "todo",
	Short: "todo is a CLI todo app",
	Long: `Todo will help you get more done in less time. It's designed to make it easy to track, update, and complete your daily tasks.`,
}

func Execute() {
  err := rootCmd.Execute(); 
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)

  // Here you will define your flags and configuration settings.
  // Cobra supports persistent flags, which, if defined here,
  // will be global for your application.

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todo.yaml)")


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
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    // Search config in home directory with name ".todo" (without extension).
    viper.AddConfigPath(home)
    viper.SetConfigName(".todo")
  }

  viper.AutomaticEnv() // read in environment variables that match

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err == nil {
    fmt.Println("Using config file:", viper.ConfigFileUsed())
  }
}

