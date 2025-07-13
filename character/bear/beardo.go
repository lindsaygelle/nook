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

// beardoBirthday represents Beardo's birthday.
var (
	// beardoBirthday represents beardo birthday.
	beardoBirthday = nook.Birthday{
		Day:   27,
		Month: time.September}
)

// beardoCode represents Beardo's unique code.
var (
	// beardoCode represents beardo code.
	beardoCode = nook.Code{
		Value: "bea13"}
)

// Different names for Beardo in various languages.
var (
	// beardoAmericanEnglishName represents beardo american english name.
	beardoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Beardo"}

	// beardoCanadianFrenchName represents beardo canadian french name.
	beardoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Eustache"}

	// beardoDutchName represents beardo dutch name.
	beardoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Beardo"}

	// beardoFrenchName represents beardo french name.
	beardoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Eustache"}

	// beardoGermanName represents beardo german name.
	beardoGermanName = nook.Name{
		Language: language.German,
		Value:    "Berthold"}

	// beardoItalianName represents beardo italian name.
	beardoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Barnaba"}

	// beardoJapaneseName represents beardo japanese name.
	beardoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ベアード"}

	// beardoKoreanName represents beardo korean name.
	beardoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "베어드"}

	// beardoLatinAmericanSpanishName represents beardo latin american spanish name.
	beardoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bernabé"}

	// beardoRussianName represents beardo russian name.
	beardoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Беардо"}

	// beardoSimplifiedChineseName represents beardo simplified chinese name.
	beardoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "熊大叔"}

	// beardoSpanishName represents beardo spanish name.
	beardoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bernabé"}

	// beardoTraditionalChineseName represents beardo traditional chinese name.
	beardoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "熊大叔"}
)

// beardoName represents Beardo's name in different languages.
var (
	// beardoName represents beardo name.
	beardoName = nook.Languages{
		language.AmericanEnglish:      beardoAmericanEnglishName,
		language.CanadianFrench:       beardoCanadianFrenchName,
		language.Dutch:                beardoDutchName,
		language.French:               beardoFrenchName,
		language.German:               beardoGermanName,
		language.Italian:              beardoItalianName,
		language.Japanese:             beardoJapaneseName,
		language.Korean:               beardoKoreanName,
		language.LatinAmericanSpanish: beardoLatinAmericanSpanishName,
		language.Russian:              beardoRussianName,
		language.SimplifiedChinese:    beardoSimplifiedChineseName,
		language.Spanish:              beardoSpanishName,
		language.TraditionalChinese:   beardoTraditionalChineseName}
)

// beardoCharacter represents Beardo's character information.
var (
	// beardoCharacter represents beardo character.
	beardoCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: beardoBirthday,
		Code:     beardoCode,
		Key:      character.Beardo,
		Gender:   gender.Male,
		Name:     beardoName,
		Special:  false}
)

// Different phrases spoken by Beardo in various languages.
var (
	// beardoAmericanEnglishPhrase represents beardo american english phrase.
	beardoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "whiskers"}

	// beardoCanadianFrenchPhrase represents beardo canadian french phrase.
	beardoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ciao"}

	// beardoDutchPhrase represents beardo dutch phrase.
	beardoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snorhaar"}

	// beardoFrenchPhrase represents beardo french phrase.
	beardoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "poilodos"}

	// beardoGermanPhrase represents beardo german phrase.
	beardoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brrroho"}

	// beardoItalianPhrase represents beardo italian phrase.
	beardoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "andale"}

	// beardoJapanesePhrase represents beardo japanese phrase.
	beardoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "オッホン"}

	// beardoKoreanPhrase represents beardo korean phrase.
	beardoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "오홍홍"}

	// beardoLatinAmericanSpanishPhrase represents beardo latin american spanish phrase.
	beardoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bigotre"}

	// beardoRussianPhrase represents beardo russian phrase.
	beardoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "усы мои"}

	// beardoSimplifiedChinesePhrase represents beardo simplified chinese phrase.
	beardoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咳咳"}

	// beardoSpanishPhrase represents beardo spanish phrase.
	beardoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bigotre"}

	// beardoTraditionalChinesePhrase represents beardo traditional chinese phrase.
	beardoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咳咳"}
)

// beardoPhrase represents Beardo's phrases in different languages.
var (
	// beardoPhrase represents beardo phrase.
	beardoPhrase = nook.Languages{
		language.AmericanEnglish:      beardoAmericanEnglishPhrase,
		language.CanadianFrench:       beardoCanadianFrenchPhrase,
		language.Dutch:                beardoDutchPhrase,
		language.French:               beardoFrenchPhrase,
		language.German:               beardoGermanPhrase,
		language.Italian:              beardoItalianPhrase,
		language.Japanese:             beardoJapanesePhrase,
		language.Korean:               beardoKoreanPhrase,
		language.LatinAmericanSpanish: beardoLatinAmericanSpanishPhrase,
		language.Russian:              beardoRussianPhrase,
		language.SimplifiedChinese:    beardoSimplifiedChinesePhrase,
		language.Spanish:              beardoSpanishPhrase,
		language.TraditionalChinese:   beardoTraditionalChinesePhrase}
)

// Beardo represents the character Beardo with his complete information.
var (
	// Beardo represents beardo.
	Beardo = nook.Villager{
		Character:   beardoCharacter,
		Personality: personality.Smug,
		Phrase:      beardoPhrase}
)
