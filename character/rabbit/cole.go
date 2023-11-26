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
	// coleBirthday represents Cole's birthday.
	coleBirthday = nook.Birthday{
		Day:   10,
		Month: time.August}
)

var (
	// coleCode represents Cole's unique code.
	coleCode = nook.Code{
		Value: "rbt18"}
)

var (
	// coleAmericanEnglishName represents Cole's name in American English.
	coleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cole"}

	// coleCanadianFrenchName represents Cole's name in Canadian French.
	coleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Épicure"}

	// coleDutchName represents Cole's name in Dutch.
	coleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cole"}

	// coleFrenchName represents Cole's name in French.
	coleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Épicure"}

	// coleGermanName represents Cole's name in German.
	coleGermanName = nook.Name{
		Language: language.German,
		Value:    "Karl"}

	// coleItalianName represents Cole's name in Italian.
	coleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lello"}

	// coleJapaneseName represents Cole's name in Japanese.
	coleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アマミン"}

	// coleKoreanName represents Cole's name in Korean.
	coleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아마민"}

	// coleLatinAmericanSpanishName represents Cole's name in Latin American Spanish.
	coleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nicolás"}

	// coleRussianName represents Cole's name in Russian.
	coleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Коул"}

	// coleSimplifiedChineseName represents Cole's name in Simplified Chinese.
	coleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿甘"}

	// coleSpanishName represents Cole's name in Spanish.
	coleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Nicolás"}

	// coleTraditionalChineseName represents Cole's name in Traditional Chinese.
	coleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿甘"}
)

var (
	// coleName represents Cole's name in different languages.
	coleName = nook.Languages{
		language.AmericanEnglish:      coleAmericanEnglishName,
		language.CanadianFrench:       coleCanadianFrenchName,
		language.Dutch:                coleDutchName,
		language.French:               coleFrenchName,
		language.German:               coleGermanName,
		language.Italian:              coleItalianName,
		language.Japanese:             coleJapaneseName,
		language.Korean:               coleKoreanName,
		language.LatinAmericanSpanish: coleLatinAmericanSpanishName,
		language.Russian:              coleRussianName,
		language.SimplifiedChinese:    coleSimplifiedChineseName,
		language.Spanish:              coleSpanishName,
		language.TraditionalChinese:   coleTraditionalChineseName}
)

var (
	// coleCharacter represents Cole's character information.
	coleCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: coleBirthday,
		Code:     coleCode,
		Key:      character.Cole,
		Gender:   gender.Male,
		Name:     coleName,
		Special:  false}
)

var (
	// coleAmericanEnglishPhrase represents Cole's phrase in American English.
	coleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "duuude"}

	// coleCanadianFrenchPhrase represents Cole's phrase in Canadian French.
	coleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chaaauuud"}

	// coleDutchPhrase represents Cole's phrase in Dutch.
	coleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "gaaast"}

	// coleFrenchPhrase represents Cole's phrase in French.
	coleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "chaaauuud"}

	// coleGermanPhrase represents Cole's phrase in German.
	coleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mümml"}

	// coleItalianPhrase represents Cole's phrase in Italian.
	coleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "niglio"}

	// coleJapanesePhrase represents Cole's phrase in Japanese.
	coleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "そうさね"}

	// coleKoreanPhrase represents Cole's phrase in Korean.
	coleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "나도나도"}

	// coleLatinAmericanSpanishPhrase represents Cole's phrase in Latin American Spanish.
	coleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "muyayo"}

	// coleRussianPhrase represents Cole's phrase in Russian.
	coleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "зайчелло"}

	// coleSimplifiedChinesePhrase represents Cole's phrase in Simplified Chinese.
	coleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "就是如此"}

	// coleSpanishPhrase represents Cole's phrase in Spanish.
	coleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "muyayo"}

	// coleTraditionalChinesePhrase represents Cole's phrase in Traditional Chinese.
	coleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "就是如此"}
)

var (
	// colePhrase represents Cole's phrase in different languages.
	colePhrase = nook.Languages{
		language.AmericanEnglish:      coleAmericanEnglishPhrase,
		language.CanadianFrench:       coleCanadianFrenchPhrase,
		language.Dutch:                coleDutchPhrase,
		language.French:               coleFrenchPhrase,
		language.German:               coleGermanPhrase,
		language.Italian:              coleItalianPhrase,
		language.Japanese:             coleJapanesePhrase,
		language.Korean:               coleKoreanPhrase,
		language.LatinAmericanSpanish: coleLatinAmericanSpanishPhrase,
		language.Russian:              coleRussianPhrase,
		language.SimplifiedChinese:    coleSimplifiedChinesePhrase,
		language.Spanish:              coleSpanishPhrase,
		language.TraditionalChinese:   coleTraditionalChinesePhrase}
)

var (
	// Cole represents the character Cole with his complete information.
	Cole = nook.Villager{
		Character:   coleCharacter,
		Personality: personality.Lazy,
		Phrase:      colePhrase}
)
