/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"keyelf/engine"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "创建一个配置文件模板",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		name := "./conf.yaml"
		if len(args) > 0 {
			name = args[0]
		}
		err := engine.StartGenerateYamlTemplateFile(name)
		if err != nil {
			fmt.Println("生成失败：", err)
		}
	},
}
