package kangaroo

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
	carrieBirthday = nook.Birthday{
		Day:   5,
		Month: time.December}
)

var (
	carrieCode = nook.Code{
		Value: "kgr02"}
)

var (
	carrieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Carrie"}

	carrieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kanga"}

	carrieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Carrie"}

	carrieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kanga"}

	carrieGermanName = nook.Name{
		Language: language.German,
		Value:    "Carola"}

	carrieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Taschina"}

	carrieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マミィ"}

	carrieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마미"}

	carrieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lola"}

	carrieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кэрри"}

	carrieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "妈咪"}

	carrieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lola"}

	carrieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "媽咪"}
)

var (
	carrieName = nook.Languages{
		language.AmericanEnglish:      carrieAmericanEnglishName,
		language.CanadianFrench:       carrieCanadianFrenchName,
		language.Dutch:                carrieDutchName,
		language.French:               carrieFrenchName,
		language.German:               carrieGermanName,
		language.Italian:              carrieItalianName,
		language.Japanese:             carrieJapaneseName,
		language.Korean:               carrieKoreanName,
		language.LatinAmericanSpanish: carrieLatinAmericanSpanishName,
		language.Russian:              carrieRussianName,
		language.SimplifiedChinese:    carrieSimplifiedChineseName,
		language.Spanish:              carrieSpanishName,
		language.TraditionalChinese:   carrieTraditionalChineseName}
)

var (
	carrieCharacter = nook.Character{
		Animal:   animal.Kangaroo,
		Birthday: carrieBirthday,
		Code:     carrieCode,
		Key:      character.Carrie,
		Gender:   gender.Female,
		Name:     carrieName}
)

var (
	carrieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "little one"}

	carrieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon pit"}

	carrieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kleine"}

	carrieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "poupon"}

	carrieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "beutel"}

	carrieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gnegné"}

	carrieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だフン"}

	carrieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "오롤로"}

	carrieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "arrorró"}

	carrieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "дитятко"}

	carrieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "亲亲"}

	carrieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bebé"}

	carrieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "親親"}
)

var (
	carriePhrase = nook.Languages{
		language.AmericanEnglish:      carrieAmericanEnglishName,
		language.CanadianFrench:       carrieCanadianFrenchName,
		language.Dutch:                carrieDutchName,
		language.French:               carrieFrenchName,
		language.German:               carrieGermanName,
		language.Italian:              carrieItalianName,
		language.Japanese:             carrieJapaneseName,
		language.Korean:               carrieKoreanName,
		language.LatinAmericanSpanish: carrieLatinAmericanSpanishName,
		language.Russian:              carrieRussianName,
		language.SimplifiedChinese:    carrieSimplifiedChineseName,
		language.Spanish:              carrieSpanishName,
		language.TraditionalChinese:   carrieTraditionalChineseName}
)

var (
	Carrie = nook.Villager{
		Character:   carrieCharacter,
		Personality: personality.Normal,
		Phrase:      carriePhrase}
)
