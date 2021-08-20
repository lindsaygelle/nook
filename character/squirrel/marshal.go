package squirrel

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
		Value:    "Mathéo"}

	marshalDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Marshal"}

	marshalFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mathéo"}

	marshalGermanName = nook.Name{
		Language: language.German,
		Value:    "Huschke"}

	marshalItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Scott"}

	marshalJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジュン"}

	marshalKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "쭈니"}

	marshalLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Munchi"}

	marshalRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Маршал"}

	marshalSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小润"}

	marshalSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Munchi"}

	marshalTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小潤"}
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
		Animal:   animal.Squirrel,
		Birthday: marshalBirthday,
		Code:     marshalCode,
		Key:      character.Marshal,
		Gender:   gender.Male,
		Name:     marshalName,
		Special:  false}
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
		language.AmericanEnglish:      marshalAmericanEnglishPhrase,
		language.CanadianFrench:       marshalCanadianFrenchPhrase,
		language.Dutch:                marshalDutchPhrase,
		language.French:               marshalFrenchPhrase,
		language.German:               marshalGermanPhrase,
		language.Italian:              marshalItalianPhrase,
		language.Japanese:             marshalJapanesePhrase,
		language.Korean:               marshalKoreanPhrase,
		language.LatinAmericanSpanish: marshalLatinAmericanSpanishPhrase,
		language.Russian:              marshalRussianPhrase,
		language.SimplifiedChinese:    marshalSimplifiedChinesePhrase,
		language.Spanish:              marshalSpanishPhrase,
		language.TraditionalChinese:   marshalTraditionalChinesePhrase}
)

var (
	Marshal = nook.Villager{
		Character:   marshalCharacter,
		Personality: personality.Smug,
		Phrase:      marshalPhrase}
)
