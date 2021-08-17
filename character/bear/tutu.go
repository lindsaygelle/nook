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
		Value:    "Tutu"}

	tutuDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tutu"}

	tutuFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tutu"}

	tutuGermanName = nook.Name{
		Language: language.German,
		Value:    "Mandy"}

	tutuItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lola"}

	tutuJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "れんにゅう"}

	tutuKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "연유"}

	tutuLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tutú"}

	tutuRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Туту"}

	tutuSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "恋恋"}

	tutuSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tutú"}

	tutuTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "戀戀"}
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
		Animal:   animal.Bear,
		Birthday: tutuBirthday,
		Code:     tutuCode,
		Key:      character.Tutu,
		Gender:   gender.Female,
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
		language.AmericanEnglish:      tutuAmericanEnglishPhrase,
		language.CanadianFrench:       tutuCanadianFrenchPhrase,
		language.Dutch:                tutuDutchPhrase,
		language.French:               tutuFrenchPhrase,
		language.German:               tutuGermanPhrase,
		language.Italian:              tutuItalianPhrase,
		language.Japanese:             tutuJapanesePhrase,
		language.Korean:               tutuKoreanPhrase,
		language.LatinAmericanSpanish: tutuLatinAmericanSpanishPhrase,
		language.Russian:              tutuRussianPhrase,
		language.SimplifiedChinese:    tutuSimplifiedChinesePhrase,
		language.Spanish:              tutuSpanishPhrase,
		language.TraditionalChinese:   tutuTraditionalChinesePhrase}
)

var (
	Tutu = nook.Villager{
		Character:   tutuCharacter,
		Personality: personality.Peppy,
		Phrase:      tutuPhrase}
)
