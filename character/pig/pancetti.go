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
		Value:    "Lydie"}

	pancettiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pancetti"}

	pancettiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lydie"}

	pancettiGermanName = nook.Name{
		Language: language.German,
		Value:    "Brigitte"}

	pancettiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cicciola"}

	pancettiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ブリトニー"}

	pancettiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "브리트니"}

	pancettiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Talía"}

	pancettiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Панчетти"}

	pancettiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "布兰妮"}

	pancettiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Talía"}

	pancettiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "布蘭妮"}
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
		Animal:   animal.Pig,
		Birthday: pancettiBirthday,
		Code:     pancettiCode,
		Key:      character.Pancetti,
		Gender:   gender.Female,
		Name:     pancettiName,
		Special:  false}
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
		language.AmericanEnglish:      pancettiAmericanEnglishPhrase,
		language.CanadianFrench:       pancettiCanadianFrenchPhrase,
		language.Dutch:                pancettiDutchPhrase,
		language.French:               pancettiFrenchPhrase,
		language.German:               pancettiGermanPhrase,
		language.Italian:              pancettiItalianPhrase,
		language.Japanese:             pancettiJapanesePhrase,
		language.Korean:               pancettiKoreanPhrase,
		language.LatinAmericanSpanish: pancettiLatinAmericanSpanishPhrase,
		language.Russian:              pancettiRussianPhrase,
		language.SimplifiedChinese:    pancettiSimplifiedChinesePhrase,
		language.Spanish:              pancettiSpanishPhrase,
		language.TraditionalChinese:   pancettiTraditionalChinesePhrase}
)

var (
	Pancetti = nook.Villager{
		Character:   pancettiCharacter,
		Personality: personality.Snooty,
		Phrase:      pancettiPhrase}
)
