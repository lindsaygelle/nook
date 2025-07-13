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
	// breeBirthday represents bree birthday.
	breeBirthday = nook.Birthday{
		Day:   7,
		Month: time.July}
)

var (
	// breeCode represents bree code.
	breeCode = nook.Code{
		Value: "mus03"}
)

var (
	// breeAmericanEnglishName represents bree american english name.
	breeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bree"}

	// breeCanadianFrenchName represents bree canadian french name.
	breeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Quenotte"}

	// breeDutchName represents bree dutch name.
	breeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bree"}

	// breeFrenchName represents bree french name.
	breeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Quenotte"}

	// breeGermanName represents bree german name.
	breeGermanName = nook.Name{
		Language: language.German,
		Value:    "Jenny"}

	// breeItalianName represents bree italian name.
	breeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ginella"}

	// breeJapaneseName represents bree japanese name.
	breeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サラ"}

	// breeKoreanName represents bree korean name.
	breeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "사라"}

	// breeLatinAmericanSpanishName represents bree latin american spanish name.
	breeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Brie"}

	// breeRussianName represents bree russian name.
	breeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бри"}

	// breeSimplifiedChineseName represents bree simplified chinese name.
	breeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "西瓜皮"}

	// breeSpanishName represents bree spanish name.
	breeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Brie"}

	// breeTraditionalChineseName represents bree traditional chinese name.
	breeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "西瓜皮"}
)

var (
	// breeName represents bree name.
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
	// breeCharacter represents bree character.
	breeCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: breeBirthday,
		Code:     breeCode,
		Key:      character.Bree,
		Gender:   gender.Female,
		Name:     breeName,
		Special:  false}
)

var (
	// breeAmericanEnglishPhrase represents bree american english phrase.
	breeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cheeseball"}

	// breeCanadianFrenchPhrase represents bree canadian french phrase.
	breeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "boulette"}

	// breeDutchPhrase represents bree dutch phrase.
	breeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kaaskoekje"}

	// breeFrenchPhrase represents bree french phrase.
	breeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "boulette"}

	// breeGermanPhrase represents bree german phrase.
	breeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kicher"}

	// breeItalianPhrase represents bree italian phrase.
	breeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "formaggino"}

	// breeJapanesePhrase represents bree japanese phrase.
	breeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なーんて"}

	// breeKoreanPhrase represents bree korean phrase.
	breeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "글치머"}

	// breeLatinAmericanSpanishPhrase represents bree latin american spanish phrase.
	breeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "quesuá"}

	// breeRussianPhrase represents bree russian phrase.
	breeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "сы-ы-ыр"}

	// breeSimplifiedChinesePhrase represents bree simplified chinese phrase.
	breeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "说说的"}

	// breeSpanishPhrase represents bree spanish phrase.
	breeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "orejas"}

	// breeTraditionalChinesePhrase represents bree traditional chinese phrase.
	breeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "說說的"}
)

var (
	// breePhrase represents bree phrase.
	breePhrase = nook.Languages{
		language.AmericanEnglish:      breeAmericanEnglishPhrase,
		language.CanadianFrench:       breeCanadianFrenchPhrase,
		language.Dutch:                breeDutchPhrase,
		language.French:               breeFrenchPhrase,
		language.German:               breeGermanPhrase,
		language.Italian:              breeItalianPhrase,
		language.Japanese:             breeJapanesePhrase,
		language.Korean:               breeKoreanPhrase,
		language.LatinAmericanSpanish: breeLatinAmericanSpanishPhrase,
		language.Russian:              breeRussianPhrase,
		language.SimplifiedChinese:    breeSimplifiedChinesePhrase,
		language.Spanish:              breeSpanishPhrase,
		language.TraditionalChinese:   breeTraditionalChinesePhrase}
)

var (
	// Bree represents bree.
	Bree = nook.Villager{
		Character:   breeCharacter,
		Personality: personality.Snooty,
		Phrase:      breePhrase}
)
