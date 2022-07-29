package main

import (
	"fmt"
	"github.com/getgort/emoji/v2/emoji"
)

func main() {
	err := emoji.Gen("./emoji/emojidata.json")
	if err != nil {
		fmt.Println(err)
		return
	}
}
