package penguin

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
	// igglyBirthday represents iggly birthday.
	igglyBirthday = nook.Birthday{
		Day:   2,
		Month: time.November}
)

var (
	// igglyCode represents iggly code.
	igglyCode = nook.Code{
		Value: "pgn11"}
)

var (
	// igglyAmericanEnglishName represents iggly american english name.
	igglyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Iggly"}

	// igglyCanadianFrenchName represents iggly canadian french name.
	igglyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Urbain"}

	// igglyDutchName represents iggly dutch name.
	igglyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Iggly"}

	// igglyFrenchName represents iggly french name.
	igglyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Urbain"}

	// igglyGermanName represents iggly german name.
	igglyGermanName = nook.Name{
		Language: language.German,
		Value:    "Pingi"}

	// igglyItalianName represents iggly italian name.
	igglyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Piccio"}

	// igglyJapaneseName represents iggly japanese name.
	igglyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "のりまき"}

	// igglyKoreanName represents iggly korean name.
	igglyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "김말이"}

	// igglyLatinAmericanSpanishName represents iggly latin american spanish name.
	igglyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bobo"}

	// igglyRussianName represents iggly russian name.
	igglyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Игли"}

	// igglySimplifiedChineseName represents iggly simplified chinese name.
	igglySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "花寿司"}

	// igglySpanishName represents iggly spanish name.
	igglySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bobo"}

	// igglyTraditionalChineseName represents iggly traditional chinese name.
	igglyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "花壽司"}
)

var (
	// igglyName represents iggly name.
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
	// igglyCharacter represents iggly character.
	igglyCharacter = nook.Character{
		Animal:   animal.Penguin,
		Birthday: igglyBirthday,
		Code:     igglyCode,
		Key:      character.Iggly,
		Gender:   gender.Male,
		Name:     igglyName,
		Special:  false}
)

var (
	// igglyAmericanEnglishPhrase represents iggly american english phrase.
	igglyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "waddler"}

	// igglyCanadianFrenchPhrase represents iggly canadian french phrase.
	igglyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "zozio"}

	// igglyDutchPhrase represents iggly dutch phrase.
	igglyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "waggelaar"}

	// igglyFrenchPhrase represents iggly french phrase.
	igglyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "zozio"}

	// igglyGermanPhrase represents iggly german phrase.
	igglyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "plitsch"}

	// igglyItalianPhrase represents iggly italian phrase.
	igglyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "tubi tubi"}

	// igglyJapanesePhrase represents iggly japanese phrase.
	igglyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "クルクル"}

	// igglyKoreanPhrase represents iggly korean phrase.
	igglyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "펭글펭글"}

	// igglyLatinAmericanSpanishPhrase represents iggly latin american spanish phrase.
	igglyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chirpi"}

	// igglyRussianPhrase represents iggly russian phrase.
	igglyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ластошлеп"}

	// igglySimplifiedChinesePhrase represents iggly simplified chinese phrase.
	igglySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "卷卷"}

	// igglySpanishPhrase represents iggly spanish phrase.
	igglySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chirpi"}

	// igglyTraditionalChinesePhrase represents iggly traditional chinese phrase.
	igglyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "捲捲"}
)

var (
	// igglyPhrase represents iggly phrase.
	igglyPhrase = nook.Languages{
		language.AmericanEnglish:      igglyAmericanEnglishPhrase,
		language.CanadianFrench:       igglyCanadianFrenchPhrase,
		language.Dutch:                igglyDutchPhrase,
		language.French:               igglyFrenchPhrase,
		language.German:               igglyGermanPhrase,
		language.Italian:              igglyItalianPhrase,
		language.Japanese:             igglyJapanesePhrase,
		language.Korean:               igglyKoreanPhrase,
		language.LatinAmericanSpanish: igglyLatinAmericanSpanishPhrase,
		language.Russian:              igglyRussianPhrase,
		language.SimplifiedChinese:    igglySimplifiedChinesePhrase,
		language.Spanish:              igglySpanishPhrase,
		language.TraditionalChinese:   igglyTraditionalChinesePhrase}
)

var (
	// Iggly represents iggly.
	Iggly = nook.Villager{
		Character:   igglyCharacter,
		Personality: personality.Jock,
		Phrase:      igglyPhrase}
)
