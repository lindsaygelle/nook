package bearcub

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
	// judyBirthday represents judy birthday.
	judyBirthday = nook.Birthday{
		Day:   10,
		Month: time.March}
)

var (
	// judyCode represents judy code.
	judyCode = nook.Code{
		Value: "cbr19"}
)

var (
	// judyAmericanEnglishName represents judy american english name.
	judyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Judy"}

	// judyCanadianFrenchName represents judy canadian french name.
	judyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Laura"}

	// judyDutchName represents judy dutch name.
	judyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Judy"}

	// judyFrenchName represents judy french name.
	judyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Laura"}

	// judyGermanName represents judy german name.
	judyGermanName = nook.Name{
		Language: language.German,
		Value:    "Misuzu"}

	// judyItalianName represents judy italian name.
	judyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Misuzu"}

	// judyJapaneseName represents judy japanese name.
	judyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "みすず"}

	// judyKoreanName represents judy korean name.
	judyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "미애"}

	// judyLatinAmericanSpanishName represents judy latin american spanish name.
	judyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rosezna"}

	// judyRussianName represents judy russian name.
	judyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джуди"}

	// judySimplifiedChineseName represents judy simplified chinese name.
	judySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "美玲"}

	// judySpanishName represents judy spanish name.
	judySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rosezna"}

	// judyTraditionalChineseName represents judy traditional chinese name.
	judyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "美玲"}
)

var (
	// judyName represents judy name.
	judyName = nook.Languages{
		language.AmericanEnglish:      judyAmericanEnglishName,
		language.CanadianFrench:       judyCanadianFrenchName,
		language.Dutch:                judyDutchName,
		language.French:               judyFrenchName,
		language.German:               judyGermanName,
		language.Italian:              judyItalianName,
		language.Japanese:             judyJapaneseName,
		language.Korean:               judyKoreanName,
		language.LatinAmericanSpanish: judyLatinAmericanSpanishName,
		language.Russian:              judyRussianName,
		language.SimplifiedChinese:    judySimplifiedChineseName,
		language.Spanish:              judySpanishName,
		language.TraditionalChinese:   judyTraditionalChineseName}
)

var (
	// judyCharacter represents judy character.
	judyCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: judyBirthday,
		Code:     judyCode,
		Key:      character.Judy,
		Gender:   gender.Female,
		Name:     judyName,
		Special:  false}
)

var (
	// judyAmericanEnglishPhrase represents judy american english phrase.
	judyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "myohmy"}

	// judyCanadianFrenchPhrase represents judy canadian french phrase.
	judyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "oh là là"}

	// judyDutchPhrase represents judy dutch phrase.
	judyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "nee maar"}

	// judyFrenchPhrase represents judy french phrase.
	judyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "oh là là"}

	// judyGermanPhrase represents judy german phrase.
	judyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "träum"}

	// judyItalianPhrase represents judy italian phrase.
	judyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "misu misu"}

	// judyJapanesePhrase represents judy japanese phrase.
	judyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "あらら"}

	// judyKoreanPhrase represents judy korean phrase.
	judyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어머머"}

	// judyLatinAmericanSpanishPhrase represents judy latin american spanish phrase.
	judyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "uyuyuy"}

	// judyRussianPhrase represents judy russian phrase.
	judyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "охохонюшки"}

	// judySimplifiedChinesePhrase represents judy simplified chinese phrase.
	judySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哎呀"}

	// judySpanishPhrase represents judy spanish phrase.
	judySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "qué cosas"}

	// judyTraditionalChinesePhrase represents judy traditional chinese phrase.
	judyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哎呀"}
)

var (
	// judyPhrase represents judy phrase.
	judyPhrase = nook.Languages{
		language.AmericanEnglish:      judyAmericanEnglishPhrase,
		language.CanadianFrench:       judyCanadianFrenchPhrase,
		language.Dutch:                judyDutchPhrase,
		language.French:               judyFrenchPhrase,
		language.German:               judyGermanPhrase,
		language.Italian:              judyItalianPhrase,
		language.Japanese:             judyJapanesePhrase,
		language.Korean:               judyKoreanPhrase,
		language.LatinAmericanSpanish: judyLatinAmericanSpanishPhrase,
		language.Russian:              judyRussianPhrase,
		language.SimplifiedChinese:    judySimplifiedChinesePhrase,
		language.Spanish:              judySpanishPhrase,
		language.TraditionalChinese:   judyTraditionalChinesePhrase}
)

var (
	// Judy represents judy.
	Judy = nook.Villager{
		Character:   judyCharacter,
		Personality: personality.Snooty,
		Phrase:      judyPhrase}
)
