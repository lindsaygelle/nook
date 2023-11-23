package rabbit

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
	// carmenBirthday represents Carmen's birthday.
	carmenBirthday = nook.Birthday{
		Day:   6,
		Month: time.January}
)

var (
	// carmenCode represents Carmen's unique code.
	carmenCode = nook.Code{
		Value: "rbt16"}
)

var (
	// carmenAmericanEnglishName represents Carmen's name in American English.
	carmenAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Carmen"}

	// carmenCanadianFrenchName represents Carmen's name in Canadian French.
	carmenCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Zoé"}

	// carmenDutchName represents Carmen's name in Dutch.
	carmenDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Carmen"}

	// carmenFrenchName represents Carmen's name in French.
	carmenFrenchName = nook.Name{
		Language: language.French,
		Value:    "Zoé"}

	// carmenGermanName represents Carmen's name in German.
	carmenGermanName = nook.Name{
		Language: language.German,
		Value:    "Hilda"}

	// carmenItalianName represents Carmen's name in Italian.
	carmenItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Conny"}

	// carmenJapaneseName represents Carmen's name in Japanese.
	carmenJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チョコ"}

	// carmenKoreanName represents Carmen's name in Korean.
	carmenKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "초코"}

	// carmenLatinAmericanSpanishName represents Carmen's name in Latin American Spanish.
	carmenLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Linda"}

	// carmenRussianName represents Carmen's name in Russian.
	carmenRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кармен"}

	// carmenSimplifiedChineseName represents Carmen's name in Simplified Chinese.
	carmenSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "巧蔻"}

	// carmenSpanishName represents Carmen's name in Spanish.
	carmenSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Linda"}

	// carmenTraditionalChineseName represents Carmen's name in Traditional Chinese.
	carmenTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "巧蔻"}
)

var (
	// carmenName represents Carmen's name in different languages.
	carmenName = nook.Languages{
		language.AmericanEnglish:      carmenAmericanEnglishName,
		language.CanadianFrench:       carmenCanadianFrenchName,
		language.Dutch:                carmenDutchName,
		language.French:               carmenFrenchName,
		language.German:               carmenGermanName,
		language.Italian:              carmenItalianName,
		language.Japanese:             carmenJapaneseName,
		language.Korean:               carmenKoreanName,
		language.LatinAmericanSpanish: carmenLatinAmericanSpanishName,
		language.Russian:              carmenRussianName,
		language.SimplifiedChinese:    carmenSimplifiedChineseName,
		language.Spanish:              carmenSpanishName,
		language.TraditionalChinese:   carmenTraditionalChineseName}
)

var (
	// carmenCharacter represents Carmen's character information.
	carmenCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: carmenBirthday,
		Code:     carmenCode,
		Key:      character.Carmen,
		Gender:   gender.Female,
		Name:     carmenName,
		Special:  false}
)

var (
	// carmenAmericanEnglishPhrase represents Carmen's phrase in American English.
	carmenAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nougat"}

	// carmenCanadianFrenchPhrase represents Carmen's phrase in Canadian French.
	carmenCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "nougat"}

	// carmenDutchPhrase represents Carmen's phrase in Dutch.
	carmenDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "noga"}

	// carmenFrenchPhrase represents Carmen's phrase in French.
	carmenFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pinou-nou"}

	// carmenGermanPhrase represents Carmen's phrase in German.
	carmenGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "langohr"}

	// carmenItalianPhrase represents Carmen's phrase in Italian.
	carmenItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gnuf gnuf"}

	// carmenJapanesePhrase represents Carmen's phrase in Japanese.
	carmenJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "まじで"}

	// carmenKoreanPhrase represents Carmen's phrase in Korean.
	carmenKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "진짜진짜"}

	// carmenLatinAmericanSpanishPhrase represents Carmen's phrase in Latin American Spanish.
	carmenLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "nufinuf"}

	// carmenRussianPhrase represents Carmen's phrase in Russian.
	carmenRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "карамелька"}

	// carmenSimplifiedChinesePhrase represents Carmen's phrase in Simplified Chinese.
	carmenSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "真的"}

	// carmenSpanishPhrase represents Carmen's phrase in Spanish.
	carmenSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "arcoíris"}

	// carmenTraditionalChinesePhrase represents Carmen's phrase in Traditional Chinese.
	carmenTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "真的"}
)

var (
	// carmenPhrase represents Carmen's phrase in different languages.
	carmenPhrase = nook.Languages{
		language.AmericanEnglish:      carmenAmericanEnglishPhrase,
		language.CanadianFrench:       carmenCanadianFrenchPhrase,
		language.Dutch:                carmenDutchPhrase,
		language.French:               carmenFrenchPhrase,
		language.German:               carmenGermanPhrase,
		language.Italian:              carmenItalianPhrase,
		language.Japanese:             carmenJapanesePhrase,
		language.Korean:               carmenKoreanPhrase,
		language.LatinAmericanSpanish: carmenLatinAmericanSpanishPhrase,
		language.Russian:              carmenRussianPhrase,
		language.SimplifiedChinese:    carmenSimplifiedChinesePhrase,
		language.Spanish:              carmenSpanishPhrase,
		language.TraditionalChinese:   carmenTraditionalChinesePhrase}
)

var (
	// Carmen represents the character Carmen with her complete information.
	Carmen = nook.Villager{
		Character:   carmenCharacter,
		Personality: personality.Peppy,
		Phrase:      carmenPhrase}
)
