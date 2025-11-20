package main

import (
	"github.com/fatih/color"
)

func main(){
	// Print with default helper functions
	color.Cyan("Prints text in cyan.")

	// A newline will be appended automatically
	color.Blue("Prints %s in blue.", "text")
	color.Red("We have red")
	color.Magenta("And many others ..")

}

