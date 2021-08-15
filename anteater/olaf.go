package anteater

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	olafBirthday = nook.Birthday{
		Day:   19,
		Month: time.May}
)

var (
	olafCode = nook.Code{
		Value: "ant09"}
)

var (
	olafAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Olaf"}

	olafCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Blairyesss"}

	olafDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Olaftoet"}

	olafFrenchName = nook.Name{
		Language: language.French,
		Value:    "Blairtartarin"}

	olafGermanName = nook.Name{
		Language: language.German,
		Value:    "Olafmwahaaa"}

	olafItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ettoreswish"}

	olafJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アントニオムーチョ"}

	olafKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "안토니오무쵸"}

	olafLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Olaffruf"}

	olafRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Олафуф-уф"}

	olafSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "安东尼欧啦"}

	olafSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Olaffruf"}

	olafTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "安東尼歐啦"}
)

var (
	olafName = nook.Languages{
		language.AmericanEnglish:      olafAmericanEnglishName,
		language.CanadianFrench:       olafCanadianFrenchName,
		language.Dutch:                olafDutchName,
		language.French:               olafFrenchName,
		language.German:               olafGermanName,
		language.Italian:              olafItalianName,
		language.Japanese:             olafJapaneseName,
		language.Korean:               olafKoreanName,
		language.LatinAmericanSpanish: olafLatinAmericanSpanishName,
		language.Russian:              olafRussianName,
		language.SimplifiedChinese:    olafSimplifiedChineseName,
		language.Spanish:              olafSpanishName,
		language.TraditionalChinese:   olafTraditionalChineseName}
)

var (
	olafCharacter = nook.Character{
		Animal:   Anteater,
		Birthday: olafBirthday,
		Code:     olafCode,
		Gender:   nook.Male,
		Name:     olafName}
)

var (
	olafAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "whiffa"}

	olafCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "yesss"}

	olafDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "toet"}

	olafFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tartarin"}

	olafGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mwahaaa"}

	olafItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "swish"}

	olafJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ムーチョ"}

	olafKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "무쵸"}

	olafLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fruf"}

	olafRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "уф-уф"}

	olafSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "欧啦"}

	olafSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fruf"}

	olafTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "歐啦"}
)

var (
	olafPhrase = nook.Languages{
		language.AmericanEnglish:      olafAmericanEnglishName,
		language.CanadianFrench:       olafCanadianFrenchName,
		language.Dutch:                olafDutchName,
		language.French:               olafFrenchName,
		language.German:               olafGermanName,
		language.Italian:              olafItalianName,
		language.Japanese:             olafJapaneseName,
		language.Korean:               olafKoreanName,
		language.LatinAmericanSpanish: olafLatinAmericanSpanishName,
		language.Russian:              olafRussianName,
		language.SimplifiedChinese:    olafSimplifiedChineseName,
		language.Spanish:              olafSpanishName,
		language.TraditionalChinese:   olafTraditionalChineseName}
)

var (
	Olaf = nook.Villager{
		Character:   olafCharacter,
		Personality: nook.Smug,
		Phrase:      olafPhrase}
)
