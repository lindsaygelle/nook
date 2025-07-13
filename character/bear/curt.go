package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// curtBirthday represents Curt's birthday.
var (
	// curtBirthday represents curt birthday.
	curtBirthday = nook.Birthday{
		Day:   1,
		Month: time.July}
)

// curtCode represents Curt's unique code.
var (
	// curtCode represents curt code.
	curtCode = nook.Code{
		Value: "bea02"}
)

// Different names for Curt in various languages.
var (
	// curtAmericanEnglishName represents curt american english name.
	curtAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Curt"}

	// curtCanadianFrenchName represents curt canadian french name.
	curtCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Curt"}

	// curtDutchName represents curt dutch name.
	curtDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Curt"}

	// curtFrenchName represents curt french name.
	curtFrenchName = nook.Name{
		Language: language.French,
		Value:    "Curt"}

	// curtGermanName represents curt german name.
	curtGermanName = nook.Name{
		Language: language.German,
		Value:    "Kurt"}

	// curtItalianName represents curt italian name.
	curtItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Curt"}

	// curtJapaneseName represents curt japanese name.
	curtJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ガンテツ"}

	// curtKoreanName represents curt korean name.
	curtKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "뚝심"}

	// curtLatinAmericanSpanishName represents curt latin american spanish name.
	curtLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gorbaché"}

	// curtRussianName represents curt russian name.
	curtRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Керт"}

	// curtSimplifiedChineseName represents curt simplified chinese name.
	curtSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "铁熊"}

	// curtSpanishName represents curt spanish name.
	curtSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gorbaché"}

	// curtTraditionalChineseName represents curt traditional chinese name.
	curtTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鐵熊"}
)

// curtName represents Curt's name in different languages.
var (
	// curtName represents curt name.
	curtName = nook.Languages{
		language.AmericanEnglish:      curtAmericanEnglishName,
		language.CanadianFrench:       curtCanadianFrenchName,
		language.Dutch:                curtDutchName,
		language.French:               curtFrenchName,
		language.German:               curtGermanName,
		language.Italian:              curtItalianName,
		language.Japanese:             curtJapaneseName,
		language.Korean:               curtKoreanName,
		language.LatinAmericanSpanish: curtLatinAmericanSpanishName,
		language.Russian:              curtRussianName,
		language.SimplifiedChinese:    curtSimplifiedChineseName,
		language.Spanish:              curtSpanishName,
		language.TraditionalChinese:   curtTraditionalChineseName}
)

// curtCharacter represents Curt's character information.
var (
	// curtCharacter represents curt character.
	curtCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: curtBirthday,
		Code:     curtCode,
		Key:      character.Curt,
		Gender:   gender.Male,
		Name:     curtName,
		Special:  false}
)

// Different phrases spoken by Curt in various languages.
var (
	// curtAmericanEnglishPhrase represents curt american english phrase.
	curtAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "fuzzball"}

	// curtCanadianFrenchPhrase represents curt canadian french phrase.
	curtCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bouboule"}

	// curtDutchPhrase represents curt dutch phrase.
	curtDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "brombeer"}

	// curtFrenchPhrase represents curt french phrase.
	curtFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bouboule"}

	// curtGermanPhrase represents curt german phrase.
	curtGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grumml"}

	// curtItalianPhrase represents curt italian phrase.
	curtItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grrroan"}

	// curtJapanesePhrase represents curt japanese phrase.
	curtJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウム"}

	// curtKoreanPhrase represents curt korean phrase.
	curtKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "음"}

	// curtLatinAmericanSpanishPhrase represents curt latin american spanish phrase.
	curtLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "gromp"}

	// curtRussianPhrase represents curt russian phrase.
	curtRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "махрово"}

	// curtSimplifiedChinesePhrase represents curt simplified chinese phrase.
	curtSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嗯唔"}

	// curtSpanishPhrase represents curt spanish phrase.
	curtSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "oye"}

	// curtTraditionalChinesePhrase represents curt traditional chinese phrase.
	curtTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗯唔"}
)

// curtPhrase represents Curt's phrases in different languages.
var (
	// curtPhrase represents curt phrase.
	curtPhrase = nook.Languages{
		language.AmericanEnglish:      curtAmericanEnglishPhrase,
		language.CanadianFrench:       curtCanadianFrenchPhrase,
		language.Dutch:                curtDutchPhrase,
		language.French:               curtFrenchPhrase,
		language.German:               curtGermanPhrase,
		language.Italian:              curtItalianPhrase,
		language.Japanese:             curtJapanesePhrase,
		language.Korean:               curtKoreanPhrase,
		language.LatinAmericanSpanish: curtLatinAmericanSpanishPhrase,
		language.Russian:              curtRussianPhrase,
		language.SimplifiedChinese:    curtSimplifiedChinesePhrase,
		language.Spanish:              curtSpanishPhrase,
		language.TraditionalChinese:   curtTraditionalChinesePhrase}
)

// Curt represents the character Curt with his complete information.
var (
	// Curt represents curt.
	Curt = nook.Villager{
		Character:   curtCharacter,
		Personality: personality.Cranky,
		Phrase:      curtPhrase}
)
