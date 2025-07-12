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
	// texBirthday represents tex birthday.
	texBirthday = nook.Birthday{
		Day:   6,
		Month: time.October}
)

var (
	// texCode represents tex code.
	texCode = nook.Code{
		Value: "pgn12"}
)

var (
	// texAmericanEnglishName represents tex american english name.
	texAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tex"}

	// texCanadianFrenchName represents tex canadian french name.
	texCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Émilien"}

	// texDutchName represents tex dutch name.
	texDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tex"}

	// texFrenchName represents tex french name.
	texFrenchName = nook.Name{
		Language: language.French,
		Value:    "Émilien"}

	// texGermanName represents tex german name.
	texGermanName = nook.Name{
		Language: language.German,
		Value:    "Matze"}

	// texItalianName represents tex italian name.
	texItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Freddy"}

	// texJapaneseName represents tex japanese name.
	texJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ボルト"}

	// texKoreanName represents tex korean name.
	texKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "볼트"}

	// texLatinAmericanSpanishName represents tex latin american spanish name.
	texLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tomeo"}

	// texRussianName represents tex russian name.
	texRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Текс"}

	// texSimplifiedChineseName represents tex simplified chinese name.
	texSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "伏特"}

	// texSpanishName represents tex spanish name.
	texSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tomeo"}

	// texTraditionalChineseName represents tex traditional chinese name.
	texTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "伏特"}
)

var (
	// texName represents tex name.
	texName = nook.Languages{
		language.AmericanEnglish:      texAmericanEnglishName,
		language.CanadianFrench:       texCanadianFrenchName,
		language.Dutch:                texDutchName,
		language.French:               texFrenchName,
		language.German:               texGermanName,
		language.Italian:              texItalianName,
		language.Japanese:             texJapaneseName,
		language.Korean:               texKoreanName,
		language.LatinAmericanSpanish: texLatinAmericanSpanishName,
		language.Russian:              texRussianName,
		language.SimplifiedChinese:    texSimplifiedChineseName,
		language.Spanish:              texSpanishName,
		language.TraditionalChinese:   texTraditionalChineseName}
)

var (
	// texCharacter represents tex character.
	texCharacter = nook.Character{
		Animal:   animal.Penguin,
		Birthday: texBirthday,
		Code:     texCode,
		Key:      character.Tex,
		Gender:   gender.Male,
		Name:     texName,
		Special:  false}
)

var (
	// texAmericanEnglishPhrase represents tex american english phrase.
	texAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "picante"}

	// texCanadianFrenchPhrase represents tex canadian french phrase.
	texCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "gla gla"}

	// texDutchPhrase represents tex dutch phrase.
	texDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "picante"}

	// texFrenchPhrase represents tex french phrase.
	texFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "gla gla"}

	// texGermanPhrase represents tex german phrase.
	texGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "watschl"}

	// texItalianPhrase represents tex italian phrase.
	texItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "frio"}

	// texJapanesePhrase represents tex japanese phrase.
	texJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ベイベッ"}

	// texKoreanPhrase represents tex korean phrase.
	texKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "베이베"}

	// texLatinAmericanSpanishPhrase represents tex latin american spanish phrase.
	texLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "siclaro"}

	// texRussianPhrase represents tex russian phrase.
	texRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "огонек"}

	// texSimplifiedChinesePhrase represents tex simplified chinese phrase.
	texSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "宝宝"}

	// texSpanishPhrase represents tex spanish phrase.
	texSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "siclaro"}

	// texTraditionalChinesePhrase represents tex traditional chinese phrase.
	texTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "寶寶"}
)

var (
	// texPhrase represents tex phrase.
	texPhrase = nook.Languages{
		language.AmericanEnglish:      texAmericanEnglishPhrase,
		language.CanadianFrench:       texCanadianFrenchPhrase,
		language.Dutch:                texDutchPhrase,
		language.French:               texFrenchPhrase,
		language.German:               texGermanPhrase,
		language.Italian:              texItalianPhrase,
		language.Japanese:             texJapanesePhrase,
		language.Korean:               texKoreanPhrase,
		language.LatinAmericanSpanish: texLatinAmericanSpanishPhrase,
		language.Russian:              texRussianPhrase,
		language.SimplifiedChinese:    texSimplifiedChinesePhrase,
		language.Spanish:              texSpanishPhrase,
		language.TraditionalChinese:   texTraditionalChinesePhrase}
)

var (
	// Tex represents tex.
	Tex = nook.Villager{
		Character:   texCharacter,
		Personality: personality.Smug,
		Phrase:      texPhrase}
)
