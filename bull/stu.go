package bull

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	stuBirthday = nook.Birthday{
		Day:   20,
		Month: time.April}
)

var (
	stuCode = nook.Code{
		Value: "bul03"}
)

var (
	stuAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Stu"}

	stuCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Beubeumeuh quoi"}

	stuDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Stuboe-jend"}

	stuFrenchName = nook.Name{
		Language: language.French,
		Value:    "Beubeumeuh quoi"}

	stuGermanName = nook.Name{
		Language: language.German,
		Value:    "Carlosmuhuhu"}

	stuItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Carlosciaumù"}

	stuJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "モーリスだなっす"}

	stuKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "모리스알것소"}

	stuLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Vitorinoolé y olé"}

	stuRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Стьюму-у"}

	stuSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "毛乐时奔驰"}

	stuSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Vitorinoolé y olé"}

	stuTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "毛樂時奔馳"}
)

var (
	stuName = nook.Languages{
		language.AmericanEnglish:      stuAmericanEnglishName,
		language.CanadianFrench:       stuCanadianFrenchName,
		language.Dutch:                stuDutchName,
		language.French:               stuFrenchName,
		language.German:               stuGermanName,
		language.Italian:              stuItalianName,
		language.Japanese:             stuJapaneseName,
		language.Korean:               stuKoreanName,
		language.LatinAmericanSpanish: stuLatinAmericanSpanishName,
		language.Russian:              stuRussianName,
		language.SimplifiedChinese:    stuSimplifiedChineseName,
		language.Spanish:              stuSpanishName,
		language.TraditionalChinese:   stuTraditionalChineseName}
)

var (
	stuCharacter = nook.Character{
		Animal:   Bull,
		Birthday: stuBirthday,
		Code:     stuCode,
		Gender:   nook.Male,
		Name:     stuName}
)

var (
	stuAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "moo-dudemrooooo"}

	stuCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "meuh quoi"}

	stuDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "boe-jend"}

	stuFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "meuh quoi"}

	stuGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "muhuhu"}

	stuItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ciaumù"}

	stuJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だなっす"}

	stuKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "알것소"}

	stuLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "olé y olé"}

	stuRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "му-у"}

	stuSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "奔驰"}

	stuSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "olé y olé"}

	stuTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "奔馳"}
)

var (
	stuPhrase = nook.Languages{
		language.AmericanEnglish:      stuAmericanEnglishName,
		language.CanadianFrench:       stuCanadianFrenchName,
		language.Dutch:                stuDutchName,
		language.French:               stuFrenchName,
		language.German:               stuGermanName,
		language.Italian:              stuItalianName,
		language.Japanese:             stuJapaneseName,
		language.Korean:               stuKoreanName,
		language.LatinAmericanSpanish: stuLatinAmericanSpanishName,
		language.Russian:              stuRussianName,
		language.SimplifiedChinese:    stuSimplifiedChineseName,
		language.Spanish:              stuSpanishName,
		language.TraditionalChinese:   stuTraditionalChineseName}
)

var (
	Stu = nook.Villager{
		Character:   stuCharacter,
		Personality: nook.Lazy,
		Phrase:      stuPhrase}
)
