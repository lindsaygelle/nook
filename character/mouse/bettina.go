package mouse

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
	bettinaBirthday = nook.Birthday{
		Day:   12,
		Month: time.June}
)

var (
	bettinaCode = nook.Code{
		Value: "mus15"}
)

var (
	bettinaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bettina"}

	bettinaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Sabrina"}

	bettinaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bettina"}

	bettinaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sabrina"}

	bettinaGermanName = nook.Name{
		Language: language.German,
		Value:    "Schoki"}

	bettinaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rattina"}

	bettinaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マルコ"}

	bettinaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마르카"}

	bettinaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ottina"}

	bettinaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Беттина"}

	bettinaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "丸子"}

	bettinaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ottina"}

	bettinaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "丸子"}
)

var (
	bettinaName = nook.Languages{
		language.AmericanEnglish:      bettinaAmericanEnglishName,
		language.CanadianFrench:       bettinaCanadianFrenchName,
		language.Dutch:                bettinaDutchName,
		language.French:               bettinaFrenchName,
		language.German:               bettinaGermanName,
		language.Italian:              bettinaItalianName,
		language.Japanese:             bettinaJapaneseName,
		language.Korean:               bettinaKoreanName,
		language.LatinAmericanSpanish: bettinaLatinAmericanSpanishName,
		language.Russian:              bettinaRussianName,
		language.SimplifiedChinese:    bettinaSimplifiedChineseName,
		language.Spanish:              bettinaSpanishName,
		language.TraditionalChinese:   bettinaTraditionalChineseName}
)

var (
	bettinaCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: bettinaBirthday,
		Code:     bettinaCode,
		Key:      character.Bettina,
		Gender:   gender.Female,
		Name:     bettinaName}
)

var (
	bettinaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "eekers"}

	bettinaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "skouik-ik"}

	bettinaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "piepers"}

	bettinaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "skouik-ik"}

	bettinaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "achja"}

	bettinaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "zink zink"}

	bettinaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですよね"}

	bettinaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "내말이요"}

	bettinaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "escui-cui"}

	bettinaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пи-пи-пи"}

	bettinaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "就说吧"}

	bettinaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "piruleta"}

	bettinaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "就說吧"}
)

var (
	bettinaPhrase = nook.Languages{
		language.AmericanEnglish:      bettinaAmericanEnglishPhrase,
		language.CanadianFrench:       bettinaCanadianFrenchPhrase,
		language.Dutch:                bettinaDutchPhrase,
		language.French:               bettinaFrenchPhrase,
		language.German:               bettinaGermanPhrase,
		language.Italian:              bettinaItalianPhrase,
		language.Japanese:             bettinaJapanesePhrase,
		language.Korean:               bettinaKoreanPhrase,
		language.LatinAmericanSpanish: bettinaLatinAmericanSpanishPhrase,
		language.Russian:              bettinaRussianPhrase,
		language.SimplifiedChinese:    bettinaSimplifiedChinesePhrase,
		language.Spanish:              bettinaSpanishPhrase,
		language.TraditionalChinese:   bettinaTraditionalChinesePhrase}
)

var (
	Bettina = nook.Villager{
		Character:   bettinaCharacter,
		Personality: personality.Normal,
		Phrase:      bettinaPhrase}
)
