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
	// pierreBirthday represents pierre birthday.
	pierreBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// pierreCode represents pierre code.
	pierreCode = nook.Code{
		Value: ""}
)

var (
	// pierreAmericanEnglishName represents pierre american english name.
	pierreAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pierre"}

	// pierreCanadianFrenchName represents pierre canadian french name.
	pierreCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// pierreDutchName represents pierre dutch name.
	pierreDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// pierreFrenchName represents pierre french name.
	pierreFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// pierreGermanName represents pierre german name.
	pierreGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// pierreItalianName represents pierre italian name.
	pierreItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// pierreJapaneseName represents pierre japanese name.
	pierreJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピエール"}

	// pierreKoreanName represents pierre korean name.
	pierreKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// pierreLatinAmericanSpanishName represents pierre latin american spanish name.
	pierreLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// pierreRussianName represents pierre russian name.
	pierreRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// pierreSimplifiedChineseName represents pierre simplified chinese name.
	pierreSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// pierreSpanishName represents pierre spanish name.
	pierreSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// pierreTraditionalChineseName represents pierre traditional chinese name.
	pierreTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// pierreName represents pierre name.
	pierreName = nook.Languages{
		language.AmericanEnglish:      pierreAmericanEnglishName,
		language.CanadianFrench:       pierreCanadianFrenchName,
		language.Dutch:                pierreDutchName,
		language.French:               pierreFrenchName,
		language.German:               pierreGermanName,
		language.Italian:              pierreItalianName,
		language.Japanese:             pierreJapaneseName,
		language.Korean:               pierreKoreanName,
		language.LatinAmericanSpanish: pierreLatinAmericanSpanishName,
		language.Russian:              pierreRussianName,
		language.SimplifiedChinese:    pierreSimplifiedChineseName,
		language.Spanish:              pierreSpanishName,
		language.TraditionalChinese:   pierreTraditionalChineseName}
)

var (
	// pierreCharacter represents pierre character.
	pierreCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: pierreBirthday,
		Code:     pierreCode,
		Key:      character.Pierre,
		Gender:   gender.Male,
		Name:     pierreName,
		Special:  false}
)

var (
	// pierreAmericanEnglishPhrase represents pierre american english phrase.
	pierreAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ってね"}

	// pierreCanadianFrenchPhrase represents pierre canadian french phrase.
	pierreCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// pierreDutchPhrase represents pierre dutch phrase.
	pierreDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// pierreFrenchPhrase represents pierre french phrase.
	pierreFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// pierreGermanPhrase represents pierre german phrase.
	pierreGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// pierreItalianPhrase represents pierre italian phrase.
	pierreItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// pierreJapanesePhrase represents pierre japanese phrase.
	pierreJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ってね"}

	// pierreKoreanPhrase represents pierre korean phrase.
	pierreKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// pierreLatinAmericanSpanishPhrase represents pierre latin american spanish phrase.
	pierreLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// pierreRussianPhrase represents pierre russian phrase.
	pierreRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// pierreSimplifiedChinesePhrase represents pierre simplified chinese phrase.
	pierreSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// pierreSpanishPhrase represents pierre spanish phrase.
	pierreSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// pierreTraditionalChinesePhrase represents pierre traditional chinese phrase.
	pierreTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// pierrePhrase represents pierre phrase.
	pierrePhrase = nook.Languages{
		language.AmericanEnglish:      pierreAmericanEnglishPhrase,
		language.CanadianFrench:       pierreCanadianFrenchPhrase,
		language.Dutch:                pierreDutchPhrase,
		language.French:               pierreFrenchPhrase,
		language.German:               pierreGermanPhrase,
		language.Italian:              pierreItalianPhrase,
		language.Japanese:             pierreJapanesePhrase,
		language.Korean:               pierreKoreanPhrase,
		language.LatinAmericanSpanish: pierreLatinAmericanSpanishPhrase,
		language.Russian:              pierreRussianPhrase,
		language.SimplifiedChinese:    pierreSimplifiedChinesePhrase,
		language.Spanish:              pierreSpanishPhrase,
		language.TraditionalChinese:   pierreTraditionalChinesePhrase}
)

var (
	// Pierre represents pierre.
	Pierre = nook.Villager{
		Character:   pierreCharacter,
		Personality: personality.Jock,
		Phrase:      pierrePhrase}
)
