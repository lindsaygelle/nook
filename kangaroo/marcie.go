package kangaroo

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Selmawallaby"}

	marcieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Marciebuidels"}

	marcieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Selmawallaby"}

	marcieGermanName = nook.Name{
		Language: language.German,
		Value:    "Marlieshüpf"}

	marcieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Adelaideme oui"}

	marcieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マリアいいのよ"}

	marcieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마리아조아요"}

	marcieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Brisanainoná"}

	marcieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Марсикармашек"}

	marcieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "玛莉亚好唷"}

	marcieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Brisanainoná"}

	marcieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "瑪莉亞好唷"}
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
		Animal:   Kangaroo,
		Birthday: marcieBirthday,
		Code:     marcieCode,
		Gender:   nook.Female,
		Name:     marcieName}
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
	Marcie = nook.Villager{
		Character:   marcieCharacter,
		Personality: nook.Normal,
		Phrase:      marciePhrase}
)
