package eagle

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
	// frankBirthday represents frank birthday.
	frankBirthday = nook.Birthday{
		Day:   30,
		Month: time.July}
)

var (
	// frankCode represents frank code.
	frankCode = nook.Code{
		Value: "pbr06"}
)

var (
	// frankAmericanEnglishName represents frank american english name.
	frankAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Frank"}

	// frankCanadianFrenchName represents frank canadian french name.
	frankCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Greggae"}

	// frankDutchName represents frank dutch name.
	frankDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Frank"}

	// frankFrenchName represents frank french name.
	frankFrenchName = nook.Name{
		Language: language.French,
		Value:    "Greggae"}

	// frankGermanName represents frank german name.
	frankGermanName = nook.Name{
		Language: language.German,
		Value:    "Arthur"}

	// frankItalianName represents frank italian name.
	frankItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Frank"}

	// frankJapaneseName represents frank japanese name.
	frankJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハルク"}

	// frankKoreanName represents frank korean name.
	frankKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "헐크"}

	// frankLatinAmericanSpanishName represents frank latin american spanish name.
	frankLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Aquilino"}

	// frankRussianName represents frank russian name.
	frankRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Франк"}

	// frankSimplifiedChineseName represents frank simplified chinese name.
	frankSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "浩克"}

	// frankSpanishName represents frank spanish name.
	frankSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Aquilino"}

	// frankTraditionalChineseName represents frank traditional chinese name.
	frankTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "浩克"}
)

var (
	// frankName represents frank name.
	frankName = nook.Languages{
		language.AmericanEnglish:      frankAmericanEnglishName,
		language.CanadianFrench:       frankCanadianFrenchName,
		language.Dutch:                frankDutchName,
		language.French:               frankFrenchName,
		language.German:               frankGermanName,
		language.Italian:              frankItalianName,
		language.Japanese:             frankJapaneseName,
		language.Korean:               frankKoreanName,
		language.LatinAmericanSpanish: frankLatinAmericanSpanishName,
		language.Russian:              frankRussianName,
		language.SimplifiedChinese:    frankSimplifiedChineseName,
		language.Spanish:              frankSpanishName,
		language.TraditionalChinese:   frankTraditionalChineseName}
)

var (
	// frankCharacter represents frank character.
	frankCharacter = nook.Character{
		Animal:   animal.Eagle,
		Birthday: frankBirthday,
		Code:     frankCode,
		Key:      character.Frank,
		Gender:   gender.Male,
		Name:     frankName,
		Special:  false}
)

var (
	// frankAmericanEnglishPhrase represents frank american english phrase.
	frankAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "crushy"}

	// frankCanadianFrenchPhrase represents frank canadian french phrase.
	frankCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "maaaarre"}

	// frankDutchPhrase represents frank dutch phrase.
	frankDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pletter"}

	// frankFrenchPhrase represents frank french phrase.
	frankFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "maaaarre"}

	// frankGermanPhrase represents frank german phrase.
	frankGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kriiieck"}

	// frankItalianPhrase represents frank italian phrase.
	frankItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "rapaciao"}

	// frankJapanesePhrase represents frank japanese phrase.
	frankJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ダンケ"}

	// frankKoreanPhrase represents frank korean phrase.
	frankKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "크헐"}

	// frankLatinAmericanSpanishPhrase represents frank latin american spanish phrase.
	frankLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "vulanico"}

	// frankRussianPhrase represents frank russian phrase.
	frankRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "красава"}

	// frankSimplifiedChinesePhrase represents frank simplified chinese phrase.
	frankSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "多谢"}

	// frankSpanishPhrase represents frank spanish phrase.
	frankSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "vulanico"}

	// frankTraditionalChinesePhrase represents frank traditional chinese phrase.
	frankTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "多謝"}
)

var (
	// frankPhrase represents frank phrase.
	frankPhrase = nook.Languages{
		language.AmericanEnglish:      frankAmericanEnglishPhrase,
		language.CanadianFrench:       frankCanadianFrenchPhrase,
		language.Dutch:                frankDutchPhrase,
		language.French:               frankFrenchPhrase,
		language.German:               frankGermanPhrase,
		language.Italian:              frankItalianPhrase,
		language.Japanese:             frankJapanesePhrase,
		language.Korean:               frankKoreanPhrase,
		language.LatinAmericanSpanish: frankLatinAmericanSpanishPhrase,
		language.Russian:              frankRussianPhrase,
		language.SimplifiedChinese:    frankSimplifiedChinesePhrase,
		language.Spanish:              frankSpanishPhrase,
		language.TraditionalChinese:   frankTraditionalChinesePhrase}
)

var (
	// Frank represents frank.
	Frank = nook.Villager{
		Character:   frankCharacter,
		Personality: personality.Cranky,
		Phrase:      frankPhrase}
)
