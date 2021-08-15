package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	harrietBirthday = nook.Birthday{
		Day:   31,
		Month: time.January}
)

var (
	harrietCode = nook.Code{
		Value: "poo"}
)

var (
	harrietAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Harriet"}

	harrietCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ginette"}

	harrietDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Harriet"}

	harrietFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ginette"}

	harrietGermanName = nook.Name{
		Language: language.German,
		Value:    "Trude"}

	harrietItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bigodina"}

	harrietJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カットリーヌ"}

	harrietKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "카트리나"}

	harrietLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Marilín"}

	harrietRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Харриет"}

	harrietSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "刀丽茹"}

	harrietSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Marilín"}

	harrietTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "刀麗茹"}
)

var (
	harrietName = nook.Languages{
		language.AmericanEnglish:      harrietAmericanEnglishName,
		language.CanadianFrench:       harrietCanadianFrenchName,
		language.Dutch:                harrietDutchName,
		language.French:               harrietFrenchName,
		language.German:               harrietGermanName,
		language.Italian:              harrietItalianName,
		language.Japanese:             harrietJapaneseName,
		language.Korean:               harrietKoreanName,
		language.LatinAmericanSpanish: harrietLatinAmericanSpanishName,
		language.Russian:              harrietRussianName,
		language.SimplifiedChinese:    harrietSimplifiedChineseName,
		language.Spanish:              harrietSpanishName,
		language.TraditionalChinese:   harrietTraditionalChineseName}
)

var (
	harrietCharacter = nook.Character{
		Animal:   Dog,
		Birthday: harrietBirthday,
		Code:     harrietCode,
		Gender:   nook.Female,
		Name:     harrietName}
)

var (
	Harriet = nook.Resident{
		Character: harrietCharacter}
)
