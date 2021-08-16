package pig

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
	ganonBirthday = nook.Birthday{
		Day:   21,
		Month: time.February}
)

var (
	ganonCode = nook.Code{
		Value: "pig18"}
)

var (
	ganonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ganon"}

	ganonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ganon"}

	ganonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	ganonFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ganon"}

	ganonGermanName = nook.Name{
		Language: language.German,
		Value:    "Ganon"}

	ganonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ganon"}

	ganonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ガノン"}

	ganonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "가논"}

	ganonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ganon"}

	ganonRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	ganonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	ganonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ganon"}

	ganonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	ganonName = nook.Languages{
		language.AmericanEnglish:      ganonAmericanEnglishName,
		language.CanadianFrench:       ganonCanadianFrenchName,
		language.Dutch:                ganonDutchName,
		language.French:               ganonFrenchName,
		language.German:               ganonGermanName,
		language.Italian:              ganonItalianName,
		language.Japanese:             ganonJapaneseName,
		language.Korean:               ganonKoreanName,
		language.LatinAmericanSpanish: ganonLatinAmericanSpanishName,
		language.Russian:              ganonRussianName,
		language.SimplifiedChinese:    ganonSimplifiedChineseName,
		language.Spanish:              ganonSpanishName,
		language.TraditionalChinese:   ganonTraditionalChineseName}
)

var (
	ganonCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: ganonBirthday,
		Code:     ganonCode,
		Key:      character.Ganon,
		Gender:   gender.Male,
		Name:     ganonName}
)

var (
	ganonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "heh heh"}

	ganonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ganouïrk"}

	ganonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	ganonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ganouïrk"}

	ganonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "ganoink"}

	ganonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "trisgrunf"}

	ganonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "フォース"}

	ganonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "포스"}

	ganonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ganoink"}

	ganonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	ganonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	ganonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ganoink"}

	ganonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	ganonPhrase = nook.Languages{
		language.AmericanEnglish:      ganonAmericanEnglishName,
		language.CanadianFrench:       ganonCanadianFrenchName,
		language.Dutch:                ganonDutchName,
		language.French:               ganonFrenchName,
		language.German:               ganonGermanName,
		language.Italian:              ganonItalianName,
		language.Japanese:             ganonJapaneseName,
		language.Korean:               ganonKoreanName,
		language.LatinAmericanSpanish: ganonLatinAmericanSpanishName,
		language.Russian:              ganonRussianName,
		language.SimplifiedChinese:    ganonSimplifiedChineseName,
		language.Spanish:              ganonSpanishName,
		language.TraditionalChinese:   ganonTraditionalChineseName}
)

var (
	Ganon = nook.Villager{
		Character:   ganonCharacter,
		Personality: personality.Cranky,
		Phrase:      ganonPhrase}
)
