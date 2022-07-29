package emoji

import (
	"encoding/json"
	"net/http"
	"os"
)

const dataUrl = "https://raw.githubusercontent.com/iamcal/emoji-data/master/emoji_pretty.json"

func Gen(out string) error {
	res, err := http.Get(dataUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var data []internalEmoji
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return err
	}
	file, err := os.Create(out)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(&data)
	if err != nil {
		return err
	}

	return nil
}
