package tapir

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	lunaBirthday = nook.Birthday{
		Day:   29,
		Month: time.February}
)

var (
	lunaCode = nook.Code{
		Value: "tap"}
)

var (
	lunaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Luna"}

	lunaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Serena"}

	lunaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Luna"}

	lunaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Serena"}

	lunaGermanName = nook.Name{
		Language: language.German,
		Value:    "Serenada"}

	lunaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sonia"}

	lunaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゆめみ"}

	lunaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "몽셰르"}

	lunaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Alakama"}

	lunaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Луна"}

	lunaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "梦美"}

	lunaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Alakama"}

	lunaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "夢美"}
)

var (
	lunaName = nook.Languages{
		language.AmericanEnglish:      lunaAmericanEnglishName,
		language.CanadianFrench:       lunaCanadianFrenchName,
		language.Dutch:                lunaDutchName,
		language.French:               lunaFrenchName,
		language.German:               lunaGermanName,
		language.Italian:              lunaItalianName,
		language.Japanese:             lunaJapaneseName,
		language.Korean:               lunaKoreanName,
		language.LatinAmericanSpanish: lunaLatinAmericanSpanishName,
		language.Russian:              lunaRussianName,
		language.SimplifiedChinese:    lunaSimplifiedChineseName,
		language.Spanish:              lunaSpanishName,
		language.TraditionalChinese:   lunaTraditionalChineseName}
)

var (
	lunaCharacter = nook.Character{
		Animal:   animal.Tapir,
		Birthday: lunaBirthday,
		Code:     lunaCode,
		Key:      character.Luna,
		Gender:   gender.Female,
		Name:     lunaName,
		Special:  true}
)

var (
	Luna = nook.Resident{
		Character: lunaCharacter}
)
