package monkey

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
	// simonBirthday represents simon birthday.
	simonBirthday = nook.Birthday{
		Day:   19,
		Month: time.January}
)

var (
	// simonCode represents simon code.
	simonCode = nook.Code{
		Value: "mnk02"}
)

var (
	// simonAmericanEnglishName represents simon american english name.
	simonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Simon"}

	// simonCanadianFrenchName represents simon canadian french name.
	simonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Simon"}

	// simonDutchName represents simon dutch name.
	simonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Simon"}

	// simonFrenchName represents simon french name.
	simonFrenchName = nook.Name{
		Language: language.French,
		Value:    "Simon"}

	// simonGermanName represents simon german name.
	simonGermanName = nook.Name{
		Language: language.German,
		Value:    "Simon"}

	// simonItalianName represents simon italian name.
	simonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Simon"}

	// simonJapaneseName represents simon japanese name.
	simonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エテキチ"}

	// simonKoreanName represents simon korean name.
	simonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "시몬"}

	// simonLatinAmericanSpanishName represents simon latin american spanish name.
	simonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Simón"}

	// simonRussianName represents simon russian name.
	simonRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Саймон"}

	// simonSimplifiedChineseName represents simon simplified chinese name.
	simonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "远仁"}

	// simonSpanishName represents simon spanish name.
	simonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Simón"}

	// simonTraditionalChineseName represents simon traditional chinese name.
	simonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "遠仁"}
)

var (
	// simonName represents simon name.
	simonName = nook.Languages{
		language.AmericanEnglish:      simonAmericanEnglishName,
		language.CanadianFrench:       simonCanadianFrenchName,
		language.Dutch:                simonDutchName,
		language.French:               simonFrenchName,
		language.German:               simonGermanName,
		language.Italian:              simonItalianName,
		language.Japanese:             simonJapaneseName,
		language.Korean:               simonKoreanName,
		language.LatinAmericanSpanish: simonLatinAmericanSpanishName,
		language.Russian:              simonRussianName,
		language.SimplifiedChinese:    simonSimplifiedChineseName,
		language.Spanish:              simonSpanishName,
		language.TraditionalChinese:   simonTraditionalChineseName}
)

var (
	// simonCharacter represents simon character.
	simonCharacter = nook.Character{
		Animal:   animal.Monkey,
		Birthday: simonBirthday,
		Code:     simonCode,
		Key:      character.Simon,
		Gender:   gender.Male,
		Name:     simonName,
		Special:  false}
)

var (
	// simonAmericanEnglishPhrase represents simon american english phrase.
	simonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "zzzook"}

	// simonCanadianFrenchPhrase represents simon canadian french phrase.
	simonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "zzziik"}

	// simonDutchPhrase represents simon dutch phrase.
	simonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "zzzoeh oeh"}

	// simonFrenchPhrase represents simon french phrase.
	simonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "zzziik"}

	// simonGermanPhrase represents simon german phrase.
	simonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bannabanna"}

	// simonItalianPhrase represents simon italian phrase.
	simonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "zzzook"}

	// simonJapanesePhrase represents simon japanese phrase.
	simonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でござる"}

	// simonKoreanPhrase represents simon korean phrase.
	simonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "하옵니다"}

	// simonLatinAmericanSpanishPhrase represents simon latin american spanish phrase.
	simonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chimpa"}

	// simonRussianPhrase represents simon russian phrase.
	simonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "дружок"}

	// simonSimplifiedChinesePhrase represents simon simplified chinese phrase.
	simonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "猿猿"}

	// simonSpanishPhrase represents simon spanish phrase.
	simonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chimpa"}

	// simonTraditionalChinesePhrase represents simon traditional chinese phrase.
	simonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "猿猿"}
)

var (
	// simonPhrase represents simon phrase.
	simonPhrase = nook.Languages{
		language.AmericanEnglish:      simonAmericanEnglishPhrase,
		language.CanadianFrench:       simonCanadianFrenchPhrase,
		language.Dutch:                simonDutchPhrase,
		language.French:               simonFrenchPhrase,
		language.German:               simonGermanPhrase,
		language.Italian:              simonItalianPhrase,
		language.Japanese:             simonJapanesePhrase,
		language.Korean:               simonKoreanPhrase,
		language.LatinAmericanSpanish: simonLatinAmericanSpanishPhrase,
		language.Russian:              simonRussianPhrase,
		language.SimplifiedChinese:    simonSimplifiedChinesePhrase,
		language.Spanish:              simonSpanishPhrase,
		language.TraditionalChinese:   simonTraditionalChinesePhrase}
)

var (
	// Simon represents simon.
	Simon = nook.Villager{
		Character:   simonCharacter,
		Personality: personality.Lazy,
		Phrase:      simonPhrase}
)
