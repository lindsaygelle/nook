package tapir

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// lunaBirthday represents luna birthday.
	lunaBirthday = nook.Birthday{
		Day:   29,
		Month: time.February}
)

var (
	// lunaCode represents luna code.
	lunaCode = nook.Code{
		Value: "tap"}
)

var (
	// lunaAmericanEnglishName represents luna american english name.
	lunaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Luna"}

	// lunaCanadianFrenchName represents luna canadian french name.
	lunaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Serena"}

	// lunaDutchName represents luna dutch name.
	lunaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Luna"}

	// lunaFrenchName represents luna french name.
	lunaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Serena"}

	// lunaGermanName represents luna german name.
	lunaGermanName = nook.Name{
		Language: language.German,
		Value:    "Serenada"}

	// lunaItalianName represents luna italian name.
	lunaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sonia"}

	// lunaJapaneseName represents luna japanese name.
	lunaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゆめみ"}

	// lunaKoreanName represents luna korean name.
	lunaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "몽셰르"}

	// lunaLatinAmericanSpanishName represents luna latin american spanish name.
	lunaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Alakama"}

	// lunaRussianName represents luna russian name.
	lunaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Луна"}

	// lunaSimplifiedChineseName represents luna simplified chinese name.
	lunaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "梦美"}

	// lunaSpanishName represents luna spanish name.
	lunaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Alakama"}

	// lunaTraditionalChineseName represents luna traditional chinese name.
	lunaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "夢美"}
)

var (
	// lunaName represents luna name.
	lunaName = nook.Languages{
		language.AmericanEnglish:      lunaAmericanEnglishName,
		language.CanadianFrench:       lunaCanadianFrenchName,
		language.Dutch:                lunaDutchName,
		language.French:               lunaFrenchName,
		language.German:               lunaGermanName,
		language.Italian:              lunaItalianName,
		language.Japanese:             lunaJapaneseName,
		language.Korean:               lunaKoreanName,
		language.LatinAmericanSpanish: lunaLatinAmericanSpanishName,
		language.Russian:              lunaRussianName,
		language.SimplifiedChinese:    lunaSimplifiedChineseName,
		language.Spanish:              lunaSpanishName,
		language.TraditionalChinese:   lunaTraditionalChineseName}
)

var (
	// lunaCharacter represents luna character.
	lunaCharacter = nook.Character{
		Animal:   animal.Tapir,
		Birthday: lunaBirthday,
		Code:     lunaCode,
		Key:      character.Luna,
		Gender:   gender.Female,
		Name:     lunaName,
		Special:  true}
)

var (
	// Luna represents luna.
	Luna = nook.Resident{
		Character: lunaCharacter}
)
