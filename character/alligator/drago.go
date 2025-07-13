package alligator

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// dragoBirthday represents Drago's birthday.
var (
	// dragoBirthday represents drago birthday.
	dragoBirthday = nook.Birthday{
		Day:   12,
		Month: time.February}
)

// dragoCode represents Drago's unique code.
var (
	// dragoCode represents drago code.
	dragoCode = nook.Code{
		Value: "crd08"}
)

// Different names for Drago in various languages.
var (
	// dragoAmericanEnglishName represents drago american english name.
	dragoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Drago"}

	// dragoCanadianFrenchName represents drago canadian french name.
	dragoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Drago"}

	// dragoDutchName represents drago dutch name.
	dragoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Drago"}

	// dragoFrenchName represents drago french name.
	dragoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Drago"}

	// dragoGermanName represents drago german name.
	dragoGermanName = nook.Name{
		Language: language.German,
		Value:    "Frederik"}

	// dragoItalianName represents drago italian name.
	dragoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Dragonio"}

	// dragoJapaneseName represents drago japanese name.
	dragoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タツオ"}

	// dragoKoreanName represents drago korean name.
	dragoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "용남이"}

	// dragoLatinAmericanSpanishName represents drago latin american spanish name.
	dragoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Dragonio"}

	// dragoRussianName represents drago russian name.
	dragoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Драго"}

	// dragoSimplifiedChineseName represents drago simplified chinese name.
	dragoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿龙"}

	// dragoSpanishName represents drago spanish name.
	dragoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Dragonio"}

	// dragoTraditionalChineseName represents drago traditional chinese name.
	dragoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿龍"}
)

// dragoName represents Drago's name in different languages.
var (
	// dragoName represents drago name.
	dragoName = nook.Languages{
		language.AmericanEnglish:      dragoAmericanEnglishName,
		language.CanadianFrench:       dragoCanadianFrenchName,
		language.Dutch:                dragoDutchName,
		language.French:               dragoFrenchName,
		language.German:               dragoGermanName,
		language.Italian:              dragoItalianName,
		language.Japanese:             dragoJapaneseName,
		language.Korean:               dragoKoreanName,
		language.LatinAmericanSpanish: dragoLatinAmericanSpanishName,
		language.Russian:              dragoRussianName,
		language.SimplifiedChinese:    dragoSimplifiedChineseName,
		language.Spanish:              dragoSpanishName,
		language.TraditionalChinese:   dragoTraditionalChineseName}
)

// dragoCharacter represents Drago's character information.
var (
	// dragoCharacter represents drago character.
	dragoCharacter = nook.Character{
		Animal:   animal.Alligator,
		Birthday: dragoBirthday,
		Code:     dragoCode,
		Key:      character.Drago,
		Gender:   gender.Male,
		Name:     dragoName,
		Special:  false}
)

// Different phrases spoken by Drago in various languages.
var (
	// dragoAmericanEnglishPhrase represents drago american english phrase.
	dragoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "burrrn"}

	// dragoCanadianFrenchPhrase represents drago canadian french phrase.
	dragoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ducroc"}

	// dragoDutchPhrase represents drago dutch phrase.
	dragoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "puf"}

	// dragoFrenchPhrase represents drago french phrase.
	dragoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ducroc"}

	// dragoGermanPhrase represents drago german phrase.
	dragoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hömpf"}

	// dragoItalianPhrase represents drago italian phrase.
	dragoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "tic tac"}

	// dragoJapanesePhrase represents drago japanese phrase.
	dragoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "えーと"}

	// dragoKoreanPhrase represents drago korean phrase.
	dragoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "띠용띠용"}

	// dragoLatinAmericanSpanishPhrase represents drago latin american spanish phrase.
	dragoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñami"}

	// dragoRussianPhrase represents drago russian phrase.
	dragoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "огонь"}

	// dragoSimplifiedChinesePhrase represents drago simplified chinese phrase.
	dragoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "然后"}

	// dragoSpanishPhrase represents drago spanish phrase.
	dragoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "quemarrr"}

	// dragoTraditionalChinesePhrase represents drago traditional chinese phrase.
	dragoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "然後"}
)

// dragoPhrase represents Drago's phrases in different languages.
var (
	// dragoPhrase represents drago phrase.
	dragoPhrase = nook.Languages{
		language.AmericanEnglish:      dragoAmericanEnglishPhrase,
		language.CanadianFrench:       dragoCanadianFrenchPhrase,
		language.Dutch:                dragoDutchPhrase,
		language.French:               dragoFrenchPhrase,
		language.German:               dragoGermanPhrase,
		language.Italian:              dragoItalianPhrase,
		language.Japanese:             dragoJapanesePhrase,
		language.Korean:               dragoKoreanPhrase,
		language.LatinAmericanSpanish: dragoLatinAmericanSpanishPhrase,
		language.Russian:              dragoRussianPhrase,
		language.SimplifiedChinese:    dragoSimplifiedChinesePhrase,
		language.Spanish:              dragoSpanishPhrase,
		language.TraditionalChinese:   dragoTraditionalChinesePhrase}
)

// Drago represents the character Drago with his complete information.
var (
	// Drago represents drago.
	Drago = nook.Villager{
		Character:   dragoCharacter,
		Personality: personality.Lazy,
		Phrase:      dragoPhrase}
)
