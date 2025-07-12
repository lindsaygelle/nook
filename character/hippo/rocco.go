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
	// roccoBirthday represents rocco birthday.
	roccoBirthday = nook.Birthday{
		Day:   18,
		Month: time.August}
)

var (
	// roccoCode represents rocco code.
	roccoCode = nook.Code{
		Value: "hip00"}
)

var (
	// roccoAmericanEnglishName represents rocco american english name.
	roccoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rocco"}

	// roccoCanadianFrenchName represents rocco canadian french name.
	roccoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bebel"}

	// roccoDutchName represents rocco dutch name.
	roccoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rocco"}

	// roccoFrenchName represents rocco french name.
	roccoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bebel"}

	// roccoGermanName represents rocco german name.
	roccoGermanName = nook.Name{
		Language: language.German,
		Value:    "Richi"}

	// roccoItalianName represents rocco italian name.
	roccoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rocco"}

	// roccoJapaneseName represents rocco japanese name.
	roccoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゴンザレス"}

	// roccoKoreanName represents rocco korean name.
	roccoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "곤잘레스"}

	// roccoLatinAmericanSpanishName represents rocco latin american spanish name.
	roccoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Roco"}

	// roccoRussianName represents rocco russian name.
	roccoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рокко"}

	// roccoSimplifiedChineseName represents rocco simplified chinese name.
	roccoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "孔在来"}

	// roccoSpanishName represents rocco spanish name.
	roccoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Roco"}

	// roccoTraditionalChineseName represents rocco traditional chinese name.
	roccoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "孔在來"}
)

var (
	// roccoName represents rocco name.
	roccoName = nook.Languages{
		language.AmericanEnglish:      roccoAmericanEnglishName,
		language.CanadianFrench:       roccoCanadianFrenchName,
		language.Dutch:                roccoDutchName,
		language.French:               roccoFrenchName,
		language.German:               roccoGermanName,
		language.Italian:              roccoItalianName,
		language.Japanese:             roccoJapaneseName,
		language.Korean:               roccoKoreanName,
		language.LatinAmericanSpanish: roccoLatinAmericanSpanishName,
		language.Russian:              roccoRussianName,
		language.SimplifiedChinese:    roccoSimplifiedChineseName,
		language.Spanish:              roccoSpanishName,
		language.TraditionalChinese:   roccoTraditionalChineseName}
)

var (
	// roccoCharacter represents rocco character.
	roccoCharacter = nook.Character{
		Animal:   animal.Hippo,
		Birthday: roccoBirthday,
		Code:     roccoCode,
		Key:      character.Rocco,
		Gender:   gender.Male,
		Name:     roccoName,
		Special:  false}
)

var (
	// roccoAmericanEnglishPhrase represents rocco american english phrase.
	roccoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hippie"}

	// roccoCanadianFrenchPhrase represents rocco canadian french phrase.
	roccoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "babi"}

	// roccoDutchPhrase represents rocco dutch phrase.
	roccoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hippo"}

	// roccoFrenchPhrase represents rocco french phrase.
	roccoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "babi"}

	// roccoGermanPhrase represents rocco german phrase.
	roccoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "pfffft"}

	// roccoItalianPhrase represents rocco italian phrase.
	roccoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ipper"}

	// roccoJapanesePhrase represents rocco japanese phrase.
	roccoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だぎゃー"}

	// roccoKoreanPhrase represents rocco korean phrase.
	roccoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그러마"}

	// roccoLatinAmericanSpanishPhrase represents rocco latin american spanish phrase.
	roccoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "hiponó"}

	// roccoRussianPhrase represents rocco russian phrase.
	roccoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "гиппи"}

	// roccoSimplifiedChinesePhrase represents rocco simplified chinese phrase.
	roccoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "觉得是啦"}

	// roccoSpanishPhrase represents rocco spanish phrase.
	roccoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "batracio"}

	// roccoTraditionalChinesePhrase represents rocco traditional chinese phrase.
	roccoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "覺得是啦"}
)

var (
	// roccoPhrase represents rocco phrase.
	roccoPhrase = nook.Languages{
		language.AmericanEnglish:      roccoAmericanEnglishPhrase,
		language.CanadianFrench:       roccoCanadianFrenchPhrase,
		language.Dutch:                roccoDutchPhrase,
		language.French:               roccoFrenchPhrase,
		language.German:               roccoGermanPhrase,
		language.Italian:              roccoItalianPhrase,
		language.Japanese:             roccoJapanesePhrase,
		language.Korean:               roccoKoreanPhrase,
		language.LatinAmericanSpanish: roccoLatinAmericanSpanishPhrase,
		language.Russian:              roccoRussianPhrase,
		language.SimplifiedChinese:    roccoSimplifiedChinesePhrase,
		language.Spanish:              roccoSpanishPhrase,
		language.TraditionalChinese:   roccoTraditionalChinesePhrase}
)

var (
	// Rocco represents rocco.
	Rocco = nook.Villager{
		Character:   roccoCharacter,
		Personality: personality.Cranky,
		Phrase:      roccoPhrase}
)
