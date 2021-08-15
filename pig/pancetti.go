package pig

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	pancettiBirthday = nook.Birthday{
		Day:   14,
		Month: time.November}
)

var (
	pancettiCode = nook.Code{
		Value: "pig16"}
)

var (
	pancettiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pancetti"}

	pancettiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lydiesabots"}

	pancettiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pancettikom kom"}

	pancettiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lydiesabots"}

	pancettiGermanName = nook.Name{
		Language: language.German,
		Value:    "Brigittegrubbel"}

	pancettiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cicciolagrufidù"}

	pancettiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ブリトニーやだも～"}

	pancettiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "브리트니어우야"}

	pancettiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Talíacuinoink"}

	pancettiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Панчеттихрюшка"}

	pancettiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "布兰妮讨厌啦"}

	pancettiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Talíamorrito"}

	pancettiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "布蘭妮討厭啦"}
)

var (
	pancettiName = nook.Languages{
		language.AmericanEnglish:      pancettiAmericanEnglishName,
		language.CanadianFrench:       pancettiCanadianFrenchName,
		language.Dutch:                pancettiDutchName,
		language.French:               pancettiFrenchName,
		language.German:               pancettiGermanName,
		language.Italian:              pancettiItalianName,
		language.Japanese:             pancettiJapaneseName,
		language.Korean:               pancettiKoreanName,
		language.LatinAmericanSpanish: pancettiLatinAmericanSpanishName,
		language.Russian:              pancettiRussianName,
		language.SimplifiedChinese:    pancettiSimplifiedChineseName,
		language.Spanish:              pancettiSpanishName,
		language.TraditionalChinese:   pancettiTraditionalChineseName}
)

var (
	pancettiCharacter = nook.Character{
		Animal:   Pig,
		Birthday: pancettiBirthday,
		Code:     pancettiCode,
		Gender:   nook.Female,
		Name:     pancettiName}
)

var (
	pancettiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sooey"}

	pancettiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "sabots"}

	pancettiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kom kom"}

	pancettiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "sabots"}

	pancettiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grubbel"}

	pancettiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grufidù"}

	pancettiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "やだも～"}

	pancettiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어우야"}

	pancettiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cuinoink"}

	pancettiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрюшка"}

	pancettiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "讨厌啦"}

	pancettiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "morrito"}

	pancettiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "討厭啦"}
)

var (
	pancettiPhrase = nook.Languages{
		language.AmericanEnglish:      pancettiAmericanEnglishName,
		language.CanadianFrench:       pancettiCanadianFrenchName,
		language.Dutch:                pancettiDutchName,
		language.French:               pancettiFrenchName,
		language.German:               pancettiGermanName,
		language.Italian:              pancettiItalianName,
		language.Japanese:             pancettiJapaneseName,
		language.Korean:               pancettiKoreanName,
		language.LatinAmericanSpanish: pancettiLatinAmericanSpanishName,
		language.Russian:              pancettiRussianName,
		language.SimplifiedChinese:    pancettiSimplifiedChineseName,
		language.Spanish:              pancettiSpanishName,
		language.TraditionalChinese:   pancettiTraditionalChineseName}
)

var (
	Pancetti = nook.Villager{
		Character:   pancettiCharacter,
		Personality: nook.Snooty,
		Phrase:      pancettiPhrase}
)
