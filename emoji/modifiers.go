package emoji

import (
	"unicode"

	"golang.org/x/text/unicode/rangetable"
)

const (
	ZeroWidthJoiner = '\u200d'

	VariantSelector16 = '\uFE0F'

	SkinToneLight       = '\U0001F3FB'
	SkinToneMediumLight = '\U0001F3Fc'
	SkinToneMedium      = '\U0001F3FD'
	SkinToneMediumDark  = '\U0001F3FE'
	SkinToneDark        = '\U0001F3FF'

	HairRed   = '\U0001F9B0'
	HairCurly = '\U0001F9B1'
	HairBald  = '\U0001F9B2'
	HairWhite = '\U0001F9B3'

	RegionalIndicatorA = '\U0001F1E6'
	RegionalIndicatorB = '\U0001F1E7'
	RegionalIndicatorC = '\U0001F1E8'
	RegionalIndicatorD = '\U0001F1E9'
	RegionalIndicatorE = '\U0001F1EA'
	RegionalIndicatorF = '\U0001F1EB'
	RegionalIndicatorG = '\U0001F1EC'
	RegionalIndicatorH = '\U0001F1ED'
	RegionalIndicatorI = '\U0001F1EE'
	RegionalIndicatorJ = '\U0001F1EF'
	RegionalIndicatorK = '\U0001F1F0'
	RegionalIndicatorL = '\U0001F1F1'
	RegionalIndicatorM = '\U0001F1F2'
	RegionalIndicatorN = '\U0001F1F3'
	RegionalIndicatorO = '\U0001F1F4'
	RegionalIndicatorP = '\U0001F1F5'
	RegionalIndicatorQ = '\U0001F1F6'
	RegionalIndicatorR = '\U0001F1F7'
	RegionalIndicatorS = '\U0001F1F8'
	RegionalIndicatorT = '\U0001F1F9'
	RegionalIndicatorU = '\U0001F1FA'
	RegionalIndicatorV = '\U0001F1FB'
	RegionalIndicatorW = '\U0001F1FC'
	RegionalIndicatorX = '\U0001F1FD'
	RegionalIndicatorY = '\U0001F1FE'
	RegionalIndicatorZ = '\U0001F1FF'

	TagLatinSmallLetterA = '\U000E0061'
	TagLatinSmallLetterB = '\U000E0062'
	TagLatinSmallLetterC = '\U000E0063'
	TagLatinSmallLetterD = '\U000E0064'
	TagLatinSmallLetterE = '\U000E0065'
	TagLatinSmallLetterF = '\U000E0066'
	TagLatinSmallLetterG = '\U000E0067'
	TagLatinSmallLetterH = '\U000E0068'
	TagLatinSmallLetterI = '\U000E0069'
	TagLatinSmallLetterJ = '\U000E006A'
	TagLatinSmallLetterK = '\U000E006B'
	TagLatinSmallLetterL = '\U000E006C'
	TagLatinSmallLetterM = '\U000E006D'
	TagLatinSmallLetterN = '\U000E006E'
	TagLatinSmallLetterO = '\U000E006F'
	TagLatinSmallLetterP = '\U000E0070'
	TagLatinSmallLetterQ = '\U000E0071'
	TagLatinSmallLetterR = '\U000E0072'
	TagLatinSmallLetterS = '\U000E0073'
	TagLatinSmallLetterT = '\U000E0074'
	TagLatinSmallLetterU = '\U000E0075'
	TagLatinSmallLetterV = '\U000E0076'
	TagLatinSmallLetterW = '\U000E0077'
	TagLatinSmallLetterX = '\U000E0078'
	TagLatinSmallLetterY = '\U000E0079'
	TagLatinSmallLetterZ = '\U000E007A'

	TagDigitZero  = '\U000E0030'
	TagDigitOne   = '\U000E0031'
	TagDigitTwo   = '\U000E0032'
	TagDigitThree = '\U000E0033'
	TagDigitFour  = '\U000E0034'
	TagDigitFive  = '\U000E0035'
	TagDigitSix   = '\U000E0036'
	TagDigitSeven = '\U000E0037'
	TagDigitEight = '\U000E0038'
	TagDigitNine  = '\U000E0039'

	TagCancel = '\U000E007F'
)

var (
	RegionalIndicators = unicode.Regional_Indicator
	Tags               = &unicode.RangeTable{
		R32: []unicode.Range32{
			{
				Lo:     0xE0020,
				Hi:     0xE007F,
				Stride: 1,
			},
		},
		LatinOffset: 0,
	}
	HairModifiers = &unicode.RangeTable{
		R32: []unicode.Range32{
			{
				Lo:     uint32(HairRed),
				Hi:     uint32(HairWhite),
				Stride: 1,
			},
		},
	}
	SkinToneModifiers = &unicode.RangeTable{
		R32: []unicode.Range32{
			{
				Lo:     uint32(SkinToneLight),
				Hi:     uint32(SkinToneDark),
				Stride: 1,
			},
		},
	}
	Modifiers = rangetable.Merge(HairModifiers, SkinToneModifiers)
)
