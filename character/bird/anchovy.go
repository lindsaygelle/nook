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
	anchovyBirthday = nook.Birthday{
		Day:   4,
		Month: time.March}
)

var (
	anchovyCode = nook.Code{
		Value: "brd02"}
)

var (
	anchovyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Anchovy"}

	anchovyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Miguel"}

	anchovyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Anchovy"}

	anchovyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Miguel"}

	anchovyGermanName = nook.Name{
		Language: language.German,
		Value:    "Armin"}

	anchovyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Acciuga"}

	anchovyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アンチョビ"}

	anchovyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "안쵸비"}

	anchovyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Boquerón"}

	anchovyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Анчоуи"}

	anchovySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "凤尾"}

	anchovySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Boquerón"}

	anchovyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鳳尾"}
)

var (
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
	anchovyCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: anchovyBirthday,
		Code:     anchovyCode,
		Key:      character.Anchovy,
		Gender:   gender.Male,
		Name:     anchovyName}
)

var (
	anchovyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "chuurp"}

	anchovyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "muf-muf"}

	anchovyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "tsjiiirp"}

	anchovyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "muf-muf"}

	anchovyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tschurp"}

	anchovyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ciripì"}

	anchovyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でシ"}

	anchovyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "이어라"}

	anchovyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "piu-piu"}

	anchovyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "чик-чирик"}

	anchovySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "如斯"}

	anchovySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "piu-piu"}

	anchovyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "如斯"}
)

var (
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
	Anchovy = nook.Villager{
		Character:   anchovyCharacter,
		Personality: personality.Lazy,
		Phrase:      anchovyPhrase}
)
