package lion

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	mottBirthday = nook.Birthday{
		Day:   10,
		Month: time.July}
)

var (
	mottCode = nook.Code{
		Value: "lon06"}
)

var (
	mottAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mott"}

	mottCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Aimégrif-grif"}

	mottDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Mottrrrrrauw"}

	mottFrenchName = nook.Name{
		Language: language.French,
		Value:    "Aimégrif-grif"}

	mottGermanName = nook.Name{
		Language: language.German,
		Value:    "Leonhardgröhl"}

	mottItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Leopoldgroarrivo"}

	mottJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リックハハハ"}

	mottKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "릭하하하"}

	mottLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jonessapristi"}

	mottRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Моттзоосад"}

	mottSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "李克哈哈哈"}

	mottSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jonessapristi"}

	mottTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "李克哈哈哈"}
)

var (
	mottName = nook.Languages{
		language.AmericanEnglish:      mottAmericanEnglishName,
		language.CanadianFrench:       mottCanadianFrenchName,
		language.Dutch:                mottDutchName,
		language.French:               mottFrenchName,
		language.German:               mottGermanName,
		language.Italian:              mottItalianName,
		language.Japanese:             mottJapaneseName,
		language.Korean:               mottKoreanName,
		language.LatinAmericanSpanish: mottLatinAmericanSpanishName,
		language.Russian:              mottRussianName,
		language.SimplifiedChinese:    mottSimplifiedChineseName,
		language.Spanish:              mottSpanishName,
		language.TraditionalChinese:   mottTraditionalChineseName}
)

var (
	mottCharacter = nook.Character{
		Animal:   Lion,
		Birthday: mottBirthday,
		Code:     mottCode,
		Gender:   nook.Male,
		Name:     mottName}
)

var (
	mottAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cagey"}

	mottCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grif-grif"}

	mottDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "rrrrrauw"}

	mottFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grif-grif"}

	mottGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "gröhl"}

	mottItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "groarrivo"}

	mottJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ハハハ"}

	mottKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "하하하"}

	mottLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "sapristi"}

	mottRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "зоосад"}

	mottSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哈哈哈"}

	mottSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "sapristi"}

	mottTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哈哈哈"}
)

var (
	mottPhrase = nook.Languages{
		language.AmericanEnglish:      mottAmericanEnglishName,
		language.CanadianFrench:       mottCanadianFrenchName,
		language.Dutch:                mottDutchName,
		language.French:               mottFrenchName,
		language.German:               mottGermanName,
		language.Italian:              mottItalianName,
		language.Japanese:             mottJapaneseName,
		language.Korean:               mottKoreanName,
		language.LatinAmericanSpanish: mottLatinAmericanSpanishName,
		language.Russian:              mottRussianName,
		language.SimplifiedChinese:    mottSimplifiedChineseName,
		language.Spanish:              mottSpanishName,
		language.TraditionalChinese:   mottTraditionalChineseName}
)

var (
	Mott = nook.Villager{
		Character:   mottCharacter,
		Personality: nook.Jock,
		Phrase:      mottPhrase}
)
