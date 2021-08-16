package dog

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
	butchBirthday = nook.Birthday{
		Day:   1,
		Month: time.November}
)

var (
	butchCode = nook.Code{
		Value: "dog01"}
)

var (
	butchAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Butch"}

	butchCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Avril"}

	butchDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Butch"}

	butchFrenchName = nook.Name{
		Language: language.French,
		Value:    "Avril"}

	butchGermanName = nook.Name{
		Language: language.German,
		Value:    "Hasso"}

	butchItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Attila"}

	butchJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジョン"}

	butchKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "존"}

	butchLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bruno"}

	butchRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бутч"}

	butchSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "约翰"}

	butchSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bruno"}

	butchTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "約翰"}
)

var (
	butchName = nook.Languages{
		language.AmericanEnglish:      butchAmericanEnglishName,
		language.CanadianFrench:       butchCanadianFrenchName,
		language.Dutch:                butchDutchName,
		language.French:               butchFrenchName,
		language.German:               butchGermanName,
		language.Italian:              butchItalianName,
		language.Japanese:             butchJapaneseName,
		language.Korean:               butchKoreanName,
		language.LatinAmericanSpanish: butchLatinAmericanSpanishName,
		language.Russian:              butchRussianName,
		language.SimplifiedChinese:    butchSimplifiedChineseName,
		language.Spanish:              butchSpanishName,
		language.TraditionalChinese:   butchTraditionalChineseName}
)

var (
	butchCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: butchBirthday,
		Code:     butchCode,
		Key:      character.Butch,
		Gender:   gender.Male,
		Name:     butchName}
)

var (
	butchAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ROOOOOWF"}

	butchCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "WROOOOUF"}

	butchDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "BROEEEF"}

	butchFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "WROOOOUF"}

	butchGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "GRRUFF"}

	butchItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ROOOOOF"}

	butchJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ノン"}

	butchKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아니"}

	butchLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "gruf-gruf"}

	butchRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ррр-гав"}

	butchSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "侬"}

	butchSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "gruf-gruf"}

	butchTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "儂"}
)

var (
	butchPhrase = nook.Languages{
		language.AmericanEnglish:      butchAmericanEnglishName,
		language.CanadianFrench:       butchCanadianFrenchName,
		language.Dutch:                butchDutchName,
		language.French:               butchFrenchName,
		language.German:               butchGermanName,
		language.Italian:              butchItalianName,
		language.Japanese:             butchJapaneseName,
		language.Korean:               butchKoreanName,
		language.LatinAmericanSpanish: butchLatinAmericanSpanishName,
		language.Russian:              butchRussianName,
		language.SimplifiedChinese:    butchSimplifiedChineseName,
		language.Spanish:              butchSpanishName,
		language.TraditionalChinese:   butchTraditionalChineseName}
)

var (
	Butch = nook.Villager{
		Character:   butchCharacter,
		Personality: personality.Cranky,
		Phrase:      butchPhrase}
)
