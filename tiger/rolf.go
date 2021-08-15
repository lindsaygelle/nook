package tiger

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	rolfBirthday = nook.Birthday{
		Day:   22,
		Month: time.August}
)

var (
	rolfCode = nook.Code{
		Value: "tig00"}
)

var (
	rolfAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rolf"}

	rolfCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ralfgrrrufff"}

	rolfDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rolfgrrrolf"}

	rolfFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ralfgrrrufff"}

	rolfGermanName = nook.Name{
		Language: language.German,
		Value:    "Borisbrüll"}

	rolfItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rolfgrrrolf"}

	rolfJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チョモランだガー"}

	rolfKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "호랭이근데"}

	rolfLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Albinogrooorf"}

	rolfRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рольфгр-р-рольф"}

	rolfSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "朱穆朗王"}

	rolfSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Albinogrooorf"}

	rolfTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "朱穆朗王"}
)

var (
	rolfName = nook.Languages{
		language.AmericanEnglish:      rolfAmericanEnglishName,
		language.CanadianFrench:       rolfCanadianFrenchName,
		language.Dutch:                rolfDutchName,
		language.French:               rolfFrenchName,
		language.German:               rolfGermanName,
		language.Italian:              rolfItalianName,
		language.Japanese:             rolfJapaneseName,
		language.Korean:               rolfKoreanName,
		language.LatinAmericanSpanish: rolfLatinAmericanSpanishName,
		language.Russian:              rolfRussianName,
		language.SimplifiedChinese:    rolfSimplifiedChineseName,
		language.Spanish:              rolfSpanishName,
		language.TraditionalChinese:   rolfTraditionalChineseName}
)

var (
	rolfCharacter = nook.Character{
		Animal:   Tiger,
		Birthday: rolfBirthday,
		Code:     rolfCode,
		Gender:   nook.Male,
		Name:     rolfName}
)

var (
	rolfAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "grrrolf"}

	rolfCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grrrufff"}

	rolfDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "grrrolf"}

	rolfFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grrrufff"}

	rolfGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brüll"}

	rolfItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grrrolf"}

	rolfJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だガー"}

	rolfKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "근데"}

	rolfLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grooorf"}

	rolfRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "гр-р-рольф"}

	rolfSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "王"}

	rolfSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "grooorf"}

	rolfTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "王"}
)

var (
	rolfPhrase = nook.Languages{
		language.AmericanEnglish:      rolfAmericanEnglishName,
		language.CanadianFrench:       rolfCanadianFrenchName,
		language.Dutch:                rolfDutchName,
		language.French:               rolfFrenchName,
		language.German:               rolfGermanName,
		language.Italian:              rolfItalianName,
		language.Japanese:             rolfJapaneseName,
		language.Korean:               rolfKoreanName,
		language.LatinAmericanSpanish: rolfLatinAmericanSpanishName,
		language.Russian:              rolfRussianName,
		language.SimplifiedChinese:    rolfSimplifiedChineseName,
		language.Spanish:              rolfSpanishName,
		language.TraditionalChinese:   rolfTraditionalChineseName}
)

var (
	Rolf = nook.Villager{
		Character:   rolfCharacter,
		Personality: nook.Cranky,
		Phrase:      rolfPhrase}
)
