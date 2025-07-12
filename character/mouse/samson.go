package mouse

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
	// samsonBirthday represents samson birthday.
	samsonBirthday = nook.Birthday{
		Day:   5,
		Month: time.July}
)

var (
	// samsonCode represents samson code.
	samsonCode = nook.Code{
		Value: "mus04"}
)

var (
	// samsonAmericanEnglishName represents samson american english name.
	samsonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Samson"}

	// samsonCanadianFrenchName represents samson canadian french name.
	samsonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Samson"}

	// samsonDutchName represents samson dutch name.
	samsonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Samson"}

	// samsonFrenchName represents samson french name.
	samsonFrenchName = nook.Name{
		Language: language.French,
		Value:    "Samson"}

	// samsonGermanName represents samson german name.
	samsonGermanName = nook.Name{
		Language: language.German,
		Value:    "Samson"}

	// samsonItalianName represents samson italian name.
	samsonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sansone"}

	// samsonJapaneseName represents samson japanese name.
	samsonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピース"}

	// samsonKoreanName represents samson korean name.
	samsonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "피스"}

	// samsonLatinAmericanSpanishName represents samson latin american spanish name.
	samsonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sansón"}

	// samsonRussianName represents samson russian name.
	samsonRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Самсон"}

	// samsonSimplifiedChineseName represents samson simplified chinese name.
	samsonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "和平"}

	// samsonSpanishName represents samson spanish name.
	samsonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sansón"}

	// samsonTraditionalChineseName represents samson traditional chinese name.
	samsonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "和平"}
)

var (
	// samsonName represents samson name.
	samsonName = nook.Languages{
		language.AmericanEnglish:      samsonAmericanEnglishName,
		language.CanadianFrench:       samsonCanadianFrenchName,
		language.Dutch:                samsonDutchName,
		language.French:               samsonFrenchName,
		language.German:               samsonGermanName,
		language.Italian:              samsonItalianName,
		language.Japanese:             samsonJapaneseName,
		language.Korean:               samsonKoreanName,
		language.LatinAmericanSpanish: samsonLatinAmericanSpanishName,
		language.Russian:              samsonRussianName,
		language.SimplifiedChinese:    samsonSimplifiedChineseName,
		language.Spanish:              samsonSpanishName,
		language.TraditionalChinese:   samsonTraditionalChineseName}
)

var (
	// samsonCharacter represents samson character.
	samsonCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: samsonBirthday,
		Code:     samsonCode,
		Key:      character.Samson,
		Gender:   gender.Male,
		Name:     samsonName,
		Special:  false}
)

var (
	// samsonAmericanEnglishPhrase represents samson american english phrase.
	samsonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pipsqueak"}

	// samsonCanadianFrenchPhrase represents samson canadian french phrase.
	samsonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "scouic"}

	// samsonDutchPhrase represents samson dutch phrase.
	samsonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "piepklein"}

	// samsonFrenchPhrase represents samson french phrase.
	samsonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "scouic"}

	// samsonGermanPhrase represents samson german phrase.
	samsonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "pieeeps"}

	// samsonItalianPhrase represents samson italian phrase.
	samsonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pisquak"}

	// samsonJapanesePhrase represents samson japanese phrase.
	samsonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "チュー"}

	// samsonKoreanPhrase represents samson korean phrase.
	samsonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쪽"}

	// samsonLatinAmericanSpanishPhrase represents samson latin american spanish phrase.
	samsonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "escuác"}

	// samsonRussianPhrase represents samson russian phrase.
	samsonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ревушка"}

	// samsonSimplifiedChinesePhrase represents samson simplified chinese phrase.
	samsonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "吱"}

	// samsonSpanishPhrase represents samson spanish phrase.
	samsonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "flojeras"}

	// samsonTraditionalChinesePhrase represents samson traditional chinese phrase.
	samsonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "吱"}
)

var (
	// samsonPhrase represents samson phrase.
	samsonPhrase = nook.Languages{
		language.AmericanEnglish:      samsonAmericanEnglishPhrase,
		language.CanadianFrench:       samsonCanadianFrenchPhrase,
		language.Dutch:                samsonDutchPhrase,
		language.French:               samsonFrenchPhrase,
		language.German:               samsonGermanPhrase,
		language.Italian:              samsonItalianPhrase,
		language.Japanese:             samsonJapanesePhrase,
		language.Korean:               samsonKoreanPhrase,
		language.LatinAmericanSpanish: samsonLatinAmericanSpanishPhrase,
		language.Russian:              samsonRussianPhrase,
		language.SimplifiedChinese:    samsonSimplifiedChinesePhrase,
		language.Spanish:              samsonSpanishPhrase,
		language.TraditionalChinese:   samsonTraditionalChinesePhrase}
)

var (
	// Samson represents samson.
	Samson = nook.Villager{
		Character:   samsonCharacter,
		Personality: personality.Jock,
		Phrase:      samsonPhrase}
)
