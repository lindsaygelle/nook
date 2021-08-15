package elephant

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	axelBirthday = nook.Birthday{
		Day:   23,
		Month: time.March}
)

var (
	axelCode = nook.Code{
		Value: "elp06"}
)

var (
	axelAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Axel"}

	axelCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Axelsplaf"}

	axelDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "AxelPWAAP"}

	axelFrenchName = nook.Name{
		Language: language.French,
		Value:    "Axelsplaf"}

	axelGermanName = nook.Name{
		Language: language.German,
		Value:    "AxelTUUUUT"}

	axelItalianName = nook.Name{
		Language: language.Italian,
		Value:    "SandroSBONK"}

	axelJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エックスエルでゴンス"}

	axelKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "엑스엘리히힛"}

	axelLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Eustakioankagua"}

	axelRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Аксельпотрубим"}

	axelSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "大大嘻嘻"}

	axelSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Eustakioankagua"}

	axelTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "大大嘻嘻"}
)

var (
	axelName = nook.Languages{
		language.AmericanEnglish:      axelAmericanEnglishName,
		language.CanadianFrench:       axelCanadianFrenchName,
		language.Dutch:                axelDutchName,
		language.French:               axelFrenchName,
		language.German:               axelGermanName,
		language.Italian:              axelItalianName,
		language.Japanese:             axelJapaneseName,
		language.Korean:               axelKoreanName,
		language.LatinAmericanSpanish: axelLatinAmericanSpanishName,
		language.Russian:              axelRussianName,
		language.SimplifiedChinese:    axelSimplifiedChineseName,
		language.Spanish:              axelSpanishName,
		language.TraditionalChinese:   axelTraditionalChineseName}
)

var (
	axelCharacter = nook.Character{
		Animal:   Elephant,
		Birthday: axelBirthday,
		Code:     axelCode,
		Gender:   nook.Male,
		Name:     axelName}
)

var (
	axelAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "WHONK"}

	axelCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "splaf"}

	axelDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "PWAAP"}

	axelFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "splaf"}

	axelGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "TUUUUT"}

	axelItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "SBONK"}

	axelJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でゴンス"}

	axelKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "히힛"}

	axelLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ankagua"}

	axelRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "потрубим"}

	axelSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘻嘻"}

	axelSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ankagua"}

	axelTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘻嘻"}
)

var (
	axelPhrase = nook.Languages{
		language.AmericanEnglish:      axelAmericanEnglishName,
		language.CanadianFrench:       axelCanadianFrenchName,
		language.Dutch:                axelDutchName,
		language.French:               axelFrenchName,
		language.German:               axelGermanName,
		language.Italian:              axelItalianName,
		language.Japanese:             axelJapaneseName,
		language.Korean:               axelKoreanName,
		language.LatinAmericanSpanish: axelLatinAmericanSpanishName,
		language.Russian:              axelRussianName,
		language.SimplifiedChinese:    axelSimplifiedChineseName,
		language.Spanish:              axelSpanishName,
		language.TraditionalChinese:   axelTraditionalChineseName}
)

var (
	Axel = nook.Villager{
		Character:   axelCharacter,
		Personality: nook.Jock,
		Phrase:      axelPhrase}
)
