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
	"os"
	"sort"
	"github.com/matt-stam/todo/data"
	"github.com/spf13/cobra"
	"text/tabwriter"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all todos",
	Long: `list all of the current todos you've added`,
	Run: listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := data.ReadItems(dataFile)
	if err != nil {
		log.Printf("%v", err);
	}

	sort.Sort(data.ByPri(items));

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0);
	for _, i := range items {
		if allOpt || i.Done == doneOpt {
			fmt.Fprintln(w, i.Label()+"\t"+i.PrettyDone()+"\t"+i.PrettyP()+"\t"+i.Text+"\t");
		}
	}

	w.Flush();
}

var doneOpt bool;
var allOpt bool;

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'Done' Todos");
	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show all Todos");
}
