package bear

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
	beardoBirthday = nook.Birthday{
		Day:   27,
		Month: time.September}
)

var (
	beardoCode = nook.Code{
		Value: "bea13"}
)

var (
	beardoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Beardo"}

	beardoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Eustache"}

	beardoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Beardo"}

	beardoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Eustache"}

	beardoGermanName = nook.Name{
		Language: language.German,
		Value:    "Berthold"}

	beardoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Barnaba"}

	beardoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ベアード"}

	beardoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "베어드"}

	beardoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bernabé"}

	beardoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Беардо"}

	beardoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "熊大叔"}

	beardoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bernabé"}

	beardoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "熊大叔"}
)

var (
	beardoName = nook.Languages{
		language.AmericanEnglish:      beardoAmericanEnglishName,
		language.CanadianFrench:       beardoCanadianFrenchName,
		language.Dutch:                beardoDutchName,
		language.French:               beardoFrenchName,
		language.German:               beardoGermanName,
		language.Italian:              beardoItalianName,
		language.Japanese:             beardoJapaneseName,
		language.Korean:               beardoKoreanName,
		language.LatinAmericanSpanish: beardoLatinAmericanSpanishName,
		language.Russian:              beardoRussianName,
		language.SimplifiedChinese:    beardoSimplifiedChineseName,
		language.Spanish:              beardoSpanishName,
		language.TraditionalChinese:   beardoTraditionalChineseName}
)

var (
	beardoCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: beardoBirthday,
		Code:     beardoCode,
		Key:      character.Beardo,
		Gender:   gender.Male,
		Name:     beardoName}
)

var (
	beardoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "whiskers"}

	beardoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ciao"}

	beardoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snorhaar"}

	beardoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "poilodos"}

	beardoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brrroho"}

	beardoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "andale"}

	beardoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "オッホン"}

	beardoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "오홍홍"}

	beardoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bigotre"}

	beardoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "усы мои"}

	beardoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咳咳"}

	beardoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bigotre"}

	beardoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咳咳"}
)

var (
	beardoPhrase = nook.Languages{
		language.AmericanEnglish:      beardoAmericanEnglishPhrase,
		language.CanadianFrench:       beardoCanadianFrenchPhrase,
		language.Dutch:                beardoDutchPhrase,
		language.French:               beardoFrenchPhrase,
		language.German:               beardoGermanPhrase,
		language.Italian:              beardoItalianPhrase,
		language.Japanese:             beardoJapanesePhrase,
		language.Korean:               beardoKoreanPhrase,
		language.LatinAmericanSpanish: beardoLatinAmericanSpanishPhrase,
		language.Russian:              beardoRussianPhrase,
		language.SimplifiedChinese:    beardoSimplifiedChinesePhrase,
		language.Spanish:              beardoSpanishPhrase,
		language.TraditionalChinese:   beardoTraditionalChinesePhrase}
)

var (
	Beardo = nook.Villager{
		Character:   beardoCharacter,
		Personality: personality.Smug,
		Phrase:      beardoPhrase}
)
