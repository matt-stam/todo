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
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long: `Add will create a new todo item to list`,
	Run: addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := data.ReadItems("/Users/Matt Stam/go/src/github.com/matt-stam/todo/todo-data.json");
	if err != nil {
		log.Printf("%v", err);
	}
	for _, x := range args {
		items = append(items, data.Item{Text:x});
	}
	
	err = data.SaveItems("/Users/Matt Stam/go/src/github.com/matt-stam/todo/todo-data.json", items);
	if err != nil {
		fmt.Errorf("%v", err);
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
