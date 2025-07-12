package horse

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
	// reneighBirthday represents reneigh birthday.
	reneighBirthday = nook.Birthday{
		Day:   4,
		Month: time.June}
)

var (
	// reneighCode represents reneigh code.
	reneighCode = nook.Code{
		Value: "hrs16"}
)

var (
	// reneighAmericanEnglishName represents reneigh american english name.
	reneighAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Reneigh"}

	// reneighCanadianFrenchName represents reneigh canadian french name.
	reneighCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Jennifer"}

	// reneighDutchName represents reneigh dutch name.
	reneighDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Reneigh"}

	// reneighFrenchName represents reneigh french name.
	reneighFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jennifer"}

	// reneighGermanName represents reneigh german name.
	reneighGermanName = nook.Name{
		Language: language.German,
		Value:    "Andrea"}

	// reneighItalianName represents reneigh italian name.
	reneighItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Requina"}

	// reneighJapaneseName represents reneigh japanese name.
	reneighJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リアーナ"}

	// reneighKoreanName represents reneigh korean name.
	reneighKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "리아나"}

	// reneighLatinAmericanSpanishName represents reneigh latin american spanish name.
	reneighLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Corcelia"}

	// reneighRussianName represents reneigh russian name.
	reneighRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рени"}

	// reneighSimplifiedChineseName represents reneigh simplified chinese name.
	reneighSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雷哈娜"}

	// reneighSpanishName represents reneigh spanish name.
	reneighSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Corcelia"}

	// reneighTraditionalChineseName represents reneigh traditional chinese name.
	reneighTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雷哈娜"}
)

var (
	// reneighName represents reneigh name.
	reneighName = nook.Languages{
		language.AmericanEnglish:      reneighAmericanEnglishName,
		language.CanadianFrench:       reneighCanadianFrenchName,
		language.Dutch:                reneighDutchName,
		language.French:               reneighFrenchName,
		language.German:               reneighGermanName,
		language.Italian:              reneighItalianName,
		language.Japanese:             reneighJapaneseName,
		language.Korean:               reneighKoreanName,
		language.LatinAmericanSpanish: reneighLatinAmericanSpanishName,
		language.Russian:              reneighRussianName,
		language.SimplifiedChinese:    reneighSimplifiedChineseName,
		language.Spanish:              reneighSpanishName,
		language.TraditionalChinese:   reneighTraditionalChineseName}
)

var (
	// reneighCharacter represents reneigh character.
	reneighCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: reneighBirthday,
		Code:     reneighCode,
		Key:      character.Reneigh,
		Gender:   gender.Female,
		Name:     reneighName,
		Special:  false}
)

var (
	// reneighAmericanEnglishPhrase represents reneigh american english phrase.
	reneighAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ayup yup"}

	// reneighCanadianFrenchPhrase represents reneigh canadian french phrase.
	reneighCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "crincrin"}

	// reneighDutchPhrase represents reneigh dutch phrase.
	reneighDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "uhuh"}

	// reneighFrenchPhrase represents reneigh french phrase.
	reneighFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "crincrin"}

	// reneighGermanPhrase represents reneigh german phrase.
	reneighGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tada"}

	// reneighItalianPhrase represents reneigh italian phrase.
	reneighItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hiii-op"}

	// reneighJapanesePhrase represents reneigh japanese phrase.
	reneighJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "スン"}

	// reneighKoreanPhrase represents reneigh korean phrase.
	reneighKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쉬익"}

	// reneighLatinAmericanSpanishPhrase represents reneigh latin american spanish phrase.
	reneighLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "juááá"}

	// reneighRussianPhrase represents reneigh russian phrase.
	reneighRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "но-но"}

	// reneighSimplifiedChinesePhrase represents reneigh simplified chinese phrase.
	reneighSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哼"}

	// reneighSpanishPhrase represents reneigh spanish phrase.
	reneighSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "alehop"}

	// reneighTraditionalChinesePhrase represents reneigh traditional chinese phrase.
	reneighTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哼"}
)

var (
	// reneighPhrase represents reneigh phrase.
	reneighPhrase = nook.Languages{
		language.AmericanEnglish:      reneighAmericanEnglishPhrase,
		language.CanadianFrench:       reneighCanadianFrenchPhrase,
		language.Dutch:                reneighDutchPhrase,
		language.French:               reneighFrenchPhrase,
		language.German:               reneighGermanPhrase,
		language.Italian:              reneighItalianPhrase,
		language.Japanese:             reneighJapanesePhrase,
		language.Korean:               reneighKoreanPhrase,
		language.LatinAmericanSpanish: reneighLatinAmericanSpanishPhrase,
		language.Russian:              reneighRussianPhrase,
		language.SimplifiedChinese:    reneighSimplifiedChinesePhrase,
		language.Spanish:              reneighSpanishPhrase,
		language.TraditionalChinese:   reneighTraditionalChinesePhrase}
)

var (
	// Reneigh represents reneigh.
	Reneigh = nook.Villager{
		Character:   reneighCharacter,
		Personality: personality.BigSister,
		Phrase:      reneighPhrase}
)
