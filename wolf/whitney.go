package wolf

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	whitneyBirthday = nook.Birthday{
		Day:   17,
		Month: time.September}
)

var (
	whitneyCode = nook.Code{
		Value: "wol03"}
)

var (
	whitneyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Whitney"}

	whitneyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Blanchebing bang"}

	whitneyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Whitneyhap"}

	whitneyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Blanchesbam"}

	whitneyGermanName = nook.Name{
		Language: language.German,
		Value:    "Lupajaulll"}

	whitneyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Biancasnappy"}

	whitneyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ビアンカステキね"}

	whitneyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "비앙카멋져"}

	whitneyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lupeauf-auf"}

	whitneyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Уитницап"}

	whitneySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "毕安卡漂亮"}

	whitneySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lupeauf-auf"}

	whitneyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "畢安卡漂亮"}
)

var (
	whitneyName = nook.Languages{
		language.AmericanEnglish:      whitneyAmericanEnglishName,
		language.CanadianFrench:       whitneyCanadianFrenchName,
		language.Dutch:                whitneyDutchName,
		language.French:               whitneyFrenchName,
		language.German:               whitneyGermanName,
		language.Italian:              whitneyItalianName,
		language.Japanese:             whitneyJapaneseName,
		language.Korean:               whitneyKoreanName,
		language.LatinAmericanSpanish: whitneyLatinAmericanSpanishName,
		language.Russian:              whitneyRussianName,
		language.SimplifiedChinese:    whitneySimplifiedChineseName,
		language.Spanish:              whitneySpanishName,
		language.TraditionalChinese:   whitneyTraditionalChineseName}
)

var (
	whitneyCharacter = nook.Character{
		Animal:   Wolf,
		Birthday: whitneyBirthday,
		Code:     whitneyCode,
		Gender:   nook.Female,
		Name:     whitneyName}
)

var (
	whitneyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snappy"}

	whitneyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bing bang"}

	whitneyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hap"}

	whitneyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "sbam"}

	whitneyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "jaulll"}

	whitneyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "snappy"}

	whitneyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ステキね"}

	whitneyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "멋져"}

	whitneyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "auf-auf"}

	whitneyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "цап"}

	whitneySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "漂亮"}

	whitneySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "auf-auf"}

	whitneyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "漂亮"}
)

var (
	whitneyPhrase = nook.Languages{
		language.AmericanEnglish:      whitneyAmericanEnglishName,
		language.CanadianFrench:       whitneyCanadianFrenchName,
		language.Dutch:                whitneyDutchName,
		language.French:               whitneyFrenchName,
		language.German:               whitneyGermanName,
		language.Italian:              whitneyItalianName,
		language.Japanese:             whitneyJapaneseName,
		language.Korean:               whitneyKoreanName,
		language.LatinAmericanSpanish: whitneyLatinAmericanSpanishName,
		language.Russian:              whitneyRussianName,
		language.SimplifiedChinese:    whitneySimplifiedChineseName,
		language.Spanish:              whitneySpanishName,
		language.TraditionalChinese:   whitneyTraditionalChineseName}
)

var (
	Whitney = nook.Villager{
		Character:   whitneyCharacter,
		Personality: nook.Snooty,
		Phrase:      whitneyPhrase}
)
