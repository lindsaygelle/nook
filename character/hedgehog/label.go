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
	// labelBirthday represents label birthday.
	labelBirthday = nook.Birthday{
		Day:   31,
		Month: time.October}
)

var (
	// labelCode represents label code.
	labelCode = nook.Code{
		Value: "hgc"}
)

var (
	// labelAmericanEnglishName represents label american english name.
	labelAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Label"}

	// labelCanadianFrenchName represents label canadian french name.
	labelCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tiquette"}

	// labelDutchName represents label dutch name.
	labelDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Label"}

	// labelFrenchName represents label french name.
	labelFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tiquette"}

	// labelGermanName represents label german name.
	labelGermanName = nook.Name{
		Language: language.German,
		Value:    "Minna"}

	// labelItalianName represents label italian name.
	labelItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Beatrice"}

	// labelJapaneseName represents label japanese name.
	labelJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ことの"}

	// labelKoreanName represents label korean name.
	labelKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "고숙이"}

	// labelLatinAmericanSpanishName represents label latin american spanish name.
	labelLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tili"}

	// labelRussianName represents label russian name.
	labelRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лэйбл"}

	// labelSimplifiedChineseName represents label simplified chinese name.
	labelSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "绵儿"}

	// labelSpanishName represents label spanish name.
	labelSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tili"}

	// labelTraditionalChineseName represents label traditional chinese name.
	labelTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "綿兒"}
)

var (
	// labelName represents label name.
	labelName = nook.Languages{
		language.AmericanEnglish:      labelAmericanEnglishName,
		language.CanadianFrench:       labelCanadianFrenchName,
		language.Dutch:                labelDutchName,
		language.French:               labelFrenchName,
		language.German:               labelGermanName,
		language.Italian:              labelItalianName,
		language.Japanese:             labelJapaneseName,
		language.Korean:               labelKoreanName,
		language.LatinAmericanSpanish: labelLatinAmericanSpanishName,
		language.Russian:              labelRussianName,
		language.SimplifiedChinese:    labelSimplifiedChineseName,
		language.Spanish:              labelSpanishName,
		language.TraditionalChinese:   labelTraditionalChineseName}
)

var (
	// labelCharacter represents label character.
	labelCharacter = nook.Character{
		Animal:   animal.Hedgehog,
		Birthday: labelBirthday,
		Code:     labelCode,
		Key:      character.Label,
		Gender:   gender.Female,
		Name:     labelName,
		Special:  true}
)

var (
	// Label represents label.
	Label = nook.Resident{
		Character: labelCharacter}
)
