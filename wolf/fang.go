package wolf

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	fangBirthday = nook.Birthday{
		Day:   18,
		Month: time.December}
)

var (
	fangCode = nook.Code{
		Value: "wol06"}
)

var (
	fangAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Fang"}

	fangCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pierrotsagouin"}

	fangDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Fangknauw"}

	fangFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pierrotsagouin"}

	fangGermanName = nook.Name{
		Language: language.German,
		Value:    "Grimmbraul"}

	fangItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Zannafabuuula"}

	fangJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シベリアですダス"}

	fangKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "시베리아콜록콜록"}

	fangLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Colmilloen-ñac"}

	fangRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фанггрызь"}

	fangSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "史培亚这倒是"}

	fangSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Colmilloen-ñac"}

	fangTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "史培亞這倒是"}
)

var (
	fangName = nook.Languages{
		language.AmericanEnglish:      fangAmericanEnglishName,
		language.CanadianFrench:       fangCanadianFrenchName,
		language.Dutch:                fangDutchName,
		language.French:               fangFrenchName,
		language.German:               fangGermanName,
		language.Italian:              fangItalianName,
		language.Japanese:             fangJapaneseName,
		language.Korean:               fangKoreanName,
		language.LatinAmericanSpanish: fangLatinAmericanSpanishName,
		language.Russian:              fangRussianName,
		language.SimplifiedChinese:    fangSimplifiedChineseName,
		language.Spanish:              fangSpanishName,
		language.TraditionalChinese:   fangTraditionalChineseName}
)

var (
	fangCharacter = nook.Character{
		Animal:   Wolf,
		Birthday: fangBirthday,
		Code:     fangCode,
		Gender:   nook.Male,
		Name:     fangName}
)

var (
	fangAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cha-chomp"}

	fangCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "sagouin"}

	fangDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "knauw"}

	fangFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "sagouin"}

	fangGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "braul"}

	fangItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "fabuuula"}

	fangJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですダス"}

	fangKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "콜록콜록"}

	fangLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "en-ñac"}

	fangRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "грызь"}

	fangSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "这倒是"}

	fangSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "en-ñac"}

	fangTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "這倒是"}
)

var (
	fangPhrase = nook.Languages{
		language.AmericanEnglish:      fangAmericanEnglishName,
		language.CanadianFrench:       fangCanadianFrenchName,
		language.Dutch:                fangDutchName,
		language.French:               fangFrenchName,
		language.German:               fangGermanName,
		language.Italian:              fangItalianName,
		language.Japanese:             fangJapaneseName,
		language.Korean:               fangKoreanName,
		language.LatinAmericanSpanish: fangLatinAmericanSpanishName,
		language.Russian:              fangRussianName,
		language.SimplifiedChinese:    fangSimplifiedChineseName,
		language.Spanish:              fangSpanishName,
		language.TraditionalChinese:   fangTraditionalChineseName}
)

var (
	Fang = nook.Villager{
		Character:   fangCharacter,
		Personality: nook.Cranky,
		Phrase:      fangPhrase}
)
