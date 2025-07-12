package cow

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
	// normaBirthday represents norma birthday.
	normaBirthday = nook.Birthday{
		Day:   20,
		Month: time.September}
)

var (
	// normaCode represents norma code.
	normaCode = nook.Code{
		Value: "cow06"}
)

var (
	// normaAmericanEnglishName represents norma american english name.
	normaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Norma"}

	// normaCanadianFrenchName represents norma canadian french name.
	normaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Norma"}

	// normaDutchName represents norma dutch name.
	normaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Norma"}

	// normaFrenchName represents norma french name.
	normaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Norma"}

	// normaGermanName represents norma german name.
	normaGermanName = nook.Name{
		Language: language.German,
		Value:    "Nelly"}

	// normaItalianName represents norma italian name.
	normaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Norma"}

	// normaJapaneseName represents norma japanese name.
	normaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "いさこ"}

	// normaKoreanName represents norma korean name.
	normaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "미자"}

	// normaLatinAmericanSpanishName represents norma latin american spanish name.
	normaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Norma"}

	// normaRussianName represents norma russian name.
	normaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Норма"}

	// normaSimplifiedChineseName represents norma simplified chinese name.
	normaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "晨曦"}

	// normaSpanishName represents norma spanish name.
	normaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Norma"}

	// normaTraditionalChineseName represents norma traditional chinese name.
	normaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "晨曦"}
)

var (
	// normaName represents norma name.
	normaName = nook.Languages{
		language.AmericanEnglish:      normaAmericanEnglishName,
		language.CanadianFrench:       normaCanadianFrenchName,
		language.Dutch:                normaDutchName,
		language.French:               normaFrenchName,
		language.German:               normaGermanName,
		language.Italian:              normaItalianName,
		language.Japanese:             normaJapaneseName,
		language.Korean:               normaKoreanName,
		language.LatinAmericanSpanish: normaLatinAmericanSpanishName,
		language.Russian:              normaRussianName,
		language.SimplifiedChinese:    normaSimplifiedChineseName,
		language.Spanish:              normaSpanishName,
		language.TraditionalChinese:   normaTraditionalChineseName}
)

var (
	// normaCharacter represents norma character.
	normaCharacter = nook.Character{
		Animal:   animal.Cow,
		Birthday: normaBirthday,
		Code:     normaCode,
		Key:      character.Norma,
		Gender:   gender.Female,
		Name:     normaName,
		Special:  false}
)

var (
	// normaAmericanEnglishPhrase represents norma american english phrase.
	normaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hoof hoo"}

	// normaCanadianFrenchPhrase represents norma canadian french phrase.
	normaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "meuh nan"}

	// normaDutchPhrase represents norma dutch phrase.
	normaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "boehoe"}

	// normaFrenchPhrase represents norma french phrase.
	normaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "meuh nan"}

	// normaGermanPhrase represents norma german phrase.
	normaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "muuuhi"}

	// normaItalianPhrase represents norma italian phrase.
	normaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "mu muuu"}

	// normaJapanesePhrase represents norma japanese phrase.
	normaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "うふ"}

	// normaKoreanPhrase represents norma korean phrase.
	normaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "에헤"}

	// normaLatinAmericanSpanishPhrase represents norma latin american spanish phrase.
	normaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "mu-mu"}

	// normaRussianPhrase represents norma russian phrase.
	normaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "до-ре-му"}

	// normaSimplifiedChinesePhrase represents norma simplified chinese phrase.
	normaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "微笑"}

	// normaSpanishPhrase represents norma spanish phrase.
	normaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mu-mu"}

	// normaTraditionalChinesePhrase represents norma traditional chinese phrase.
	normaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "微笑"}
)

var (
	// normaPhrase represents norma phrase.
	normaPhrase = nook.Languages{
		language.AmericanEnglish:      normaAmericanEnglishPhrase,
		language.CanadianFrench:       normaCanadianFrenchPhrase,
		language.Dutch:                normaDutchPhrase,
		language.French:               normaFrenchPhrase,
		language.German:               normaGermanPhrase,
		language.Italian:              normaItalianPhrase,
		language.Japanese:             normaJapanesePhrase,
		language.Korean:               normaKoreanPhrase,
		language.LatinAmericanSpanish: normaLatinAmericanSpanishPhrase,
		language.Russian:              normaRussianPhrase,
		language.SimplifiedChinese:    normaSimplifiedChinesePhrase,
		language.Spanish:              normaSpanishPhrase,
		language.TraditionalChinese:   normaTraditionalChinesePhrase}
)

var (
	// Norma represents norma.
	Norma = nook.Villager{
		Character:   normaCharacter,
		Personality: personality.Normal,
		Phrase:      normaPhrase}
)
