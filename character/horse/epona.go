package horse

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
	// eponaBirthday represents epona birthday.
	eponaBirthday = nook.Birthday{
		Day:   21,
		Month: time.November}
)

var (
	// eponaCode represents epona code.
	eponaCode = nook.Code{
		Value: "hrs15"}
)

var (
	// eponaAmericanEnglishName represents epona american english name.
	eponaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Epona"}

	// eponaCanadianFrenchName represents epona canadian french name.
	eponaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Épona"}

	// eponaDutchName represents epona dutch name.
	eponaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// eponaFrenchName represents epona french name.
	eponaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Épona"}

	// eponaGermanName represents epona german name.
	eponaGermanName = nook.Name{
		Language: language.German,
		Value:    "Epona"}

	// eponaItalianName represents epona italian name.
	eponaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Epona"}

	// eponaJapaneseName represents epona japanese name.
	eponaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エポナ"}

	// eponaKoreanName represents epona korean name.
	eponaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "에포나"}

	// eponaLatinAmericanSpanishName represents epona latin american spanish name.
	eponaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Epona"}

	// eponaRussianName represents epona russian name.
	eponaRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// eponaSimplifiedChineseName represents epona simplified chinese name.
	eponaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// eponaSpanishName represents epona spanish name.
	eponaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Epona"}

	// eponaTraditionalChineseName represents epona traditional chinese name.
	eponaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// eponaName represents epona name.
	eponaName = nook.Languages{
		language.AmericanEnglish:      eponaAmericanEnglishName,
		language.CanadianFrench:       eponaCanadianFrenchName,
		language.Dutch:                eponaDutchName,
		language.French:               eponaFrenchName,
		language.German:               eponaGermanName,
		language.Italian:              eponaItalianName,
		language.Japanese:             eponaJapaneseName,
		language.Korean:               eponaKoreanName,
		language.LatinAmericanSpanish: eponaLatinAmericanSpanishName,
		language.Russian:              eponaRussianName,
		language.SimplifiedChinese:    eponaSimplifiedChineseName,
		language.Spanish:              eponaSpanishName,
		language.TraditionalChinese:   eponaTraditionalChineseName}
)

var (
	// eponaCharacter represents epona character.
	eponaCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: eponaBirthday,
		Code:     eponaCode,
		Key:      character.Epona,
		Gender:   gender.Female,
		Name:     eponaName,
		Special:  false}
)

var (
	// eponaAmericanEnglishPhrase represents epona american english phrase.
	eponaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "neigh"}

	// eponaCanadianFrenchPhrase represents epona canadian french phrase.
	eponaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "hippépique"}

	// eponaDutchPhrase represents epona dutch phrase.
	eponaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// eponaFrenchPhrase represents epona french phrase.
	eponaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "hippépique"}

	// eponaGermanPhrase represents epona german phrase.
	eponaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "lonlon"}

	// eponaItalianPhrase represents epona italian phrase.
	eponaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "liiinnn"}

	// eponaJapanesePhrase represents epona japanese phrase.
	eponaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ヒンヒン"}

	// eponaKoreanPhrase represents epona korean phrase.
	eponaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "히힝"}

	// eponaLatinAmericanSpanishPhrase represents epona latin american spanish phrase.
	eponaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "lon lon"}

	// eponaRussianPhrase represents epona russian phrase.
	eponaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// eponaSimplifiedChinesePhrase represents epona simplified chinese phrase.
	eponaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// eponaSpanishPhrase represents epona spanish phrase.
	eponaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "lon lon"}

	// eponaTraditionalChinesePhrase represents epona traditional chinese phrase.
	eponaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// eponaPhrase represents epona phrase.
	eponaPhrase = nook.Languages{
		language.AmericanEnglish:      eponaAmericanEnglishPhrase,
		language.CanadianFrench:       eponaCanadianFrenchPhrase,
		language.Dutch:                eponaDutchPhrase,
		language.French:               eponaFrenchPhrase,
		language.German:               eponaGermanPhrase,
		language.Italian:              eponaItalianPhrase,
		language.Japanese:             eponaJapanesePhrase,
		language.Korean:               eponaKoreanPhrase,
		language.LatinAmericanSpanish: eponaLatinAmericanSpanishPhrase,
		language.Russian:              eponaRussianPhrase,
		language.SimplifiedChinese:    eponaSimplifiedChinesePhrase,
		language.Spanish:              eponaSpanishPhrase,
		language.TraditionalChinese:   eponaTraditionalChinesePhrase}
)

var (
	// Epona represents epona.
	Epona = nook.Villager{
		Character:   eponaCharacter,
		Personality: personality.Peppy,
		Phrase:      eponaPhrase}
)
