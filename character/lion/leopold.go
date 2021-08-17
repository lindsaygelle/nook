package lion

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
	leopoldBirthday = nook.Birthday{
		Day:   14,
		Month: time.August}
)

var (
	leopoldCode = nook.Code{
		Value: "lon04"}
)

var (
	leopoldAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Leopold"}

	leopoldCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Leandro"}

	leopoldDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Leopold"}

	leopoldFrenchName = nook.Name{
		Language: language.French,
		Value:    "Leandro"}

	leopoldGermanName = nook.Name{
		Language: language.German,
		Value:    "Leandro"}

	leopoldItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Leandro"}

	leopoldJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ティーチャー"}

	leopoldKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "티처"}

	leopoldLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Leopoldo"}

	leopoldRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Леопольд"}

	leopoldSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "老狮"}

	leopoldSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Leopoldo"}

	leopoldTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "老獅"}
)

var (
	leopoldName = nook.Languages{
		language.AmericanEnglish:      leopoldAmericanEnglishName,
		language.CanadianFrench:       leopoldCanadianFrenchName,
		language.Dutch:                leopoldDutchName,
		language.French:               leopoldFrenchName,
		language.German:               leopoldGermanName,
		language.Italian:              leopoldItalianName,
		language.Japanese:             leopoldJapaneseName,
		language.Korean:               leopoldKoreanName,
		language.LatinAmericanSpanish: leopoldLatinAmericanSpanishName,
		language.Russian:              leopoldRussianName,
		language.SimplifiedChinese:    leopoldSimplifiedChineseName,
		language.Spanish:              leopoldSpanishName,
		language.TraditionalChinese:   leopoldTraditionalChineseName}
)

var (
	leopoldCharacter = nook.Character{
		Animal:   animal.Lion,
		Birthday: leopoldBirthday,
		Code:     leopoldCode,
		Key:      character.Leopold,
		Gender:   gender.Male,
		Name:     leopoldName}
)

var (
	leopoldAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "lion cub"}

	leopoldCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "minus"}

	leopoldDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "jong"}

	leopoldFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "minus"}

	leopoldGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "löwieh"}

	leopoldItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sonré"}

	leopoldJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "よな"}

	leopoldKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "졸지마"}

	leopoldLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "gro-grau"}

	leopoldRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "левушка"}

	leopoldSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "上课了"}

	leopoldSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fiera"}

	leopoldTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "上課了"}
)

var (
	leopoldPhrase = nook.Languages{
		language.AmericanEnglish:      leopoldAmericanEnglishPhrase,
		language.CanadianFrench:       leopoldCanadianFrenchPhrase,
		language.Dutch:                leopoldDutchPhrase,
		language.French:               leopoldFrenchPhrase,
		language.German:               leopoldGermanPhrase,
		language.Italian:              leopoldItalianPhrase,
		language.Japanese:             leopoldJapanesePhrase,
		language.Korean:               leopoldKoreanPhrase,
		language.LatinAmericanSpanish: leopoldLatinAmericanSpanishPhrase,
		language.Russian:              leopoldRussianPhrase,
		language.SimplifiedChinese:    leopoldSimplifiedChinesePhrase,
		language.Spanish:              leopoldSpanishPhrase,
		language.TraditionalChinese:   leopoldTraditionalChinesePhrase}
)

var (
	Leopold = nook.Villager{
		Character:   leopoldCharacter,
		Personality: personality.Smug,
		Phrase:      leopoldPhrase}
)
