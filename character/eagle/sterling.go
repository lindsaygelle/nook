package eagle

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
		Value:    "Manfred"}

	sterlingDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sterling"}

	sterlingFrenchName = nook.Name{
		Language: language.French,
		Value:    "Manfred"}

	sterlingGermanName = nook.Name{
		Language: language.German,
		Value:    "Horst"}

	sterlingItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Beccodì"}

	sterlingJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ギンカク"}

	sterlingKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "은수리"}

	sterlingLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Arni"}

	sterlingRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Стерлинг"}

	sterlingSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "银阁"}

	sterlingSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Arni"}

	sterlingTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "銀閣"}
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
		Animal:   animal.Eagle,
		Birthday: sterlingBirthday,
		Code:     sterlingCode,
		Key:      character.Sterling,
		Gender:   gender.Male,
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
		Personality: personality.Jock,
		Phrase:      sterlingPhrase}
)
