package output

import (
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func GetWindowWith() int {
	w, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0
	}
	return w
}
