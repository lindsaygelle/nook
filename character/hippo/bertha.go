package hippo

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
	// berthaBirthday represents bertha birthday.
	berthaBirthday = nook.Birthday{
		Day:   25,
		Month: time.April}
)

var (
	// berthaCode represents bertha code.
	berthaCode = nook.Code{
		Value: "hip03"}
)

var (
	// berthaAmericanEnglishName represents bertha american english name.
	berthaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bertha"}

	// berthaCanadianFrenchName represents bertha canadian french name.
	berthaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bertha"}

	// berthaDutchName represents bertha dutch name.
	berthaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bertha"}

	// berthaFrenchName represents bertha french name.
	berthaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bertha"}

	// berthaGermanName represents bertha german name.
	berthaGermanName = nook.Name{
		Language: language.German,
		Value:    "Berta"}

	// berthaItalianName represents bertha italian name.
	berthaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Umberta"}

	// berthaJapaneseName represents bertha japanese name.
	berthaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "あんこ"}

	// berthaKoreanName represents bertha korean name.
	berthaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "베티"}

	// berthaLatinAmericanSpanishName represents bertha latin american spanish name.
	berthaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Berta"}

	// berthaRussianName represents bertha russian name.
	berthaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Берта"}

	// berthaSimplifiedChineseName represents bertha simplified chinese name.
	berthaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "豆沙"}

	// berthaSpanishName represents bertha spanish name.
	berthaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Berta"}

	// berthaTraditionalChineseName represents bertha traditional chinese name.
	berthaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "豆沙"}
)

var (
	// berthaName represents bertha name.
	berthaName = nook.Languages{
		language.AmericanEnglish:      berthaAmericanEnglishName,
		language.CanadianFrench:       berthaCanadianFrenchName,
		language.Dutch:                berthaDutchName,
		language.French:               berthaFrenchName,
		language.German:               berthaGermanName,
		language.Italian:              berthaItalianName,
		language.Japanese:             berthaJapaneseName,
		language.Korean:               berthaKoreanName,
		language.LatinAmericanSpanish: berthaLatinAmericanSpanishName,
		language.Russian:              berthaRussianName,
		language.SimplifiedChinese:    berthaSimplifiedChineseName,
		language.Spanish:              berthaSpanishName,
		language.TraditionalChinese:   berthaTraditionalChineseName}
)

var (
	// berthaCharacter represents bertha character.
	berthaCharacter = nook.Character{
		Animal:   animal.Hippo,
		Birthday: berthaBirthday,
		Code:     berthaCode,
		Key:      character.Bertha,
		Gender:   gender.Female,
		Name:     berthaName,
		Special:  false}
)

var (
	// berthaAmericanEnglishPhrase represents bertha american english phrase.
	berthaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bloop"}

	// berthaCanadianFrenchPhrase represents bertha canadian french phrase.
	berthaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "gloups"}

	// berthaDutchPhrase represents bertha dutch phrase.
	berthaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "blubber"}

	// berthaFrenchPhrase represents bertha french phrase.
	berthaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "rololol"}

	// berthaGermanPhrase represents bertha german phrase.
	berthaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "blubber"}

	// berthaItalianPhrase represents bertha italian phrase.
	berthaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "blup"}

	// berthaJapanesePhrase represents bertha japanese phrase.
	berthaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "そうです"}

	// berthaKoreanPhrase represents bertha korean phrase.
	berthaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "맞습니다"}

	// berthaLatinAmericanSpanishPhrase represents bertha latin american spanish phrase.
	berthaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "glub-glub"}

	// berthaRussianPhrase represents bertha russian phrase.
	berthaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бултых"}

	// berthaSimplifiedChinesePhrase represents bertha simplified chinese phrase.
	berthaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "没错"}

	// berthaSpanishPhrase represents bertha spanish phrase.
	berthaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "glub-glub"}

	// berthaTraditionalChinesePhrase represents bertha traditional chinese phrase.
	berthaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "沒錯"}
)

var (
	// berthaPhrase represents bertha phrase.
	berthaPhrase = nook.Languages{
		language.AmericanEnglish:      berthaAmericanEnglishPhrase,
		language.CanadianFrench:       berthaCanadianFrenchPhrase,
		language.Dutch:                berthaDutchPhrase,
		language.French:               berthaFrenchPhrase,
		language.German:               berthaGermanPhrase,
		language.Italian:              berthaItalianPhrase,
		language.Japanese:             berthaJapanesePhrase,
		language.Korean:               berthaKoreanPhrase,
		language.LatinAmericanSpanish: berthaLatinAmericanSpanishPhrase,
		language.Russian:              berthaRussianPhrase,
		language.SimplifiedChinese:    berthaSimplifiedChinesePhrase,
		language.Spanish:              berthaSpanishPhrase,
		language.TraditionalChinese:   berthaTraditionalChinesePhrase}
)

var (
	// Bertha represents bertha.
	Bertha = nook.Villager{
		Character:   berthaCharacter,
		Personality: personality.Normal,
		Phrase:      berthaPhrase}
)
