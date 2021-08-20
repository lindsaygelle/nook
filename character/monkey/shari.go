package monkey

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
		Value:    "Luna"}

	shariDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Shari"}

	shariFrenchName = nook.Name{
		Language: language.French,
		Value:    "Luna"}

	shariGermanName = nook.Name{
		Language: language.German,
		Value:    "Uta"}

	shariItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ninì"}

	shariJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シェリー"}

	shariKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "젤리"}

	shariLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Montse"}

	shariRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Шари"}

	shariSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雪莉"}

	shariSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Montse"}

	shariTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雪莉"}
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
		Animal:   animal.Monkey,
		Birthday: shariBirthday,
		Code:     shariCode,
		Key:      character.Shari,
		Gender:   gender.Female,
		Name:     shariName,
		Special:  false}
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
		language.AmericanEnglish:      shariAmericanEnglishPhrase,
		language.CanadianFrench:       shariCanadianFrenchPhrase,
		language.Dutch:                shariDutchPhrase,
		language.French:               shariFrenchPhrase,
		language.German:               shariGermanPhrase,
		language.Italian:              shariItalianPhrase,
		language.Japanese:             shariJapanesePhrase,
		language.Korean:               shariKoreanPhrase,
		language.LatinAmericanSpanish: shariLatinAmericanSpanishPhrase,
		language.Russian:              shariRussianPhrase,
		language.SimplifiedChinese:    shariSimplifiedChinesePhrase,
		language.Spanish:              shariSpanishPhrase,
		language.TraditionalChinese:   shariTraditionalChinesePhrase}
)

var (
	Shari = nook.Villager{
		Character:   shariCharacter,
		Personality: personality.BigSister,
		Phrase:      shariPhrase}
)
