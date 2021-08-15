package frog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	frobertBirthday = nook.Birthday{
		Day:   8,
		Month: time.February}
)

var (
	frobertCode = nook.Code{
		Value: "flg02"}
)

var (
	frobertAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Frobert"}

	frobertCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Verbert"}

	frobertDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Frobert"}

	frobertFrenchName = nook.Name{
		Language: language.French,
		Value:    "Verbert"}

	frobertGermanName = nook.Name{
		Language: language.German,
		Value:    "Fritz"}

	frobertItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Froberto"}

	frobertJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "コージィ"}

	frobertKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "구리구리"}

	frobertLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Croaldo"}

	frobertRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фроберт"}

	frobertSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "江适"}

	frobertSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Croaldo"}

	frobertTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "江適"}
)

var (
	frobertName = nook.Languages{
		language.AmericanEnglish:      frobertAmericanEnglishName,
		language.CanadianFrench:       frobertCanadianFrenchName,
		language.Dutch:                frobertDutchName,
		language.French:               frobertFrenchName,
		language.German:               frobertGermanName,
		language.Italian:              frobertItalianName,
		language.Japanese:             frobertJapaneseName,
		language.Korean:               frobertKoreanName,
		language.LatinAmericanSpanish: frobertLatinAmericanSpanishName,
		language.Russian:              frobertRussianName,
		language.SimplifiedChinese:    frobertSimplifiedChineseName,
		language.Spanish:              frobertSpanishName,
		language.TraditionalChinese:   frobertTraditionalChineseName}
)

var (
	frobertCharacter = nook.Character{
		Animal:   Frog,
		Birthday: frobertBirthday,
		Code:     frobertCode,
		Gender:   nook.Male,
		Name:     frobertName}
)

var (
	frobertAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "fribbit"}

	frobertCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "breup"}

	frobertDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "fffwaak"}

	frobertFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "breup"}

	frobertGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quäääk"}

	frobertItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "fribbit"}

	frobertJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "クルリ"}

	frobertKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "객울"}

	frobertLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fribit"}

	frobertRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "квак"}

	frobertSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咕噜"}

	frobertSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "anquitas"}

	frobertTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咕嚕"}
)

var (
	frobertPhrase = nook.Languages{
		language.AmericanEnglish:      frobertAmericanEnglishName,
		language.CanadianFrench:       frobertCanadianFrenchName,
		language.Dutch:                frobertDutchName,
		language.French:               frobertFrenchName,
		language.German:               frobertGermanName,
		language.Italian:              frobertItalianName,
		language.Japanese:             frobertJapaneseName,
		language.Korean:               frobertKoreanName,
		language.LatinAmericanSpanish: frobertLatinAmericanSpanishName,
		language.Russian:              frobertRussianName,
		language.SimplifiedChinese:    frobertSimplifiedChineseName,
		language.Spanish:              frobertSpanishName,
		language.TraditionalChinese:   frobertTraditionalChineseName}
)

var (
	Frobert = nook.Villager{
		Character:   frobertCharacter,
		Personality: nook.Jock,
		Phrase:      frobertPhrase}
)
