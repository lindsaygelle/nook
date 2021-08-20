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
	chrissyBirthday = nook.Birthday{
		Day:   28,
		Month: time.August}
)

var (
	chrissyCode = nook.Code{
		Value: "rbt13"}
)

var (
	chrissyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chrissy"}

	chrissyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kristine"}

	chrissyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chrissy"}

	chrissyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kristine"}

	chrissyGermanName = nook.Name{
		Language: language.German,
		Value:    "Anne"}

	chrissyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Natascha"}

	chrissyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クリスチーヌ"}

	chrissyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "크리스틴"}

	chrissyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lali"}

	chrissyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Крисси"}

	chrissySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "克莉琪"}

	chrissySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lali"}

	chrissyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "克莉琪"}
)

var (
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
	chrissyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sparkles"}

	chrissyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "hop"}

	chrissyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "sprankel"}

	chrissyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "hop"}

	chrissyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knuddl"}

	chrissyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "lilalà"}

	chrissyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "リララ"}

	chrissyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "샤방"}

	chrissyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "mimomimo"}

	chrissyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "золотце"}

	chrissySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哩啦啦"}

	chrissySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mimomimo"}

	chrissyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哩啦啦"}
)

var (
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
	Chrissy = nook.Villager{
		Character:   chrissyCharacter,
		Personality: personality.Peppy,
		Phrase:      chrissyPhrase}
)
