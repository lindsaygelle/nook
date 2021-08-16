package rabbit

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	bonbonBirthday = nook.Birthday{
		Day:   3,
		Month: time.March}
)

var (
	bonbonCode = nook.Code{
		Value: "rbt17"}
)

var (
	bonbonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bonbon"}

	bonbonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Sylvette"}

	bonbonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bonbon"}

	bonbonFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sylvette"}

	bonbonGermanName = nook.Name{
		Language: language.German,
		Value:    "Henrike"}

	bonbonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lolly"}

	bonbonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ミミィ"}

	bonbonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "미미"}

	bonbonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Chocolat"}

	bonbonRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бонбон"}

	bonbonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "妙妙"}

	bonbonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Chocolat"}

	bonbonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "妙妙"}
)

var (
	bonbonName = nook.Languages{
		language.AmericanEnglish:      bonbonAmericanEnglishName,
		language.CanadianFrench:       bonbonCanadianFrenchName,
		language.Dutch:                bonbonDutchName,
		language.French:               bonbonFrenchName,
		language.German:               bonbonGermanName,
		language.Italian:              bonbonItalianName,
		language.Japanese:             bonbonJapaneseName,
		language.Korean:               bonbonKoreanName,
		language.LatinAmericanSpanish: bonbonLatinAmericanSpanishName,
		language.Russian:              bonbonRussianName,
		language.SimplifiedChinese:    bonbonSimplifiedChineseName,
		language.Spanish:              bonbonSpanishName,
		language.TraditionalChinese:   bonbonTraditionalChineseName}
)

var (
	bonbonCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: bonbonBirthday,
		Code:     bonbonCode,
		Gender:   gender.Female,
		Name:     bonbonName}
)

var (
	bonbonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "deelish"}

	bonbonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "nananère"}

	bonbonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "heer-lijk"}

	bonbonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "nananère"}

	bonbonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hüppi"}

	bonbonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hurrah"}

	bonbonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ヤバッ"}

	bonbonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어뜨케"}

	bonbonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "oh lalá"}

	bonbonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "вкусняшка"}

	bonbonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "不妙"}

	bonbonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "oh lalá"}

	bonbonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "不妙"}
)

var (
	bonbonPhrase = nook.Languages{
		language.AmericanEnglish:      bonbonAmericanEnglishName,
		language.CanadianFrench:       bonbonCanadianFrenchName,
		language.Dutch:                bonbonDutchName,
		language.French:               bonbonFrenchName,
		language.German:               bonbonGermanName,
		language.Italian:              bonbonItalianName,
		language.Japanese:             bonbonJapaneseName,
		language.Korean:               bonbonKoreanName,
		language.LatinAmericanSpanish: bonbonLatinAmericanSpanishName,
		language.Russian:              bonbonRussianName,
		language.SimplifiedChinese:    bonbonSimplifiedChineseName,
		language.Spanish:              bonbonSpanishName,
		language.TraditionalChinese:   bonbonTraditionalChineseName}
)

var (
	Bonbon = nook.Villager{
		Character:   bonbonCharacter,
		Personality: personality.Peppy,
		Phrase:      bonbonPhrase}
)
