package cmd

import (
	"github.com/desertbit/grumble"
	"github.com/fatih/color"
)

var App = grumble.New(&grumble.Config{
	Name:                  "MetaTOC",
	Description:           "MetaTOC is an open source project under the data acquisition and model based on edge computing, and it is the best practice for data decentralized authentication and tracking with blockchain violas.",
	Prompt:                "MetaTOC » ",
	PromptColor:           color.New(color.FgGreen, color.Bold),
	HelpHeadlineColor:     color.New(color.FgGreen),
	HelpHeadlineUnderline: true,
	HelpSubCommands:       true,
})

func init() {
	App.SetPrintASCIILogo(func(a *grumble.App) {
		a.Println()
		a.Println("         _____   _____")
		a.Println(" _____ _|_   _|_|_   _|_ ___")
		a.Println("|     | -_| |  _  | | . |  _|")
		a.Println("|_|_|_|___|_|_| |_|_|___|___|")
		a.Println()
	})
}
