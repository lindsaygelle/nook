package deer

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
	// bamBirthday represents bam birthday.
	bamBirthday = nook.Birthday{
		Day:   7,
		Month: time.November}
)

var (
	// bamCode represents bam code.
	bamCode = nook.Code{
		Value: "der01"}
)

var (
	// bamAmericanEnglishName represents bam american english name.
	bamAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bam"}

	// bamCanadianFrenchName represents bam canadian french name.
	bamCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Nacer"}

	// bamDutchName represents bam dutch name.
	bamDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bam"}

	// bamFrenchName represents bam french name.
	bamFrenchName = nook.Name{
		Language: language.French,
		Value:    "Nacer"}

	// bamGermanName represents bam german name.
	bamGermanName = nook.Name{
		Language: language.German,
		Value:    "Benjamin"}

	// bamItalianName represents bam italian name.
	bamItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cornelio"}

	// bamJapaneseName represents bam japanese name.
	bamJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タケル"}

	// bamKoreanName represents bam korean name.
	bamKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "록키"}

	// bamLatinAmericanSpanishName represents bam latin american spanish name.
	bamLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cornelio"}

	// bamRussianName represents bam russian name.
	bamRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бам"}

	// bamSimplifiedChineseName represents bam simplified chinese name.
	bamSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小健"}

	// bamSpanishName represents bam spanish name.
	bamSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cornelio"}

	// bamTraditionalChineseName represents bam traditional chinese name.
	bamTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小健"}
)

var (
	// bamName represents bam name.
	bamName = nook.Languages{
		language.AmericanEnglish:      bamAmericanEnglishName,
		language.CanadianFrench:       bamCanadianFrenchName,
		language.Dutch:                bamDutchName,
		language.French:               bamFrenchName,
		language.German:               bamGermanName,
		language.Italian:              bamItalianName,
		language.Japanese:             bamJapaneseName,
		language.Korean:               bamKoreanName,
		language.LatinAmericanSpanish: bamLatinAmericanSpanishName,
		language.Russian:              bamRussianName,
		language.SimplifiedChinese:    bamSimplifiedChineseName,
		language.Spanish:              bamSpanishName,
		language.TraditionalChinese:   bamTraditionalChineseName}
)

var (
	// bamCharacter represents bam character.
	bamCharacter = nook.Character{
		Animal:   animal.Deer,
		Birthday: bamBirthday,
		Code:     bamCode,
		Key:      character.Bam,
		Gender:   gender.Male,
		Name:     bamName,
		Special:  false}
)

var (
	// bamAmericanEnglishPhrase represents bam american english phrase.
	bamAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "boosh"}

	// bamCanadianFrenchPhrase represents bam canadian french phrase.
	bamCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "têtedebois"}

	// bamDutchPhrase represents bam dutch phrase.
	bamDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "gaffel"}

	// bamFrenchPhrase represents bam french phrase.
	bamFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "têtedebois"}

	// bamGermanPhrase represents bam german phrase.
	bamGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "weihweih"}

	// bamItalianPhrase represents bam italian phrase.
	bamItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hop hop"}

	// bamJapanesePhrase represents bam japanese phrase.
	bamJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ナラね"}

	// bamKoreanPhrase represents bam korean phrase.
	bamKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "슉슉"}

	// bamLatinAmericanSpanishPhrase represents bam latin american spanish phrase.
	bamLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "caplúúú"}

	// bamRussianPhrase represents bam russian phrase.
	bamRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "скок-скок"}

	// bamSimplifiedChinesePhrase represents bam simplified chinese phrase.
	bamSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "梅花"}

	// bamSpanishPhrase represents bam spanish phrase.
	bamSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "galilla"}

	// bamTraditionalChinesePhrase represents bam traditional chinese phrase.
	bamTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "梅花"}
)

var (
	// bamPhrase represents bam phrase.
	bamPhrase = nook.Languages{
		language.AmericanEnglish:      bamAmericanEnglishPhrase,
		language.CanadianFrench:       bamCanadianFrenchPhrase,
		language.Dutch:                bamDutchPhrase,
		language.French:               bamFrenchPhrase,
		language.German:               bamGermanPhrase,
		language.Italian:              bamItalianPhrase,
		language.Japanese:             bamJapanesePhrase,
		language.Korean:               bamKoreanPhrase,
		language.LatinAmericanSpanish: bamLatinAmericanSpanishPhrase,
		language.Russian:              bamRussianPhrase,
		language.SimplifiedChinese:    bamSimplifiedChinesePhrase,
		language.Spanish:              bamSpanishPhrase,
		language.TraditionalChinese:   bamTraditionalChinesePhrase}
)

var (
	// Bam represents bam.
	Bam = nook.Villager{
		Character:   bamCharacter,
		Personality: personality.Jock,
		Phrase:      bamPhrase}
)
