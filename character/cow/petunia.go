package cow

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
	// petuniaBirthday represents petunia birthday.
	petuniaBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// petuniaCode represents petunia code.
	petuniaCode = nook.Code{
		Value: ""}
)

var (
	// petuniaAmericanEnglishName represents petunia american english name.
	petuniaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Petunia"}

	// petuniaCanadianFrenchName represents petunia canadian french name.
	petuniaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// petuniaDutchName represents petunia dutch name.
	petuniaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// petuniaFrenchName represents petunia french name.
	petuniaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pétunia"}

	// petuniaGermanName represents petunia german name.
	petuniaGermanName = nook.Name{
		Language: language.German,
		Value:    "Petunia"}

	// petuniaItalianName represents petunia italian name.
	petuniaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Petunia"}

	// petuniaJapaneseName represents petunia japanese name.
	petuniaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "しもふり"}

	// petuniaKoreanName represents petunia korean name.
	petuniaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// petuniaLatinAmericanSpanishName represents petunia latin american spanish name.
	petuniaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// petuniaRussianName represents petunia russian name.
	petuniaRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// petuniaSimplifiedChineseName represents petunia simplified chinese name.
	petuniaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "彤彤"}

	// petuniaSpanishName represents petunia spanish name.
	petuniaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Petunia"}

	// petuniaTraditionalChineseName represents petunia traditional chinese name.
	petuniaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// petuniaName represents petunia name.
	petuniaName = nook.Languages{
		language.AmericanEnglish:      petuniaAmericanEnglishName,
		language.CanadianFrench:       petuniaCanadianFrenchName,
		language.Dutch:                petuniaDutchName,
		language.French:               petuniaFrenchName,
		language.German:               petuniaGermanName,
		language.Italian:              petuniaItalianName,
		language.Japanese:             petuniaJapaneseName,
		language.Korean:               petuniaKoreanName,
		language.LatinAmericanSpanish: petuniaLatinAmericanSpanishName,
		language.Russian:              petuniaRussianName,
		language.SimplifiedChinese:    petuniaSimplifiedChineseName,
		language.Spanish:              petuniaSpanishName,
		language.TraditionalChinese:   petuniaTraditionalChineseName}
)

var (
	// petuniaCharacter represents petunia character.
	petuniaCharacter = nook.Character{
		Animal:   animal.Cow,
		Birthday: petuniaBirthday,
		Code:     petuniaCode,
		Key:      character.Petunia,
		Gender:   gender.Female,
		Name:     petuniaName,
		Special:  false}
)

var (
	// petuniaAmericanEnglishPhrase represents petunia american english phrase.
	petuniaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "moo la la"}

	// petuniaCanadianFrenchPhrase represents petunia canadian french phrase.
	petuniaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// petuniaDutchPhrase represents petunia dutch phrase.
	petuniaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// petuniaFrenchPhrase represents petunia french phrase.
	petuniaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mouh mouh"}

	// petuniaGermanPhrase represents petunia german phrase.
	petuniaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "muh-la-la"}

	// petuniaItalianPhrase represents petunia italian phrase.
	petuniaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "muu la la"}

	// petuniaJapanesePhrase represents petunia japanese phrase.
	petuniaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ザマス"}

	// petuniaKoreanPhrase represents petunia korean phrase.
	petuniaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// petuniaLatinAmericanSpanishPhrase represents petunia latin american spanish phrase.
	petuniaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// petuniaRussianPhrase represents petunia russian phrase.
	petuniaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// petuniaSimplifiedChinesePhrase represents petunia simplified chinese phrase.
	petuniaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哞呜"}

	// petuniaSpanishPhrase represents petunia spanish phrase.
	petuniaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "muu-la-la"}

	// petuniaTraditionalChinesePhrase represents petunia traditional chinese phrase.
	petuniaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// petuniaPhrase represents petunia phrase.
	petuniaPhrase = nook.Languages{
		language.AmericanEnglish:      petuniaAmericanEnglishPhrase,
		language.CanadianFrench:       petuniaCanadianFrenchPhrase,
		language.Dutch:                petuniaDutchPhrase,
		language.French:               petuniaFrenchPhrase,
		language.German:               petuniaGermanPhrase,
		language.Italian:              petuniaItalianPhrase,
		language.Japanese:             petuniaJapanesePhrase,
		language.Korean:               petuniaKoreanPhrase,
		language.LatinAmericanSpanish: petuniaLatinAmericanSpanishPhrase,
		language.Russian:              petuniaRussianPhrase,
		language.SimplifiedChinese:    petuniaSimplifiedChinesePhrase,
		language.Spanish:              petuniaSpanishPhrase,
		language.TraditionalChinese:   petuniaTraditionalChinesePhrase}
)

var (
	// Petunia represents petunia.
	Petunia = nook.Villager{
		Character:   petuniaCharacter,
		Personality: personality.Snooty,
		Phrase:      petuniaPhrase}
)
