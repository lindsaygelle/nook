package elephant

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
		Value:    "Axel"}

	axelDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Axel"}

	axelFrenchName = nook.Name{
		Language: language.French,
		Value:    "Axel"}

	axelGermanName = nook.Name{
		Language: language.German,
		Value:    "Axel"}

	axelItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sandro"}

	axelJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エックスエル"}

	axelKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "엑스엘리"}

	axelLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Eustakio"}

	axelRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Аксель"}

	axelSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "大大"}

	axelSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Eustakio"}

	axelTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "大大"}
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
		Animal:   animal.Elephant,
		Birthday: axelBirthday,
		Code:     axelCode,
		Key:      character.Axel,
		Gender:   gender.Male,
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
		language.AmericanEnglish:      axelAmericanEnglishPhrase,
		language.CanadianFrench:       axelCanadianFrenchPhrase,
		language.Dutch:                axelDutchPhrase,
		language.French:               axelFrenchPhrase,
		language.German:               axelGermanPhrase,
		language.Italian:              axelItalianPhrase,
		language.Japanese:             axelJapanesePhrase,
		language.Korean:               axelKoreanPhrase,
		language.LatinAmericanSpanish: axelLatinAmericanSpanishPhrase,
		language.Russian:              axelRussianPhrase,
		language.SimplifiedChinese:    axelSimplifiedChinesePhrase,
		language.Spanish:              axelSpanishPhrase,
		language.TraditionalChinese:   axelTraditionalChinesePhrase}
)

var (
	Axel = nook.Villager{
		Character:   axelCharacter,
		Personality: personality.Jock,
		Phrase:      axelPhrase}
)
