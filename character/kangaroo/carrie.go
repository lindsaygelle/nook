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
	// carrieBirthday represents carrie birthday.
	carrieBirthday = nook.Birthday{
		Day:   5,
		Month: time.December}
)

var (
	// carrieCode represents carrie code.
	carrieCode = nook.Code{
		Value: "kgr02"}
)

var (
	// carrieAmericanEnglishName represents carrie american english name.
	carrieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Carrie"}

	// carrieCanadianFrenchName represents carrie canadian french name.
	carrieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kanga"}

	// carrieDutchName represents carrie dutch name.
	carrieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Carrie"}

	// carrieFrenchName represents carrie french name.
	carrieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kanga"}

	// carrieGermanName represents carrie german name.
	carrieGermanName = nook.Name{
		Language: language.German,
		Value:    "Carola"}

	// carrieItalianName represents carrie italian name.
	carrieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Taschina"}

	// carrieJapaneseName represents carrie japanese name.
	carrieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マミィ"}

	// carrieKoreanName represents carrie korean name.
	carrieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마미"}

	// carrieLatinAmericanSpanishName represents carrie latin american spanish name.
	carrieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lola"}

	// carrieRussianName represents carrie russian name.
	carrieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кэрри"}

	// carrieSimplifiedChineseName represents carrie simplified chinese name.
	carrieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "妈咪"}

	// carrieSpanishName represents carrie spanish name.
	carrieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lola"}

	// carrieTraditionalChineseName represents carrie traditional chinese name.
	carrieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "媽咪"}
)

var (
	// carrieName represents carrie name.
	carrieName = nook.Languages{
		language.AmericanEnglish:      carrieAmericanEnglishName,
		language.CanadianFrench:       carrieCanadianFrenchName,
		language.Dutch:                carrieDutchName,
		language.French:               carrieFrenchName,
		language.German:               carrieGermanName,
		language.Italian:              carrieItalianName,
		language.Japanese:             carrieJapaneseName,
		language.Korean:               carrieKoreanName,
		language.LatinAmericanSpanish: carrieLatinAmericanSpanishName,
		language.Russian:              carrieRussianName,
		language.SimplifiedChinese:    carrieSimplifiedChineseName,
		language.Spanish:              carrieSpanishName,
		language.TraditionalChinese:   carrieTraditionalChineseName}
)

var (
	// carrieCharacter represents carrie character.
	carrieCharacter = nook.Character{
		Animal:   animal.Kangaroo,
		Birthday: carrieBirthday,
		Code:     carrieCode,
		Key:      character.Carrie,
		Gender:   gender.Female,
		Name:     carrieName,
		Special:  false}
)

var (
	// carrieAmericanEnglishPhrase represents carrie american english phrase.
	carrieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "little one"}

	// carrieCanadianFrenchPhrase represents carrie canadian french phrase.
	carrieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon pit"}

	// carrieDutchPhrase represents carrie dutch phrase.
	carrieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kleine"}

	// carrieFrenchPhrase represents carrie french phrase.
	carrieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "poupon"}

	// carrieGermanPhrase represents carrie german phrase.
	carrieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "beutel"}

	// carrieItalianPhrase represents carrie italian phrase.
	carrieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gnegné"}

	// carrieJapanesePhrase represents carrie japanese phrase.
	carrieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だフン"}

	// carrieKoreanPhrase represents carrie korean phrase.
	carrieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "오롤로"}

	// carrieLatinAmericanSpanishPhrase represents carrie latin american spanish phrase.
	carrieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "arrorró"}

	// carrieRussianPhrase represents carrie russian phrase.
	carrieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "дитятко"}

	// carrieSimplifiedChinesePhrase represents carrie simplified chinese phrase.
	carrieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "亲亲"}

	// carrieSpanishPhrase represents carrie spanish phrase.
	carrieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bebé"}

	// carrieTraditionalChinesePhrase represents carrie traditional chinese phrase.
	carrieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "親親"}
)

var (
	// carriePhrase represents carrie phrase.
	carriePhrase = nook.Languages{
		language.AmericanEnglish:      carrieAmericanEnglishPhrase,
		language.CanadianFrench:       carrieCanadianFrenchPhrase,
		language.Dutch:                carrieDutchPhrase,
		language.French:               carrieFrenchPhrase,
		language.German:               carrieGermanPhrase,
		language.Italian:              carrieItalianPhrase,
		language.Japanese:             carrieJapanesePhrase,
		language.Korean:               carrieKoreanPhrase,
		language.LatinAmericanSpanish: carrieLatinAmericanSpanishPhrase,
		language.Russian:              carrieRussianPhrase,
		language.SimplifiedChinese:    carrieSimplifiedChinesePhrase,
		language.Spanish:              carrieSpanishPhrase,
		language.TraditionalChinese:   carrieTraditionalChinesePhrase}
)

var (
	// Carrie represents carrie.
	Carrie = nook.Villager{
		Character:   carrieCharacter,
		Personality: personality.Normal,
		Phrase:      carriePhrase}
)
