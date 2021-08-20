package rabbit

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
	docBirthday = nook.Birthday{
		Day:   16,
		Month: time.March}
)

var (
	docCode = nook.Code{
		Value: "rbt10"}
)

var (
	docAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Doc"}

	docCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Doc"}

	docDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Doc"}

	docFrenchName = nook.Name{
		Language: language.French,
		Value:    "Doc"}

	docGermanName = nook.Name{
		Language: language.German,
		Value:    "Gustl"}

	docItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Doc"}

	docJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "トビオ"}

	docKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "토니"}

	docLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Crispín"}

	docRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Док"}

	docSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿飞"}

	docSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Crispín"}

	docTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿飛"}
)

var (
	docName = nook.Languages{
		language.AmericanEnglish:      docAmericanEnglishName,
		language.CanadianFrench:       docCanadianFrenchName,
		language.Dutch:                docDutchName,
		language.French:               docFrenchName,
		language.German:               docGermanName,
		language.Italian:              docItalianName,
		language.Japanese:             docJapaneseName,
		language.Korean:               docKoreanName,
		language.LatinAmericanSpanish: docLatinAmericanSpanishName,
		language.Russian:              docRussianName,
		language.SimplifiedChinese:    docSimplifiedChineseName,
		language.Spanish:              docSpanishName,
		language.TraditionalChinese:   docTraditionalChineseName}
)

var (
	docCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: docBirthday,
		Code:     docCode,
		Key:      character.Doc,
		Gender:   gender.Male,
		Name:     docName,
		Special:  false}
)

var (
	docAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "old bunny"}

	docCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tout bleu"}

	docDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ouwe reus"}

	docFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "morbleu"}

	docGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hasi"}

	docItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "soufflé"}

	docJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですね"}

	docKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그렇잖아"}

	docLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bingboing"}

	docRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "зайчище"}

	docSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是啦"}

	docSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "y tal"}

	docTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是啦"}
)

var (
	docPhrase = nook.Languages{
		language.AmericanEnglish:      docAmericanEnglishPhrase,
		language.CanadianFrench:       docCanadianFrenchPhrase,
		language.Dutch:                docDutchPhrase,
		language.French:               docFrenchPhrase,
		language.German:               docGermanPhrase,
		language.Italian:              docItalianPhrase,
		language.Japanese:             docJapanesePhrase,
		language.Korean:               docKoreanPhrase,
		language.LatinAmericanSpanish: docLatinAmericanSpanishPhrase,
		language.Russian:              docRussianPhrase,
		language.SimplifiedChinese:    docSimplifiedChinesePhrase,
		language.Spanish:              docSpanishPhrase,
		language.TraditionalChinese:   docTraditionalChinesePhrase}
)

var (
	Doc = nook.Villager{
		Character:   docCharacter,
		Personality: personality.Lazy,
		Phrase:      docPhrase}
)
