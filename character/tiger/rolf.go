package tiger

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
		Value:    "Ralf"}

	rolfDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rolf"}

	rolfFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ralf"}

	rolfGermanName = nook.Name{
		Language: language.German,
		Value:    "Boris"}

	rolfItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rolf"}

	rolfJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チョモラン"}

	rolfKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "호랭이"}

	rolfLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Albino"}

	rolfRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рольф"}

	rolfSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "朱穆朗"}

	rolfSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Albino"}

	rolfTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "朱穆朗"}
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
		Animal:   animal.Tiger,
		Birthday: rolfBirthday,
		Code:     rolfCode,
		Key:      character.Rolf,
		Gender:   gender.Male,
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
		Personality: personality.Cranky,
		Phrase:      rolfPhrase}
)
