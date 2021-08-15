package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	blancaBirthday = nook.Birthday{
		Day:   8,
		Month: time.February}
)

var (
	blancaCode = nook.Code{
		Value: "mka"}
)

var (
	blancaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Blanca"}

	blancaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Blanca"}

	blancaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Blanca"}

	blancaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Blanca"}

	blancaGermanName = nook.Name{
		Language: language.German,
		Value:    "Blanka"}

	blancaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Blanca"}

	blancaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "あやしいネコ"}

	blancaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아트고양이"}

	blancaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Blanca"}

	blancaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бланка"}

	blancaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "无脸猫"}

	blancaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Blanca"}

	blancaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "無臉貓"}
)

var (
	blancaName = nook.Languages{
		language.AmericanEnglish:      blancaAmericanEnglishName,
		language.CanadianFrench:       blancaCanadianFrenchName,
		language.Dutch:                blancaDutchName,
		language.French:               blancaFrenchName,
		language.German:               blancaGermanName,
		language.Italian:              blancaItalianName,
		language.Japanese:             blancaJapaneseName,
		language.Korean:               blancaKoreanName,
		language.LatinAmericanSpanish: blancaLatinAmericanSpanishName,
		language.Russian:              blancaRussianName,
		language.SimplifiedChinese:    blancaSimplifiedChineseName,
		language.Spanish:              blancaSpanishName,
		language.TraditionalChinese:   blancaTraditionalChineseName}
)

var (
	blancaCharacter = nook.Character{
		Animal:   Cat,
		Birthday: blancaBirthday,
		Code:     blancaCode,
		Gender:   nook.Female,
		Name:     blancaName}
)

var (
	Blanca = nook.Resident{
		Character: blancaCharacter}
)
