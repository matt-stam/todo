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
	"log"
	"github.com/matt-stam/todo/data"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long: `Add will create a new todo item to list`,
	Run: addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := data.ReadItems(viper.GetString("datafile"));
	if err != nil {
		log.Printf("%v", err);
	}
	for _, x := range args {
		item := data.Item{Text: x};
		item.SetPriority(priority);
		items = append(items, item);
	}
	
	err = data.SaveItems(viper.GetString("datafile"), items);
	if err != nil {
		fmt.Errorf("%v", err);
	}
}

var priority int;

func init() {
	rootCmd.AddCommand(addCmd)
	
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority: 1, 2, 3");
}
