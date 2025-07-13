package frog

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
	// wartjrBirthday represents wartjr birthday.
	wartjrBirthday = nook.Birthday{
		Day:   21,
		Month: time.August}
)

var (
	// wartjrCode represents wartjr code.
	wartjrCode = nook.Code{
		Value: "flg05"}
)

var (
	// wartjrAmericanEnglishName represents wartjr american english name.
	wartjrAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Wart Jr."}

	// wartjrCanadianFrenchName represents wartjr canadian french name.
	wartjrCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Crakos"}

	// wartjrDutchName represents wartjr dutch name.
	wartjrDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Wart Jr."}

	// wartjrFrenchName represents wartjr french name.
	wartjrFrenchName = nook.Name{
		Language: language.French,
		Value:    "Crakos"}

	// wartjrGermanName represents wartjr german name.
	wartjrGermanName = nook.Name{
		Language: language.German,
		Value:    "Warzi"}

	// wartjrItalianName represents wartjr italian name.
	wartjrItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Porro"}

	// wartjrJapaneseName represents wartjr japanese name.
	wartjrJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サム"}

	// wartjrKoreanName represents wartjr korean name.
	wartjrKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "샘"}

	// wartjrLatinAmericanSpanishName represents wartjr latin american spanish name.
	wartjrLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Saponcio"}

	// wartjrRussianName represents wartjr russian name.
	wartjrRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Варт-мл."}

	// wartjrSimplifiedChineseName represents wartjr simplified chinese name.
	wartjrSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "山姆"}

	// wartjrSpanishName represents wartjr spanish name.
	wartjrSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Saponcio"}

	// wartjrTraditionalChineseName represents wartjr traditional chinese name.
	wartjrTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "山姆"}
)

var (
	// wartjrName represents wartjr name.
	wartjrName = nook.Languages{
		language.AmericanEnglish:      wartjrAmericanEnglishName,
		language.CanadianFrench:       wartjrCanadianFrenchName,
		language.Dutch:                wartjrDutchName,
		language.French:               wartjrFrenchName,
		language.German:               wartjrGermanName,
		language.Italian:              wartjrItalianName,
		language.Japanese:             wartjrJapaneseName,
		language.Korean:               wartjrKoreanName,
		language.LatinAmericanSpanish: wartjrLatinAmericanSpanishName,
		language.Russian:              wartjrRussianName,
		language.SimplifiedChinese:    wartjrSimplifiedChineseName,
		language.Spanish:              wartjrSpanishName,
		language.TraditionalChinese:   wartjrTraditionalChineseName}
)

var (
	// wartjrCharacter represents wartjr character.
	wartjrCharacter = nook.Character{
		Animal:   animal.Frog,
		Birthday: wartjrBirthday,
		Code:     wartjrCode,
		Key:      character.WartJr,
		Gender:   gender.Male,
		Name:     wartjrName,
		Special:  false}
)

var (
	// wartjrAmericanEnglishPhrase represents wartjr american english phrase.
	wartjrAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "grr-ribbit"}

	// wartjrCanadianFrenchPhrase represents wartjr canadian french phrase.
	wartjrCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "croa-croa"}

	// wartjrDutchPhrase represents wartjr dutch phrase.
	wartjrDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "brulkikker"}

	// wartjrFrenchPhrase represents wartjr french phrase.
	wartjrFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "croa-croa"}

	// wartjrGermanPhrase represents wartjr german phrase.
	wartjrGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grupap"}

	// wartjrItalianPhrase represents wartjr italian phrase.
	wartjrItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cracragnam"}

	// wartjrJapanesePhrase represents wartjr japanese phrase.
	wartjrJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だぎゃ"}

	// wartjrKoreanPhrase represents wartjr korean phrase.
	wartjrKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "므흣"}

	// wartjrLatinAmericanSpanishPhrase represents wartjr latin american spanish phrase.
	wartjrLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grroac"}

	// wartjrRussianPhrase represents wartjr russian phrase.
	wartjrRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ква-ква"}

	// wartjrSimplifiedChinesePhrase represents wartjr simplified chinese phrase.
	wartjrSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "山山"}

	// wartjrSpanishPhrase represents wartjr spanish phrase.
	wartjrSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "grroac"}

	// wartjrTraditionalChinesePhrase represents wartjr traditional chinese phrase.
	wartjrTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "山山"}
)

var (
	// wartjrPhrase represents wartjr phrase.
	wartjrPhrase = nook.Languages{
		language.AmericanEnglish:      wartjrAmericanEnglishPhrase,
		language.CanadianFrench:       wartjrCanadianFrenchPhrase,
		language.Dutch:                wartjrDutchPhrase,
		language.French:               wartjrFrenchPhrase,
		language.German:               wartjrGermanPhrase,
		language.Italian:              wartjrItalianPhrase,
		language.Japanese:             wartjrJapanesePhrase,
		language.Korean:               wartjrKoreanPhrase,
		language.LatinAmericanSpanish: wartjrLatinAmericanSpanishPhrase,
		language.Russian:              wartjrRussianPhrase,
		language.SimplifiedChinese:    wartjrSimplifiedChinesePhrase,
		language.Spanish:              wartjrSpanishPhrase,
		language.TraditionalChinese:   wartjrTraditionalChinesePhrase}
)

var (
	// WartJr represents wart jr.
	WartJr = nook.Villager{
		Character:   wartjrCharacter,
		Personality: personality.Cranky,
		Phrase:      wartjrPhrase}
)
