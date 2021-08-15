package horse

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Licowestern"}

	julianDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Julianglitter"}

	julianFrenchName = nook.Name{
		Language: language.French,
		Value:    "Licotrottoune"}

	julianGermanName = nook.Name{
		Language: language.German,
		Value:    "Jimmywihihihi"}

	julianItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Giulianovamos"}

	julianJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジュリーね、キミ"}

	julianKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "유니오그대여"}

	julianLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Azulinotrotót"}

	julianRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джулианблеск"}

	julianSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "朱黎喂！你"}

	julianSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Azulinopinchito"}

	julianTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "朱黎喂！你"}
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
		Animal:   Horse,
		Birthday: julianBirthday,
		Code:     julianCode,
		Gender:   nook.Male,
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
		Personality: nook.Smug,
		Phrase:      julianPhrase}
)
