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
	// stinkyBirthday represents stinky birthday.
	stinkyBirthday = nook.Birthday{
		Day:   17,
		Month: time.August}
)

var (
	// stinkyCode represents stinky code.
	stinkyCode = nook.Code{
		Value: "cat13"}
)

var (
	// stinkyAmericanEnglishName represents stinky american english name.
	stinkyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Stinky"}

	// stinkyCanadianFrenchName represents stinky canadian french name.
	stinkyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tupux"}

	// stinkyDutchName represents stinky dutch name.
	stinkyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Stinky"}

	// stinkyFrenchName represents stinky french name.
	stinkyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tupux"}

	// stinkyGermanName represents stinky german name.
	stinkyGermanName = nook.Name{
		Language: language.German,
		Value:    "Stefan"}

	// stinkyItalianName represents stinky italian name.
	stinkyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Puzzolo"}

	// stinkyJapaneseName represents stinky japanese name.
	stinkyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アセクサ"}

	// stinkyKoreanName represents stinky korean name.
	stinkyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "땀띠"}

	// stinkyLatinAmericanSpanishName represents stinky latin american spanish name.
	stinkyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Micifuz"}

	// stinkyRussianName represents stinky russian name.
	stinkyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Стинки"}

	// stinkySimplifiedChineseName represents stinky simplified chinese name.
	stinkySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "男子汗"}

	// stinkySpanishName represents stinky spanish name.
	stinkySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Micifuz"}

	// stinkyTraditionalChineseName represents stinky traditional chinese name.
	stinkyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "男子汗"}
)

var (
	// stinkyName represents stinky name.
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
	// stinkyCharacter represents stinky character.
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
	// stinkyAmericanEnglishPhrase represents stinky american english phrase.
	stinkyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "GAAHHH"}

	// stinkyCanadianFrenchPhrase represents stinky canadian french phrase.
	stinkyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "miaoum"}

	// stinkyDutchPhrase represents stinky dutch phrase.
	stinkyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "BAH"}

	// stinkyFrenchPhrase represents stinky french phrase.
	stinkyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "miaoum"}

	// stinkyGermanPhrase represents stinky german phrase.
	stinkyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "fauchch"}

	// stinkyItalianPhrase represents stinky italian phrase.
	stinkyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "GAAHHH"}

	// stinkyJapanesePhrase represents stinky japanese phrase.
	stinkyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ダァーッ"}

	// stinkyKoreanPhrase represents stinky korean phrase.
	stinkyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아뵤"}

	// stinkyLatinAmericanSpanishPhrase represents stinky latin american spanish phrase.
	stinkyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "marramiau"}

	// stinkyRussianPhrase represents stinky russian phrase.
	stinkyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "фу-у-у"}

	// stinkySimplifiedChinesePhrase represents stinky simplified chinese phrase.
	stinkySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "打啊"}

	// stinkySpanishPhrase represents stinky spanish phrase.
	stinkySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "marramiau"}

	// stinkyTraditionalChinesePhrase represents stinky traditional chinese phrase.
	stinkyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "打啊"}
)

var (
	// stinkyPhrase represents stinky phrase.
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
	// Stinky represents stinky.
	Stinky = nook.Villager{
		Character:   stinkyCharacter,
		Personality: personality.Jock,
		Phrase:      stinkyPhrase}
)
