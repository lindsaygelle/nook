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
	// borisBirthday represents boris birthday.
	borisBirthday = nook.Birthday{
		Day:   6,
		Month: time.November}
)

var (
	// borisCode represents boris code.
	borisCode = nook.Code{
		Value: "pig09"}
)

var (
	// borisAmericanEnglishName represents boris american english name.
	borisAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Boris"}

	// borisCanadianFrenchName represents boris canadian french name.
	borisCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Boris"}

	// borisDutchName represents boris dutch name.
	borisDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Boris"}

	// borisFrenchName represents boris french name.
	borisFrenchName = nook.Name{
		Language: language.French,
		Value:    "Boris"}

	// borisGermanName represents boris german name.
	borisGermanName = nook.Name{
		Language: language.German,
		Value:    "Bolle"}

	// borisItalianName represents boris italian name.
	borisItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Boris"}

	// borisJapaneseName represents boris japanese name.
	borisJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ダリー"}

	// borisKoreanName represents boris korean name.
	borisKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "보리"}

	// borisLatinAmericanSpanishName represents boris latin american spanish name.
	borisLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Boris"}

	// borisRussianName represents boris russian name.
	borisRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Борис"}

	// borisSimplifiedChineseName represents boris simplified chinese name.
	borisSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "达利"}

	// borisSpanishName represents boris spanish name.
	borisSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Boris"}

	// borisTraditionalChineseName represents boris traditional chinese name.
	borisTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "達利"}
)

var (
	// borisName represents boris name.
	borisName = nook.Languages{
		language.AmericanEnglish:      borisAmericanEnglishName,
		language.CanadianFrench:       borisCanadianFrenchName,
		language.Dutch:                borisDutchName,
		language.French:               borisFrenchName,
		language.German:               borisGermanName,
		language.Italian:              borisItalianName,
		language.Japanese:             borisJapaneseName,
		language.Korean:               borisKoreanName,
		language.LatinAmericanSpanish: borisLatinAmericanSpanishName,
		language.Russian:              borisRussianName,
		language.SimplifiedChinese:    borisSimplifiedChineseName,
		language.Spanish:              borisSpanishName,
		language.TraditionalChinese:   borisTraditionalChineseName}
)

var (
	// borisCharacter represents boris character.
	borisCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: borisBirthday,
		Code:     borisCode,
		Key:      character.Boris,
		Gender:   gender.Male,
		Name:     borisName,
		Special:  false}
)

var (
	// borisAmericanEnglishPhrase represents boris american english phrase.
	borisAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "schnort"}

	// borisCanadianFrenchPhrase represents boris canadian french phrase.
	borisCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "minot"}

	// borisDutchPhrase represents boris dutch phrase.
	borisDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "oink"}

	// borisFrenchPhrase represents boris french phrase.
	borisFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "minot"}

	// borisGermanPhrase represents boris german phrase.
	borisGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quiek"}

	// borisItalianPhrase represents boris italian phrase.
	borisItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sgrunf"}

	// borisJapanesePhrase represents boris japanese phrase.
	borisJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ブヒッ"}

	// borisKoreanPhrase represents boris korean phrase.
	borisKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쌀"}

	// borisLatinAmericanSpanishPhrase represents boris latin american spanish phrase.
	borisLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "oink"}

	// borisRussianPhrase represents boris russian phrase.
	borisRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрю-хряк"}

	// borisSimplifiedChinesePhrase represents boris simplified chinese phrase.
	borisSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "噗嘻"}

	// borisSpanishPhrase represents boris spanish phrase.
	borisSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "oink"}

	// borisTraditionalChinesePhrase represents boris traditional chinese phrase.
	borisTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "噗嘻"}
)

var (
	// borisPhrase represents boris phrase.
	borisPhrase = nook.Languages{
		language.AmericanEnglish:      borisAmericanEnglishPhrase,
		language.CanadianFrench:       borisCanadianFrenchPhrase,
		language.Dutch:                borisDutchPhrase,
		language.French:               borisFrenchPhrase,
		language.German:               borisGermanPhrase,
		language.Italian:              borisItalianPhrase,
		language.Japanese:             borisJapanesePhrase,
		language.Korean:               borisKoreanPhrase,
		language.LatinAmericanSpanish: borisLatinAmericanSpanishPhrase,
		language.Russian:              borisRussianPhrase,
		language.SimplifiedChinese:    borisSimplifiedChinesePhrase,
		language.Spanish:              borisSpanishPhrase,
		language.TraditionalChinese:   borisTraditionalChinesePhrase}
)

var (
	// Boris represents boris.
	Boris = nook.Villager{
		Character:   borisCharacter,
		Personality: personality.Cranky,
		Phrase:      borisPhrase}
)
