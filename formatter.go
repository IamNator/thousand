package thousand

import (
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)



func ThousandFormatter(num float64) string {
	p := message.NewPrinter(language.English)
	return RemoveTrailingZeros(p.Sprintf("%f", num))
}

func RemoveTrailingZeros(num string) string {
	var result string
	if !strings.Contains(num, ".") { // if num is not a float
		return num + ".00"
	}

	parts := strings.SplitN(num, ".", 2)
	if parts == nil || len(parts) != 2 {
		return num
	}
	wholeNum := parts[0]
	decimalNum := parts[1]

	//run through it
	for i := len(decimalNum) - 1; i >= 0; i-- {
		if decimalNum[i] == '0' {
			decimalNum = decimalNum[:len(decimalNum)-1]
		} else {
			break
		}
	}

	for len(decimalNum) < 2 {
		decimalNum = decimalNum + "0"
	}

	result = wholeNum + "." + decimalNum

	return result
}
