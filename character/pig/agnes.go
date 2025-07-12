package pig

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
	// agnesBirthday represents agnes birthday.
	agnesBirthday = nook.Birthday{
		Day:   21,
		Month: time.April}
)

var (
	// agnesCode represents agnes code.
	agnesCode = nook.Code{
		Value: "pig17"}
)

var (
	// agnesAmericanEnglishName represents agnes american english name.
	agnesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Agnes"}

	// agnesCanadianFrenchName represents agnes canadian french name.
	agnesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pansy"}

	// agnesDutchName represents agnes dutch name.
	agnesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Agnes"}

	// agnesFrenchName represents agnes french name.
	agnesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pansy"}

	// agnesGermanName represents agnes german name.
	agnesGermanName = nook.Name{
		Language: language.German,
		Value:    "Nora"}

	// agnesItalianName represents agnes italian name.
	agnesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Maia"}

	// agnesJapaneseName represents agnes japanese name.
	agnesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アグネス"}

	// agnesKoreanName represents agnes korean name.
	agnesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아그네스"}

	// agnesLatinAmericanSpanishName represents agnes latin american spanish name.
	agnesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nerea"}

	// agnesRussianName represents agnes russian name.
	agnesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Агнес"}

	// agnesSimplifiedChineseName represents agnes simplified chinese name.
	agnesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "古乃欣"}

	// agnesSpanishName represents agnes spanish name.
	agnesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Negrea"}

	// agnesTraditionalChineseName represents agnes traditional chinese name.
	agnesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "古乃欣"}
)

var (
	// agnesName represents agnes name.
	agnesName = nook.Languages{
		language.AmericanEnglish:      agnesAmericanEnglishName,
		language.CanadianFrench:       agnesCanadianFrenchName,
		language.Dutch:                agnesDutchName,
		language.French:               agnesFrenchName,
		language.German:               agnesGermanName,
		language.Italian:              agnesItalianName,
		language.Japanese:             agnesJapaneseName,
		language.Korean:               agnesKoreanName,
		language.LatinAmericanSpanish: agnesLatinAmericanSpanishName,
		language.Russian:              agnesRussianName,
		language.SimplifiedChinese:    agnesSimplifiedChineseName,
		language.Spanish:              agnesSpanishName,
		language.TraditionalChinese:   agnesTraditionalChineseName}
)

var (
	// agnesCharacter represents agnes character.
	agnesCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: agnesBirthday,
		Code:     agnesCode,
		Key:      character.Agnes,
		Gender:   gender.Female,
		Name:     agnesName,
		Special:  false}
)

var (
	// agnesAmericanEnglishPhrase represents agnes american english phrase.
	agnesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snuffle"}

	// agnesCanadianFrenchPhrase represents agnes canadian french phrase.
	agnesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "sauciflard"}

	// agnesDutchPhrase represents agnes dutch phrase.
	agnesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snuffelaar"}

	// agnesFrenchPhrase represents agnes french phrase.
	agnesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "sauciflard"}

	// agnesGermanPhrase represents agnes german phrase.
	agnesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnorrz"}

	// agnesItalianPhrase represents agnes italian phrase.
	agnesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "splosh"}

	// agnesJapanesePhrase represents agnes japanese phrase.
	agnesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ブフフ"}

	// agnesKoreanPhrase represents agnes korean phrase.
	agnesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꿀꾸루"}

	// agnesLatinAmericanSpanishPhrase represents agnes latin american spanish phrase.
	agnesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "dindoink"}

	// agnesRussianPhrase represents agnes russian phrase.
	agnesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрю-хрю"}

	// agnesSimplifiedChinesePhrase represents agnes simplified chinese phrase.
	agnesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "噗呼呼"}

	// agnesSpanishPhrase represents agnes spanish phrase.
	agnesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "dindoink"}

	// agnesTraditionalChinesePhrase represents agnes traditional chinese phrase.
	agnesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "噗呼呼"}
)

var (
	// agnesPhrase represents agnes phrase.
	agnesPhrase = nook.Languages{
		language.AmericanEnglish:      agnesAmericanEnglishPhrase,
		language.CanadianFrench:       agnesCanadianFrenchPhrase,
		language.Dutch:                agnesDutchPhrase,
		language.French:               agnesFrenchPhrase,
		language.German:               agnesGermanPhrase,
		language.Italian:              agnesItalianPhrase,
		language.Japanese:             agnesJapanesePhrase,
		language.Korean:               agnesKoreanPhrase,
		language.LatinAmericanSpanish: agnesLatinAmericanSpanishPhrase,
		language.Russian:              agnesRussianPhrase,
		language.SimplifiedChinese:    agnesSimplifiedChinesePhrase,
		language.Spanish:              agnesSpanishPhrase,
		language.TraditionalChinese:   agnesTraditionalChinesePhrase}
)

var (
	// Agnes represents agnes.
	Agnes = nook.Villager{
		Character:   agnesCharacter,
		Personality: personality.BigSister,
		Phrase:      agnesPhrase}
)
