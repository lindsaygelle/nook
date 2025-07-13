package kangaroo

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
	// astridBirthday represents astrid birthday.
	astridBirthday = nook.Birthday{
		Day:   8,
		Month: time.September}
)

var (
	// astridCode represents astrid code.
	astridCode = nook.Code{
		Value: "kgr05"}
)

var (
	// astridAmericanEnglishName represents astrid american english name.
	astridAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Astrid"}

	// astridCanadianFrenchName represents astrid canadian french name.
	astridCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rhona"}

	// astridDutchName represents astrid dutch name.
	astridDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Astrid"}

	// astridFrenchName represents astrid french name.
	astridFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rhona"}

	// astridGermanName represents astrid german name.
	astridGermanName = nook.Name{
		Language: language.German,
		Value:    "Astrid"}

	// astridItalianName represents astrid italian name.
	astridItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Stella"}

	// astridJapaneseName represents astrid japanese name.
	astridJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キッズ"}

	// astridKoreanName represents astrid korean name.
	astridKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "펑키맘"}

	// astridLatinAmericanSpanishName represents astrid latin american spanish name.
	astridLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Astrid"}

	// astridRussianName represents astrid russian name.
	astridRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Астрид"}

	// astridSimplifiedChineseName represents astrid simplified chinese name.
	astridSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "童儿"}

	// astridSpanishName represents astrid spanish name.
	astridSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Astrid"}

	// astridTraditionalChineseName represents astrid traditional chinese name.
	astridTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "童兒"}
)

var (
	// astridName represents astrid name.
	astridName = nook.Languages{
		language.AmericanEnglish:      astridAmericanEnglishName,
		language.CanadianFrench:       astridCanadianFrenchName,
		language.Dutch:                astridDutchName,
		language.French:               astridFrenchName,
		language.German:               astridGermanName,
		language.Italian:              astridItalianName,
		language.Japanese:             astridJapaneseName,
		language.Korean:               astridKoreanName,
		language.LatinAmericanSpanish: astridLatinAmericanSpanishName,
		language.Russian:              astridRussianName,
		language.SimplifiedChinese:    astridSimplifiedChineseName,
		language.Spanish:              astridSpanishName,
		language.TraditionalChinese:   astridTraditionalChineseName}
)

var (
	// astridCharacter represents astrid character.
	astridCharacter = nook.Character{
		Animal:   animal.Kangaroo,
		Birthday: astridBirthday,
		Code:     astridCode,
		Key:      character.Astrid,
		Gender:   gender.Female,
		Name:     astridName,
		Special:  false}
)

var (
	// astridAmericanEnglishPhrase represents astrid american english phrase.
	astridAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "my pet"}

	// astridCanadianFrenchPhrase represents astrid canadian french phrase.
	astridCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon étoile"}

	// astridDutchPhrase represents astrid dutch phrase.
	astridDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "troetel"}

	// astridFrenchPhrase represents astrid french phrase.
	astridFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon étoile"}

	// astridGermanPhrase represents astrid german phrase.
	astridGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schmusi"}

	// astridItalianPhrase represents astrid italian phrase.
	astridItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "fagotto"}

	// astridJapanesePhrase represents astrid japanese phrase.
	astridJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だパンク"}

	// astridKoreanPhrase represents astrid korean phrase.
	astridKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "뭘봐"}

	// astridLatinAmericanSpanishPhrase represents astrid latin american spanish phrase.
	astridLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "brillabrí"}

	// astridRussianPhrase represents astrid russian phrase.
	astridRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "лапочка"}

	// astridSimplifiedChinesePhrase represents astrid simplified chinese phrase.
	astridSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "庞克"}

	// astridSpanishPhrase represents astrid spanish phrase.
	astridSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mascotita"}

	// astridTraditionalChinesePhrase represents astrid traditional chinese phrase.
	astridTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "龐克"}
)

var (
	// astridPhrase represents astrid phrase.
	astridPhrase = nook.Languages{
		language.AmericanEnglish:      astridAmericanEnglishPhrase,
		language.CanadianFrench:       astridCanadianFrenchPhrase,
		language.Dutch:                astridDutchPhrase,
		language.French:               astridFrenchPhrase,
		language.German:               astridGermanPhrase,
		language.Italian:              astridItalianPhrase,
		language.Japanese:             astridJapanesePhrase,
		language.Korean:               astridKoreanPhrase,
		language.LatinAmericanSpanish: astridLatinAmericanSpanishPhrase,
		language.Russian:              astridRussianPhrase,
		language.SimplifiedChinese:    astridSimplifiedChinesePhrase,
		language.Spanish:              astridSpanishPhrase,
		language.TraditionalChinese:   astridTraditionalChinesePhrase}
)

var (
	// Astrid represents astrid.
	Astrid = nook.Villager{
		Character:   astridCharacter,
		Personality: personality.Snooty,
		Phrase:      astridPhrase}
)
