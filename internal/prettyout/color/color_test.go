package color

import (
	"os"
	"testing"
)

/*
* @Author: zouyx
* @Email: zouyx@knowsec.com
* @Date:   2025/9/9 下午4:49
* @Package:
 */
func TestColor(t *testing.T) {
	Output = os.Stdout

	NoColor = false

	testColors := []struct {
		text string
		code Attribute
	}{
		{text: "black", code: FgBlack},
		{text: "red", code: FgRed},
		{text: "green", code: FgGreen},
		{text: "yellow", code: FgYellow},
		{text: "blue", code: FgBlue},
		{text: "magent", code: FgMagenta},
		{text: "cyan", code: FgCyan},
		{text: "white", code: FgWhite},
		{text: "hblack", code: FgHiBlack},
		{text: "hred", code: FgHiRed},
		{text: "hgreen", code: FgHiGreen},
		{text: "hyellow", code: FgHiYellow},
		{text: "hblue", code: FgHiBlue},
		{text: "hmagent", code: FgHiMagenta},
		{text: "hcyan", code: FgHiCyan},
		{text: "hwhite", code: FgHiWhite},
	}

	for _, c := range testColors {
		New(c.code).Fprintf(Output, "%s\n", c.text)
		//New(c.code).Print("hello")
		//line, _ := rb.ReadString('\n')
		//scannedLine := fmt.Sprintf("%q", line)
		//colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", c.code, c.text)
		//escapedForm := fmt.Sprintf("%q", colored)
		//
		//fmt.Printf("%s\t: %s\n", c.text, line)
		//
		//if scannedLine != escapedForm {
		//	t.Errorf("Expecting %s, got '%s'\n", escapedForm, scannedLine)
		//}
	}

	//for _, c := range testColors {
	//	line := New(c.code).Sprintf("%s", c.text)
	//	scannedLine := fmt.Sprintf("%q", line)
	//	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", c.code, c.text)
	//	escapedForm := fmt.Sprintf("%q", colored)
	//
	//	fmt.Printf("%s\t: %s\n", c.text, line)
	//
	//	if scannedLine != escapedForm {
	//		t.Errorf("Expecting %s, got '%s'\n", escapedForm, scannedLine)
	//	}
	//}
}
