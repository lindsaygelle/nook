package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
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
		Value:    "Mistigri"}

	mitziDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Mitzi"}

	mitziFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mistigri"}

	mitziGermanName = nook.Name{
		Language: language.German,
		Value:    "Miezi"}

	mitziItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Zampetta"}

	mitziJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マール"}

	mitziKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마르"}

	mitziLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Misi"}

	mitziRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Митци"}

	mitziSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "圆圆"}

	mitziSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Misi"}

	mitziTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "圓圓"}
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
		Animal:   animal.Cat,
		Birthday: mitziBirthday,
		Code:     mitziCode,
		Gender:   gender.Female,
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
		Personality: personality.Normal,
		Phrase:      mitziPhrase}
)
