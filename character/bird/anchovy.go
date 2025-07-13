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
	// anchovyBirthday represents anchovy birthday.
	anchovyBirthday = nook.Birthday{
		Day:   4,
		Month: time.March}
)

var (
	// anchovyCode represents anchovy code.
	anchovyCode = nook.Code{
		Value: "brd02"}
)

var (
	// anchovyAmericanEnglishName represents anchovy american english name.
	anchovyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Anchovy"}

	// anchovyCanadianFrenchName represents anchovy canadian french name.
	anchovyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Miguel"}

	// anchovyDutchName represents anchovy dutch name.
	anchovyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Anchovy"}

	// anchovyFrenchName represents anchovy french name.
	anchovyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Miguel"}

	// anchovyGermanName represents anchovy german name.
	anchovyGermanName = nook.Name{
		Language: language.German,
		Value:    "Armin"}

	// anchovyItalianName represents anchovy italian name.
	anchovyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Acciuga"}

	// anchovyJapaneseName represents anchovy japanese name.
	anchovyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アンチョビ"}

	// anchovyKoreanName represents anchovy korean name.
	anchovyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "안쵸비"}

	// anchovyLatinAmericanSpanishName represents anchovy latin american spanish name.
	anchovyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Boquerón"}

	// anchovyRussianName represents anchovy russian name.
	anchovyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Анчоуи"}

	// anchovySimplifiedChineseName represents anchovy simplified chinese name.
	anchovySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "凤尾"}

	// anchovySpanishName represents anchovy spanish name.
	anchovySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Boquerón"}

	// anchovyTraditionalChineseName represents anchovy traditional chinese name.
	anchovyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鳳尾"}
)

var (
	// anchovyName represents anchovy name.
	anchovyName = nook.Languages{
		language.AmericanEnglish:      anchovyAmericanEnglishName,
		language.CanadianFrench:       anchovyCanadianFrenchName,
		language.Dutch:                anchovyDutchName,
		language.French:               anchovyFrenchName,
		language.German:               anchovyGermanName,
		language.Italian:              anchovyItalianName,
		language.Japanese:             anchovyJapaneseName,
		language.Korean:               anchovyKoreanName,
		language.LatinAmericanSpanish: anchovyLatinAmericanSpanishName,
		language.Russian:              anchovyRussianName,
		language.SimplifiedChinese:    anchovySimplifiedChineseName,
		language.Spanish:              anchovySpanishName,
		language.TraditionalChinese:   anchovyTraditionalChineseName}
)

var (
	// anchovyCharacter represents anchovy character.
	anchovyCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: anchovyBirthday,
		Code:     anchovyCode,
		Key:      character.Anchovy,
		Gender:   gender.Male,
		Name:     anchovyName,
		Special:  false}
)

var (
	// anchovyAmericanEnglishPhrase represents anchovy american english phrase.
	anchovyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "chuurp"}

	// anchovyCanadianFrenchPhrase represents anchovy canadian french phrase.
	anchovyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "muf-muf"}

	// anchovyDutchPhrase represents anchovy dutch phrase.
	anchovyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "tsjiiirp"}

	// anchovyFrenchPhrase represents anchovy french phrase.
	anchovyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "muf-muf"}

	// anchovyGermanPhrase represents anchovy german phrase.
	anchovyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tschurp"}

	// anchovyItalianPhrase represents anchovy italian phrase.
	anchovyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ciripì"}

	// anchovyJapanesePhrase represents anchovy japanese phrase.
	anchovyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でシ"}

	// anchovyKoreanPhrase represents anchovy korean phrase.
	anchovyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "이어라"}

	// anchovyLatinAmericanSpanishPhrase represents anchovy latin american spanish phrase.
	anchovyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "piu-piu"}

	// anchovyRussianPhrase represents anchovy russian phrase.
	anchovyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "чик-чирик"}

	// anchovySimplifiedChinesePhrase represents anchovy simplified chinese phrase.
	anchovySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "如斯"}

	// anchovySpanishPhrase represents anchovy spanish phrase.
	anchovySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "piu-piu"}

	// anchovyTraditionalChinesePhrase represents anchovy traditional chinese phrase.
	anchovyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "如斯"}
)

var (
	// anchovyPhrase represents anchovy phrase.
	anchovyPhrase = nook.Languages{
		language.AmericanEnglish:      anchovyAmericanEnglishPhrase,
		language.CanadianFrench:       anchovyCanadianFrenchPhrase,
		language.Dutch:                anchovyDutchPhrase,
		language.French:               anchovyFrenchPhrase,
		language.German:               anchovyGermanPhrase,
		language.Italian:              anchovyItalianPhrase,
		language.Japanese:             anchovyJapanesePhrase,
		language.Korean:               anchovyKoreanPhrase,
		language.LatinAmericanSpanish: anchovyLatinAmericanSpanishPhrase,
		language.Russian:              anchovyRussianPhrase,
		language.SimplifiedChinese:    anchovySimplifiedChinesePhrase,
		language.Spanish:              anchovySpanishPhrase,
		language.TraditionalChinese:   anchovyTraditionalChinesePhrase}
)

var (
	// Anchovy represents anchovy.
	Anchovy = nook.Villager{
		Character:   anchovyCharacter,
		Personality: personality.Lazy,
		Phrase:      anchovyPhrase}
)
