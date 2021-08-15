package chicken

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	bettyBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	bettyCode = nook.Code{
		Value: ""}
)

var (
	bettyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Betty"}

	bettyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	bettyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	bettyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Betty"}

	bettyGermanName = nook.Name{
		Language: language.German,
		Value:    "Margot"}

	bettyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Crestina"}

	bettyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ホイップ"}

	bettyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	bettyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	bettyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	bettySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嗒嗒"}

	bettySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Clueca"}

	bettyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	bettyName = nook.Languages{
		language.AmericanEnglish:      bettyAmericanEnglishName,
		language.CanadianFrench:       bettyCanadianFrenchName,
		language.Dutch:                bettyDutchName,
		language.French:               bettyFrenchName,
		language.German:               bettyGermanName,
		language.Italian:              bettyItalianName,
		language.Japanese:             bettyJapaneseName,
		language.Korean:               bettyKoreanName,
		language.LatinAmericanSpanish: bettyLatinAmericanSpanishName,
		language.Russian:              bettyRussianName,
		language.SimplifiedChinese:    bettySimplifiedChineseName,
		language.Spanish:              bettySpanishName,
		language.TraditionalChinese:   bettyTraditionalChineseName}
)

var (
	bettyCharacter = nook.Character{
		Animal:   Chicken,
		Birthday: bettyBirthday,
		Code:     bettyCode,
		Gender:   nook.Female,
		Name:     bettyName}
)

var (
	bettyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	bettyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	bettyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	bettyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "côt côt"}

	bettyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "klackl"}

	bettyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "coccodè"}

	bettyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だがね"}

	bettyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	bettyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	bettyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	bettySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "ok"}

	bettySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "clo-clo"}

	bettyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	bettyPhrase = nook.Languages{
		language.AmericanEnglish:      bettyAmericanEnglishName,
		language.CanadianFrench:       bettyCanadianFrenchName,
		language.Dutch:                bettyDutchName,
		language.French:               bettyFrenchName,
		language.German:               bettyGermanName,
		language.Italian:              bettyItalianName,
		language.Japanese:             bettyJapaneseName,
		language.Korean:               bettyKoreanName,
		language.LatinAmericanSpanish: bettyLatinAmericanSpanishName,
		language.Russian:              bettyRussianName,
		language.SimplifiedChinese:    bettySimplifiedChineseName,
		language.Spanish:              bettySpanishName,
		language.TraditionalChinese:   bettyTraditionalChineseName}
)

var (
	Betty = nook.Villager{
		Character:   bettyCharacter,
		Personality: nook.Normal,
		Phrase:      bettyPhrase}
)
