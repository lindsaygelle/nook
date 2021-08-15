package monkey

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	shariBirthday = nook.Birthday{
		Day:   10,
		Month: time.April}
)

var (
	shariCode = nook.Code{
		Value: "mnk07"}
)

var (
	shariAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Shari"}

	shariCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lunatavu"}

	shariDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sharibrutaaltje"}

	shariFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lunatavu"}

	shariGermanName = nook.Name{
		Language: language.German,
		Value:    "Utamakimaki"}

	shariItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ninìbabbuh"}

	shariJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シェリーウキキ"}

	shariKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "젤리우끼네"}

	shariLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Montsebabuí"}

	shariRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Шарищекасто"}

	shariSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雪莉呜吱吱"}

	shariSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Montsetunante"}

	shariTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雪莉嗚吱吱"}
)

var (
	shariName = nook.Languages{
		language.AmericanEnglish:      shariAmericanEnglishName,
		language.CanadianFrench:       shariCanadianFrenchName,
		language.Dutch:                shariDutchName,
		language.French:               shariFrenchName,
		language.German:               shariGermanName,
		language.Italian:              shariItalianName,
		language.Japanese:             shariJapaneseName,
		language.Korean:               shariKoreanName,
		language.LatinAmericanSpanish: shariLatinAmericanSpanishName,
		language.Russian:              shariRussianName,
		language.SimplifiedChinese:    shariSimplifiedChineseName,
		language.Spanish:              shariSpanishName,
		language.TraditionalChinese:   shariTraditionalChineseName}
)

var (
	shariCharacter = nook.Character{
		Animal:   Monkey,
		Birthday: shariBirthday,
		Code:     shariCode,
		Gender:   nook.Female,
		Name:     shariName}
)

var (
	shariAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cheeky"}

	shariCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tavu"}

	shariDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "brutaaltje"}

	shariFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tavu"}

	shariGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "makimaki"}

	shariItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "babbuh"}

	shariJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウキキ"}

	shariKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우끼네"}

	shariLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "babuí"}

	shariRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "щекасто"}

	shariSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呜吱吱"}

	shariSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tunante"}

	shariTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗚吱吱"}
)

var (
	shariPhrase = nook.Languages{
		language.AmericanEnglish:      shariAmericanEnglishName,
		language.CanadianFrench:       shariCanadianFrenchName,
		language.Dutch:                shariDutchName,
		language.French:               shariFrenchName,
		language.German:               shariGermanName,
		language.Italian:              shariItalianName,
		language.Japanese:             shariJapaneseName,
		language.Korean:               shariKoreanName,
		language.LatinAmericanSpanish: shariLatinAmericanSpanishName,
		language.Russian:              shariRussianName,
		language.SimplifiedChinese:    shariSimplifiedChineseName,
		language.Spanish:              shariSpanishName,
		language.TraditionalChinese:   shariTraditionalChineseName}
)

var (
	Shari = nook.Villager{
		Character:   shariCharacter,
		Personality: nook.BigSister,
		Phrase:      shariPhrase}
)
