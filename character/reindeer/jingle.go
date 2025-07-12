package reindeer

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// jingleBirthday represents jingle birthday.
	jingleBirthday = nook.Birthday{
		Day:   24,
		Month: time.December}
)

var (
	// jingleCode represents jingle code.
	jingleCode = nook.Code{
		Value: "snt/rei"}
)

var (
	// jingleAmericanEnglishName represents jingle american english name.
	jingleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Jingle"}

	// jingleCanadianFrenchName represents jingle canadian french name.
	jingleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rodolphe"}

	// jingleDutchName represents jingle dutch name.
	jingleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Jingle"}

	// jingleFrenchName represents jingle french name.
	jingleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rodolphe"}

	// jingleGermanName represents jingle german name.
	jingleGermanName = nook.Name{
		Language: language.German,
		Value:    "Chris"}

	// jingleItalianName represents jingle italian name.
	jingleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Jingle"}

	// jingleJapaneseName represents jingle japanese name.
	jingleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジングル"}

	// jingleKoreanName represents jingle korean name.
	jingleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "루돌"}

	// jingleLatinAmericanSpanishName represents jingle latin american spanish name.
	jingleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Renato"}

	// jingleRussianName represents jingle russian name.
	jingleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джингл"}

	// jingleSimplifiedChineseName represents jingle simplified chinese name.
	jingleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "铃铃鹿"}

	// jingleSpanishName represents jingle spanish name.
	jingleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Renato"}

	// jingleTraditionalChineseName represents jingle traditional chinese name.
	jingleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鈴鈴鹿"}
)

var (
	// jingleName represents jingle name.
	jingleName = nook.Languages{
		language.AmericanEnglish:      jingleAmericanEnglishName,
		language.CanadianFrench:       jingleCanadianFrenchName,
		language.Dutch:                jingleDutchName,
		language.French:               jingleFrenchName,
		language.German:               jingleGermanName,
		language.Italian:              jingleItalianName,
		language.Japanese:             jingleJapaneseName,
		language.Korean:               jingleKoreanName,
		language.LatinAmericanSpanish: jingleLatinAmericanSpanishName,
		language.Russian:              jingleRussianName,
		language.SimplifiedChinese:    jingleSimplifiedChineseName,
		language.Spanish:              jingleSpanishName,
		language.TraditionalChinese:   jingleTraditionalChineseName}
)

var (
	// jingleCharacter represents jingle character.
	jingleCharacter = nook.Character{
		Animal:   animal.Reindeer,
		Birthday: jingleBirthday,
		Code:     jingleCode,
		Key:      character.Jingle,
		Gender:   gender.Male,
		Name:     jingleName,
		Special:  true}
)

var (
	// Jingle represents jingle.
	Jingle = nook.Resident{
		Character: jingleCharacter}
)
