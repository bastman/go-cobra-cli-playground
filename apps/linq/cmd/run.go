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
	"encoding/json"
	"fmt"
	. "github.com/ahmetb/go-linq"
	"github.com/spf13/cobra"
	"strings"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")
		runPipeline()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runPipeline() {
	fmt.Println("runPipeline called")
	var results []string
	var sentences []string = make([]string, 3)
	sentences = append(sentences, "The fox jumps. over the fence.")

	From(sentences).
		// split sentences to words
		SelectManyT(func(sentence string) Query {
			return From(strings.Split(sentence, " "))
		}).
		// group the words
		GroupByT(
			func(word string) string { return word },
			func(word string) string { return word },
		).
		// order by count
		OrderByDescendingT(func(wordGroup Group) int {
			return len(wordGroup.Group)
		}).
		// order by the word
		ThenByT(func(wordGroup Group) string {
			return wordGroup.Key.(string)
		}).
		Take(5). // take the top 5
		// project the words using the index as rank
		SelectIndexedT(func(index int, wordGroup Group) string {
			return fmt.Sprintf("Rank: #%d, Word: %s, Counts: %d", index+1, wordGroup.Key, len(wordGroup.Group))
		}).
		ToSlice(&results)

	fmt.Println("results: " + strings.Join(results[:], ","))

	json, _ := json.Marshal(results)
	fmt.Println("results (json): " + string(json))

}
