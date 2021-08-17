package monkey

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
	tammiBirthday = nook.Birthday{
		Day:   23,
		Month: time.April}
)

var (
	tammiCode = nook.Code{
		Value: "mnk03"}
)

var (
	tammiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tammi"}

	tammiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lili"}

	tammiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tammi"}

	tammiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lili"}

	tammiGermanName = nook.Name{
		Language: language.German,
		Value:    "Bonni"}

	tammiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tammi"}

	tammiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エイプリル"}

	tammiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "에이프릴"}

	tammiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tami"}

	tammiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тэмми"}

	tammiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "四月"}

	tammiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tami"}

	tammiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "四月"}
)

var (
	tammiName = nook.Languages{
		language.AmericanEnglish:      tammiAmericanEnglishName,
		language.CanadianFrench:       tammiCanadianFrenchName,
		language.Dutch:                tammiDutchName,
		language.French:               tammiFrenchName,
		language.German:               tammiGermanName,
		language.Italian:              tammiItalianName,
		language.Japanese:             tammiJapaneseName,
		language.Korean:               tammiKoreanName,
		language.LatinAmericanSpanish: tammiLatinAmericanSpanishName,
		language.Russian:              tammiRussianName,
		language.SimplifiedChinese:    tammiSimplifiedChineseName,
		language.Spanish:              tammiSpanishName,
		language.TraditionalChinese:   tammiTraditionalChineseName}
)

var (
	tammiCharacter = nook.Character{
		Animal:   animal.Monkey,
		Birthday: tammiBirthday,
		Code:     tammiCode,
		Key:      character.Tammi,
		Gender:   gender.Female,
		Name:     tammiName}
)

var (
	tammiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "chimpy"}

	tammiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "booof"}

	tammiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "chimpie"}

	tammiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "booof"}

	tammiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "agaga"}

	tammiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "scimmietta"}

	tammiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ワオ"}

	tammiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "와우"}

	tammiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "uki-uki"}

	tammiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "шимпатяга"}

	tammiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哇呜"}

	tammiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "uki-uki"}

	tammiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哇嗚"}
)

var (
	tammiPhrase = nook.Languages{
		language.AmericanEnglish:      tammiAmericanEnglishPhrase,
		language.CanadianFrench:       tammiCanadianFrenchPhrase,
		language.Dutch:                tammiDutchPhrase,
		language.French:               tammiFrenchPhrase,
		language.German:               tammiGermanPhrase,
		language.Italian:              tammiItalianPhrase,
		language.Japanese:             tammiJapanesePhrase,
		language.Korean:               tammiKoreanPhrase,
		language.LatinAmericanSpanish: tammiLatinAmericanSpanishPhrase,
		language.Russian:              tammiRussianPhrase,
		language.SimplifiedChinese:    tammiSimplifiedChinesePhrase,
		language.Spanish:              tammiSpanishPhrase,
		language.TraditionalChinese:   tammiTraditionalChinesePhrase}
)

var (
	Tammi = nook.Villager{
		Character:   tammiCharacter,
		Personality: personality.Peppy,
		Phrase:      tammiPhrase}
)
