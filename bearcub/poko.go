package bearcub

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	pokoBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	pokoCode = nook.Code{
		Value: ""}
)

var (
	pokoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Poko"}

	pokoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	pokoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	pokoFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	pokoGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	pokoItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	pokoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ポコへへ"}

	pokoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	pokoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	pokoRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	pokoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	pokoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	pokoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	pokoName = nook.Languages{
		language.AmericanEnglish:      pokoAmericanEnglishName,
		language.CanadianFrench:       pokoCanadianFrenchName,
		language.Dutch:                pokoDutchName,
		language.French:               pokoFrenchName,
		language.German:               pokoGermanName,
		language.Italian:              pokoItalianName,
		language.Japanese:             pokoJapaneseName,
		language.Korean:               pokoKoreanName,
		language.LatinAmericanSpanish: pokoLatinAmericanSpanishName,
		language.Russian:              pokoRussianName,
		language.SimplifiedChinese:    pokoSimplifiedChineseName,
		language.Spanish:              pokoSpanishName,
		language.TraditionalChinese:   pokoTraditionalChineseName}
)

var (
	pokoCharacter = nook.Character{
		Animal:   Bearcub,
		Birthday: pokoBirthday,
		Code:     pokoCode,
		Gender:   nook.Male,
		Name:     pokoName}
)

var (
	pokoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	pokoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	pokoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	pokoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	pokoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	pokoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	pokoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "へへ"}

	pokoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	pokoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	pokoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	pokoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	pokoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	pokoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	pokoPhrase = nook.Languages{
		language.AmericanEnglish:      pokoAmericanEnglishName,
		language.CanadianFrench:       pokoCanadianFrenchName,
		language.Dutch:                pokoDutchName,
		language.French:               pokoFrenchName,
		language.German:               pokoGermanName,
		language.Italian:              pokoItalianName,
		language.Japanese:             pokoJapaneseName,
		language.Korean:               pokoKoreanName,
		language.LatinAmericanSpanish: pokoLatinAmericanSpanishName,
		language.Russian:              pokoRussianName,
		language.SimplifiedChinese:    pokoSimplifiedChineseName,
		language.Spanish:              pokoSpanishName,
		language.TraditionalChinese:   pokoTraditionalChineseName}
)

var (
	Poko = nook.Villager{
		Character:   pokoCharacter,
		Personality: nook.Jock,
		Phrase:      pokoPhrase}
)
