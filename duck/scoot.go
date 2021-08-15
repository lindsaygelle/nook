package duck

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	scootBirthday = nook.Birthday{
		Day:   13,
		Month: time.June}
)

var (
	scootCode = nook.Code{
		Value: "duk10"}
)

var (
	scootAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Scoot"}

	scootCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Scooter"}

	scootDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Scoot"}

	scootFrenchName = nook.Name{
		Language: language.French,
		Value:    "Scooter"}

	scootGermanName = nook.Name{
		Language: language.German,
		Value:    "Helmut"}

	scootItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Guizzo"}

	scootJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マモル"}

	scootKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "지키미"}

	scootLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Chema"}

	scootRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Скут"}

	scootSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿守"}

	scootSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Chema"}

	scootTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿守"}
)

var (
	scootName = nook.Languages{
		language.AmericanEnglish:      scootAmericanEnglishName,
		language.CanadianFrench:       scootCanadianFrenchName,
		language.Dutch:                scootDutchName,
		language.French:               scootFrenchName,
		language.German:               scootGermanName,
		language.Italian:              scootItalianName,
		language.Japanese:             scootJapaneseName,
		language.Korean:               scootKoreanName,
		language.LatinAmericanSpanish: scootLatinAmericanSpanishName,
		language.Russian:              scootRussianName,
		language.SimplifiedChinese:    scootSimplifiedChineseName,
		language.Spanish:              scootSpanishName,
		language.TraditionalChinese:   scootTraditionalChineseName}
)

var (
	scootCharacter = nook.Character{
		Animal:   Duck,
		Birthday: scootBirthday,
		Code:     scootCode,
		Gender:   nook.Male,
		Name:     scootName}
)

var (
	scootAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "zip zoom"}

	scootCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "zak zak"}

	scootDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "broem"}

	scootFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "zak zak"}

	scootGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "flitzer"}

	scootItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "zip zoom"}

	scootJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "グワッ"}

	scootKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꾸왁"}

	scootLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "zapizum"}

	scootRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "брум-брум"}

	scootSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呱呱"}

	scootSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "zapizum"}

	scootTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "呱呱"}
)

var (
	scootPhrase = nook.Languages{
		language.AmericanEnglish:      scootAmericanEnglishName,
		language.CanadianFrench:       scootCanadianFrenchName,
		language.Dutch:                scootDutchName,
		language.French:               scootFrenchName,
		language.German:               scootGermanName,
		language.Italian:              scootItalianName,
		language.Japanese:             scootJapaneseName,
		language.Korean:               scootKoreanName,
		language.LatinAmericanSpanish: scootLatinAmericanSpanishName,
		language.Russian:              scootRussianName,
		language.SimplifiedChinese:    scootSimplifiedChineseName,
		language.Spanish:              scootSpanishName,
		language.TraditionalChinese:   scootTraditionalChineseName}
)

var (
	Scoot = nook.Villager{
		Character:   scootCharacter,
		Personality: nook.Jock,
		Phrase:      scootPhrase}
)
