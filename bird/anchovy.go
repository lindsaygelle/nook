package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Miguelmuf-muf"}

	anchovyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Anchovytsjiiirp"}

	anchovyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Miguelmuf-muf"}

	anchovyGermanName = nook.Name{
		Language: language.German,
		Value:    "Armintschurp"}

	anchovyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Acciugaciripì"}

	anchovyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アンチョビでシ"}

	anchovyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "안쵸비이어라"}

	anchovyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Boquerónpiu-piu"}

	anchovyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Анчоуичик-чирик"}

	anchovySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "凤尾如斯"}

	anchovySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Boquerónpiu-piu"}

	anchovyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鳳尾如斯"}
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
		Animal:   Bird,
		Birthday: anchovyBirthday,
		Code:     anchovyCode,
		Gender:   nook.Male,
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
	Anchovy = nook.Villager{
		Character:   anchovyCharacter,
		Personality: nook.Lazy,
		Phrase:      anchovyPhrase}
)
