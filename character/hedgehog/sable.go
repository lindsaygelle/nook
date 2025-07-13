package hedgehog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// sableBirthday represents sable birthday.
	sableBirthday = nook.Birthday{
		Day:   22,
		Month: time.November}
)

var (
	// sableCode represents sable code.
	sableCode = nook.Code{
		Value: "hgs"}
)

var (
	// sableAmericanEnglishName represents sable american english name.
	sableAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sable"}

	// sableCanadianFrenchName represents sable canadian french name.
	sableCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cousette"}

	// sableDutchName represents sable dutch name.
	sableDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sable"}

	// sableFrenchName represents sable french name.
	sableFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cousette"}

	// sableGermanName represents sable german name.
	sableGermanName = nook.Name{
		Language: language.German,
		Value:    "Sina"}

	// sableItalianName represents sable italian name.
	sableItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Filomena"}

	// sableJapaneseName represents sable japanese name.
	sableJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "あさみ"}

	// sableKoreanName represents sable korean name.
	sableKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "고옥이"}

	// sableLatinAmericanSpanishName represents sable latin american spanish name.
	sableLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mili"}

	// sableRussianName represents sable russian name.
	sableRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Сэйбл"}

	// sableSimplifiedChineseName represents sable simplified chinese name.
	sableSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "麻儿"}

	// sableSpanishName represents sable spanish name.
	sableSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mili"}

	// sableTraditionalChineseName represents sable traditional chinese name.
	sableTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "麻兒"}
)

var (
	// sableName represents sable name.
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
	// sableCharacter represents sable character.
	sableCharacter = nook.Character{
		Animal:   animal.Hedgehog,
		Birthday: sableBirthday,
		Code:     sableCode,
		Key:      character.Sable,
		Gender:   gender.Female,
		Name:     sableName,
		Special:  true}
)

var (
	// Sable represents sable.
	Sable = nook.Resident{
		Character: sableCharacter}
)
