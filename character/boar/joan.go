package boar

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// joanBirthday represents joan birthday.
	joanBirthday = nook.Birthday{
		Day:   8,
		Month: time.January}
)

var (
	// joanCode represents joan code.
	joanCode = nook.Code{
		Value: "boa"}
)

var (
	// joanAmericanEnglishName represents joan american english name.
	joanAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Joan"}

	// joanCanadianFrenchName represents joan canadian french name.
	joanCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Porcella"}

	// joanDutchName represents joan dutch name.
	joanDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Joan"}

	// joanFrenchName represents joan french name.
	joanFrenchName = nook.Name{
		Language: language.French,
		Value:    "Porcella"}

	// joanGermanName represents joan german name.
	joanGermanName = nook.Name{
		Language: language.German,
		Value:    "Sigrid"}

	// joanItalianName represents joan italian name.
	joanItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Nella"}

	// joanJapaneseName represents joan japanese name.
	joanJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カブリバ"}

	// joanKoreanName represents joan korean name.
	joanKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "무파라"}

	// joanLatinAmericanSpanishName represents joan latin american spanish name.
	joanLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Juana"}

	// joanRussianName represents joan russian name.
	joanRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джоан"}

	// joanSimplifiedChineseName represents joan simplified chinese name.
	joanSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "曹谷"}

	// joanSpanishName represents joan spanish name.
	joanSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Juana"}

	// joanTraditionalChineseName represents joan traditional chinese name.
	joanTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "曹谷"}
)

var (
	// joanName represents joan name.
	joanName = nook.Languages{
		language.AmericanEnglish:      joanAmericanEnglishName,
		language.CanadianFrench:       joanCanadianFrenchName,
		language.Dutch:                joanDutchName,
		language.French:               joanFrenchName,
		language.German:               joanGermanName,
		language.Italian:              joanItalianName,
		language.Japanese:             joanJapaneseName,
		language.Korean:               joanKoreanName,
		language.LatinAmericanSpanish: joanLatinAmericanSpanishName,
		language.Russian:              joanRussianName,
		language.SimplifiedChinese:    joanSimplifiedChineseName,
		language.Spanish:              joanSpanishName,
		language.TraditionalChinese:   joanTraditionalChineseName}
)

var (
	// joanCharacter represents joan character.
	joanCharacter = nook.Character{
		Animal:   animal.Boar,
		Birthday: joanBirthday,
		Code:     joanCode,
		Key:      character.Joan,
		Gender:   gender.Female,
		Name:     joanName,
		Special:  true}
)

var (
	// Joan represents joan.
	Joan = nook.Resident{
		Character: joanCharacter}
)
