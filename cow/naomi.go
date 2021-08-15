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
		Value:    "Maikomerengue"}

	naomiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Naomiloei"}

	naomiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maikopistouille"}

	naomiGermanName = nook.Name{
		Language: language.German,
		Value:    "Jolandajoggi"}

	naomiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Melinamoumou"}

	naomiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハナコセボーン"}

	naomiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "화자메르시"}

	naomiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Reginamumumu"}

	naomiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Наомимуля"}

	naomiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "玉兰这样很好"}

	naomiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Reginamerengue"}

	naomiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "玉蘭這樣很好"}
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
