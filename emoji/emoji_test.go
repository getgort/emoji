package emoji

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmojiBasic(t *testing.T) {
	e1 := From("thumbsup")
	assert.Equal(t, "thumbsup", e1.Shortname())
	assert.Equal(t, "👍", e1.Unicode())

	e2 := From("+1")
	assert.Equal(t, "+1", e2.Shortname())
	assert.Equal(t, "👍", e2.Unicode())

	e3 := From("👍")
	assert.Equal(t, "👍", e3.Unicode())
	assert.Equal(t, "+1", e3.Shortname())
}
