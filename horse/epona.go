package horse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	eponaBirthday = nook.Birthday{
		Day:   21,
		Month: time.November}
)

var (
	eponaCode = nook.Code{
		Value: "hrs15"}
)

var (
	eponaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Epona"}

	eponaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Éponahippépique"}

	eponaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	eponaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Éponahippépique"}

	eponaGermanName = nook.Name{
		Language: language.German,
		Value:    "Eponalonlon"}

	eponaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Eponaliiinnn"}

	eponaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エポナヒンヒン"}

	eponaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "에포나히힝"}

	eponaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Eponalon lon"}

	eponaRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	eponaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	eponaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Eponalon lon"}

	eponaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
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
	eponaCharacter = nook.Character{
		Animal:   Horse,
		Birthday: eponaBirthday,
		Code:     eponaCode,
		Gender:   nook.Female,
		Name:     eponaName}
)

var (
	eponaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "neigh"}

	eponaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "hippépique"}

	eponaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	eponaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "hippépique"}

	eponaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "lonlon"}

	eponaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "liiinnn"}

	eponaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ヒンヒン"}

	eponaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "히힝"}

	eponaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "lon lon"}

	eponaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	eponaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	eponaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "lon lon"}

	eponaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	eponaPhrase = nook.Languages{
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
	Epona = nook.Villager{
		Character:   eponaCharacter,
		Personality: nook.Peppy,
		Phrase:      eponaPhrase}
)
