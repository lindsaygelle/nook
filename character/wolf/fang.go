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
	// fangBirthday represents Fang's birthday (December 18th).
	fangBirthday = nook.Birthday{
		Day:   18,
		Month: time.December}
)

var (
	// fangCode represents Fang's unique code ("wol06").
	fangCode = nook.Code{
		Value: "wol06"}
)

var (
	// fangAmericanEnglishName represents Fang's name in American English.
	fangAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Fang"}

	// fangCanadianFrenchName represents Fang's name in Canadian French.
	fangCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pierrot"}

	// fangDutchName represents Fang's name in Dutch.
	fangDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Fang"}

	// fangFrenchName represents Fang's name in French.
	fangFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pierrot"}

	// fangGermanName represents Fang's name in German.
	fangGermanName = nook.Name{
		Language: language.German,
		Value:    "Grimm"}

	// fangItalianName represents Fang's name in Italian.
	fangItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Zanna"}

	// fangJapaneseName represents Fang's name in Japanese.
	fangJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シベリア"}

	// fangKoreanName represents Fang's name in Korean.
	fangKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "시베리아"}

	// fangLatinAmericanSpanishName represents Fang's name in Latin American Spanish.
	fangLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Colmillo"}

	// fangRussianName represents Fang's name in Russian.
	fangRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фанг"}

	// fangSimplifiedChineseName represents Fang's name in Simplified Chinese.
	fangSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "史培亚"}

	// fangSpanishName represents Fang's name in Spanish.
	fangSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Colmillo"}

	// fangTraditionalChineseName represents Fang's name in Traditional Chinese.
	fangTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "史培亞"}
)

var (
	// fangName represents Fang's name in different languages.
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
	// fangCharacter represents Fang's character information.
	fangCharacter = nook.Character{
		Animal:   animal.Wolf,
		Birthday: fangBirthday,
		Code:     fangCode,
		Key:      character.Fang,
		Gender:   gender.Male,
		Name:     fangName,
		Special:  false}
)

var (
	// fangAmericanEnglishPhrase represents Fang's phrase in American English.
	fangAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cha-chomp"}

	// fangCanadianFrenchPhrase represents Fang's phrase in Canadian French.
	fangCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "sagouin"}

	// fangDutchPhrase represents Fang's phrase in Dutch.
	fangDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "knauw"}

	// fangFrenchPhrase represents Fang's phrase in French.
	fangFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "sagouin"}

	// fangGermanPhrase represents Fang's phrase in German.
	fangGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "braul"}

	// fangItalianPhrase represents Fang's phrase in Italian.
	fangItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "fabuuula"}

	// fangJapanesePhrase represents Fang's phrase in Japanese.
	fangJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですダス"}

	// fangKoreanPhrase represents Fang's phrase in Korean.
	fangKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "콜록콜록"}

	// fangLatinAmericanSpanishPhrase represents Fang's phrase in Latin American Spanish.
	fangLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "en-ñac"}

	// fangRussianPhrase represents Fang's phrase in Russian.
	fangRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "грызь"}

	// fangSimplifiedChinesePhrase represents Fang's phrase in Simplified Chinese.
	fangSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "这倒是"}

	// fangSpanishPhrase represents Fang's phrase in Spanish.
	fangSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "en-ñac"}

	// fangTraditionalChinesePhrase represents Fang's phrase in Traditional Chinese.
	fangTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "這倒是"}
)

var (
	// fangPhrase represents Fang's phrase in different languages.
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
	// Fang represents the character Fang with his complete information.
	Fang = nook.Villager{
		Character:   fangCharacter,
		Personality: personality.Cranky,
		Phrase:      fangPhrase}
)
