package cat

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
	// tomBirthday represents tom birthday.
	tomBirthday = nook.Birthday{
		Day:   10,
		Month: time.December}
)

var (
	// tomCode represents tom code.
	tomCode = nook.Code{
		Value: "cat15"}
)

var (
	// tomAmericanEnglishName represents tom american english name.
	tomAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tom"}

	// tomCanadianFrenchName represents tom canadian french name.
	tomCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tom"}

	// tomDutchName represents tom dutch name.
	tomDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tom"}

	// tomFrenchName represents tom french name.
	tomFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tom"}

	// tomGermanName represents tom german name.
	tomGermanName = nook.Name{
		Language: language.German,
		Value:    "Timo"}

	// tomItalianName represents tom italian name.
	tomItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Isidoro"}

	// tomJapaneseName represents tom japanese name.
	tomJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "バンタム"}

	// tomKoreanName represents tom korean name.
	tomKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "밴덤"}

	// tomLatinAmericanSpanishName represents tom latin american spanish name.
	tomLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Zapirón"}

	// tomRussianName represents tom russian name.
	tomRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Том"}

	// tomSimplifiedChineseName represents tom simplified chinese name.
	tomSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿邦"}

	// tomSpanishName represents tom spanish name.
	tomSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Zapirón"}

	// tomTraditionalChineseName represents tom traditional chinese name.
	tomTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿邦"}
)

var (
	// tomName represents tom name.
	tomName = nook.Languages{
		language.AmericanEnglish:      tomAmericanEnglishName,
		language.CanadianFrench:       tomCanadianFrenchName,
		language.Dutch:                tomDutchName,
		language.French:               tomFrenchName,
		language.German:               tomGermanName,
		language.Italian:              tomItalianName,
		language.Japanese:             tomJapaneseName,
		language.Korean:               tomKoreanName,
		language.LatinAmericanSpanish: tomLatinAmericanSpanishName,
		language.Russian:              tomRussianName,
		language.SimplifiedChinese:    tomSimplifiedChineseName,
		language.Spanish:              tomSpanishName,
		language.TraditionalChinese:   tomTraditionalChineseName}
)

var (
	// tomCharacter represents tom character.
	tomCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: tomBirthday,
		Code:     tomCode,
		Key:      character.Tom,
		Gender:   gender.Male,
		Name:     tomName,
		Special:  false}
)

var (
	// tomAmericanEnglishPhrase represents tom american english phrase.
	tomAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "me-YOWZA"}

	// tomCanadianFrenchPhrase represents tom canadian french phrase.
	tomCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mistigri"}

	// tomDutchPhrase represents tom dutch phrase.
	tomDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "miauw moe"}

	// tomFrenchPhrase represents tom french phrase.
	tomFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mistigri"}

	// tomGermanPhrase represents tom german phrase.
	tomGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "fauch"}

	// tomItalianPhrase represents tom italian phrase.
	tomItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "mi-IAO"}

	// tomJapanesePhrase represents tom japanese phrase.
	tomJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ナックル"}

	// tomKoreanPhrase represents tom korean phrase.
	tomKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쳇"}

	// tomLatinAmericanSpanishPhrase represents tom latin american spanish phrase.
	tomLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "miooou"}

	// tomRussianPhrase represents tom russian phrase.
	tomRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мя-я-яоза"}

	// tomSimplifiedChinesePhrase represents tom simplified chinese phrase.
	tomSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "切"}

	// tomSpanishPhrase represents tom spanish phrase.
	tomSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "miooou"}

	// tomTraditionalChinesePhrase represents tom traditional chinese phrase.
	tomTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "關節"}
)

var (
	// tomPhrase represents tom phrase.
	tomPhrase = nook.Languages{
		language.AmericanEnglish:      tomAmericanEnglishPhrase,
		language.CanadianFrench:       tomCanadianFrenchPhrase,
		language.Dutch:                tomDutchPhrase,
		language.French:               tomFrenchPhrase,
		language.German:               tomGermanPhrase,
		language.Italian:              tomItalianPhrase,
		language.Japanese:             tomJapanesePhrase,
		language.Korean:               tomKoreanPhrase,
		language.LatinAmericanSpanish: tomLatinAmericanSpanishPhrase,
		language.Russian:              tomRussianPhrase,
		language.SimplifiedChinese:    tomSimplifiedChinesePhrase,
		language.Spanish:              tomSpanishPhrase,
		language.TraditionalChinese:   tomTraditionalChinesePhrase}
)

var (
	// Tom represents tom.
	Tom = nook.Villager{
		Character:   tomCharacter,
		Personality: personality.Cranky,
		Phrase:      tomPhrase}
)
