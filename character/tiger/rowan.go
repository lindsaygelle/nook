package tiger

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
	// rowanBirthday represents Rowan's birthday.
	rowanBirthday = nook.Birthday{
		Day:   26,
		Month: time.August}
)

var (
	// rowanCode represents Rowan's unique code.
	rowanCode = nook.Code{
		Value: "tig01"}
)

var (
	// rowanAmericanEnglishName represents Rowan's name in American English.
	rowanAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rowan"}

	// rowanCanadianFrenchName represents Rowan's name in Canadian French.
	rowanCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Marito"}

	// rowanDutchName represents Rowan's name in Dutch.
	rowanDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rowan"}

	// rowanFrenchName represents Rowan's name in French.
	rowanFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marito"}

	// rowanGermanName represents Rowan's name in German.
	rowanGermanName = nook.Name{
		Language: language.German,
		Value:    "Gisbert"}

	// rowanItalianName represents Rowan's name in Italian.
	rowanItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ernesto"}

	// rowanJapaneseName represents Rowan's name in Japanese.
	rowanJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゴメス"}

	// rowanKoreanName represents Rowan's name in Korean.
	rowanKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "고메스"}

	// rowanLatinAmericanSpanishName represents Rowan's name in Latin American Spanish.
	rowanLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Miguelón"}

	// rowanRussianName represents Rowan's name in Russian.
	rowanRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Роуэн"}

	// rowanSimplifiedChineseName represents Rowan's name in Simplified Chinese.
	rowanSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "戈麦斯"}

	// rowanSpanishName represents Rowan's name in Spanish.
	rowanSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Miguelón"}

	// rowanTraditionalChineseName represents Rowan's name in Traditional Chinese.
	rowanTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "戈麥斯"}
)

var (
	// rowanName represents Rowan's name in different languages.
	rowanName = nook.Languages{
		language.AmericanEnglish:      rowanAmericanEnglishName,
		language.CanadianFrench:       rowanCanadianFrenchName,
		language.Dutch:                rowanDutchName,
		language.French:               rowanFrenchName,
		language.German:               rowanGermanName,
		language.Italian:              rowanItalianName,
		language.Japanese:             rowanJapaneseName,
		language.Korean:               rowanKoreanName,
		language.LatinAmericanSpanish: rowanLatinAmericanSpanishName,
		language.Russian:              rowanRussianName,
		language.SimplifiedChinese:    rowanSimplifiedChineseName,
		language.Spanish:              rowanSpanishName,
		language.TraditionalChinese:   rowanTraditionalChineseName}
)

var (
	// rowanCharacter represents Rowan's character information.
	rowanCharacter = nook.Character{
		Animal:   animal.Tiger,
		Birthday: rowanBirthday,
		Code:     rowanCode,
		Key:      character.Rowan,
		Gender:   gender.Male,
		Name:     rowanName,
		Special:  false}
)

var (
	// rowanAmericanEnglishPhrase represents Rowan's phrase in American English.
	rowanAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "mango"}

	// rowanCanadianFrenchPhrase represents Rowan's phrase in Canadian French.
	rowanCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "banane"}

	// rowanDutchPhrase represents Rowan's phrase in Dutch.
	rowanDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mango"}

	// rowanFrenchPhrase represents Rowan's phrase in French.
	rowanFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "banane"}

	// rowanGermanPhrase represents Rowan's phrase in German.
	rowanGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mango"}

	// rowanItalianPhrase represents Rowan's phrase in Italian.
	rowanItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "mango"}

	// rowanJapanesePhrase represents Rowan's phrase in Japanese.
	rowanJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "まったく"}

	// rowanKoreanPhrase represents Rowan's phrase in Korean.
	rowanKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "내참"}

	// rowanLatinAmericanSpanishPhrase represents Rowan's phrase in Latin American Spanish.
	rowanLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "gurruf"}

	// rowanRussianPhrase represents Rowan's phrase in Russian.
	rowanRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "манго"}

	// rowanSimplifiedChinesePhrase represents Rowan's phrase in Simplified Chinese.
	rowanSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "完全"}

	// rowanSpanishPhrase represents Rowan's phrase in Spanish.
	rowanSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "turista"}

	// rowanTraditionalChinesePhrase represents Rowan's phrase in Traditional Chinese.
	rowanTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "完全"}
)

var (
	// rowanPhrase represents Rowan's phrase in different languages.
	rowanPhrase = nook.Languages{
		language.AmericanEnglish:      rowanAmericanEnglishPhrase,
		language.CanadianFrench:       rowanCanadianFrenchPhrase,
		language.Dutch:                rowanDutchPhrase,
		language.French:               rowanFrenchPhrase,
		language.German:               rowanGermanPhrase,
		language.Italian:              rowanItalianPhrase,
		language.Japanese:             rowanJapanesePhrase,
		language.Korean:               rowanKoreanPhrase,
		language.LatinAmericanSpanish: rowanLatinAmericanSpanishPhrase,
		language.Russian:              rowanRussianPhrase,
		language.SimplifiedChinese:    rowanSimplifiedChinesePhrase,
		language.Spanish:              rowanSpanishPhrase,
		language.TraditionalChinese:   rowanTraditionalChinesePhrase}
)

var (
	// Rowan represents the character Rowan with his complete information.
	Rowan = nook.Villager{
		Character:   rowanCharacter,
		Personality: personality.Jock,
		Phrase:      rowanPhrase}
)
