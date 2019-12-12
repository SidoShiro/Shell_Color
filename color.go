package shell_color

import (
	"fmt"
	"strconv"
)

const (
	redShellColorNum    = 31
	greenShellColorNum  = 32
	cyanShellColorNum   = 36
	yellowShellColorNum = 33
	blueShellColorNum   = 34
)

func GreenPrintln(str string) {
	fmt.Println("\033[0;" + strconv.Itoa(greenShellColorNum) + "m" + str + "\033[0m")
}

func CyanPrintln(str string) {
	fmt.Println("\033[0;" + strconv.Itoa(cyanShellColorNum) + "m" + str + "\033[0m")
}

func RedPrintln(str string) {
	fmt.Println("\033[0;" + strconv.Itoa(redShellColorNum) + "m" + str + "\033[0m")
}

func YellowPrintln(str string) {
	fmt.Println("\033[0;" + strconv.Itoa(yellowShellColorNum) + "m" + str + "\033[0m")
}

func BluePrintln(str string) {
	fmt.Println("\033[0;" + strconv.Itoa(blueShellColorNum) + "m" + str + "\033[0m")
}
