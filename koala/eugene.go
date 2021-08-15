package koala

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	eugeneBirthday = nook.Birthday{
		Day:   26,
		Month: time.October}
)

var (
	eugeneCode = nook.Code{
		Value: "kal10"}
)

var (
	eugeneAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Eugene"}

	eugeneCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Jamy"}

	eugeneDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Eugene"}

	eugeneFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jamy"}

	eugeneGermanName = nook.Name{
		Language: language.German,
		Value:    "Sunny"}

	eugeneItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Corrado"}

	eugeneJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ロッキー"}

	eugeneKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "코알"}

	eugeneLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Eucalín"}

	eugeneRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Юджин"}

	eugeneSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗奇"}

	eugeneSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Eucalín"}

	eugeneTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羅奇"}
)

var (
	eugeneName = nook.Languages{
		language.AmericanEnglish:      eugeneAmericanEnglishName,
		language.CanadianFrench:       eugeneCanadianFrenchName,
		language.Dutch:                eugeneDutchName,
		language.French:               eugeneFrenchName,
		language.German:               eugeneGermanName,
		language.Italian:              eugeneItalianName,
		language.Japanese:             eugeneJapaneseName,
		language.Korean:               eugeneKoreanName,
		language.LatinAmericanSpanish: eugeneLatinAmericanSpanishName,
		language.Russian:              eugeneRussianName,
		language.SimplifiedChinese:    eugeneSimplifiedChineseName,
		language.Spanish:              eugeneSpanishName,
		language.TraditionalChinese:   eugeneTraditionalChineseName}
)

var (
	eugeneCharacter = nook.Character{
		Animal:   Koala,
		Birthday: eugeneBirthday,
		Code:     eugeneCode,
		Gender:   nook.Male,
		Name:     eugeneName}
)

var (
	eugeneAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "yeah buddy"}

	eugeneCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "oh yeah"}

	eugeneDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "yeah"}

	eugeneFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "yapaphoto"}

	eugeneGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "yeah"}

	eugeneItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "koahola"}

	eugeneJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "おいおい"}

	eugeneKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "야야야"}

	eugeneLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ouyeah"}

	eugeneRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "приятель"}

	eugeneSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喂喂"}

	eugeneSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ouyeah"}

	eugeneTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喂喂"}
)

var (
	eugenePhrase = nook.Languages{
		language.AmericanEnglish:      eugeneAmericanEnglishName,
		language.CanadianFrench:       eugeneCanadianFrenchName,
		language.Dutch:                eugeneDutchName,
		language.French:               eugeneFrenchName,
		language.German:               eugeneGermanName,
		language.Italian:              eugeneItalianName,
		language.Japanese:             eugeneJapaneseName,
		language.Korean:               eugeneKoreanName,
		language.LatinAmericanSpanish: eugeneLatinAmericanSpanishName,
		language.Russian:              eugeneRussianName,
		language.SimplifiedChinese:    eugeneSimplifiedChineseName,
		language.Spanish:              eugeneSpanishName,
		language.TraditionalChinese:   eugeneTraditionalChineseName}
)

var (
	Eugene = nook.Villager{
		Character:   eugeneCharacter,
		Personality: nook.Smug,
		Phrase:      eugenePhrase}
)
