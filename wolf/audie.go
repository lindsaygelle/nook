package wolf

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Monicatrottine"}

	audieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Audiefoxtrot"}

	audieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Monicatrottine"}

	audieGermanName = nook.Name{
		Language: language.German,
		Value:    "Katharinafoxtrott"}

	audieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lupiliafoxtrot"}

	audieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "モニカアハッ"}

	audieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "모니카아하핫"}

	audieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mónicaulalilá"}

	audieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Одифокстрот"}

	audieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "莫妮卡呀哈"}

	audieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mónicamaravilla"}

	audieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "莫妮卡呀哈"}
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
		Animal:   Wolf,
		Birthday: audieBirthday,
		Code:     audieCode,
		Gender:   nook.Female,
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
		Personality: nook.Peppy,
		Phrase:      audiePhrase}
)
