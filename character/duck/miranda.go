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
	// mirandaBirthday represents miranda birthday.
	mirandaBirthday = nook.Birthday{
		Day:   23,
		Month: time.April}
)

var (
	// mirandaCode represents miranda code.
	mirandaCode = nook.Code{
		Value: "duk12"}
)

var (
	// mirandaAmericanEnglishName represents miranda american english name.
	mirandaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Miranda"}

	// mirandaCanadianFrenchName represents miranda canadian french name.
	mirandaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Maëllis"}

	// mirandaDutchName represents miranda dutch name.
	mirandaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Miranda"}

	// mirandaFrenchName represents miranda french name.
	mirandaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maëllis"}

	// mirandaGermanName represents miranda german name.
	mirandaGermanName = nook.Name{
		Language: language.German,
		Value:    "Tanya"}

	// mirandaItalianName represents miranda italian name.
	mirandaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Miranda"}

	// mirandaJapaneseName represents miranda japanese name.
	mirandaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ミランダ"}

	// mirandaKoreanName represents miranda korean name.
	mirandaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "미란다"}

	// mirandaLatinAmericanSpanishName represents miranda latin american spanish name.
	mirandaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Miranda"}

	// mirandaRussianName represents miranda russian name.
	mirandaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Миранда"}

	// mirandaSimplifiedChineseName represents miranda simplified chinese name.
	mirandaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "米兰达"}

	// mirandaSpanishName represents miranda spanish name.
	mirandaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Miranda"}

	// mirandaTraditionalChineseName represents miranda traditional chinese name.
	mirandaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "米蘭達"}
)

var (
	// mirandaName represents miranda name.
	mirandaName = nook.Languages{
		language.AmericanEnglish:      mirandaAmericanEnglishName,
		language.CanadianFrench:       mirandaCanadianFrenchName,
		language.Dutch:                mirandaDutchName,
		language.French:               mirandaFrenchName,
		language.German:               mirandaGermanName,
		language.Italian:              mirandaItalianName,
		language.Japanese:             mirandaJapaneseName,
		language.Korean:               mirandaKoreanName,
		language.LatinAmericanSpanish: mirandaLatinAmericanSpanishName,
		language.Russian:              mirandaRussianName,
		language.SimplifiedChinese:    mirandaSimplifiedChineseName,
		language.Spanish:              mirandaSpanishName,
		language.TraditionalChinese:   mirandaTraditionalChineseName}
)

var (
	// mirandaCharacter represents miranda character.
	mirandaCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: mirandaBirthday,
		Code:     mirandaCode,
		Key:      character.Miranda,
		Gender:   gender.Female,
		Name:     mirandaName,
		Special:  false}
)

var (
	// mirandaAmericanEnglishPhrase represents miranda american english phrase.
	mirandaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "quackulous"}

	// mirandaCanadianFrenchPhrase represents miranda canadian french phrase.
	mirandaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "couac"}

	// mirandaDutchPhrase represents miranda dutch phrase.
	mirandaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kwektakel"}

	// mirandaFrenchPhrase represents miranda french phrase.
	mirandaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "couac"}

	// mirandaGermanPhrase represents miranda german phrase.
	mirandaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brotbrot"}

	// mirandaItalianPhrase represents miranda italian phrase.
	mirandaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quaquao"}

	// mirandaJapanesePhrase represents miranda japanese phrase.
	mirandaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なにさ"}

	// mirandaKoreanPhrase represents miranda korean phrase.
	mirandaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "괜찮아"}

	// mirandaLatinAmericanSpanishPhrase represents miranda latin american spanish phrase.
	mirandaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cua-cuá"}

	// mirandaRussianPhrase represents miranda russian phrase.
	mirandaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "покрясающе"}

	// mirandaSimplifiedChinesePhrase represents miranda simplified chinese phrase.
	mirandaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "怎么回事"}

	// mirandaSpanishPhrase represents miranda spanish phrase.
	mirandaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cua-cuá"}

	// mirandaTraditionalChinesePhrase represents miranda traditional chinese phrase.
	mirandaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "怎麼回事"}
)

var (
	// mirandaPhrase represents miranda phrase.
	mirandaPhrase = nook.Languages{
		language.AmericanEnglish:      mirandaAmericanEnglishPhrase,
		language.CanadianFrench:       mirandaCanadianFrenchPhrase,
		language.Dutch:                mirandaDutchPhrase,
		language.French:               mirandaFrenchPhrase,
		language.German:               mirandaGermanPhrase,
		language.Italian:              mirandaItalianPhrase,
		language.Japanese:             mirandaJapanesePhrase,
		language.Korean:               mirandaKoreanPhrase,
		language.LatinAmericanSpanish: mirandaLatinAmericanSpanishPhrase,
		language.Russian:              mirandaRussianPhrase,
		language.SimplifiedChinese:    mirandaSimplifiedChinesePhrase,
		language.Spanish:              mirandaSpanishPhrase,
		language.TraditionalChinese:   mirandaTraditionalChinesePhrase}
)

var (
	// Miranda represents miranda.
	Miranda = nook.Villager{
		Character:   mirandaCharacter,
		Personality: personality.Snooty,
		Phrase:      mirandaPhrase}
)
