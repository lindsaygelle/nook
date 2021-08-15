package rabbit

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Kristinehop"}

	chrissyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chrissysprankel"}

	chrissyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kristinehop"}

	chrissyGermanName = nook.Name{
		Language: language.German,
		Value:    "Anneknuddl"}

	chrissyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Nataschalilalà"}

	chrissyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クリスチーヌリララ"}

	chrissyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "크리스틴샤방"}

	chrissyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lalimimomimo"}

	chrissyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Криссизолотце"}

	chrissySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "克莉琪哩啦啦"}

	chrissySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lalimimomimo"}

	chrissyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "克莉琪哩啦啦"}
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
		Animal:   Rabbit,
		Birthday: chrissyBirthday,
		Code:     chrissyCode,
		Gender:   nook.Female,
		Name:     chrissyName}
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
	Chrissy = nook.Villager{
		Character:   chrissyCharacter,
		Personality: nook.Peppy,
		Phrase:      chrissyPhrase}
)
