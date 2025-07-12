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
	// jayBirthday represents jay birthday.
	jayBirthday = nook.Birthday{
		Day:   17,
		Month: time.July}
)

var (
	// jayCode represents jay code.
	jayCode = nook.Code{
		Value: "brd00"}
)

var (
	// jayAmericanEnglishName represents jay american english name.
	jayAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Jay"}

	// jayCanadianFrenchName represents jay canadian french name.
	jayCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gérard"}

	// jayDutchName represents jay dutch name.
	jayDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Jay"}

	// jayFrenchName represents jay french name.
	jayFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gérard"}

	// jayGermanName represents jay german name.
	jayGermanName = nook.Name{
		Language: language.German,
		Value:    "Jan"}

	// jayItalianName represents jay italian name.
	jayItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Jet"}

	// jayJapaneseName represents jay japanese name.
	jayJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ツバクロ"}

	// jayKoreanName represents jay korean name.
	jayKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "참돌이"}

	// jayLatinAmericanSpanishName represents jay latin american spanish name.
	jayLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jairo"}

	// jayRussianName represents jay russian name.
	jayRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джей"}

	// jaySimplifiedChineseName represents jay simplified chinese name.
	jaySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿燕"}

	// jaySpanishName represents jay spanish name.
	jaySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jota"}

	// jayTraditionalChineseName represents jay traditional chinese name.
	jayTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿燕"}
)

var (
	// jayName represents jay name.
	jayName = nook.Languages{
		language.AmericanEnglish:      jayAmericanEnglishName,
		language.CanadianFrench:       jayCanadianFrenchName,
		language.Dutch:                jayDutchName,
		language.French:               jayFrenchName,
		language.German:               jayGermanName,
		language.Italian:              jayItalianName,
		language.Japanese:             jayJapaneseName,
		language.Korean:               jayKoreanName,
		language.LatinAmericanSpanish: jayLatinAmericanSpanishName,
		language.Russian:              jayRussianName,
		language.SimplifiedChinese:    jaySimplifiedChineseName,
		language.Spanish:              jaySpanishName,
		language.TraditionalChinese:   jayTraditionalChineseName}
)

var (
	// jayCharacter represents jay character.
	jayCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: jayBirthday,
		Code:     jayCode,
		Key:      character.Jay,
		Gender:   gender.Male,
		Name:     jayName,
		Special:  false}
)

var (
	// jayAmericanEnglishPhrase represents jay american english phrase.
	jayAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "heeeeeyy"}

	// jayCanadianFrenchPhrase represents jay canadian french phrase.
	jayCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "maaaaaais"}

	// jayDutchPhrase represents jay dutch phrase.
	jayDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "heeeeeyy"}

	// jayFrenchPhrase represents jay french phrase.
	jayFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "maaaaaais"}

	// jayGermanPhrase represents jay german phrase.
	jayGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hiiiijjj"}

	// jayItalianPhrase represents jay italian phrase.
	jayItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cipinci"}

	// jayJapanesePhrase represents jay japanese phrase.
	jayJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でおます"}

	// jayKoreanPhrase represents jay korean phrase.
	jayKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그쵸"}

	// jayLatinAmericanSpanishPhrase represents jay latin american spanish phrase.
	jayLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "eeey"}

	// jayRussianPhrase represents jay russian phrase.
	jayRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хей-хей"}

	// jaySimplifiedChinesePhrase represents jay simplified chinese phrase.
	jaySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呢喃"}

	// jaySpanishPhrase represents jay spanish phrase.
	jaySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "heeey"}

	// jayTraditionalChinesePhrase represents jay traditional chinese phrase.
	jayTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "呢喃"}
)

var (
	// jayPhrase represents jay phrase.
	jayPhrase = nook.Languages{
		language.AmericanEnglish:      jayAmericanEnglishPhrase,
		language.CanadianFrench:       jayCanadianFrenchPhrase,
		language.Dutch:                jayDutchPhrase,
		language.French:               jayFrenchPhrase,
		language.German:               jayGermanPhrase,
		language.Italian:              jayItalianPhrase,
		language.Japanese:             jayJapanesePhrase,
		language.Korean:               jayKoreanPhrase,
		language.LatinAmericanSpanish: jayLatinAmericanSpanishPhrase,
		language.Russian:              jayRussianPhrase,
		language.SimplifiedChinese:    jaySimplifiedChinesePhrase,
		language.Spanish:              jaySpanishPhrase,
		language.TraditionalChinese:   jayTraditionalChinesePhrase}
)

var (
	// Jay represents jay.
	Jay = nook.Villager{
		Character:   jayCharacter,
		Personality: personality.Jock,
		Phrase:      jayPhrase}
)
