package hedgehog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	mabelBirthday = nook.Birthday{
		Day:   22,
		Month: time.May}
)

var (
	mabelCode = nook.Code{
		Value: "hgh"}
)

var (
	mabelAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mabel"}

	mabelCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Layette"}

	mabelDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Mabel"}

	mabelFrenchName = nook.Name{
		Language: language.French,
		Value:    "Layette"}

	mabelGermanName = nook.Name{
		Language: language.German,
		Value:    "Tina"}

	mabelItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Agostina"}

	mabelJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "きぬよ"}

	mabelKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "고순이"}

	mabelLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pili"}

	mabelRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мэйбл"}

	mabelSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "绢儿"}

	mabelSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pili"}

	mabelTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "絹兒"}
)

var (
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
	mabelCharacter = nook.Character{
		Animal:   Hedgehog,
		Birthday: mabelBirthday,
		Code:     mabelCode,
		Gender:   nook.Female,
		Name:     mabelName}
)

var (
	Mabel = nook.Resident{
		Character: mabelCharacter}
)
