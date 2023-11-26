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
	// chrissyBirthday represents Chrissy's birthday.
	chrissyBirthday = nook.Birthday{
		Day:   28,
		Month: time.August}
)

var (
	// chrissyCode represents Chrissy's unique code.
	chrissyCode = nook.Code{
		Value: "rbt13"}
)

var (
	// chrissyAmericanEnglishName represents Chrissy's name in American English.
	chrissyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chrissy"}

	// chrissyCanadianFrenchName represents Chrissy's name in Canadian French.
	chrissyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kristine"}

	// chrissyDutchName represents Chrissy's name in Dutch.
	chrissyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chrissy"}

	// chrissyFrenchName represents Chrissy's name in French.
	chrissyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kristine"}

	// chrissyGermanName represents Chrissy's name in German.
	chrissyGermanName = nook.Name{
		Language: language.German,
		Value:    "Anne"}

	// chrissyItalianName represents Chrissy's name in Italian.
	chrissyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Natascha"}

	// chrissyJapaneseName represents Chrissy's name in Japanese.
	chrissyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クリスチーヌ"}

	// chrissyKoreanName represents Chrissy's name in Korean.
	chrissyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "크리스틴"}

	// chrissyLatinAmericanSpanishName represents Chrissy's name in Latin American Spanish.
	chrissyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lali"}

	// chrissyRussianName represents Chrissy's name in Russian.
	chrissyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Крисси"}

	// chrissySimplifiedChineseName represents Chrissy's name in Simplified Chinese.
	chrissySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "克莉琪"}

	// chrissySpanishName represents Chrissy's name in Spanish.
	chrissySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lali"}

	// chrissyTraditionalChineseName represents Chrissy's name in Traditional Chinese.
	chrissyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "克莉琪"}
)

var (
	// chrissyName represents Chrissy's name in different languages.
	chrissyName = nook.Languages{
		language.AmericanEnglish:      chrissyAmericanEnglishName,
		language.CanadianFrench:       chrissyCanadianFrenchName,
		language.Dutch:                chrissyDutchName,
		language.French:               chrissyFrenchName,
		language.German:               chrissyGermanName,
		language.Italian:              chrissyItalianName,
		language.Japanese:             chrissyJapaneseName,
		language.Korean:               chrissyKoreanName,
		language.LatinAmericanSpanish: chrissyLatinAmericanSpanishName,
		language.Russian:              chrissyRussianName,
		language.SimplifiedChinese:    chrissySimplifiedChineseName,
		language.Spanish:              chrissySpanishName,
		language.TraditionalChinese:   chrissyTraditionalChineseName}
)

var (
	// chrissyCharacter represents Chrissy's character information.
	chrissyCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: chrissyBirthday,
		Code:     chrissyCode,
		Key:      character.Chrissy,
		Gender:   gender.Female,
		Name:     chrissyName,
		Special:  false}
)

var (
	// chrissyAmericanEnglishPhrase represents Chrissy's phrase in American English.
	chrissyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sparkles"}

	// chrissyCanadianFrenchPhrase represents Chrissy's phrase in Canadian French.
	chrissyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "hop"}

	// chrissyDutchPhrase represents Chrissy's phrase in Dutch.
	chrissyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "sprankel"}

	// chrissyFrenchPhrase represents Chrissy's phrase in French.
	chrissyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "hop"}

	// chrissyGermanPhrase represents Chrissy's phrase in German.
	chrissyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knuddl"}

	// chrissyItalianPhrase represents Chrissy's phrase in Italian.
	chrissyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "lilalà"}

	// chrissyJapanesePhrase represents Chrissy's phrase in Japanese.
	chrissyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "リララ"}

	// chrissyKoreanPhrase represents Chrissy's phrase in Korean.
	chrissyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "샤방"}

	// chrissyLatinAmericanSpanishPhrase represents Chrissy's phrase in Latin American Spanish.
	chrissyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "mimomimo"}

	// chrissyRussianPhrase represents Chrissy's phrase in Russian.
	chrissyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "золотце"}

	// chrissySimplifiedChinesePhrase represents Chrissy's phrase in Simplified Chinese.
	chrissySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哩啦啦"}

	// chrissySpanishPhrase represents Chrissy's phrase in Spanish.
	chrissySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mimomimo"}

	// chrissyTraditionalChinesePhrase represents Chrissy's phrase in Traditional Chinese.
	chrissyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哩啦啦"}
)

var (
	// chrissyPhrase represents Chrissy's phrase in different languages.
	chrissyPhrase = nook.Languages{
		language.AmericanEnglish:      chrissyAmericanEnglishPhrase,
		language.CanadianFrench:       chrissyCanadianFrenchPhrase,
		language.Dutch:                chrissyDutchPhrase,
		language.French:               chrissyFrenchPhrase,
		language.German:               chrissyGermanPhrase,
		language.Italian:              chrissyItalianPhrase,
		language.Japanese:             chrissyJapanesePhrase,
		language.Korean:               chrissyKoreanPhrase,
		language.LatinAmericanSpanish: chrissyLatinAmericanSpanishPhrase,
		language.Russian:              chrissyRussianPhrase,
		language.SimplifiedChinese:    chrissySimplifiedChinesePhrase,
		language.Spanish:              chrissySpanishPhrase,
		language.TraditionalChinese:   chrissyTraditionalChinesePhrase}
)

var (
	// Chrissy represents the character Chrissy with her complete information.
	Chrissy = nook.Villager{
		Character:   chrissyCharacter,
		Personality: personality.Peppy,
		Phrase:      chrissyPhrase}
)
