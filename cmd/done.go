/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"log"
	"sort"
	"strconv";
	"github.com/matt-stam/todo/data"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Aliases : []string{"do"},
	Short: "Mark Item as Done",
	Run: doneRun,
}

func doneRun(cmd *cobra.Command, args []string) {
	items, err := data.ReadItems(viper.GetString("datafile"));
	i, err := strconv.Atoi(args[0]);

	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err);
	}

	if i > 0 && i < len(items) {
		items[i-1].Done = true;
		fmt.Printf("%q %v\n", items[i-1].Text, "marked done");
		
		sort.Sort(data.ByPri(items));
		data.SaveItems(viper.GetString("datafile"), items);
	} else {
		log.Println(i, "doesn't match any items");
	}
}

func init() {
	rootCmd.AddCommand(doneCmd);
}
