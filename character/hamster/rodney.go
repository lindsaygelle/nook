package hamster

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
	// rodneyBirthday represents rodney birthday.
	rodneyBirthday = nook.Birthday{
		Day:   10,
		Month: time.November}
)

var (
	// rodneyCode represents rodney code.
	rodneyCode = nook.Code{
		Value: "ham03"}
)

var (
	// rodneyAmericanEnglishName represents rodney american english name.
	rodneyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rodney"}

	// rodneyCanadianFrenchName represents rodney canadian french name.
	rodneyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Chico"}

	// rodneyDutchName represents rodney dutch name.
	rodneyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rodney"}

	// rodneyFrenchName represents rodney french name.
	rodneyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chico"}

	// rodneyGermanName represents rodney german name.
	rodneyGermanName = nook.Name{
		Language: language.German,
		Value:    "Reinhold"}

	// rodneyItalianName represents rodney italian name.
	rodneyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rosikkio"}

	// rodneyJapaneseName represents rodney japanese name.
	rodneyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジミー"}

	// rodneyKoreanName represents rodney korean name.
	rodneyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "지미"}

	// rodneyLatinAmericanSpanishName represents rodney latin american spanish name.
	rodneyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Chusquis"}

	// rodneyRussianName represents rodney russian name.
	rodneyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Родни"}

	// rodneySimplifiedChineseName represents rodney simplified chinese name.
	rodneySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "米米"}

	// rodneySpanishName represents rodney spanish name.
	rodneySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Chusquis"}

	// rodneyTraditionalChineseName represents rodney traditional chinese name.
	rodneyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "米米"}
)

var (
	// rodneyName represents rodney name.
	rodneyName = nook.Languages{
		language.AmericanEnglish:      rodneyAmericanEnglishName,
		language.CanadianFrench:       rodneyCanadianFrenchName,
		language.Dutch:                rodneyDutchName,
		language.French:               rodneyFrenchName,
		language.German:               rodneyGermanName,
		language.Italian:              rodneyItalianName,
		language.Japanese:             rodneyJapaneseName,
		language.Korean:               rodneyKoreanName,
		language.LatinAmericanSpanish: rodneyLatinAmericanSpanishName,
		language.Russian:              rodneyRussianName,
		language.SimplifiedChinese:    rodneySimplifiedChineseName,
		language.Spanish:              rodneySpanishName,
		language.TraditionalChinese:   rodneyTraditionalChineseName}
)

var (
	// rodneyCharacter represents rodney character.
	rodneyCharacter = nook.Character{
		Animal:   animal.Hamster,
		Birthday: rodneyBirthday,
		Code:     rodneyCode,
		Key:      character.Rodney,
		Gender:   gender.Male,
		Name:     rodneyName,
		Special:  false}
)

var (
	// rodneyAmericanEnglishPhrase represents rodney american english phrase.
	rodneyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "le ham"}

	// rodneyCanadianFrenchPhrase represents rodney canadian french phrase.
	rodneyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "yes sir"}

	// rodneyDutchPhrase represents rodney dutch phrase.
	rodneyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "le ham"}

	// rodneyFrenchPhrase represents rodney french phrase.
	rodneyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "caries"}

	// rodneyGermanPhrase represents rodney german phrase.
	rodneyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mjamjam"}

	// rodneyItalianPhrase represents rodney italian phrase.
	rodneyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "flop"}

	// rodneyJapanesePhrase represents rodney japanese phrase.
	rodneyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ヒュー"}

	// rodneyKoreanPhrase represents rodney korean phrase.
	rodneyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "휴우"}

	// rodneyLatinAmericanSpanishPhrase represents rodney latin american spanish phrase.
	rodneyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "adoya"}

	// rodneyRussianPhrase represents rodney russian phrase.
	rodneyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ле хом"}

	// rodneySimplifiedChinesePhrase represents rodney simplified chinese phrase.
	rodneySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咻咻"}

	// rodneySpanishPhrase represents rodney spanish phrase.
	rodneySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "adoya"}

	// rodneyTraditionalChinesePhrase represents rodney traditional chinese phrase.
	rodneyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咻咻"}
)

var (
	// rodneyPhrase represents rodney phrase.
	rodneyPhrase = nook.Languages{
		language.AmericanEnglish:      rodneyAmericanEnglishPhrase,
		language.CanadianFrench:       rodneyCanadianFrenchPhrase,
		language.Dutch:                rodneyDutchPhrase,
		language.French:               rodneyFrenchPhrase,
		language.German:               rodneyGermanPhrase,
		language.Italian:              rodneyItalianPhrase,
		language.Japanese:             rodneyJapanesePhrase,
		language.Korean:               rodneyKoreanPhrase,
		language.LatinAmericanSpanish: rodneyLatinAmericanSpanishPhrase,
		language.Russian:              rodneyRussianPhrase,
		language.SimplifiedChinese:    rodneySimplifiedChinesePhrase,
		language.Spanish:              rodneySpanishPhrase,
		language.TraditionalChinese:   rodneyTraditionalChinesePhrase}
)

var (
	// Rodney represents rodney.
	Rodney = nook.Villager{
		Character:   rodneyCharacter,
		Personality: personality.Smug,
		Phrase:      rodneyPhrase}
)
