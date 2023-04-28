/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"lictl/tools/osx"
	"strings"

	"github.com/spf13/cobra"
)

// osCmd represents the os command
var osCmd = &cobra.Command{
	Use:   "os",
	Short: "For file operations",
	Long:  `Learn how Golang utilizes OS packages to copy, move, rename, delete, and other operations on files`,
}

// cpCmd represents the os command
var cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "copy file",
	Long:  `copy file example lictl os cp [source file]->[target file] `,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			// arg 字符串分割
			arr := strings.Split(arg, "->")
			if len(arr) != 2 {
				fmt.Printf("\"%s\" does not meet the defined format len %d\n", arg, len(arr))
				continue
			}
			err := osx.Copy(arr[0], arr[1])
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				continue
			}
			fmt.Printf("copy %s -> %s success \n", arr[0], arr[1])
		}
	},
}

// mvCmd represents the os command
var mvCmd = &cobra.Command{
	Use:   "mv",
	Short: "move file",
	Long:  `move file example lictl os mv [source file]->[target file] `,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			// arg 字符串分割
			arr := strings.Split(arg, "->")
			if len(arr) != 2 {
				fmt.Printf("\"%s\" does not meet the defined format len %d\n", arg, len(arr))
				continue
			}
			err := osx.Move(arr[0], arr[1])
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				continue
			}
			fmt.Printf("move %s -> %s success \n", arr[0], arr[1])
		}
	},
}

// rnCmd represents the os command
var rnCmd = &cobra.Command{
	Use:   "rn",
	Short: "rename file",
	Long:  `rename file example lictl os rn [source file]->[target file] `,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			// arg 字符串分割
			arr := strings.Split(arg, "->")
			if len(arr) != 2 {
				fmt.Printf("\"%s\" does not meet the defined format len %d\n", arg, len(arr))
				continue
			}
			err := osx.Rename(arr[0], arr[1])
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				continue
			}
			fmt.Printf("rename %s -> %s success \n", arr[0], arr[1])
		}
	},
}

// rmCmd represents the os command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove file",
	Long:  `remove file example lictl os rm [source file]->[target file] `,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			err := osx.Remove(arg)
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				continue
			}
			fmt.Printf("remove %s success \n", arg)
		}
	},
}

func init() {
	osCmd.AddCommand(cpCmd, mvCmd, rnCmd, rmCmd)
	rootCmd.AddCommand(osCmd)
}
