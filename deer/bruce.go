package deer

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	bruceBirthday = nook.Birthday{
		Day:   26,
		Month: time.May}
)

var (
	bruceCode = nook.Code{
		Value: "der03"}
)

var (
	bruceAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bruce"}

	bruceCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Boubou"}

	bruceDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bruce"}

	bruceFrenchName = nook.Name{
		Language: language.French,
		Value:    "Boubou"}

	bruceGermanName = nook.Name{
		Language: language.German,
		Value:    "Oswald"}

	bruceItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Arturo"}

	bruceJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ブルース"}

	bruceKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "브루스"}

	bruceLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Aristo"}

	bruceRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Брюс"}

	bruceSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "布鹿斯"}

	bruceSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Aristo"}

	bruceTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "布鹿斯"}
)

var (
	bruceName = nook.Languages{
		language.AmericanEnglish:      bruceAmericanEnglishName,
		language.CanadianFrench:       bruceCanadianFrenchName,
		language.Dutch:                bruceDutchName,
		language.French:               bruceFrenchName,
		language.German:               bruceGermanName,
		language.Italian:              bruceItalianName,
		language.Japanese:             bruceJapaneseName,
		language.Korean:               bruceKoreanName,
		language.LatinAmericanSpanish: bruceLatinAmericanSpanishName,
		language.Russian:              bruceRussianName,
		language.SimplifiedChinese:    bruceSimplifiedChineseName,
		language.Spanish:              bruceSpanishName,
		language.TraditionalChinese:   bruceTraditionalChineseName}
)

var (
	bruceCharacter = nook.Character{
		Animal:   Deer,
		Birthday: bruceBirthday,
		Code:     bruceCode,
		Gender:   nook.Male,
		Name:     bruceName}
)

var (
	bruceAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "gruff"}

	bruceCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ouain"}

	bruceDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "burl"}

	bruceFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "nunul"}

	bruceGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "röööhr"}

	bruceItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "uauwaii"}

	bruceJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "しかしな"}

	bruceKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아비옹"}

	bruceLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "aquesí"}

	bruceRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "забияка"}

	bruceSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "福鹿寿"}

	bruceSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "aquesí"}

	bruceTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "福鹿壽"}
)

var (
	brucePhrase = nook.Languages{
		language.AmericanEnglish:      bruceAmericanEnglishName,
		language.CanadianFrench:       bruceCanadianFrenchName,
		language.Dutch:                bruceDutchName,
		language.French:               bruceFrenchName,
		language.German:               bruceGermanName,
		language.Italian:              bruceItalianName,
		language.Japanese:             bruceJapaneseName,
		language.Korean:               bruceKoreanName,
		language.LatinAmericanSpanish: bruceLatinAmericanSpanishName,
		language.Russian:              bruceRussianName,
		language.SimplifiedChinese:    bruceSimplifiedChineseName,
		language.Spanish:              bruceSpanishName,
		language.TraditionalChinese:   bruceTraditionalChineseName}
)

var (
	Bruce = nook.Villager{
		Character:   bruceCharacter,
		Personality: nook.Cranky,
		Phrase:      brucePhrase}
)
