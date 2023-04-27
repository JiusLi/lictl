/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//wd, err := os.Getwd()
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}
		//fmt.Println(wd)
		//
		//err = os.Mkdir("aaa", os.ModePerm)
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}
		//err = os.Chdir(wd + "/aaa")
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}
		//err = os.Mkdir("bbb", os.ModePerm)
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}

		data := []byte("告诉小卢，他就是个垃圾~")

		os.WriteFile("root.go", data, 0644)

		fmt.Println("File created successfully.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
