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
	// marshalBirthday represents Marshal's birthday (September 29th).
	marshalBirthday = nook.Birthday{
		Day:   29,
		Month: time.September}
)

var (
	// marshalCode represents Marshal's unique code ("squ17").
	marshalCode = nook.Code{
		Value: "squ17"}
)

var (
	// marshalAmericanEnglishName represents Marshal's name in American English.
	marshalAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Marshal"}

	// marshalCanadianFrenchName represents Marshal's name in Canadian French.
	marshalCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mathéo"}

	// marshalDutchName represents Marshal's name in Dutch.
	marshalDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Marshal"}

	// marshalFrenchName represents Marshal's name in French.
	marshalFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mathéo"}

	// marshalGermanName represents Marshal's name in German.
	marshalGermanName = nook.Name{
		Language: language.German,
		Value:    "Huschke"}

	// marshalItalianName represents Marshal's name in Italian.
	marshalItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Scott"}

	// marshalJapaneseName represents Marshal's name in Japanese.
	marshalJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジュン"}

	// marshalKoreanName represents Marshal's name in Korean.
	marshalKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "쭈니"}

	// marshalLatinAmericanSpanishName represents Marshal's name in Latin American Spanish.
	marshalLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Munchi"}

	// marshalRussianName represents Marshal's name in Russian.
	marshalRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Маршал"}

	// marshalSimplifiedChineseName represents Marshal's name in Simplified Chinese.
	marshalSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小润"}

	// marshalSpanishName represents Marshal's name in Spanish.
	marshalSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Munchi"}

	// marshalTraditionalChineseName represents Marshal's name in Traditional Chinese.
	marshalTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小潤"}
)

var (
	// marshalName represents Marshal's name in different languages.
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
	// marshalCharacter represents Marshal's character information.
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
	// marshalAmericanEnglishPhrase represents Marshal's phrase in American English.
	marshalAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sulky"}

	// marshalCanadianFrenchPhrase represents Marshal's phrase in Canadian French.
	marshalCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "carpe diem"}

	// marshalDutchPhrase represents Marshal's phrase in Dutch.
	marshalDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "knorrepot"}

	// marshalFrenchPhrase represents Marshal's phrase in French.
	marshalFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "nanache"}

	// marshalGermanPhrase represents Marshal's phrase in German.
	marshalGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "huschhusch"}

	// marshalItalianPhrase represents Marshal's phrase in Italian.
	marshalItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ardillà"}

	// marshalJapanesePhrase represents Marshal's phrase in Japanese.
	marshalJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "あながち"}

	// marshalKoreanPhrase represents Marshal's phrase in Korean.
	marshalKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어차피"}

	// marshalLatinAmericanSpanishPhrase represents Marshal's phrase in Latin American Spanish.
	marshalLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tecuén"}

	// marshalRussianPhrase represents Marshal's phrase in Russian.
	marshalRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тихоня"}

	// marshalSimplifiedChinesePhrase represents Marshal's phrase in Simplified Chinese.
	marshalSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "不管怎样"}

	// marshalSpanishPhrase represents Marshal's phrase in Spanish.
	marshalSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tecuén"}

	// marshalTraditionalChinesePhrase represents Marshal's phrase in Traditional Chinese.
	marshalTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "不管怎樣"}
)

var (
	// marshalPhrase represents Marshal's phrase in different languages.
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
	// Marshal represents the character Marshal with his complete information.
	Marshal = nook.Villager{
		Character:   marshalCharacter,
		Personality: personality.Smug,
		Phrase:      marshalPhrase}
)
