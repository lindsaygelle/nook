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
	// deirdreBirthday represents deirdre birthday.
	deirdreBirthday = nook.Birthday{
		Day:   4,
		Month: time.May}
)

var (
	// deirdreCode represents deirdre code.
	deirdreCode = nook.Code{
		Value: "der04"}
)

var (
	// deirdreAmericanEnglishName represents deirdre american english name.
	deirdreAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Deirdre"}

	// deirdreCanadianFrenchName represents deirdre canadian french name.
	deirdreCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bichoune"}

	// deirdreDutchName represents deirdre dutch name.
	deirdreDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Deirdre"}

	// deirdreFrenchName represents deirdre french name.
	deirdreFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bichoune"}

	// deirdreGermanName represents deirdre german name.
	deirdreGermanName = nook.Name{
		Language: language.German,
		Value:    "Dina"}

	// deirdreItalianName represents deirdre italian name.
	deirdreItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Daria"}

	// deirdreJapaneseName represents deirdre japanese name.
	deirdreJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ナディア"}

	// deirdreKoreanName represents deirdre korean name.
	deirdreKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "나디아"}

	// deirdreLatinAmericanSpanishName represents deirdre latin american spanish name.
	deirdreLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Venada"}

	// deirdreRussianName represents deirdre russian name.
	deirdreRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дейрдре"}

	// deirdreSimplifiedChineseName represents deirdre simplified chinese name.
	deirdreSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "娣雅"}

	// deirdreSpanishName represents deirdre spanish name.
	deirdreSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Venada"}

	// deirdreTraditionalChineseName represents deirdre traditional chinese name.
	deirdreTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "娣雅"}
)

var (
	// deirdreName represents deirdre name.
	deirdreName = nook.Languages{
		language.AmericanEnglish:      deirdreAmericanEnglishName,
		language.CanadianFrench:       deirdreCanadianFrenchName,
		language.Dutch:                deirdreDutchName,
		language.French:               deirdreFrenchName,
		language.German:               deirdreGermanName,
		language.Italian:              deirdreItalianName,
		language.Japanese:             deirdreJapaneseName,
		language.Korean:               deirdreKoreanName,
		language.LatinAmericanSpanish: deirdreLatinAmericanSpanishName,
		language.Russian:              deirdreRussianName,
		language.SimplifiedChinese:    deirdreSimplifiedChineseName,
		language.Spanish:              deirdreSpanishName,
		language.TraditionalChinese:   deirdreTraditionalChineseName}
)

var (
	// deirdreCharacter represents deirdre character.
	deirdreCharacter = nook.Character{
		Animal:   animal.Deer,
		Birthday: deirdreBirthday,
		Code:     deirdreCode,
		Key:      character.Deirdre,
		Gender:   gender.Female,
		Name:     deirdreName,
		Special:  false}
)

var (
	// deirdreAmericanEnglishPhrase represents deirdre american english phrase.
	deirdreAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "whatevs"}

	// deirdreCanadianFrenchPhrase represents deirdre canadian french phrase.
	deirdreCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pan"}

	// deirdreDutchPhrase represents deirdre dutch phrase.
	deirdreDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "nou ja"}

	// deirdreFrenchPhrase represents deirdre french phrase.
	deirdreFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pan"}

	// deirdreGermanPhrase represents deirdre german phrase.
	deirdreGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tripptripp"}

	// deirdreItalianPhrase represents deirdre italian phrase.
	deirdreItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "babam"}

	// deirdreJapanesePhrase represents deirdre japanese phrase.
	deirdreJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "まいっか"}

	// deirdreKoreanPhrase represents deirdre korean phrase.
	deirdreKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "됐고"}

	// deirdreLatinAmericanSpanishPhrase represents deirdre latin american spanish phrase.
	deirdreLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "uatever"}

	// deirdreRussianPhrase represents deirdre russian phrase.
	deirdreRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "как-то так"}

	// deirdreSimplifiedChinesePhrase represents deirdre simplified chinese phrase.
	deirdreSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "随便啦"}

	// deirdreSpanishPhrase represents deirdre spanish phrase.
	deirdreSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "rumiante"}

	// deirdreTraditionalChinesePhrase represents deirdre traditional chinese phrase.
	deirdreTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "隨便啦"}
)

var (
	// deirdrePhrase represents deirdre phrase.
	deirdrePhrase = nook.Languages{
		language.AmericanEnglish:      deirdreAmericanEnglishPhrase,
		language.CanadianFrench:       deirdreCanadianFrenchPhrase,
		language.Dutch:                deirdreDutchPhrase,
		language.French:               deirdreFrenchPhrase,
		language.German:               deirdreGermanPhrase,
		language.Italian:              deirdreItalianPhrase,
		language.Japanese:             deirdreJapanesePhrase,
		language.Korean:               deirdreKoreanPhrase,
		language.LatinAmericanSpanish: deirdreLatinAmericanSpanishPhrase,
		language.Russian:              deirdreRussianPhrase,
		language.SimplifiedChinese:    deirdreSimplifiedChinesePhrase,
		language.Spanish:              deirdreSpanishPhrase,
		language.TraditionalChinese:   deirdreTraditionalChinesePhrase}
)

var (
	// Deirdre represents deirdre.
	Deirdre = nook.Villager{
		Character:   deirdreCharacter,
		Personality: personality.BigSister,
		Phrase:      deirdrePhrase}
)
