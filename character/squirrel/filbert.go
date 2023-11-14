package squirrel

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
	// filbertBirthday represents Filbert's birthday (June 3rd).
	filbertBirthday = nook.Birthday{
		Day:   3,
		Month: time.June}
)

var (
	// filbertCode represents Filbert's unique code ("squ02").
	filbertCode = nook.Code{
		Value: "squ02"}
)

var (
	// filbertAmericanEnglishName represents Filbert's name in American English.
	filbertAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Filbert"}

	// filbertCanadianFrenchName represents Filbert's name in Canadian French.
	filbertCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Filibert"}

	// filbertDutchName represents Filbert's name in Dutch.
	filbertDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Filbert"}

	// filbertFrenchName represents Filbert's name in French.
	filbertFrenchName = nook.Name{
		Language: language.French,
		Value:    "Filibert"}

	// filbertGermanName represents Filbert's name in German.
	filbertGermanName = nook.Name{
		Language: language.German,
		Value:    "Felix"}

	// filbertItalianName represents Filbert's name in Italian.
	filbertItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Scriccio"}

	// filbertJapaneseName represents Filbert's name in Japanese.
	filbertJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リッキー"}

	// filbertKoreanName represents Filbert's name in Korean.
	filbertKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "리키"}

	// filbertLatinAmericanSpanishName represents Filbert's name in Latin American Spanish.
	filbertLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Filberto"}

	// filbertRussianName represents Filbert's name in Russian.
	filbertRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Филберт"}

	// filbertSimplifiedChineseName represents Filbert's name in Simplified Chinese.
	filbertSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "瑞奇"}

	// filbertSpanishName represents Filbert's name in Spanish.
	filbertSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Filberto"}

	// filbertTraditionalChineseName represents Filbert's name in Traditional Chinese.
	filbertTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "瑞奇"}
)

var (
	// filbertName represents Filbert's name in different languages.
	filbertName = nook.Languages{
		language.AmericanEnglish:      filbertAmericanEnglishName,
		language.CanadianFrench:       filbertCanadianFrenchName,
		language.Dutch:                filbertDutchName,
		language.French:               filbertFrenchName,
		language.German:               filbertGermanName,
		language.Italian:              filbertItalianName,
		language.Japanese:             filbertJapaneseName,
		language.Korean:               filbertKoreanName,
		language.LatinAmericanSpanish: filbertLatinAmericanSpanishName,
		language.Russian:              filbertRussianName,
		language.SimplifiedChinese:    filbertSimplifiedChineseName,
		language.Spanish:              filbertSpanishName,
		language.TraditionalChinese:   filbertTraditionalChineseName}
)

var (
	// filbertCharacter represents Filbert's character information.
	filbertCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: filbertBirthday,
		Code:     filbertCode,
		Key:      character.Filbert,
		Gender:   gender.Male,
		Name:     filbertName,
		Special:  false}
)

var (
	// filbertAmericanEnglishPhrase represents Filbert's phrase in American English.
	filbertAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bucko"}

	// filbertCanadianFrenchPhrase represents Filbert's phrase in Canadian French.
	filbertCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "gogo"}

	// filbertDutchPhrase represents Filbert's phrase in Dutch.
	filbertDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "vent"}

	// filbertFrenchPhrase represents Filbert's phrase in French.
	filbertFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "gogo"}

	// filbertGermanPhrase represents Filbert's phrase in German.
	filbertGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "manno"}

	// filbertItalianPhrase represents Filbert's phrase in Italian.
	filbertItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gnic gnic"}

	// filbertJapanesePhrase represents Filbert's phrase in Japanese.
	filbertJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でしゅ"}

	// filbertKoreanPhrase represents Filbert's phrase in Korean.
	filbertKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "예용"}

	// filbertLatinAmericanSpanishPhrase represents Filbert's phrase in Latin American Spanish.
	filbertLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "guirlache"}

	// filbertRussianPhrase represents Filbert's phrase in Russian.
	filbertRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "дружище"}

	// filbertSimplifiedChinesePhrase represents Filbert's phrase in Simplified Chinese.
	filbertSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是哦"}

	// filbertSpanishPhrase represents Filbert's phrase in Spanish.
	filbertSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "guirlache"}

	// filbertTraditionalChinesePhrase represents Filbert's phrase in Traditional Chinese.
	filbertTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是囉"}
)

var (
	// filbertPhrase represents Filbert's phrases in different languages.
	filbertPhrase = nook.Languages{
		language.AmericanEnglish:      filbertAmericanEnglishPhrase,
		language.CanadianFrench:       filbertCanadianFrenchPhrase,
		language.Dutch:                filbertDutchPhrase,
		language.French:               filbertFrenchPhrase,
		language.German:               filbertGermanPhrase,
		language.Italian:              filbertItalianPhrase,
		language.Japanese:             filbertJapanesePhrase,
		language.Korean:               filbertKoreanPhrase,
		language.LatinAmericanSpanish: filbertLatinAmericanSpanishPhrase,
		language.Russian:              filbertRussianPhrase,
		language.SimplifiedChinese:    filbertSimplifiedChinesePhrase,
		language.Spanish:              filbertSpanishPhrase,
		language.TraditionalChinese:   filbertTraditionalChinesePhrase}
)

var (
	// Filbert represents the character Filbert with his complete information.
	Filbert = nook.Villager{
		Character:   filbertCharacter,
		Personality: personality.Lazy,
		Phrase:      filbertPhrase}
)
