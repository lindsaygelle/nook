package koala

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
	// melbaBirthday represents melba birthday.
	melbaBirthday = nook.Birthday{
		Day:   12,
		Month: time.April}
)

var (
	// melbaCode represents melba code.
	melbaCode = nook.Code{
		Value: "kal02"}
)

var (
	// melbaAmericanEnglishName represents melba american english name.
	melbaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Melba"}

	// melbaCanadianFrenchName represents melba canadian french name.
	melbaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Melba"}

	// melbaDutchName represents melba dutch name.
	melbaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Melba"}

	// melbaFrenchName represents melba french name.
	melbaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Melba"}

	// melbaGermanName represents melba german name.
	melbaGermanName = nook.Name{
		Language: language.German,
		Value:    "Kornelia"}

	// melbaItalianName represents melba italian name.
	melbaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Melba"}

	// melbaJapaneseName represents melba japanese name.
	melbaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アデレード"}

	// melbaKoreanName represents melba korean name.
	melbaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아델레이드"}

	// melbaLatinAmericanSpanishName represents melba latin american spanish name.
	melbaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Melba"}

	// melbaRussianName represents melba russian name.
	melbaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мельба"}

	// melbaSimplifiedChineseName represents melba simplified chinese name.
	melbaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿得"}

	// melbaSpanishName represents melba spanish name.
	melbaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Melba"}

	// melbaTraditionalChineseName represents melba traditional chinese name.
	melbaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿得"}
)

var (
	// melbaName represents melba name.
	melbaName = nook.Languages{
		language.AmericanEnglish:      melbaAmericanEnglishName,
		language.CanadianFrench:       melbaCanadianFrenchName,
		language.Dutch:                melbaDutchName,
		language.French:               melbaFrenchName,
		language.German:               melbaGermanName,
		language.Italian:              melbaItalianName,
		language.Japanese:             melbaJapaneseName,
		language.Korean:               melbaKoreanName,
		language.LatinAmericanSpanish: melbaLatinAmericanSpanishName,
		language.Russian:              melbaRussianName,
		language.SimplifiedChinese:    melbaSimplifiedChineseName,
		language.Spanish:              melbaSpanishName,
		language.TraditionalChinese:   melbaTraditionalChineseName}
)

var (
	// melbaCharacter represents melba character.
	melbaCharacter = nook.Character{
		Animal:   animal.Koala,
		Birthday: melbaBirthday,
		Code:     melbaCode,
		Key:      character.Melba,
		Gender:   gender.Female,
		Name:     melbaName,
		Special:  false}
)

var (
	// melbaAmericanEnglishPhrase represents melba american english phrase.
	melbaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "toasty"}

	// melbaCanadianFrenchPhrase represents melba canadian french phrase.
	melbaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pêchou"}

	// melbaDutchPhrase represents melba dutch phrase.
	melbaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "toastje"}

	// melbaFrenchPhrase represents melba french phrase.
	melbaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pêchou"}

	// melbaGermanPhrase represents melba german phrase.
	melbaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zubba"}

	// melbaItalianPhrase represents melba italian phrase.
	melbaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "toasty"}

	// melbaJapanesePhrase represents melba japanese phrase.
	melbaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "とっても"}

	// melbaKoreanPhrase represents melba korean phrase.
	melbaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "정말로요"}

	// melbaLatinAmericanSpanishPhrase represents melba latin american spanish phrase.
	melbaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "pichú"}

	// melbaRussianPhrase represents melba russian phrase.
	melbaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрустяшка"}

	// melbaSimplifiedChinesePhrase represents melba simplified chinese phrase.
	melbaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "不得了"}

	// melbaSpanishPhrase represents melba spanish phrase.
	melbaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tesoro"}

	// melbaTraditionalChinesePhrase represents melba traditional chinese phrase.
	melbaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "不得了"}
)

var (
	// melbaPhrase represents melba phrase.
	melbaPhrase = nook.Languages{
		language.AmericanEnglish:      melbaAmericanEnglishPhrase,
		language.CanadianFrench:       melbaCanadianFrenchPhrase,
		language.Dutch:                melbaDutchPhrase,
		language.French:               melbaFrenchPhrase,
		language.German:               melbaGermanPhrase,
		language.Italian:              melbaItalianPhrase,
		language.Japanese:             melbaJapanesePhrase,
		language.Korean:               melbaKoreanPhrase,
		language.LatinAmericanSpanish: melbaLatinAmericanSpanishPhrase,
		language.Russian:              melbaRussianPhrase,
		language.SimplifiedChinese:    melbaSimplifiedChinesePhrase,
		language.Spanish:              melbaSpanishPhrase,
		language.TraditionalChinese:   melbaTraditionalChinesePhrase}
)

var (
	// Melba represents melba.
	Melba = nook.Villager{
		Character:   melbaCharacter,
		Personality: personality.Normal,
		Phrase:      melbaPhrase}
)
