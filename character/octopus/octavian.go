package octopus

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
	// octavianBirthday represents octavian birthday.
	octavianBirthday = nook.Birthday{
		Day:   20,
		Month: time.September}
)

var (
	// octavianCode represents octavian code.
	octavianCode = nook.Code{
		Value: "ocp00"}
)

var (
	// octavianAmericanEnglishName represents octavian american english name.
	octavianAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Octavian"}

	// octavianCanadianFrenchName represents octavian canadian french name.
	octavianCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Octave"}

	// octavianDutchName represents octavian dutch name.
	octavianDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Octavian"}

	// octavianFrenchName represents octavian french name.
	octavianFrenchName = nook.Name{
		Language: language.French,
		Value:    "Octave"}

	// octavianGermanName represents octavian german name.
	octavianGermanName = nook.Name{
		Language: language.German,
		Value:    "Ottfried"}

	// octavianItalianName represents octavian italian name.
	octavianItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ottavio"}

	// octavianJapaneseName represents octavian japanese name.
	octavianJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "おくたろう"}

	// octavianKoreanName represents octavian korean name.
	octavianKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "문복"}

	// octavianLatinAmericanSpanishName represents octavian latin american spanish name.
	octavianLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Octavio"}

	// octavianRussianName represents octavian russian name.
	octavianRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Октавиан"}

	// octavianSimplifiedChineseName represents octavian simplified chinese name.
	octavianSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "张瑜"}

	// octavianSpanishName represents octavian spanish name.
	octavianSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Octavio"}

	// octavianTraditionalChineseName represents octavian traditional chinese name.
	octavianTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "張瑜"}
)

var (
	// octavianName represents octavian name.
	octavianName = nook.Languages{
		language.AmericanEnglish:      octavianAmericanEnglishName,
		language.CanadianFrench:       octavianCanadianFrenchName,
		language.Dutch:                octavianDutchName,
		language.French:               octavianFrenchName,
		language.German:               octavianGermanName,
		language.Italian:              octavianItalianName,
		language.Japanese:             octavianJapaneseName,
		language.Korean:               octavianKoreanName,
		language.LatinAmericanSpanish: octavianLatinAmericanSpanishName,
		language.Russian:              octavianRussianName,
		language.SimplifiedChinese:    octavianSimplifiedChineseName,
		language.Spanish:              octavianSpanishName,
		language.TraditionalChinese:   octavianTraditionalChineseName}
)

var (
	// octavianCharacter represents octavian character.
	octavianCharacter = nook.Character{
		Animal:   animal.Octopus,
		Birthday: octavianBirthday,
		Code:     octavianCode,
		Key:      character.Octavian,
		Gender:   gender.Male,
		Name:     octavianName,
		Special:  false}
)

var (
	// octavianAmericanEnglishPhrase represents octavian american english phrase.
	octavianAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sucker"}

	// octavianCanadianFrenchPhrase represents octavian canadian french phrase.
	octavianCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chenapan"}

	// octavianDutchPhrase represents octavian dutch phrase.
	octavianDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "zuignap"}

	// octavianFrenchPhrase represents octavian french phrase.
	octavianFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "chenapan"}

	// octavianGermanPhrase represents octavian german phrase.
	octavianGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schling"}

	// octavianItalianPhrase represents octavian italian phrase.
	octavianItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "perdindi"}

	// octavianJapanesePhrase represents octavian japanese phrase.
	octavianJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "タコ"}

	// octavianKoreanPhrase represents octavian korean phrase.
	octavianKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쭉쭉"}

	// octavianLatinAmericanSpanishPhrase represents octavian latin american spanish phrase.
	octavianLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "glop-glop"}

	// octavianRussianPhrase represents octavian russian phrase.
	octavianRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "все пропало"}

	// octavianSimplifiedChinesePhrase represents octavian simplified chinese phrase.
	octavianSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "章鱼"}

	// octavianSpanishPhrase represents octavian spanish phrase.
	octavianSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chupatinta"}

	// octavianTraditionalChinesePhrase represents octavian traditional chinese phrase.
	octavianTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "章魚"}
)

var (
	// octavianPhrase represents octavian phrase.
	octavianPhrase = nook.Languages{
		language.AmericanEnglish:      octavianAmericanEnglishPhrase,
		language.CanadianFrench:       octavianCanadianFrenchPhrase,
		language.Dutch:                octavianDutchPhrase,
		language.French:               octavianFrenchPhrase,
		language.German:               octavianGermanPhrase,
		language.Italian:              octavianItalianPhrase,
		language.Japanese:             octavianJapanesePhrase,
		language.Korean:               octavianKoreanPhrase,
		language.LatinAmericanSpanish: octavianLatinAmericanSpanishPhrase,
		language.Russian:              octavianRussianPhrase,
		language.SimplifiedChinese:    octavianSimplifiedChinesePhrase,
		language.Spanish:              octavianSpanishPhrase,
		language.TraditionalChinese:   octavianTraditionalChinesePhrase}
)

var (
	// Octavian represents octavian.
	Octavian = nook.Villager{
		Character:   octavianCharacter,
		Personality: personality.Cranky,
		Phrase:      octavianPhrase}
)
