package ostrich

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
	queenieBirthday = nook.Birthday{
		Day:   13,
		Month: time.November}
)

var (
	queenieCode = nook.Code{
		Value: "ost00"}
)

var (
	queenieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Queenie"}

	queenieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Reine"}

	queenieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Queenie"}

	queenieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Reine"}

	queenieGermanName = nook.Name{
		Language: language.German,
		Value:    "Isabella"}

	queenieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Regina"}

	queenieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タキュ"}

	queenieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "택주"}

	queenieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Marujita"}

	queenieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Квини"}

	queenieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "果果"}

	queenieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Marujita"}

	queenieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "果果"}
)

var (
	queenieName = nook.Languages{
		language.AmericanEnglish:      queenieAmericanEnglishName,
		language.CanadianFrench:       queenieCanadianFrenchName,
		language.Dutch:                queenieDutchName,
		language.French:               queenieFrenchName,
		language.German:               queenieGermanName,
		language.Italian:              queenieItalianName,
		language.Japanese:             queenieJapaneseName,
		language.Korean:               queenieKoreanName,
		language.LatinAmericanSpanish: queenieLatinAmericanSpanishName,
		language.Russian:              queenieRussianName,
		language.SimplifiedChinese:    queenieSimplifiedChineseName,
		language.Spanish:              queenieSpanishName,
		language.TraditionalChinese:   queenieTraditionalChineseName}
)

var (
	queenieCharacter = nook.Character{
		Animal:   animal.Ostrich,
		Birthday: queenieBirthday,
		Code:     queenieCode,
		Key:      character.Queenie,
		Gender:   gender.Female,
		Name:     queenieName}
)

var (
	queenieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "chicken"}

	queenieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "poupoule"}

	queenieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kippie"}

	queenieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "poupoule"}

	queenieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hühnchen"}

	queenieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "polpetta"}

	queenieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "やっぱし"}

	queenieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "역시"}

	queenieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "plumifá"}

	queenieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "цыпленок"}

	queenieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "果然"}

	queenieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "gallina"}

	queenieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "果然"}
)

var (
	queeniePhrase = nook.Languages{
		language.AmericanEnglish:      queenieAmericanEnglishName,
		language.CanadianFrench:       queenieCanadianFrenchName,
		language.Dutch:                queenieDutchName,
		language.French:               queenieFrenchName,
		language.German:               queenieGermanName,
		language.Italian:              queenieItalianName,
		language.Japanese:             queenieJapaneseName,
		language.Korean:               queenieKoreanName,
		language.LatinAmericanSpanish: queenieLatinAmericanSpanishName,
		language.Russian:              queenieRussianName,
		language.SimplifiedChinese:    queenieSimplifiedChineseName,
		language.Spanish:              queenieSpanishName,
		language.TraditionalChinese:   queenieTraditionalChineseName}
)

var (
	Queenie = nook.Villager{
		Character:   queenieCharacter,
		Personality: personality.Snooty,
		Phrase:      queeniePhrase}
)
