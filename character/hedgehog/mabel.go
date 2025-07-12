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
	// mabelBirthday represents mabel birthday.
	mabelBirthday = nook.Birthday{
		Day:   22,
		Month: time.May}
)

var (
	// mabelCode represents mabel code.
	mabelCode = nook.Code{
		Value: "hgh"}
)

var (
	// mabelAmericanEnglishName represents mabel american english name.
	mabelAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mabel"}

	// mabelCanadianFrenchName represents mabel canadian french name.
	mabelCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Layette"}

	// mabelDutchName represents mabel dutch name.
	mabelDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Mabel"}

	// mabelFrenchName represents mabel french name.
	mabelFrenchName = nook.Name{
		Language: language.French,
		Value:    "Layette"}

	// mabelGermanName represents mabel german name.
	mabelGermanName = nook.Name{
		Language: language.German,
		Value:    "Tina"}

	// mabelItalianName represents mabel italian name.
	mabelItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Agostina"}

	// mabelJapaneseName represents mabel japanese name.
	mabelJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "きぬよ"}

	// mabelKoreanName represents mabel korean name.
	mabelKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "고순이"}

	// mabelLatinAmericanSpanishName represents mabel latin american spanish name.
	mabelLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pili"}

	// mabelRussianName represents mabel russian name.
	mabelRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мэйбл"}

	// mabelSimplifiedChineseName represents mabel simplified chinese name.
	mabelSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "绢儿"}

	// mabelSpanishName represents mabel spanish name.
	mabelSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pili"}

	// mabelTraditionalChineseName represents mabel traditional chinese name.
	mabelTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "絹兒"}
)

var (
	// mabelName represents mabel name.
	mabelName = nook.Languages{
		language.AmericanEnglish:      mabelAmericanEnglishName,
		language.CanadianFrench:       mabelCanadianFrenchName,
		language.Dutch:                mabelDutchName,
		language.French:               mabelFrenchName,
		language.German:               mabelGermanName,
		language.Italian:              mabelItalianName,
		language.Japanese:             mabelJapaneseName,
		language.Korean:               mabelKoreanName,
		language.LatinAmericanSpanish: mabelLatinAmericanSpanishName,
		language.Russian:              mabelRussianName,
		language.SimplifiedChinese:    mabelSimplifiedChineseName,
		language.Spanish:              mabelSpanishName,
		language.TraditionalChinese:   mabelTraditionalChineseName}
)

var (
	// mabelCharacter represents mabel character.
	mabelCharacter = nook.Character{
		Animal:   animal.Hedgehog,
		Birthday: mabelBirthday,
		Code:     mabelCode,
		Key:      character.Mabel,
		Gender:   gender.Female,
		Name:     mabelName,
		Special:  true}
)

var (
	// Mabel represents mabel.
	Mabel = nook.Resident{
		Character: mabelCharacter}
)
