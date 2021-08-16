package tiger

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
	leonardoBirthday = nook.Birthday{
		Day:   15,
		Month: time.May}
)

var (
	leonardoCode = nook.Code{
		Value: "tig04"}
)

var (
	leonardoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Leonardo"}

	leonardoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Dolph"}

	leonardoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Leonardo"}

	leonardoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Dolph"}

	leonardoGermanName = nook.Name{
		Language: language.German,
		Value:    "Tim"}

	leonardoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bengalo"}

	leonardoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヒョウタ"}

	leonardoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "범호"}

	leonardoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Leocadio"}

	leonardoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Леонардо"}

	leonardoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿彪"}

	leonardoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Leocadio"}

	leonardoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿彪"}
)

var (
	leonardoName = nook.Languages{
		language.AmericanEnglish:      leonardoAmericanEnglishName,
		language.CanadianFrench:       leonardoCanadianFrenchName,
		language.Dutch:                leonardoDutchName,
		language.French:               leonardoFrenchName,
		language.German:               leonardoGermanName,
		language.Italian:              leonardoItalianName,
		language.Japanese:             leonardoJapaneseName,
		language.Korean:               leonardoKoreanName,
		language.LatinAmericanSpanish: leonardoLatinAmericanSpanishName,
		language.Russian:              leonardoRussianName,
		language.SimplifiedChinese:    leonardoSimplifiedChineseName,
		language.Spanish:              leonardoSpanishName,
		language.TraditionalChinese:   leonardoTraditionalChineseName}
)

var (
	leonardoCharacter = nook.Character{
		Animal:   animal.Tiger,
		Birthday: leonardoBirthday,
		Code:     leonardoCode,
		Key:      character.Leonardo,
		Gender:   gender.Male,
		Name:     leonardoName}
)

var (
	leonardoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "flexin'"}

	leonardoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mes abdos"}

	leonardoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "flex"}

	leonardoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "rebelote"}

	leonardoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "krrrrrr"}

	leonardoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "joggingrrr"}

	leonardoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "じゃが"}

	leonardoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어흥"}

	leonardoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "macoque"}

	leonardoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "р-разминка"}

	leonardoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "但是"}

	leonardoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "macoque"}

	leonardoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "但是"}
)

var (
	leonardoPhrase = nook.Languages{
		language.AmericanEnglish:      leonardoAmericanEnglishName,
		language.CanadianFrench:       leonardoCanadianFrenchName,
		language.Dutch:                leonardoDutchName,
		language.French:               leonardoFrenchName,
		language.German:               leonardoGermanName,
		language.Italian:              leonardoItalianName,
		language.Japanese:             leonardoJapaneseName,
		language.Korean:               leonardoKoreanName,
		language.LatinAmericanSpanish: leonardoLatinAmericanSpanishName,
		language.Russian:              leonardoRussianName,
		language.SimplifiedChinese:    leonardoSimplifiedChineseName,
		language.Spanish:              leonardoSpanishName,
		language.TraditionalChinese:   leonardoTraditionalChineseName}
)

var (
	Leonardo = nook.Villager{
		Character:   leonardoCharacter,
		Personality: personality.Jock,
		Phrase:      leonardoPhrase}
)
