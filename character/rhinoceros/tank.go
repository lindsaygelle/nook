package rhinoceros

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
	// tankBirthday represents Tank's birthday (May 6th).
	tankBirthday = nook.Birthday{
		Day:   6,
		Month: time.May}
)

var (
	// tankCode represents Tank's unique code ("rhn00").
	tankCode = nook.Code{
		Value: "rhn00"}
)

var (
	// tankAmericanEnglishName represents Tank's name in American English.
	tankAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tank"}

	// tankCanadianFrenchName represents Tank's name in Canadian French.
	tankCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ben"}

	// tankDutchName represents Tank's name in Dutch.
	tankDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tank"}

	// tankFrenchName represents Tank's name in French.
	tankFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ben"}

	// tankGermanName represents Tank's name in German.
	tankGermanName = nook.Name{
		Language: language.German,
		Value:    "Frank"}

	// tankItalianName represents Tank's name in Italian.
	tankItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rinaldo"}

	// tankJapaneseName represents Tank's name in Japanese.
	tankJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "くるぶし"}

	// tankKoreanName represents Tank's name in Korean.
	tankKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "탱크"}

	// tankLatinAmericanSpanishName represents Tank's name in Latin American Spanish.
	tankLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Aníbal"}

	// tankRussianName represents Tank's name in Russian.
	tankRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Танк"}

	// tankSimplifiedChineseName represents Tank's name in Simplified Chinese.
	tankSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿足"}

	// tankSpanishName represents Tank's name in Spanish.
	tankSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Aníbal"}

	// tankTraditionalChineseName represents Tank's name in Traditional Chinese.
	tankTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿足"}
)

var (
	// tankName represents Tank's name in different languages.
	tankName = nook.Languages{
		language.AmericanEnglish:      tankAmericanEnglishName,
		language.CanadianFrench:       tankCanadianFrenchName,
		language.Dutch:                tankDutchName,
		language.French:               tankFrenchName,
		language.German:               tankGermanName,
		language.Italian:              tankItalianName,
		language.Japanese:             tankJapaneseName,
		language.Korean:               tankKoreanName,
		language.LatinAmericanSpanish: tankLatinAmericanSpanishName,
		language.Russian:              tankRussianName,
		language.SimplifiedChinese:    tankSimplifiedChineseName,
		language.Spanish:              tankSpanishName,
		language.TraditionalChinese:   tankTraditionalChineseName}
)

var (
	// tankCharacter represents Tank's character information.
	tankCharacter = nook.Character{
		Animal:   animal.Rhinoceros,
		Birthday: tankBirthday,
		Code:     tankCode,
		Key:      character.Tank,
		Gender:   gender.Male,
		Name:     tankName,
		Special:  false}
)

var (
	// tankAmericanEnglishPhrase represents Tank's phrase in American English.
	tankAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "kerPOW"}

	// tankCanadianFrenchPhrase represents Tank's phrase in Canadian French.
	tankCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "kaboum"}

	// tankDutchPhrase represents Tank's phrase in Dutch.
	tankDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "klaBAM"}

	// tankFrenchPhrase represents Tank's phrase in French.
	tankFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "kaboum"}

	// tankGermanPhrase represents Tank's phrase in German.
	tankGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "rumpel"}

	// tankItalianPhrase represents Tank's phrase in Italian.
	tankItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "zamperù"}

	// tankJapanesePhrase represents Tank's phrase in Japanese.
	tankJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですサイ"}

	// tankKoreanPhrase represents Tank's phrase in Korean.
	tankKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아뿔소"}

	// tankLatinAmericanSpanishPhrase represents Tank's phrase in Latin American Spanish.
	tankLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tapumba"}

	// tankRussianPhrase represents Tank's phrase in Russian.
	tankRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бабах"}

	// tankSimplifiedChinesePhrase represents Tank's phrase in Simplified Chinese.
	tankSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "犀犀"}

	// tankSpanishPhrase represents Tank's phrase in Spanish.
	tankSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tapumba"}

	// tankTraditionalChinesePhrase represents Tank's phrase in Traditional Chinese.
	tankTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "犀犀"}
)

var (
	// tankPhrase represents Tank's phrase in different languages.
	tankPhrase = nook.Languages{
		language.AmericanEnglish:      tankAmericanEnglishPhrase,
		language.CanadianFrench:       tankCanadianFrenchPhrase,
		language.Dutch:                tankDutchPhrase,
		language.French:               tankFrenchPhrase,
		language.German:               tankGermanPhrase,
		language.Italian:              tankItalianPhrase,
		language.Japanese:             tankJapanesePhrase,
		language.Korean:               tankKoreanPhrase,
		language.LatinAmericanSpanish: tankLatinAmericanSpanishPhrase,
		language.Russian:              tankRussianPhrase,
		language.SimplifiedChinese:    tankSimplifiedChinesePhrase,
		language.Spanish:              tankSpanishPhrase,
		language.TraditionalChinese:   tankTraditionalChinesePhrase}
)

var (
	// Tank represents the character Tank with his complete information.
	Tank = nook.Villager{
		Character:   tankCharacter,
		Personality: personality.Jock,
		Phrase:      tankPhrase}
)
