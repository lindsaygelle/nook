package ostrich

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
	// philBirthday represents phil birthday.
	philBirthday = nook.Birthday{
		Day:   27,
		Month: time.November}
)

var (
	// philCode represents phil code.
	philCode = nook.Code{
		Value: "ost07"}
)

var (
	// philAmericanEnglishName represents phil american english name.
	philAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Phil"}

	// philCanadianFrenchName represents phil canadian french name.
	philCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Phil"}

	// philDutchName represents phil dutch name.
	philDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Phil"}

	// philFrenchName represents phil french name.
	philFrenchName = nook.Name{
		Language: language.French,
		Value:    "Phil"}

	// philGermanName represents phil german name.
	philGermanName = nook.Name{
		Language: language.German,
		Value:    "Ingo"}

	// philItalianName represents phil italian name.
	philItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Enzo"}

	// philJapaneseName represents phil japanese name.
	philJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ケイン"}

	// philKoreanName represents phil korean name.
	philKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "케인"}

	// philLatinAmericanSpanishName represents phil latin american spanish name.
	philLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Amalio"}

	// philRussianName represents phil russian name.
	philRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фил"}

	// philSimplifiedChineseName represents phil simplified chinese name.
	philSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "凯恩"}

	// philSpanishName represents phil spanish name.
	philSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Amalio"}

	// philTraditionalChineseName represents phil traditional chinese name.
	philTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "凱恩"}
)

var (
	// philName represents phil name.
	philName = nook.Languages{
		language.AmericanEnglish:      philAmericanEnglishName,
		language.CanadianFrench:       philCanadianFrenchName,
		language.Dutch:                philDutchName,
		language.French:               philFrenchName,
		language.German:               philGermanName,
		language.Italian:              philItalianName,
		language.Japanese:             philJapaneseName,
		language.Korean:               philKoreanName,
		language.LatinAmericanSpanish: philLatinAmericanSpanishName,
		language.Russian:              philRussianName,
		language.SimplifiedChinese:    philSimplifiedChineseName,
		language.Spanish:              philSpanishName,
		language.TraditionalChinese:   philTraditionalChineseName}
)

var (
	// philCharacter represents phil character.
	philCharacter = nook.Character{
		Animal:   animal.Ostrich,
		Birthday: philBirthday,
		Code:     philCode,
		Key:      character.Phil,
		Gender:   gender.Male,
		Name:     philName,
		Special:  false}
)

var (
	// philAmericanEnglishPhrase represents phil american english phrase.
	philAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hurk"}

	// philCanadianFrenchPhrase represents phil canadian french phrase.
	philCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bécot"}

	// philDutchPhrase represents phil dutch phrase.
	philDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "struis"}

	// philFrenchPhrase represents phil french phrase.
	philFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bécot"}

	// philGermanPhrase represents phil german phrase.
	philGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zwoing"}

	// philItalianPhrase represents phil italian phrase.
	philItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "tambien"}

	// philJapanesePhrase represents phil japanese phrase.
	philJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ホロロ"}

	// philKoreanPhrase represents phil korean phrase.
	philKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "호롤로"}

	// philLatinAmericanSpanishPhrase represents phil latin american spanish phrase.
	philLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "jurk"}

	// philRussianPhrase represents phil russian phrase.
	philRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хех"}

	// philSimplifiedChinesePhrase represents phil simplified chinese phrase.
	philSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "轰隆隆"}

	// philSpanishPhrase represents phil spanish phrase.
	philSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "jurk"}

	// philTraditionalChinesePhrase represents phil traditional chinese phrase.
	philTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "轟隆隆"}
)

var (
	// philPhrase represents phil phrase.
	philPhrase = nook.Languages{
		language.AmericanEnglish:      philAmericanEnglishPhrase,
		language.CanadianFrench:       philCanadianFrenchPhrase,
		language.Dutch:                philDutchPhrase,
		language.French:               philFrenchPhrase,
		language.German:               philGermanPhrase,
		language.Italian:              philItalianPhrase,
		language.Japanese:             philJapanesePhrase,
		language.Korean:               philKoreanPhrase,
		language.LatinAmericanSpanish: philLatinAmericanSpanishPhrase,
		language.Russian:              philRussianPhrase,
		language.SimplifiedChinese:    philSimplifiedChinesePhrase,
		language.Spanish:              philSpanishPhrase,
		language.TraditionalChinese:   philTraditionalChinesePhrase}
)

var (
	// Phil represents phil.
	Phil = nook.Villager{
		Character:   philCharacter,
		Personality: personality.Smug,
		Phrase:      philPhrase}
)
