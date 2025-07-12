package ostrich

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	// floraBirthday represents flora birthday.
	floraBirthday = nook.Birthday{
		Day:   9,
		Month: time.February}
)

var (
	// floraCode represents flora code.
	floraCode = nook.Code{
		Value: "ost09"}
)

var (
	// floraAmericanEnglishName represents flora american english name.
	floraAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Flora"}

	// floraCanadianFrenchName represents flora canadian french name.
	floraCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Justine"}

	// floraDutchName represents flora dutch name.
	floraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Flora"}

	// floraFrenchName represents flora french name.
	floraFrenchName = nook.Name{
		Language: language.French,
		Value:    "Justine"}

	// floraGermanName represents flora german name.
	floraGermanName = nook.Name{
		Language: language.German,
		Value:    "Flora"}

	// floraItalianName represents flora italian name.
	floraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rosalba"}

	// floraJapaneseName represents flora japanese name.
	floraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フララ"}

	// floraKoreanName represents flora korean name.
	floraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "플라라"}

	// floraLatinAmericanSpanishName represents flora latin american spanish name.
	floraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Azucena"}

	// floraRussianName represents flora russian name.
	floraRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Флора"}

	// floraSimplifiedChineseName represents flora simplified chinese name.
	floraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "胡拉拉"}

	// floraSpanishName represents flora spanish name.
	floraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Azucena"}

	// floraTraditionalChineseName represents flora traditional chinese name.
	floraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "胡拉拉"}
)

var (
	// floraName represents flora name.
	floraName = nook.Languages{
		language.AmericanEnglish:      floraAmericanEnglishName,
		language.CanadianFrench:       floraCanadianFrenchName,
		language.Dutch:                floraDutchName,
		language.French:               floraFrenchName,
		language.German:               floraGermanName,
		language.Italian:              floraItalianName,
		language.Japanese:             floraJapaneseName,
		language.Korean:               floraKoreanName,
		language.LatinAmericanSpanish: floraLatinAmericanSpanishName,
		language.Russian:              floraRussianName,
		language.SimplifiedChinese:    floraSimplifiedChineseName,
		language.Spanish:              floraSpanishName,
		language.TraditionalChinese:   floraTraditionalChineseName}
)

var (
	// floraCharacter represents flora character.
	floraCharacter = nook.Character{
		Animal:   animal.Ostrich,
		Birthday: floraBirthday,
		Code:     floraCode,
		Key:      character.Flora,
		Gender:   gender.Female,
		Name:     floraName,
		Special:  false}
)

var (
	// floraAmericanEnglishPhrase represents flora american english phrase.
	floraAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pinky"}

	// floraCanadianFrenchPhrase represents flora canadian french phrase.
	floraCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "rosi rosa"}

	// floraDutchPhrase represents flora dutch phrase.
	floraDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "roze"}

	// floraFrenchPhrase represents flora french phrase.
	floraFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "guiboles"}

	// floraGermanPhrase represents flora german phrase.
	floraGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "flappflapp"}

	// floraItalianPhrase represents flora italian phrase.
	floraItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sivuplé"}

	// floraJapanesePhrase represents flora japanese phrase.
	floraJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だミン"}

	// floraKoreanPhrase represents flora korean phrase.
	floraKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "밍고밍고"}

	// floraLatinAmericanSpanishPhrase represents flora latin american spanish phrase.
	floraLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "croqui"}

	// floraRussianPhrase represents flora russian phrase.
	floraRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "розочка"}

	// floraSimplifiedChinesePhrase represents flora simplified chinese phrase.
	floraSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "红鹤"}

	// floraSpanishPhrase represents flora spanish phrase.
	floraSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "croqui"}

	// floraTraditionalChinesePhrase represents flora traditional chinese phrase.
	floraTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "紅鶴"}
)

var (
	// floraPhrase represents flora phrase.
	floraPhrase = nook.Languages{
		language.AmericanEnglish:      floraAmericanEnglishPhrase,
		language.CanadianFrench:       floraCanadianFrenchPhrase,
		language.Dutch:                floraDutchPhrase,
		language.French:               floraFrenchPhrase,
		language.German:               floraGermanPhrase,
		language.Italian:              floraItalianPhrase,
		language.Japanese:             floraJapanesePhrase,
		language.Korean:               floraKoreanPhrase,
		language.LatinAmericanSpanish: floraLatinAmericanSpanishPhrase,
		language.Russian:              floraRussianPhrase,
		language.SimplifiedChinese:    floraSimplifiedChinesePhrase,
		language.Spanish:              floraSpanishPhrase,
		language.TraditionalChinese:   floraTraditionalChinesePhrase}
)

var (
	// Flora represents flora.
	Flora = nook.Villager{
		Character:   floraCharacter,
		Personality: personality.Peppy,
		Phrase:      floraPhrase}
)
