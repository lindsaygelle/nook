package rabbit

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
	// bonbonBirthday represents Bonbon's birthday.
	bonbonBirthday = nook.Birthday{
		Day:   3,
		Month: time.March}
)

var (
	// bonbonCode represents Bonbon's unique code ("rbt17").
	bonbonCode = nook.Code{
		Value: "rbt17"}
)

var (
	// bonbonAmericanEnglishName represents Bonbon's name in American English.
	bonbonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bonbon"}

	// bonbonCanadianFrenchName represents Bonbon's name in Canadian French.
	bonbonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Sylvette"}

	// bonbonDutchName represents Bonbon's name in Dutch.
	bonbonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bonbon"}

	// bonbonFrenchName represents Bonbon's name in French.
	bonbonFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sylvette"}

	// bonbonGermanName represents Bonbon's name in German.
	bonbonGermanName = nook.Name{
		Language: language.German,
		Value:    "Henrike"}

	// bonbonItalianName represents Bonbon's name in Italian.
	bonbonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lolly"}

	// bonbonJapaneseName represents Bonbon's name in Japanese.
	bonbonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ミミィ"}

	// bonbonKoreanName represents Bonbon's name in Korean.
	bonbonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "미미"}

	// bonbonLatinAmericanSpanishName represents Bonbon's name in Latin American Spanish.
	bonbonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Chocolat"}

	// bonbonRussianName represents Bonbon's name in Russian.
	bonbonRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бонбон"}

	// bonbonSimplifiedChineseName represents Bonbon's name in Simplified Chinese.
	bonbonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "妙妙"}

	// bonbonSpanishName represents Bonbon's name in Spanish.
	bonbonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Chocolat"}

	// bonbonTraditionalChineseName represents Bonbon's name in Traditional Chinese.
	bonbonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "妙妙"}
)

var (
	// bonbonName represents Bonbon's name in different languages.
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
	// bonbonCharacter represents Bonbon's character information.
	bonbonCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: bonbonBirthday,
		Code:     bonbonCode,
		Key:      character.Bonbon,
		Gender:   gender.Female,
		Name:     bonbonName,
		Special:  false}
)

var (
	// bonbonAmericanEnglishPhrase represents Bonbon's phrase in American English.
	bonbonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "deelish"}

	// bonbonCanadianFrenchPhrase represents Bonbon's phrase in Canadian French.
	bonbonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "nananère"}

	// bonbonDutchPhrase represents Bonbon's phrase in Dutch.
	bonbonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "heer-lijk"}

	// bonbonFrenchPhrase represents Bonbon's phrase in French.
	bonbonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "nananère"}

	// bonbonGermanPhrase represents Bonbon's phrase in German.
	bonbonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hüppi"}

	// bonbonItalianPhrase represents Bonbon's phrase in Italian.
	bonbonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hurrah"}

	// bonbonJapanesePhrase represents Bonbon's phrase in Japanese.
	bonbonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ヤバッ"}

	// bonbonKoreanPhrase represents Bonbon's phrase in Korean.
	bonbonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어뜨케"}

	// bonbonLatinAmericanSpanishPhrase represents Bonbon's phrase in Latin American Spanish.
	bonbonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "oh lalá"}

	// bonbonRussianPhrase represents Bonbon's phrase in Russian.
	bonbonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "вкусняшка"}

	// bonbonSimplifiedChinesePhrase represents Bonbon's phrase in Simplified Chinese.
	bonbonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "不妙"}

	// bonbonSpanishPhrase represents Bonbon's phrase in Spanish.
	bonbonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "oh lalá"}

	// bonbonTraditionalChinesePhrase represents Bonbon's phrase in Traditional Chinese.
	bonbonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "不妙"}
)

var (
	// bonbonPhrase represents Bonbon's phrase in different languages.
	bonbonPhrase = nook.Languages{
		language.AmericanEnglish:      bonbonAmericanEnglishPhrase,
		language.CanadianFrench:       bonbonCanadianFrenchPhrase,
		language.Dutch:                bonbonDutchPhrase,
		language.French:               bonbonFrenchPhrase,
		language.German:               bonbonGermanPhrase,
		language.Italian:              bonbonItalianPhrase,
		language.Japanese:             bonbonJapanesePhrase,
		language.Korean:               bonbonKoreanPhrase,
		language.LatinAmericanSpanish: bonbonLatinAmericanSpanishPhrase,
		language.Russian:              bonbonRussianPhrase,
		language.SimplifiedChinese:    bonbonSimplifiedChinesePhrase,
		language.Spanish:              bonbonSpanishPhrase,
		language.TraditionalChinese:   bonbonTraditionalChinesePhrase}
)

var (
	// Bonbon represents the character Bonbon with her complete information.
	Bonbon = nook.Villager{
		Character:   bonbonCharacter,
		Personality: personality.Peppy,
		Phrase:      bonbonPhrase}
)
