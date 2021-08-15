package penguin

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	igglyBirthday = nook.Birthday{
		Day:   2,
		Month: time.November}
)

var (
	igglyCode = nook.Code{
		Value: "pgn11"}
)

var (
	igglyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Iggly"}

	igglyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Urbain"}

	igglyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Iggly"}

	igglyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Urbain"}

	igglyGermanName = nook.Name{
		Language: language.German,
		Value:    "Pingi"}

	igglyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Piccio"}

	igglyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "のりまき"}

	igglyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "김말이"}

	igglyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bobo"}

	igglyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Игли"}

	igglySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "花寿司"}

	igglySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bobo"}

	igglyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "花壽司"}
)

var (
	igglyName = nook.Languages{
		language.AmericanEnglish:      igglyAmericanEnglishName,
		language.CanadianFrench:       igglyCanadianFrenchName,
		language.Dutch:                igglyDutchName,
		language.French:               igglyFrenchName,
		language.German:               igglyGermanName,
		language.Italian:              igglyItalianName,
		language.Japanese:             igglyJapaneseName,
		language.Korean:               igglyKoreanName,
		language.LatinAmericanSpanish: igglyLatinAmericanSpanishName,
		language.Russian:              igglyRussianName,
		language.SimplifiedChinese:    igglySimplifiedChineseName,
		language.Spanish:              igglySpanishName,
		language.TraditionalChinese:   igglyTraditionalChineseName}
)

var (
	igglyCharacter = nook.Character{
		Animal:   Penguin,
		Birthday: igglyBirthday,
		Code:     igglyCode,
		Gender:   nook.Male,
		Name:     igglyName}
)

var (
	igglyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "waddler"}

	igglyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "zozio"}

	igglyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "waggelaar"}

	igglyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "zozio"}

	igglyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "plitsch"}

	igglyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "tubi tubi"}

	igglyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "クルクル"}

	igglyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "펭글펭글"}

	igglyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chirpi"}

	igglyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ластошлеп"}

	igglySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "卷卷"}

	igglySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chirpi"}

	igglyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "捲捲"}
)

var (
	igglyPhrase = nook.Languages{
		language.AmericanEnglish:      igglyAmericanEnglishName,
		language.CanadianFrench:       igglyCanadianFrenchName,
		language.Dutch:                igglyDutchName,
		language.French:               igglyFrenchName,
		language.German:               igglyGermanName,
		language.Italian:              igglyItalianName,
		language.Japanese:             igglyJapaneseName,
		language.Korean:               igglyKoreanName,
		language.LatinAmericanSpanish: igglyLatinAmericanSpanishName,
		language.Russian:              igglyRussianName,
		language.SimplifiedChinese:    igglySimplifiedChineseName,
		language.Spanish:              igglySpanishName,
		language.TraditionalChinese:   igglyTraditionalChineseName}
)

var (
	Iggly = nook.Villager{
		Character:   igglyCharacter,
		Personality: nook.Jock,
		Phrase:      igglyPhrase}
)
