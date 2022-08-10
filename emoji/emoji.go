package emoji

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/getgort/emoji/v2/emoji/iamcal"
	"github.com/getgort/emoji/v2/emoji/joypixels"
)

// Emoji represents an emoji.
type Emoji struct {
	unicode string
}

// IamcalShortname returns a string shortname used to refer to the emoji in many
// chat providers, including slack.
func (e Emoji) IamcalShortname() string {
	s, _ := iamcal.UnicodeToShortname(e.NoSkinTones())
	return s
}

// JoypixelsShortname returns a string shortname used to refer to the emoji in
// chat providers that use JoyPixels (formerly EmojiOne), including discord.
func (e Emoji) JoypixelsShortname() string {
	s, _ := joypixels.UnicodeToShortname(e.FullEmoji())
	return s
}

// FullEmoji returns the actual unicode rune that is the emoji.
func (e Emoji) FullEmoji() string {
	return e.unicode
}

func (e Emoji) NoSkinTones() string {
	var sb strings.Builder
	for _, r := range e.unicode {
		if !unicode.Is(SkinToneModifiers, r) {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func (e Emoji) SkinTone() rune {
	for _, r := range e.unicode {
		if unicode.Is(SkinToneModifiers, r) {
			return r
		}
	}
	return rune(0)
}

func From(s string) (Emoji, error) {
	_, present := joypixels.UnicodeToShortname(s)
	if present {
		return FromUnicode(s), nil
	}
	u, present := iamcal.ShortnameToUnicode(s)
	if present {
		return FromUnicode(u), nil
	}
	u, present = joypixels.ShortnameToUnicode(s)
	if present {
		return FromUnicode(u), nil
	}

	return Emoji{}, fmt.Errorf("unknown emoji")
}

func FromUnicode(u string) Emoji {
	//var base []rune
	//var modifiers []rune
	//for _, r := range []rune(u) {
	//	if r == ZeroWidthJoiner {
	//		continue
	//	} else if unicode.Is(unicode.Sk, r) {
	//		modifiers = append(modifiers, r)
	//	} else {
	//		base = append(base, r)
	//	}
	//}
	//
	//return Emoji{
	//	base:      base,
	//	modifiers: modifiers,
	//}
	return Emoji{unicode: u}
}
