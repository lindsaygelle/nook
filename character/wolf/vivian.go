package wolf

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
	// vivianBirthday represents Vivian's birthday (January 26th).
	vivianBirthday = nook.Birthday{
		Day:   26,
		Month: time.January}
)

var (
	// vivianCode represents Vivian's unique code ("wol08").
	vivianCode = nook.Code{
		Value: "wol08"}
)

var (
	// vivianAmericanEnglishName represents Vivian's name in American English.
	vivianAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Vivian"}

	// vivianCanadianFrenchName represents Vivian's name in Canadian French.
	vivianCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Viviane"}

	// vivianDutchName represents Vivian's name in Dutch.
	vivianDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Vivian"}

	// vivianFrenchName represents Vivian's name in French.
	vivianFrenchName = nook.Name{
		Language: language.French,
		Value:    "Viviane"}

	// vivianGermanName represents Vivian's name in German.
	vivianGermanName = nook.Name{
		Language: language.German,
		Value:    "Viviane"}

	// vivianItalianName represents Vivian's name in Italian.
	vivianItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Viviana"}

	// vivianJapaneseName represents Vivian's name in Japanese.
	vivianJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヴァネッサ"}

	// vivianKoreanName represents Vivian's name in Korean.
	vivianKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "바네사"}

	// vivianLatinAmericanSpanishName represents Vivian's name in Latin American Spanish.
	vivianLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Viviana"}

	// vivianRussianName represents Vivian's name in Russian.
	vivianRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Вивьен"}

	// vivianSimplifiedChineseName represents Vivian's name in Simplified Chinese.
	vivianSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "范妮沙"}

	// vivianSpanishName represents Vivian's name in Spanish.
	vivianSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Viviana"}

	// vivianTraditionalChineseName represents Vivian's name in Traditional Chinese.
	vivianTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "范妮莎"}
)

var (
	// vivianName represents Vivian's name in different languages.
	vivianName = nook.Languages{
		language.AmericanEnglish:      vivianAmericanEnglishName,
		language.CanadianFrench:       vivianCanadianFrenchName,
		language.Dutch:                vivianDutchName,
		language.French:               vivianFrenchName,
		language.German:               vivianGermanName,
		language.Italian:              vivianItalianName,
		language.Japanese:             vivianJapaneseName,
		language.Korean:               vivianKoreanName,
		language.LatinAmericanSpanish: vivianLatinAmericanSpanishName,
		language.Russian:              vivianRussianName,
		language.SimplifiedChinese:    vivianSimplifiedChineseName,
		language.Spanish:              vivianSpanishName,
		language.TraditionalChinese:   vivianTraditionalChineseName}
)

var (
	// vivianCharacter represents Vivian's character information.
	vivianCharacter = nook.Character{
		Animal:   animal.Wolf,
		Birthday: vivianBirthday,
		Code:     vivianCode,
		Key:      character.Vivian,
		Gender:   gender.Female,
		Name:     vivianName,
		Special:  false}
)

var (
	// vivianAmericanEnglishPhrase represents Vivian's phrase in American English.
	vivianAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "piffle"}

	// vivianCanadianFrenchPhrase represents Vivian's phrase in Canadian French.
	vivianCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "crocroque"}

	// vivianDutchPhrase represents Vivian's phrase in Dutch.
	vivianDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "nonsens"}

	// vivianFrenchPhrase represents Vivian's phrase in French.
	vivianFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "crocroque"}

	// vivianGermanPhrase represents Vivian's phrase in German.
	vivianGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "scholli"}

	// vivianItalianPhrase represents Vivian's phrase in Italian.
	vivianItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "arruuu"}

	// vivianJapanesePhrase represents Vivian's phrase in Japanese.
	vivianJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だわよ"}

	// vivianKoreanPhrase represents Vivian's phrase in Korean.
	vivianKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그러하다"}

	// vivianLatinAmericanSpanishPhrase represents Vivian's phrase in Latin American Spanish.
	vivianLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "aú-aú"}

	// vivianRussianPhrase represents Vivian's phrase in Russian.
	vivianRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "вздор"}

	// vivianSimplifiedChinesePhrase represents Vivian's phrase in Simplified Chinese.
	vivianSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喔呵呵"}

	// vivianSpanishPhrase represents Vivian's phrase in Spanish.
	vivianSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "aú-aú"}

	// vivianTraditionalChinesePhrase represents Vivian's phrase in Traditional Chinese.
	vivianTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喔呵呵"}
)

var (
	// vivianPhrase represents Vivian's phrase in different languages.
	vivianPhrase = nook.Languages{
		language.AmericanEnglish:      vivianAmericanEnglishPhrase,
		language.CanadianFrench:       vivianCanadianFrenchPhrase,
		language.Dutch:                vivianDutchPhrase,
		language.French:               vivianFrenchPhrase,
		language.German:               vivianGermanPhrase,
		language.Italian:              vivianItalianPhrase,
		language.Japanese:             vivianJapanesePhrase,
		language.Korean:               vivianKoreanPhrase,
		language.LatinAmericanSpanish: vivianLatinAmericanSpanishPhrase,
		language.Russian:              vivianRussianPhrase,
		language.SimplifiedChinese:    vivianSimplifiedChinesePhrase,
		language.Spanish:              vivianSpanishPhrase,
		language.TraditionalChinese:   vivianTraditionalChinesePhrase}
)

var (
	// Vivian represents the character Vivian with her complete information.
	Vivian = nook.Villager{
		Character:   vivianCharacter,
		Personality: personality.Snooty,
		Phrase:      vivianPhrase}
)
