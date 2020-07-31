package main

import (
	"fmt"
	"os"

	"github.com/rp185145/goprojectlayout"
	"github.com/rp185145/goprojectlayout/emitter"
	"github.com/rp185145/goprojectlayout/idutils"
	"github.com/spf13/cobra"
)

const greeting = "greeting"

var exitCode int = 0

var rootCmd = &cobra.Command{
	Use:   "gplcli",
	Short: "Go Project Layout Cobra Demo",
}

var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "Execute package goprojectlayout (root)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("demo")
		goprojectlayout.Demo()
	},
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Execute package emitter from the world",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello")
		emitter.Emitter("Hello World")
	},
}

var howdyCmd = &cobra.Command{
	Use:   "howdy",
	Short: "Execute package emitter from Texas",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("howdy")
		emitter.Emitter("Howdy Pardner")
	},
}

var customCmd = &cobra.Command{
	Use:   "custom",
	Short: "Execute package emitter using a greeting option",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("custom")

		if greeting, err := cmd.Flags().GetString(greeting); err != nil {
			fmt.Fprintf(os.Stderr, "Error processing the %s option: %s\n", greeting, err)
			exitCode = 1
			return
		}

		emitter.Emitter(greeting)
	},
}

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Execute package uuid",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("uuid")
		fmt.Println(idutils.GetUUID())
	},
}

func init() {
	rootCmd.AddCommand(demoCmd)
	rootCmd.AddCommand(helloCmd)
	rootCmd.AddCommand(howdyCmd)
	rootCmd.AddCommand(uuidCmd)

	customCmd.Flags().StringP(greeting, "g", "", "Greeting to emit")
	customCmd.MarkFlagRequired(greeting)
	rootCmd.AddCommand(customCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		exitCode = 1
	}

	os.Exit(exitCode)
}
