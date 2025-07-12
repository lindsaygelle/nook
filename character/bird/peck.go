package bird

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
	// peckBirthday represents peck birthday.
	peckBirthday = nook.Birthday{
		Day:   25,
		Month: time.July}
)

var (
	// peckCode represents peck code.
	peckCode = nook.Code{
		Value: "brd17"}
)

var (
	// peckAmericanEnglishName represents peck american english name.
	peckAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Peck"}

	// peckCanadianFrenchName represents peck canadian french name.
	peckCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pec"}

	// peckDutchName represents peck dutch name.
	peckDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Peck"}

	// peckFrenchName represents peck french name.
	peckFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pec"}

	// peckGermanName represents peck german name.
	peckGermanName = nook.Name{
		Language: language.German,
		Value:    "Picko"}

	// peckItalianName represents peck italian name.
	peckItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gheppio"}

	// peckJapaneseName represents peck japanese name.
	peckJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ふみたろう"}

	// peckKoreanName represents peck korean name.
	peckKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "문대"}

	// peckLatinAmericanSpanishName represents peck latin american spanish name.
	peckLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Picuet"}

	// peckRussianName represents peck russian name.
	peckRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пек"}

	// peckSimplifiedChineseName represents peck simplified chinese name.
	peckSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "文雄"}

	// peckSpanishName represents peck spanish name.
	peckSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Picuet"}

	// peckTraditionalChineseName represents peck traditional chinese name.
	peckTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "文雄"}
)

var (
	// peckName represents peck name.
	peckName = nook.Languages{
		language.AmericanEnglish:      peckAmericanEnglishName,
		language.CanadianFrench:       peckCanadianFrenchName,
		language.Dutch:                peckDutchName,
		language.French:               peckFrenchName,
		language.German:               peckGermanName,
		language.Italian:              peckItalianName,
		language.Japanese:             peckJapaneseName,
		language.Korean:               peckKoreanName,
		language.LatinAmericanSpanish: peckLatinAmericanSpanishName,
		language.Russian:              peckRussianName,
		language.SimplifiedChinese:    peckSimplifiedChineseName,
		language.Spanish:              peckSpanishName,
		language.TraditionalChinese:   peckTraditionalChineseName}
)

var (
	// peckCharacter represents peck character.
	peckCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: peckBirthday,
		Code:     peckCode,
		Key:      character.Peck,
		Gender:   gender.Male,
		Name:     peckName,
		Special:  false}
)

var (
	// peckAmericanEnglishPhrase represents peck american english phrase.
	peckAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "crunch"}

	// peckCanadianFrenchPhrase represents peck canadian french phrase.
	peckCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "poupousse"}

	// peckDutchPhrase represents peck dutch phrase.
	peckDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "wasbordje"}

	// peckFrenchPhrase represents peck french phrase.
	peckFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "poupousse"}

	// peckGermanPhrase represents peck german phrase.
	peckGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "flatter"}

	// peckItalianPhrase represents peck italian phrase.
	peckItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sprint"}

	// peckJapanesePhrase represents peck japanese phrase.
	peckJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ブン"}

	// peckKoreanPhrase represents peck korean phrase.
	peckKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "츠츠츠"}

	// peckLatinAmericanSpanishPhrase represents peck latin american spanish phrase.
	peckLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cachirp"}

	// peckRussianPhrase represents peck russian phrase.
	peckRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тук-тук"}

	// peckSimplifiedChinesePhrase represents peck simplified chinese phrase.
	peckSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "噗噗"}

	// peckSpanishPhrase represents peck spanish phrase.
	peckSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "muy fuerte"}

	// peckTraditionalChinesePhrase represents peck traditional chinese phrase.
	peckTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "噗噗"}
)

var (
	// peckPhrase represents peck phrase.
	peckPhrase = nook.Languages{
		language.AmericanEnglish:      peckAmericanEnglishPhrase,
		language.CanadianFrench:       peckCanadianFrenchPhrase,
		language.Dutch:                peckDutchPhrase,
		language.French:               peckFrenchPhrase,
		language.German:               peckGermanPhrase,
		language.Italian:              peckItalianPhrase,
		language.Japanese:             peckJapanesePhrase,
		language.Korean:               peckKoreanPhrase,
		language.LatinAmericanSpanish: peckLatinAmericanSpanishPhrase,
		language.Russian:              peckRussianPhrase,
		language.SimplifiedChinese:    peckSimplifiedChinesePhrase,
		language.Spanish:              peckSpanishPhrase,
		language.TraditionalChinese:   peckTraditionalChinesePhrase}
)

var (
	// Peck represents peck.
	Peck = nook.Villager{
		Character:   peckCharacter,
		Personality: personality.Jock,
		Phrase:      peckPhrase}
)
