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
	// sueeBirthday represents suee birthday.
	sueeBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// sueeCode represents suee code.
	sueeCode = nook.Code{
		Value: ""}
)

var (
	// sueeAmericanEnglishName represents suee american english name.
	sueeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sue E"}

	// sueeCanadianFrenchName represents suee canadian french name.
	sueeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// sueeDutchName represents suee dutch name.
	sueeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// sueeFrenchName represents suee french name.
	sueeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Peguy"}

	// sueeGermanName represents suee german name.
	sueeGermanName = nook.Name{
		Language: language.German,
		Value:    "Sabrina"}

	// sueeItalianName represents suee italian name.
	sueeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Piggy"}

	// sueeJapaneseName represents suee japanese name.
	sueeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ブルジョア"}

	// sueeKoreanName represents suee korean name.
	sueeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// sueeLatinAmericanSpanishName represents suee latin american spanish name.
	sueeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// sueeRussianName represents suee russian name.
	sueeRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// sueeSimplifiedChineseName represents suee simplified chinese name.
	sueeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "娇娇"}

	// sueeSpanishName represents suee spanish name.
	sueeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Porcia"}

	// sueeTraditionalChineseName represents suee traditional chinese name.
	sueeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// sueeName represents suee name.
	sueeName = nook.Languages{
		language.AmericanEnglish:      sueeAmericanEnglishName,
		language.CanadianFrench:       sueeCanadianFrenchName,
		language.Dutch:                sueeDutchName,
		language.French:               sueeFrenchName,
		language.German:               sueeGermanName,
		language.Italian:              sueeItalianName,
		language.Japanese:             sueeJapaneseName,
		language.Korean:               sueeKoreanName,
		language.LatinAmericanSpanish: sueeLatinAmericanSpanishName,
		language.Russian:              sueeRussianName,
		language.SimplifiedChinese:    sueeSimplifiedChineseName,
		language.Spanish:              sueeSpanishName,
		language.TraditionalChinese:   sueeTraditionalChineseName}
)

var (
	// sueeCharacter represents suee character.
	sueeCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: sueeBirthday,
		Code:     sueeCode,
		Key:      character.SueE,
		Gender:   gender.Female,
		Name:     sueeName,
		Special:  false}
)

var (
	// sueeAmericanEnglishPhrase represents suee american english phrase.
	sueeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snort"}

	// sueeCanadianFrenchPhrase represents suee canadian french phrase.
	sueeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// sueeDutchPhrase represents suee dutch phrase.
	sueeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// sueeFrenchPhrase represents suee french phrase.
	sueeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "groingroin"}

	// sueeGermanPhrase represents suee german phrase.
	sueeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnaub"}

	// sueeItalianPhrase represents suee italian phrase.
	sueeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "oinksnoink"}

	// sueeJapanesePhrase represents suee japanese phrase.
	sueeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ざんす"}

	// sueeKoreanPhrase represents suee korean phrase.
	sueeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// sueeLatinAmericanSpanishPhrase represents suee latin american spanish phrase.
	sueeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// sueeRussianPhrase represents suee russian phrase.
	sueeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// sueeSimplifiedChinesePhrase represents suee simplified chinese phrase.
	sueeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "烦死"}

	// sueeSpanishPhrase represents suee spanish phrase.
	sueeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "piojo"}

	// sueeTraditionalChinesePhrase represents suee traditional chinese phrase.
	sueeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// sueePhrase represents suee phrase.
	sueePhrase = nook.Languages{
		language.AmericanEnglish:      sueeAmericanEnglishPhrase,
		language.CanadianFrench:       sueeCanadianFrenchPhrase,
		language.Dutch:                sueeDutchPhrase,
		language.French:               sueeFrenchPhrase,
		language.German:               sueeGermanPhrase,
		language.Italian:              sueeItalianPhrase,
		language.Japanese:             sueeJapanesePhrase,
		language.Korean:               sueeKoreanPhrase,
		language.LatinAmericanSpanish: sueeLatinAmericanSpanishPhrase,
		language.Russian:              sueeRussianPhrase,
		language.SimplifiedChinese:    sueeSimplifiedChinesePhrase,
		language.Spanish:              sueeSpanishPhrase,
		language.TraditionalChinese:   sueeTraditionalChinesePhrase}
)

var (
	// SueE represents sue e.
	SueE = nook.Villager{
		Character:   sueeCharacter,
		Personality: personality.Snooty,
		Phrase:      sueePhrase}
)
