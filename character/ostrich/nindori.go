package ostrich

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	nindoriBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	nindoriCode = nook.Code{
		Value: ""}
)

var (
	nindoriAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Nindori"}

	nindoriCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	nindoriDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	nindoriFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	nindoriGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	nindoriItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	nindoriJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ニンドリ"}

	nindoriKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	nindoriLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	nindoriRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	nindoriSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	nindoriSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	nindoriTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	nindoriName = nook.Languages{
		language.AmericanEnglish:      nindoriAmericanEnglishName,
		language.CanadianFrench:       nindoriCanadianFrenchName,
		language.Dutch:                nindoriDutchName,
		language.French:               nindoriFrenchName,
		language.German:               nindoriGermanName,
		language.Italian:              nindoriItalianName,
		language.Japanese:             nindoriJapaneseName,
		language.Korean:               nindoriKoreanName,
		language.LatinAmericanSpanish: nindoriLatinAmericanSpanishName,
		language.Russian:              nindoriRussianName,
		language.SimplifiedChinese:    nindoriSimplifiedChineseName,
		language.Spanish:              nindoriSpanishName,
		language.TraditionalChinese:   nindoriTraditionalChineseName}
)

var (
	nindoriCharacter = nook.Character{
		Animal:   animal.Ostrich,
		Birthday: nindoriBirthday,
		Code:     nindoriCode,
		Gender:   gender.Male,
		Name:     nindoriName}
)

var (
	nindoriAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ドクソウ"}

	nindoriCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	nindoriDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	nindoriFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	nindoriGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	nindoriItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	nindoriJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ドクソウ"}

	nindoriKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	nindoriLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	nindoriRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	nindoriSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	nindoriSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	nindoriTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	nindoriPhrase = nook.Languages{
		language.AmericanEnglish:      nindoriAmericanEnglishName,
		language.CanadianFrench:       nindoriCanadianFrenchName,
		language.Dutch:                nindoriDutchName,
		language.French:               nindoriFrenchName,
		language.German:               nindoriGermanName,
		language.Italian:              nindoriItalianName,
		language.Japanese:             nindoriJapaneseName,
		language.Korean:               nindoriKoreanName,
		language.LatinAmericanSpanish: nindoriLatinAmericanSpanishName,
		language.Russian:              nindoriRussianName,
		language.SimplifiedChinese:    nindoriSimplifiedChineseName,
		language.Spanish:              nindoriSpanishName,
		language.TraditionalChinese:   nindoriTraditionalChineseName}
)

var (
	Nindori = nook.Villager{
		Character:   nindoriCharacter,
		Personality: personality.Jock,
		Phrase:      nindoriPhrase}
)
