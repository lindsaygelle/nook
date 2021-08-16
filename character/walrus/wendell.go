package walrus

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	wendellBirthday = nook.Birthday{
		Day:   25,
		Month: time.February}
)

var (
	wendellCode = nook.Code{
		Value: "wls/wrl"}
)

var (
	wendellAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Wendell"}

	wendellCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Morsicus"}

	wendellDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Wendell"}

	wendellFrenchName = nook.Name{
		Language: language.French,
		Value:    "Morsicus"}

	wendellGermanName = nook.Name{
		Language: language.German,
		Value:    "Winci"}

	wendellItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Van Trik"}

	wendellJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "セイイチ"}

	wendellKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "고파유"}

	wendellLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Da Morsi"}

	wendellRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Уэнделл"}

	wendellSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "海项"}

	wendellSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Da Morsi"}

	wendellTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "海項"}
)

var (
	wendellName = nook.Languages{
		language.AmericanEnglish:      wendellAmericanEnglishName,
		language.CanadianFrench:       wendellCanadianFrenchName,
		language.Dutch:                wendellDutchName,
		language.French:               wendellFrenchName,
		language.German:               wendellGermanName,
		language.Italian:              wendellItalianName,
		language.Japanese:             wendellJapaneseName,
		language.Korean:               wendellKoreanName,
		language.LatinAmericanSpanish: wendellLatinAmericanSpanishName,
		language.Russian:              wendellRussianName,
		language.SimplifiedChinese:    wendellSimplifiedChineseName,
		language.Spanish:              wendellSpanishName,
		language.TraditionalChinese:   wendellTraditionalChineseName}
)

var (
	wendellCharacter = nook.Character{
		Animal:   animal.Walrus,
		Birthday: wendellBirthday,
		Code:     wendellCode,
		Gender:   gender.Male,
		Name:     wendellName}
)

var (
	Wendell = nook.Resident{
		Character: wendellCharacter}
)
