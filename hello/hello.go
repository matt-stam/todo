package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use: "todo"
	Short: "todo is a CLI todo app"
	Long: `Todo will help you get more done in less time. It's designed to make it easy to track, update, and complete your daily tasks.`
}

func main() {
	fmt.Println("Hello, world.")
}
