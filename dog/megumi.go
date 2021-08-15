package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	megumiBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	megumiCode = nook.Code{
		Value: ""}
)

var (
	megumiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Megumi"}

	megumiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	megumiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	megumiFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	megumiGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	megumiItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	megumiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "メグミ"}

	megumiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	megumiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	megumiRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	megumiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	megumiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	megumiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	megumiName = nook.Languages{
		language.AmericanEnglish:      megumiAmericanEnglishName,
		language.CanadianFrench:       megumiCanadianFrenchName,
		language.Dutch:                megumiDutchName,
		language.French:               megumiFrenchName,
		language.German:               megumiGermanName,
		language.Italian:              megumiItalianName,
		language.Japanese:             megumiJapaneseName,
		language.Korean:               megumiKoreanName,
		language.LatinAmericanSpanish: megumiLatinAmericanSpanishName,
		language.Russian:              megumiRussianName,
		language.SimplifiedChinese:    megumiSimplifiedChineseName,
		language.Spanish:              megumiSpanishName,
		language.TraditionalChinese:   megumiTraditionalChineseName}
)

var (
	megumiCharacter = nook.Character{
		Animal:   Dog,
		Birthday: megumiBirthday,
		Code:     megumiCode,
		Gender:   nook.Female,
		Name:     megumiName}
)

var (
	megumiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "きゅん"}

	megumiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	megumiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	megumiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	megumiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	megumiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	megumiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "きゅん"}

	megumiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	megumiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	megumiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	megumiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	megumiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	megumiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	megumiPhrase = nook.Languages{
		language.AmericanEnglish:      megumiAmericanEnglishName,
		language.CanadianFrench:       megumiCanadianFrenchName,
		language.Dutch:                megumiDutchName,
		language.French:               megumiFrenchName,
		language.German:               megumiGermanName,
		language.Italian:              megumiItalianName,
		language.Japanese:             megumiJapaneseName,
		language.Korean:               megumiKoreanName,
		language.LatinAmericanSpanish: megumiLatinAmericanSpanishName,
		language.Russian:              megumiRussianName,
		language.SimplifiedChinese:    megumiSimplifiedChineseName,
		language.Spanish:              megumiSpanishName,
		language.TraditionalChinese:   megumiTraditionalChineseName}
)

var (
	Megumi = nook.Villager{
		Character:   megumiCharacter,
		Personality: nook.Peppy,
		Phrase:      megumiPhrase}
)
