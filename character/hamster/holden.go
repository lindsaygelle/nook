package hamster

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
	holdenBirthday = nook.Birthday{
		Day:   11,
		Month: time.June}
)

var (
	holdenCode = nook.Code{
		Value: "ham08"}
)

var (
	holdenAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Holden"}

	holdenCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Glutin"}

	holdenDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	holdenFrenchName = nook.Name{
		Language: language.French,
		Value:    "Glutin"}

	holdenGermanName = nook.Name{
		Language: language.German,
		Value:    "Holden"}

	holdenItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Collotto"}

	holdenJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "のりぼう"}

	holdenKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "노리보"}

	holdenLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Holden"}

	holdenRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	holdenSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	holdenSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Holden"}

	holdenTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	holdenName = nook.Languages{
		language.AmericanEnglish:      holdenAmericanEnglishName,
		language.CanadianFrench:       holdenCanadianFrenchName,
		language.Dutch:                holdenDutchName,
		language.French:               holdenFrenchName,
		language.German:               holdenGermanName,
		language.Italian:              holdenItalianName,
		language.Japanese:             holdenJapaneseName,
		language.Korean:               holdenKoreanName,
		language.LatinAmericanSpanish: holdenLatinAmericanSpanishName,
		language.Russian:              holdenRussianName,
		language.SimplifiedChinese:    holdenSimplifiedChineseName,
		language.Spanish:              holdenSpanishName,
		language.TraditionalChinese:   holdenTraditionalChineseName}
)

var (
	holdenCharacter = nook.Character{
		Animal:   animal.Hamster,
		Birthday: holdenBirthday,
		Code:     holdenCode,
		Key:      character.Holden,
		Gender:   gender.Male,
		Name:     holdenName}
)

var (
	holdenAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "glue stick"}

	holdenCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "gluglu"}

	holdenDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	holdenFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "gluglu"}

	holdenGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "klebstift"}

	holdenItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "incollo"}

	holdenJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ノリノリ"}

	holdenKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "둠칫둠칫"}

	holdenLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tudiquesí"}

	holdenRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	holdenSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	holdenSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tudiquesí"}

	holdenTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	holdenPhrase = nook.Languages{
		language.AmericanEnglish:      holdenAmericanEnglishName,
		language.CanadianFrench:       holdenCanadianFrenchName,
		language.Dutch:                holdenDutchName,
		language.French:               holdenFrenchName,
		language.German:               holdenGermanName,
		language.Italian:              holdenItalianName,
		language.Japanese:             holdenJapaneseName,
		language.Korean:               holdenKoreanName,
		language.LatinAmericanSpanish: holdenLatinAmericanSpanishName,
		language.Russian:              holdenRussianName,
		language.SimplifiedChinese:    holdenSimplifiedChineseName,
		language.Spanish:              holdenSpanishName,
		language.TraditionalChinese:   holdenTraditionalChineseName}
)

var (
	Holden = nook.Villager{
		Character:   holdenCharacter,
		Personality: personality.Jock,
		Phrase:      holdenPhrase}
)
