package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	flashBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	flashCode = nook.Code{
		Value: ""}
)

var (
	flashAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Flash"}

	flashCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	flashDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	flashFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maurice"}

	flashGermanName = nook.Name{
		Language: language.German,
		Value:    "Flo"}

	flashItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Baleno"}

	flashJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "みちる"}

	flashKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	flashLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	flashRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	flashSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	flashSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mauro"}

	flashTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	flashName = nook.Languages{
		language.AmericanEnglish:      flashAmericanEnglishName,
		language.CanadianFrench:       flashCanadianFrenchName,
		language.Dutch:                flashDutchName,
		language.French:               flashFrenchName,
		language.German:               flashGermanName,
		language.Italian:              flashItalianName,
		language.Japanese:             flashJapaneseName,
		language.Korean:               flashKoreanName,
		language.LatinAmericanSpanish: flashLatinAmericanSpanishName,
		language.Russian:              flashRussianName,
		language.SimplifiedChinese:    flashSimplifiedChineseName,
		language.Spanish:              flashSpanishName,
		language.TraditionalChinese:   flashTraditionalChineseName}
)

var (
	flashCharacter = nook.Character{
		Animal:   Bird,
		Birthday: flashBirthday,
		Code:     flashCode,
		Gender:   nook.Male,
		Name:     flashName}
)

var (
	flashAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	flashCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	flashDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	flashFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "cuicui"}

	flashGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schätzchen"}

	flashItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "flash"}

	flashJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "チルチル"}

	flashKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	flashLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	flashRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	flashSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	flashSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cari"}

	flashTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	flashPhrase = nook.Languages{
		language.AmericanEnglish:      flashAmericanEnglishName,
		language.CanadianFrench:       flashCanadianFrenchName,
		language.Dutch:                flashDutchName,
		language.French:               flashFrenchName,
		language.German:               flashGermanName,
		language.Italian:              flashItalianName,
		language.Japanese:             flashJapaneseName,
		language.Korean:               flashKoreanName,
		language.LatinAmericanSpanish: flashLatinAmericanSpanishName,
		language.Russian:              flashRussianName,
		language.SimplifiedChinese:    flashSimplifiedChineseName,
		language.Spanish:              flashSpanishName,
		language.TraditionalChinese:   flashTraditionalChineseName}
)

var (
	Flash = nook.Villager{
		Character:   flashCharacter,
		Personality: nook.Cranky,
		Phrase:      flashPhrase}
)
