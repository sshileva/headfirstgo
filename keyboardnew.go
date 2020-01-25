//Package keyboard reads user input from the keyboard.
package keyboardnew

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//GetFloat function reads a floating-point number from keyboard.
//It returns the number read and any error encountered.
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
