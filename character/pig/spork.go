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
	// sporkBirthday represents spork birthday.
	sporkBirthday = nook.Birthday{
		Day:   3,
		Month: time.September}
)

var (
	// sporkCode represents spork code.
	sporkCode = nook.Code{
		Value: "pig05"}
)

var (
	// sporkAmericanEnglishName represents spork american english name.
	sporkAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Spork Crackle"}

	// sporkCanadianFrenchName represents spork canadian french name.
	sporkCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Justin"}

	// sporkDutchName represents spork dutch name.
	sporkDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Crackle"}

	// sporkFrenchName represents spork french name.
	sporkFrenchName = nook.Name{
		Language: language.French,
		Value:    "Justin"}

	// sporkGermanName represents spork german name.
	sporkGermanName = nook.Name{
		Language: language.German,
		Value:    "Schwarte"}

	// sporkItalianName represents spork italian name.
	sporkItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cicciolo"}

	// sporkJapaneseName represents spork japanese name.
	sporkJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ポーク"}

	// sporkKoreanName represents spork korean name.
	sporkKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "포크"}

	// sporkLatinAmericanSpanishName represents spork latin american spanish name.
	sporkLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Espork"}

	// sporkRussianName represents spork russian name.
	sporkRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Крэкл"}

	// sporkSimplifiedChineseName represents spork simplified chinese name.
	sporkSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "肉肉"}

	// sporkSpanishName represents spork spanish name.
	sporkSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Espork"}

	// sporkTraditionalChineseName represents spork traditional chinese name.
	sporkTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "肉肉"}
)

var (
	// sporkName represents spork name.
	sporkName = nook.Languages{
		language.AmericanEnglish:      sporkAmericanEnglishName,
		language.CanadianFrench:       sporkCanadianFrenchName,
		language.Dutch:                sporkDutchName,
		language.French:               sporkFrenchName,
		language.German:               sporkGermanName,
		language.Italian:              sporkItalianName,
		language.Japanese:             sporkJapaneseName,
		language.Korean:               sporkKoreanName,
		language.LatinAmericanSpanish: sporkLatinAmericanSpanishName,
		language.Russian:              sporkRussianName,
		language.SimplifiedChinese:    sporkSimplifiedChineseName,
		language.Spanish:              sporkSpanishName,
		language.TraditionalChinese:   sporkTraditionalChineseName}
)

var (
	// sporkCharacter represents spork character.
	sporkCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: sporkBirthday,
		Code:     sporkCode,
		Key:      character.Spork,
		Gender:   gender.Male,
		Name:     sporkName,
		Special:  false}
)

var (
	// sporkAmericanEnglishPhrase represents spork american english phrase.
	sporkAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snork"}

	// sporkCanadianFrenchPhrase represents spork canadian french phrase.
	sporkCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "diantre"}

	// sporkDutchPhrase represents spork dutch phrase.
	sporkDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snurk"}

	// sporkFrenchPhrase represents spork french phrase.
	sporkFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "diantre"}

	// sporkGermanPhrase represents spork german phrase.
	sporkGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grunz"}

	// sporkItalianPhrase represents spork italian phrase.
	sporkItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "snork"}

	// sporkJapanesePhrase represents spork japanese phrase.
	sporkJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だブー"}

	// sporkKoreanPhrase represents spork korean phrase.
	sporkKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꼬르륵"}

	// sporkLatinAmericanSpanishPhrase represents spork latin american spanish phrase.
	sporkLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "pork"}

	// sporkRussianPhrase represents spork russian phrase.
	sporkRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрю"}

	// sporkSimplifiedChinesePhrase represents spork simplified chinese phrase.
	sporkSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "噗"}

	// sporkSpanishPhrase represents spork spanish phrase.
	sporkSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pork"}

	// sporkTraditionalChinesePhrase represents spork traditional chinese phrase.
	sporkTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "噗"}
)

var (
	// sporkPhrase represents spork phrase.
	sporkPhrase = nook.Languages{
		language.AmericanEnglish:      sporkAmericanEnglishPhrase,
		language.CanadianFrench:       sporkCanadianFrenchPhrase,
		language.Dutch:                sporkDutchPhrase,
		language.French:               sporkFrenchPhrase,
		language.German:               sporkGermanPhrase,
		language.Italian:              sporkItalianPhrase,
		language.Japanese:             sporkJapanesePhrase,
		language.Korean:               sporkKoreanPhrase,
		language.LatinAmericanSpanish: sporkLatinAmericanSpanishPhrase,
		language.Russian:              sporkRussianPhrase,
		language.SimplifiedChinese:    sporkSimplifiedChinesePhrase,
		language.Spanish:              sporkSpanishPhrase,
		language.TraditionalChinese:   sporkTraditionalChinesePhrase}
)

var (
	// Spork represents spork.
	Spork = nook.Villager{
		Character:   sporkCharacter,
		Personality: personality.Lazy,
		Phrase:      sporkPhrase}
)
