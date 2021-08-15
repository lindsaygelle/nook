package tiger

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Maritobanane"}

	rowanDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rowanmango"}

	rowanFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maritobanane"}

	rowanGermanName = nook.Name{
		Language: language.German,
		Value:    "Gisbertmango"}

	rowanItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ernestomango"}

	rowanJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゴメスまったく"}

	rowanKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "고메스내참"}

	rowanLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Miguelóngurruf"}

	rowanRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Роуэнманго"}

	rowanSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "戈麦斯完全"}

	rowanSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Miguelónturista"}

	rowanTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "戈麥斯完全"}
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
		Animal:   Tiger,
		Birthday: rowanBirthday,
		Code:     rowanCode,
		Gender:   nook.Male,
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
	Rowan = nook.Villager{
		Character:   rowanCharacter,
		Personality: nook.Jock,
		Phrase:      rowanPhrase}
)
