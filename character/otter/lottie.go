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
	lottieBirthday = nook.Birthday{
		Day:   12,
		Month: time.September}
)

var (
	lottieCode = nook.Code{
		Value: "otg"}
)

var (
	lottieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lottie"}

	lottieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lou"}

	lottieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lottie"}

	lottieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lou"}

	lottieGermanName = nook.Name{
		Language: language.German,
		Value:    "Karlotta"}

	lottieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Casimira"}

	lottieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タクミ"}

	lottieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "솜이"}

	lottieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nuria"}

	lottieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лотти"}

	lottieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "巧美"}

	lottieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Nuria"}

	lottieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "巧美"}
)

var (
	lottieName = nook.Languages{
		language.AmericanEnglish:      lottieAmericanEnglishName,
		language.CanadianFrench:       lottieCanadianFrenchName,
		language.Dutch:                lottieDutchName,
		language.French:               lottieFrenchName,
		language.German:               lottieGermanName,
		language.Italian:              lottieItalianName,
		language.Japanese:             lottieJapaneseName,
		language.Korean:               lottieKoreanName,
		language.LatinAmericanSpanish: lottieLatinAmericanSpanishName,
		language.Russian:              lottieRussianName,
		language.SimplifiedChinese:    lottieSimplifiedChineseName,
		language.Spanish:              lottieSpanishName,
		language.TraditionalChinese:   lottieTraditionalChineseName}
)

var (
	lottieCharacter = nook.Character{
		Animal:   animal.Otter,
		Birthday: lottieBirthday,
		Code:     lottieCode,
		Key:      character.Lottie,
		Gender:   gender.Female,
		Name:     lottieName,
		Special:  true}
)

var (
	Lottie = nook.Resident{
		Character: lottieCharacter}
)
