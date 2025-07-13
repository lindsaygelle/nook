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
	// madamrosaBirthday represents madamrosa birthday.
	madamrosaBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// madamrosaCode represents madamrosa code.
	madamrosaCode = nook.Code{
		Value: ""}
)

var (
	// madamrosaAmericanEnglishName represents madamrosa american english name.
	madamrosaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Madam Rosa"}

	// madamrosaCanadianFrenchName represents madamrosa canadian french name.
	madamrosaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// madamrosaDutchName represents madamrosa dutch name.
	madamrosaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// madamrosaFrenchName represents madamrosa french name.
	madamrosaFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// madamrosaGermanName represents madamrosa german name.
	madamrosaGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// madamrosaItalianName represents madamrosa italian name.
	madamrosaItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// madamrosaJapaneseName represents madamrosa japanese name.
	madamrosaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マダムローザ"}

	// madamrosaKoreanName represents madamrosa korean name.
	madamrosaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// madamrosaLatinAmericanSpanishName represents madamrosa latin american spanish name.
	madamrosaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// madamrosaRussianName represents madamrosa russian name.
	madamrosaRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// madamrosaSimplifiedChineseName represents madamrosa simplified chinese name.
	madamrosaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// madamrosaSpanishName represents madamrosa spanish name.
	madamrosaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// madamrosaTraditionalChineseName represents madamrosa traditional chinese name.
	madamrosaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// madamrosaName represents madamrosa name.
	madamrosaName = nook.Languages{
		language.AmericanEnglish:      madamrosaAmericanEnglishName,
		language.CanadianFrench:       madamrosaCanadianFrenchName,
		language.Dutch:                madamrosaDutchName,
		language.French:               madamrosaFrenchName,
		language.German:               madamrosaGermanName,
		language.Italian:              madamrosaItalianName,
		language.Japanese:             madamrosaJapaneseName,
		language.Korean:               madamrosaKoreanName,
		language.LatinAmericanSpanish: madamrosaLatinAmericanSpanishName,
		language.Russian:              madamrosaRussianName,
		language.SimplifiedChinese:    madamrosaSimplifiedChineseName,
		language.Spanish:              madamrosaSpanishName,
		language.TraditionalChinese:   madamrosaTraditionalChineseName}
)

var (
	// madamrosaCharacter represents madamrosa character.
	madamrosaCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: madamrosaBirthday,
		Code:     madamrosaCode,
		Key:      character.MadamRosa,
		Gender:   gender.Female,
		Name:     madamrosaName,
		Special:  false}
)

var (
	// madamrosaAmericanEnglishPhrase represents madamrosa american english phrase.
	madamrosaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ほほほ"}

	// madamrosaCanadianFrenchPhrase represents madamrosa canadian french phrase.
	madamrosaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// madamrosaDutchPhrase represents madamrosa dutch phrase.
	madamrosaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// madamrosaFrenchPhrase represents madamrosa french phrase.
	madamrosaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// madamrosaGermanPhrase represents madamrosa german phrase.
	madamrosaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// madamrosaItalianPhrase represents madamrosa italian phrase.
	madamrosaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// madamrosaJapanesePhrase represents madamrosa japanese phrase.
	madamrosaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ほほほ"}

	// madamrosaKoreanPhrase represents madamrosa korean phrase.
	madamrosaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// madamrosaLatinAmericanSpanishPhrase represents madamrosa latin american spanish phrase.
	madamrosaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// madamrosaRussianPhrase represents madamrosa russian phrase.
	madamrosaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// madamrosaSimplifiedChinesePhrase represents madamrosa simplified chinese phrase.
	madamrosaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// madamrosaSpanishPhrase represents madamrosa spanish phrase.
	madamrosaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// madamrosaTraditionalChinesePhrase represents madamrosa traditional chinese phrase.
	madamrosaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// madamrosaPhrase represents madamrosa phrase.
	madamrosaPhrase = nook.Languages{
		language.AmericanEnglish:      madamrosaAmericanEnglishPhrase,
		language.CanadianFrench:       madamrosaCanadianFrenchPhrase,
		language.Dutch:                madamrosaDutchPhrase,
		language.French:               madamrosaFrenchPhrase,
		language.German:               madamrosaGermanPhrase,
		language.Italian:              madamrosaItalianPhrase,
		language.Japanese:             madamrosaJapanesePhrase,
		language.Korean:               madamrosaKoreanPhrase,
		language.LatinAmericanSpanish: madamrosaLatinAmericanSpanishPhrase,
		language.Russian:              madamrosaRussianPhrase,
		language.SimplifiedChinese:    madamrosaSimplifiedChinesePhrase,
		language.Spanish:              madamrosaSpanishPhrase,
		language.TraditionalChinese:   madamrosaTraditionalChinesePhrase}
)

var (
	// MadamRosa represents madam rosa.
	MadamRosa = nook.Villager{
		Character:   madamrosaCharacter,
		Personality: personality.Snooty,
		Phrase:      madamrosaPhrase}
)
