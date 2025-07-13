package hippo

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
	// hippeuxBirthday represents hippeux birthday.
	hippeuxBirthday = nook.Birthday{
		Day:   15,
		Month: time.October}
)

var (
	// hippeuxCode represents hippeux code.
	hippeuxCode = nook.Code{
		Value: "hip09"}
)

var (
	// hippeuxAmericanEnglishName represents hippeux american english name.
	hippeuxAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hippeux"}

	// hippeuxCanadianFrenchName represents hippeux canadian french name.
	hippeuxCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Paulito"}

	// hippeuxDutchName represents hippeux dutch name.
	hippeuxDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hippeux"}

	// hippeuxFrenchName represents hippeux french name.
	hippeuxFrenchName = nook.Name{
		Language: language.French,
		Value:    "Paulito"}

	// hippeuxGermanName represents hippeux german name.
	hippeuxGermanName = nook.Name{
		Language: language.German,
		Value:    "Herbert"}

	// hippeuxItalianName represents hippeux italian name.
	hippeuxItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Poppy"}

	// hippeuxJapaneseName represents hippeux japanese name.
	hippeuxJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ディビッド"}

	// hippeuxKoreanName represents hippeux korean name.
	hippeuxKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "데이빗"}

	// hippeuxLatinAmericanSpanishName represents hippeux latin american spanish name.
	hippeuxLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Hipóleo"}

	// hippeuxRussianName represents hippeux russian name.
	hippeuxRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гиппо"}

	// hippeuxSimplifiedChineseName represents hippeux simplified chinese name.
	hippeuxSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "戴维"}

	// hippeuxSpanishName represents hippeux spanish name.
	hippeuxSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Hipóleo"}

	// hippeuxTraditionalChineseName represents hippeux traditional chinese name.
	hippeuxTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "戴維"}
)

var (
	// hippeuxName represents hippeux name.
	hippeuxName = nook.Languages{
		language.AmericanEnglish:      hippeuxAmericanEnglishName,
		language.CanadianFrench:       hippeuxCanadianFrenchName,
		language.Dutch:                hippeuxDutchName,
		language.French:               hippeuxFrenchName,
		language.German:               hippeuxGermanName,
		language.Italian:              hippeuxItalianName,
		language.Japanese:             hippeuxJapaneseName,
		language.Korean:               hippeuxKoreanName,
		language.LatinAmericanSpanish: hippeuxLatinAmericanSpanishName,
		language.Russian:              hippeuxRussianName,
		language.SimplifiedChinese:    hippeuxSimplifiedChineseName,
		language.Spanish:              hippeuxSpanishName,
		language.TraditionalChinese:   hippeuxTraditionalChineseName}
)

var (
	// hippeuxCharacter represents hippeux character.
	hippeuxCharacter = nook.Character{
		Animal:   animal.Hippo,
		Birthday: hippeuxBirthday,
		Code:     hippeuxCode,
		Key:      character.Hippeux,
		Gender:   gender.Male,
		Name:     hippeuxName,
		Special:  false}
)

var (
	// hippeuxAmericanEnglishPhrase represents hippeux american english phrase.
	hippeuxAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "natch"}

	// hippeuxCanadianFrenchPhrase represents hippeux canadian french phrase.
	hippeuxCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tam-tam"}

	// hippeuxDutchPhrase represents hippeux dutch phrase.
	hippeuxDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "tuurlijk"}

	// hippeuxFrenchPhrase represents hippeux french phrase.
	hippeuxFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tam-tam"}

	// hippeuxGermanPhrase represents hippeux german phrase.
	hippeuxGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hipphipp"}

	// hippeuxItalianPhrase represents hippeux italian phrase.
	hippeuxItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "vale"}

	// hippeuxJapanesePhrase represents hippeux japanese phrase.
	hippeuxJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "イェス"}

	// hippeuxKoreanPhrase represents hippeux korean phrase.
	hippeuxKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "예스"}

	// hippeuxLatinAmericanSpanishPhrase represents hippeux latin american spanish phrase.
	hippeuxLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chofchof"}

	// hippeuxRussianPhrase represents hippeux russian phrase.
	hippeuxRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "естессно"}

	// hippeuxSimplifiedChinesePhrase represents hippeux simplified chinese phrase.
	hippeuxSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Yes"}

	// hippeuxSpanishPhrase represents hippeux spanish phrase.
	hippeuxSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chofchof"}

	// hippeuxTraditionalChinesePhrase represents hippeux traditional chinese phrase.
	hippeuxTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Yes"}
)

var (
	// hippeuxPhrase represents hippeux phrase.
	hippeuxPhrase = nook.Languages{
		language.AmericanEnglish:      hippeuxAmericanEnglishPhrase,
		language.CanadianFrench:       hippeuxCanadianFrenchPhrase,
		language.Dutch:                hippeuxDutchPhrase,
		language.French:               hippeuxFrenchPhrase,
		language.German:               hippeuxGermanPhrase,
		language.Italian:              hippeuxItalianPhrase,
		language.Japanese:             hippeuxJapanesePhrase,
		language.Korean:               hippeuxKoreanPhrase,
		language.LatinAmericanSpanish: hippeuxLatinAmericanSpanishPhrase,
		language.Russian:              hippeuxRussianPhrase,
		language.SimplifiedChinese:    hippeuxSimplifiedChinesePhrase,
		language.Spanish:              hippeuxSpanishPhrase,
		language.TraditionalChinese:   hippeuxTraditionalChinesePhrase}
)

var (
	// Hippeux represents hippeux.
	Hippeux = nook.Villager{
		Character:   hippeuxCharacter,
		Personality: personality.Smug,
		Phrase:      hippeuxPhrase}
)
