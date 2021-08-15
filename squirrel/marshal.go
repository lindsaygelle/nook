package squirrel

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	marshalBirthday = nook.Birthday{
		Day:   29,
		Month: time.September}
)

var (
	marshalCode = nook.Code{
		Value: "squ17"}
)

var (
	marshalAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Marshal"}

	marshalCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mathéocarpe diem"}

	marshalDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Marshalknorrepot"}

	marshalFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mathéonanache"}

	marshalGermanName = nook.Name{
		Language: language.German,
		Value:    "Huschkehuschhusch"}

	marshalItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Scottardillà"}

	marshalJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジュンあながち"}

	marshalKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "쭈니어차피"}

	marshalLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Munchitecuén"}

	marshalRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Маршалтихоня"}

	marshalSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小润不管怎样"}

	marshalSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Munchitecuén"}

	marshalTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小潤不管怎樣"}
)

var (
	marshalName = nook.Languages{
		language.AmericanEnglish:      marshalAmericanEnglishName,
		language.CanadianFrench:       marshalCanadianFrenchName,
		language.Dutch:                marshalDutchName,
		language.French:               marshalFrenchName,
		language.German:               marshalGermanName,
		language.Italian:              marshalItalianName,
		language.Japanese:             marshalJapaneseName,
		language.Korean:               marshalKoreanName,
		language.LatinAmericanSpanish: marshalLatinAmericanSpanishName,
		language.Russian:              marshalRussianName,
		language.SimplifiedChinese:    marshalSimplifiedChineseName,
		language.Spanish:              marshalSpanishName,
		language.TraditionalChinese:   marshalTraditionalChineseName}
)

var (
	marshalCharacter = nook.Character{
		Animal:   Squirrel,
		Birthday: marshalBirthday,
		Code:     marshalCode,
		Gender:   nook.Male,
		Name:     marshalName}
)

var (
	marshalAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sulky"}

	marshalCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "carpe diem"}

	marshalDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "knorrepot"}

	marshalFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "nanache"}

	marshalGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "huschhusch"}

	marshalItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ardillà"}

	marshalJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "あながち"}

	marshalKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어차피"}

	marshalLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tecuén"}

	marshalRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тихоня"}

	marshalSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "不管怎样"}

	marshalSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tecuén"}

	marshalTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "不管怎樣"}
)

var (
	marshalPhrase = nook.Languages{
		language.AmericanEnglish:      marshalAmericanEnglishName,
		language.CanadianFrench:       marshalCanadianFrenchName,
		language.Dutch:                marshalDutchName,
		language.French:               marshalFrenchName,
		language.German:               marshalGermanName,
		language.Italian:              marshalItalianName,
		language.Japanese:             marshalJapaneseName,
		language.Korean:               marshalKoreanName,
		language.LatinAmericanSpanish: marshalLatinAmericanSpanishName,
		language.Russian:              marshalRussianName,
		language.SimplifiedChinese:    marshalSimplifiedChineseName,
		language.Spanish:              marshalSpanishName,
		language.TraditionalChinese:   marshalTraditionalChineseName}
)

var (
	Marshal = nook.Villager{
		Character:   marshalCharacter,
		Personality: nook.Smug,
		Phrase:      marshalPhrase}
)
