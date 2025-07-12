package horse

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
	// savannahBirthday represents savannah birthday.
	savannahBirthday = nook.Birthday{
		Day:   25,
		Month: time.January}
)

var (
	// savannahCode represents savannah code.
	savannahCode = nook.Code{
		Value: "hrs02"}
)

var (
	// savannahAmericanEnglishName represents savannah american english name.
	savannahAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Savannah"}

	// savannahCanadianFrenchName represents savannah canadian french name.
	savannahCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Savana"}

	// savannahDutchName represents savannah dutch name.
	savannahDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Savannah"}

	// savannahFrenchName represents savannah french name.
	savannahFrenchName = nook.Name{
		Language: language.French,
		Value:    "Savana"}

	// savannahGermanName represents savannah german name.
	savannahGermanName = nook.Name{
		Language: language.German,
		Value:    "Zara"}

	// savannahItalianName represents savannah italian name.
	savannahItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Savannah"}

	// savannahJapaneseName represents savannah japanese name.
	savannahJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サバンナ"}

	// savannahKoreanName represents savannah korean name.
	savannahKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "사반나"}

	// savannahLatinAmericanSpanishName represents savannah latin american spanish name.
	savannahLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sabana"}

	// savannahRussianName represents savannah russian name.
	savannahRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Саванна"}

	// savannahSimplifiedChineseName represents savannah simplified chinese name.
	savannahSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "草斑娜"}

	// savannahSpanishName represents savannah spanish name.
	savannahSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sabana"}

	// savannahTraditionalChineseName represents savannah traditional chinese name.
	savannahTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "草斑娜"}
)

var (
	// savannahName represents savannah name.
	savannahName = nook.Languages{
		language.AmericanEnglish:      savannahAmericanEnglishName,
		language.CanadianFrench:       savannahCanadianFrenchName,
		language.Dutch:                savannahDutchName,
		language.French:               savannahFrenchName,
		language.German:               savannahGermanName,
		language.Italian:              savannahItalianName,
		language.Japanese:             savannahJapaneseName,
		language.Korean:               savannahKoreanName,
		language.LatinAmericanSpanish: savannahLatinAmericanSpanishName,
		language.Russian:              savannahRussianName,
		language.SimplifiedChinese:    savannahSimplifiedChineseName,
		language.Spanish:              savannahSpanishName,
		language.TraditionalChinese:   savannahTraditionalChineseName}
)

var (
	// savannahCharacter represents savannah character.
	savannahCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: savannahBirthday,
		Code:     savannahCode,
		Key:      character.Savannah,
		Gender:   gender.Female,
		Name:     savannahName,
		Special:  false}
)

var (
	// savannahAmericanEnglishPhrase represents savannah american english phrase.
	savannahAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "y'all"}

	// savannahCanadianFrenchPhrase represents savannah canadian french phrase.
	savannahCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "wimbowé"}

	// savannahDutchPhrase represents savannah dutch phrase.
	savannahDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pony"}

	// savannahFrenchPhrase represents savannah french phrase.
	savannahFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "wimbowé"}

	// savannahGermanPhrase represents savannah german phrase.
	savannahGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "wichichi"}

	// savannahItalianPhrase represents savannah italian phrase.
	savannahItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hip hip"}

	// savannahJapanesePhrase represents savannah japanese phrase.
	savannahJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ってば"}

	// savannahKoreanPhrase represents savannah korean phrase.
	savannahKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "맞아요"}

	// savannahLatinAmericanSpanishPhrase represents savannah latin american spanish phrase.
	savannahLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "azuquita"}

	// savannahRussianPhrase represents savannah russian phrase.
	savannahRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "и все"}

	// savannahSimplifiedChinesePhrase represents savannah simplified chinese phrase.
	savannahSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "说到"}

	// savannahSpanishPhrase represents savannah spanish phrase.
	savannahSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "azucarillo"}

	// savannahTraditionalChinesePhrase represents savannah traditional chinese phrase.
	savannahTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "說到"}
)

var (
	// savannahPhrase represents savannah phrase.
	savannahPhrase = nook.Languages{
		language.AmericanEnglish:      savannahAmericanEnglishPhrase,
		language.CanadianFrench:       savannahCanadianFrenchPhrase,
		language.Dutch:                savannahDutchPhrase,
		language.French:               savannahFrenchPhrase,
		language.German:               savannahGermanPhrase,
		language.Italian:              savannahItalianPhrase,
		language.Japanese:             savannahJapanesePhrase,
		language.Korean:               savannahKoreanPhrase,
		language.LatinAmericanSpanish: savannahLatinAmericanSpanishPhrase,
		language.Russian:              savannahRussianPhrase,
		language.SimplifiedChinese:    savannahSimplifiedChinesePhrase,
		language.Spanish:              savannahSpanishPhrase,
		language.TraditionalChinese:   savannahTraditionalChinesePhrase}
)

var (
	// Savannah represents savannah.
	Savannah = nook.Villager{
		Character:   savannahCharacter,
		Personality: personality.Normal,
		Phrase:      savannahPhrase}
)
