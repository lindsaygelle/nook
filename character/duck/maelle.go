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
	// maelleBirthday represents maelle birthday.
	maelleBirthday = nook.Birthday{
		Day:   8,
		Month: time.April}
)

var (
	// maelleCode represents maelle code.
	maelleCode = nook.Code{
		Value: "duk03"}
)

var (
	// maelleAmericanEnglishName represents maelle american english name.
	maelleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Maelle"}

	// maelleCanadianFrenchName represents maelle canadian french name.
	maelleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Maëlle"}

	// maelleDutchName represents maelle dutch name.
	maelleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Maelle"}

	// maelleFrenchName represents maelle french name.
	maelleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maëlle"}

	// maelleGermanName represents maelle german name.
	maelleGermanName = nook.Name{
		Language: language.German,
		Value:    "Sissi"}

	// maelleItalianName represents maelle italian name.
	maelleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Palmita"}

	// maelleJapaneseName represents maelle japanese name.
	maelleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アンヌ"}

	// maelleKoreanName represents maelle korean name.
	maelleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "앤"}

	// maelleLatinAmericanSpanishName represents maelle latin american spanish name.
	maelleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Patidifú"}

	// maelleRussianName represents maelle russian name.
	maelleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Маэль"}

	// maelleSimplifiedChineseName represents maelle simplified chinese name.
	maelleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "安妮"}

	// maelleSpanishName represents maelle spanish name.
	maelleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Patidifú"}

	// maelleTraditionalChineseName represents maelle traditional chinese name.
	maelleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "安妮"}
)

var (
	// maelleName represents maelle name.
	maelleName = nook.Languages{
		language.AmericanEnglish:      maelleAmericanEnglishName,
		language.CanadianFrench:       maelleCanadianFrenchName,
		language.Dutch:                maelleDutchName,
		language.French:               maelleFrenchName,
		language.German:               maelleGermanName,
		language.Italian:              maelleItalianName,
		language.Japanese:             maelleJapaneseName,
		language.Korean:               maelleKoreanName,
		language.LatinAmericanSpanish: maelleLatinAmericanSpanishName,
		language.Russian:              maelleRussianName,
		language.SimplifiedChinese:    maelleSimplifiedChineseName,
		language.Spanish:              maelleSpanishName,
		language.TraditionalChinese:   maelleTraditionalChineseName}
)

var (
	// maelleCharacter represents maelle character.
	maelleCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: maelleBirthday,
		Code:     maelleCode,
		Key:      character.Maelle,
		Gender:   gender.Female,
		Name:     maelleName,
		Special:  false}
)

var (
	// maelleAmericanEnglishPhrase represents maelle american english phrase.
	maelleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "duckling"}

	// maelleCanadianFrenchPhrase represents maelle canadian french phrase.
	maelleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "canardou"}

	// maelleDutchPhrase represents maelle dutch phrase.
	maelleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pulletje"}

	// maelleFrenchPhrase represents maelle french phrase.
	maelleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "canardou"}

	// maelleGermanPhrase represents maelle german phrase.
	maelleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "entchen"}

	// maelleItalianPhrase represents maelle italian phrase.
	maelleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "paperotto"}

	// maelleJapanesePhrase represents maelle japanese phrase.
	maelleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ふぅ"}

	// maelleKoreanPhrase represents maelle korean phrase.
	maelleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우아"}

	// maelleLatinAmericanSpanishPhrase represents maelle latin american spanish phrase.
	maelleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "plumi-piú"}

	// maelleRussianPhrase represents maelle russian phrase.
	maelleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "желторотик"}

	// maelleSimplifiedChinesePhrase represents maelle simplified chinese phrase.
	maelleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "叹"}

	// maelleSpanishPhrase represents maelle spanish phrase.
	maelleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "perla"}

	// maelleTraditionalChinesePhrase represents maelle traditional chinese phrase.
	maelleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘆"}
)

var (
	// maellePhrase represents maelle phrase.
	maellePhrase = nook.Languages{
		language.AmericanEnglish:      maelleAmericanEnglishPhrase,
		language.CanadianFrench:       maelleCanadianFrenchPhrase,
		language.Dutch:                maelleDutchPhrase,
		language.French:               maelleFrenchPhrase,
		language.German:               maelleGermanPhrase,
		language.Italian:              maelleItalianPhrase,
		language.Japanese:             maelleJapanesePhrase,
		language.Korean:               maelleKoreanPhrase,
		language.LatinAmericanSpanish: maelleLatinAmericanSpanishPhrase,
		language.Russian:              maelleRussianPhrase,
		language.SimplifiedChinese:    maelleSimplifiedChinesePhrase,
		language.Spanish:              maelleSpanishPhrase,
		language.TraditionalChinese:   maelleTraditionalChinesePhrase}
)

var (
	// Maelle represents maelle.
	Maelle = nook.Villager{
		Character:   maelleCharacter,
		Personality: personality.Snooty,
		Phrase:      maellePhrase}
)
