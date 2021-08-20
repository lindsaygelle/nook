package rabbit

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
	ohareBirthday = nook.Birthday{
		Day:   24,
		Month: time.July}
)

var (
	ohareCode = nook.Code{
		Value: "rbt15"}
)

var (
	ohareAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "O'Hare"}

	ohareCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Civet"}

	ohareDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "O'Hare"}

	ohareFrenchName = nook.Name{
		Language: language.French,
		Value:    "Civet"}

	ohareGermanName = nook.Name{
		Language: language.German,
		Value:    "Nico"}

	ohareItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fritz"}

	ohareJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サントス"}

	ohareKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "산토스"}

	ohareLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ernesto"}

	ohareRussianName = nook.Name{
		Language: language.Russian,
		Value:    "О'Хэар"}

	ohareSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "桑托斯"}

	ohareSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ernesto"}

	ohareTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "桑托斯"}
)

var (
	ohareName = nook.Languages{
		language.AmericanEnglish:      ohareAmericanEnglishName,
		language.CanadianFrench:       ohareCanadianFrenchName,
		language.Dutch:                ohareDutchName,
		language.French:               ohareFrenchName,
		language.German:               ohareGermanName,
		language.Italian:              ohareItalianName,
		language.Japanese:             ohareJapaneseName,
		language.Korean:               ohareKoreanName,
		language.LatinAmericanSpanish: ohareLatinAmericanSpanishName,
		language.Russian:              ohareRussianName,
		language.SimplifiedChinese:    ohareSimplifiedChineseName,
		language.Spanish:              ohareSpanishName,
		language.TraditionalChinese:   ohareTraditionalChineseName}
)

var (
	ohareCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: ohareBirthday,
		Code:     ohareCode,
		Key:      character.OHare,
		Gender:   gender.Male,
		Name:     ohareName,
		Special:  false}
)

var (
	ohareAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "amigo"}

	ohareCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "amigo"}

	ohareDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "amigo"}

	ohareFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pécaïre"}

	ohareGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "amigo"}

	ohareItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "aquesí"}

	ohareJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "アミーゴ"}

	ohareKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아미고"}

	ohareLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tralarí"}

	ohareRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "амиго"}

	ohareSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "朋友"}

	ohareSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tralarí"}

	ohareTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "朋友"}
)

var (
	oharePhrase = nook.Languages{
		language.AmericanEnglish:      ohareAmericanEnglishPhrase,
		language.CanadianFrench:       ohareCanadianFrenchPhrase,
		language.Dutch:                ohareDutchPhrase,
		language.French:               ohareFrenchPhrase,
		language.German:               ohareGermanPhrase,
		language.Italian:              ohareItalianPhrase,
		language.Japanese:             ohareJapanesePhrase,
		language.Korean:               ohareKoreanPhrase,
		language.LatinAmericanSpanish: ohareLatinAmericanSpanishPhrase,
		language.Russian:              ohareRussianPhrase,
		language.SimplifiedChinese:    ohareSimplifiedChinesePhrase,
		language.Spanish:              ohareSpanishPhrase,
		language.TraditionalChinese:   ohareTraditionalChinesePhrase}
)

var (
	OHare = nook.Villager{
		Character:   ohareCharacter,
		Personality: personality.Smug,
		Phrase:      oharePhrase}
)
