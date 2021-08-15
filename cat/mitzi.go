package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	mitziBirthday = nook.Birthday{
		Day:   25,
		Month: time.September}
)

var (
	mitziCode = nook.Code{
		Value: "cat01"}
)

var (
	mitziAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mitzi"}

	mitziCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mistigrimeeeeeh"}

	mitziDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Mitzimieuw"}

	mitziFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mistigrimeeeeeh"}

	mitziGermanName = nook.Name{
		Language: language.German,
		Value:    "Miezimuh"}

	mitziItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Zampettamiao"}

	mitziJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マールニャー"}

	mitziKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마르야～옹"}

	mitziLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Misimisi-misi"}

	mitziRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Митцимяу-мяу"}

	mitziSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "圆圆喵"}

	mitziSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Misimisi-misi"}

	mitziTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "圓圓喵"}
)

var (
	mitziName = nook.Languages{
		language.AmericanEnglish:      mitziAmericanEnglishName,
		language.CanadianFrench:       mitziCanadianFrenchName,
		language.Dutch:                mitziDutchName,
		language.French:               mitziFrenchName,
		language.German:               mitziGermanName,
		language.Italian:              mitziItalianName,
		language.Japanese:             mitziJapaneseName,
		language.Korean:               mitziKoreanName,
		language.LatinAmericanSpanish: mitziLatinAmericanSpanishName,
		language.Russian:              mitziRussianName,
		language.SimplifiedChinese:    mitziSimplifiedChineseName,
		language.Spanish:              mitziSpanishName,
		language.TraditionalChinese:   mitziTraditionalChineseName}
)

var (
	mitziCharacter = nook.Character{
		Animal:   Cat,
		Birthday: mitziBirthday,
		Code:     mitziCode,
		Gender:   nook.Female,
		Name:     mitziName}
)

var (
	mitziAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "mew"}

	mitziCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "meeeeeh"}

	mitziDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mieuw"}

	mitziFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "meeeeeh"}

	mitziGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "muh"}

	mitziItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "miao"}

	mitziJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ニャー"}

	mitziKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "야～옹"}

	mitziLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "misi-misi"}

	mitziRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мяу-мяу"}

	mitziSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喵"}

	mitziSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "misi-misi"}

	mitziTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喵"}
)

var (
	mitziPhrase = nook.Languages{
		language.AmericanEnglish:      mitziAmericanEnglishName,
		language.CanadianFrench:       mitziCanadianFrenchName,
		language.Dutch:                mitziDutchName,
		language.French:               mitziFrenchName,
		language.German:               mitziGermanName,
		language.Italian:              mitziItalianName,
		language.Japanese:             mitziJapaneseName,
		language.Korean:               mitziKoreanName,
		language.LatinAmericanSpanish: mitziLatinAmericanSpanishName,
		language.Russian:              mitziRussianName,
		language.SimplifiedChinese:    mitziSimplifiedChineseName,
		language.Spanish:              mitziSpanishName,
		language.TraditionalChinese:   mitziTraditionalChineseName}
)

var (
	Mitzi = nook.Villager{
		Character:   mitziCharacter,
		Personality: nook.Normal,
		Phrase:      mitziPhrase}
)
