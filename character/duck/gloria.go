package duck

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
	// gloriaBirthday represents gloria birthday.
	gloriaBirthday = nook.Birthday{
		Day:   12,
		Month: time.August}
)

var (
	// gloriaCode represents gloria code.
	gloriaCode = nook.Code{
		Value: "duk15"}
)

var (
	// gloriaAmericanEnglishName represents gloria american english name.
	gloriaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gloria"}

	// gloriaCanadianFrenchName represents gloria canadian french name.
	gloriaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Déborah"}

	// gloriaDutchName represents gloria dutch name.
	gloriaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gloria"}

	// gloriaFrenchName represents gloria french name.
	gloriaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Déborah"}

	// gloriaGermanName represents gloria german name.
	gloriaGermanName = nook.Name{
		Language: language.German,
		Value:    "Gustavia"}

	// gloriaItalianName represents gloria italian name.
	gloriaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pappy"}

	// gloriaJapaneseName represents gloria japanese name.
	gloriaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "スワンソン"}

	// gloriaKoreanName represents gloria korean name.
	gloriaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마릴린"}

	// gloriaLatinAmericanSpanishName represents gloria latin american spanish name.
	gloriaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jarrea"}

	// gloriaRussianName represents gloria russian name.
	gloriaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Глория"}

	// gloriaSimplifiedChineseName represents gloria simplified chinese name.
	gloriaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "素返真"}

	// gloriaSpanishName represents gloria spanish name.
	gloriaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jarrea"}

	// gloriaTraditionalChineseName represents gloria traditional chinese name.
	gloriaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "素返真"}
)

var (
	// gloriaName represents gloria name.
	gloriaName = nook.Languages{
		language.AmericanEnglish:      gloriaAmericanEnglishName,
		language.CanadianFrench:       gloriaCanadianFrenchName,
		language.Dutch:                gloriaDutchName,
		language.French:               gloriaFrenchName,
		language.German:               gloriaGermanName,
		language.Italian:              gloriaItalianName,
		language.Japanese:             gloriaJapaneseName,
		language.Korean:               gloriaKoreanName,
		language.LatinAmericanSpanish: gloriaLatinAmericanSpanishName,
		language.Russian:              gloriaRussianName,
		language.SimplifiedChinese:    gloriaSimplifiedChineseName,
		language.Spanish:              gloriaSpanishName,
		language.TraditionalChinese:   gloriaTraditionalChineseName}
)

var (
	// gloriaCharacter represents gloria character.
	gloriaCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: gloriaBirthday,
		Code:     gloriaCode,
		Key:      character.Gloria,
		Gender:   gender.Female,
		Name:     gloriaName,
		Special:  false}
)

var (
	// gloriaAmericanEnglishPhrase represents gloria american english phrase.
	gloriaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "quacker"}

	// gloriaCanadianFrenchPhrase represents gloria canadian french phrase.
	gloriaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "kwakos"}

	// gloriaDutchPhrase represents gloria dutch phrase.
	gloriaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kwebbel"}

	// gloriaFrenchPhrase represents gloria french phrase.
	gloriaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "kwakos"}

	// gloriaGermanPhrase represents gloria german phrase.
	gloriaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "nerv"}

	// gloriaItalianPhrase represents gloria italian phrase.
	gloriaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "paperonz"}

	// gloriaJapanesePhrase represents gloria japanese phrase.
	gloriaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ぎゃくに"}

	// gloriaKoreanPhrase represents gloria korean phrase.
	gloriaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "유～후"}

	// gloriaLatinAmericanSpanishPhrase represents gloria latin american spanish phrase.
	gloriaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "picodoro"}

	// gloriaRussianPhrase represents gloria russian phrase.
	gloriaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "не кряк ли"}

	// gloriaSimplifiedChinesePhrase represents gloria simplified chinese phrase.
	gloriaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "相反"}

	// gloriaSpanishPhrase represents gloria spanish phrase.
	gloriaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "picodoro"}

	// gloriaTraditionalChinesePhrase represents gloria traditional chinese phrase.
	gloriaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "相反"}
)

var (
	// gloriaPhrase represents gloria phrase.
	gloriaPhrase = nook.Languages{
		language.AmericanEnglish:      gloriaAmericanEnglishPhrase,
		language.CanadianFrench:       gloriaCanadianFrenchPhrase,
		language.Dutch:                gloriaDutchPhrase,
		language.French:               gloriaFrenchPhrase,
		language.German:               gloriaGermanPhrase,
		language.Italian:              gloriaItalianPhrase,
		language.Japanese:             gloriaJapanesePhrase,
		language.Korean:               gloriaKoreanPhrase,
		language.LatinAmericanSpanish: gloriaLatinAmericanSpanishPhrase,
		language.Russian:              gloriaRussianPhrase,
		language.SimplifiedChinese:    gloriaSimplifiedChinesePhrase,
		language.Spanish:              gloriaSpanishPhrase,
		language.TraditionalChinese:   gloriaTraditionalChinesePhrase}
)

var (
	// Gloria represents gloria.
	Gloria = nook.Villager{
		Character:   gloriaCharacter,
		Personality: personality.Snooty,
		Phrase:      gloriaPhrase}
)
