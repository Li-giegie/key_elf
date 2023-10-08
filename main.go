/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"sync"
	"keyelf/cmd"
)

var w sync.WaitGroup

func main() {
	cmd.Execute()
}
