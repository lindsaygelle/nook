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
	// etoileBirthday represents Étoile's birthday (December 25th).
	etoileBirthday = nook.Birthday{
		Day:   25,
		Month: time.December}
)

var (
	// etoileCode represents Étoile's unique code ("shp14").
	etoileCode = nook.Code{
		Value: "shp14"}
)

var (
	// etoileAmericanEnglishName represents Étoile's name in American English.
	etoileAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Étoile"}

	// etoileCanadianFrenchName represents Étoile's name in Canadian French.
	etoileCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Étoile"}

	// etoileDutchName represents Étoile's name in Dutch.
	etoileDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Étoile"}

	// etoileFrenchName represents Étoile's name in French.
	etoileFrenchName = nook.Name{
		Language: language.French,
		Value:    "Étoile"}

	// etoileGermanName represents Étoile's name in German.
	etoileGermanName = nook.Name{
		Language: language.German,
		Value:    "Etoile"}

	// etoileItalianName represents Étoile's name in Italian.
	etoileItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Étoile"}

	// etoileJapaneseName represents Étoile's name in Japanese.
	etoileJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エトワール"}

	// etoileKoreanName represents Étoile's name in Korean.
	etoileKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "에뜨와르"}

	// etoileLatinAmericanSpanishName represents Étoile's name in Latin American Spanish.
	etoileLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Etoile"}

	// etoileRussianName represents Étoile's name in Russian.
	etoileRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Этуаль"}

	// etoileSimplifiedChineseName represents Étoile's name in Simplified Chinese.
	etoileSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "彩星"}

	// etoileSpanishName represents Étoile's name in Spanish.
	etoileSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Etoile"}

	// etoileTraditionalChineseName represents Étoile's name in Traditional Chinese.
	etoileTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "彩星"}
)

var (
	// etoileName represents Étoile's name in different languages.
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
	// etoileCharacter represents Étoile's character information.
	etoileCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: etoileBirthday,
		Code:     etoileCode,
		Key:      character.Etoile,
		Gender:   gender.Female,
		Name:     etoileName,
		Special:  false}
)

var (
	// etoileAmericanEnglishPhrase represents Étoile's phrase in American English.
	etoileAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "fuzzy"}

	// etoileCanadianFrenchPhrase represents Étoile's phrase in Canadian French.
	etoileCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mohair"}

	// etoileDutchPhrase represents Étoile's phrase in Dutch.
	etoileDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kikilala"}

	// etoileFrenchPhrase represents Étoile's phrase in French.
	etoileFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mohair"}

	// etoileGermanPhrase represents Étoile's phrase in German.
	etoileGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kikilala"}

	// etoileItalianPhrase represents Étoile's phrase in Italian.
	etoileItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ovinin"}

	// etoileJapanesePhrase represents Étoile's phrase in Japanese.
	etoileJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "メロメロ"}

	// etoileKoreanPhrase represents Étoile's phrase in Korean.
	etoileKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "폭신폭신"}

	// etoileLatinAmericanSpanishPhrase represents Étoile's phrase in Latin American Spanish.
	etoileLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "kikilala"}

	// etoileRussianPhrase represents Étoile's phrase in Russian.
	etoileRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мягенько"}

	// etoileSimplifiedChinesePhrase represents Étoile's phrase in Simplified Chinese.
	etoileSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "毛茸茸"}

	// etoileSpanishPhrase represents Étoile's phrase in Spanish.
	etoileSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "kikilala"}

	// etoileTraditionalChinesePhrase represents Étoile's phrase in Traditional Chinese.
	etoileTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "毛茸茸"}
)

var (
	// etoilePhrase represents Étoile's phrases in different languages.
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
	// Etoile represents the character Étoile with her complete information.
	Etoile = nook.Villager{
		Character:   etoileCharacter,
		Personality: personality.Normal,
		Phrase:      etoilePhrase}
)
