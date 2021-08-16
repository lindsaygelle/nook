package frog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
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
		Value:    "Gloria"}

	gigiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gigi"}

	gigiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gloria"}

	gigiGermanName = nook.Name{
		Language: language.German,
		Value:    "Violetta"}

	gigiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Raganà"}

	gigiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リンダ"}

	gigiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "린다"}

	gigiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cleo"}

	gigiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джиджи"}

	gigiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "林妲"}

	gigiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cleo"}

	gigiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "林妲"}
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
		Animal:   animal.Frog,
		Birthday: gigiBirthday,
		Code:     gigiCode,
		Gender:   gender.Female,
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
		Personality: personality.Snooty,
		Phrase:      gigiPhrase}
)
