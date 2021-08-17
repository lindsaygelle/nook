package bearcub

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
	ponchoBirthday = nook.Birthday{
		Day:   2,
		Month: time.January}
)

var (
	ponchoCode = nook.Code{
		Value: "cbr02"}
)

var (
	ponchoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Poncho"}

	ponchoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Théo"}

	ponchoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Poncho"}

	ponchoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Théo"}

	ponchoGermanName = nook.Name{
		Language: language.German,
		Value:    "Toni"}

	ponchoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Poncho"}

	ponchoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ポンチョ"}

	ponchoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "봉추"}

	ponchoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Poncho"}

	ponchoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пончо"}

	ponchoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蓬蓬"}

	ponchoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Poncho"}

	ponchoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蓬蓬"}
)

var (
	ponchoName = nook.Languages{
		language.AmericanEnglish:      ponchoAmericanEnglishName,
		language.CanadianFrench:       ponchoCanadianFrenchName,
		language.Dutch:                ponchoDutchName,
		language.French:               ponchoFrenchName,
		language.German:               ponchoGermanName,
		language.Italian:              ponchoItalianName,
		language.Japanese:             ponchoJapaneseName,
		language.Korean:               ponchoKoreanName,
		language.LatinAmericanSpanish: ponchoLatinAmericanSpanishName,
		language.Russian:              ponchoRussianName,
		language.SimplifiedChinese:    ponchoSimplifiedChineseName,
		language.Spanish:              ponchoSpanishName,
		language.TraditionalChinese:   ponchoTraditionalChineseName}
)

var (
	ponchoCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: ponchoBirthday,
		Code:     ponchoCode,
		Key:      character.Poncho,
		Gender:   gender.Male,
		Name:     ponchoName}
)

var (
	ponchoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "li'l bear"}

	ponchoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "nounours"}

	ponchoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "beertje"}

	ponchoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "nounours"}

	ponchoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bärli"}

	ponchoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "babà"}

	ponchoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "モン"}

	ponchoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "몽"}

	ponchoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "pruah"}

	ponchoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "медвежонок"}

	ponchoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "萌"}

	ponchoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bolita"}

	ponchoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "萌"}
)

var (
	ponchoPhrase = nook.Languages{
		language.AmericanEnglish:      ponchoAmericanEnglishPhrase,
		language.CanadianFrench:       ponchoCanadianFrenchPhrase,
		language.Dutch:                ponchoDutchPhrase,
		language.French:               ponchoFrenchPhrase,
		language.German:               ponchoGermanPhrase,
		language.Italian:              ponchoItalianPhrase,
		language.Japanese:             ponchoJapanesePhrase,
		language.Korean:               ponchoKoreanPhrase,
		language.LatinAmericanSpanish: ponchoLatinAmericanSpanishPhrase,
		language.Russian:              ponchoRussianPhrase,
		language.SimplifiedChinese:    ponchoSimplifiedChinesePhrase,
		language.Spanish:              ponchoSpanishPhrase,
		language.TraditionalChinese:   ponchoTraditionalChinesePhrase}
)

var (
	Poncho = nook.Villager{
		Character:   ponchoCharacter,
		Personality: personality.Jock,
		Phrase:      ponchoPhrase}
)
