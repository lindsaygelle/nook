package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// pinkyBirthday represents Pinky's birthday.
var (
	// pinkyBirthday represents pinky birthday.
	pinkyBirthday = nook.Birthday{
		Day:   9,
		Month: time.September}
)

// pinkyCode represents Pinky's unique code.
var (
	// pinkyCode represents pinky code.
	pinkyCode = nook.Code{
		Value: "bea01"}
)

// Different names for Pinky in various languages.
var (
	// pinkyAmericanEnglishName represents pinky american english name.
	pinkyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pinky"}

	// pinkyCanadianFrenchName represents pinky canadian french name.
	pinkyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rosine"}

	// pinkyDutchName represents pinky dutch name.
	pinkyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pinky"}

	// pinkyFrenchName represents pinky french name.
	pinkyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rosine"}

	// pinkyGermanName represents pinky german name.
	pinkyGermanName = nook.Name{
		Language: language.German,
		Value:    "Pia"}

	// pinkyItalianName represents pinky italian name.
	pinkyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pinky"}

	// pinkyJapaneseName represents pinky japanese name.
	pinkyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タンタン"}

	// pinkyKoreanName represents pinky korean name.
	pinkyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "링링"}

	// pinkyLatinAmericanSpanishName represents pinky latin american spanish name.
	pinkyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Violeta"}

	// pinkyRussianName represents pinky russian name.
	pinkyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пинки"}

	// pinkySimplifiedChineseName represents pinky simplified chinese name.
	pinkySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "丹丹"}

	// pinkySpanishName represents pinky spanish name.
	pinkySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Violeta"}

	// pinkyTraditionalChineseName represents pinky traditional chinese name.
	pinkyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "丹丹"}
)

// pinkyName represents Pinky's name in different languages.
var (
	// pinkyName represents pinky name.
	pinkyName = nook.Languages{
		language.AmericanEnglish:      pinkyAmericanEnglishName,
		language.CanadianFrench:       pinkyCanadianFrenchName,
		language.Dutch:                pinkyDutchName,
		language.French:               pinkyFrenchName,
		language.German:               pinkyGermanName,
		language.Italian:              pinkyItalianName,
		language.Japanese:             pinkyJapaneseName,
		language.Korean:               pinkyKoreanName,
		language.LatinAmericanSpanish: pinkyLatinAmericanSpanishName,
		language.Russian:              pinkyRussianName,
		language.SimplifiedChinese:    pinkySimplifiedChineseName,
		language.Spanish:              pinkySpanishName,
		language.TraditionalChinese:   pinkyTraditionalChineseName}
)

// pinkyCharacter represents Pinky's character information.
var (
	// pinkyCharacter represents pinky character.
	pinkyCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: pinkyBirthday,
		Code:     pinkyCode,
		Key:      character.Pinky,
		Gender:   gender.Female,
		Name:     pinkyName,
		Special:  false}
)

// Different phrases spoken by Pinky in various languages.
var (
	// pinkyAmericanEnglishPhrase represents pinky american english phrase.
	pinkyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cubbie"}

	// pinkyCanadianFrenchPhrase represents pinky canadian french phrase.
	pinkyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ouah"}

	// pinkyDutchPhrase represents pinky dutch phrase.
	pinkyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "wauw"}

	// pinkyFrenchPhrase represents pinky french phrase.
	pinkyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ouah"}

	// pinkyGermanPhrase represents pinky german phrase.
	pinkyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schätzchen"}

	// pinkyItalianPhrase represents pinky italian phrase.
	pinkyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cubetto"}

	// pinkyJapanesePhrase represents pinky japanese phrase.
	pinkyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "わぉ"}

	// pinkyKoreanPhrase represents pinky korean phrase.
	pinkyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "와"}

	// pinkyLatinAmericanSpanishPhrase represents pinky latin american spanish phrase.
	pinkyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "guah"}

	// pinkyRussianPhrase represents pinky russian phrase.
	pinkyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ну и ну"}

	// pinkySimplifiedChinesePhrase represents pinky simplified chinese phrase.
	pinkySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哇喔"}

	// pinkySpanishPhrase represents pinky spanish phrase.
	pinkySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "terronillo"}

	// pinkyTraditionalChinesePhrase represents pinky traditional chinese phrase.
	pinkyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哇喔"}
)

// pinkyPhrase represents Pinky's phrases in different languages.
var (
	// pinkyPhrase represents pinky phrase.
	pinkyPhrase = nook.Languages{
		language.AmericanEnglish:      pinkyAmericanEnglishPhrase,
		language.CanadianFrench:       pinkyCanadianFrenchPhrase,
		language.Dutch:                pinkyDutchPhrase,
		language.French:               pinkyFrenchPhrase,
		language.German:               pinkyGermanPhrase,
		language.Italian:              pinkyItalianPhrase,
		language.Japanese:             pinkyJapanesePhrase,
		language.Korean:               pinkyKoreanPhrase,
		language.LatinAmericanSpanish: pinkyLatinAmericanSpanishPhrase,
		language.Russian:              pinkyRussianPhrase,
		language.SimplifiedChinese:    pinkySimplifiedChinesePhrase,
		language.Spanish:              pinkySpanishPhrase,
		language.TraditionalChinese:   pinkyTraditionalChinesePhrase}
)

// Pinky represents the character Pinky with her complete information.
var (
	// Pinky represents pinky.
	Pinky = nook.Villager{
		Character:   pinkyCharacter,
		Personality: personality.Peppy,
		Phrase:      pinkyPhrase}
)
