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
	peckBirthday = nook.Birthday{
		Day:   25,
		Month: time.July}
)

var (
	peckCode = nook.Code{
		Value: "brd17"}
)

var (
	peckAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Peck"}

	peckCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pec"}

	peckDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Peck"}

	peckFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pec"}

	peckGermanName = nook.Name{
		Language: language.German,
		Value:    "Picko"}

	peckItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gheppio"}

	peckJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ふみたろう"}

	peckKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "문대"}

	peckLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Picuet"}

	peckRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пек"}

	peckSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "文雄"}

	peckSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Picuet"}

	peckTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "文雄"}
)

var (
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
	peckCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: peckBirthday,
		Code:     peckCode,
		Key:      character.Peck,
		Gender:   gender.Male,
		Name:     peckName}
)

var (
	peckAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "crunch"}

	peckCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "poupousse"}

	peckDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "wasbordje"}

	peckFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "poupousse"}

	peckGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "flatter"}

	peckItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sprint"}

	peckJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ブン"}

	peckKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "츠츠츠"}

	peckLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cachirp"}

	peckRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тук-тук"}

	peckSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "噗噗"}

	peckSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "muy fuerte"}

	peckTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "噗噗"}
)

var (
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
	Peck = nook.Villager{
		Character:   peckCharacter,
		Personality: personality.Jock,
		Phrase:      peckPhrase}
)
