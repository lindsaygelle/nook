package cow

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	belleBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	belleCode = nook.Code{
		Value: ""}
)

var (
	belleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Belle"}

	belleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	belleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	belleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Belle"}

	belleGermanName = nook.Name{
		Language: language.German,
		Value:    "Heidi"}

	belleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Isabella"}

	belleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ミルク"}

	belleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	belleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	belleRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	belleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "洁洁"}

	belleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bella"}

	belleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	belleName = nook.Languages{
		language.AmericanEnglish:      belleAmericanEnglishName,
		language.CanadianFrench:       belleCanadianFrenchName,
		language.Dutch:                belleDutchName,
		language.French:               belleFrenchName,
		language.German:               belleGermanName,
		language.Italian:              belleItalianName,
		language.Japanese:             belleJapaneseName,
		language.Korean:               belleKoreanName,
		language.LatinAmericanSpanish: belleLatinAmericanSpanishName,
		language.Russian:              belleRussianName,
		language.SimplifiedChinese:    belleSimplifiedChineseName,
		language.Spanish:              belleSpanishName,
		language.TraditionalChinese:   belleTraditionalChineseName}
)

var (
	belleCharacter = nook.Character{
		Animal:   Cow,
		Birthday: belleBirthday,
		Code:     belleCode,
		Gender:   nook.Female,
		Name:     belleName}
)

var (
	belleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cuddles"}

	belleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "meuh-euh"}

	belleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	belleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "meuh-euh"}

	belleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knutsch"}

	belleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "amuuur"}

	belleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "にゅう"}

	belleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	belleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	belleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	belleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Unknown"}

	belleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cencerrito"}

	belleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	bellePhrase = nook.Languages{
		language.AmericanEnglish:      belleAmericanEnglishName,
		language.CanadianFrench:       belleCanadianFrenchName,
		language.Dutch:                belleDutchName,
		language.French:               belleFrenchName,
		language.German:               belleGermanName,
		language.Italian:              belleItalianName,
		language.Japanese:             belleJapaneseName,
		language.Korean:               belleKoreanName,
		language.LatinAmericanSpanish: belleLatinAmericanSpanishName,
		language.Russian:              belleRussianName,
		language.SimplifiedChinese:    belleSimplifiedChineseName,
		language.Spanish:              belleSpanishName,
		language.TraditionalChinese:   belleTraditionalChineseName}
)

var (
	Belle = nook.Villager{
		Character:   belleCharacter,
		Personality: nook.Peppy,
		Phrase:      bellePhrase}
)
