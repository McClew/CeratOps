package util

import (
	"fmt"
	"os"
)

func GetCSSUnix() string {
	file, err := os.Stat("./static/css/output.css")
	if err != nil {
		return "1"
	}
	return fmt.Sprintf("%d", file.ModTime().Unix())
}
