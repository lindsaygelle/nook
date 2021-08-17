package frog

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
		Animal:   animal.Frog,
		Birthday: frobertBirthday,
		Code:     frobertCode,
		Key:      character.Frobert,
		Gender:   gender.Male,
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
		language.AmericanEnglish:      frobertAmericanEnglishPhrase,
		language.CanadianFrench:       frobertCanadianFrenchPhrase,
		language.Dutch:                frobertDutchPhrase,
		language.French:               frobertFrenchPhrase,
		language.German:               frobertGermanPhrase,
		language.Italian:              frobertItalianPhrase,
		language.Japanese:             frobertJapanesePhrase,
		language.Korean:               frobertKoreanPhrase,
		language.LatinAmericanSpanish: frobertLatinAmericanSpanishPhrase,
		language.Russian:              frobertRussianPhrase,
		language.SimplifiedChinese:    frobertSimplifiedChinesePhrase,
		language.Spanish:              frobertSpanishPhrase,
		language.TraditionalChinese:   frobertTraditionalChinesePhrase}
)

var (
	Frobert = nook.Villager{
		Character:   frobertCharacter,
		Personality: personality.Jock,
		Phrase:      frobertPhrase}
)
