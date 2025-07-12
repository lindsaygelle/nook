package bearcub

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
	// pekoeBirthday represents pekoe birthday.
	pekoeBirthday = nook.Birthday{
		Day:   18,
		Month: time.May}
)

var (
	// pekoeCode represents pekoe code.
	pekoeCode = nook.Code{
		Value: "cbr14"}
)

var (
	// pekoeAmericanEnglishName represents pekoe american english name.
	pekoeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pekoe"}

	// pekoeCanadianFrenchName represents pekoe canadian french name.
	pekoeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pauline"}

	// pekoeDutchName represents pekoe dutch name.
	pekoeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pekoe"}

	// pekoeFrenchName represents pekoe french name.
	pekoeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pauline"}

	// pekoeGermanName represents pekoe german name.
	pekoeGermanName = nook.Name{
		Language: language.German,
		Value:    "Sandrine"}

	// pekoeItalianName represents pekoe italian name.
	pekoeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Mina"}

	// pekoeJapaneseName represents pekoe japanese name.
	pekoeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジャスミン"}

	// pekoeKoreanName represents pekoe korean name.
	pekoeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "재스민"}

	// pekoeLatinAmericanSpanishName represents pekoe latin american spanish name.
	pekoeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Vera"}

	// pekoeRussianName represents pekoe russian name.
	pekoeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пеко"}

	// pekoeSimplifiedChineseName represents pekoe simplified chinese name.
	pekoeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "佳敏"}

	// pekoeSpanishName represents pekoe spanish name.
	pekoeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Vera"}

	// pekoeTraditionalChineseName represents pekoe traditional chinese name.
	pekoeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "佳敏"}
)

var (
	// pekoeName represents pekoe name.
	pekoeName = nook.Languages{
		language.AmericanEnglish:      pekoeAmericanEnglishName,
		language.CanadianFrench:       pekoeCanadianFrenchName,
		language.Dutch:                pekoeDutchName,
		language.French:               pekoeFrenchName,
		language.German:               pekoeGermanName,
		language.Italian:              pekoeItalianName,
		language.Japanese:             pekoeJapaneseName,
		language.Korean:               pekoeKoreanName,
		language.LatinAmericanSpanish: pekoeLatinAmericanSpanishName,
		language.Russian:              pekoeRussianName,
		language.SimplifiedChinese:    pekoeSimplifiedChineseName,
		language.Spanish:              pekoeSpanishName,
		language.TraditionalChinese:   pekoeTraditionalChineseName}
)

var (
	// pekoeCharacter represents pekoe character.
	pekoeCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: pekoeBirthday,
		Code:     pekoeCode,
		Key:      character.Pekoe,
		Gender:   gender.Female,
		Name:     pekoeName,
		Special:  false}
)

var (
	// pekoeAmericanEnglishPhrase represents pekoe american english phrase.
	pekoeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bud"}

	// pekoeCanadianFrenchPhrase represents pekoe canadian french phrase.
	pekoeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "linou"}

	// pekoeDutchPhrase represents pekoe dutch phrase.
	pekoeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "bloempje"}

	// pekoeFrenchPhrase represents pekoe french phrase.
	pekoeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "linou"}

	// pekoeGermanPhrase represents pekoe german phrase.
	pekoeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "flüster"}

	// pekoeItalianPhrase represents pekoe italian phrase.
	pekoeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "purrrimba"}

	// pekoeJapanesePhrase represents pekoe japanese phrase.
	pekoeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "チャ"}

	// pekoeKoreanPhrase represents pekoe korean phrase.
	pekoeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "띵호아"}

	// pekoeLatinAmericanSpanishPhrase represents pekoe latin american spanish phrase.
	pekoeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "pitiminí"}

	// pekoeRussianPhrase represents pekoe russian phrase.
	pekoeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "цветочек"}

	// pekoeSimplifiedChinesePhrase represents pekoe simplified chinese phrase.
	pekoeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "佳佳"}

	// pekoeSpanishPhrase represents pekoe spanish phrase.
	pekoeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pitiminí"}

	// pekoeTraditionalChinesePhrase represents pekoe traditional chinese phrase.
	pekoeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "佳佳"}
)

var (
	// pekoePhrase represents pekoe phrase.
	pekoePhrase = nook.Languages{
		language.AmericanEnglish:      pekoeAmericanEnglishPhrase,
		language.CanadianFrench:       pekoeCanadianFrenchPhrase,
		language.Dutch:                pekoeDutchPhrase,
		language.French:               pekoeFrenchPhrase,
		language.German:               pekoeGermanPhrase,
		language.Italian:              pekoeItalianPhrase,
		language.Japanese:             pekoeJapanesePhrase,
		language.Korean:               pekoeKoreanPhrase,
		language.LatinAmericanSpanish: pekoeLatinAmericanSpanishPhrase,
		language.Russian:              pekoeRussianPhrase,
		language.SimplifiedChinese:    pekoeSimplifiedChinesePhrase,
		language.Spanish:              pekoeSpanishPhrase,
		language.TraditionalChinese:   pekoeTraditionalChinesePhrase}
)

var (
	// Pekoe represents pekoe.
	Pekoe = nook.Villager{
		Character:   pekoeCharacter,
		Personality: personality.Normal,
		Phrase:      pekoePhrase}
)
