package alligator

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	alliBirthday = nook.Birthday{
		Day:   8,
		Month: time.November}
)

var (
	alliCode = nook.Code{
		Value: "crd01"}
)

var (
	alliAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Alli"}

	alliCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Alliegraaagh"}

	alliDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Alligraaach"}

	alliFrenchName = nook.Name{
		Language: language.French,
		Value:    "Alliegraaagh"}

	alliGermanName = nook.Name{
		Language: language.German,
		Value:    "Aligraaach"}

	alliItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Alligraaag"}

	alliJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クロコどすえ"}

	alliKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "크로크얘야"}

	alliLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Codrilamuamú"}

	alliRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Эллигра-а-а"}

	alliSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鳄罗思鳄鱼皮"}

	alliSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Codrilalagartija"}

	alliTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鱷羅思鱷魚皮"}
)

var (
	alliName = nook.Languages{
		language.AmericanEnglish:      alliAmericanEnglishName,
		language.CanadianFrench:       alliCanadianFrenchName,
		language.Dutch:                alliDutchName,
		language.French:               alliFrenchName,
		language.German:               alliGermanName,
		language.Italian:              alliItalianName,
		language.Japanese:             alliJapaneseName,
		language.Korean:               alliKoreanName,
		language.LatinAmericanSpanish: alliLatinAmericanSpanishName,
		language.Russian:              alliRussianName,
		language.SimplifiedChinese:    alliSimplifiedChineseName,
		language.Spanish:              alliSpanishName,
		language.TraditionalChinese:   alliTraditionalChineseName}
)

var (
	alliCharacter = nook.Character{
		Animal:   Alligator,
		Birthday: alliBirthday,
		Code:     alliCode,
		Gender:   nook.Female,
		Name:     alliName}
)

var (
	alliAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "graaagh"}

	alliCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "graaagh"}

	alliDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "graaach"}

	alliFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "graaagh"}

	alliGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "graaach"}

	alliItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "graaag"}

	alliJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "どすえ"}

	alliKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "얘야"}

	alliLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "muamú"}

	alliRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "гра-а-а"}

	alliSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鳄鱼皮"}

	alliSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "lagartija"}

	alliTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鱷魚皮"}
)

var (
	alliPhrase = nook.Languages{
		language.AmericanEnglish:      alliAmericanEnglishName,
		language.CanadianFrench:       alliCanadianFrenchName,
		language.Dutch:                alliDutchName,
		language.French:               alliFrenchName,
		language.German:               alliGermanName,
		language.Italian:              alliItalianName,
		language.Japanese:             alliJapaneseName,
		language.Korean:               alliKoreanName,
		language.LatinAmericanSpanish: alliLatinAmericanSpanishName,
		language.Russian:              alliRussianName,
		language.SimplifiedChinese:    alliSimplifiedChineseName,
		language.Spanish:              alliSpanishName,
		language.TraditionalChinese:   alliTraditionalChineseName}
)

var (
	Alli = nook.Villager{
		Character:   alliCharacter,
		Personality: nook.Snooty,
		Phrase:      alliPhrase}
)
