package eagle

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
	quetzalBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	quetzalCode = nook.Code{
		Value: ""}
)

var (
	quetzalAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Quetzal"}

	quetzalCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	quetzalDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	quetzalFrenchName = nook.Name{
		Language: language.French,
		Value:    "Quetzal"}

	quetzalGermanName = nook.Name{
		Language: language.German,
		Value:    "Enrique"}

	quetzalItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Quetzal"}

	quetzalJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハチェット"}

	quetzalKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	quetzalLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	quetzalRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	quetzalSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "战斧"}

	quetzalSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Quetzal"}

	quetzalTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	quetzalName = nook.Languages{
		language.AmericanEnglish:      quetzalAmericanEnglishName,
		language.CanadianFrench:       quetzalCanadianFrenchName,
		language.Dutch:                quetzalDutchName,
		language.French:               quetzalFrenchName,
		language.German:               quetzalGermanName,
		language.Italian:              quetzalItalianName,
		language.Japanese:             quetzalJapaneseName,
		language.Korean:               quetzalKoreanName,
		language.LatinAmericanSpanish: quetzalLatinAmericanSpanishName,
		language.Russian:              quetzalRussianName,
		language.SimplifiedChinese:    quetzalSimplifiedChineseName,
		language.Spanish:              quetzalSpanishName,
		language.TraditionalChinese:   quetzalTraditionalChineseName}
)

var (
	quetzalCharacter = nook.Character{
		Animal:   animal.Eagle,
		Birthday: quetzalBirthday,
		Code:     quetzalCode,
		Key:      character.Quetzal,
		Gender:   gender.Male,
		Name:     quetzalName}
)

var (
	quetzalAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "SKREEE"}

	quetzalCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	quetzalDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	quetzalFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "couêêêtzal"}

	quetzalGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "KRIKRI"}

	quetzalItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "OCCHIO"}

	quetzalJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ゲロッパ"}

	quetzalKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	quetzalLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	quetzalRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	quetzalSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喳喳"}

	quetzalSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cuiii"}

	quetzalTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	quetzalPhrase = nook.Languages{
		language.AmericanEnglish:      quetzalAmericanEnglishPhrase,
		language.CanadianFrench:       quetzalCanadianFrenchPhrase,
		language.Dutch:                quetzalDutchPhrase,
		language.French:               quetzalFrenchPhrase,
		language.German:               quetzalGermanPhrase,
		language.Italian:              quetzalItalianPhrase,
		language.Japanese:             quetzalJapanesePhrase,
		language.Korean:               quetzalKoreanPhrase,
		language.LatinAmericanSpanish: quetzalLatinAmericanSpanishPhrase,
		language.Russian:              quetzalRussianPhrase,
		language.SimplifiedChinese:    quetzalSimplifiedChinesePhrase,
		language.Spanish:              quetzalSpanishPhrase,
		language.TraditionalChinese:   quetzalTraditionalChinesePhrase}
)

var (
	Quetzal = nook.Villager{
		Character:   quetzalCharacter,
		Personality: personality.Jock,
		Phrase:      quetzalPhrase}
)
