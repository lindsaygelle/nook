package duck

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	derwinBirthday = nook.Birthday{
		Day:   25,
		Month: time.May}
)

var (
	derwinCode = nook.Code{
		Value: "duk08"}
)

var (
	derwinAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Derwin"}

	derwinCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Profderrrr"}

	derwinDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Derwinergo"}

	derwinFrenchName = nook.Name{
		Language: language.French,
		Value:    "Profderrrr"}

	derwinGermanName = nook.Name{
		Language: language.German,
		Value:    "Erwinkrook"}

	derwinItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Mikequaccolo"}

	derwinJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ボンね！ママ"}

	derwinKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "봉엄마"}

	derwinLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Torcuatocuato"}

	derwinRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дервинкррря"}

	derwinSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿坊喂！妈妈"}

	derwinSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Torcuatocuato"}

	derwinTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿坊喂！媽媽"}
)

var (
	derwinName = nook.Languages{
		language.AmericanEnglish:      derwinAmericanEnglishName,
		language.CanadianFrench:       derwinCanadianFrenchName,
		language.Dutch:                derwinDutchName,
		language.French:               derwinFrenchName,
		language.German:               derwinGermanName,
		language.Italian:              derwinItalianName,
		language.Japanese:             derwinJapaneseName,
		language.Korean:               derwinKoreanName,
		language.LatinAmericanSpanish: derwinLatinAmericanSpanishName,
		language.Russian:              derwinRussianName,
		language.SimplifiedChinese:    derwinSimplifiedChineseName,
		language.Spanish:              derwinSpanishName,
		language.TraditionalChinese:   derwinTraditionalChineseName}
)

var (
	derwinCharacter = nook.Character{
		Animal:   Duck,
		Birthday: derwinBirthday,
		Code:     derwinCode,
		Gender:   nook.Male,
		Name:     derwinName}
)

var (
	derwinAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "derrrr"}

	derwinCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "derrrr"}

	derwinDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ergo"}

	derwinFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "derrrr"}

	derwinGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "krook"}

	derwinItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quaccolo"}

	derwinJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ね！ママ"}

	derwinKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "엄마"}

	derwinLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cuato"}

	derwinRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кррря"}

	derwinSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喂！妈妈"}

	derwinSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cuato"}

	derwinTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喂！媽媽"}
)

var (
	derwinPhrase = nook.Languages{
		language.AmericanEnglish:      derwinAmericanEnglishName,
		language.CanadianFrench:       derwinCanadianFrenchName,
		language.Dutch:                derwinDutchName,
		language.French:               derwinFrenchName,
		language.German:               derwinGermanName,
		language.Italian:              derwinItalianName,
		language.Japanese:             derwinJapaneseName,
		language.Korean:               derwinKoreanName,
		language.LatinAmericanSpanish: derwinLatinAmericanSpanishName,
		language.Russian:              derwinRussianName,
		language.SimplifiedChinese:    derwinSimplifiedChineseName,
		language.Spanish:              derwinSpanishName,
		language.TraditionalChinese:   derwinTraditionalChineseName}
)

var (
	Derwin = nook.Villager{
		Character:   derwinCharacter,
		Personality: nook.Lazy,
		Phrase:      derwinPhrase}
)
