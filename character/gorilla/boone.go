package gorilla

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
	// booneBirthday represents boone birthday.
	booneBirthday = nook.Birthday{
		Day:   12,
		Month: time.September}
)

var (
	// booneCode represents boone code.
	booneCode = nook.Code{
		Value: "gor02"}
)

var (
	// booneAmericanEnglishName represents boone american english name.
	booneAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Boone"}

	// booneCanadianFrenchName represents boone canadian french name.
	booneCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Babouin"}

	// booneDutchName represents boone dutch name.
	booneDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Boone"}

	// booneFrenchName represents boone french name.
	booneFrenchName = nook.Name{
		Language: language.French,
		Value:    "Babouin"}

	// booneGermanName represents boone german name.
	booneGermanName = nook.Name{
		Language: language.German,
		Value:    "Kong"}

	// booneItalianName represents boone italian name.
	booneItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Babu"}

	// booneJapaneseName represents boone japanese name.
	booneJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "まんたろう"}

	// booneKoreanName represents boone korean name.
	booneKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "만복이"}

	// booneLatinAmericanSpanishName represents boone latin american spanish name.
	booneLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Babú"}

	// booneRussianName represents boone russian name.
	booneRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Буин"}

	// booneSimplifiedChineseName represents boone simplified chinese name.
	booneSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "万泰"}

	// booneSpanishName represents boone spanish name.
	booneSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Babú"}

	// booneTraditionalChineseName represents boone traditional chinese name.
	booneTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "萬泰"}
)

var (
	// booneName represents boone name.
	booneName = nook.Languages{
		language.AmericanEnglish:      booneAmericanEnglishName,
		language.CanadianFrench:       booneCanadianFrenchName,
		language.Dutch:                booneDutchName,
		language.French:               booneFrenchName,
		language.German:               booneGermanName,
		language.Italian:              booneItalianName,
		language.Japanese:             booneJapaneseName,
		language.Korean:               booneKoreanName,
		language.LatinAmericanSpanish: booneLatinAmericanSpanishName,
		language.Russian:              booneRussianName,
		language.SimplifiedChinese:    booneSimplifiedChineseName,
		language.Spanish:              booneSpanishName,
		language.TraditionalChinese:   booneTraditionalChineseName}
)

var (
	// booneCharacter represents boone character.
	booneCharacter = nook.Character{
		Animal:   animal.Gorilla,
		Birthday: booneBirthday,
		Code:     booneCode,
		Key:      character.Boone,
		Gender:   gender.Male,
		Name:     booneName,
		Special:  false}
)

var (
	// booneAmericanEnglishPhrase represents boone american english phrase.
	booneAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "baboom"}

	// booneCanadianFrenchPhrase represents boone canadian french phrase.
	booneCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ma noix"}

	// booneDutchPhrase represents boone dutch phrase.
	booneDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "bavi-AAN"}

	// booneFrenchPhrase represents boone french phrase.
	booneFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ma noix"}

	// booneGermanPhrase represents boone german phrase.
	booneGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "ugh"}

	// booneItalianPhrase represents boone italian phrase.
	booneItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "baboom"}

	// booneJapanesePhrase represents boone japanese phrase.
	booneJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウルトラ"}

	// booneKoreanPhrase represents boone korean phrase.
	booneKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "야임마"}

	// booneLatinAmericanSpanishPhrase represents boone latin american spanish phrase.
	booneLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "angaua"}

	// booneRussianPhrase represents boone russian phrase.
	booneRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бабум"}

	// booneSimplifiedChinesePhrase represents boone simplified chinese phrase.
	booneSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "这家伙"}

	// booneSpanishPhrase represents boone spanish phrase.
	booneSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "angaua"}

	// booneTraditionalChinesePhrase represents boone traditional chinese phrase.
	booneTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "超級"}
)

var (
	// boonePhrase represents boone phrase.
	boonePhrase = nook.Languages{
		language.AmericanEnglish:      booneAmericanEnglishPhrase,
		language.CanadianFrench:       booneCanadianFrenchPhrase,
		language.Dutch:                booneDutchPhrase,
		language.French:               booneFrenchPhrase,
		language.German:               booneGermanPhrase,
		language.Italian:              booneItalianPhrase,
		language.Japanese:             booneJapanesePhrase,
		language.Korean:               booneKoreanPhrase,
		language.LatinAmericanSpanish: booneLatinAmericanSpanishPhrase,
		language.Russian:              booneRussianPhrase,
		language.SimplifiedChinese:    booneSimplifiedChinesePhrase,
		language.Spanish:              booneSpanishPhrase,
		language.TraditionalChinese:   booneTraditionalChinesePhrase}
)

var (
	// Boone represents boone.
	Boone = nook.Villager{
		Character:   booneCharacter,
		Personality: personality.Jock,
		Phrase:      boonePhrase}
)
