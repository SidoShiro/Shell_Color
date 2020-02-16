package shell_color

import (
	"fmt"
)

const (
	// whiteShellColorNum  = 30
	redShellColorNum    = "31"
	greenShellColorNum  = "32"
	yellowShellColorNum = "33"
	blueShellColorNum   = "34"
	cyanShellColorNum   = "36"
	start               = "\033[0;"
	end                 = "\033[0m"
)

// output uses fmt.Sprint and a set of colors to output a colored string
func output(color string, v ...interface{}) string {
	ss := start + color + "m" + fmt.Sprint(v...) + end
	return ss
}

// GreenSprint formats using the default formats for its operands and returns
// the resulting string. Spaces are added between operand when neither is a
// string. The string is now Green colored.
func GreenSprint(v ...interface{}) string { return output(greenShellColorNum, v...) }

// CyanSprint formats using the default formats for its operands and returns
// the resulting string. Spaces are added between operand when neither is a
// string. The string is now Cyan colored.
func CyanSprint(v ...interface{}) string { return output(cyanShellColorNum, v...) }

// RedSprint formats using the default formats for its operands and returns
// the resulting string. Spaces are added between operand when neither is a
// string. The string is now Red colored.
func RedSprint(v ...interface{}) string { return output(redShellColorNum, v...) }

// YellowSprint formats using the default formats for its operands and returns
// the resulting string. Spaces are added between operand when neither is a
// string. The string is now Yellow colored.
func YellowSprint(v ...interface{}) string { return output(yellowShellColorNum, v...) }

// BlueSprint formats using the default formats for its operands and returns
// the resulting string. Spaces are added between operand when neither is a
// string. The string is now Blue colored.
func BlueSprint(v ...interface{}) string { return output(blueShellColorNum, v...) }

// GreenPrintln formats using the default formats for its operands and writes
// to standard output. Spaces are always added bewteen operands and a newline is
// appended. The string is now Green colored. It returns the number of bytes
// written and any write error encountered.
func GreenPrintln(v ...interface{}) (int, error) { return fmt.Println(GreenSprint(v...)) }

// CyanPrintln formats using the default formats for its operands and writes
// to standard output. Spaces are always added bewteen operands and a newline is
// appended. The string is now Cyan colored. It returns the number of bytes
// written and any write error encountered.
func CyanPrintln(v ...interface{}) (int, error) { return fmt.Println(CyanSprint(v...)) }

// RedPrintln formats using the default formats for its operands and writes
// to standard output. Spaces are always added bewteen operands and a newline is
// appended. The string is now Red colored. It returns the number of bytes
// written and any write error encountered.
func RedPrintln(v ...interface{}) (int, error) { return fmt.Println(RedSprint(v...)) }

// YellowPrintln formats using the default formats for its operands and writes
// to standard output. Spaces are always added bewteen operands and a newline is
// appended. The string is now Yellow colored. It returns the number of bytes
// written and any write error encountered.
func YellowPrintln(v ...interface{}) (int, error) { return fmt.Println(YellowSprint(v...)) }

// BluePrintln formats using the default formats for its operands and writes
// to standard output. Spaces are always added bewteen operands and a newline is
// appended. The string is now Blue colored. It returns the number of bytes
// written and any write error encountered.
func BluePrintln(v ...interface{}) (int, error) { return fmt.Println(BlueSprint(v...)) }

// func Print(v ...interface{}) { fmt.Print(Sprint(v...)) }

// GreenPrint formats using the default formats for its operands and writes to
// standard output. Spaces are added between operands when neither is a string.
// The string is now Green colored. It returns the number of bytes written and
// anny write error encountered.
func GreenPrint(v ...interface{}) (int, error) { return fmt.Print(GreenSprint(v...)) }

// CyanPrint formats using the default formats for its operands and writes to
// standard output. Spaces are added between operands when neither is a string.
// The string is now Cyan colored. It returns the number of bytes written and
// anny write error encountered.
func CyanPrint(v ...interface{}) (int, error) { return fmt.Print(CyanSprint(v...)) }

// RedPrint formats using the default formats for its operands and writes to
// standard output. Spaces are added between operands when neither is a string.
// The string is now Red colored. It returns the number of bytes written and
// anny write error encountered.
func RedPrint(v ...interface{}) (int, error) { return fmt.Print(RedSprint(v...)) }

// YellowPrint formats using the default formats for its operands and writes to
// standard output. Spaces are added between operands when neither is a string.
// The string is now Yellow colored. It returns the number of bytes written and
// anny write error encountered.
func YellowPrint(v ...interface{}) (int, error) { return fmt.Print(YellowSprint(v...)) }

// BluePrint formats using the default formats for its operands and writes to
// standard output. Spaces are added between operands when neither is a string.
// The string is now Blue colored. It returns the number of bytes written and
// anny write error encountered.
func BluePrint(v ...interface{}) (int, error) { return fmt.Print(BlueSprint(v...)) }
