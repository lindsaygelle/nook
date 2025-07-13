package cat

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
	// felicityBirthday represents felicity birthday.
	felicityBirthday = nook.Birthday{
		Day:   30,
		Month: time.March}
)

var (
	// felicityCode represents felicity code.
	felicityCode = nook.Code{
		Value: "cat17"}
)

var (
	// felicityAmericanEnglishName represents felicity american english name.
	felicityAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Felicity"}

	// felicityCanadianFrenchName represents felicity canadian french name.
	felicityCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Maud"}

	// felicityDutchName represents felicity dutch name.
	felicityDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Felicity"}

	// felicityFrenchName represents felicity french name.
	felicityFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maud"}

	// felicityGermanName represents felicity german name.
	felicityGermanName = nook.Name{
		Language: language.German,
		Value:    "Minka"}

	// felicityItalianName represents felicity italian name.
	felicityItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Milly"}

	// felicityJapaneseName represents felicity japanese name.
	felicityJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "みかっち"}

	// felicityKoreanName represents felicity korean name.
	felicityKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "예링"}

	// felicityLatinAmericanSpanishName represents felicity latin american spanish name.
	felicityLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Micha"}

	// felicityRussianName represents felicity russian name.
	felicityRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фелисити"}

	// felicitySimplifiedChineseName represents felicity simplified chinese name.
	felicitySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "美佳"}

	// felicitySpanishName represents felicity spanish name.
	felicitySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Micha"}

	// felicityTraditionalChineseName represents felicity traditional chinese name.
	felicityTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "美佳"}
)

var (
	// felicityName represents felicity name.
	felicityName = nook.Languages{
		language.AmericanEnglish:      felicityAmericanEnglishName,
		language.CanadianFrench:       felicityCanadianFrenchName,
		language.Dutch:                felicityDutchName,
		language.French:               felicityFrenchName,
		language.German:               felicityGermanName,
		language.Italian:              felicityItalianName,
		language.Japanese:             felicityJapaneseName,
		language.Korean:               felicityKoreanName,
		language.LatinAmericanSpanish: felicityLatinAmericanSpanishName,
		language.Russian:              felicityRussianName,
		language.SimplifiedChinese:    felicitySimplifiedChineseName,
		language.Spanish:              felicitySpanishName,
		language.TraditionalChinese:   felicityTraditionalChineseName}
)

var (
	// felicityCharacter represents felicity character.
	felicityCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: felicityBirthday,
		Code:     felicityCode,
		Key:      character.Felicity,
		Gender:   gender.Female,
		Name:     felicityName,
		Special:  false}
)

var (
	// felicityAmericanEnglishPhrase represents felicity american english phrase.
	felicityAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "mimimi"}

	// felicityCanadianFrenchPhrase represents felicity canadian french phrase.
	felicityCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chatounet"}

	// felicityDutchPhrase represents felicity dutch phrase.
	felicityDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mimimi"}

	// felicityFrenchPhrase represents felicity french phrase.
	felicityFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "chatounet"}

	// felicityGermanPhrase represents felicity german phrase.
	felicityGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mauzi"}

	// felicityItalianPhrase represents felicity italian phrase.
	felicityItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "meeeow"}

	// felicityJapanesePhrase represents felicity japanese phrase.
	felicityJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ね"}

	// felicityKoreanPhrase represents felicity korean phrase.
	felicityKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우앙"}

	// felicityLatinAmericanSpanishPhrase represents felicity latin american spanish phrase.
	felicityLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "miauito"}

	// felicityRussianPhrase represents felicity russian phrase.
	felicityRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ми-ми-ми"}

	// felicitySimplifiedChinesePhrase represents felicity simplified chinese phrase.
	felicitySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哪"}

	// felicitySpanishPhrase represents felicity spanish phrase.
	felicitySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "miauito"}

	// felicityTraditionalChinesePhrase represents felicity traditional chinese phrase.
	felicityTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哪"}
)

var (
	// felicityPhrase represents felicity phrase.
	felicityPhrase = nook.Languages{
		language.AmericanEnglish:      felicityAmericanEnglishPhrase,
		language.CanadianFrench:       felicityCanadianFrenchPhrase,
		language.Dutch:                felicityDutchPhrase,
		language.French:               felicityFrenchPhrase,
		language.German:               felicityGermanPhrase,
		language.Italian:              felicityItalianPhrase,
		language.Japanese:             felicityJapanesePhrase,
		language.Korean:               felicityKoreanPhrase,
		language.LatinAmericanSpanish: felicityLatinAmericanSpanishPhrase,
		language.Russian:              felicityRussianPhrase,
		language.SimplifiedChinese:    felicitySimplifiedChinesePhrase,
		language.Spanish:              felicitySpanishPhrase,
		language.TraditionalChinese:   felicityTraditionalChinesePhrase}
)

var (
	// Felicity represents felicity.
	Felicity = nook.Villager{
		Character:   felicityCharacter,
		Personality: personality.Peppy,
		Phrase:      felicityPhrase}
)
