package wolf

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
		Value:    "Pierrot"}

	fangDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Fang"}

	fangFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pierrot"}

	fangGermanName = nook.Name{
		Language: language.German,
		Value:    "Grimm"}

	fangItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Zanna"}

	fangJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シベリア"}

	fangKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "시베리아"}

	fangLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Colmillo"}

	fangRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фанг"}

	fangSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "史培亚"}

	fangSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Colmillo"}

	fangTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "史培亞"}
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
		Animal:   animal.Wolf,
		Birthday: fangBirthday,
		Code:     fangCode,
		Key:      character.Fang,
		Gender:   gender.Male,
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
		language.AmericanEnglish:      fangAmericanEnglishPhrase,
		language.CanadianFrench:       fangCanadianFrenchPhrase,
		language.Dutch:                fangDutchPhrase,
		language.French:               fangFrenchPhrase,
		language.German:               fangGermanPhrase,
		language.Italian:              fangItalianPhrase,
		language.Japanese:             fangJapanesePhrase,
		language.Korean:               fangKoreanPhrase,
		language.LatinAmericanSpanish: fangLatinAmericanSpanishPhrase,
		language.Russian:              fangRussianPhrase,
		language.SimplifiedChinese:    fangSimplifiedChinesePhrase,
		language.Spanish:              fangSpanishPhrase,
		language.TraditionalChinese:   fangTraditionalChinesePhrase}
)

var (
	Fang = nook.Villager{
		Character:   fangCharacter,
		Personality: personality.Cranky,
		Phrase:      fangPhrase}
)
