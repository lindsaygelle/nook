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
	// admiralBirthday represents admiral birthday.
	admiralBirthday = nook.Birthday{
		Day:   27,
		Month: time.January}
)

var (
	// admiralCode represents admiral code.
	admiralCode = nook.Code{
		Value: "brd06"}
)

var (
	// admiralAmericanEnglishName represents admiral american english name.
	admiralAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Admiral"}

	// admiralCanadianFrenchName represents admiral canadian french name.
	admiralCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Maréchal"}

	// admiralDutchName represents admiral dutch name.
	admiralDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Admiral"}

	// admiralFrenchName represents admiral french name.
	admiralFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maréchal"}

	// admiralGermanName represents admiral german name.
	admiralGermanName = nook.Name{
		Language: language.German,
		Value:    "Ansgar"}

	// admiralItalianName represents admiral italian name.
	admiralItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Adolfo"}

	// admiralJapaneseName represents admiral japanese name.
	admiralJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "イッテツ"}

	// admiralKoreanName represents admiral korean name.
	admiralKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "일섭"}

	// admiralLatinAmericanSpanishName represents admiral latin american spanish name.
	admiralLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Avutardo"}

	// admiralRussianName represents admiral russian name.
	admiralRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Адмирал"}

	// admiralSimplifiedChineseName represents admiral simplified chinese name.
	admiralSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "李彻"}

	// admiralSpanishName represents admiral spanish name.
	admiralSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Avutardo"}

	// admiralTraditionalChineseName represents admiral traditional chinese name.
	admiralTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "李徹"}
)

var (
	// admiralName represents admiral name.
	admiralName = nook.Languages{
		language.AmericanEnglish:      admiralAmericanEnglishName,
		language.CanadianFrench:       admiralCanadianFrenchName,
		language.Dutch:                admiralDutchName,
		language.French:               admiralFrenchName,
		language.German:               admiralGermanName,
		language.Italian:              admiralItalianName,
		language.Japanese:             admiralJapaneseName,
		language.Korean:               admiralKoreanName,
		language.LatinAmericanSpanish: admiralLatinAmericanSpanishName,
		language.Russian:              admiralRussianName,
		language.SimplifiedChinese:    admiralSimplifiedChineseName,
		language.Spanish:              admiralSpanishName,
		language.TraditionalChinese:   admiralTraditionalChineseName}
)

var (
	// admiralCharacter represents admiral character.
	admiralCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: admiralBirthday,
		Code:     admiralCode,
		Key:      character.Admiral,
		Gender:   gender.Male,
		Name:     admiralName,
		Special:  false}
)

var (
	// admiralAmericanEnglishPhrase represents admiral american english phrase.
	admiralAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "aye aye"}

	// admiralCanadianFrenchPhrase represents admiral canadian french phrase.
	admiralCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "cornichon"}

	// admiralDutchPhrase represents admiral dutch phrase.
	admiralDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "yep yep"}

	// admiralFrenchPhrase represents admiral french phrase.
	admiralFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "cornichon"}

	// admiralGermanPhrase represents admiral german phrase.
	admiralGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "oh, ja"}

	// admiralItalianPhrase represents admiral italian phrase.
	admiralItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cipissi"}

	// admiralJapanesePhrase represents admiral japanese phrase.
	admiralJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ってか"}

	// admiralKoreanPhrase represents admiral korean phrase.
	admiralKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "오예"}

	// admiralLatinAmericanSpanishPhrase represents admiral latin american spanish phrase.
	admiralLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "piuc-piuc"}

	// admiralRussianPhrase represents admiral russian phrase.
	admiralRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "чиф-чиф"}

	// admiralSimplifiedChinesePhrase represents admiral simplified chinese phrase.
	admiralSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "贯彻"}

	// admiralSpanishPhrase represents admiral spanish phrase.
	admiralSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "avechucho"}

	// admiralTraditionalChinesePhrase represents admiral traditional chinese phrase.
	admiralTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "貫徹"}
)

var (
	// admiralPhrase represents admiral phrase.
	admiralPhrase = nook.Languages{
		language.AmericanEnglish:      admiralAmericanEnglishPhrase,
		language.CanadianFrench:       admiralCanadianFrenchPhrase,
		language.Dutch:                admiralDutchPhrase,
		language.French:               admiralFrenchPhrase,
		language.German:               admiralGermanPhrase,
		language.Italian:              admiralItalianPhrase,
		language.Japanese:             admiralJapanesePhrase,
		language.Korean:               admiralKoreanPhrase,
		language.LatinAmericanSpanish: admiralLatinAmericanSpanishPhrase,
		language.Russian:              admiralRussianPhrase,
		language.SimplifiedChinese:    admiralSimplifiedChinesePhrase,
		language.Spanish:              admiralSpanishPhrase,
		language.TraditionalChinese:   admiralTraditionalChinesePhrase}
)

var (
	// Admiral represents admiral.
	Admiral = nook.Villager{
		Character:   admiralCharacter,
		Personality: personality.Cranky,
		Phrase:      admiralPhrase}
)
