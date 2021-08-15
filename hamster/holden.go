package hamster

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Glutingluglu"}

	holdenDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	holdenFrenchName = nook.Name{
		Language: language.French,
		Value:    "Glutingluglu"}

	holdenGermanName = nook.Name{
		Language: language.German,
		Value:    "Holdenklebstift"}

	holdenItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Collottoincollo"}

	holdenJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "のりぼうノリノリ"}

	holdenKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "노리보둠칫둠칫"}

	holdenLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Holdentudiquesí"}

	holdenRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	holdenSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	holdenSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Holdentudiquesí"}

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
		Animal:   Hamster,
		Birthday: holdenBirthday,
		Code:     holdenCode,
		Gender:   nook.Male,
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
		Value:    "N/A"}

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
		Value:    "N/A"}

	holdenSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	holdenSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tudiquesí"}

	holdenTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Personality: nook.Jock,
		Phrase:      holdenPhrase}
)
