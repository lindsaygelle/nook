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
	// frobertBirthday represents frobert birthday.
	frobertBirthday = nook.Birthday{
		Day:   8,
		Month: time.February}
)

var (
	// frobertCode represents frobert code.
	frobertCode = nook.Code{
		Value: "flg02"}
)

var (
	// frobertAmericanEnglishName represents frobert american english name.
	frobertAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Frobert"}

	// frobertCanadianFrenchName represents frobert canadian french name.
	frobertCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Verbert"}

	// frobertDutchName represents frobert dutch name.
	frobertDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Frobert"}

	// frobertFrenchName represents frobert french name.
	frobertFrenchName = nook.Name{
		Language: language.French,
		Value:    "Verbert"}

	// frobertGermanName represents frobert german name.
	frobertGermanName = nook.Name{
		Language: language.German,
		Value:    "Fritz"}

	// frobertItalianName represents frobert italian name.
	frobertItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Froberto"}

	// frobertJapaneseName represents frobert japanese name.
	frobertJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "コージィ"}

	// frobertKoreanName represents frobert korean name.
	frobertKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "구리구리"}

	// frobertLatinAmericanSpanishName represents frobert latin american spanish name.
	frobertLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Croaldo"}

	// frobertRussianName represents frobert russian name.
	frobertRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фроберт"}

	// frobertSimplifiedChineseName represents frobert simplified chinese name.
	frobertSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "江适"}

	// frobertSpanishName represents frobert spanish name.
	frobertSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Croaldo"}

	// frobertTraditionalChineseName represents frobert traditional chinese name.
	frobertTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "江適"}
)

var (
	// frobertName represents frobert name.
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
	// frobertCharacter represents frobert character.
	frobertCharacter = nook.Character{
		Animal:   animal.Frog,
		Birthday: frobertBirthday,
		Code:     frobertCode,
		Key:      character.Frobert,
		Gender:   gender.Male,
		Name:     frobertName,
		Special:  false}
)

var (
	// frobertAmericanEnglishPhrase represents frobert american english phrase.
	frobertAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "fribbit"}

	// frobertCanadianFrenchPhrase represents frobert canadian french phrase.
	frobertCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "breup"}

	// frobertDutchPhrase represents frobert dutch phrase.
	frobertDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "fffwaak"}

	// frobertFrenchPhrase represents frobert french phrase.
	frobertFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "breup"}

	// frobertGermanPhrase represents frobert german phrase.
	frobertGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quäääk"}

	// frobertItalianPhrase represents frobert italian phrase.
	frobertItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "fribbit"}

	// frobertJapanesePhrase represents frobert japanese phrase.
	frobertJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "クルリ"}

	// frobertKoreanPhrase represents frobert korean phrase.
	frobertKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "객울"}

	// frobertLatinAmericanSpanishPhrase represents frobert latin american spanish phrase.
	frobertLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fribit"}

	// frobertRussianPhrase represents frobert russian phrase.
	frobertRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "квак"}

	// frobertSimplifiedChinesePhrase represents frobert simplified chinese phrase.
	frobertSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咕噜"}

	// frobertSpanishPhrase represents frobert spanish phrase.
	frobertSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "anquitas"}

	// frobertTraditionalChinesePhrase represents frobert traditional chinese phrase.
	frobertTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咕嚕"}
)

var (
	// frobertPhrase represents frobert phrase.
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
	// Frobert represents frobert.
	Frobert = nook.Villager{
		Character:   frobertCharacter,
		Personality: personality.Jock,
		Phrase:      frobertPhrase}
)
