package penguin

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
	// floBirthday represents flo birthday.
	floBirthday = nook.Birthday{
		Day:   2,
		Month: time.September}
)

var (
	// floCode represents flo code.
	floCode = nook.Code{
		Value: "pgn13"}
)

var (
	// floAmericanEnglishName represents flo american english name.
	floAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Flo"}

	// floCanadianFrenchName represents flo canadian french name.
	floCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Nora"}

	// floDutchName represents flo dutch name.
	floDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Flo"}

	// floFrenchName represents flo french name.
	floFrenchName = nook.Name{
		Language: language.French,
		Value:    "Nora"}

	// floGermanName represents flo german name.
	floGermanName = nook.Name{
		Language: language.German,
		Value:    "Susanne"}

	// floItalianName represents flo italian name.
	floItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Nives"}

	// floJapaneseName represents flo japanese name.
	floJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "レイラ"}

	// floKoreanName represents flo korean name.
	floKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "레이라"}

	// floLatinAmericanSpanishName represents flo latin american spanish name.
	floLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nieves"}

	// floRussianName represents flo russian name.
	floRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фло"}

	// floSimplifiedChineseName represents flo simplified chinese name.
	floSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蕾拉"}

	// floSpanishName represents flo spanish name.
	floSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Nieves"}

	// floTraditionalChineseName represents flo traditional chinese name.
	floTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蕾拉"}
)

var (
	// floName represents flo name.
	floName = nook.Languages{
		language.AmericanEnglish:      floAmericanEnglishName,
		language.CanadianFrench:       floCanadianFrenchName,
		language.Dutch:                floDutchName,
		language.French:               floFrenchName,
		language.German:               floGermanName,
		language.Italian:              floItalianName,
		language.Japanese:             floJapaneseName,
		language.Korean:               floKoreanName,
		language.LatinAmericanSpanish: floLatinAmericanSpanishName,
		language.Russian:              floRussianName,
		language.SimplifiedChinese:    floSimplifiedChineseName,
		language.Spanish:              floSpanishName,
		language.TraditionalChinese:   floTraditionalChineseName}
)

var (
	// floCharacter represents flo character.
	floCharacter = nook.Character{
		Animal:   animal.Penguin,
		Birthday: floBirthday,
		Code:     floCode,
		Key:      character.Flo,
		Gender:   gender.Female,
		Name:     floName,
		Special:  false}
)

var (
	// floAmericanEnglishPhrase represents flo american english phrase.
	floAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cha"}

	// floCanadianFrenchPhrase represents flo canadian french phrase.
	floCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "frigogo"}

	// floDutchPhrase represents flo dutch phrase.
	floDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "duh"}

	// floFrenchPhrase represents flo french phrase.
	floFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "golri"}

	// floGermanPhrase represents flo german phrase.
	floGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "flatsch"}

	// floItalianPhrase represents flo italian phrase.
	floItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "iglù"}

	// floJapanesePhrase represents flo japanese phrase.
	floJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "じゃんよ"}

	// floKoreanPhrase represents flo korean phrase.
	floKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "맞잖아"}

	// floLatinAmericanSpanishPhrase represents flo latin american spanish phrase.
	floLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fresqui"}

	// floRussianPhrase represents flo russian phrase.
	floRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "чхи"}

	// floSimplifiedChinesePhrase represents flo simplified chinese phrase.
	floSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "这样唷"}

	// floSpanishPhrase represents flo spanish phrase.
	floSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fresqui"}

	// floTraditionalChinesePhrase represents flo traditional chinese phrase.
	floTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "這樣唷"}
)

var (
	// floPhrase represents flo phrase.
	floPhrase = nook.Languages{
		language.AmericanEnglish:      floAmericanEnglishPhrase,
		language.CanadianFrench:       floCanadianFrenchPhrase,
		language.Dutch:                floDutchPhrase,
		language.French:               floFrenchPhrase,
		language.German:               floGermanPhrase,
		language.Italian:              floItalianPhrase,
		language.Japanese:             floJapanesePhrase,
		language.Korean:               floKoreanPhrase,
		language.LatinAmericanSpanish: floLatinAmericanSpanishPhrase,
		language.Russian:              floRussianPhrase,
		language.SimplifiedChinese:    floSimplifiedChinesePhrase,
		language.Spanish:              floSpanishPhrase,
		language.TraditionalChinese:   floTraditionalChinesePhrase}
)

var (
	// Flo represents flo.
	Flo = nook.Villager{
		Character:   floCharacter,
		Personality: personality.BigSister,
		Phrase:      floPhrase}
)
