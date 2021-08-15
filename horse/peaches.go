package horse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	peachesBirthday = nook.Birthday{
		Day:   28,
		Month: time.November}
)

var (
	peachesCode = nook.Code{
		Value: "hrs08"}
)

var (
	peachesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Peaches"}

	peachesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Prunemustang"}

	peachesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Peacheshinnik"}

	peachesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Prunemustang"}

	peachesGermanName = nook.Name{
		Language: language.German,
		Value:    "Clairenachbar"}

	peachesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ronzinaoh mamma"}

	peachesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ドサコだポン"}

	peachesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "말자몰라요"}

	peachesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Perlaniiiih"}

	peachesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пичессоседушка"}

	peachesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小薰蹦"}

	peachesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Perlaniiiih"}

	peachesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小薰蹦"}
)

var (
	peachesName = nook.Languages{
		language.AmericanEnglish:      peachesAmericanEnglishName,
		language.CanadianFrench:       peachesCanadianFrenchName,
		language.Dutch:                peachesDutchName,
		language.French:               peachesFrenchName,
		language.German:               peachesGermanName,
		language.Italian:              peachesItalianName,
		language.Japanese:             peachesJapaneseName,
		language.Korean:               peachesKoreanName,
		language.LatinAmericanSpanish: peachesLatinAmericanSpanishName,
		language.Russian:              peachesRussianName,
		language.SimplifiedChinese:    peachesSimplifiedChineseName,
		language.Spanish:              peachesSpanishName,
		language.TraditionalChinese:   peachesTraditionalChineseName}
)

var (
	peachesCharacter = nook.Character{
		Animal:   Horse,
		Birthday: peachesBirthday,
		Code:     peachesCode,
		Gender:   nook.Female,
		Name:     peachesName}
)

var (
	peachesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "neighbor"}

	peachesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mustang"}

	peachesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hinnik"}

	peachesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mustang"}

	peachesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "nachbar"}

	peachesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "oh mamma"}

	peachesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だポン"}

	peachesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "몰라요"}

	peachesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "niiiih"}

	peachesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "соседушка"}

	peachesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蹦"}

	peachesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "niiiih"}

	peachesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蹦"}
)

var (
	peachesPhrase = nook.Languages{
		language.AmericanEnglish:      peachesAmericanEnglishName,
		language.CanadianFrench:       peachesCanadianFrenchName,
		language.Dutch:                peachesDutchName,
		language.French:               peachesFrenchName,
		language.German:               peachesGermanName,
		language.Italian:              peachesItalianName,
		language.Japanese:             peachesJapaneseName,
		language.Korean:               peachesKoreanName,
		language.LatinAmericanSpanish: peachesLatinAmericanSpanishName,
		language.Russian:              peachesRussianName,
		language.SimplifiedChinese:    peachesSimplifiedChineseName,
		language.Spanish:              peachesSpanishName,
		language.TraditionalChinese:   peachesTraditionalChineseName}
)

var (
	Peaches = nook.Villager{
		Character:   peachesCharacter,
		Personality: nook.Normal,
		Phrase:      peachesPhrase}
)
