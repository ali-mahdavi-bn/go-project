package logger

import (
	"fmt"
	"go-tel/src/backbone/helpers/colors"
	"time"
)

func Info(message any) {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	fmt.Printf("%s %s %s %s %s %s\n",
		colors.GreenColor(formattedTime),
		colors.LineRed,
		"INFO    ",
		colors.LineRed,
		colors.MindedRed,
		message,
	)
}

func Error(message any) {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	fmt.Printf("%s %s %s %s %s %s\n",
		colors.GreenColor(formattedTime),
		colors.LineRed,
		colors.RedColor("ERROR   "),
		colors.LineRed,
		colors.MindedRed,
		colors.RedColor(message),
	)
}

func Debug(message any) {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	fmt.Printf("%s %s %s %s %s %s\n",
		colors.GreenColor(formattedTime),
		colors.LineRed,
		colors.PrimaryColor("DEBUG   "),
		colors.LineRed,
		colors.MindedRed,
		colors.PrimaryColor(message),
	)
}
