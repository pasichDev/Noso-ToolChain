package main

import (
	cmds "github.com/pasichDev/nosotc/cmd/nosotc-cli/commands"
)

func main() {
	/*	logo := figure.NewFigure("N-ToolChain", "", false)
		logo.Print()
		fmt.Println("\n\n")

	*/
	cmds.Execute()
}
