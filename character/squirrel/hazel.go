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
	// hazelBirthday represents Hazel's birthday.
	hazelBirthday = nook.Birthday{
		Day:   30,
		Month: time.August}
)

var (
	// hazelCode represents Hazel's unique code.
	hazelCode = nook.Code{
		Value: "squ18"}
)

var (
	// hazelAmericanEnglishName represents Hazel's name in American English.
	hazelAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hazel"}

	// hazelCanadianFrenchName represents Hazel's name in Canadian French.
	hazelCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pamela"}

	// hazelDutchName represents Hazel's name in Dutch.
	hazelDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hazel"}

	// hazelFrenchName represents Hazel's name in French.
	hazelFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pamela"}

	// hazelGermanName represents Hazel's name in German.
	hazelGermanName = nook.Name{
		Language: language.German,
		Value:    "Nadine"}

	// hazelItalianName represents Hazel's name in Italian.
	hazelItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cigliola"}

	// hazelJapaneseName represents Hazel's name in Japanese.
	hazelJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アイリス"}

	// hazelKoreanName represents Hazel's name in Korean.
	hazelKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아이리스"}

	// hazelLatinAmericanSpanishName represents Hazel's name in Latin American Spanish.
	hazelLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Anacarda"}

	// hazelRussianName represents Hazel's name in Russian.
	hazelRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Хейзел"}

	// hazelSimplifiedChineseName represents Hazel's name in Simplified Chinese.
	hazelSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "艾栗丝"}

	// hazelSpanishName represents Hazel's name in Spanish.
	hazelSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Anacarda"}

	// hazelTraditionalChineseName represents Hazel's name in Traditional Chinese.
	hazelTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "艾栗絲"}
)

var (
	// hazelName represents Hazel's name in different languages.
	hazelName = nook.Languages{
		language.AmericanEnglish:      hazelAmericanEnglishName,
		language.CanadianFrench:       hazelCanadianFrenchName,
		language.Dutch:                hazelDutchName,
		language.French:               hazelFrenchName,
		language.German:               hazelGermanName,
		language.Italian:              hazelItalianName,
		language.Japanese:             hazelJapaneseName,
		language.Korean:               hazelKoreanName,
		language.LatinAmericanSpanish: hazelLatinAmericanSpanishName,
		language.Russian:              hazelRussianName,
		language.SimplifiedChinese:    hazelSimplifiedChineseName,
		language.Spanish:              hazelSpanishName,
		language.TraditionalChinese:   hazelTraditionalChineseName}
)

var (
	// hazelCharacter represents Hazel's character information.
	hazelCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: hazelBirthday,
		Code:     hazelCode,
		Key:      character.Hazel,
		Gender:   gender.Female,
		Name:     hazelName,
		Special:  false}
)

var (
	// hazelAmericanEnglishPhrase represents Hazel's phrase in American English.
	hazelAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "uni-wow"}

	// hazelCanadianFrenchPhrase represents Hazel's phrase in Canadian French.
	hazelCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tac tac"}

	// hazelDutchPhrase represents Hazel's phrase in Dutch.
	hazelDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "brauw"}

	// hazelFrenchPhrase represents Hazel's phrase in French.
	hazelFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tac tac"}

	// hazelGermanPhrase represents Hazel's phrase in German.
	hazelGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kekeke"}

	// hazelItalianPhrase represents Hazel's phrase in Italian.
	hazelItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "blink"}

	// hazelJapanesePhrase represents Hazel's phrase in Japanese.
	hazelJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だもんね"}

	// hazelKoreanPhrase represents Hazel's phrase in Korean.
	hazelKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "알아알아"}

	// hazelLatinAmericanSpanishPhrase represents Hazel's phrase in Latin American Spanish.
	hazelLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tac-tac"}

	// hazelRussianPhrase represents Hazel's phrase in Russian.
	hazelRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "фундук"}

	// hazelSimplifiedChinesePhrase represents Hazel's phrase in Simplified Chinese.
	hazelSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "当然哦"}

	// hazelSpanishPhrase represents Hazel's phrase in Spanish.
	hazelSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "castaña"}

	// hazelTraditionalChinesePhrase represents Hazel's phrase in Traditional Chinese.
	hazelTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "當然囉"}
)

var (
	// hazelPhrase represents Hazel's phrase in different languages.
	hazelPhrase = nook.Languages{
		language.AmericanEnglish:      hazelAmericanEnglishPhrase,
		language.CanadianFrench:       hazelCanadianFrenchPhrase,
		language.Dutch:                hazelDutchPhrase,
		language.French:               hazelFrenchPhrase,
		language.German:               hazelGermanPhrase,
		language.Italian:              hazelItalianPhrase,
		language.Japanese:             hazelJapanesePhrase,
		language.Korean:               hazelKoreanPhrase,
		language.LatinAmericanSpanish: hazelLatinAmericanSpanishPhrase,
		language.Russian:              hazelRussianPhrase,
		language.SimplifiedChinese:    hazelSimplifiedChinesePhrase,
		language.Spanish:              hazelSpanishPhrase,
		language.TraditionalChinese:   hazelTraditionalChinesePhrase}
)

var (
	// Hazel represents the character Hazel with her complete information.
	Hazel = nook.Villager{
		Character:   hazelCharacter,
		Personality: personality.BigSister,
		Phrase:      hazelPhrase}
)
