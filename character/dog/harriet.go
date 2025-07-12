package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// harrietBirthday represents harriet birthday.
	harrietBirthday = nook.Birthday{
		Day:   31,
		Month: time.January}
)

var (
	// harrietCode represents harriet code.
	harrietCode = nook.Code{
		Value: "poo"}
)

var (
	// harrietAmericanEnglishName represents harriet american english name.
	harrietAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Harriet"}

	// harrietCanadianFrenchName represents harriet canadian french name.
	harrietCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ginette"}

	// harrietDutchName represents harriet dutch name.
	harrietDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Harriet"}

	// harrietFrenchName represents harriet french name.
	harrietFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ginette"}

	// harrietGermanName represents harriet german name.
	harrietGermanName = nook.Name{
		Language: language.German,
		Value:    "Trude"}

	// harrietItalianName represents harriet italian name.
	harrietItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bigodina"}

	// harrietJapaneseName represents harriet japanese name.
	harrietJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カットリーヌ"}

	// harrietKoreanName represents harriet korean name.
	harrietKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "카트리나"}

	// harrietLatinAmericanSpanishName represents harriet latin american spanish name.
	harrietLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Marilín"}

	// harrietRussianName represents harriet russian name.
	harrietRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Харриет"}

	// harrietSimplifiedChineseName represents harriet simplified chinese name.
	harrietSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "刀丽茹"}

	// harrietSpanishName represents harriet spanish name.
	harrietSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Marilín"}

	// harrietTraditionalChineseName represents harriet traditional chinese name.
	harrietTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "刀麗茹"}
)

var (
	// harrietName represents harriet name.
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
	// harrietCharacter represents harriet character.
	harrietCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: harrietBirthday,
		Code:     harrietCode,
		Key:      character.Harriet,
		Gender:   gender.Female,
		Name:     harrietName,
		Special:  true}
)

var (
	// Harriet represents harriet.
	Harriet = nook.Resident{
		Character: harrietCharacter}
)
