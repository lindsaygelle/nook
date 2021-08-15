package frog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	raddleBirthday = nook.Birthday{
		Day:   6,
		Month: time.June}
)

var (
	raddleCode = nook.Code{
		Value: "flg15"}
)

var (
	raddleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Raddle"}

	raddleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Fabienkroah"}

	raddleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Raddlekraah"}

	raddleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Fabienkroah"}

	raddleGermanName = nook.Name{
		Language: language.German,
		Value:    "Sanitatü"}

	raddleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Massimorrruuuak"}

	raddleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カックンへくしっ"}

	raddleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "개군콜록"}

	raddleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Radiolocroaqui"}

	raddleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рэддлах-ох"}

	raddleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘉俊哈啾"}

	raddleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Radiolocroaqui"}

	raddleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘉俊哈啾"}
)

var (
	raddleName = nook.Languages{
		language.AmericanEnglish:      raddleAmericanEnglishName,
		language.CanadianFrench:       raddleCanadianFrenchName,
		language.Dutch:                raddleDutchName,
		language.French:               raddleFrenchName,
		language.German:               raddleGermanName,
		language.Italian:              raddleItalianName,
		language.Japanese:             raddleJapaneseName,
		language.Korean:               raddleKoreanName,
		language.LatinAmericanSpanish: raddleLatinAmericanSpanishName,
		language.Russian:              raddleRussianName,
		language.SimplifiedChinese:    raddleSimplifiedChineseName,
		language.Spanish:              raddleSpanishName,
		language.TraditionalChinese:   raddleTraditionalChineseName}
)

var (
	raddleCharacter = nook.Character{
		Animal:   Frog,
		Birthday: raddleBirthday,
		Code:     raddleCode,
		Gender:   nook.Male,
		Name:     raddleName}
)

var (
	raddleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "aaach—"}

	raddleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "kroah"}

	raddleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kraah"}

	raddleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "kroah"}

	raddleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tatü"}

	raddleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "rrruuuak"}

	raddleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "へくしっ"}

	raddleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "콜록"}

	raddleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "croaqui"}

	raddleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ах-ох"}

	raddleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哈啾"}

	raddleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "croaqui"}

	raddleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哈啾"}
)

var (
	raddlePhrase = nook.Languages{
		language.AmericanEnglish:      raddleAmericanEnglishName,
		language.CanadianFrench:       raddleCanadianFrenchName,
		language.Dutch:                raddleDutchName,
		language.French:               raddleFrenchName,
		language.German:               raddleGermanName,
		language.Italian:              raddleItalianName,
		language.Japanese:             raddleJapaneseName,
		language.Korean:               raddleKoreanName,
		language.LatinAmericanSpanish: raddleLatinAmericanSpanishName,
		language.Russian:              raddleRussianName,
		language.SimplifiedChinese:    raddleSimplifiedChineseName,
		language.Spanish:              raddleSpanishName,
		language.TraditionalChinese:   raddleTraditionalChineseName}
)

var (
	Raddle = nook.Villager{
		Character:   raddleCharacter,
		Personality: nook.Lazy,
		Phrase:      raddlePhrase}
)
