/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"keyelf/engine"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "keyelf",
	Short:   "windows平台工具：按键行为管理",
	Long:    `当触发了某个已经侦听的键时，执行一组命令，在配置文件中（conf.yaml）进行定义`,
	Version: "2023.10.8:8.58",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if err := engine.StartTaskManager(Config); err != nil {
			log.Fatalln(err)
		}
		select {}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//reader := bufio.NewReader(os.Stdin)
	//for {
	//	fmt.Print(">>")
	//	//进入命令行模式
	//	text, err := reader.ReadString('\n')
	//	if err != nil {
	//		fmt.Printf("输入格式错误：%v\n",err)
	//		continue
	//	}
	//	args:=strings.Fields(text)
	//	if len(args) < 1 {
	//		continue
	//	}
	//	rootCmd.SetArgs(args)
	//
	//	if err = rootCmd.Execute();err != nil {
	//		fmt.Printf("err：%v\n",err)
	//		continue
	//	}
	//}
}

var Config string

func init() {
	rootCmd.AddCommand(createCmd)
	rootCmd.Flags().StringVarP(&Config, "config", "c", "./conf.yaml", "指定配置文件启动")
}
