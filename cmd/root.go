/*
Copyright © 2023 fuming
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"gofinger/core/banner"
	"gofinger/core/options"
	"gofinger/core/utils"
	"gofinger/runner"
	"log"
	"os"
	"time"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gofinger",
	Short:   "一款指纹识别工具",
	Version: "0.9",
	Long:    banner.Banner,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(banner.Banner)
		if len(url) == 0 && len(file) == 0 && !stdin {
			return
		}
		log.Println("start fingerprint rule matching ...")
		start := time.Now()
		var urls []string
		if url != "" {
			urls = append(urls, url)
		}
		if file != "" {
			lines, err := utils.ReadLines(file)
			lines = utils.DeduplicateEmptyStrings(lines)
			if err != nil {
				log.Fatalln(err)
			}
			for _, line := range lines {
				urls = append(urls, line)
			}
		}
		if stdin {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Split(bufio.ScanLines)
			for scanner.Scan() {
				urls = append(urls, scanner.Text())
			}
		}
		option := &options.Options{
			Output: output,
			Thread: thread,
			Proxy:  proxy,
			Level:  level,
			Urls:   urls,
		}
		runner := runner.NewRunner(option)
		runner.RunEnumeration()
		elapsed := time.Since(start)
		log.Printf("this time it takes %v .", elapsed)
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

var (
	url, file, proxy, output string
	thread, level            int
	stdin                    bool
)

func init() {
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "-u https://www.baidu.com")
	rootCmd.Flags().StringVarP(&file, "file", "f", "", "-f targets.txt")
	rootCmd.Flags().StringVarP(&output, "output", "o", "", "-o results.csv")
	rootCmd.Flags().IntVarP(&thread, "thread", "t", 50, "-t 25")
	rootCmd.Flags().StringVarP(&proxy, "proxy", "p", "", "-p http://127.0.0.1:8080")
	rootCmd.Flags().IntVarP(&level, "level", "l", 1, "-l 1-3")
	rootCmd.Flags().BoolVarP(&stdin, "stdin", "", false, "--stdin true")
}
