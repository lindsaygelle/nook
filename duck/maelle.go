package duck

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	maelleBirthday = nook.Birthday{
		Day:   8,
		Month: time.April}
)

var (
	maelleCode = nook.Code{
		Value: "duk03"}
)

var (
	maelleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Maelle"}

	maelleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Maëllecanardou"}

	maelleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Maellepulletje"}

	maelleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maëllecanardou"}

	maelleGermanName = nook.Name{
		Language: language.German,
		Value:    "Sissientchen"}

	maelleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Palmitapaperotto"}

	maelleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アンヌふぅ"}

	maelleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "앤우아"}

	maelleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Patidifúplumi-piú"}

	maelleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Маэльжелторотик"}

	maelleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "安妮叹"}

	maelleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Patidifúperla"}

	maelleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "安妮嘆"}
)

var (
	maelleName = nook.Languages{
		language.AmericanEnglish:      maelleAmericanEnglishName,
		language.CanadianFrench:       maelleCanadianFrenchName,
		language.Dutch:                maelleDutchName,
		language.French:               maelleFrenchName,
		language.German:               maelleGermanName,
		language.Italian:              maelleItalianName,
		language.Japanese:             maelleJapaneseName,
		language.Korean:               maelleKoreanName,
		language.LatinAmericanSpanish: maelleLatinAmericanSpanishName,
		language.Russian:              maelleRussianName,
		language.SimplifiedChinese:    maelleSimplifiedChineseName,
		language.Spanish:              maelleSpanishName,
		language.TraditionalChinese:   maelleTraditionalChineseName}
)

var (
	maelleCharacter = nook.Character{
		Animal:   Duck,
		Birthday: maelleBirthday,
		Code:     maelleCode,
		Gender:   nook.Female,
		Name:     maelleName}
)

var (
	maelleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "duckling"}

	maelleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "canardou"}

	maelleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pulletje"}

	maelleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "canardou"}

	maelleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "entchen"}

	maelleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "paperotto"}

	maelleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ふぅ"}

	maelleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우아"}

	maelleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "plumi-piú"}

	maelleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "желторотик"}

	maelleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "叹"}

	maelleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "perla"}

	maelleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘆"}
)

var (
	maellePhrase = nook.Languages{
		language.AmericanEnglish:      maelleAmericanEnglishName,
		language.CanadianFrench:       maelleCanadianFrenchName,
		language.Dutch:                maelleDutchName,
		language.French:               maelleFrenchName,
		language.German:               maelleGermanName,
		language.Italian:              maelleItalianName,
		language.Japanese:             maelleJapaneseName,
		language.Korean:               maelleKoreanName,
		language.LatinAmericanSpanish: maelleLatinAmericanSpanishName,
		language.Russian:              maelleRussianName,
		language.SimplifiedChinese:    maelleSimplifiedChineseName,
		language.Spanish:              maelleSpanishName,
		language.TraditionalChinese:   maelleTraditionalChineseName}
)

var (
	Maelle = nook.Villager{
		Character:   maelleCharacter,
		Personality: nook.Snooty,
		Phrase:      maellePhrase}
)
