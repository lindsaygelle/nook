package eagle

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	sterlingBirthday = nook.Birthday{
		Day:   11,
		Month: time.December}
)

var (
	sterlingCode = nook.Code{
		Value: "pbr07"}
)

var (
	sterlingAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sterling"}

	sterlingCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Manfredkaputt"}

	sterlingDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sterlinghavik"}

	sterlingFrenchName = nook.Name{
		Language: language.French,
		Value:    "Manfredkaputt"}

	sterlingGermanName = nook.Name{
		Language: language.German,
		Value:    "Horstkrahkrah"}

	sterlingItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Beccodìchicchirò"}

	sterlingJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ギンカクやあッ"}

	sterlingKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "은수리호옷"}

	sterlingLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Arniescróóó"}

	sterlingRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Стерлингклекот"}

	sterlingSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "银阁呀啊"}

	sterlingSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Arnimancuernas"}

	sterlingTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "銀閣呀啊"}
)

var (
	sterlingName = nook.Languages{
		language.AmericanEnglish:      sterlingAmericanEnglishName,
		language.CanadianFrench:       sterlingCanadianFrenchName,
		language.Dutch:                sterlingDutchName,
		language.French:               sterlingFrenchName,
		language.German:               sterlingGermanName,
		language.Italian:              sterlingItalianName,
		language.Japanese:             sterlingJapaneseName,
		language.Korean:               sterlingKoreanName,
		language.LatinAmericanSpanish: sterlingLatinAmericanSpanishName,
		language.Russian:              sterlingRussianName,
		language.SimplifiedChinese:    sterlingSimplifiedChineseName,
		language.Spanish:              sterlingSpanishName,
		language.TraditionalChinese:   sterlingTraditionalChineseName}
)

var (
	sterlingCharacter = nook.Character{
		Animal:   Eagle,
		Birthday: sterlingBirthday,
		Code:     sterlingCode,
		Gender:   nook.Male,
		Name:     sterlingName}
)

var (
	sterlingAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "skraaaaw"}

	sterlingCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "kaputt"}

	sterlingDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "havik"}

	sterlingFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "kaputt"}

	sterlingGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "krahkrah"}

	sterlingItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "chicchirò"}

	sterlingJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "やあッ"}

	sterlingKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "호옷"}

	sterlingLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "escróóó"}

	sterlingRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "клекот"}

	sterlingSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呀啊"}

	sterlingSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mancuernas"}

	sterlingTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "呀啊"}
)

var (
	sterlingPhrase = nook.Languages{
		language.AmericanEnglish:      sterlingAmericanEnglishName,
		language.CanadianFrench:       sterlingCanadianFrenchName,
		language.Dutch:                sterlingDutchName,
		language.French:               sterlingFrenchName,
		language.German:               sterlingGermanName,
		language.Italian:              sterlingItalianName,
		language.Japanese:             sterlingJapaneseName,
		language.Korean:               sterlingKoreanName,
		language.LatinAmericanSpanish: sterlingLatinAmericanSpanishName,
		language.Russian:              sterlingRussianName,
		language.SimplifiedChinese:    sterlingSimplifiedChineseName,
		language.Spanish:              sterlingSpanishName,
		language.TraditionalChinese:   sterlingTraditionalChineseName}
)

var (
	Sterling = nook.Villager{
		Character:   sterlingCharacter,
		Personality: nook.Jock,
		Phrase:      sterlingPhrase}
)
