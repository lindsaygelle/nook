package monkey

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	eliseBirthday = nook.Birthday{
		Day:   21,
		Month: time.March}
)

var (
	eliseCode = nook.Code{
		Value: "mnk05"}
)

var (
	eliseAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Elise"}

	eliseCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Élisesteup"}

	eliseDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Elisekom op zeg"}

	eliseFrenchName = nook.Name{
		Language: language.French,
		Value:    "Élisesteup"}

	eliseGermanName = nook.Name{
		Language: language.German,
		Value:    "Steffistrebstreb"}

	eliseItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Elisaper favore"}

	eliseJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "モンこだモン"}

	eliseKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "몽자몽몽"}

	eliseLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mayrauh-lalá"}

	eliseRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Элизасамо собой"}

	eliseSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "孟琪孟孟"}

	eliseSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mayraricura"}

	eliseTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "孟琪孟孟"}
)

var (
	eliseName = nook.Languages{
		language.AmericanEnglish:      eliseAmericanEnglishName,
		language.CanadianFrench:       eliseCanadianFrenchName,
		language.Dutch:                eliseDutchName,
		language.French:               eliseFrenchName,
		language.German:               eliseGermanName,
		language.Italian:              eliseItalianName,
		language.Japanese:             eliseJapaneseName,
		language.Korean:               eliseKoreanName,
		language.LatinAmericanSpanish: eliseLatinAmericanSpanishName,
		language.Russian:              eliseRussianName,
		language.SimplifiedChinese:    eliseSimplifiedChineseName,
		language.Spanish:              eliseSpanishName,
		language.TraditionalChinese:   eliseTraditionalChineseName}
)

var (
	eliseCharacter = nook.Character{
		Animal:   Monkey,
		Birthday: eliseBirthday,
		Code:     eliseCode,
		Gender:   nook.Female,
		Name:     eliseName}
)

var (
	eliseAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "puh-lease"}

	eliseCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "steup"}

	eliseDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kom op zeg"}

	eliseFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "steup"}

	eliseGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "strebstreb"}

	eliseItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "per favore"}

	eliseJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だモン"}

	eliseKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "몽몽"}

	eliseLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "uh-lalá"}

	eliseRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "само собой"}

	eliseSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "孟孟"}

	eliseSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ricura"}

	eliseTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "孟孟"}
)

var (
	elisePhrase = nook.Languages{
		language.AmericanEnglish:      eliseAmericanEnglishName,
		language.CanadianFrench:       eliseCanadianFrenchName,
		language.Dutch:                eliseDutchName,
		language.French:               eliseFrenchName,
		language.German:               eliseGermanName,
		language.Italian:              eliseItalianName,
		language.Japanese:             eliseJapaneseName,
		language.Korean:               eliseKoreanName,
		language.LatinAmericanSpanish: eliseLatinAmericanSpanishName,
		language.Russian:              eliseRussianName,
		language.SimplifiedChinese:    eliseSimplifiedChineseName,
		language.Spanish:              eliseSpanishName,
		language.TraditionalChinese:   eliseTraditionalChineseName}
)

var (
	Elise = nook.Villager{
		Character:   eliseCharacter,
		Personality: nook.Snooty,
		Phrase:      elisePhrase}
)
