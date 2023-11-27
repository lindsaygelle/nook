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
	// ohareBirthday represents O'Hare's birthday.
	ohareBirthday = nook.Birthday{
		Day:   24,
		Month: time.July}
)

var (
	// ohareCode represents O'Hare's unique code.
	ohareCode = nook.Code{
		Value: "rbt15"}
)

var (
	// ohareAmericanEnglishName represents O'Hare's name in American English.
	ohareAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "O'Hare"}

	// ohareCanadianFrenchName represents O'Hare's name in Canadian French.
	ohareCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Civet"}

	// ohareDutchName represents O'Hare's name in Dutch.
	ohareDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "O'Hare"}

	// ohareFrenchName represents O'Hare's name in French.
	ohareFrenchName = nook.Name{
		Language: language.French,
		Value:    "Civet"}

	// ohareGermanName represents O'Hare's name in German.
	ohareGermanName = nook.Name{
		Language: language.German,
		Value:    "Nico"}

	// ohareItalianName represents O'Hare's name in Italian.
	ohareItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fritz"}

	// ohareJapaneseName represents O'Hare's name in Japanese.
	ohareJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サントス"}

	// ohareKoreanName represents O'Hare's name in Korean.
	ohareKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "산토스"}

	// ohareLatinAmericanSpanishName represents O'Hare's name in Latin American Spanish.
	ohareLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ernesto"}

	// ohareRussianName represents O'Hare's name in Russian.
	ohareRussianName = nook.Name{
		Language: language.Russian,
		Value:    "О'Хэар"}

	// ohareSimplifiedChineseName represents O'Hare's name in Simplified Chinese.
	ohareSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "桑托斯"}

	// ohareSpanishName represents O'Hare's name in Spanish.
	ohareSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ernesto"}

	// ohareTraditionalChineseName represents O'Hare's name in Traditional Chinese.
	ohareTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "桑托斯"}
)

var (
	// ohareName represents O'Hare's name in different languages.
	ohareName = nook.Languages{
		language.AmericanEnglish:      ohareAmericanEnglishName,
		language.CanadianFrench:       ohareCanadianFrenchName,
		language.Dutch:                ohareDutchName,
		language.French:               ohareFrenchName,
		language.German:               ohareGermanName,
		language.Italian:              ohareItalianName,
		language.Japanese:             ohareJapaneseName,
		language.Korean:               ohareKoreanName,
		language.LatinAmericanSpanish: ohareLatinAmericanSpanishName,
		language.Russian:              ohareRussianName,
		language.SimplifiedChinese:    ohareSimplifiedChineseName,
		language.Spanish:              ohareSpanishName,
		language.TraditionalChinese:   ohareTraditionalChineseName}
)

var (
	// ohareCharacter represents O'Hare's character information.
	ohareCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: ohareBirthday,
		Code:     ohareCode,
		Key:      character.OHare,
		Gender:   gender.Male,
		Name:     ohareName,
		Special:  false}
)

var (
	// ohareAmericanEnglishPhrase represents O'Hare's phrase in American English.
	ohareAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "amigo"}

	// ohareCanadianFrenchPhrase represents O'Hare's phrase in Canadian French.
	ohareCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "amigo"}

	// ohareDutchPhrase represents O'Hare's phrase in Dutch.
	ohareDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "amigo"}

	// ohareFrenchPhrase represents O'Hare's phrase in French.
	ohareFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pécaïre"}

	// ohareGermanPhrase represents O'Hare's phrase in German.
	ohareGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "amigo"}

	// ohareItalianPhrase represents O'Hare's phrase in Italian.
	ohareItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "aquesí"}

	// ohareJapanesePhrase represents O'Hare's phrase in Japanese.
	ohareJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "アミーゴ"}

	// ohareKoreanPhrase represents O'Hare's phrase in Korean.
	ohareKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아미고"}

	// ohareLatinAmericanSpanishPhrase represents O'Hare's phrase in Latin American Spanish.
	ohareLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tralarí"}

	// ohareRussianPhrase represents O'Hare's phrase in Russian.
	ohareRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "амиго"}

	// ohareSimplifiedChinesePhrase represents O'Hare's phrase in Simplified Chinese.
	ohareSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "朋友"}

	// ohareSpanishPhrase represents O'Hare's phrase in Spanish.
	ohareSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tralarí"}

	// ohareTraditionalChinesePhrase represents O'Hare's phrase in Traditional Chinese.
	ohareTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "朋友"}
)

var (
	// oharePhrase represents O'Hare's phrases in different languages.
	oharePhrase = nook.Languages{
		language.AmericanEnglish:      ohareAmericanEnglishPhrase,
		language.CanadianFrench:       ohareCanadianFrenchPhrase,
		language.Dutch:                ohareDutchPhrase,
		language.French:               ohareFrenchPhrase,
		language.German:               ohareGermanPhrase,
		language.Italian:              ohareItalianPhrase,
		language.Japanese:             ohareJapanesePhrase,
		language.Korean:               ohareKoreanPhrase,
		language.LatinAmericanSpanish: ohareLatinAmericanSpanishPhrase,
		language.Russian:              ohareRussianPhrase,
		language.SimplifiedChinese:    ohareSimplifiedChinesePhrase,
		language.Spanish:              ohareSpanishPhrase,
		language.TraditionalChinese:   ohareTraditionalChinesePhrase}
)

var (
	// OHare represents the character O'Hare with his complete information.
	OHare = nook.Villager{
		Character:   ohareCharacter,
		Personality: personality.Smug,
		Phrase:      oharePhrase}
)
