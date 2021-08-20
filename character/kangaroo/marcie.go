package kangaroo

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
	marcieBirthday = nook.Birthday{
		Day:   31,
		Month: time.May}
)

var (
	marcieCode = nook.Code{
		Value: "kgr10"}
)

var (
	marcieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Marcie"}

	marcieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Selma"}

	marcieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Marcie"}

	marcieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Selma"}

	marcieGermanName = nook.Name{
		Language: language.German,
		Value:    "Marlies"}

	marcieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Adelaide"}

	marcieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マリア"}

	marcieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마리아"}

	marcieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Brisa"}

	marcieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Марси"}

	marcieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "玛莉亚"}

	marcieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Brisa"}

	marcieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "瑪莉亞"}
)

var (
	marcieName = nook.Languages{
		language.AmericanEnglish:      marcieAmericanEnglishName,
		language.CanadianFrench:       marcieCanadianFrenchName,
		language.Dutch:                marcieDutchName,
		language.French:               marcieFrenchName,
		language.German:               marcieGermanName,
		language.Italian:              marcieItalianName,
		language.Japanese:             marcieJapaneseName,
		language.Korean:               marcieKoreanName,
		language.LatinAmericanSpanish: marcieLatinAmericanSpanishName,
		language.Russian:              marcieRussianName,
		language.SimplifiedChinese:    marcieSimplifiedChineseName,
		language.Spanish:              marcieSpanishName,
		language.TraditionalChinese:   marcieTraditionalChineseName}
)

var (
	marcieCharacter = nook.Character{
		Animal:   animal.Kangaroo,
		Birthday: marcieBirthday,
		Code:     marcieCode,
		Key:      character.Marcie,
		Gender:   gender.Female,
		Name:     marcieName,
		Special:  false}
)

var (
	marcieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pouches"}

	marcieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "wallaby"}

	marcieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "buidels"}

	marcieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "wallaby"}

	marcieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hüpf"}

	marcieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "me oui"}

	marcieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "いいのよ"}

	marcieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "조아요"}

	marcieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "nainoná"}

	marcieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кармашек"}

	marcieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "好唷"}

	marcieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "nainoná"}

	marcieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "好唷"}
)

var (
	marciePhrase = nook.Languages{
		language.AmericanEnglish:      marcieAmericanEnglishPhrase,
		language.CanadianFrench:       marcieCanadianFrenchPhrase,
		language.Dutch:                marcieDutchPhrase,
		language.French:               marcieFrenchPhrase,
		language.German:               marcieGermanPhrase,
		language.Italian:              marcieItalianPhrase,
		language.Japanese:             marcieJapanesePhrase,
		language.Korean:               marcieKoreanPhrase,
		language.LatinAmericanSpanish: marcieLatinAmericanSpanishPhrase,
		language.Russian:              marcieRussianPhrase,
		language.SimplifiedChinese:    marcieSimplifiedChinesePhrase,
		language.Spanish:              marcieSpanishPhrase,
		language.TraditionalChinese:   marcieTraditionalChinesePhrase}
)

var (
	Marcie = nook.Villager{
		Character:   marcieCharacter,
		Personality: personality.Normal,
		Phrase:      marciePhrase}
)
