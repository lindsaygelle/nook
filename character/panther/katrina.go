package panther

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	katrinaBirthday = nook.Birthday{
		Day:   28,
		Month: time.October}
)

var (
	katrinaCode = nook.Code{
		Value: "bpt"}
)

var (
	katrinaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Katrina"}

	katrinaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Astrid"}

	katrinaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Katrina"}

	katrinaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Astrid"}

	katrinaGermanName = nook.Name{
		Language: language.German,
		Value:    "Smeralda"}

	katrinaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Vanda"}

	katrinaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハッケミィ"}

	katrinaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마추릴라"}

	katrinaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Katrina"}

	katrinaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Катрина"}

	katrinaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "星薇"}

	katrinaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Katrina"}

	katrinaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "星薇"}
)

var (
	katrinaName = nook.Languages{
		language.AmericanEnglish:      katrinaAmericanEnglishName,
		language.CanadianFrench:       katrinaCanadianFrenchName,
		language.Dutch:                katrinaDutchName,
		language.French:               katrinaFrenchName,
		language.German:               katrinaGermanName,
		language.Italian:              katrinaItalianName,
		language.Japanese:             katrinaJapaneseName,
		language.Korean:               katrinaKoreanName,
		language.LatinAmericanSpanish: katrinaLatinAmericanSpanishName,
		language.Russian:              katrinaRussianName,
		language.SimplifiedChinese:    katrinaSimplifiedChineseName,
		language.Spanish:              katrinaSpanishName,
		language.TraditionalChinese:   katrinaTraditionalChineseName}
)

var (
	katrinaCharacter = nook.Character{
		Animal:   animal.Panther,
		Birthday: katrinaBirthday,
		Code:     katrinaCode,
		Gender:   gender.Female,
		Name:     katrinaName}
)

var (
	Katrina = nook.Resident{
		Character: katrinaCharacter}
)
