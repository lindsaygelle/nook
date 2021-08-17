package squirrel

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
	blaireBirthday = nook.Birthday{
		Day:   3,
		Month: time.July}
)

var (
	blaireCode = nook.Code{
		Value: "squ01"}
)

var (
	blaireAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Blaire"}

	blaireCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cachou"}

	blaireDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Blaire"}

	blaireFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cachou"}

	blaireGermanName = nook.Name{
		Language: language.German,
		Value:    "Klara"}

	blaireItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ghianda"}

	blaireJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シルエット"}

	blaireKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "실루엣"}

	blaireLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Azabache"}

	blaireRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Блер"}

	blaireSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小影"}

	blaireSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Azabache"}

	blaireTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小影"}
)

var (
	blaireName = nook.Languages{
		language.AmericanEnglish:      blaireAmericanEnglishName,
		language.CanadianFrench:       blaireCanadianFrenchName,
		language.Dutch:                blaireDutchName,
		language.French:               blaireFrenchName,
		language.German:               blaireGermanName,
		language.Italian:              blaireItalianName,
		language.Japanese:             blaireJapaneseName,
		language.Korean:               blaireKoreanName,
		language.LatinAmericanSpanish: blaireLatinAmericanSpanishName,
		language.Russian:              blaireRussianName,
		language.SimplifiedChinese:    blaireSimplifiedChineseName,
		language.Spanish:              blaireSpanishName,
		language.TraditionalChinese:   blaireTraditionalChineseName}
)

var (
	blaireCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: blaireBirthday,
		Code:     blaireCode,
		Key:      character.Blaire,
		Gender:   gender.Female,
		Name:     blaireName}
)

var (
	blaireAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nutlet"}

	blaireCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "coquille"}

	blaireDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "notendop"}

	blaireFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "coquille"}

	blaireGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knack"}

	blaireItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "scrick"}

	blaireJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ふふん"}

	blaireKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "헤헤헤"}

	blaireLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñicuá"}

	blaireRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "крошка"}

	blaireSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哼哼"}

	blaireSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bayita"}

	blaireTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哼哼"}
)

var (
	blairePhrase = nook.Languages{
		language.AmericanEnglish:      blaireAmericanEnglishPhrase,
		language.CanadianFrench:       blaireCanadianFrenchPhrase,
		language.Dutch:                blaireDutchPhrase,
		language.French:               blaireFrenchPhrase,
		language.German:               blaireGermanPhrase,
		language.Italian:              blaireItalianPhrase,
		language.Japanese:             blaireJapanesePhrase,
		language.Korean:               blaireKoreanPhrase,
		language.LatinAmericanSpanish: blaireLatinAmericanSpanishPhrase,
		language.Russian:              blaireRussianPhrase,
		language.SimplifiedChinese:    blaireSimplifiedChinesePhrase,
		language.Spanish:              blaireSpanishPhrase,
		language.TraditionalChinese:   blaireTraditionalChinesePhrase}
)

var (
	Blaire = nook.Villager{
		Character:   blaireCharacter,
		Personality: personality.Snooty,
		Phrase:      blairePhrase}
)
