package walrus

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// wendellBirthday represents wendell birthday.
	wendellBirthday = nook.Birthday{
		Day:   25,
		Month: time.February}
)

var (
	// wendellCode represents wendell code.
	wendellCode = nook.Code{
		Value: "wls/wrl"}
)

var (
	// wendellAmericanEnglishName represents wendell american english name.
	wendellAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Wendell"}

	// wendellCanadianFrenchName represents wendell canadian french name.
	wendellCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Morsicus"}

	// wendellDutchName represents wendell dutch name.
	wendellDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Wendell"}

	// wendellFrenchName represents wendell french name.
	wendellFrenchName = nook.Name{
		Language: language.French,
		Value:    "Morsicus"}

	// wendellGermanName represents wendell german name.
	wendellGermanName = nook.Name{
		Language: language.German,
		Value:    "Winci"}

	// wendellItalianName represents wendell italian name.
	wendellItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Van Trik"}

	// wendellJapaneseName represents wendell japanese name.
	wendellJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "セイイチ"}

	// wendellKoreanName represents wendell korean name.
	wendellKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "고파유"}

	// wendellLatinAmericanSpanishName represents wendell latin american spanish name.
	wendellLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Da Morsi"}

	// wendellRussianName represents wendell russian name.
	wendellRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Уэнделл"}

	// wendellSimplifiedChineseName represents wendell simplified chinese name.
	wendellSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "海项"}

	// wendellSpanishName represents wendell spanish name.
	wendellSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Da Morsi"}

	// wendellTraditionalChineseName represents wendell traditional chinese name.
	wendellTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "海項"}
)

var (
	// wendellName represents wendell name.
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
	// wendellCharacter represents wendell character.
	wendellCharacter = nook.Character{
		Animal:   animal.Walrus,
		Birthday: wendellBirthday,
		Code:     wendellCode,
		Key:      character.Wendell,
		Gender:   gender.Male,
		Name:     wendellName,
		Special:  true}
)

var (
	// Wendell represents wendell.
	Wendell = nook.Resident{
		Character: wendellCharacter}
)
