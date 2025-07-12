package elephant

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
	// tiaBirthday represents tia birthday.
	tiaBirthday = nook.Birthday{
		Day:   18,
		Month: time.November}
)

var (
	// tiaCode represents tia code.
	tiaCode = nook.Code{
		Value: "elp10"}
)

var (
	// tiaAmericanEnglishName represents tia american english name.
	tiaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tia"}

	// tiaCanadianFrenchName represents tia canadian french name.
	tiaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Fanny"}

	// tiaDutchName represents tia dutch name.
	tiaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tia"}

	// tiaFrenchName represents tia french name.
	tiaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Fanny"}

	// tiaGermanName represents tia german name.
	tiaGermanName = nook.Name{
		Language: language.German,
		Value:    "Eleonore"}

	// tiaItalianName represents tia italian name.
	tiaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fanny"}

	// tiaJapaneseName represents tia japanese name.
	tiaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ティーナ"}

	// tiaKoreanName represents tia korean name.
	tiaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "티나"}

	// tiaLatinAmericanSpanishName represents tia latin american spanish name.
	tiaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Meralda"}

	// tiaRussianName represents tia russian name.
	tiaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тиа"}

	// tiaSimplifiedChineseName represents tia simplified chinese name.
	tiaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "茉莉"}

	// tiaSpanishName represents tia spanish name.
	tiaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Meralda"}

	// tiaTraditionalChineseName represents tia traditional chinese name.
	tiaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "茉莉"}
)

var (
	// tiaName represents tia name.
	tiaName = nook.Languages{
		language.AmericanEnglish:      tiaAmericanEnglishName,
		language.CanadianFrench:       tiaCanadianFrenchName,
		language.Dutch:                tiaDutchName,
		language.French:               tiaFrenchName,
		language.German:               tiaGermanName,
		language.Italian:              tiaItalianName,
		language.Japanese:             tiaJapaneseName,
		language.Korean:               tiaKoreanName,
		language.LatinAmericanSpanish: tiaLatinAmericanSpanishName,
		language.Russian:              tiaRussianName,
		language.SimplifiedChinese:    tiaSimplifiedChineseName,
		language.Spanish:              tiaSpanishName,
		language.TraditionalChinese:   tiaTraditionalChineseName}
)

var (
	// tiaCharacter represents tia character.
	tiaCharacter = nook.Character{
		Animal:   animal.Elephant,
		Birthday: tiaBirthday,
		Code:     tiaCode,
		Key:      character.Tia,
		Gender:   gender.Female,
		Name:     tiaName,
		Special:  false}
)

var (
	// tiaAmericanEnglishPhrase represents tia american english phrase.
	tiaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "teacup"}

	// tiaCanadianFrenchPhrase represents tia canadian french phrase.
	tiaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "énoooorme"}

	// tiaDutchPhrase represents tia dutch phrase.
	tiaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kopje thee"}

	// tiaFrenchPhrase represents tia french phrase.
	tiaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "énoooorme"}

	// tiaGermanPhrase represents tia german phrase.
	tiaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "huiuiui"}

	// tiaItalianPhrase represents tia italian phrase.
	tiaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "parbleu"}

	// tiaJapanesePhrase represents tia japanese phrase.
	tiaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ホッ"}

	// tiaKoreanPhrase represents tia korean phrase.
	tiaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "따끈따끈"}

	// tiaLatinAmericanSpanishPhrase represents tia latin american spanish phrase.
	tiaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "snuuuf"}

	// tiaRussianPhrase represents tia russian phrase.
	tiaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "чайку-у"}

	// tiaSimplifiedChinesePhrase represents tia simplified chinese phrase.
	tiaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呼"}

	// tiaSpanishPhrase represents tia spanish phrase.
	tiaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "snuuuf"}

	// tiaTraditionalChinesePhrase represents tia traditional chinese phrase.
	tiaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "呼"}
)

var (
	// tiaPhrase represents tia phrase.
	tiaPhrase = nook.Languages{
		language.AmericanEnglish:      tiaAmericanEnglishPhrase,
		language.CanadianFrench:       tiaCanadianFrenchPhrase,
		language.Dutch:                tiaDutchPhrase,
		language.French:               tiaFrenchPhrase,
		language.German:               tiaGermanPhrase,
		language.Italian:              tiaItalianPhrase,
		language.Japanese:             tiaJapanesePhrase,
		language.Korean:               tiaKoreanPhrase,
		language.LatinAmericanSpanish: tiaLatinAmericanSpanishPhrase,
		language.Russian:              tiaRussianPhrase,
		language.SimplifiedChinese:    tiaSimplifiedChinesePhrase,
		language.Spanish:              tiaSpanishPhrase,
		language.TraditionalChinese:   tiaTraditionalChinesePhrase}
)

var (
	// Tia represents tia.
	Tia = nook.Villager{
		Character:   tiaCharacter,
		Personality: personality.Normal,
		Phrase:      tiaPhrase}
)
