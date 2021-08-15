package mouse

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Sabrinaskouik-ik"}

	bettinaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bettinapiepers"}

	bettinaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sabrinaskouik-ik"}

	bettinaGermanName = nook.Name{
		Language: language.German,
		Value:    "Schokiachja"}

	bettinaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rattinazink zink"}

	bettinaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マルコですよね"}

	bettinaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마르카내말이요"}

	bettinaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ottinaescui-cui"}

	bettinaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Беттинапи-пи-пи"}

	bettinaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "丸子就说吧"}

	bettinaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ottinapiruleta"}

	bettinaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "丸子就說吧"}
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
		Animal:   Mouse,
		Birthday: bettinaBirthday,
		Code:     bettinaCode,
		Gender:   nook.Female,
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
	Bettina = nook.Villager{
		Character:   bettinaCharacter,
		Personality: nook.Normal,
		Phrase:      bettinaPhrase}
)
