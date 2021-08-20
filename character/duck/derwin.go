package duck

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
		Value:    "Prof"}

	derwinDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Derwin"}

	derwinFrenchName = nook.Name{
		Language: language.French,
		Value:    "Prof"}

	derwinGermanName = nook.Name{
		Language: language.German,
		Value:    "Erwin"}

	derwinItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Mike"}

	derwinJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ボン"}

	derwinKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "봉"}

	derwinLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Torcuato"}

	derwinRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дервин"}

	derwinSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿坊"}

	derwinSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Torcuato"}

	derwinTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿坊"}
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
		Animal:   animal.Duck,
		Birthday: derwinBirthday,
		Code:     derwinCode,
		Key:      character.Derwin,
		Gender:   gender.Male,
		Name:     derwinName,
		Special:  false}
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
		language.AmericanEnglish:      derwinAmericanEnglishPhrase,
		language.CanadianFrench:       derwinCanadianFrenchPhrase,
		language.Dutch:                derwinDutchPhrase,
		language.French:               derwinFrenchPhrase,
		language.German:               derwinGermanPhrase,
		language.Italian:              derwinItalianPhrase,
		language.Japanese:             derwinJapanesePhrase,
		language.Korean:               derwinKoreanPhrase,
		language.LatinAmericanSpanish: derwinLatinAmericanSpanishPhrase,
		language.Russian:              derwinRussianPhrase,
		language.SimplifiedChinese:    derwinSimplifiedChinesePhrase,
		language.Spanish:              derwinSpanishPhrase,
		language.TraditionalChinese:   derwinTraditionalChinesePhrase}
)

var (
	Derwin = nook.Villager{
		Character:   derwinCharacter,
		Personality: personality.Lazy,
		Phrase:      derwinPhrase}
)
