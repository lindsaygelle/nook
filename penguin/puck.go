package penguin

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	puckBirthday = nook.Birthday{
		Day:   21,
		Month: time.February}
)

var (
	puckCode = nook.Code{
		Value: "pgn06"}
)

var (
	puckAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Puck"}

	puckCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Hervébrrrrrrrrr"}

	puckDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Puckbrrrrrrrrr"}

	puckFrenchName = nook.Name{
		Language: language.French,
		Value:    "Hervébrrrrrrrrr"}

	puckGermanName = nook.Name{
		Language: language.German,
		Value:    "Puckbrrrrrrrrr"}

	puckItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Glacidobrrrrrrrrr"}

	puckJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ホッケーさぶー"}

	puckKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "하키엣취"}

	puckLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Fredobrrr-uah"}

	puckRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пакбрр"}

	puckSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哈奇候补"}

	puckSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Fredosardinilla"}

	puckTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哈奇候補"}
)

var (
	puckName = nook.Languages{
		language.AmericanEnglish:      puckAmericanEnglishName,
		language.CanadianFrench:       puckCanadianFrenchName,
		language.Dutch:                puckDutchName,
		language.French:               puckFrenchName,
		language.German:               puckGermanName,
		language.Italian:              puckItalianName,
		language.Japanese:             puckJapaneseName,
		language.Korean:               puckKoreanName,
		language.LatinAmericanSpanish: puckLatinAmericanSpanishName,
		language.Russian:              puckRussianName,
		language.SimplifiedChinese:    puckSimplifiedChineseName,
		language.Spanish:              puckSpanishName,
		language.TraditionalChinese:   puckTraditionalChineseName}
)

var (
	puckCharacter = nook.Character{
		Animal:   Penguin,
		Birthday: puckBirthday,
		Code:     puckCode,
		Gender:   nook.Male,
		Name:     puckName}
)

var (
	puckAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "brrrrrrrrr"}

	puckCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "brrrrrrrrr"}

	puckDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "brrrrrrrrr"}

	puckFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "brrrrrrrrr"}

	puckGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brrrrrrrrr"}

	puckItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "brrrrrrrrr"}

	puckJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "さぶー"}

	puckKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "엣취"}

	puckLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "brrr-uah"}

	puckRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "брр"}

	puckSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "候补"}

	puckSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "sardinilla"}

	puckTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "候補"}
)

var (
	puckPhrase = nook.Languages{
		language.AmericanEnglish:      puckAmericanEnglishName,
		language.CanadianFrench:       puckCanadianFrenchName,
		language.Dutch:                puckDutchName,
		language.French:               puckFrenchName,
		language.German:               puckGermanName,
		language.Italian:              puckItalianName,
		language.Japanese:             puckJapaneseName,
		language.Korean:               puckKoreanName,
		language.LatinAmericanSpanish: puckLatinAmericanSpanishName,
		language.Russian:              puckRussianName,
		language.SimplifiedChinese:    puckSimplifiedChineseName,
		language.Spanish:              puckSpanishName,
		language.TraditionalChinese:   puckTraditionalChineseName}
)

var (
	Puck = nook.Villager{
		Character:   puckCharacter,
		Personality: nook.Lazy,
		Phrase:      puckPhrase}
)
