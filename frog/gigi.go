package frog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	gigiBirthday = nook.Birthday{
		Day:   11,
		Month: time.August}
)

var (
	gigiCode = nook.Code{
		Value: "flg16"}
)

var (
	gigiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gigi"}

	gigiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gloriacrôazou"}

	gigiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gigikwaakette"}

	gigiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gloriacrôazou"}

	gigiGermanName = nook.Name{
		Language: language.German,
		Value:    "Violettaquaki"}

	gigiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Raganàgurghigù"}

	gigiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リンダベイビィ"}

	gigiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "린다베이비"}

	gigiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cleocroá-croá"}

	gigiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джиджиквакет"}

	gigiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "林妲宝贝"}

	gigiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cleotrebolito"}

	gigiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "林妲寶貝"}
)

var (
	gigiName = nook.Languages{
		language.AmericanEnglish:      gigiAmericanEnglishName,
		language.CanadianFrench:       gigiCanadianFrenchName,
		language.Dutch:                gigiDutchName,
		language.French:               gigiFrenchName,
		language.German:               gigiGermanName,
		language.Italian:              gigiItalianName,
		language.Japanese:             gigiJapaneseName,
		language.Korean:               gigiKoreanName,
		language.LatinAmericanSpanish: gigiLatinAmericanSpanishName,
		language.Russian:              gigiRussianName,
		language.SimplifiedChinese:    gigiSimplifiedChineseName,
		language.Spanish:              gigiSpanishName,
		language.TraditionalChinese:   gigiTraditionalChineseName}
)

var (
	gigiCharacter = nook.Character{
		Animal:   Frog,
		Birthday: gigiBirthday,
		Code:     gigiCode,
		Gender:   nook.Female,
		Name:     gigiName}
)

var (
	gigiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ribbette"}

	gigiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "crôazou"}

	gigiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kwaakette"}

	gigiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "crôazou"}

	gigiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quaki"}

	gigiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gurghigù"}

	gigiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ベイビィ"}

	gigiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "베이비"}

	gigiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "croá-croá"}

	gigiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "квакет"}

	gigiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "宝贝"}

	gigiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "trebolito"}

	gigiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "寶貝"}
)

var (
	gigiPhrase = nook.Languages{
		language.AmericanEnglish:      gigiAmericanEnglishName,
		language.CanadianFrench:       gigiCanadianFrenchName,
		language.Dutch:                gigiDutchName,
		language.French:               gigiFrenchName,
		language.German:               gigiGermanName,
		language.Italian:              gigiItalianName,
		language.Japanese:             gigiJapaneseName,
		language.Korean:               gigiKoreanName,
		language.LatinAmericanSpanish: gigiLatinAmericanSpanishName,
		language.Russian:              gigiRussianName,
		language.SimplifiedChinese:    gigiSimplifiedChineseName,
		language.Spanish:              gigiSpanishName,
		language.TraditionalChinese:   gigiTraditionalChineseName}
)

var (
	Gigi = nook.Villager{
		Character:   gigiCharacter,
		Personality: nook.Snooty,
		Phrase:      gigiPhrase}
)
