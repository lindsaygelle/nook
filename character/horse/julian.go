package horse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	julianBirthday = nook.Birthday{
		Day:   15,
		Month: time.March}
)

var (
	julianCode = nook.Code{
		Value: "hrs13"}
)

var (
	julianAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Julian"}

	julianCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lico"}

	julianDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Julian"}

	julianFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lico"}

	julianGermanName = nook.Name{
		Language: language.German,
		Value:    "Jimmy"}

	julianItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Giuliano"}

	julianJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジュリー"}

	julianKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "유니오"}

	julianLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Azulino"}

	julianRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джулиан"}

	julianSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "朱黎"}

	julianSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Azulino"}

	julianTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "朱黎"}
)

var (
	julianName = nook.Languages{
		language.AmericanEnglish:      julianAmericanEnglishName,
		language.CanadianFrench:       julianCanadianFrenchName,
		language.Dutch:                julianDutchName,
		language.French:               julianFrenchName,
		language.German:               julianGermanName,
		language.Italian:              julianItalianName,
		language.Japanese:             julianJapaneseName,
		language.Korean:               julianKoreanName,
		language.LatinAmericanSpanish: julianLatinAmericanSpanishName,
		language.Russian:              julianRussianName,
		language.SimplifiedChinese:    julianSimplifiedChineseName,
		language.Spanish:              julianSpanishName,
		language.TraditionalChinese:   julianTraditionalChineseName}
)

var (
	julianCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: julianBirthday,
		Code:     julianCode,
		Gender:   gender.Male,
		Name:     julianName}
)

var (
	julianAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "glitter"}

	julianCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "western"}

	julianDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "glitter"}

	julianFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "trottoune"}

	julianGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "wihihihi"}

	julianItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "vamos"}

	julianJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ね、キミ"}

	julianKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그대여"}

	julianLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "trotót"}

	julianRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "блеск"}

	julianSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喂！你"}

	julianSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pinchito"}

	julianTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喂！你"}
)

var (
	julianPhrase = nook.Languages{
		language.AmericanEnglish:      julianAmericanEnglishName,
		language.CanadianFrench:       julianCanadianFrenchName,
		language.Dutch:                julianDutchName,
		language.French:               julianFrenchName,
		language.German:               julianGermanName,
		language.Italian:              julianItalianName,
		language.Japanese:             julianJapaneseName,
		language.Korean:               julianKoreanName,
		language.LatinAmericanSpanish: julianLatinAmericanSpanishName,
		language.Russian:              julianRussianName,
		language.SimplifiedChinese:    julianSimplifiedChineseName,
		language.Spanish:              julianSpanishName,
		language.TraditionalChinese:   julianTraditionalChineseName}
)

var (
	Julian = nook.Villager{
		Character:   julianCharacter,
		Personality: personality.Smug,
		Phrase:      julianPhrase}
)
