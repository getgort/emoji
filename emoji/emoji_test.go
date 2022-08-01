package emoji

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmojiBasic(t *testing.T) {
	e1, err := From("thumbsup")
	assert.NoError(t, err)
	assert.Equal(t, "+1", e1.IamcalShortname())
	assert.Equal(t, "👍", e1.FullyQualifiedUnicode())

	e2, err := From("+1")
	assert.NoError(t, err)
	assert.Equal(t, "+1", e2.IamcalShortname())
	assert.Equal(t, "👍", e2.FullyQualifiedUnicode())

	e3, err := From("👍")
	assert.NoError(t, err)
	assert.Equal(t, "👍", e3.FullyQualifiedUnicode())
	assert.Equal(t, "+1", e3.IamcalShortname())
}

func TestEmojiMulti(t *testing.T) {
	england, err := From("flag-england")
	assert.NoError(t, err)
	assert.Equal(t, "🏴󠁧󠁢󠁥󠁮󠁧󠁿", england.FullyQualifiedUnicode())

	womanWomanBoyBoy, err := From("woman-woman-boy-boy")
	assert.NoError(t, err)
	assert.Equal(t, "👩‍👩‍👦‍👦", womanWomanBoyBoy.FullyQualifiedUnicode())
}

func TestPlatformSpecificShortnames(t *testing.T) {
	wwbb, err := From("woman-woman-boy-boy")
	assert.NoError(t, err)
	assert.Equal(t, "woman-woman-boy-boy", wwbb.IamcalShortname())
	assert.Equal(t, "family_wwbb", wwbb.JoypixelsShortname())
}
