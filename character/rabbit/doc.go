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
	// docBirthday represents Doc's birthday.
	docBirthday = nook.Birthday{
		Day:   16,
		Month: time.March}
)

var (
	// docCode represents Doc's unique code ("rbt10").
	docCode = nook.Code{
		Value: "rbt10"}
)

var (
	// docAmericanEnglishName represents Doc's name in American English.
	docAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Doc"}

	// docCanadianFrenchName represents Doc's name in Canadian French.
	docCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Doc"}

	// docDutchName represents Doc's name in Dutch.
	docDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Doc"}

	// docFrenchName represents Doc's name in French.
	docFrenchName = nook.Name{
		Language: language.French,
		Value:    "Doc"}

	// docGermanName represents Doc's name in German.
	docGermanName = nook.Name{
		Language: language.German,
		Value:    "Gustl"}

	// docItalianName represents Doc's name in Italian.
	docItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Doc"}

	// docJapaneseName represents Doc's name in Japanese.
	docJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "トビオ"}

	// docKoreanName represents Doc's name in Korean.
	docKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "토니"}

	// docLatinAmericanSpanishName represents Doc's name in Latin American Spanish.
	docLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Crispín"}

	// docRussianName represents Doc's name in Russian.
	docRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Док"}

	// docSimplifiedChineseName represents Doc's name in Simplified Chinese.
	docSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿飞"}

	// docSpanishName represents Doc's name in Spanish.
	docSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Crispín"}

	// docTraditionalChineseName represents Doc's name in Traditional Chinese.
	docTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿飛"}
)

var (
	// docName represents Doc's name in different languages.
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
	// docCharacter represents Doc's character information.
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
	// docAmericanEnglishPhrase represents Doc's phrase in American English.
	docAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "old bunny"}

	// docCanadianFrenchPhrase represents Doc's phrase in Canadian French.
	docCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tout bleu"}

	// docDutchPhrase represents Doc's phrase in Dutch.
	docDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ouwe reus"}

	// docFrenchPhrase represents Doc's phrase in French.
	docFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "morbleu"}

	// docGermanPhrase represents Doc's phrase in German.
	docGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hasi"}

	// docItalianPhrase represents Doc's phrase in Italian.
	docItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "soufflé"}

	// docJapanesePhrase represents Doc's phrase in Japanese.
	docJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですね"}

	// docKoreanPhrase represents Doc's phrase in Korean.
	docKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그렇잖아"}

	// docLatinAmericanSpanishPhrase represents Doc's phrase in Latin American Spanish.
	docLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bingboing"}

	// docRussianPhrase represents Doc's phrase in Russian.
	docRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "зайчище"}

	// docSimplifiedChinesePhrase represents Doc's phrase in Simplified Chinese.
	docSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是啦"}

	// docSpanishPhrase represents Doc's phrase in Spanish.
	docSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "y tal"}

	// docTraditionalChinesePhrase represents Doc's phrase in Traditional Chinese.
	docTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是啦"}
)

var (
	// docPhrase represents Doc's phrase in different languages.
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
	// Doc represents the character Doc with his complete information.
	Doc = nook.Villager{
		Character:   docCharacter,
		Personality: personality.Lazy,
		Phrase:      docPhrase}
)
