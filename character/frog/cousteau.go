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
	// cousteauBirthday represents cousteau birthday.
	cousteauBirthday = nook.Birthday{
		Day:   17,
		Month: time.December}
)

var (
	// cousteauCode represents cousteau code.
	cousteauCode = nook.Code{
		Value: "flg10"}
)

var (
	// cousteauAmericanEnglishName represents cousteau american english name.
	cousteauAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cousteau"}

	// cousteauCanadianFrenchName represents cousteau canadian french name.
	cousteauCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Figaro"}

	// cousteauDutchName represents cousteau dutch name.
	cousteauDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cousteau"}

	// cousteauFrenchName represents cousteau french name.
	cousteauFrenchName = nook.Name{
		Language: language.French,
		Value:    "Figaro"}

	// cousteauGermanName represents cousteau german name.
	cousteauGermanName = nook.Name{
		Language: language.German,
		Value:    "Jacques"}

	// cousteauItalianName represents cousteau italian name.
	cousteauItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Jacques"}

	// cousteauJapaneseName represents cousteau japanese name.
	cousteauJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハルマキ"}

	// cousteauKoreanName represents cousteau korean name.
	cousteauKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "왕서방"}

	// cousteauLatinAmericanSpanishName represents cousteau latin american spanish name.
	cousteauLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cousteau"}

	// cousteauRussianName represents cousteau russian name.
	cousteauRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кусто"}

	// cousteauSimplifiedChineseName represents cousteau simplified chinese name.
	cousteauSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "春卷"}

	// cousteauSpanishName represents cousteau spanish name.
	cousteauSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cousteau"}

	// cousteauTraditionalChineseName represents cousteau traditional chinese name.
	cousteauTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "春卷"}
)

var (
	// cousteauName represents cousteau name.
	cousteauName = nook.Languages{
		language.AmericanEnglish:      cousteauAmericanEnglishName,
		language.CanadianFrench:       cousteauCanadianFrenchName,
		language.Dutch:                cousteauDutchName,
		language.French:               cousteauFrenchName,
		language.German:               cousteauGermanName,
		language.Italian:              cousteauItalianName,
		language.Japanese:             cousteauJapaneseName,
		language.Korean:               cousteauKoreanName,
		language.LatinAmericanSpanish: cousteauLatinAmericanSpanishName,
		language.Russian:              cousteauRussianName,
		language.SimplifiedChinese:    cousteauSimplifiedChineseName,
		language.Spanish:              cousteauSpanishName,
		language.TraditionalChinese:   cousteauTraditionalChineseName}
)

var (
	// cousteauCharacter represents cousteau character.
	cousteauCharacter = nook.Character{
		Animal:   animal.Frog,
		Birthday: cousteauBirthday,
		Code:     cousteauCode,
		Key:      character.Cousteau,
		Gender:   gender.Male,
		Name:     cousteauName,
		Special:  false}
)

var (
	// cousteauAmericanEnglishPhrase represents cousteau american english phrase.
	cousteauAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "oui oui"}

	// cousteauCanadianFrenchPhrase represents cousteau canadian french phrase.
	cousteauCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bidule"}

	// cousteauDutchPhrase represents cousteau dutch phrase.
	cousteauDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "oui oui"}

	// cousteauFrenchPhrase represents cousteau french phrase.
	cousteauFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bidule"}

	// cousteauGermanPhrase represents cousteau german phrase.
	cousteauGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "oui, oui"}

	// cousteauItalianPhrase represents cousteau italian phrase.
	cousteauItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "et voilà"}

	// cousteauJapanesePhrase represents cousteau japanese phrase.
	cousteauJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "のことよ"}

	// cousteauKoreanPhrase represents cousteau korean phrase.
	cousteauKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "라이라이"}

	// cousteauLatinAmericanSpanishPhrase represents cousteau latin american spanish phrase.
	cousteauLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "oui oui"}

	// cousteauRussianPhrase represents cousteau russian phrase.
	cousteauRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "уи-уи"}

	// cousteauSimplifiedChinesePhrase represents cousteau simplified chinese phrase.
	cousteauSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "不客气"}

	// cousteauSpanishPhrase represents cousteau spanish phrase.
	cousteauSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "oui oui"}

	// cousteauTraditionalChinesePhrase represents cousteau traditional chinese phrase.
	cousteauTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "不客氣"}
)

var (
	// cousteauPhrase represents cousteau phrase.
	cousteauPhrase = nook.Languages{
		language.AmericanEnglish:      cousteauAmericanEnglishPhrase,
		language.CanadianFrench:       cousteauCanadianFrenchPhrase,
		language.Dutch:                cousteauDutchPhrase,
		language.French:               cousteauFrenchPhrase,
		language.German:               cousteauGermanPhrase,
		language.Italian:              cousteauItalianPhrase,
		language.Japanese:             cousteauJapanesePhrase,
		language.Korean:               cousteauKoreanPhrase,
		language.LatinAmericanSpanish: cousteauLatinAmericanSpanishPhrase,
		language.Russian:              cousteauRussianPhrase,
		language.SimplifiedChinese:    cousteauSimplifiedChinesePhrase,
		language.Spanish:              cousteauSpanishPhrase,
		language.TraditionalChinese:   cousteauTraditionalChinesePhrase}
)

var (
	// Cousteau represents cousteau.
	Cousteau = nook.Villager{
		Character:   cousteauCharacter,
		Personality: personality.Jock,
		Phrase:      cousteauPhrase}
)
