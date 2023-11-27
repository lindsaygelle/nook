package sheep

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
	// willowBirthday represents Willow's birthday.
	willowBirthday = nook.Birthday{
		Day:   26,
		Month: time.November}
)

var (
	// willowCode represents Willow's unique code.
	willowCode = nook.Code{
		Value: "shp07"}
)

var (
	// willowAmericanEnglishName represents Willow's name in American English.
	willowAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Willow"}

	// willowCanadianFrenchName represents Willow's name in Canadian French.
	willowCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Maï"}

	// willowDutchName represents Willow's name in Dutch.
	willowDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Willow"}

	// willowFrenchName represents Willow's name in French.
	willowFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maï"}

	// willowGermanName represents Willow's name in German.
	willowGermanName = nook.Name{
		Language: language.German,
		Value:    "Natascha"}

	// willowItalianName represents Willow's name in Italian.
	willowItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Belinda"}

	// willowJapaneseName represents Willow's name in Japanese.
	willowJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マリー"}

	// willowKoreanName represents Willow's name in Korean.
	willowKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마리"}

	// willowLatinAmericanSpanishName represents Willow's name in Latin American Spanish.
	willowLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cuqui"}

	// willowRussianName represents Willow's name in Russian.
	willowRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Уиллоу"}

	// willowSimplifiedChineseName represents Willow's name in Simplified Chinese.
	willowSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "梅丽诺"}

	// willowSpanishName represents Willow's name in Spanish.
	willowSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cuqui"}

	// willowTraditionalChineseName represents Willow's name in Traditional Chinese.
	willowTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "梅麗諾"}
)

var (
	// willowName represents Willow's name in different languages.
	willowName = nook.Languages{
		language.AmericanEnglish:      willowAmericanEnglishName,
		language.CanadianFrench:       willowCanadianFrenchName,
		language.Dutch:                willowDutchName,
		language.French:               willowFrenchName,
		language.German:               willowGermanName,
		language.Italian:              willowItalianName,
		language.Japanese:             willowJapaneseName,
		language.Korean:               willowKoreanName,
		language.LatinAmericanSpanish: willowLatinAmericanSpanishName,
		language.Russian:              willowRussianName,
		language.SimplifiedChinese:    willowSimplifiedChineseName,
		language.Spanish:              willowSpanishName,
		language.TraditionalChinese:   willowTraditionalChineseName}
)

var (
	// willowCharacter represents Willow's character information.
	willowCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: willowBirthday,
		Code:     willowCode,
		Key:      character.Willow,
		Gender:   gender.Female,
		Name:     willowName,
		Special:  false}
)

var (
	// willowAmericanEnglishPhrase represents Willow's phrase in American English.
	willowAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bo peep"}

	// willowCanadianFrenchPhrase represents Willow's phrase in Canadian French.
	willowCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tourni"}

	// willowDutchPhrase represents Willow's phrase in Dutch.
	willowDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "wollebol"}

	// willowFrenchPhrase represents Willow's phrase in French.
	willowFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tourni"}

	// willowGermanPhrase represents Willow's phrase in German.
	willowGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "blö-ö-ök"}

	// willowItalianPhrase represents Willow's phrase in Italian.
	willowItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "babeeebí"}

	// willowJapanesePhrase represents Willow's phrase in Japanese.
	willowJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですのよ"}

	// willowKoreanPhrase represents Willow's phrase in Korean.
	willowKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "됐거든"}

	// willowLatinAmericanSpanishPhrase represents Willow's phrase in Latin American Spanish.
	willowLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "babeeebí"}

	// willowRussianPhrase represents Willow's phrase in Russian.
	willowRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пасемся"}

	// willowSimplifiedChinesePhrase represents Willow's phrase in Simplified Chinese.
	willowSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是的唷"}

	// willowSpanishPhrase represents Willow's phrase in Spanish.
	willowSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "babeeebí"}

	// willowTraditionalChinesePhrase represents Willow's phrase in Traditional Chinese.
	willowTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是的唷"}
)

var (
	// willowPhrase represents Willow's phrases in different languages.
	willowPhrase = nook.Languages{
		language.AmericanEnglish:      willowAmericanEnglishPhrase,
		language.CanadianFrench:       willowCanadianFrenchPhrase,
		language.Dutch:                willowDutchPhrase,
		language.French:               willowFrenchPhrase,
		language.German:               willowGermanPhrase,
		language.Italian:              willowItalianPhrase,
		language.Japanese:             willowJapanesePhrase,
		language.Korean:               willowKoreanPhrase,
		language.LatinAmericanSpanish: willowLatinAmericanSpanishPhrase,
		language.Russian:              willowRussianPhrase,
		language.SimplifiedChinese:    willowSimplifiedChinesePhrase,
		language.Spanish:              willowSpanishPhrase,
		language.TraditionalChinese:   willowTraditionalChinesePhrase}
)

var (
	// Willow represents the character Willow with her complete information.
	Willow = nook.Villager{
		Character:   willowCharacter,
		Personality: personality.Snooty,
		Phrase:      willowPhrase}
)
