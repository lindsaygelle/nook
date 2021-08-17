package pig

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
	kevinBirthday = nook.Birthday{
		Day:   26,
		Month: time.April}
)

var (
	kevinCode = nook.Code{
		Value: "pig15"}
)

var (
	kevinAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kevin"}

	kevinCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Jean-Bon"}

	kevinDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kevin"}

	kevinFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jean-Bon"}

	kevinGermanName = nook.Name{
		Language: language.German,
		Value:    "Kevin"}

	kevinItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Kotekiño"}

	kevinJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "イノッチ"}

	kevinKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "멧지"}

	kevinLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Porcinio"}

	kevinRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кевин"}

	kevinSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "伊比利"}

	kevinSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Porcinio"}

	kevinTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "伊比利"}
)

var (
	kevinName = nook.Languages{
		language.AmericanEnglish:      kevinAmericanEnglishName,
		language.CanadianFrench:       kevinCanadianFrenchName,
		language.Dutch:                kevinDutchName,
		language.French:               kevinFrenchName,
		language.German:               kevinGermanName,
		language.Italian:              kevinItalianName,
		language.Japanese:             kevinJapaneseName,
		language.Korean:               kevinKoreanName,
		language.LatinAmericanSpanish: kevinLatinAmericanSpanishName,
		language.Russian:              kevinRussianName,
		language.SimplifiedChinese:    kevinSimplifiedChineseName,
		language.Spanish:              kevinSpanishName,
		language.TraditionalChinese:   kevinTraditionalChineseName}
)

var (
	kevinCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: kevinBirthday,
		Code:     kevinCode,
		Key:      character.Kevin,
		Gender:   gender.Male,
		Name:     kevinName}
)

var (
	kevinAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "weeweewee"}

	kevinCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "toutébon"}

	kevinDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ieieieie"}

	kevinFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "toutébon"}

	kevinGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quiiiek"}

	kevinItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "boink"}

	kevinJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウリウリ"}

	kevinKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "냠냠"}

	kevinLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bidoink"}

	kevinRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "уи-и-и"}

	kevinSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "比利比利"}

	kevinSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bidoink"}

	kevinTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "比利比利"}
)

var (
	kevinPhrase = nook.Languages{
		language.AmericanEnglish:      kevinAmericanEnglishPhrase,
		language.CanadianFrench:       kevinCanadianFrenchPhrase,
		language.Dutch:                kevinDutchPhrase,
		language.French:               kevinFrenchPhrase,
		language.German:               kevinGermanPhrase,
		language.Italian:              kevinItalianPhrase,
		language.Japanese:             kevinJapanesePhrase,
		language.Korean:               kevinKoreanPhrase,
		language.LatinAmericanSpanish: kevinLatinAmericanSpanishPhrase,
		language.Russian:              kevinRussianPhrase,
		language.SimplifiedChinese:    kevinSimplifiedChinesePhrase,
		language.Spanish:              kevinSpanishPhrase,
		language.TraditionalChinese:   kevinTraditionalChinesePhrase}
)

var (
	Kevin = nook.Villager{
		Character:   kevinCharacter,
		Personality: personality.Jock,
		Phrase:      kevinPhrase}
)
