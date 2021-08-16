package tiger

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
	tybaltBirthday = nook.Birthday{
		Day:   19,
		Month: time.August}
)

var (
	tybaltCode = nook.Code{
		Value: "tig02"}
)

var (
	tybaltAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tybalt"}

	tybaltCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Jeff"}

	tybaltDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tybalt"}

	tybaltFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jeff"}

	tybaltGermanName = nook.Name{
		Language: language.German,
		Value:    "Arne"}

	tybaltItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ubaldo"}

	tybaltJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハリマオ"}

	tybaltKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "티볼트"}

	tybaltLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Teobaldo"}

	tybaltRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тибальт"}

	tybaltSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "谷丰"}

	tybaltSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Teobaldo"}

	tybaltTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "谷豐"}
)

var (
	tybaltName = nook.Languages{
		language.AmericanEnglish:      tybaltAmericanEnglishName,
		language.CanadianFrench:       tybaltCanadianFrenchName,
		language.Dutch:                tybaltDutchName,
		language.French:               tybaltFrenchName,
		language.German:               tybaltGermanName,
		language.Italian:              tybaltItalianName,
		language.Japanese:             tybaltJapaneseName,
		language.Korean:               tybaltKoreanName,
		language.LatinAmericanSpanish: tybaltLatinAmericanSpanishName,
		language.Russian:              tybaltRussianName,
		language.SimplifiedChinese:    tybaltSimplifiedChineseName,
		language.Spanish:              tybaltSpanishName,
		language.TraditionalChinese:   tybaltTraditionalChineseName}
)

var (
	tybaltCharacter = nook.Character{
		Animal:   animal.Tiger,
		Birthday: tybaltBirthday,
		Code:     tybaltCode,
		Key:      character.Tybalt,
		Gender:   gender.Male,
		Name:     tybaltName}
)

var (
	tybaltAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "grrrRAH"}

	tybaltCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grouaaah"}

	tybaltDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "brrrrulll"}

	tybaltFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grouaaah"}

	tybaltGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grrrummel"}

	tybaltItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grrrRAH"}

	tybaltJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だトラ"}

	tybaltKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "으르렁"}

	tybaltLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "raaarruf"}

	tybaltRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "р-рык"}

	tybaltSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "虎虎"}

	tybaltSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bengala"}

	tybaltTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "虎虎"}
)

var (
	tybaltPhrase = nook.Languages{
		language.AmericanEnglish:      tybaltAmericanEnglishName,
		language.CanadianFrench:       tybaltCanadianFrenchName,
		language.Dutch:                tybaltDutchName,
		language.French:               tybaltFrenchName,
		language.German:               tybaltGermanName,
		language.Italian:              tybaltItalianName,
		language.Japanese:             tybaltJapaneseName,
		language.Korean:               tybaltKoreanName,
		language.LatinAmericanSpanish: tybaltLatinAmericanSpanishName,
		language.Russian:              tybaltRussianName,
		language.SimplifiedChinese:    tybaltSimplifiedChineseName,
		language.Spanish:              tybaltSpanishName,
		language.TraditionalChinese:   tybaltTraditionalChineseName}
)

var (
	Tybalt = nook.Villager{
		Character:   tybaltCharacter,
		Personality: personality.Jock,
		Phrase:      tybaltPhrase}
)
