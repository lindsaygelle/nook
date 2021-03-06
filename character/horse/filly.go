package horse

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
	fillyBirthday = nook.Birthday{
		Day:   11,
		Month: time.July}
)

var (
	fillyCode = nook.Code{
		Value: "hrs14"}
)

var (
	fillyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Filly"}

	fillyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Palomina"}

	fillyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	fillyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Palomina"}

	fillyGermanName = nook.Name{
		Language: language.German,
		Value:    "Beate"}

	fillyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Filippa"}

	fillyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "7ごう"}

	fillyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "7호"}

	fillyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Séptima"}

	fillyRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	fillySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	fillySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Séptima"}

	fillyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	fillyName = nook.Languages{
		language.AmericanEnglish:      fillyAmericanEnglishName,
		language.CanadianFrench:       fillyCanadianFrenchName,
		language.Dutch:                fillyDutchName,
		language.French:               fillyFrenchName,
		language.German:               fillyGermanName,
		language.Italian:              fillyItalianName,
		language.Japanese:             fillyJapaneseName,
		language.Korean:               fillyKoreanName,
		language.LatinAmericanSpanish: fillyLatinAmericanSpanishName,
		language.Russian:              fillyRussianName,
		language.SimplifiedChinese:    fillySimplifiedChineseName,
		language.Spanish:              fillySpanishName,
		language.TraditionalChinese:   fillyTraditionalChineseName}
)

var (
	fillyCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: fillyBirthday,
		Code:     fillyCode,
		Key:      character.Filly,
		Gender:   gender.Female,
		Name:     fillyName,
		Special:  false}
)

var (
	fillyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hah"}

	fillyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "papaye"}

	fillyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	fillyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "papaye"}

	fillyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hühott"}

	fillyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "galop"}

	fillyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "はぁっ！"}

	fillyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "타앗"}

	fillyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "galope"}

	fillyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	fillySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	fillySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "galope"}

	fillyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	fillyPhrase = nook.Languages{
		language.AmericanEnglish:      fillyAmericanEnglishPhrase,
		language.CanadianFrench:       fillyCanadianFrenchPhrase,
		language.Dutch:                fillyDutchPhrase,
		language.French:               fillyFrenchPhrase,
		language.German:               fillyGermanPhrase,
		language.Italian:              fillyItalianPhrase,
		language.Japanese:             fillyJapanesePhrase,
		language.Korean:               fillyKoreanPhrase,
		language.LatinAmericanSpanish: fillyLatinAmericanSpanishPhrase,
		language.Russian:              fillyRussianPhrase,
		language.SimplifiedChinese:    fillySimplifiedChinesePhrase,
		language.Spanish:              fillySpanishPhrase,
		language.TraditionalChinese:   fillyTraditionalChinesePhrase}
)

var (
	Filly = nook.Villager{
		Character:   fillyCharacter,
		Personality: personality.Normal,
		Phrase:      fillyPhrase}
)
