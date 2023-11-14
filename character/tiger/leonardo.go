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
	// leonardoBirthday represents Leonardo's birthday (May 15th).
	leonardoBirthday = nook.Birthday{
		Day:   15,
		Month: time.May}
)

var (
	// leonardoCode represents Leonardo's unique code ("tig04").
	leonardoCode = nook.Code{
		Value: "tig04"}
)

var (
	// leonardoAmericanEnglishName represents Leonardo's name in American English.
	leonardoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Leonardo"}

	// leonardoCanadianFrenchName represents Leonardo's name in Canadian French.
	leonardoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Dolph"}

	// leonardoDutchName represents Leonardo's name in Dutch.
	leonardoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Leonardo"}

	// leonardoFrenchName represents Leonardo's name in French.
	leonardoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Dolph"}

	// leonardoGermanName represents Leonardo's name in German.
	leonardoGermanName = nook.Name{
		Language: language.German,
		Value:    "Tim"}

	// leonardoItalianName represents Leonardo's name in Italian.
	leonardoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bengalo"}

	// leonardoJapaneseName represents Leonardo's name in Japanese.
	leonardoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヒョウタ"}

	// leonardoKoreanName represents Leonardo's name in Korean.
	leonardoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "범호"}

	// leonardoLatinAmericanSpanishName represents Leonardo's name in Latin American Spanish.
	leonardoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Leocadio"}

	// leonardoRussianName represents Leonardo's name in Russian.
	leonardoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Леонардо"}

	// leonardoSimplifiedChineseName represents Leonardo's name in Simplified Chinese.
	leonardoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿彪"}

	// leonardoSpanishName represents Leonardo's name in Spanish.
	leonardoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Leocadio"}

	// leonardoTraditionalChineseName represents Leonardo's name in Traditional Chinese.
	leonardoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿彪"}
)

var (
	// leonardoName represents Leonardo's name in different languages.
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
	// leonardoCharacter represents Leonardo's character information.
	leonardoCharacter = nook.Character{
		Animal:   animal.Tiger,
		Birthday: leonardoBirthday,
		Code:     leonardoCode,
		Key:      character.Leonardo,
		Gender:   gender.Male,
		Name:     leonardoName,
		Special:  false}
)

var (
	// leonardoAmericanEnglishPhrase represents Leonardo's phrase in American English.
	leonardoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "flexin'"}

	// leonardoCanadianFrenchPhrase represents Leonardo's phrase in Canadian French.
	leonardoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mes abdos"}

	// leonardoDutchPhrase represents Leonardo's phrase in Dutch.
	leonardoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "flex"}

	// leonardoFrenchPhrase represents Leonardo's phrase in French.
	leonardoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "rebelote"}

	// leonardoGermanPhrase represents Leonardo's phrase in German.
	leonardoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "krrrrrr"}

	// leonardoItalianPhrase represents Leonardo's phrase in Italian.
	leonardoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "joggingrrr"}

	// leonardoJapanesePhrase represents Leonardo's phrase in Japanese.
	leonardoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "じゃが"}

	// leonardoKoreanPhrase represents Leonardo's phrase in Korean.
	leonardoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어흥"}

	// leonardoLatinAmericanSpanishPhrase represents Leonardo's phrase in Latin American Spanish.
	leonardoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "macoque"}

	// leonardoRussianPhrase represents Leonardo's phrase in Russian.
	leonardoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "р-разминка"}

	// leonardoSimplifiedChinesePhrase represents Leonardo's phrase in Simplified Chinese.
	leonardoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "但是"}

	// leonardoSpanishPhrase represents Leonardo's phrase in Spanish.
	leonardoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "macoque"}

	// leonardoTraditionalChinesePhrase represents Leonardo's phrase in Traditional Chinese.
	leonardoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "但是"}
)

var (
	// leonardoPhrase represents Leonardo's phrase in different languages.
	leonardoPhrase = nook.Languages{
		language.AmericanEnglish:      leonardoAmericanEnglishPhrase,
		language.CanadianFrench:       leonardoCanadianFrenchPhrase,
		language.Dutch:                leonardoDutchPhrase,
		language.French:               leonardoFrenchPhrase,
		language.German:               leonardoGermanPhrase,
		language.Italian:              leonardoItalianPhrase,
		language.Japanese:             leonardoJapanesePhrase,
		language.Korean:               leonardoKoreanPhrase,
		language.LatinAmericanSpanish: leonardoLatinAmericanSpanishPhrase,
		language.Russian:              leonardoRussianPhrase,
		language.SimplifiedChinese:    leonardoSimplifiedChinesePhrase,
		language.Spanish:              leonardoSpanishPhrase,
		language.TraditionalChinese:   leonardoTraditionalChinesePhrase}
)

var (
	// Leonardo represents the character Leonardo with his complete information.
	Leonardo = nook.Villager{
		Character:   leonardoCharacter,
		Personality: personality.Jock,
		Phrase:      leonardoPhrase}
)
