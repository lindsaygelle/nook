package bull

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
	// angusBirthday represents angus birthday.
	angusBirthday = nook.Birthday{
		Day:   30,
		Month: time.April}
)

var (
	// angusCode represents angus code.
	angusCode = nook.Code{
		Value: "bul00"}
)

var (
	// angusAmericanEnglishName represents angus american english name.
	angusAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Angus"}

	// angusCanadianFrenchName represents angus canadian french name.
	angusCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Angus"}

	// angusDutchName represents angus dutch name.
	angusDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Angus"}

	// angusFrenchName represents angus french name.
	angusFrenchName = nook.Name{
		Language: language.French,
		Value:    "Angus"}

	// angusGermanName represents angus german name.
	angusGermanName = nook.Name{
		Language: language.German,
		Value:    "Angus"}

	// angusItalianName represents angus italian name.
	angusItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Angus"}

	// angusJapaneseName represents angus japanese name.
	angusJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "セルバンテス"}

	// angusKoreanName represents angus korean name.
	angusKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "반데스"}

	// angusLatinAmericanSpanishName represents angus latin american spanish name.
	angusLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Aliste"}

	// angusRussianName represents angus russian name.
	angusRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Ангус"}

	// angusSimplifiedChineseName represents angus simplified chinese name.
	angusSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "施万德"}

	// angusSpanishName represents angus spanish name.
	angusSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Aliste"}

	// angusTraditionalChineseName represents angus traditional chinese name.
	angusTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "施萬德"}
)

var (
	// angusName represents angus name.
	angusName = nook.Languages{
		language.AmericanEnglish:      angusAmericanEnglishName,
		language.CanadianFrench:       angusCanadianFrenchName,
		language.Dutch:                angusDutchName,
		language.French:               angusFrenchName,
		language.German:               angusGermanName,
		language.Italian:              angusItalianName,
		language.Japanese:             angusJapaneseName,
		language.Korean:               angusKoreanName,
		language.LatinAmericanSpanish: angusLatinAmericanSpanishName,
		language.Russian:              angusRussianName,
		language.SimplifiedChinese:    angusSimplifiedChineseName,
		language.Spanish:              angusSpanishName,
		language.TraditionalChinese:   angusTraditionalChineseName}
)

var (
	// angusCharacter represents angus character.
	angusCharacter = nook.Character{
		Animal:   animal.Bull,
		Birthday: angusBirthday,
		Code:     angusCode,
		Key:      character.Angus,
		Gender:   gender.Male,
		Name:     angusName,
		Special:  false}
)

var (
	// angusAmericanEnglishPhrase represents angus american english phrase.
	angusAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "macmoo"}

	// angusCanadianFrenchPhrase represents angus canadian french phrase.
	angusCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "meuh meuh"}

	// angusDutchPhrase represents angus dutch phrase.
	angusDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "denderund"}

	// angusFrenchPhrase represents angus french phrase.
	angusFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "meuh meuh"}

	// angusGermanPhrase represents angus german phrase.
	angusGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brrruzzl"}

	// angusItalianPhrase represents angus italian phrase.
	angusItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "macmoo"}

	// angusJapanesePhrase represents angus japanese phrase.
	angusJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ふふ"}

	// angusKoreanPhrase represents angus korean phrase.
	angusKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "후후"}

	// angusLatinAmericanSpanishPhrase represents angus latin american spanish phrase.
	angusLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bufff"}

	// angusRussianPhrase represents angus russian phrase.
	angusRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "макму-у-у"}

	// angusSimplifiedChinesePhrase represents angus simplified chinese phrase.
	angusSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呼呼"}

	// angusSpanishPhrase represents angus spanish phrase.
	angusSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cuernos"}

	// angusTraditionalChinesePhrase represents angus traditional chinese phrase.
	angusTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "呼呼"}
)

var (
	// angusPhrase represents angus phrase.
	angusPhrase = nook.Languages{
		language.AmericanEnglish:      angusAmericanEnglishPhrase,
		language.CanadianFrench:       angusCanadianFrenchPhrase,
		language.Dutch:                angusDutchPhrase,
		language.French:               angusFrenchPhrase,
		language.German:               angusGermanPhrase,
		language.Italian:              angusItalianPhrase,
		language.Japanese:             angusJapanesePhrase,
		language.Korean:               angusKoreanPhrase,
		language.LatinAmericanSpanish: angusLatinAmericanSpanishPhrase,
		language.Russian:              angusRussianPhrase,
		language.SimplifiedChinese:    angusSimplifiedChinesePhrase,
		language.Spanish:              angusSpanishPhrase,
		language.TraditionalChinese:   angusTraditionalChinesePhrase}
)

var (
	// Angus represents angus.
	Angus = nook.Villager{
		Character:   angusCharacter,
		Personality: personality.Cranky,
		Phrase:      angusPhrase}
)
