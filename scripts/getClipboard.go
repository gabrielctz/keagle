package scripts

import (
	"github.com/atotto/clipboard"
)

func GetClipboard() string {

	clipboard, _ := clipboard.ReadAll()
	return clipboard

}
