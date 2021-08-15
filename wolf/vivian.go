package wolf

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	vivianBirthday = nook.Birthday{
		Day:   26,
		Month: time.January}
)

var (
	vivianCode = nook.Code{
		Value: "wol08"}
)

var (
	vivianAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Vivian"}

	vivianCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Viviane"}

	vivianDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Vivian"}

	vivianFrenchName = nook.Name{
		Language: language.French,
		Value:    "Viviane"}

	vivianGermanName = nook.Name{
		Language: language.German,
		Value:    "Viviane"}

	vivianItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Viviana"}

	vivianJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヴァネッサ"}

	vivianKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "바네사"}

	vivianLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Viviana"}

	vivianRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Вивьен"}

	vivianSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "范妮沙"}

	vivianSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Viviana"}

	vivianTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "范妮莎"}
)

var (
	vivianName = nook.Languages{
		language.AmericanEnglish:      vivianAmericanEnglishName,
		language.CanadianFrench:       vivianCanadianFrenchName,
		language.Dutch:                vivianDutchName,
		language.French:               vivianFrenchName,
		language.German:               vivianGermanName,
		language.Italian:              vivianItalianName,
		language.Japanese:             vivianJapaneseName,
		language.Korean:               vivianKoreanName,
		language.LatinAmericanSpanish: vivianLatinAmericanSpanishName,
		language.Russian:              vivianRussianName,
		language.SimplifiedChinese:    vivianSimplifiedChineseName,
		language.Spanish:              vivianSpanishName,
		language.TraditionalChinese:   vivianTraditionalChineseName}
)

var (
	vivianCharacter = nook.Character{
		Animal:   Wolf,
		Birthday: vivianBirthday,
		Code:     vivianCode,
		Gender:   nook.Female,
		Name:     vivianName}
)

var (
	vivianAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "piffle"}

	vivianCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "crocroque"}

	vivianDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "nonsens"}

	vivianFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "crocroque"}

	vivianGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "scholli"}

	vivianItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "arruuu"}

	vivianJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だわよ"}

	vivianKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그러하다"}

	vivianLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "aú-aú"}

	vivianRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "вздор"}

	vivianSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喔呵呵"}

	vivianSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "aú-aú"}

	vivianTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喔呵呵"}
)

var (
	vivianPhrase = nook.Languages{
		language.AmericanEnglish:      vivianAmericanEnglishName,
		language.CanadianFrench:       vivianCanadianFrenchName,
		language.Dutch:                vivianDutchName,
		language.French:               vivianFrenchName,
		language.German:               vivianGermanName,
		language.Italian:              vivianItalianName,
		language.Japanese:             vivianJapaneseName,
		language.Korean:               vivianKoreanName,
		language.LatinAmericanSpanish: vivianLatinAmericanSpanishName,
		language.Russian:              vivianRussianName,
		language.SimplifiedChinese:    vivianSimplifiedChineseName,
		language.Spanish:              vivianSpanishName,
		language.TraditionalChinese:   vivianTraditionalChineseName}
)

var (
	Vivian = nook.Villager{
		Character:   vivianCharacter,
		Personality: nook.Snooty,
		Phrase:      vivianPhrase}
)
