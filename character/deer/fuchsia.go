package deer

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
	// fuchsiaBirthday represents fuchsia birthday.
	fuchsiaBirthday = nook.Birthday{
		Day:   19,
		Month: time.September}
)

var (
	// fuchsiaCode represents fuchsia code.
	fuchsiaCode = nook.Code{
		Value: "der06"}
)

var (
	// fuchsiaAmericanEnglishName represents fuchsia american english name.
	fuchsiaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Fuchsia"}

	// fuchsiaCanadianFrenchName represents fuchsia canadian french name.
	fuchsiaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rosanne"}

	// fuchsiaDutchName represents fuchsia dutch name.
	fuchsiaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Fuchsia"}

	// fuchsiaFrenchName represents fuchsia french name.
	fuchsiaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rosanne"}

	// fuchsiaGermanName represents fuchsia german name.
	fuchsiaGermanName = nook.Name{
		Language: language.German,
		Value:    "Selina"}

	// fuchsiaItalianName represents fuchsia italian name.
	fuchsiaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fuxia"}

	// fuchsiaJapaneseName represents fuchsia japanese name.
	fuchsiaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジェシカ"}

	// fuchsiaKoreanName represents fuchsia korean name.
	fuchsiaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "제시카"}

	// fuchsiaLatinAmericanSpanishName represents fuchsia latin american spanish name.
	fuchsiaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rosalina"}

	// fuchsiaRussianName represents fuchsia russian name.
	fuchsiaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фуксия"}

	// fuchsiaSimplifiedChineseName represents fuchsia simplified chinese name.
	fuchsiaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "洁西卡"}

	// fuchsiaSpanishName represents fuchsia spanish name.
	fuchsiaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rosalina"}

	// fuchsiaTraditionalChineseName represents fuchsia traditional chinese name.
	fuchsiaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "潔西卡"}
)

var (
	// fuchsiaName represents fuchsia name.
	fuchsiaName = nook.Languages{
		language.AmericanEnglish:      fuchsiaAmericanEnglishName,
		language.CanadianFrench:       fuchsiaCanadianFrenchName,
		language.Dutch:                fuchsiaDutchName,
		language.French:               fuchsiaFrenchName,
		language.German:               fuchsiaGermanName,
		language.Italian:              fuchsiaItalianName,
		language.Japanese:             fuchsiaJapaneseName,
		language.Korean:               fuchsiaKoreanName,
		language.LatinAmericanSpanish: fuchsiaLatinAmericanSpanishName,
		language.Russian:              fuchsiaRussianName,
		language.SimplifiedChinese:    fuchsiaSimplifiedChineseName,
		language.Spanish:              fuchsiaSpanishName,
		language.TraditionalChinese:   fuchsiaTraditionalChineseName}
)

var (
	// fuchsiaCharacter represents fuchsia character.
	fuchsiaCharacter = nook.Character{
		Animal:   animal.Deer,
		Birthday: fuchsiaBirthday,
		Code:     fuchsiaCode,
		Key:      character.Fuchsia,
		Gender:   gender.Female,
		Name:     fuchsiaName,
		Special:  false}
)

var (
	// fuchsiaAmericanEnglishPhrase represents fuchsia american english phrase.
	fuchsiaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "girlfriend"}

	// fuchsiaCanadianFrenchPhrase represents fuchsia canadian french phrase.
	fuchsiaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "frambou"}

	// fuchsiaDutchPhrase represents fuchsia dutch phrase.
	fuchsiaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "meid"}

	// fuchsiaFrenchPhrase represents fuchsia french phrase.
	fuchsiaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "rosace"}

	// fuchsiaGermanPhrase represents fuchsia german phrase.
	fuchsiaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zick"}

	// fuchsiaItalianPhrase represents fuchsia italian phrase.
	fuchsiaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cervilà"}

	// fuchsiaJapanesePhrase represents fuchsia japanese phrase.
	fuchsiaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なんしか"}

	// fuchsiaKoreanPhrase represents fuchsia korean phrase.
	fuchsiaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "핑크크"}

	// fuchsiaLatinAmericanSpanishPhrase represents fuchsia latin american spanish phrase.
	fuchsiaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "braaaam"}

	// fuchsiaRussianPhrase represents fuchsia russian phrase.
	fuchsiaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "зазноба"}

	// fuchsiaSimplifiedChinesePhrase represents fuchsia simplified chinese phrase.
	fuchsiaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "西卡"}

	// fuchsiaSpanishPhrase represents fuchsia spanish phrase.
	fuchsiaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "braaaam"}

	// fuchsiaTraditionalChinesePhrase represents fuchsia traditional chinese phrase.
	fuchsiaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "西卡"}
)

var (
	// fuchsiaPhrase represents fuchsia phrase.
	fuchsiaPhrase = nook.Languages{
		language.AmericanEnglish:      fuchsiaAmericanEnglishPhrase,
		language.CanadianFrench:       fuchsiaCanadianFrenchPhrase,
		language.Dutch:                fuchsiaDutchPhrase,
		language.French:               fuchsiaFrenchPhrase,
		language.German:               fuchsiaGermanPhrase,
		language.Italian:              fuchsiaItalianPhrase,
		language.Japanese:             fuchsiaJapanesePhrase,
		language.Korean:               fuchsiaKoreanPhrase,
		language.LatinAmericanSpanish: fuchsiaLatinAmericanSpanishPhrase,
		language.Russian:              fuchsiaRussianPhrase,
		language.SimplifiedChinese:    fuchsiaSimplifiedChinesePhrase,
		language.Spanish:              fuchsiaSpanishPhrase,
		language.TraditionalChinese:   fuchsiaTraditionalChinesePhrase}
)

var (
	// Fuchsia represents fuchsia.
	Fuchsia = nook.Villager{
		Character:   fuchsiaCharacter,
		Personality: personality.BigSister,
		Phrase:      fuchsiaPhrase}
)
