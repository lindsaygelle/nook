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
	// lottieBirthday represents lottie birthday.
	lottieBirthday = nook.Birthday{
		Day:   12,
		Month: time.September}
)

var (
	// lottieCode represents lottie code.
	lottieCode = nook.Code{
		Value: "otg"}
)

var (
	// lottieAmericanEnglishName represents lottie american english name.
	lottieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lottie"}

	// lottieCanadianFrenchName represents lottie canadian french name.
	lottieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lou"}

	// lottieDutchName represents lottie dutch name.
	lottieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lottie"}

	// lottieFrenchName represents lottie french name.
	lottieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lou"}

	// lottieGermanName represents lottie german name.
	lottieGermanName = nook.Name{
		Language: language.German,
		Value:    "Karlotta"}

	// lottieItalianName represents lottie italian name.
	lottieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Casimira"}

	// lottieJapaneseName represents lottie japanese name.
	lottieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タクミ"}

	// lottieKoreanName represents lottie korean name.
	lottieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "솜이"}

	// lottieLatinAmericanSpanishName represents lottie latin american spanish name.
	lottieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nuria"}

	// lottieRussianName represents lottie russian name.
	lottieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лотти"}

	// lottieSimplifiedChineseName represents lottie simplified chinese name.
	lottieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "巧美"}

	// lottieSpanishName represents lottie spanish name.
	lottieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Nuria"}

	// lottieTraditionalChineseName represents lottie traditional chinese name.
	lottieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "巧美"}
)

var (
	// lottieName represents lottie name.
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
	// lottieCharacter represents lottie character.
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
	// Lottie represents lottie.
	Lottie = nook.Resident{
		Character: lottieCharacter}
)
