package cow

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	naomiBirthday = nook.Birthday{
		Day:   28,
		Month: time.February}
)

var (
	naomiCode = nook.Code{
		Value: "cow07"}
)

var (
	naomiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Naomi"}

	naomiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Maiko"}

	naomiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Naomi"}

	naomiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maiko"}

	naomiGermanName = nook.Name{
		Language: language.German,
		Value:    "Jolanda"}

	naomiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Melina"}

	naomiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハナコ"}

	naomiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "화자"}

	naomiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Regina"}

	naomiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Наоми"}

	naomiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "玉兰"}

	naomiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Regina"}

	naomiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "玉蘭"}
)

var (
	naomiName = nook.Languages{
		language.AmericanEnglish:      naomiAmericanEnglishName,
		language.CanadianFrench:       naomiCanadianFrenchName,
		language.Dutch:                naomiDutchName,
		language.French:               naomiFrenchName,
		language.German:               naomiGermanName,
		language.Italian:              naomiItalianName,
		language.Japanese:             naomiJapaneseName,
		language.Korean:               naomiKoreanName,
		language.LatinAmericanSpanish: naomiLatinAmericanSpanishName,
		language.Russian:              naomiRussianName,
		language.SimplifiedChinese:    naomiSimplifiedChineseName,
		language.Spanish:              naomiSpanishName,
		language.TraditionalChinese:   naomiTraditionalChineseName}
)

var (
	naomiCharacter = nook.Character{
		Animal:   Cow,
		Birthday: naomiBirthday,
		Code:     naomiCode,
		Gender:   nook.Female,
		Name:     naomiName}
)

var (
	naomiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "moolah"}

	naomiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "merengue"}

	naomiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "loei"}

	naomiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pistouille"}

	naomiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "joggi"}

	naomiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "moumou"}

	naomiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "セボーン"}

	naomiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "메르시"}

	naomiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "mumumu"}

	naomiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "муля"}

	naomiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "这样很好"}

	naomiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "merengue"}

	naomiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "這樣很好"}
)

var (
	naomiPhrase = nook.Languages{
		language.AmericanEnglish:      naomiAmericanEnglishName,
		language.CanadianFrench:       naomiCanadianFrenchName,
		language.Dutch:                naomiDutchName,
		language.French:               naomiFrenchName,
		language.German:               naomiGermanName,
		language.Italian:              naomiItalianName,
		language.Japanese:             naomiJapaneseName,
		language.Korean:               naomiKoreanName,
		language.LatinAmericanSpanish: naomiLatinAmericanSpanishName,
		language.Russian:              naomiRussianName,
		language.SimplifiedChinese:    naomiSimplifiedChineseName,
		language.Spanish:              naomiSpanishName,
		language.TraditionalChinese:   naomiTraditionalChineseName}
)

var (
	Naomi = nook.Villager{
		Character:   naomiCharacter,
		Personality: nook.Snooty,
		Phrase:      naomiPhrase}
)
