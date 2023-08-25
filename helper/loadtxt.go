package helper

import (
	"log"
	"os"
)

func ReadText(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(b)
	}
	txtAll := string(b)
	return txtAll
}
