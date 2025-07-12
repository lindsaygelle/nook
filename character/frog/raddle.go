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
	// raddleBirthday represents raddle birthday.
	raddleBirthday = nook.Birthday{
		Day:   6,
		Month: time.June}
)

var (
	// raddleCode represents raddle code.
	raddleCode = nook.Code{
		Value: "flg15"}
)

var (
	// raddleAmericanEnglishName represents raddle american english name.
	raddleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Raddle"}

	// raddleCanadianFrenchName represents raddle canadian french name.
	raddleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Fabien"}

	// raddleDutchName represents raddle dutch name.
	raddleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Raddle"}

	// raddleFrenchName represents raddle french name.
	raddleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Fabien"}

	// raddleGermanName represents raddle german name.
	raddleGermanName = nook.Name{
		Language: language.German,
		Value:    "Sani"}

	// raddleItalianName represents raddle italian name.
	raddleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Massimo"}

	// raddleJapaneseName represents raddle japanese name.
	raddleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カックン"}

	// raddleKoreanName represents raddle korean name.
	raddleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "개군"}

	// raddleLatinAmericanSpanishName represents raddle latin american spanish name.
	raddleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Radiolo"}

	// raddleRussianName represents raddle russian name.
	raddleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рэддл"}

	// raddleSimplifiedChineseName represents raddle simplified chinese name.
	raddleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘉俊"}

	// raddleSpanishName represents raddle spanish name.
	raddleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Radiolo"}

	// raddleTraditionalChineseName represents raddle traditional chinese name.
	raddleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘉俊"}
)

var (
	// raddleName represents raddle name.
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
	// raddleCharacter represents raddle character.
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
	// raddleAmericanEnglishPhrase represents raddle american english phrase.
	raddleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "aaach—"}

	// raddleCanadianFrenchPhrase represents raddle canadian french phrase.
	raddleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "kroah"}

	// raddleDutchPhrase represents raddle dutch phrase.
	raddleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kraah"}

	// raddleFrenchPhrase represents raddle french phrase.
	raddleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "kroah"}

	// raddleGermanPhrase represents raddle german phrase.
	raddleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tatü"}

	// raddleItalianPhrase represents raddle italian phrase.
	raddleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "rrruuuak"}

	// raddleJapanesePhrase represents raddle japanese phrase.
	raddleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "へくしっ"}

	// raddleKoreanPhrase represents raddle korean phrase.
	raddleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "콜록"}

	// raddleLatinAmericanSpanishPhrase represents raddle latin american spanish phrase.
	raddleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "croaqui"}

	// raddleRussianPhrase represents raddle russian phrase.
	raddleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ах-ох"}

	// raddleSimplifiedChinesePhrase represents raddle simplified chinese phrase.
	raddleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哈啾"}

	// raddleSpanishPhrase represents raddle spanish phrase.
	raddleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "croaqui"}

	// raddleTraditionalChinesePhrase represents raddle traditional chinese phrase.
	raddleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哈啾"}
)

var (
	// raddlePhrase represents raddle phrase.
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
	// Raddle represents raddle.
	Raddle = nook.Villager{
		Character:   raddleCharacter,
		Personality: personality.Lazy,
		Phrase:      raddlePhrase}
)
