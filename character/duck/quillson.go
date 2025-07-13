package duck

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
	// quillsonBirthday represents quillson birthday.
	quillsonBirthday = nook.Birthday{
		Day:   22,
		Month: time.December}
)

var (
	// quillsonCode represents quillson code.
	quillsonCode = nook.Code{
		Value: "duk17"}
)

var (
	// quillsonAmericanEnglishName represents quillson american english name.
	quillsonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Quillson"}

	// quillsonCanadianFrenchName represents quillson canadian french name.
	quillsonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Narcisse"}

	// quillsonDutchName represents quillson dutch name.
	quillsonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Quillson"}

	// quillsonFrenchName represents quillson french name.
	quillsonFrenchName = nook.Name{
		Language: language.French,
		Value:    "Narcisse"}

	// quillsonGermanName represents quillson german name.
	quillsonGermanName = nook.Name{
		Language: language.German,
		Value:    "Quentin"}

	// quillsonItalianName represents quillson italian name.
	quillsonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Verdonio"}

	// quillsonJapaneseName represents quillson japanese name.
	quillsonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タックン"}

	// quillsonKoreanName represents quillson korean name.
	quillsonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "덕근"}

	// quillsonLatinAmericanSpanishName represents quillson latin american spanish name.
	quillsonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cuálter"}

	// quillsonRussianName represents quillson russian name.
	quillsonRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Квилсон"}

	// quillsonSimplifiedChineseName represents quillson simplified chinese name.
	quillsonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "何童"}

	// quillsonSpanishName represents quillson spanish name.
	quillsonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cuálter"}

	// quillsonTraditionalChineseName represents quillson traditional chinese name.
	quillsonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "何童"}
)

var (
	// quillsonName represents quillson name.
	quillsonName = nook.Languages{
		language.AmericanEnglish:      quillsonAmericanEnglishName,
		language.CanadianFrench:       quillsonCanadianFrenchName,
		language.Dutch:                quillsonDutchName,
		language.French:               quillsonFrenchName,
		language.German:               quillsonGermanName,
		language.Italian:              quillsonItalianName,
		language.Japanese:             quillsonJapaneseName,
		language.Korean:               quillsonKoreanName,
		language.LatinAmericanSpanish: quillsonLatinAmericanSpanishName,
		language.Russian:              quillsonRussianName,
		language.SimplifiedChinese:    quillsonSimplifiedChineseName,
		language.Spanish:              quillsonSpanishName,
		language.TraditionalChinese:   quillsonTraditionalChineseName}
)

var (
	// quillsonCharacter represents quillson character.
	quillsonCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: quillsonBirthday,
		Code:     quillsonCode,
		Key:      character.Quillson,
		Gender:   gender.Male,
		Name:     quillsonName,
		Special:  false}
)

var (
	// quillsonAmericanEnglishPhrase represents quillson american english phrase.
	quillsonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ridukulous"}

	// quillsonCanadianFrenchPhrase represents quillson canadian french phrase.
	quillsonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "arrêêêête"}

	// quillsonDutchPhrase represents quillson dutch phrase.
	quillsonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "smient"}

	// quillsonFrenchPhrase represents quillson french phrase.
	quillsonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "arrêêêête"}

	// quillsonGermanPhrase represents quillson german phrase.
	quillsonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "pluster"}

	// quillsonItalianPhrase represents quillson italian phrase.
	quillsonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cua cua"}

	// quillsonJapanesePhrase represents quillson japanese phrase.
	quillsonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "マジかよ"}

	// quillsonKoreanPhrase represents quillson korean phrase.
	quillsonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "리얼리"}

	// quillsonLatinAmericanSpanishPhrase represents quillson latin american spanish phrase.
	quillsonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cuacoac"}

	// quillsonRussianPhrase represents quillson russian phrase.
	quillsonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "шу-утка"}

	// quillsonSimplifiedChinesePhrase represents quillson simplified chinese phrase.
	quillsonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "说真的"}

	// quillsonSpanishPhrase represents quillson spanish phrase.
	quillsonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "verdoncho"}

	// quillsonTraditionalChinesePhrase represents quillson traditional chinese phrase.
	quillsonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "說真的"}
)

var (
	// quillsonPhrase represents quillson phrase.
	quillsonPhrase = nook.Languages{
		language.AmericanEnglish:      quillsonAmericanEnglishPhrase,
		language.CanadianFrench:       quillsonCanadianFrenchPhrase,
		language.Dutch:                quillsonDutchPhrase,
		language.French:               quillsonFrenchPhrase,
		language.German:               quillsonGermanPhrase,
		language.Italian:              quillsonItalianPhrase,
		language.Japanese:             quillsonJapanesePhrase,
		language.Korean:               quillsonKoreanPhrase,
		language.LatinAmericanSpanish: quillsonLatinAmericanSpanishPhrase,
		language.Russian:              quillsonRussianPhrase,
		language.SimplifiedChinese:    quillsonSimplifiedChinesePhrase,
		language.Spanish:              quillsonSpanishPhrase,
		language.TraditionalChinese:   quillsonTraditionalChinesePhrase}
)

var (
	// Quillson represents quillson.
	Quillson = nook.Villager{
		Character:   quillsonCharacter,
		Personality: personality.Smug,
		Phrase:      quillsonPhrase}
)
