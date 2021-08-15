package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Pecpoupousse"}

	peckDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Peckwasbordje"}

	peckFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pecpoupousse"}

	peckGermanName = nook.Name{
		Language: language.German,
		Value:    "Pickoflatter"}

	peckItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gheppiosprint"}

	peckJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ふみたろうブン"}

	peckKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "문대츠츠츠"}

	peckLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Picuetcachirp"}

	peckRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пектук-тук"}

	peckSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "文雄噗噗"}

	peckSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Picuetmuy fuerte"}

	peckTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "文雄噗噗"}
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
		Animal:   Bird,
		Birthday: peckBirthday,
		Code:     peckCode,
		Gender:   nook.Male,
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
	Peck = nook.Villager{
		Character:   peckCharacter,
		Personality: nook.Jock,
		Phrase:      peckPhrase}
)
