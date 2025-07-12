package mouse

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
	// gretaBirthday represents greta birthday.
	gretaBirthday = nook.Birthday{
		Day:   5,
		Month: time.September}
)

var (
	// gretaCode represents greta code.
	gretaCode = nook.Code{
		Value: "mus16"}
)

var (
	// gretaAmericanEnglishName represents greta american english name.
	gretaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Greta"}

	// gretaCanadianFrenchName represents greta canadian french name.
	gretaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Greta"}

	// gretaDutchName represents greta dutch name.
	gretaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Greta"}

	// gretaFrenchName represents greta french name.
	gretaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Greta"}

	// gretaGermanName represents greta german name.
	gretaGermanName = nook.Name{
		Language: language.German,
		Value:    "Gretel"}

	// gretaItalianName represents greta italian name.
	gretaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Greta"}

	// gretaJapaneseName represents greta japanese name.
	gretaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ふくこ"}

	// gretaKoreanName represents greta korean name.
	gretaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "복자"}

	// gretaLatinAmericanSpanishName represents greta latin american spanish name.
	gretaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jimena"}

	// gretaRussianName represents greta russian name.
	gretaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Грета"}

	// gretaSimplifiedChineseName represents greta simplified chinese name.
	gretaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "如意"}

	// gretaSpanishName represents greta spanish name.
	gretaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jimena"}

	// gretaTraditionalChineseName represents greta traditional chinese name.
	gretaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "如意"}
)

var (
	// gretaName represents greta name.
	gretaName = nook.Languages{
		language.AmericanEnglish:      gretaAmericanEnglishName,
		language.CanadianFrench:       gretaCanadianFrenchName,
		language.Dutch:                gretaDutchName,
		language.French:               gretaFrenchName,
		language.German:               gretaGermanName,
		language.Italian:              gretaItalianName,
		language.Japanese:             gretaJapaneseName,
		language.Korean:               gretaKoreanName,
		language.LatinAmericanSpanish: gretaLatinAmericanSpanishName,
		language.Russian:              gretaRussianName,
		language.SimplifiedChinese:    gretaSimplifiedChineseName,
		language.Spanish:              gretaSpanishName,
		language.TraditionalChinese:   gretaTraditionalChineseName}
)

var (
	// gretaCharacter represents greta character.
	gretaCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: gretaBirthday,
		Code:     gretaCode,
		Key:      character.Greta,
		Gender:   gender.Female,
		Name:     gretaName,
		Special:  false}
)

var (
	// gretaAmericanEnglishPhrase represents greta american english phrase.
	gretaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "yelp"}

	// gretaCanadianFrenchPhrase represents greta canadian french phrase.
	gretaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "okka"}

	// gretaDutchPhrase represents greta dutch phrase.
	gretaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "piepzak"}

	// gretaFrenchPhrase represents greta french phrase.
	gretaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "téma"}

	// gretaGermanPhrase represents greta german phrase.
	gretaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "nagnag"}

	// gretaItalianPhrase represents greta italian phrase.
	gretaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "rattaplan"}

	// gretaJapanesePhrase represents greta japanese phrase.
	gretaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "おほほ"}

	// gretaKoreanPhrase represents greta korean phrase.
	gretaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "오호호"}

	// gretaLatinAmericanSpanishPhrase represents greta latin american spanish phrase.
	gretaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñiñi-ñiá"}

	// gretaRussianPhrase represents greta russian phrase.
	gretaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "писк"}

	// gretaSimplifiedChinesePhrase represents greta simplified chinese phrase.
	gretaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呵呵呵"}

	// gretaSpanishPhrase represents greta spanish phrase.
	gretaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bocadito"}

	// gretaTraditionalChinesePhrase represents greta traditional chinese phrase.
	gretaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "呵呵呵"}
)

var (
	// gretaPhrase represents greta phrase.
	gretaPhrase = nook.Languages{
		language.AmericanEnglish:      gretaAmericanEnglishPhrase,
		language.CanadianFrench:       gretaCanadianFrenchPhrase,
		language.Dutch:                gretaDutchPhrase,
		language.French:               gretaFrenchPhrase,
		language.German:               gretaGermanPhrase,
		language.Italian:              gretaItalianPhrase,
		language.Japanese:             gretaJapanesePhrase,
		language.Korean:               gretaKoreanPhrase,
		language.LatinAmericanSpanish: gretaLatinAmericanSpanishPhrase,
		language.Russian:              gretaRussianPhrase,
		language.SimplifiedChinese:    gretaSimplifiedChinesePhrase,
		language.Spanish:              gretaSpanishPhrase,
		language.TraditionalChinese:   gretaTraditionalChinesePhrase}
)

var (
	// Greta represents greta.
	Greta = nook.Villager{
		Character:   gretaCharacter,
		Personality: personality.Snooty,
		Phrase:      gretaPhrase}
)
