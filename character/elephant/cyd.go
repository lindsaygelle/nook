package elephant

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	cydBirthday = nook.Birthday{
		Day:   9,
		Month: time.June}
)

var (
	cydCode = nook.Code{
		Value: "elp12"}
)

var (
	cydAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cyd"}

	cydCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Punk"}

	cydDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cyd"}

	cydFrenchName = nook.Name{
		Language: language.French,
		Value:    "Punk"}

	cydGermanName = nook.Name{
		Language: language.German,
		Value:    "Sid"}

	cydItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sid"}

	cydJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パンクス"}

	cydKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "펑크스"}

	cydLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ramón"}

	cydRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Сид"}

	cydSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "庞克斯"}

	cydSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ramón"}

	cydTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "龐克斯"}
)

var (
	cydName = nook.Languages{
		language.AmericanEnglish:      cydAmericanEnglishName,
		language.CanadianFrench:       cydCanadianFrenchName,
		language.Dutch:                cydDutchName,
		language.French:               cydFrenchName,
		language.German:               cydGermanName,
		language.Italian:              cydItalianName,
		language.Japanese:             cydJapaneseName,
		language.Korean:               cydKoreanName,
		language.LatinAmericanSpanish: cydLatinAmericanSpanishName,
		language.Russian:              cydRussianName,
		language.SimplifiedChinese:    cydSimplifiedChineseName,
		language.Spanish:              cydSpanishName,
		language.TraditionalChinese:   cydTraditionalChineseName}
)

var (
	cydCharacter = nook.Character{
		Animal:   animal.Elephant,
		Birthday: cydBirthday,
		Code:     cydCode,
		Gender:   gender.Male,
		Name:     cydName}
)

var (
	cydAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "rockin'"}

	cydCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "barda"}

	cydDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mieters"}

	cydFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bastringue"}

	cydGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hau"}

	cydItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "rocking"}

	cydJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウス"}

	cydKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "피스"}

	cydLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "rocanrol"}

	cydRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "рок-н-ролл"}

	cydSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "摇滚"}

	cydSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "rocanrol"}

	cydTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "搖滾"}
)

var (
	cydPhrase = nook.Languages{
		language.AmericanEnglish:      cydAmericanEnglishName,
		language.CanadianFrench:       cydCanadianFrenchName,
		language.Dutch:                cydDutchName,
		language.French:               cydFrenchName,
		language.German:               cydGermanName,
		language.Italian:              cydItalianName,
		language.Japanese:             cydJapaneseName,
		language.Korean:               cydKoreanName,
		language.LatinAmericanSpanish: cydLatinAmericanSpanishName,
		language.Russian:              cydRussianName,
		language.SimplifiedChinese:    cydSimplifiedChineseName,
		language.Spanish:              cydSpanishName,
		language.TraditionalChinese:   cydTraditionalChineseName}
)

var (
	Cyd = nook.Villager{
		Character:   cydCharacter,
		Personality: personality.Cranky,
		Phrase:      cydPhrase}
)
