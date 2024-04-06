package colors

import "github.com/fatih/color"

var (
	WhiteColor   = color.New(color.FgWhite).SprintFunc()
	GreenColor   = color.New(color.FgGreen).SprintFunc()
	PrimaryColor = color.New(color.FgCyan).SprintFunc()
	YellowColor  = color.New(color.FgYellow).SprintFunc()
	MagentaColor = color.New(color.FgMagenta).SprintFunc()
	RedColor     = color.New(color.FgRed).SprintFunc()
	MindedRed    = RedColor("-")
	LineRed      = RedColor("|")
)
