package pig

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	curlyBirthday = nook.Birthday{
		Day:   26,
		Month: time.July}
)

var (
	curlyCode = nook.Code{
		Value: "pig00"}
)

var (
	curlyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Curly"}

	curlyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tirboucoink"}

	curlyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Curlyhupsaknor"}

	curlyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tirboucoink"}

	curlyGermanName = nook.Name{
		Language: language.German,
		Value:    "Oinkoinkoink"}

	curlyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Riccioloniooink"}

	curlyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハムカツどもども"}

	curlyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "햄까스꿀꿀"}

	curlyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rufuetoñoink"}

	curlyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Керлихряк-бряк"}

	curlySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "胜利感谢"}

	curlySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rufuetoñoink"}

	curlyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "勝利感謝"}
)

var (
	curlyName = nook.Languages{
		language.AmericanEnglish:      curlyAmericanEnglishName,
		language.CanadianFrench:       curlyCanadianFrenchName,
		language.Dutch:                curlyDutchName,
		language.French:               curlyFrenchName,
		language.German:               curlyGermanName,
		language.Italian:              curlyItalianName,
		language.Japanese:             curlyJapaneseName,
		language.Korean:               curlyKoreanName,
		language.LatinAmericanSpanish: curlyLatinAmericanSpanishName,
		language.Russian:              curlyRussianName,
		language.SimplifiedChinese:    curlySimplifiedChineseName,
		language.Spanish:              curlySpanishName,
		language.TraditionalChinese:   curlyTraditionalChineseName}
)

var (
	curlyCharacter = nook.Character{
		Animal:   Pig,
		Birthday: curlyBirthday,
		Code:     curlyCode,
		Gender:   nook.Male,
		Name:     curlyName}
)

var (
	curlyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nyoink"}

	curlyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "coink"}

	curlyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hupsaknor"}

	curlyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "coink"}

	curlyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "oinkoink"}

	curlyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "niooink"}

	curlyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "どもども"}

	curlyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꿀꿀"}

	curlyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñoink"}

	curlyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хряк-бряк"}

	curlySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "感谢"}

	curlySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ñoink"}

	curlyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "感謝"}
)

var (
	curlyPhrase = nook.Languages{
		language.AmericanEnglish:      curlyAmericanEnglishName,
		language.CanadianFrench:       curlyCanadianFrenchName,
		language.Dutch:                curlyDutchName,
		language.French:               curlyFrenchName,
		language.German:               curlyGermanName,
		language.Italian:              curlyItalianName,
		language.Japanese:             curlyJapaneseName,
		language.Korean:               curlyKoreanName,
		language.LatinAmericanSpanish: curlyLatinAmericanSpanishName,
		language.Russian:              curlyRussianName,
		language.SimplifiedChinese:    curlySimplifiedChineseName,
		language.Spanish:              curlySpanishName,
		language.TraditionalChinese:   curlyTraditionalChineseName}
)

var (
	Curly = nook.Villager{
		Character:   curlyCharacter,
		Personality: nook.Jock,
		Phrase:      curlyPhrase}
)
