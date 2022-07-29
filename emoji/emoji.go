package emoji

// Emoji represents an emoji.
type Emoji struct {
	shortname string
	unicode   string
}

// Shortname returns a string shortname used to refer to the emoji.
func (e Emoji) Shortname() string {
	return e.shortname
}

// Unicode returns the actual unicode rune that is the emoji.
func (e Emoji) Unicode() string {
	return e.unicode
}

func From(s string) Emoji {
	short, present := unicodeToShortname[s]
	if present {
		return Emoji{
			shortname: short,
			unicode:   s,
		}
	}
	unicode, present := shortnameToUnicode[s]
	if present {
		return Emoji{
			shortname: s,
			unicode:   unicode,
		}
	}

	return Emoji{}
}
