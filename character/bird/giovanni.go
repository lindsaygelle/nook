package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	giovanniBirthday = nook.Birthday{
		Day:   6,
		Month: time.September}
)

var (
	giovanniCode = nook.Code{
		Value: "cwa"}
)

var (
	giovanniAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Giovanni"}

	giovanniCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Corbac"}

	giovanniDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	giovanniFrenchName = nook.Name{
		Language: language.French,
		Value:    "Corbac"}

	giovanniGermanName = nook.Name{
		Language: language.German,
		Value:    "Fiete"}

	giovanniItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Corvanni"}

	giovanniJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キャンタロー"}

	giovanniKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "캠식"}

	giovanniLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Corvín"}

	giovanniRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	giovanniSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	giovanniSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Corvín"}

	giovanniTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "露哥"}
)

var (
	giovanniName = nook.Languages{
		language.AmericanEnglish:      giovanniAmericanEnglishName,
		language.CanadianFrench:       giovanniCanadianFrenchName,
		language.Dutch:                giovanniDutchName,
		language.French:               giovanniFrenchName,
		language.German:               giovanniGermanName,
		language.Italian:              giovanniItalianName,
		language.Japanese:             giovanniJapaneseName,
		language.Korean:               giovanniKoreanName,
		language.LatinAmericanSpanish: giovanniLatinAmericanSpanishName,
		language.Russian:              giovanniRussianName,
		language.SimplifiedChinese:    giovanniSimplifiedChineseName,
		language.Spanish:              giovanniSpanishName,
		language.TraditionalChinese:   giovanniTraditionalChineseName}
)

var (
	giovanniCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: giovanniBirthday,
		Code:     giovanniCode,
		Key:      character.Giovanni,
		Gender:   gender.Male,
		Name:     giovanniName,
		Special:  true}
)

var (
	Giovanni = nook.Resident{
		Character: giovanniCharacter}
)
