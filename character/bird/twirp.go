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
	// twirpBirthday represents twirp birthday.
	twirpBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// twirpCode represents twirp code.
	twirpCode = nook.Code{
		Value: ""}
)

var (
	// twirpAmericanEnglishName represents twirp american english name.
	twirpAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Twirp"}

	// twirpCanadianFrenchName represents twirp canadian french name.
	twirpCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// twirpDutchName represents twirp dutch name.
	twirpDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// twirpFrenchName represents twirp french name.
	twirpFrenchName = nook.Name{
		Language: language.French,
		Value:    "Neuneu"}

	// twirpGermanName represents twirp german name.
	twirpGermanName = nook.Name{
		Language: language.German,
		Value:    "Meik"}

	// twirpItalianName represents twirp italian name.
	twirpItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Furbecco"}

	// twirpJapaneseName represents twirp japanese name.
	twirpJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "でチュン"}

	// twirpKoreanName represents twirp korean name.
	twirpKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// twirpLatinAmericanSpanishName represents twirp latin american spanish name.
	twirpLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// twirpRussianName represents twirp russian name.
	twirpRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// twirpSimplifiedChineseName represents twirp simplified chinese name.
	twirpSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "子谈"}

	// twirpSpanishName represents twirp spanish name.
	twirpSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tomás"}

	// twirpTraditionalChineseName represents twirp traditional chinese name.
	twirpTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// twirpName represents twirp name.
	twirpName = nook.Languages{
		language.AmericanEnglish:      twirpAmericanEnglishName,
		language.CanadianFrench:       twirpCanadianFrenchName,
		language.Dutch:                twirpDutchName,
		language.French:               twirpFrenchName,
		language.German:               twirpGermanName,
		language.Italian:              twirpItalianName,
		language.Japanese:             twirpJapaneseName,
		language.Korean:               twirpKoreanName,
		language.LatinAmericanSpanish: twirpLatinAmericanSpanishName,
		language.Russian:              twirpRussianName,
		language.SimplifiedChinese:    twirpSimplifiedChineseName,
		language.Spanish:              twirpSpanishName,
		language.TraditionalChinese:   twirpTraditionalChineseName}
)

var (
	// twirpCharacter represents twirp character.
	twirpCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: twirpBirthday,
		Code:     twirpCode,
		Key:      character.Twirp,
		Gender:   gender.Male,
		Name:     twirpName,
		Special:  false}
)

var (
	// twirpAmericanEnglishPhrase represents twirp american english phrase.
	twirpAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "CHIRP"}

	// twirpCanadianFrenchPhrase represents twirp canadian french phrase.
	twirpCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// twirpDutchPhrase represents twirp dutch phrase.
	twirpDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// twirpFrenchPhrase represents twirp french phrase.
	twirpFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "oufffff"}

	// twirpGermanPhrase represents twirp german phrase.
	twirpGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tschirp"}

	// twirpItalianPhrase represents twirp italian phrase.
	twirpItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "CIRP"}

	// twirpJapanesePhrase represents twirp japanese phrase.
	twirpJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でチュン"}

	// twirpKoreanPhrase represents twirp korean phrase.
	twirpKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// twirpLatinAmericanSpanishPhrase represents twirp latin american spanish phrase.
	twirpLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// twirpRussianPhrase represents twirp russian phrase.
	twirpRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// twirpSimplifiedChinesePhrase represents twirp simplified chinese phrase.
	twirpSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咻"}

	// twirpSpanishPhrase represents twirp spanish phrase.
	twirpSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "PÍO-PÍO"}

	// twirpTraditionalChinesePhrase represents twirp traditional chinese phrase.
	twirpTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// twirpPhrase represents twirp phrase.
	twirpPhrase = nook.Languages{
		language.AmericanEnglish:      twirpAmericanEnglishPhrase,
		language.CanadianFrench:       twirpCanadianFrenchPhrase,
		language.Dutch:                twirpDutchPhrase,
		language.French:               twirpFrenchPhrase,
		language.German:               twirpGermanPhrase,
		language.Italian:              twirpItalianPhrase,
		language.Japanese:             twirpJapanesePhrase,
		language.Korean:               twirpKoreanPhrase,
		language.LatinAmericanSpanish: twirpLatinAmericanSpanishPhrase,
		language.Russian:              twirpRussianPhrase,
		language.SimplifiedChinese:    twirpSimplifiedChinesePhrase,
		language.Spanish:              twirpSpanishPhrase,
		language.TraditionalChinese:   twirpTraditionalChinesePhrase}
)

var (
	// Twirp represents twirp.
	Twirp = nook.Villager{
		Character:   twirpCharacter,
		Personality: personality.Cranky,
		Phrase:      twirpPhrase}
)
