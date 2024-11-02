package main

import (
	"fmt"

	"github.com/b-s-a-1985/colored"
)

func main() {
	// test colored package

	fmt.Println(colored.Blue + "Blue txt")
	fmt.Println(colored.Cyan + "Cyan txt")
	fmt.Println(colored.Gray + "Gray txt")
	fmt.Println(colored.Green + "Green txt")
	fmt.Println(colored.Magenta + "Magenta txt")
	fmt.Println(colored.White + "White txt")
	fmt.Println(colored.Red + "Red txt")
	fmt.Println(colored.Yellow + "Yellow txt")
	fmt.Println()

	fmt.Println(colored.BoldBlue + "Bold blue text")
	fmt.Println(colored.BoldCyan + "Bold cyan text")
	fmt.Println(colored.BoldGray + "Bold gray text")
	fmt.Println(colored.Green + "Bold green text")
	fmt.Println(colored.BoldMagenta + "Bold magenta text")
	fmt.Println(colored.BoldWhite + "Bold white text")
	fmt.Println(colored.BoldYellow + "Bold yellow text" + colored.Reset)
	fmt.Println()
}
