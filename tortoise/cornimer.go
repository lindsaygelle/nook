package tortoise

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	cornimerBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	cornimerCode = nook.Code{
		Value: "dnk"}
)

var (
	cornimerAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cornimer"}

	cornimerCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cornimer"}

	cornimerDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	cornimerFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cornimer"}

	cornimerGermanName = nook.Name{
		Language: language.German,
		Value:    "Eichenm."}

	cornimerItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ghiandy"}

	cornimerJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ドンどんぐり"}

	cornimerKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "도토리옹"}

	cornimerLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Córnimer"}

	cornimerRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	cornimerSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	cornimerSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Córnimer"}

	cornimerTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	cornimerName = nook.Languages{
		language.AmericanEnglish:      cornimerAmericanEnglishName,
		language.CanadianFrench:       cornimerCanadianFrenchName,
		language.Dutch:                cornimerDutchName,
		language.French:               cornimerFrenchName,
		language.German:               cornimerGermanName,
		language.Italian:              cornimerItalianName,
		language.Japanese:             cornimerJapaneseName,
		language.Korean:               cornimerKoreanName,
		language.LatinAmericanSpanish: cornimerLatinAmericanSpanishName,
		language.Russian:              cornimerRussianName,
		language.SimplifiedChinese:    cornimerSimplifiedChineseName,
		language.Spanish:              cornimerSpanishName,
		language.TraditionalChinese:   cornimerTraditionalChineseName}
)

var (
	cornimerCharacter = nook.Character{
		Animal:   Tortoise,
		Birthday: cornimerBirthday,
		Code:     cornimerCode,
		Gender:   nook.Male,
		Name:     cornimerName}
)

var (
	Cornimer = nook.Resident{
		Character: cornimerCharacter}
)
