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
		Value:    "Fabien"}

	raddleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Raddle"}

	raddleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Fabien"}

	raddleGermanName = nook.Name{
		Language: language.German,
		Value:    "Sani"}

	raddleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Massimo"}

	raddleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カックン"}

	raddleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "개군"}

	raddleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Radiolo"}

	raddleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рэддл"}

	raddleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘉俊"}

	raddleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Radiolo"}

	raddleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘉俊"}
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
		Animal:   animal.Frog,
		Birthday: raddleBirthday,
		Code:     raddleCode,
		Key:      character.Raddle,
		Gender:   gender.Male,
		Name:     raddleName,
		Special:  false}
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
		language.AmericanEnglish:      raddleAmericanEnglishPhrase,
		language.CanadianFrench:       raddleCanadianFrenchPhrase,
		language.Dutch:                raddleDutchPhrase,
		language.French:               raddleFrenchPhrase,
		language.German:               raddleGermanPhrase,
		language.Italian:              raddleItalianPhrase,
		language.Japanese:             raddleJapanesePhrase,
		language.Korean:               raddleKoreanPhrase,
		language.LatinAmericanSpanish: raddleLatinAmericanSpanishPhrase,
		language.Russian:              raddleRussianPhrase,
		language.SimplifiedChinese:    raddleSimplifiedChinesePhrase,
		language.Spanish:              raddleSpanishPhrase,
		language.TraditionalChinese:   raddleTraditionalChinesePhrase}
)

var (
	Raddle = nook.Villager{
		Character:   raddleCharacter,
		Personality: personality.Lazy,
		Phrase:      raddlePhrase}
)
