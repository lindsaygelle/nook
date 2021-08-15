package rabbit

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	coleBirthday = nook.Birthday{
		Day:   10,
		Month: time.August}
)

var (
	coleCode = nook.Code{
		Value: "rbt18"}
)

var (
	coleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cole"}

	coleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Épicurechaaauuud"}

	coleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Colegaaast"}

	coleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Épicurechaaauuud"}

	coleGermanName = nook.Name{
		Language: language.German,
		Value:    "Karlmümml"}

	coleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lelloniglio"}

	coleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アマミンそうさね"}

	coleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아마민나도나도"}

	coleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nicolásmuyayo"}

	coleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Коулзайчелло"}

	coleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿甘就是如此"}

	coleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Nicolásmuyayo"}

	coleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿甘就是如此"}
)

var (
	coleName = nook.Languages{
		language.AmericanEnglish:      coleAmericanEnglishName,
		language.CanadianFrench:       coleCanadianFrenchName,
		language.Dutch:                coleDutchName,
		language.French:               coleFrenchName,
		language.German:               coleGermanName,
		language.Italian:              coleItalianName,
		language.Japanese:             coleJapaneseName,
		language.Korean:               coleKoreanName,
		language.LatinAmericanSpanish: coleLatinAmericanSpanishName,
		language.Russian:              coleRussianName,
		language.SimplifiedChinese:    coleSimplifiedChineseName,
		language.Spanish:              coleSpanishName,
		language.TraditionalChinese:   coleTraditionalChineseName}
)

var (
	coleCharacter = nook.Character{
		Animal:   Rabbit,
		Birthday: coleBirthday,
		Code:     coleCode,
		Gender:   nook.Male,
		Name:     coleName}
)

var (
	coleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "duuudecoooooool"}

	coleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chaaauuud"}

	coleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "gaaast"}

	coleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "chaaauuud"}

	coleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mümml"}

	coleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "niglio"}

	coleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "そうさね"}

	coleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "나도나도"}

	coleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "muyayo"}

	coleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "зайчелло"}

	coleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "就是如此"}

	coleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "muyayo"}

	coleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "就是如此"}
)

var (
	colePhrase = nook.Languages{
		language.AmericanEnglish:      coleAmericanEnglishName,
		language.CanadianFrench:       coleCanadianFrenchName,
		language.Dutch:                coleDutchName,
		language.French:               coleFrenchName,
		language.German:               coleGermanName,
		language.Italian:              coleItalianName,
		language.Japanese:             coleJapaneseName,
		language.Korean:               coleKoreanName,
		language.LatinAmericanSpanish: coleLatinAmericanSpanishName,
		language.Russian:              coleRussianName,
		language.SimplifiedChinese:    coleSimplifiedChineseName,
		language.Spanish:              coleSpanishName,
		language.TraditionalChinese:   coleTraditionalChineseName}
)

var (
	Cole = nook.Villager{
		Character:   coleCharacter,
		Personality: nook.Lazy,
		Phrase:      colePhrase}
)
