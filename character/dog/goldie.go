package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	// goldieBirthday represents goldie birthday.
	goldieBirthday = nook.Birthday{
		Day:   27,
		Month: time.December}
)

var (
	// goldieCode represents goldie code.
	goldieCode = nook.Code{
		Value: "dog00"}
)

var (
	// goldieAmericanEnglishName represents goldie american english name.
	goldieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Goldie"}

	// goldieCanadianFrenchName represents goldie canadian french name.
	goldieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mirza"}

	// goldieDutchName represents goldie dutch name.
	goldieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Goldie"}

	// goldieFrenchName represents goldie french name.
	goldieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mirza"}

	// goldieGermanName represents goldie german name.
	goldieGermanName = nook.Name{
		Language: language.German,
		Value:    "Bienchen"}

	// goldieItalianName represents goldie italian name.
	goldieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Dora"}

	// goldieJapaneseName represents goldie japanese name.
	goldieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キャラメル"}

	// goldieKoreanName represents goldie korean name.
	goldieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "카라멜"}

	// goldieLatinAmericanSpanishName represents goldie latin american spanish name.
	goldieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tere"}

	// goldieRussianName represents goldie russian name.
	goldieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Голди"}

	// goldieSimplifiedChineseName represents goldie simplified chinese name.
	goldieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "牛奶糖"}

	// goldieSpanishName represents goldie spanish name.
	goldieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tere"}

	// goldieTraditionalChineseName represents goldie traditional chinese name.
	goldieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "牛奶糖"}
)

var (
	// goldieName represents goldie name.
	goldieName = nook.Languages{
		language.AmericanEnglish:      goldieAmericanEnglishName,
		language.CanadianFrench:       goldieCanadianFrenchName,
		language.Dutch:                goldieDutchName,
		language.French:               goldieFrenchName,
		language.German:               goldieGermanName,
		language.Italian:              goldieItalianName,
		language.Japanese:             goldieJapaneseName,
		language.Korean:               goldieKoreanName,
		language.LatinAmericanSpanish: goldieLatinAmericanSpanishName,
		language.Russian:              goldieRussianName,
		language.SimplifiedChinese:    goldieSimplifiedChineseName,
		language.Spanish:              goldieSpanishName,
		language.TraditionalChinese:   goldieTraditionalChineseName}
)

var (
	// goldieCharacter represents goldie character.
	goldieCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: goldieBirthday,
		Code:     goldieCode,
		Key:      character.Goldie,
		Gender:   gender.Female,
		Name:     goldieName,
		Special:  false}
)

var (
	// goldieAmericanEnglishPhrase represents goldie american english phrase.
	goldieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "woof"}

	// goldieCanadianFrenchPhrase represents goldie canadian french phrase.
	goldieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ouaf"}

	// goldieDutchPhrase represents goldie dutch phrase.
	goldieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "woef"}

	// goldieFrenchPhrase represents goldie french phrase.
	goldieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ouaf"}

	// goldieGermanPhrase represents goldie german phrase.
	goldieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "wuff"}

	// goldieItalianPhrase represents goldie italian phrase.
	goldieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "wuf wuf"}

	// goldieJapanesePhrase represents goldie japanese phrase.
	goldieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ワン"}

	// goldieKoreanPhrase represents goldie korean phrase.
	goldieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "왈왈"}

	// goldieLatinAmericanSpanishPhrase represents goldie latin american spanish phrase.
	goldieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "guau-guau"}

	// goldieRussianPhrase represents goldie russian phrase.
	goldieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тяв"}

	// goldieSimplifiedChinesePhrase represents goldie simplified chinese phrase.
	goldieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "汪"}

	// goldieSpanishPhrase represents goldie spanish phrase.
	goldieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "guau-guau"}

	// goldieTraditionalChinesePhrase represents goldie traditional chinese phrase.
	goldieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "汪"}
)

var (
	// goldiePhrase represents goldie phrase.
	goldiePhrase = nook.Languages{
		language.AmericanEnglish:      goldieAmericanEnglishPhrase,
		language.CanadianFrench:       goldieCanadianFrenchPhrase,
		language.Dutch:                goldieDutchPhrase,
		language.French:               goldieFrenchPhrase,
		language.German:               goldieGermanPhrase,
		language.Italian:              goldieItalianPhrase,
		language.Japanese:             goldieJapanesePhrase,
		language.Korean:               goldieKoreanPhrase,
		language.LatinAmericanSpanish: goldieLatinAmericanSpanishPhrase,
		language.Russian:              goldieRussianPhrase,
		language.SimplifiedChinese:    goldieSimplifiedChinesePhrase,
		language.Spanish:              goldieSpanishPhrase,
		language.TraditionalChinese:   goldieTraditionalChinesePhrase}
)

var (
	// Goldie represents goldie.
	Goldie = nook.Villager{
		Character:   goldieCharacter,
		Personality: personality.Normal,
		Phrase:      goldiePhrase}
)
