package bull

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
		Value:    "Beubeu"}

	stuDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Stu"}

	stuFrenchName = nook.Name{
		Language: language.French,
		Value:    "Beubeu"}

	stuGermanName = nook.Name{
		Language: language.German,
		Value:    "Carlos"}

	stuItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Carlos"}

	stuJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "モーリス"}

	stuKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "모리스"}

	stuLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Vitorino"}

	stuRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Стью"}

	stuSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "毛乐时"}

	stuSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Vitorino"}

	stuTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "毛樂時"}
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
		Animal:   animal.Bull,
		Birthday: stuBirthday,
		Code:     stuCode,
		Key:      character.Stu,
		Gender:   gender.Male,
		Name:     stuName}
)

var (
	stuAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "moo-dude"}

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
		Personality: personality.Lazy,
		Phrase:      stuPhrase}
)
