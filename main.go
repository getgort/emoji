package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/getgort/emoji/v2/emoji"
	"go/format"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
)

func main() {
	err := GenIamcal()
	if err != nil {
		fmt.Println("Failed to generate iamcal emoji map:", err)
		return
	}

	err = GenJoypixels()
	if err != nil {
		fmt.Println("Failed to generate joypixels emoji map:", err)
		return
	}
}

const (
	iamcalDataUrl    = "https://raw.githubusercontent.com/iamcal/emoji-data/master/emoji_pretty.json"
	joypixelsDataUrl = "https://raw.githubusercontent.com/joypixels/emoji-toolkit/master/emoji.json"
)

func GenIamcal() error {
	res, err := http.Get(iamcalDataUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var data []iamcalEmoji
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return err
	}

	var idata []internalEmoji
	for _, e := range data {
		u, _ := unifiedToUnicode(e.Unified) // TODO
		idata = append(idata, internalEmoji{
			Shortname:  e.Shortname,
			Shortnames: e.Shortnames,
			Unicode:    u,
		})
	}

	err = Gen(idata, "iamcal", "./emoji/iamcal/map.go")
	if err != nil {
		return err
	}

	return nil
}

func GenJoypixels() error {
	res, err := http.Get(joypixelsDataUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var data map[string]joypixelsEmoji
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return err
	}

	var idata []internalEmoji
	for _, e := range data {
		u, _ := unifiedToUnicode(e.Codepoints.Unified) // TODO
		var shortnamesAlt []string
		for _, s := range shortnamesAlt {
			shortnamesAlt = append(shortnamesAlt, strings.Trim(s, ":"))
		}
		idata = append(idata, internalEmoji{
			Shortname:  strings.Trim(e.Shortname, ":"),
			Shortnames: append(shortnamesAlt, e.Shortname),
			Unicode:    u,
		})
	}

	err = Gen(idata, "joypixels", "./emoji/joypixels/map.go")
	if err != nil {
		return err
	}

	return nil
}

func Gen(data []internalEmoji, packageName, out string) error {
	file, err := os.Create(out)
	if err != nil {
		return err
	}
	defer file.Close()

	t, err := template.New("map.gotemplate").Funcs(template.FuncMap{
		"quote": strconv.Quote,
	}).ParseFiles("map.gotemplate")
	if err != nil {
		return err
	}

	for i := range data {
		data[i].Unicode = emoji.FromUnicode(data[i].Unicode).FullyQualifiedUnicode()
	}

	b := bytes.NewBuffer([]byte{})

	err = t.Execute(b, templateParams{
		Emoji:       data,
		PackageName: packageName,
	})
	if err != nil {
		return err
	}

	formatted, err := format.Source(b.Bytes())
	if err != nil {
		return err
	}

	_, err = file.Write(formatted)
	if err != nil {
		return err
	}

	return nil
}

type templateParams struct {
	Emoji       []internalEmoji
	PackageName string
}

type internalEmoji struct {
	Shortname  string
	Shortnames []string
	Unicode    string
}

type iamcalEmoji struct {
	Shortname  string   `json:"short_name"`
	Shortnames []string `json:"short_names"`
	Unified    string
}

type joypixelsEmoji struct {
	// Shortname and ShortnamesAlt have colons surrounding them
	Shortname     string   `json:"shortname"`
	ShortnamesAlt []string `json:"shortname_alternates"`
	Codepoints    struct {
		Unified string `json:"fully_qualified"`
	} `json:"code_points"`
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
