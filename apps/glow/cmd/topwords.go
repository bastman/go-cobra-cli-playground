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
	"strings"

	"github.com/chrislusf/glow/flow"
	"github.com/spf13/cobra"
)

// topwordsCmd represents the topwords command
var topwordsCmd = &cobra.Command{
	Use:   "topwords",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("topwords called")
		runPipeline()
	},
}

func init() {
	rootCmd.AddCommand(topwordsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// topwordsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// topwordsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runPipeline() {
	flow.New().TextFile(
		"/etc/passwd", 3,
	).Filter(func(line string) bool {
		return !strings.HasPrefix(line, "#")
	}).Map(func(line string, ch chan string) {
		for _, token := range strings.Split(line, ":") {
			ch <- token
		}
	}).Map(func(key string) int {
		return 1
	}).Reduce(func(x int, y int) int {
		return x + y
	}).Map(func(x int) {
		println("count:", x)
	}).Run()
}
