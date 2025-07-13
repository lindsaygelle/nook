package eagle

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
	// celiaBirthday represents celia birthday.
	celiaBirthday = nook.Birthday{
		Day:   25,
		Month: time.March}
)

var (
	// celiaCode represents celia code.
	celiaCode = nook.Code{
		Value: "pbr09"}
)

var (
	// celiaAmericanEnglishName represents celia american english name.
	celiaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Celia"}

	// celiaCanadianFrenchName represents celia canadian french name.
	celiaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Garance"}

	// celiaDutchName represents celia dutch name.
	celiaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Celia"}

	// celiaFrenchName represents celia french name.
	celiaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Garance"}

	// celiaGermanName represents celia german name.
	celiaGermanName = nook.Name{
		Language: language.German,
		Value:    "Lora"}

	// celiaItalianName represents celia italian name.
	celiaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Celia"}

	// celiaJapaneseName represents celia japanese name.
	celiaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ティファニー"}

	// celiaKoreanName represents celia korean name.
	celiaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "티파니"}

	// celiaLatinAmericanSpanishName represents celia latin american spanish name.
	celiaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jazmín"}

	// celiaRussianName represents celia russian name.
	celiaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Селия"}

	// celiaSimplifiedChineseName represents celia simplified chinese name.
	celiaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蒂芬妮"}

	// celiaSpanishName represents celia spanish name.
	celiaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jazmín"}

	// celiaTraditionalChineseName represents celia traditional chinese name.
	celiaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蒂芬妮"}
)

var (
	// celiaName represents celia name.
	celiaName = nook.Languages{
		language.AmericanEnglish:      celiaAmericanEnglishName,
		language.CanadianFrench:       celiaCanadianFrenchName,
		language.Dutch:                celiaDutchName,
		language.French:               celiaFrenchName,
		language.German:               celiaGermanName,
		language.Italian:              celiaItalianName,
		language.Japanese:             celiaJapaneseName,
		language.Korean:               celiaKoreanName,
		language.LatinAmericanSpanish: celiaLatinAmericanSpanishName,
		language.Russian:              celiaRussianName,
		language.SimplifiedChinese:    celiaSimplifiedChineseName,
		language.Spanish:              celiaSpanishName,
		language.TraditionalChinese:   celiaTraditionalChineseName}
)

var (
	// celiaCharacter represents celia character.
	celiaCharacter = nook.Character{
		Animal:   animal.Eagle,
		Birthday: celiaBirthday,
		Code:     celiaCode,
		Key:      character.Celia,
		Gender:   gender.Female,
		Name:     celiaName,
		Special:  false}
)

var (
	// celiaAmericanEnglishPhrase represents celia american english phrase.
	celiaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "feathers"}

	// celiaCanadianFrenchPhrase represents celia canadian french phrase.
	celiaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bebec"}

	// celiaDutchPhrase represents celia dutch phrase.
	celiaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pluimpje"}

	// celiaFrenchPhrase represents celia french phrase.
	celiaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bebec"}

	// celiaGermanPhrase represents celia german phrase.
	celiaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kraaack"}

	// celiaItalianPhrase represents celia italian phrase.
	celiaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ah bon"}

	// celiaJapanesePhrase represents celia japanese phrase.
	celiaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "そうね"}

	// celiaKoreanPhrase represents celia korean phrase.
	celiaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그래맞아"}

	// celiaLatinAmericanSpanishPhrase represents celia latin american spanish phrase.
	celiaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "plumí"}

	// celiaRussianPhrase represents celia russian phrase.
	celiaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "перышко"}

	// celiaSimplifiedChinesePhrase represents celia simplified chinese phrase.
	celiaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是呢"}

	// celiaSpanishPhrase represents celia spanish phrase.
	celiaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "plumis"}

	// celiaTraditionalChinesePhrase represents celia traditional chinese phrase.
	celiaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是呢"}
)

var (
	// celiaPhrase represents celia phrase.
	celiaPhrase = nook.Languages{
		language.AmericanEnglish:      celiaAmericanEnglishPhrase,
		language.CanadianFrench:       celiaCanadianFrenchPhrase,
		language.Dutch:                celiaDutchPhrase,
		language.French:               celiaFrenchPhrase,
		language.German:               celiaGermanPhrase,
		language.Italian:              celiaItalianPhrase,
		language.Japanese:             celiaJapanesePhrase,
		language.Korean:               celiaKoreanPhrase,
		language.LatinAmericanSpanish: celiaLatinAmericanSpanishPhrase,
		language.Russian:              celiaRussianPhrase,
		language.SimplifiedChinese:    celiaSimplifiedChinesePhrase,
		language.Spanish:              celiaSpanishPhrase,
		language.TraditionalChinese:   celiaTraditionalChinesePhrase}
)

var (
	// Celia represents celia.
	Celia = nook.Villager{
		Character:   celiaCharacter,
		Personality: personality.Normal,
		Phrase:      celiaPhrase}
)
