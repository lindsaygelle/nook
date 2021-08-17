package sheep

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
	etoileBirthday = nook.Birthday{
		Day:   25,
		Month: time.December}
)

var (
	etoileCode = nook.Code{
		Value: "shp14"}
)

var (
	etoileAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Étoile"}

	etoileCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Étoile"}

	etoileDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Étoile"}

	etoileFrenchName = nook.Name{
		Language: language.French,
		Value:    "Étoile"}

	etoileGermanName = nook.Name{
		Language: language.German,
		Value:    "Etoile"}

	etoileItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Étoile"}

	etoileJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エトワール"}

	etoileKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "에뜨와르"}

	etoileLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Etoile"}

	etoileRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Этуаль"}

	etoileSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "彩星"}

	etoileSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Etoile"}

	etoileTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "彩星"}
)

var (
	etoileName = nook.Languages{
		language.AmericanEnglish:      etoileAmericanEnglishName,
		language.CanadianFrench:       etoileCanadianFrenchName,
		language.Dutch:                etoileDutchName,
		language.French:               etoileFrenchName,
		language.German:               etoileGermanName,
		language.Italian:              etoileItalianName,
		language.Japanese:             etoileJapaneseName,
		language.Korean:               etoileKoreanName,
		language.LatinAmericanSpanish: etoileLatinAmericanSpanishName,
		language.Russian:              etoileRussianName,
		language.SimplifiedChinese:    etoileSimplifiedChineseName,
		language.Spanish:              etoileSpanishName,
		language.TraditionalChinese:   etoileTraditionalChineseName}
)

var (
	etoileCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: etoileBirthday,
		Code:     etoileCode,
		Key:      character.Etoile,
		Gender:   gender.Female,
		Name:     etoileName}
)

var (
	etoileAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "fuzzy"}

	etoileCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mohair"}

	etoileDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kikilala"}

	etoileFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mohair"}

	etoileGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kikilala"}

	etoileItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ovinin"}

	etoileJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "メロメロ"}

	etoileKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "폭신폭신"}

	etoileLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "kikilala"}

	etoileRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мягенько"}

	etoileSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "毛茸茸"}

	etoileSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "kikilala"}

	etoileTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "毛茸茸"}
)

var (
	etoilePhrase = nook.Languages{
		language.AmericanEnglish:      etoileAmericanEnglishPhrase,
		language.CanadianFrench:       etoileCanadianFrenchPhrase,
		language.Dutch:                etoileDutchPhrase,
		language.French:               etoileFrenchPhrase,
		language.German:               etoileGermanPhrase,
		language.Italian:              etoileItalianPhrase,
		language.Japanese:             etoileJapanesePhrase,
		language.Korean:               etoileKoreanPhrase,
		language.LatinAmericanSpanish: etoileLatinAmericanSpanishPhrase,
		language.Russian:              etoileRussianPhrase,
		language.SimplifiedChinese:    etoileSimplifiedChinesePhrase,
		language.Spanish:              etoileSpanishPhrase,
		language.TraditionalChinese:   etoileTraditionalChinesePhrase}
)

var (
	Etoile = nook.Villager{
		Character:   etoileCharacter,
		Personality: personality.Normal,
		Phrase:      etoilePhrase}
)
