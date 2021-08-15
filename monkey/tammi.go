package monkey

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Lilibooof"}

	tammiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tammichimpie"}

	tammiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lilibooof"}

	tammiGermanName = nook.Name{
		Language: language.German,
		Value:    "Bonniagaga"}

	tammiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tammiscimmietta"}

	tammiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エイプリルワオ"}

	tammiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "에이프릴와우"}

	tammiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tamiuki-uki"}

	tammiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тэммишимпатяга"}

	tammiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "四月哇呜"}

	tammiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tamiuki-uki"}

	tammiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "四月哇嗚"}
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
		Animal:   Monkey,
		Birthday: tammiBirthday,
		Code:     tammiCode,
		Gender:   nook.Female,
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
	Tammi = nook.Villager{
		Character:   tammiCharacter,
		Personality: nook.Peppy,
		Phrase:      tammiPhrase}
)
