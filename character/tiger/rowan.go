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
	rowanBirthday = nook.Birthday{
		Day:   26,
		Month: time.August}
)

var (
	rowanCode = nook.Code{
		Value: "tig01"}
)

var (
	rowanAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rowan"}

	rowanCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Marito"}

	rowanDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rowan"}

	rowanFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marito"}

	rowanGermanName = nook.Name{
		Language: language.German,
		Value:    "Gisbert"}

	rowanItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ernesto"}

	rowanJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゴメス"}

	rowanKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "고메스"}

	rowanLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Miguelón"}

	rowanRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Роуэн"}

	rowanSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "戈麦斯"}

	rowanSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Miguelón"}

	rowanTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "戈麥斯"}
)

var (
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
	rowanCharacter = nook.Character{
		Animal:   animal.Tiger,
		Birthday: rowanBirthday,
		Code:     rowanCode,
		Key:      character.Rowan,
		Gender:   gender.Male,
		Name:     rowanName}
)

var (
	rowanAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "mango"}

	rowanCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "banane"}

	rowanDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mango"}

	rowanFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "banane"}

	rowanGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mango"}

	rowanItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "mango"}

	rowanJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "まったく"}

	rowanKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "내참"}

	rowanLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "gurruf"}

	rowanRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "манго"}

	rowanSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "完全"}

	rowanSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "turista"}

	rowanTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "完全"}
)

var (
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
	Rowan = nook.Villager{
		Character:   rowanCharacter,
		Personality: personality.Jock,
		Phrase:      rowanPhrase}
)
