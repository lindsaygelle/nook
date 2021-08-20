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
	jingleBirthday = nook.Birthday{
		Day:   24,
		Month: time.December}
)

var (
	jingleCode = nook.Code{
		Value: "snt/rei"}
)

var (
	jingleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Jingle"}

	jingleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rodolphe"}

	jingleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Jingle"}

	jingleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rodolphe"}

	jingleGermanName = nook.Name{
		Language: language.German,
		Value:    "Chris"}

	jingleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Jingle"}

	jingleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジングル"}

	jingleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "루돌"}

	jingleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Renato"}

	jingleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джингл"}

	jingleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "铃铃鹿"}

	jingleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Renato"}

	jingleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鈴鈴鹿"}
)

var (
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
	Jingle = nook.Resident{
		Character: jingleCharacter}
)
