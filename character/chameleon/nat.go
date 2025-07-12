package chameleon

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// natBirthday represents nat birthday.
	natBirthday = nook.Birthday{
		Day:   25,
		Month: time.July}
)

var (
	// natCode represents nat code.
	natCode = nook.Code{
		Value: "chm"}
)

var (
	// natAmericanEnglishName represents nat american english name.
	natAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Nat"}

	// natCanadianFrenchName represents nat canadian french name.
	natCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Djarod"}

	// natDutchName represents nat dutch name.
	natDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Nat"}

	// natFrenchName represents nat french name.
	natFrenchName = nook.Name{
		Language: language.French,
		Value:    "Djarod"}

	// natGermanName represents nat german name.
	natGermanName = nook.Name{
		Language: language.German,
		Value:    "Carleon"}

	// natItalianName represents nat italian name.
	natItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gustavo"}

	// natJapaneseName represents nat japanese name.
	natJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カメヤマさん"}

	// natKoreanName represents nat korean name.
	natKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "멜레옹"}

	// natLatinAmericanSpanishName represents nat latin american spanish name.
	natLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Papilo"}

	// natRussianName represents nat russian name.
	natRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Нат"}

	// natSimplifiedChineseName represents nat simplified chinese name.
	natSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "龙山先生"}

	// natSpanishName represents nat spanish name.
	natSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Papilo"}

	// natTraditionalChineseName represents nat traditional chinese name.
	natTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "龍山先生"}
)

var (
	// natName represents nat name.
	natName = nook.Languages{
		language.AmericanEnglish:      natAmericanEnglishName,
		language.CanadianFrench:       natCanadianFrenchName,
		language.Dutch:                natDutchName,
		language.French:               natFrenchName,
		language.German:               natGermanName,
		language.Italian:              natItalianName,
		language.Japanese:             natJapaneseName,
		language.Korean:               natKoreanName,
		language.LatinAmericanSpanish: natLatinAmericanSpanishName,
		language.Russian:              natRussianName,
		language.SimplifiedChinese:    natSimplifiedChineseName,
		language.Spanish:              natSpanishName,
		language.TraditionalChinese:   natTraditionalChineseName}
)

var (
	// natCharacter represents nat character.
	natCharacter = nook.Character{
		Animal:   animal.Chameleon,
		Birthday: natBirthday,
		Code:     natCode,
		Key:      character.Nat,
		Gender:   gender.Male,
		Name:     natName,
		Special:  true}
)

var (
	// Nat represents nat.
	Nat = nook.Resident{
		Character: natCharacter}
)
