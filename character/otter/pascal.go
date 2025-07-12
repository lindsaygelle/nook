package otter

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// pascalBirthday represents pascal birthday.
	pascalBirthday = nook.Birthday{
		Day:   19,
		Month: time.July}
)

var (
	// pascalCode represents pascal code.
	pascalCode = nook.Code{
		Value: "seo"}
)

var (
	// pascalAmericanEnglishName represents pascal american english name.
	pascalAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pascal"}

	// pascalCanadianFrenchName represents pascal canadian french name.
	pascalCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pascal"}

	// pascalDutchName represents pascal dutch name.
	pascalDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pascal"}

	// pascalFrenchName represents pascal french name.
	pascalFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pascal"}

	// pascalGermanName represents pascal german name.
	pascalGermanName = nook.Name{
		Language: language.German,
		Value:    "Johannes"}

	// pascalItalianName represents pascal italian name.
	pascalItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pasqualo"}

	// pascalJapaneseName represents pascal japanese name.
	pascalJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ラコスケ"}

	// pascalKoreanName represents pascal korean name.
	pascalKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "해탈한"}

	// pascalLatinAmericanSpanishName represents pascal latin american spanish name.
	pascalLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pascal"}

	// pascalRussianName represents pascal russian name.
	pascalRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Паскаль"}

	// pascalSimplifiedChineseName represents pascal simplified chinese name.
	pascalSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿獭"}

	// pascalSpanishName represents pascal spanish name.
	pascalSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pascal"}

	// pascalTraditionalChineseName represents pascal traditional chinese name.
	pascalTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿獺"}
)

var (
	// pascalName represents pascal name.
	pascalName = nook.Languages{
		language.AmericanEnglish:      pascalAmericanEnglishName,
		language.CanadianFrench:       pascalCanadianFrenchName,
		language.Dutch:                pascalDutchName,
		language.French:               pascalFrenchName,
		language.German:               pascalGermanName,
		language.Italian:              pascalItalianName,
		language.Japanese:             pascalJapaneseName,
		language.Korean:               pascalKoreanName,
		language.LatinAmericanSpanish: pascalLatinAmericanSpanishName,
		language.Russian:              pascalRussianName,
		language.SimplifiedChinese:    pascalSimplifiedChineseName,
		language.Spanish:              pascalSpanishName,
		language.TraditionalChinese:   pascalTraditionalChineseName}
)

var (
	// pascalCharacter represents pascal character.
	pascalCharacter = nook.Character{
		Animal:   animal.Otter,
		Birthday: pascalBirthday,
		Code:     pascalCode,
		Key:      character.Pascal,
		Gender:   gender.Male,
		Name:     pascalName,
		Special:  true}
)

var (
	// Pascal represents pascal.
	Pascal = nook.Resident{
		Character: pascalCharacter}
)
