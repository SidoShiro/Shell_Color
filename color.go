package shell_color

import (
	"fmt"
	"strconv"
)

const (
	whiteShellColorNum  = 30
	redShellColorNum    = 31
	greenShellColorNum  = 32
	cyanShellColorNum   = 36
	yellowShellColorNum = 33
	blueShellColorNum   = 34
)

const start string = "\033[0;"
const end string = "\033[0m"

func Output(color int, str string) string {
	ss := "\033[0;" + strconv.Itoa(color) + "m" + str + "\033[0m"
	return ss
}

func GreenPrintln(str string) {
	fmt.Println(Output(greenShellColorNum, str))
}

func CyanPrintln(str string) {
	fmt.Println(Output(cyanShellColorNum, str))
}

func RedPrintln(str string) {
	fmt.Println(Output(redShellColorNum, str))
}

func YellowPrintln(str string) {
	fmt.Println(Output(yellowShellColorNum, str))
}

func BluePrintln(str string) {
	fmt.Println(Output(blueShellColorNum, str))
}

func GreenPrint(str string) {
	fmt.Print(Output(greenShellColorNum, str))
}

func CyanPrint(str string) {
	fmt.Print(Output(cyanShellColorNum, str))
}

func RedPrint(str string) {
	fmt.Print(Output(redShellColorNum, str))
}

func YellowPrint(str string) {
	fmt.Print(Output(yellowShellColorNum, str))
}

func BluePrint(str string) {
	fmt.Print(Output(blueShellColorNum, str))
}
