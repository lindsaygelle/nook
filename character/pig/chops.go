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
	// chopsBirthday represents chops birthday.
	chopsBirthday = nook.Birthday{
		Day:   13,
		Month: time.October}
)

var (
	// chopsCode represents chops code.
	chopsCode = nook.Code{
		Value: "pig14"}
)

var (
	// chopsAmericanEnglishName represents chops american english name.
	chopsAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chops"}

	// chopsCanadianFrenchName represents chops canadian french name.
	chopsCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Aaron"}

	// chopsDutchName represents chops dutch name.
	chopsDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chops"}

	// chopsFrenchName represents chops french name.
	chopsFrenchName = nook.Name{
		Language: language.French,
		Value:    "Aaron"}

	// chopsGermanName represents chops german name.
	chopsGermanName = nook.Name{
		Language: language.German,
		Value:    "Clemens"}

	// chopsItalianName represents chops italian name.
	chopsItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lino"}

	// chopsJapaneseName represents chops japanese name.
	chopsJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "トンファン"}

	// chopsKoreanName represents chops korean name.
	chopsKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "돈후앙"}

	// chopsLatinAmericanSpanishName represents chops latin american spanish name.
	chopsLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Solomín"}

	// chopsRussianName represents chops russian name.
	chopsRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Чопс"}

	// chopsSimplifiedChineseName represents chops simplified chinese name.
	chopsSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "豚皇"}

	// chopsSpanishName represents chops spanish name.
	chopsSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Solomín"}

	// chopsTraditionalChineseName represents chops traditional chinese name.
	chopsTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "豚皇"}
)

var (
	// chopsName represents chops name.
	chopsName = nook.Languages{
		language.AmericanEnglish:      chopsAmericanEnglishName,
		language.CanadianFrench:       chopsCanadianFrenchName,
		language.Dutch:                chopsDutchName,
		language.French:               chopsFrenchName,
		language.German:               chopsGermanName,
		language.Italian:              chopsItalianName,
		language.Japanese:             chopsJapaneseName,
		language.Korean:               chopsKoreanName,
		language.LatinAmericanSpanish: chopsLatinAmericanSpanishName,
		language.Russian:              chopsRussianName,
		language.SimplifiedChinese:    chopsSimplifiedChineseName,
		language.Spanish:              chopsSpanishName,
		language.TraditionalChinese:   chopsTraditionalChineseName}
)

var (
	// chopsCharacter represents chops character.
	chopsCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: chopsBirthday,
		Code:     chopsCode,
		Key:      character.Chops,
		Gender:   gender.Male,
		Name:     chopsName,
		Special:  false}
)

var (
	// chopsAmericanEnglishPhrase represents chops american english phrase.
	chopsAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "zoink"}

	// chopsCanadianFrenchPhrase represents chops canadian french phrase.
	chopsCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "lardon"}

	// chopsDutchPhrase represents chops dutch phrase.
	chopsDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "moddervet"}

	// chopsFrenchPhrase represents chops french phrase.
	chopsFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "lardon"}

	// chopsGermanPhrase represents chops german phrase.
	chopsGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "borsti"}

	// chopsItalianPhrase represents chops italian phrase.
	chopsItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "oinc"}

	// chopsJapanesePhrase represents chops japanese phrase.
	chopsJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だトン"}

	// chopsKoreanPhrase represents chops korean phrase.
	chopsKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "떼끼에로"}

	// chopsLatinAmericanSpanishPhrase represents chops latin american spanish phrase.
	chopsLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "zoink"}

	// chopsRussianPhrase represents chops russian phrase.
	chopsRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хряк-с"}

	// chopsSimplifiedChinesePhrase represents chops simplified chinese phrase.
	chopsSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "豚"}

	// chopsSpanishPhrase represents chops spanish phrase.
	chopsSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "zoink"}

	// chopsTraditionalChinesePhrase represents chops traditional chinese phrase.
	chopsTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "豚"}
)

var (
	// chopsPhrase represents chops phrase.
	chopsPhrase = nook.Languages{
		language.AmericanEnglish:      chopsAmericanEnglishPhrase,
		language.CanadianFrench:       chopsCanadianFrenchPhrase,
		language.Dutch:                chopsDutchPhrase,
		language.French:               chopsFrenchPhrase,
		language.German:               chopsGermanPhrase,
		language.Italian:              chopsItalianPhrase,
		language.Japanese:             chopsJapanesePhrase,
		language.Korean:               chopsKoreanPhrase,
		language.LatinAmericanSpanish: chopsLatinAmericanSpanishPhrase,
		language.Russian:              chopsRussianPhrase,
		language.SimplifiedChinese:    chopsSimplifiedChinesePhrase,
		language.Spanish:              chopsSpanishPhrase,
		language.TraditionalChinese:   chopsTraditionalChinesePhrase}
)

var (
	// Chops represents chops.
	Chops = nook.Villager{
		Character:   chopsCharacter,
		Personality: personality.Smug,
		Phrase:      chopsPhrase}
)
