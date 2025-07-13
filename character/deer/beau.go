package deer

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
	// beauBirthday represents beau birthday.
	beauBirthday = nook.Birthday{
		Day:   5,
		Month: time.April}
)

var (
	// beauCode represents beau code.
	beauCode = nook.Code{
		Value: "der07"}
)

var (
	// beauAmericanEnglishName represents beau american english name.
	beauAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Beau"}

	// beauCanadianFrenchName represents beau canadian french name.
	beauCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Stefaon"}

	// beauDutchName represents beau dutch name.
	beauDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Beau"}

	// beauFrenchName represents beau french name.
	beauFrenchName = nook.Name{
		Language: language.French,
		Value:    "Stefaon"}

	// beauGermanName represents beau german name.
	beauGermanName = nook.Name{
		Language: language.German,
		Value:    "Martin"}

	// beauItalianName represents beau italian name.
	beauItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gigi"}

	// beauJapaneseName represents beau japanese name.
	beauJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ペーター"}

	// beauKoreanName represents beau korean name.
	beauKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "피터"}

	// beauLatinAmericanSpanishName represents beau latin american spanish name.
	beauLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lope"}

	// beauRussianName represents beau russian name.
	beauRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бью"}

	// beauSimplifiedChineseName represents beau simplified chinese name.
	beauSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "彼得"}

	// beauSpanishName represents beau spanish name.
	beauSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lope"}

	// beauTraditionalChineseName represents beau traditional chinese name.
	beauTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "彼得"}
)

var (
	// beauName represents beau name.
	beauName = nook.Languages{
		language.AmericanEnglish:      beauAmericanEnglishName,
		language.CanadianFrench:       beauCanadianFrenchName,
		language.Dutch:                beauDutchName,
		language.French:               beauFrenchName,
		language.German:               beauGermanName,
		language.Italian:              beauItalianName,
		language.Japanese:             beauJapaneseName,
		language.Korean:               beauKoreanName,
		language.LatinAmericanSpanish: beauLatinAmericanSpanishName,
		language.Russian:              beauRussianName,
		language.SimplifiedChinese:    beauSimplifiedChineseName,
		language.Spanish:              beauSpanishName,
		language.TraditionalChinese:   beauTraditionalChineseName}
)

var (
	// beauCharacter represents beau character.
	beauCharacter = nook.Character{
		Animal:   animal.Deer,
		Birthday: beauBirthday,
		Code:     beauCode,
		Key:      character.Beau,
		Gender:   gender.Male,
		Name:     beauName,
		Special:  false}
)

var (
	// beauAmericanEnglishPhrase represents beau american english phrase.
	beauAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "saltlick"}

	// beauCanadianFrenchPhrase represents beau canadian french phrase.
	beauCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "feufeuille"}

	// beauDutchPhrase represents beau dutch phrase.
	beauDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "zoutsteen"}

	// beauFrenchPhrase represents beau french phrase.
	beauFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "feufeuille"}

	// beauGermanPhrase represents beau german phrase.
	beauGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knospel"}

	// beauItalianPhrase represents beau italian phrase.
	beauItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gnam gnam"}

	// beauJapanesePhrase represents beau japanese phrase.
	beauJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "おろおろ"}

	// beauKoreanPhrase represents beau korean phrase.
	beauKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우왕"}

	// beauLatinAmericanSpanishPhrase represents beau latin american spanish phrase.
	beauLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "babum"}

	// beauRussianPhrase represents beau russian phrase.
	beauRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "соль-соль"}

	// beauSimplifiedChinesePhrase represents beau simplified chinese phrase.
	beauSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "怎么办"}

	// beauSpanishPhrase represents beau spanish phrase.
	beauSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "babum"}

	// beauTraditionalChinesePhrase represents beau traditional chinese phrase.
	beauTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "怎麼辦"}
)

var (
	// beauPhrase represents beau phrase.
	beauPhrase = nook.Languages{
		language.AmericanEnglish:      beauAmericanEnglishPhrase,
		language.CanadianFrench:       beauCanadianFrenchPhrase,
		language.Dutch:                beauDutchPhrase,
		language.French:               beauFrenchPhrase,
		language.German:               beauGermanPhrase,
		language.Italian:              beauItalianPhrase,
		language.Japanese:             beauJapanesePhrase,
		language.Korean:               beauKoreanPhrase,
		language.LatinAmericanSpanish: beauLatinAmericanSpanishPhrase,
		language.Russian:              beauRussianPhrase,
		language.SimplifiedChinese:    beauSimplifiedChinesePhrase,
		language.Spanish:              beauSpanishPhrase,
		language.TraditionalChinese:   beauTraditionalChinesePhrase}
)

var (
	// Beau represents beau.
	Beau = nook.Villager{
		Character:   beauCharacter,
		Personality: personality.Lazy,
		Phrase:      beauPhrase}
)
