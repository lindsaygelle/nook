package ostrich

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "N/A"}

	nindoriDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	nindoriFrenchName = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	nindoriGermanName = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	nindoriItalianName = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	nindoriJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ニンドリ"}

	nindoriKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	nindoriLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	nindoriRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	nindoriSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	nindoriSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	nindoriTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Animal:   Ostrich,
		Birthday: nindoriBirthday,
		Code:     nindoriCode,
		Gender:   nook.Male,
		Name:     nindoriName}
)

var (
	nindoriAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	nindoriCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	nindoriDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	nindoriFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	nindoriGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	nindoriItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	nindoriJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ドクソウ"}

	nindoriKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	nindoriLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	nindoriRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	nindoriSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	nindoriSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	nindoriTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Personality: nook.Jock,
		Phrase:      nindoriPhrase}
)
