package bearcub

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
	// ponchoBirthday represents poncho birthday.
	ponchoBirthday = nook.Birthday{
		Day:   2,
		Month: time.January}
)

var (
	// ponchoCode represents poncho code.
	ponchoCode = nook.Code{
		Value: "cbr02"}
)

var (
	// ponchoAmericanEnglishName represents poncho american english name.
	ponchoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Poncho"}

	// ponchoCanadianFrenchName represents poncho canadian french name.
	ponchoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Théo"}

	// ponchoDutchName represents poncho dutch name.
	ponchoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Poncho"}

	// ponchoFrenchName represents poncho french name.
	ponchoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Théo"}

	// ponchoGermanName represents poncho german name.
	ponchoGermanName = nook.Name{
		Language: language.German,
		Value:    "Toni"}

	// ponchoItalianName represents poncho italian name.
	ponchoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Poncho"}

	// ponchoJapaneseName represents poncho japanese name.
	ponchoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ポンチョ"}

	// ponchoKoreanName represents poncho korean name.
	ponchoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "봉추"}

	// ponchoLatinAmericanSpanishName represents poncho latin american spanish name.
	ponchoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Poncho"}

	// ponchoRussianName represents poncho russian name.
	ponchoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пончо"}

	// ponchoSimplifiedChineseName represents poncho simplified chinese name.
	ponchoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蓬蓬"}

	// ponchoSpanishName represents poncho spanish name.
	ponchoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Poncho"}

	// ponchoTraditionalChineseName represents poncho traditional chinese name.
	ponchoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蓬蓬"}
)

var (
	// ponchoName represents poncho name.
	ponchoName = nook.Languages{
		language.AmericanEnglish:      ponchoAmericanEnglishName,
		language.CanadianFrench:       ponchoCanadianFrenchName,
		language.Dutch:                ponchoDutchName,
		language.French:               ponchoFrenchName,
		language.German:               ponchoGermanName,
		language.Italian:              ponchoItalianName,
		language.Japanese:             ponchoJapaneseName,
		language.Korean:               ponchoKoreanName,
		language.LatinAmericanSpanish: ponchoLatinAmericanSpanishName,
		language.Russian:              ponchoRussianName,
		language.SimplifiedChinese:    ponchoSimplifiedChineseName,
		language.Spanish:              ponchoSpanishName,
		language.TraditionalChinese:   ponchoTraditionalChineseName}
)

var (
	// ponchoCharacter represents poncho character.
	ponchoCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: ponchoBirthday,
		Code:     ponchoCode,
		Key:      character.Poncho,
		Gender:   gender.Male,
		Name:     ponchoName,
		Special:  false}
)

var (
	// ponchoAmericanEnglishPhrase represents poncho american english phrase.
	ponchoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "li'l bear"}

	// ponchoCanadianFrenchPhrase represents poncho canadian french phrase.
	ponchoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "nounours"}

	// ponchoDutchPhrase represents poncho dutch phrase.
	ponchoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "beertje"}

	// ponchoFrenchPhrase represents poncho french phrase.
	ponchoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "nounours"}

	// ponchoGermanPhrase represents poncho german phrase.
	ponchoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bärli"}

	// ponchoItalianPhrase represents poncho italian phrase.
	ponchoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "babà"}

	// ponchoJapanesePhrase represents poncho japanese phrase.
	ponchoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "モン"}

	// ponchoKoreanPhrase represents poncho korean phrase.
	ponchoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "몽"}

	// ponchoLatinAmericanSpanishPhrase represents poncho latin american spanish phrase.
	ponchoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "pruah"}

	// ponchoRussianPhrase represents poncho russian phrase.
	ponchoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "медвежонок"}

	// ponchoSimplifiedChinesePhrase represents poncho simplified chinese phrase.
	ponchoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "萌"}

	// ponchoSpanishPhrase represents poncho spanish phrase.
	ponchoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bolita"}

	// ponchoTraditionalChinesePhrase represents poncho traditional chinese phrase.
	ponchoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "萌"}
)

var (
	// ponchoPhrase represents poncho phrase.
	ponchoPhrase = nook.Languages{
		language.AmericanEnglish:      ponchoAmericanEnglishPhrase,
		language.CanadianFrench:       ponchoCanadianFrenchPhrase,
		language.Dutch:                ponchoDutchPhrase,
		language.French:               ponchoFrenchPhrase,
		language.German:               ponchoGermanPhrase,
		language.Italian:              ponchoItalianPhrase,
		language.Japanese:             ponchoJapanesePhrase,
		language.Korean:               ponchoKoreanPhrase,
		language.LatinAmericanSpanish: ponchoLatinAmericanSpanishPhrase,
		language.Russian:              ponchoRussianPhrase,
		language.SimplifiedChinese:    ponchoSimplifiedChinesePhrase,
		language.Spanish:              ponchoSpanishPhrase,
		language.TraditionalChinese:   ponchoTraditionalChinesePhrase}
)

var (
	// Poncho represents poncho.
	Poncho = nook.Villager{
		Character:   ponchoCharacter,
		Personality: personality.Jock,
		Phrase:      ponchoPhrase}
)
