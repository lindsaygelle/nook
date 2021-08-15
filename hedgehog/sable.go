package hedgehog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	sableBirthday = nook.Birthday{
		Day:   22,
		Month: time.November}
)

var (
	sableCode = nook.Code{
		Value: "hgs"}
)

var (
	sableAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sable"}

	sableCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cousette"}

	sableDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sable"}

	sableFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cousette"}

	sableGermanName = nook.Name{
		Language: language.German,
		Value:    "Sina"}

	sableItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Filomena"}

	sableJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "あさみ"}

	sableKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "고옥이"}

	sableLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mili"}

	sableRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Сэйбл"}

	sableSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "麻儿"}

	sableSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mili"}

	sableTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "麻兒"}
)

var (
	sableName = nook.Languages{
		language.AmericanEnglish:      sableAmericanEnglishName,
		language.CanadianFrench:       sableCanadianFrenchName,
		language.Dutch:                sableDutchName,
		language.French:               sableFrenchName,
		language.German:               sableGermanName,
		language.Italian:              sableItalianName,
		language.Japanese:             sableJapaneseName,
		language.Korean:               sableKoreanName,
		language.LatinAmericanSpanish: sableLatinAmericanSpanishName,
		language.Russian:              sableRussianName,
		language.SimplifiedChinese:    sableSimplifiedChineseName,
		language.Spanish:              sableSpanishName,
		language.TraditionalChinese:   sableTraditionalChineseName}
)

var (
	sableCharacter = nook.Character{
		Animal:   Hedgehog,
		Birthday: sableBirthday,
		Code:     sableCode,
		Gender:   nook.Female,
		Name:     sableName}
)

var (
	Sable = nook.Resident{
		Character: sableCharacter}
)
