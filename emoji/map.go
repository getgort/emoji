package emoji

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

//go:embed emojidata.json
var emojiData string

var unicodeToShortname = make(map[string]string)
var shortnameToUnicode = make(map[string]string)

func init() {
	var iEmoji []internalEmoji
	err := json.NewDecoder(strings.NewReader(emojiData)).Decode(&iEmoji)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, e := range iEmoji {
		unicode, err := unifiedToUnicode(e.Unified)
		if err != nil {
			continue
		}
		unicodeToShortname[unicode] = e.ShortName
		shortnameToUnicode[e.ShortName] = unicode
		for _, n := range e.ShortNames {
			shortnameToUnicode[n] = unicode
		}
	}
}

type internalEmoji struct {
	ShortName  string   `json:"short_name"`
	ShortNames []string `json:"short_names"`
	Unified    string
}

func unifiedToUnicode(unified string) (string, error) {
	codes := strings.Split(unified, "-")
	var s strings.Builder
	for _, c := range codes {
		i, err := strconv.ParseInt(c, 16, 32)
		if err != nil {
			return "", err
		}
		s.WriteRune(rune(i))
	}
	return s.String(), nil
}
