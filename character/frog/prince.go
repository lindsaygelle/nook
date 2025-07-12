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
	// princeBirthday represents prince birthday.
	princeBirthday = nook.Birthday{
		Day:   21,
		Month: time.July}
)

var (
	// princeCode represents prince code.
	princeCode = nook.Code{
		Value: "flg12"}
)

var (
	// princeAmericanEnglishName represents prince american english name.
	princeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Prince"}

	// princeCanadianFrenchName represents prince canadian french name.
	princeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Prince"}

	// princeDutchName represents prince dutch name.
	princeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Prince"}

	// princeFrenchName represents prince french name.
	princeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Prince"}

	// princeGermanName represents prince german name.
	princeGermanName = nook.Name{
		Language: language.German,
		Value:    "Prinz"}

	// princeItalianName represents prince italian name.
	princeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Principe"}

	// princeJapaneseName represents prince japanese name.
	princeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カール"}

	// princeKoreanName represents prince korean name.
	princeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "카일"}

	// princeLatinAmericanSpanishName represents prince latin american spanish name.
	princeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Felipe"}

	// princeRussianName represents prince russian name.
	princeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Принс"}

	// princeSimplifiedChineseName represents prince simplified chinese name.
	princeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "青呱"}

	// princeSpanishName represents prince spanish name.
	princeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Felipe"}

	// princeTraditionalChineseName represents prince traditional chinese name.
	princeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "青呱"}
)

var (
	// princeName represents prince name.
	princeName = nook.Languages{
		language.AmericanEnglish:      princeAmericanEnglishName,
		language.CanadianFrench:       princeCanadianFrenchName,
		language.Dutch:                princeDutchName,
		language.French:               princeFrenchName,
		language.German:               princeGermanName,
		language.Italian:              princeItalianName,
		language.Japanese:             princeJapaneseName,
		language.Korean:               princeKoreanName,
		language.LatinAmericanSpanish: princeLatinAmericanSpanishName,
		language.Russian:              princeRussianName,
		language.SimplifiedChinese:    princeSimplifiedChineseName,
		language.Spanish:              princeSpanishName,
		language.TraditionalChinese:   princeTraditionalChineseName}
)

var (
	// princeCharacter represents prince character.
	princeCharacter = nook.Character{
		Animal:   animal.Frog,
		Birthday: princeBirthday,
		Code:     princeCode,
		Key:      character.Prince,
		Gender:   gender.Male,
		Name:     princeName,
		Special:  false}
)

var (
	// princeAmericanEnglishPhrase represents prince american english phrase.
	princeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "burrup"}

	// princeCanadianFrenchPhrase represents prince canadian french phrase.
	princeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "boiiiiing"}

	// princeDutchPhrase represents prince dutch phrase.
	princeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "blurp"}

	// princeFrenchPhrase represents prince french phrase.
	princeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "boiiiiing"}

	// princeGermanPhrase represents prince german phrase.
	princeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "gurp"}

	// princeItalianPhrase represents prince italian phrase.
	princeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "berp"}

	// princeJapanesePhrase represents prince japanese phrase.
	princeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですだ"}

	// princeKoreanPhrase represents prince korean phrase.
	princeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "이리오슈"}

	// princeLatinAmericanSpanishPhrase represents prince latin american spanish phrase.
	princeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "buuurp"}

	// princeRussianPhrase represents prince russian phrase.
	princeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "квак-квак"}

	// princeSimplifiedChinesePhrase represents prince simplified chinese phrase.
	princeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是的"}

	// princeSpanishPhrase represents prince spanish phrase.
	princeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "buuurp"}

	// princeTraditionalChinesePhrase represents prince traditional chinese phrase.
	princeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是的"}
)

var (
	// princePhrase represents prince phrase.
	princePhrase = nook.Languages{
		language.AmericanEnglish:      princeAmericanEnglishPhrase,
		language.CanadianFrench:       princeCanadianFrenchPhrase,
		language.Dutch:                princeDutchPhrase,
		language.French:               princeFrenchPhrase,
		language.German:               princeGermanPhrase,
		language.Italian:              princeItalianPhrase,
		language.Japanese:             princeJapanesePhrase,
		language.Korean:               princeKoreanPhrase,
		language.LatinAmericanSpanish: princeLatinAmericanSpanishPhrase,
		language.Russian:              princeRussianPhrase,
		language.SimplifiedChinese:    princeSimplifiedChinesePhrase,
		language.Spanish:              princeSpanishPhrase,
		language.TraditionalChinese:   princeTraditionalChinesePhrase}
)

var (
	// Prince represents prince.
	Prince = nook.Villager{
		Character:   princeCharacter,
		Personality: personality.Lazy,
		Phrase:      princePhrase}
)
