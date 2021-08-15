package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	admiralBirthday = nook.Birthday{
		Day:   27,
		Month: time.January}
)

var (
	admiralCode = nook.Code{
		Value: "brd06"}
)

var (
	admiralAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Admiral"}

	admiralCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Maréchal"}

	admiralDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Admiral"}

	admiralFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maréchal"}

	admiralGermanName = nook.Name{
		Language: language.German,
		Value:    "Ansgar"}

	admiralItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Adolfo"}

	admiralJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "イッテツ"}

	admiralKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "일섭"}

	admiralLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Avutardo"}

	admiralRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Адмирал"}

	admiralSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "李彻"}

	admiralSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Avutardo"}

	admiralTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "李徹"}
)

var (
	admiralName = nook.Languages{
		language.AmericanEnglish:      admiralAmericanEnglishName,
		language.CanadianFrench:       admiralCanadianFrenchName,
		language.Dutch:                admiralDutchName,
		language.French:               admiralFrenchName,
		language.German:               admiralGermanName,
		language.Italian:              admiralItalianName,
		language.Japanese:             admiralJapaneseName,
		language.Korean:               admiralKoreanName,
		language.LatinAmericanSpanish: admiralLatinAmericanSpanishName,
		language.Russian:              admiralRussianName,
		language.SimplifiedChinese:    admiralSimplifiedChineseName,
		language.Spanish:              admiralSpanishName,
		language.TraditionalChinese:   admiralTraditionalChineseName}
)

var (
	admiralCharacter = nook.Character{
		Animal:   Bird,
		Birthday: admiralBirthday,
		Code:     admiralCode,
		Gender:   nook.Male,
		Name:     admiralName}
)

var (
	admiralAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "aye aye"}

	admiralCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "cornichon"}

	admiralDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "yep yep"}

	admiralFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "cornichon"}

	admiralGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "oh, ja"}

	admiralItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cipissi"}

	admiralJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ってか"}

	admiralKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "오예"}

	admiralLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "piuc-piuc"}

	admiralRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "чиф-чиф"}

	admiralSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "贯彻"}

	admiralSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "avechucho"}

	admiralTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "貫徹"}
)

var (
	admiralPhrase = nook.Languages{
		language.AmericanEnglish:      admiralAmericanEnglishName,
		language.CanadianFrench:       admiralCanadianFrenchName,
		language.Dutch:                admiralDutchName,
		language.French:               admiralFrenchName,
		language.German:               admiralGermanName,
		language.Italian:              admiralItalianName,
		language.Japanese:             admiralJapaneseName,
		language.Korean:               admiralKoreanName,
		language.LatinAmericanSpanish: admiralLatinAmericanSpanishName,
		language.Russian:              admiralRussianName,
		language.SimplifiedChinese:    admiralSimplifiedChineseName,
		language.Spanish:              admiralSpanishName,
		language.TraditionalChinese:   admiralTraditionalChineseName}
)

var (
	Admiral = nook.Villager{
		Character:   admiralCharacter,
		Personality: nook.Cranky,
		Phrase:      admiralPhrase}
)
