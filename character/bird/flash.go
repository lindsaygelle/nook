package bird

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
	// flashBirthday represents flash birthday.
	flashBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// flashCode represents flash code.
	flashCode = nook.Code{
		Value: ""}
)

var (
	// flashAmericanEnglishName represents flash american english name.
	flashAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Flash"}

	// flashCanadianFrenchName represents flash canadian french name.
	flashCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// flashDutchName represents flash dutch name.
	flashDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// flashFrenchName represents flash french name.
	flashFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maurice"}

	// flashGermanName represents flash german name.
	flashGermanName = nook.Name{
		Language: language.German,
		Value:    "Flo"}

	// flashItalianName represents flash italian name.
	flashItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Baleno"}

	// flashJapaneseName represents flash japanese name.
	flashJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "みちる"}

	// flashKoreanName represents flash korean name.
	flashKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// flashLatinAmericanSpanishName represents flash latin american spanish name.
	flashLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// flashRussianName represents flash russian name.
	flashRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// flashSimplifiedChineseName represents flash simplified chinese name.
	flashSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// flashSpanishName represents flash spanish name.
	flashSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mauro"}

	// flashTraditionalChineseName represents flash traditional chinese name.
	flashTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// flashName represents flash name.
	flashName = nook.Languages{
		language.AmericanEnglish:      flashAmericanEnglishName,
		language.CanadianFrench:       flashCanadianFrenchName,
		language.Dutch:                flashDutchName,
		language.French:               flashFrenchName,
		language.German:               flashGermanName,
		language.Italian:              flashItalianName,
		language.Japanese:             flashJapaneseName,
		language.Korean:               flashKoreanName,
		language.LatinAmericanSpanish: flashLatinAmericanSpanishName,
		language.Russian:              flashRussianName,
		language.SimplifiedChinese:    flashSimplifiedChineseName,
		language.Spanish:              flashSpanishName,
		language.TraditionalChinese:   flashTraditionalChineseName}
)

var (
	// flashCharacter represents flash character.
	flashCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: flashBirthday,
		Code:     flashCode,
		Key:      character.Flash,
		Gender:   gender.Male,
		Name:     flashName,
		Special:  false}
)

var (
	// flashAmericanEnglishPhrase represents flash american english phrase.
	flashAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "babe"}

	// flashCanadianFrenchPhrase represents flash canadian french phrase.
	flashCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// flashDutchPhrase represents flash dutch phrase.
	flashDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// flashFrenchPhrase represents flash french phrase.
	flashFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "cuicui"}

	// flashGermanPhrase represents flash german phrase.
	flashGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schätzchen"}

	// flashItalianPhrase represents flash italian phrase.
	flashItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "flash"}

	// flashJapanesePhrase represents flash japanese phrase.
	flashJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "チルチル"}

	// flashKoreanPhrase represents flash korean phrase.
	flashKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// flashLatinAmericanSpanishPhrase represents flash latin american spanish phrase.
	flashLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// flashRussianPhrase represents flash russian phrase.
	flashRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// flashSimplifiedChinesePhrase represents flash simplified chinese phrase.
	flashSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// flashSpanishPhrase represents flash spanish phrase.
	flashSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cari"}

	// flashTraditionalChinesePhrase represents flash traditional chinese phrase.
	flashTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// flashPhrase represents flash phrase.
	flashPhrase = nook.Languages{
		language.AmericanEnglish:      flashAmericanEnglishPhrase,
		language.CanadianFrench:       flashCanadianFrenchPhrase,
		language.Dutch:                flashDutchPhrase,
		language.French:               flashFrenchPhrase,
		language.German:               flashGermanPhrase,
		language.Italian:              flashItalianPhrase,
		language.Japanese:             flashJapanesePhrase,
		language.Korean:               flashKoreanPhrase,
		language.LatinAmericanSpanish: flashLatinAmericanSpanishPhrase,
		language.Russian:              flashRussianPhrase,
		language.SimplifiedChinese:    flashSimplifiedChinesePhrase,
		language.Spanish:              flashSpanishPhrase,
		language.TraditionalChinese:   flashTraditionalChinesePhrase}
)

var (
	// Flash represents flash.
	Flash = nook.Villager{
		Character:   flashCharacter,
		Personality: personality.Cranky,
		Phrase:      flashPhrase}
)
