package duck

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	ketchupBirthday = nook.Birthday{
		Day:   27,
		Month: time.July}
)

var (
	ketchupCode = nook.Code{
		Value: "duk13"}
)

var (
	ketchupAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ketchup"}

	ketchupCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ketchup"}

	ketchupDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ketchup"}

	ketchupFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ketchup"}

	ketchupGermanName = nook.Name{
		Language: language.German,
		Value:    "Pullunda"}

	ketchupItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ketchup"}

	ketchupJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ケチャップ"}

	ketchupKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "케첩"}

	ketchupLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Kétchup"}

	ketchupRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кетчуп"}

	ketchupSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "番茄酱"}

	ketchupSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Kétchup"}

	ketchupTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "番茄醬"}
)

var (
	ketchupName = nook.Languages{
		language.AmericanEnglish:      ketchupAmericanEnglishName,
		language.CanadianFrench:       ketchupCanadianFrenchName,
		language.Dutch:                ketchupDutchName,
		language.French:               ketchupFrenchName,
		language.German:               ketchupGermanName,
		language.Italian:              ketchupItalianName,
		language.Japanese:             ketchupJapaneseName,
		language.Korean:               ketchupKoreanName,
		language.LatinAmericanSpanish: ketchupLatinAmericanSpanishName,
		language.Russian:              ketchupRussianName,
		language.SimplifiedChinese:    ketchupSimplifiedChineseName,
		language.Spanish:              ketchupSpanishName,
		language.TraditionalChinese:   ketchupTraditionalChineseName}
)

var (
	ketchupCharacter = nook.Character{
		Animal:   Duck,
		Birthday: ketchupBirthday,
		Code:     ketchupCode,
		Gender:   nook.Female,
		Name:     ketchupName}
)

var (
	ketchupAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bitty"}

	ketchupCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "plouf"}

	ketchupDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "tomaatje"}

	ketchupFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "plouf"}

	ketchupGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "pullilol"}

	ketchupItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quaraquack"}

	ketchupJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "プチ"}

	ketchupKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "찌익"}

	ketchupLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cuic-cuic"}

	ketchupRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "покрякушка"}

	ketchupSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "噗吱"}

	ketchupSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cuic-cuic"}

	ketchupTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "噗吱"}
)

var (
	ketchupPhrase = nook.Languages{
		language.AmericanEnglish:      ketchupAmericanEnglishName,
		language.CanadianFrench:       ketchupCanadianFrenchName,
		language.Dutch:                ketchupDutchName,
		language.French:               ketchupFrenchName,
		language.German:               ketchupGermanName,
		language.Italian:              ketchupItalianName,
		language.Japanese:             ketchupJapaneseName,
		language.Korean:               ketchupKoreanName,
		language.LatinAmericanSpanish: ketchupLatinAmericanSpanishName,
		language.Russian:              ketchupRussianName,
		language.SimplifiedChinese:    ketchupSimplifiedChineseName,
		language.Spanish:              ketchupSpanishName,
		language.TraditionalChinese:   ketchupTraditionalChineseName}
)

var (
	Ketchup = nook.Villager{
		Character:   ketchupCharacter,
		Personality: nook.Peppy,
		Phrase:      ketchupPhrase}
)
