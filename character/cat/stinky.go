package cat

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
		Value:    "Tupux"}

	stinkyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Stinky"}

	stinkyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tupux"}

	stinkyGermanName = nook.Name{
		Language: language.German,
		Value:    "Stefan"}

	stinkyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Puzzolo"}

	stinkyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アセクサ"}

	stinkyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "땀띠"}

	stinkyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Micifuz"}

	stinkyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Стинки"}

	stinkySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "男子汗"}

	stinkySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Micifuz"}

	stinkyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "男子汗"}
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
		Animal:   animal.Cat,
		Birthday: stinkyBirthday,
		Code:     stinkyCode,
		Key:      character.Stinky,
		Gender:   gender.Male,
		Name:     stinkyName,
		Special:  false}
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
		language.AmericanEnglish:      stinkyAmericanEnglishPhrase,
		language.CanadianFrench:       stinkyCanadianFrenchPhrase,
		language.Dutch:                stinkyDutchPhrase,
		language.French:               stinkyFrenchPhrase,
		language.German:               stinkyGermanPhrase,
		language.Italian:              stinkyItalianPhrase,
		language.Japanese:             stinkyJapanesePhrase,
		language.Korean:               stinkyKoreanPhrase,
		language.LatinAmericanSpanish: stinkyLatinAmericanSpanishPhrase,
		language.Russian:              stinkyRussianPhrase,
		language.SimplifiedChinese:    stinkySimplifiedChinesePhrase,
		language.Spanish:              stinkySpanishPhrase,
		language.TraditionalChinese:   stinkyTraditionalChinesePhrase}
)

var (
	Stinky = nook.Villager{
		Character:   stinkyCharacter,
		Personality: personality.Jock,
		Phrase:      stinkyPhrase}
)
