package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	stinkyBirthday = nook.Birthday{
		Day:   17,
		Month: time.August}
)

var (
	stinkyCode = nook.Code{
		Value: "cat13"}
)

var (
	stinkyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Stinky"}

	stinkyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tupuxmiaoum"}

	stinkyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "StinkyBAH"}

	stinkyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tupuxmiaoum"}

	stinkyGermanName = nook.Name{
		Language: language.German,
		Value:    "Stefanfauchch"}

	stinkyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "PuzzoloGAAHHH"}

	stinkyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アセクサダァーッ"}

	stinkyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "땀띠아뵤"}

	stinkyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Micifuzmarramiau"}

	stinkyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Стинкифу-у-у"}

	stinkySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "男子汗打啊"}

	stinkySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Micifuzmarramiau"}

	stinkyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "男子汗打啊"}
)

var (
	stinkyName = nook.Languages{
		language.AmericanEnglish:      stinkyAmericanEnglishName,
		language.CanadianFrench:       stinkyCanadianFrenchName,
		language.Dutch:                stinkyDutchName,
		language.French:               stinkyFrenchName,
		language.German:               stinkyGermanName,
		language.Italian:              stinkyItalianName,
		language.Japanese:             stinkyJapaneseName,
		language.Korean:               stinkyKoreanName,
		language.LatinAmericanSpanish: stinkyLatinAmericanSpanishName,
		language.Russian:              stinkyRussianName,
		language.SimplifiedChinese:    stinkySimplifiedChineseName,
		language.Spanish:              stinkySpanishName,
		language.TraditionalChinese:   stinkyTraditionalChineseName}
)

var (
	stinkyCharacter = nook.Character{
		Animal:   Cat,
		Birthday: stinkyBirthday,
		Code:     stinkyCode,
		Gender:   nook.Male,
		Name:     stinkyName}
)

var (
	stinkyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "GAAHHH"}

	stinkyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "miaoum"}

	stinkyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "BAH"}

	stinkyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "miaoum"}

	stinkyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "fauchch"}

	stinkyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "GAAHHH"}

	stinkyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ダァーッ"}

	stinkyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아뵤"}

	stinkyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "marramiau"}

	stinkyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "фу-у-у"}

	stinkySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "打啊"}

	stinkySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "marramiau"}

	stinkyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "打啊"}
)

var (
	stinkyPhrase = nook.Languages{
		language.AmericanEnglish:      stinkyAmericanEnglishName,
		language.CanadianFrench:       stinkyCanadianFrenchName,
		language.Dutch:                stinkyDutchName,
		language.French:               stinkyFrenchName,
		language.German:               stinkyGermanName,
		language.Italian:              stinkyItalianName,
		language.Japanese:             stinkyJapaneseName,
		language.Korean:               stinkyKoreanName,
		language.LatinAmericanSpanish: stinkyLatinAmericanSpanishName,
		language.Russian:              stinkyRussianName,
		language.SimplifiedChinese:    stinkySimplifiedChineseName,
		language.Spanish:              stinkySpanishName,
		language.TraditionalChinese:   stinkyTraditionalChineseName}
)

var (
	Stinky = nook.Villager{
		Character:   stinkyCharacter,
		Personality: nook.Jock,
		Phrase:      stinkyPhrase}
)
