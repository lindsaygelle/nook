package chicken

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
	// broffinaBirthday represents broffina birthday.
	broffinaBirthday = nook.Birthday{
		Day:   24,
		Month: time.October}
)

var (
	// broffinaCode represents broffina code.
	broffinaCode = nook.Code{
		Value: "chn12"}
)

var (
	// broffinaAmericanEnglishName represents broffina american english name.
	broffinaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Broffina"}

	// broffinaCanadianFrenchName represents broffina canadian french name.
	broffinaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Jo"}

	// broffinaDutchName represents broffina dutch name.
	broffinaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Broffina"}

	// broffinaFrenchName represents broffina french name.
	broffinaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jo"}

	// broffinaGermanName represents broffina german name.
	broffinaGermanName = nook.Name{
		Language: language.German,
		Value:    "Elfriede"}

	// broffinaItalianName represents broffina italian name.
	broffinaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Concetta"}

	// broffinaJapaneseName represents broffina japanese name.
	broffinaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カサンドラ"}

	// broffinaKoreanName represents broffina korean name.
	broffinaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "히킨"}

	// broffinaLatinAmericanSpanishName represents broffina latin american spanish name.
	broffinaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Brunilda"}

	// broffinaRussianName represents broffina russian name.
	broffinaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Броффина"}

	// broffinaSimplifiedChineseName represents broffina simplified chinese name.
	broffinaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "凯西"}

	// broffinaSpanishName represents broffina spanish name.
	broffinaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Brunilda"}

	// broffinaTraditionalChineseName represents broffina traditional chinese name.
	broffinaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "凱西"}
)

var (
	// broffinaName represents broffina name.
	broffinaName = nook.Languages{
		language.AmericanEnglish:      broffinaAmericanEnglishName,
		language.CanadianFrench:       broffinaCanadianFrenchName,
		language.Dutch:                broffinaDutchName,
		language.French:               broffinaFrenchName,
		language.German:               broffinaGermanName,
		language.Italian:              broffinaItalianName,
		language.Japanese:             broffinaJapaneseName,
		language.Korean:               broffinaKoreanName,
		language.LatinAmericanSpanish: broffinaLatinAmericanSpanishName,
		language.Russian:              broffinaRussianName,
		language.SimplifiedChinese:    broffinaSimplifiedChineseName,
		language.Spanish:              broffinaSpanishName,
		language.TraditionalChinese:   broffinaTraditionalChineseName}
)

var (
	// broffinaCharacter represents broffina character.
	broffinaCharacter = nook.Character{
		Animal:   animal.Chicken,
		Birthday: broffinaBirthday,
		Code:     broffinaCode,
		Key:      character.Broffina,
		Gender:   gender.Female,
		Name:     broffinaName,
		Special:  false}
)

var (
	// broffinaAmericanEnglishPhrase represents broffina american english phrase.
	broffinaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cluckadoo"}

	// broffinaCanadianFrenchPhrase represents broffina canadian french phrase.
	broffinaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon œuf"}

	// broffinaDutchPhrase represents broffina dutch phrase.
	broffinaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kakel"}

	// broffinaFrenchPhrase represents broffina french phrase.
	broffinaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon œuf"}

	// broffinaGermanPhrase represents broffina german phrase.
	broffinaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "eieiei"}

	// broffinaItalianPhrase represents broffina italian phrase.
	broffinaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ovodè"}

	// broffinaJapanesePhrase represents broffina japanese phrase.
	broffinaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ケッコー"}

	// broffinaKoreanPhrase represents broffina korean phrase.
	broffinaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "오싹오싹"}

	// broffinaLatinAmericanSpanishPhrase represents broffina latin american spanish phrase.
	broffinaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "kiriquicó"}

	// broffinaRussianPhrase represents broffina russian phrase.
	broffinaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ки-ки-ки"}

	// broffinaSimplifiedChinesePhrase represents broffina simplified chinese phrase.
	broffinaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咯咯"}

	// broffinaSpanishPhrase represents broffina spanish phrase.
	broffinaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pitas"}

	// broffinaTraditionalChinesePhrase represents broffina traditional chinese phrase.
	broffinaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咯咯"}
)

var (
	// broffinaPhrase represents broffina phrase.
	broffinaPhrase = nook.Languages{
		language.AmericanEnglish:      broffinaAmericanEnglishPhrase,
		language.CanadianFrench:       broffinaCanadianFrenchPhrase,
		language.Dutch:                broffinaDutchPhrase,
		language.French:               broffinaFrenchPhrase,
		language.German:               broffinaGermanPhrase,
		language.Italian:              broffinaItalianPhrase,
		language.Japanese:             broffinaJapanesePhrase,
		language.Korean:               broffinaKoreanPhrase,
		language.LatinAmericanSpanish: broffinaLatinAmericanSpanishPhrase,
		language.Russian:              broffinaRussianPhrase,
		language.SimplifiedChinese:    broffinaSimplifiedChinesePhrase,
		language.Spanish:              broffinaSpanishPhrase,
		language.TraditionalChinese:   broffinaTraditionalChinesePhrase}
)

var (
	// Broffina represents broffina.
	Broffina = nook.Villager{
		Character:   broffinaCharacter,
		Personality: personality.Snooty,
		Phrase:      broffinaPhrase}
)
