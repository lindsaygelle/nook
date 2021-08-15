package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	tutuBirthday = nook.Birthday{
		Day:   15,
		Month: time.September}
)

var (
	tutuCode = nook.Code{
		Value: "bea07"}
)

var (
	tutuAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tutu"}

	tutuCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tutupic pic"}

	tutuDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tutuhoningbij"}

	tutuFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tutupic pic"}

	tutuGermanName = nook.Name{
		Language: language.German,
		Value:    "Mandyhooonig"}

	tutuItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lolagnifa"}

	tutuJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "れんにゅうファイト"}

	tutuKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "연유사르르"}

	tutuLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tutúfisflash"}

	tutuRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тутуме-е-ед"}

	tutuSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "恋恋加油"}

	tutuSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tutúfisflash"}

	tutuTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "戀戀加油"}
)

var (
	tutuName = nook.Languages{
		language.AmericanEnglish:      tutuAmericanEnglishName,
		language.CanadianFrench:       tutuCanadianFrenchName,
		language.Dutch:                tutuDutchName,
		language.French:               tutuFrenchName,
		language.German:               tutuGermanName,
		language.Italian:              tutuItalianName,
		language.Japanese:             tutuJapaneseName,
		language.Korean:               tutuKoreanName,
		language.LatinAmericanSpanish: tutuLatinAmericanSpanishName,
		language.Russian:              tutuRussianName,
		language.SimplifiedChinese:    tutuSimplifiedChineseName,
		language.Spanish:              tutuSpanishName,
		language.TraditionalChinese:   tutuTraditionalChineseName}
)

var (
	tutuCharacter = nook.Character{
		Animal:   Bear,
		Birthday: tutuBirthday,
		Code:     tutuCode,
		Gender:   nook.Female,
		Name:     tutuName}
)

var (
	tutuAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "twinkles"}

	tutuCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pic pic"}

	tutuDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "honingbij"}

	tutuFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pic pic"}

	tutuGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hooonig"}

	tutuItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gnifa"}

	tutuJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ファイト"}

	tutuKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "사르르"}

	tutuLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fisflash"}

	tutuRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ме-е-ед"}

	tutuSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "加油"}

	tutuSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fisflash"}

	tutuTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "加油"}
)

var (
	tutuPhrase = nook.Languages{
		language.AmericanEnglish:      tutuAmericanEnglishName,
		language.CanadianFrench:       tutuCanadianFrenchName,
		language.Dutch:                tutuDutchName,
		language.French:               tutuFrenchName,
		language.German:               tutuGermanName,
		language.Italian:              tutuItalianName,
		language.Japanese:             tutuJapaneseName,
		language.Korean:               tutuKoreanName,
		language.LatinAmericanSpanish: tutuLatinAmericanSpanishName,
		language.Russian:              tutuRussianName,
		language.SimplifiedChinese:    tutuSimplifiedChineseName,
		language.Spanish:              tutuSpanishName,
		language.TraditionalChinese:   tutuTraditionalChineseName}
)

var (
	Tutu = nook.Villager{
		Character:   tutuCharacter,
		Personality: nook.Peppy,
		Phrase:      tutuPhrase}
)
