package turtle

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// kappnBirthday represents kappn birthday.
	kappnBirthday = nook.Birthday{
		Day:   12,
		Month: time.July}
)

var (
	// kappnCode represents kappn code.
	kappnCode = nook.Code{
		Value: "wip/kpp"}
)

var (
	// kappnAmericanEnglishName represents kappn american english name.
	kappnAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kapp'n"}

	// kappnCanadianFrenchName represents kappn canadian french name.
	kappnCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Amiral"}

	// kappnDutchName represents kappn dutch name.
	kappnDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kapp'n"}

	// kappnFrenchName represents kappn french name.
	kappnFrenchName = nook.Name{
		Language: language.French,
		Value:    "Amiral"}

	// kappnGermanName represents kappn german name.
	kappnGermanName = nook.Name{
		Language: language.German,
		Value:    "Käpten"}

	// kappnItalianName represents kappn italian name.
	kappnItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Remo"}

	// kappnJapaneseName represents kappn japanese name.
	kappnJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "かっぺい"}

	// kappnKoreanName represents kappn korean name.
	kappnKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "갑돌"}

	// kappnLatinAmericanSpanishName represents kappn latin american spanish name.
	kappnLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Capitán"}

	// kappnRussianName represents kappn russian name.
	kappnRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кэпн"}

	// kappnSimplifiedChineseName represents kappn simplified chinese name.
	kappnSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "航平"}

	// kappnSpanishName represents kappn spanish name.
	kappnSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Capitán"}

	// kappnTraditionalChineseName represents kappn traditional chinese name.
	kappnTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "航平"}
)

var (
	// kappnName represents kappn name.
	kappnName = nook.Languages{
		language.AmericanEnglish:      kappnAmericanEnglishName,
		language.CanadianFrench:       kappnCanadianFrenchName,
		language.Dutch:                kappnDutchName,
		language.French:               kappnFrenchName,
		language.German:               kappnGermanName,
		language.Italian:              kappnItalianName,
		language.Japanese:             kappnJapaneseName,
		language.Korean:               kappnKoreanName,
		language.LatinAmericanSpanish: kappnLatinAmericanSpanishName,
		language.Russian:              kappnRussianName,
		language.SimplifiedChinese:    kappnSimplifiedChineseName,
		language.Spanish:              kappnSpanishName,
		language.TraditionalChinese:   kappnTraditionalChineseName}
)

var (
	// kappnCharacter represents kappn character.
	kappnCharacter = nook.Character{
		Animal:   animal.Turtle,
		Birthday: kappnBirthday,
		Code:     kappnCode,
		Key:      character.Kappn,
		Gender:   gender.Male,
		Name:     kappnName,
		Special:  true}
)

var (
	// Kappn represents kappn.
	Kappn = nook.Resident{
		Character: kappnCharacter}
)
