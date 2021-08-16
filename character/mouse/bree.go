package mouse

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
	breeBirthday = nook.Birthday{
		Day:   7,
		Month: time.July}
)

var (
	breeCode = nook.Code{
		Value: "mus03"}
)

var (
	breeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bree"}

	breeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Quenotte"}

	breeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bree"}

	breeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Quenotte"}

	breeGermanName = nook.Name{
		Language: language.German,
		Value:    "Jenny"}

	breeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ginella"}

	breeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サラ"}

	breeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "사라"}

	breeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Brie"}

	breeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бри"}

	breeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "西瓜皮"}

	breeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Brie"}

	breeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "西瓜皮"}
)

var (
	breeName = nook.Languages{
		language.AmericanEnglish:      breeAmericanEnglishName,
		language.CanadianFrench:       breeCanadianFrenchName,
		language.Dutch:                breeDutchName,
		language.French:               breeFrenchName,
		language.German:               breeGermanName,
		language.Italian:              breeItalianName,
		language.Japanese:             breeJapaneseName,
		language.Korean:               breeKoreanName,
		language.LatinAmericanSpanish: breeLatinAmericanSpanishName,
		language.Russian:              breeRussianName,
		language.SimplifiedChinese:    breeSimplifiedChineseName,
		language.Spanish:              breeSpanishName,
		language.TraditionalChinese:   breeTraditionalChineseName}
)

var (
	breeCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: breeBirthday,
		Code:     breeCode,
		Key:      character.Bree,
		Gender:   gender.Female,
		Name:     breeName}
)

var (
	breeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cheeseball"}

	breeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "boulette"}

	breeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kaaskoekje"}

	breeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "boulette"}

	breeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kicher"}

	breeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "formaggino"}

	breeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なーんて"}

	breeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "글치머"}

	breeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "quesuá"}

	breeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "сы-ы-ыр"}

	breeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "说说的"}

	breeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "orejas"}

	breeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "說說的"}
)

var (
	breePhrase = nook.Languages{
		language.AmericanEnglish:      breeAmericanEnglishName,
		language.CanadianFrench:       breeCanadianFrenchName,
		language.Dutch:                breeDutchName,
		language.French:               breeFrenchName,
		language.German:               breeGermanName,
		language.Italian:              breeItalianName,
		language.Japanese:             breeJapaneseName,
		language.Korean:               breeKoreanName,
		language.LatinAmericanSpanish: breeLatinAmericanSpanishName,
		language.Russian:              breeRussianName,
		language.SimplifiedChinese:    breeSimplifiedChineseName,
		language.Spanish:              breeSpanishName,
		language.TraditionalChinese:   breeTraditionalChineseName}
)

var (
	Bree = nook.Villager{
		Character:   breeCharacter,
		Personality: personality.Snooty,
		Phrase:      breePhrase}
)
