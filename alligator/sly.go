package alligator

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	slyBirthday = nook.Birthday{
		Day:   15,
		Month: time.November}
)

var (
	slyCode = nook.Code{
		Value: "crd06"}
)

var (
	slyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sly"}

	slyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Chuckrepos"}

	slyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Slyingerukt"}

	slyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chuckrepos"}

	slyGermanName = nook.Name{
		Language: language.German,
		Value:    "Steveschnapp"}

	slyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Dribertociaf"}

	slyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハイドカサコソ"}

	slyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "하이드부스럭"}

	slyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Estallónsmack"}

	slyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Слайура-а-а"}

	slySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "海德唰唰"}

	slySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Estallóncapón"}

	slyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "海德唰唰"}
)

var (
	slyName = nook.Languages{
		language.AmericanEnglish:      slyAmericanEnglishName,
		language.CanadianFrench:       slyCanadianFrenchName,
		language.Dutch:                slyDutchName,
		language.French:               slyFrenchName,
		language.German:               slyGermanName,
		language.Italian:              slyItalianName,
		language.Japanese:             slyJapaneseName,
		language.Korean:               slyKoreanName,
		language.LatinAmericanSpanish: slyLatinAmericanSpanishName,
		language.Russian:              slyRussianName,
		language.SimplifiedChinese:    slySimplifiedChineseName,
		language.Spanish:              slySpanishName,
		language.TraditionalChinese:   slyTraditionalChineseName}
)

var (
	slyCharacter = nook.Character{
		Animal:   Alligator,
		Birthday: slyBirthday,
		Code:     slyCode,
		Gender:   nook.Male,
		Name:     slyName}
)

var (
	slyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hoo-rah"}

	slyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "repos"}

	slyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ingerukt"}

	slyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "repos"}

	slyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnapp"}

	slyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ciaf"}

	slyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "カサコソ"}

	slyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "부스럭"}

	slyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "smack"}

	slyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ура-а-а"}

	slySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唰唰"}

	slySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "capón"}

	slyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唰唰"}
)

var (
	slyPhrase = nook.Languages{
		language.AmericanEnglish:      slyAmericanEnglishName,
		language.CanadianFrench:       slyCanadianFrenchName,
		language.Dutch:                slyDutchName,
		language.French:               slyFrenchName,
		language.German:               slyGermanName,
		language.Italian:              slyItalianName,
		language.Japanese:             slyJapaneseName,
		language.Korean:               slyKoreanName,
		language.LatinAmericanSpanish: slyLatinAmericanSpanishName,
		language.Russian:              slyRussianName,
		language.SimplifiedChinese:    slySimplifiedChineseName,
		language.Spanish:              slySpanishName,
		language.TraditionalChinese:   slyTraditionalChineseName}
)

var (
	Sly = nook.Villager{
		Character:   slyCharacter,
		Personality: nook.Jock,
		Phrase:      slyPhrase}
)
