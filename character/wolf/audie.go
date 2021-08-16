package wolf

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	audieBirthday = nook.Birthday{
		Day:   31,
		Month: time.August}
)

var (
	audieCode = nook.Code{
		Value: "wol12"}
)

var (
	audieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Audie"}

	audieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Monica"}

	audieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Audie"}

	audieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Monica"}

	audieGermanName = nook.Name{
		Language: language.German,
		Value:    "Katharina"}

	audieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lupilia"}

	audieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "モニカ"}

	audieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "모니카"}

	audieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mónica"}

	audieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Оди"}

	audieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "莫妮卡"}

	audieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mónica"}

	audieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "莫妮卡"}
)

var (
	audieName = nook.Languages{
		language.AmericanEnglish:      audieAmericanEnglishName,
		language.CanadianFrench:       audieCanadianFrenchName,
		language.Dutch:                audieDutchName,
		language.French:               audieFrenchName,
		language.German:               audieGermanName,
		language.Italian:              audieItalianName,
		language.Japanese:             audieJapaneseName,
		language.Korean:               audieKoreanName,
		language.LatinAmericanSpanish: audieLatinAmericanSpanishName,
		language.Russian:              audieRussianName,
		language.SimplifiedChinese:    audieSimplifiedChineseName,
		language.Spanish:              audieSpanishName,
		language.TraditionalChinese:   audieTraditionalChineseName}
)

var (
	audieCharacter = nook.Character{
		Animal:   animal.Wolf,
		Birthday: audieBirthday,
		Code:     audieCode,
		Gender:   gender.Female,
		Name:     audieName}
)

var (
	audieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "foxtrot"}

	audieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "trottine"}

	audieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "foxtrot"}

	audieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "trottine"}

	audieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "foxtrott"}

	audieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "foxtrot"}

	audieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "アハッ"}

	audieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아하핫"}

	audieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ulalilá"}

	audieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "фокстрот"}

	audieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呀哈"}

	audieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "maravilla"}

	audieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "呀哈"}
)

var (
	audiePhrase = nook.Languages{
		language.AmericanEnglish:      audieAmericanEnglishName,
		language.CanadianFrench:       audieCanadianFrenchName,
		language.Dutch:                audieDutchName,
		language.French:               audieFrenchName,
		language.German:               audieGermanName,
		language.Italian:              audieItalianName,
		language.Japanese:             audieJapaneseName,
		language.Korean:               audieKoreanName,
		language.LatinAmericanSpanish: audieLatinAmericanSpanishName,
		language.Russian:              audieRussianName,
		language.SimplifiedChinese:    audieSimplifiedChineseName,
		language.Spanish:              audieSpanishName,
		language.TraditionalChinese:   audieTraditionalChineseName}
)

var (
	Audie = nook.Villager{
		Character:   audieCharacter,
		Personality: personality.Peppy,
		Phrase:      audiePhrase}
)
