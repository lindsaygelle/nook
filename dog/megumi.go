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
		Value:    "メグミきゅん"}

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
		Value:    ""}

	megumiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	megumiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	megumiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	megumiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	megumiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	megumiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "きゅん"}

	megumiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	megumiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	megumiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	megumiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	megumiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	megumiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
