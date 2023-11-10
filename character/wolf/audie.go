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
	// audieBirthday represents Audie's birthday (August 31st).
	audieBirthday = nook.Birthday{
		Day:   31,
		Month: time.August}
)

var (
	// audieCode represents Audie's unique code ("wol12").
	audieCode = nook.Code{
		Value: "wol12"}
)

var (
	// audieAmericanEnglishName represents Audie's name in American English.
	audieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Audie"}

	// audieCanadianFrenchName represents Audie's name in Canadian French.
	audieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Monica"}

	// audieDutchName represents Audie's name in Dutch.
	audieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Audie"}

	// audieFrenchName represents Audie's name in French.
	audieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Monica"}

	// audieGermanName represents Audie's name in German.
	audieGermanName = nook.Name{
		Language: language.German,
		Value:    "Katharina"}

	// audieItalianName represents Audie's name in Italian.
	audieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lupilia"}

	// audieJapaneseName represents Audie's name in Japanese.
	audieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "モニカ"}

	// audieKoreanName represents Audie's name in Korean.
	audieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "모니카"}

	// audieLatinAmericanSpanishName represents Audie's name in Latin American Spanish.
	audieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mónica"}

	// audieRussianName represents Audie's name in Russian.
	audieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Оди"}

	// audieSimplifiedChineseName represents Audie's name in Simplified Chinese.
	audieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "莫妮卡"}

	// audieSpanishName represents Audie's name in Spanish.
	audieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mónica"}

	// audieTraditionalChineseName represents Audie's name in Traditional Chinese.
	audieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "莫妮卡"}
)

var (
	// audieName represents Audie's name in different languages.
	audieName = nook.Languages{
		language.AmericanEnglish:      audieAmericanEnglishName,
		language.CanadianFrench:       audieCanadianFrenchName,
		language.Dutch:                audieDutchName,
		language.French:               audieFrenchName,
		language.German:               audieGermanName,
		language.Italian:              audieItalianName,
		language.Japanese:             audieJapaneseName,
		language.Korean:               audieKoreanName,
		language.LatinAmericanSpanish: audieLatinAmericanSpanishName,
		language.Russian:              audieRussianName,
		language.SimplifiedChinese:    audieSimplifiedChineseName,
		language.Spanish:              audieSpanishName,
		language.TraditionalChinese:   audieTraditionalChineseName}
)

var (
	// audieCharacter represents Audie's character information.
	audieCharacter = nook.Character{
		Animal:   animal.Wolf,
		Birthday: audieBirthday,
		Code:     audieCode,
		Key:      character.Audie,
		Gender:   gender.Female,
		Name:     audieName,
		Special:  false}
)

var (
	// audieAmericanEnglishPhrase represents Audie's phrase in American English.
	audieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "foxtrot"}

	// audieCanadianFrenchPhrase represents Audie's phrase in Canadian French.
	audieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "trottine"}

	// audieDutchPhrase represents Audie's phrase in Dutch.
	audieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "foxtrot"}

	// audieFrenchPhrase represents Audie's phrase in French.
	audieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "trottine"}

	// audieGermanPhrase represents Audie's phrase in German.
	audieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "foxtrott"}

	// audieItalianPhrase represents Audie's phrase in Italian.
	audieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "foxtrot"}

	// audieJapanesePhrase represents Audie's phrase in Japanese.
	audieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "アハッ"}

	// audieKoreanPhrase represents Audie's phrase in Korean.
	audieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아하핫"}

	// audieLatinAmericanSpanishPhrase represents Audie's phrase in Latin American Spanish.
	audieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ulalilá"}

	// audieRussianPhrase represents Audie's phrase in Russian.
	audieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "фокстрот"}

	// audieSimplifiedChinesePhrase represents Audie's phrase in Simplified Chinese.
	audieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呀哈"}

	// audieSpanishPhrase represents Audie's phrase in Spanish.
	audieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "maravilla"}

	// audieTraditionalChinesePhrase represents Audie's phrase in Traditional Chinese.
	audieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "呀哈"}
)

var (
	// audiePhrase represents Audie's phrase in different languages.
	audiePhrase = nook.Languages{
		language.AmericanEnglish:      audieAmericanEnglishPhrase,
		language.CanadianFrench:       audieCanadianFrenchPhrase,
		language.Dutch:                audieDutchPhrase,
		language.French:               audieFrenchPhrase,
		language.German:               audieGermanPhrase,
		language.Italian:              audieItalianPhrase,
		language.Japanese:             audieJapanesePhrase,
		language.Korean:               audieKoreanPhrase,
		language.LatinAmericanSpanish: audieLatinAmericanSpanishPhrase,
		language.Russian:              audieRussianPhrase,
		language.SimplifiedChinese:    audieSimplifiedChinesePhrase,
		language.Spanish:              audieSpanishPhrase,
		language.TraditionalChinese:   audieTraditionalChinesePhrase}
)

var (
	// Audie represents the character Audie with her complete information.
	Audie = nook.Villager{
		Character:   audieCharacter,
		Personality: personality.Peppy,
		Phrase:      audiePhrase}
)
