package monkey

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	simonBirthday = nook.Birthday{
		Day:   19,
		Month: time.January}
)

var (
	simonCode = nook.Code{
		Value: "mnk02"}
)

var (
	simonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Simon"}

	simonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Simon"}

	simonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Simon"}

	simonFrenchName = nook.Name{
		Language: language.French,
		Value:    "Simon"}

	simonGermanName = nook.Name{
		Language: language.German,
		Value:    "Simon"}

	simonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Simon"}

	simonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エテキチ"}

	simonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "시몬"}

	simonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Simón"}

	simonRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Саймон"}

	simonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "远仁"}

	simonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Simón"}

	simonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "遠仁"}
)

var (
	simonName = nook.Languages{
		language.AmericanEnglish:      simonAmericanEnglishName,
		language.CanadianFrench:       simonCanadianFrenchName,
		language.Dutch:                simonDutchName,
		language.French:               simonFrenchName,
		language.German:               simonGermanName,
		language.Italian:              simonItalianName,
		language.Japanese:             simonJapaneseName,
		language.Korean:               simonKoreanName,
		language.LatinAmericanSpanish: simonLatinAmericanSpanishName,
		language.Russian:              simonRussianName,
		language.SimplifiedChinese:    simonSimplifiedChineseName,
		language.Spanish:              simonSpanishName,
		language.TraditionalChinese:   simonTraditionalChineseName}
)

var (
	simonCharacter = nook.Character{
		Animal:   animal.Monkey,
		Birthday: simonBirthday,
		Code:     simonCode,
		Gender:   gender.Male,
		Name:     simonName}
)

var (
	simonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "zzzook"}

	simonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "zzziik"}

	simonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "zzzoeh oeh"}

	simonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "zzziik"}

	simonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bannabanna"}

	simonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "zzzook"}

	simonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でござる"}

	simonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "하옵니다"}

	simonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chimpa"}

	simonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "дружок"}

	simonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "猿猿"}

	simonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chimpa"}

	simonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "猿猿"}
)

var (
	simonPhrase = nook.Languages{
		language.AmericanEnglish:      simonAmericanEnglishName,
		language.CanadianFrench:       simonCanadianFrenchName,
		language.Dutch:                simonDutchName,
		language.French:               simonFrenchName,
		language.German:               simonGermanName,
		language.Italian:              simonItalianName,
		language.Japanese:             simonJapaneseName,
		language.Korean:               simonKoreanName,
		language.LatinAmericanSpanish: simonLatinAmericanSpanishName,
		language.Russian:              simonRussianName,
		language.SimplifiedChinese:    simonSimplifiedChineseName,
		language.Spanish:              simonSpanishName,
		language.TraditionalChinese:   simonTraditionalChineseName}
)

var (
	Simon = nook.Villager{
		Character:   simonCharacter,
		Personality: personality.Lazy,
		Phrase:      simonPhrase}
)
