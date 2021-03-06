package horse

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
	clydeBirthday = nook.Birthday{
		Day:   1,
		Month: time.May}
)

var (
	clydeCode = nook.Code{
		Value: "hrs10"}
)

var (
	clydeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Clyde"}

	clydeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Dorian"}

	clydeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Clyde"}

	clydeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Dorian"}

	clydeGermanName = nook.Name{
		Language: language.German,
		Value:    "Tommi"}

	clydeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lallo"}

	clydeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "デースケ"}

	clydeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마철이"}

	clydeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Moe"}

	clydeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Клайд"}

	clydeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "笛笛"}

	clydeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Moe"}

	clydeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "笛笛"}
)

var (
	clydeName = nook.Languages{
		language.AmericanEnglish:      clydeAmericanEnglishName,
		language.CanadianFrench:       clydeCanadianFrenchName,
		language.Dutch:                clydeDutchName,
		language.French:               clydeFrenchName,
		language.German:               clydeGermanName,
		language.Italian:              clydeItalianName,
		language.Japanese:             clydeJapaneseName,
		language.Korean:               clydeKoreanName,
		language.LatinAmericanSpanish: clydeLatinAmericanSpanishName,
		language.Russian:              clydeRussianName,
		language.SimplifiedChinese:    clydeSimplifiedChineseName,
		language.Spanish:              clydeSpanishName,
		language.TraditionalChinese:   clydeTraditionalChineseName}
)

var (
	clydeCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: clydeBirthday,
		Code:     clydeCode,
		Key:      character.Clyde,
		Gender:   gender.Male,
		Name:     clydeName,
		Special:  false}
)

var (
	clydeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "clip clawp"}

	clydeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mouh-ouh"}

	clydeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "klipklop"}

	clydeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mouh-ouh"}

	clydeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schauder"}

	clydeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "clop clop"}

	clydeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ぴろぴろ"}

	clydeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "다그닥"}

	clydeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "clip-clop"}

	clydeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "цок-цок"}

	clydeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "吹卷卷"}

	clydeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "troteras"}

	clydeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "吹捲捲"}
)

var (
	clydePhrase = nook.Languages{
		language.AmericanEnglish:      clydeAmericanEnglishPhrase,
		language.CanadianFrench:       clydeCanadianFrenchPhrase,
		language.Dutch:                clydeDutchPhrase,
		language.French:               clydeFrenchPhrase,
		language.German:               clydeGermanPhrase,
		language.Italian:              clydeItalianPhrase,
		language.Japanese:             clydeJapanesePhrase,
		language.Korean:               clydeKoreanPhrase,
		language.LatinAmericanSpanish: clydeLatinAmericanSpanishPhrase,
		language.Russian:              clydeRussianPhrase,
		language.SimplifiedChinese:    clydeSimplifiedChinesePhrase,
		language.Spanish:              clydeSpanishPhrase,
		language.TraditionalChinese:   clydeTraditionalChinesePhrase}
)

var (
	Clyde = nook.Villager{
		Character:   clydeCharacter,
		Personality: personality.Lazy,
		Phrase:      clydePhrase}
)
