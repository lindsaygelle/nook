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
	// chelseaBirthday represents chelsea birthday.
	chelseaBirthday = nook.Birthday{
		Day:   18,
		Month: time.January}
)

var (
	// chelseaCode represents chelsea code.
	chelseaCode = nook.Code{
		Value: "der10"}
)

var (
	// chelseaAmericanEnglishName represents chelsea american english name.
	chelseaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chelsea"}

	// chelseaCanadianFrenchName represents chelsea canadian french name.
	chelseaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Chelsea"}

	// chelseaDutchName represents chelsea dutch name.
	chelseaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chelsea"}

	// chelseaFrenchName represents chelsea french name.
	chelseaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chelsea"}

	// chelseaGermanName represents chelsea german name.
	chelseaGermanName = nook.Name{
		Language: language.German,
		Value:    "Chelsea"}

	// chelseaItalianName represents chelsea italian name.
	chelseaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Chelsea"}

	// chelseaJapaneseName represents chelsea japanese name.
	chelseaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チェルシー"}

	// chelseaKoreanName represents chelsea korean name.
	chelseaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "첼시"}

	// chelseaLatinAmericanSpanishName represents chelsea latin american spanish name.
	chelseaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Chelsea"}

	// chelseaRussianName represents chelsea russian name.
	chelseaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Челси"}

	// chelseaSimplifiedChineseName represents chelsea simplified chinese name.
	chelseaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雀儿喜"}

	// chelseaSpanishName represents chelsea spanish name.
	chelseaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Chelsea"}

	// chelseaTraditionalChineseName represents chelsea traditional chinese name.
	chelseaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雀兒喜"}
)

var (
	// chelseaName represents chelsea name.
	chelseaName = nook.Languages{
		language.AmericanEnglish:      chelseaAmericanEnglishName,
		language.CanadianFrench:       chelseaCanadianFrenchName,
		language.Dutch:                chelseaDutchName,
		language.French:               chelseaFrenchName,
		language.German:               chelseaGermanName,
		language.Italian:              chelseaItalianName,
		language.Japanese:             chelseaJapaneseName,
		language.Korean:               chelseaKoreanName,
		language.LatinAmericanSpanish: chelseaLatinAmericanSpanishName,
		language.Russian:              chelseaRussianName,
		language.SimplifiedChinese:    chelseaSimplifiedChineseName,
		language.Spanish:              chelseaSpanishName,
		language.TraditionalChinese:   chelseaTraditionalChineseName}
)

var (
	// chelseaCharacter represents chelsea character.
	chelseaCharacter = nook.Character{
		Animal:   animal.Deer,
		Birthday: chelseaBirthday,
		Code:     chelseaCode,
		Key:      character.Chelsea,
		Gender:   gender.Female,
		Name:     chelseaName,
		Special:  false}
)

var (
	// chelseaAmericanEnglishPhrase represents chelsea american english phrase.
	chelseaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pound cake"}

	// chelseaCanadianFrenchPhrase represents chelsea canadian french phrase.
	chelseaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bibiche"}

	// chelseaDutchPhrase represents chelsea dutch phrase.
	chelseaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "melodietje"}

	// chelseaFrenchPhrase represents chelsea french phrase.
	chelseaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bibiche"}

	// chelseaGermanPhrase represents chelsea german phrase.
	chelseaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "denkpink"}

	// chelseaItalianPhrase represents chelsea italian phrase.
	chelseaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "melodì"}

	// chelseaJapanesePhrase represents chelsea japanese phrase.
	chelseaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "メロメロ"}

	// chelseaKoreanPhrase represents chelsea korean phrase.
	chelseaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "하트하트"}

	// chelseaLatinAmericanSpanishPhrase represents chelsea latin american spanish phrase.
	chelseaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "melomelo"}

	// chelseaRussianPhrase represents chelsea russian phrase.
	chelseaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бисквит"}

	// chelseaSimplifiedChinesePhrase represents chelsea simplified chinese phrase.
	chelseaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喜喜"}

	// chelseaSpanishPhrase represents chelsea spanish phrase.
	chelseaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "melomelo"}

	// chelseaTraditionalChinesePhrase represents chelsea traditional chinese phrase.
	chelseaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喜喜"}
)

var (
	// chelseaPhrase represents chelsea phrase.
	chelseaPhrase = nook.Languages{
		language.AmericanEnglish:      chelseaAmericanEnglishPhrase,
		language.CanadianFrench:       chelseaCanadianFrenchPhrase,
		language.Dutch:                chelseaDutchPhrase,
		language.French:               chelseaFrenchPhrase,
		language.German:               chelseaGermanPhrase,
		language.Italian:              chelseaItalianPhrase,
		language.Japanese:             chelseaJapanesePhrase,
		language.Korean:               chelseaKoreanPhrase,
		language.LatinAmericanSpanish: chelseaLatinAmericanSpanishPhrase,
		language.Russian:              chelseaRussianPhrase,
		language.SimplifiedChinese:    chelseaSimplifiedChinesePhrase,
		language.Spanish:              chelseaSpanishPhrase,
		language.TraditionalChinese:   chelseaTraditionalChinesePhrase}
)

var (
	// Chelsea represents chelsea.
	Chelsea = nook.Villager{
		Character:   chelseaCharacter,
		Personality: personality.Normal,
		Phrase:      chelseaPhrase}
)
