package kangaroo

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	astridBirthday = nook.Birthday{
		Day:   8,
		Month: time.September}
)

var (
	astridCode = nook.Code{
		Value: "kgr05"}
)

var (
	astridAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Astrid"}

	astridCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rhona"}

	astridDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Astrid"}

	astridFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rhona"}

	astridGermanName = nook.Name{
		Language: language.German,
		Value:    "Astrid"}

	astridItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Stella"}

	astridJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キッズ"}

	astridKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "펑키맘"}

	astridLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Astrid"}

	astridRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Астрид"}

	astridSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "童儿"}

	astridSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Astrid"}

	astridTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "童兒"}
)

var (
	astridName = nook.Languages{
		language.AmericanEnglish:      astridAmericanEnglishName,
		language.CanadianFrench:       astridCanadianFrenchName,
		language.Dutch:                astridDutchName,
		language.French:               astridFrenchName,
		language.German:               astridGermanName,
		language.Italian:              astridItalianName,
		language.Japanese:             astridJapaneseName,
		language.Korean:               astridKoreanName,
		language.LatinAmericanSpanish: astridLatinAmericanSpanishName,
		language.Russian:              astridRussianName,
		language.SimplifiedChinese:    astridSimplifiedChineseName,
		language.Spanish:              astridSpanishName,
		language.TraditionalChinese:   astridTraditionalChineseName}
)

var (
	astridCharacter = nook.Character{
		Animal:   Kangaroo,
		Birthday: astridBirthday,
		Code:     astridCode,
		Gender:   nook.Female,
		Name:     astridName}
)

var (
	astridAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "my pet"}

	astridCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon étoile"}

	astridDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "troetel"}

	astridFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon étoile"}

	astridGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schmusi"}

	astridItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "fagotto"}

	astridJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だパンク"}

	astridKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "뭘봐"}

	astridLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "brillabrí"}

	astridRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "лапочка"}

	astridSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "庞克"}

	astridSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mascotita"}

	astridTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "龐克"}
)

var (
	astridPhrase = nook.Languages{
		language.AmericanEnglish:      astridAmericanEnglishName,
		language.CanadianFrench:       astridCanadianFrenchName,
		language.Dutch:                astridDutchName,
		language.French:               astridFrenchName,
		language.German:               astridGermanName,
		language.Italian:              astridItalianName,
		language.Japanese:             astridJapaneseName,
		language.Korean:               astridKoreanName,
		language.LatinAmericanSpanish: astridLatinAmericanSpanishName,
		language.Russian:              astridRussianName,
		language.SimplifiedChinese:    astridSimplifiedChineseName,
		language.Spanish:              astridSpanishName,
		language.TraditionalChinese:   astridTraditionalChineseName}
)

var (
	Astrid = nook.Villager{
		Character:   astridCharacter,
		Personality: nook.Snooty,
		Phrase:      astridPhrase}
)
