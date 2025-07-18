package penguin

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
	// puckBirthday represents puck birthday.
	puckBirthday = nook.Birthday{
		Day:   21,
		Month: time.February}
)

var (
	// puckCode represents puck code.
	puckCode = nook.Code{
		Value: "pgn06"}
)

var (
	// puckAmericanEnglishName represents puck american english name.
	puckAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Puck"}

	// puckCanadianFrenchName represents puck canadian french name.
	puckCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Hervé"}

	// puckDutchName represents puck dutch name.
	puckDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Puck"}

	// puckFrenchName represents puck french name.
	puckFrenchName = nook.Name{
		Language: language.French,
		Value:    "Hervé"}

	// puckGermanName represents puck german name.
	puckGermanName = nook.Name{
		Language: language.German,
		Value:    "Puck"}

	// puckItalianName represents puck italian name.
	puckItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Glacido"}

	// puckJapaneseName represents puck japanese name.
	puckJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ホッケー"}

	// puckKoreanName represents puck korean name.
	puckKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "하키"}

	// puckLatinAmericanSpanishName represents puck latin american spanish name.
	puckLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Fredo"}

	// puckRussianName represents puck russian name.
	puckRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пак"}

	// puckSimplifiedChineseName represents puck simplified chinese name.
	puckSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哈奇"}

	// puckSpanishName represents puck spanish name.
	puckSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Fredo"}

	// puckTraditionalChineseName represents puck traditional chinese name.
	puckTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哈奇"}
)

var (
	// puckName represents puck name.
	puckName = nook.Languages{
		language.AmericanEnglish:      puckAmericanEnglishName,
		language.CanadianFrench:       puckCanadianFrenchName,
		language.Dutch:                puckDutchName,
		language.French:               puckFrenchName,
		language.German:               puckGermanName,
		language.Italian:              puckItalianName,
		language.Japanese:             puckJapaneseName,
		language.Korean:               puckKoreanName,
		language.LatinAmericanSpanish: puckLatinAmericanSpanishName,
		language.Russian:              puckRussianName,
		language.SimplifiedChinese:    puckSimplifiedChineseName,
		language.Spanish:              puckSpanishName,
		language.TraditionalChinese:   puckTraditionalChineseName}
)

var (
	// puckCharacter represents puck character.
	puckCharacter = nook.Character{
		Animal:   animal.Penguin,
		Birthday: puckBirthday,
		Code:     puckCode,
		Key:      character.Puck,
		Gender:   gender.Male,
		Name:     puckName,
		Special:  false}
)

var (
	// puckAmericanEnglishPhrase represents puck american english phrase.
	puckAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "brrrrrrrrr"}

	// puckCanadianFrenchPhrase represents puck canadian french phrase.
	puckCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "brrrrrrrrr"}

	// puckDutchPhrase represents puck dutch phrase.
	puckDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "brrrrrrrrr"}

	// puckFrenchPhrase represents puck french phrase.
	puckFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "brrrrrrrrr"}

	// puckGermanPhrase represents puck german phrase.
	puckGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brrrrrrrrr"}

	// puckItalianPhrase represents puck italian phrase.
	puckItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "brrrrrrrrr"}

	// puckJapanesePhrase represents puck japanese phrase.
	puckJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "さぶー"}

	// puckKoreanPhrase represents puck korean phrase.
	puckKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "엣취"}

	// puckLatinAmericanSpanishPhrase represents puck latin american spanish phrase.
	puckLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "brrr-uah"}

	// puckRussianPhrase represents puck russian phrase.
	puckRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "брр"}

	// puckSimplifiedChinesePhrase represents puck simplified chinese phrase.
	puckSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "候补"}

	// puckSpanishPhrase represents puck spanish phrase.
	puckSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "sardinilla"}

	// puckTraditionalChinesePhrase represents puck traditional chinese phrase.
	puckTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "候補"}
)

var (
	// puckPhrase represents puck phrase.
	puckPhrase = nook.Languages{
		language.AmericanEnglish:      puckAmericanEnglishPhrase,
		language.CanadianFrench:       puckCanadianFrenchPhrase,
		language.Dutch:                puckDutchPhrase,
		language.French:               puckFrenchPhrase,
		language.German:               puckGermanPhrase,
		language.Italian:              puckItalianPhrase,
		language.Japanese:             puckJapanesePhrase,
		language.Korean:               puckKoreanPhrase,
		language.LatinAmericanSpanish: puckLatinAmericanSpanishPhrase,
		language.Russian:              puckRussianPhrase,
		language.SimplifiedChinese:    puckSimplifiedChinesePhrase,
		language.Spanish:              puckSpanishPhrase,
		language.TraditionalChinese:   puckTraditionalChinesePhrase}
)

var (
	// Puck represents puck.
	Puck = nook.Villager{
		Character:   puckCharacter,
		Personality: personality.Lazy,
		Phrase:      puckPhrase}
)
