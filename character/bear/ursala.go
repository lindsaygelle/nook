package bear

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
	ursalaBirthday = nook.Birthday{
		Day:   16,
		Month: time.January}
)

var (
	ursalaCode = nook.Code{
		Value: "bea08"}
)

var (
	ursalaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ursala"}

	ursalaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Oursula"}

	ursalaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ursala"}

	ursalaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Oursula"}

	ursalaGermanName = nook.Name{
		Language: language.German,
		Value:    "Ursula"}

	ursalaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ursula"}

	ursalaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ネーヤ"}

	ursalaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "네이아"}

	ursalaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Úrsula"}

	ursalaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Урсала"}

	ursalaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "妮雅"}

	ursalaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Úrsula"}

	ursalaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "妮雅"}
)

var (
	ursalaName = nook.Languages{
		language.AmericanEnglish:      ursalaAmericanEnglishName,
		language.CanadianFrench:       ursalaCanadianFrenchName,
		language.Dutch:                ursalaDutchName,
		language.French:               ursalaFrenchName,
		language.German:               ursalaGermanName,
		language.Italian:              ursalaItalianName,
		language.Japanese:             ursalaJapaneseName,
		language.Korean:               ursalaKoreanName,
		language.LatinAmericanSpanish: ursalaLatinAmericanSpanishName,
		language.Russian:              ursalaRussianName,
		language.SimplifiedChinese:    ursalaSimplifiedChineseName,
		language.Spanish:              ursalaSpanishName,
		language.TraditionalChinese:   ursalaTraditionalChineseName}
)

var (
	ursalaCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: ursalaBirthday,
		Code:     ursalaCode,
		Key:      character.Ursala,
		Gender:   gender.Female,
		Name:     ursalaName}
)

var (
	ursalaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "grooomph"}

	ursalaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "groumpf"}

	ursalaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "oemf"}

	ursalaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "groumpf"}

	ursalaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brumbrum"}

	ursalaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gruuunf"}

	ursalaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "やーねぇ"}

	ursalaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그라지"}

	ursalaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grumf"}

	ursalaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "гррум"}

	ursalaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呀呐"}

	ursalaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "grumf"}

	ursalaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "呀吶"}
)

var (
	ursalaPhrase = nook.Languages{
		language.AmericanEnglish:      ursalaAmericanEnglishPhrase,
		language.CanadianFrench:       ursalaCanadianFrenchPhrase,
		language.Dutch:                ursalaDutchPhrase,
		language.French:               ursalaFrenchPhrase,
		language.German:               ursalaGermanPhrase,
		language.Italian:              ursalaItalianPhrase,
		language.Japanese:             ursalaJapanesePhrase,
		language.Korean:               ursalaKoreanPhrase,
		language.LatinAmericanSpanish: ursalaLatinAmericanSpanishPhrase,
		language.Russian:              ursalaRussianPhrase,
		language.SimplifiedChinese:    ursalaSimplifiedChinesePhrase,
		language.Spanish:              ursalaSpanishPhrase,
		language.TraditionalChinese:   ursalaTraditionalChinesePhrase}
)

var (
	Ursala = nook.Villager{
		Character:   ursalaCharacter,
		Personality: personality.BigSister,
		Phrase:      ursalaPhrase}
)
