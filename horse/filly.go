package horse

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Palominapapaye"}

	fillyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	fillyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Palominapapaye"}

	fillyGermanName = nook.Name{
		Language: language.German,
		Value:    "Beatehühott"}

	fillyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Filippagalop"}

	fillyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "7ごうはぁっ！"}

	fillyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "7호타앗"}

	fillyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Séptimagalope"}

	fillyRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	fillySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	fillySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Séptimagalope"}

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
		Animal:   Horse,
		Birthday: fillyBirthday,
		Code:     fillyCode,
		Gender:   nook.Female,
		Name:     fillyName}
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
		Value:    "N/A"}

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
		Value:    "N/A"}

	fillySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	fillySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "galope"}

	fillyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	fillyPhrase = nook.Languages{
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
	Filly = nook.Villager{
		Character:   fillyCharacter,
		Personality: nook.Normal,
		Phrase:      fillyPhrase}
)
