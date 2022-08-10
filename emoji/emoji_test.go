package emoji

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmojiBasic(t *testing.T) {
	e1, err := From("thumbsup")
	assert.NoError(t, err)
	assert.Equal(t, "+1", e1.IamcalShortname())
	assert.Equal(t, "👍", e1.FullEmoji())

	e2, err := From("+1")
	assert.NoError(t, err)
	assert.Equal(t, "+1", e2.IamcalShortname())
	assert.Equal(t, "👍", e2.FullEmoji())

	e3, err := From("👍")
	assert.NoError(t, err)
	assert.Equal(t, "👍", e3.FullEmoji())
	assert.Equal(t, "+1", e3.IamcalShortname())
}

func TestEmojiMulti(t *testing.T) {
	england, err := From("flag-england")
	assert.NoError(t, err)
	assert.Equal(t, "🏴󠁧󠁢󠁥󠁮󠁧󠁿", england.FullEmoji())

	womanWomanBoyBoy, err := From("woman-woman-boy-boy")
	assert.NoError(t, err)
	assert.Equal(t, "👩‍👩‍👦‍👦", womanWomanBoyBoy.FullEmoji())
}

func TestPlatformSpecificShortnames(t *testing.T) {
	wwbb, err := From("woman-woman-boy-boy")
	assert.NoError(t, err)
	assert.Equal(t, "woman-woman-boy-boy", wwbb.IamcalShortname())
	assert.Equal(t, "family_wwbb", wwbb.JoypixelsShortname())
}

func TestEmojiSkinTone(t *testing.T) {

}
