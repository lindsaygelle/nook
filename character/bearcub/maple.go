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
	mapleBirthday = nook.Birthday{
		Day:   15,
		Month: time.June}
)

var (
	mapleCode = nook.Code{
		Value: "cbr01"}
)

var (
	mapleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Maple"}

	mapleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Léa"}

	mapleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Maple"}

	mapleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Léa"}

	mapleGermanName = nook.Name{
		Language: language.German,
		Value:    "Mona"}

	mapleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Dulcinea"}

	mapleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "メープル"}

	mapleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "메이첼"}

	mapleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Dulce"}

	mapleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мейпл"}

	mapleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小枫"}

	mapleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Dulce"}

	mapleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小楓"}
)

var (
	mapleName = nook.Languages{
		language.AmericanEnglish:      mapleAmericanEnglishName,
		language.CanadianFrench:       mapleCanadianFrenchName,
		language.Dutch:                mapleDutchName,
		language.French:               mapleFrenchName,
		language.German:               mapleGermanName,
		language.Italian:              mapleItalianName,
		language.Japanese:             mapleJapaneseName,
		language.Korean:               mapleKoreanName,
		language.LatinAmericanSpanish: mapleLatinAmericanSpanishName,
		language.Russian:              mapleRussianName,
		language.SimplifiedChinese:    mapleSimplifiedChineseName,
		language.Spanish:              mapleSpanishName,
		language.TraditionalChinese:   mapleTraditionalChineseName}
)

var (
	mapleCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: mapleBirthday,
		Code:     mapleCode,
		Key:      character.Maple,
		Gender:   gender.Female,
		Name:     mapleName}
)

var (
	mapleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "honey"}

	mapleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chouchou"}

	mapleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "zoeterd"}

	mapleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "chouchou"}

	mapleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "darling"}

	mapleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "miele"}

	mapleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だベア"}

	mapleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "저기요"}

	mapleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "mielmiel"}

	mapleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "лапушка"}

	mapleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "熊仔"}

	mapleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "caramelito"}

	mapleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "熊仔"}
)

var (
	maplePhrase = nook.Languages{
		language.AmericanEnglish:      mapleAmericanEnglishName,
		language.CanadianFrench:       mapleCanadianFrenchName,
		language.Dutch:                mapleDutchName,
		language.French:               mapleFrenchName,
		language.German:               mapleGermanName,
		language.Italian:              mapleItalianName,
		language.Japanese:             mapleJapaneseName,
		language.Korean:               mapleKoreanName,
		language.LatinAmericanSpanish: mapleLatinAmericanSpanishName,
		language.Russian:              mapleRussianName,
		language.SimplifiedChinese:    mapleSimplifiedChineseName,
		language.Spanish:              mapleSpanishName,
		language.TraditionalChinese:   mapleTraditionalChineseName}
)

var (
	Maple = nook.Villager{
		Character:   mapleCharacter,
		Personality: personality.Normal,
		Phrase:      maplePhrase}
)
