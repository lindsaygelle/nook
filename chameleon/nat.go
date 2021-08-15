package chameleon

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	natBirthday = nook.Birthday{
		Day:   25,
		Month: time.July}
)

var (
	natCode = nook.Code{
		Value: "chm"}
)

var (
	natAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Nat"}

	natCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Djarod"}

	natDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Nat"}

	natFrenchName = nook.Name{
		Language: language.French,
		Value:    "Djarod"}

	natGermanName = nook.Name{
		Language: language.German,
		Value:    "Carleon"}

	natItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gustavo"}

	natJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カメヤマさん"}

	natKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "멜레옹"}

	natLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Papilo"}

	natRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Нат"}

	natSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "龙山先生"}

	natSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Papilo"}

	natTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "龍山先生"}
)

var (
	natName = nook.Languages{
		language.AmericanEnglish:      natAmericanEnglishName,
		language.CanadianFrench:       natCanadianFrenchName,
		language.Dutch:                natDutchName,
		language.French:               natFrenchName,
		language.German:               natGermanName,
		language.Italian:              natItalianName,
		language.Japanese:             natJapaneseName,
		language.Korean:               natKoreanName,
		language.LatinAmericanSpanish: natLatinAmericanSpanishName,
		language.Russian:              natRussianName,
		language.SimplifiedChinese:    natSimplifiedChineseName,
		language.Spanish:              natSpanishName,
		language.TraditionalChinese:   natTraditionalChineseName}
)

var (
	natCharacter = nook.Character{
		Animal:   Chameleon,
		Birthday: natBirthday,
		Code:     natCode,
		Gender:   nook.Male,
		Name:     natName}
)

var (
	Nat = nook.Resident{
		Character: natCharacter}
)
