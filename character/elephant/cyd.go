package elephant

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
	// cydBirthday represents cyd birthday.
	cydBirthday = nook.Birthday{
		Day:   9,
		Month: time.June}
)

var (
	// cydCode represents cyd code.
	cydCode = nook.Code{
		Value: "elp12"}
)

var (
	// cydAmericanEnglishName represents cyd american english name.
	cydAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cyd"}

	// cydCanadianFrenchName represents cyd canadian french name.
	cydCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Punk"}

	// cydDutchName represents cyd dutch name.
	cydDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cyd"}

	// cydFrenchName represents cyd french name.
	cydFrenchName = nook.Name{
		Language: language.French,
		Value:    "Punk"}

	// cydGermanName represents cyd german name.
	cydGermanName = nook.Name{
		Language: language.German,
		Value:    "Sid"}

	// cydItalianName represents cyd italian name.
	cydItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sid"}

	// cydJapaneseName represents cyd japanese name.
	cydJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パンクス"}

	// cydKoreanName represents cyd korean name.
	cydKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "펑크스"}

	// cydLatinAmericanSpanishName represents cyd latin american spanish name.
	cydLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ramón"}

	// cydRussianName represents cyd russian name.
	cydRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Сид"}

	// cydSimplifiedChineseName represents cyd simplified chinese name.
	cydSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "庞克斯"}

	// cydSpanishName represents cyd spanish name.
	cydSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ramón"}

	// cydTraditionalChineseName represents cyd traditional chinese name.
	cydTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "龐克斯"}
)

var (
	// cydName represents cyd name.
	cydName = nook.Languages{
		language.AmericanEnglish:      cydAmericanEnglishName,
		language.CanadianFrench:       cydCanadianFrenchName,
		language.Dutch:                cydDutchName,
		language.French:               cydFrenchName,
		language.German:               cydGermanName,
		language.Italian:              cydItalianName,
		language.Japanese:             cydJapaneseName,
		language.Korean:               cydKoreanName,
		language.LatinAmericanSpanish: cydLatinAmericanSpanishName,
		language.Russian:              cydRussianName,
		language.SimplifiedChinese:    cydSimplifiedChineseName,
		language.Spanish:              cydSpanishName,
		language.TraditionalChinese:   cydTraditionalChineseName}
)

var (
	// cydCharacter represents cyd character.
	cydCharacter = nook.Character{
		Animal:   animal.Elephant,
		Birthday: cydBirthday,
		Code:     cydCode,
		Key:      character.Cyd,
		Gender:   gender.Male,
		Name:     cydName,
		Special:  false}
)

var (
	// cydAmericanEnglishPhrase represents cyd american english phrase.
	cydAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "rockin'"}

	// cydCanadianFrenchPhrase represents cyd canadian french phrase.
	cydCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "barda"}

	// cydDutchPhrase represents cyd dutch phrase.
	cydDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mieters"}

	// cydFrenchPhrase represents cyd french phrase.
	cydFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bastringue"}

	// cydGermanPhrase represents cyd german phrase.
	cydGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hau"}

	// cydItalianPhrase represents cyd italian phrase.
	cydItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "rocking"}

	// cydJapanesePhrase represents cyd japanese phrase.
	cydJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウス"}

	// cydKoreanPhrase represents cyd korean phrase.
	cydKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "피스"}

	// cydLatinAmericanSpanishPhrase represents cyd latin american spanish phrase.
	cydLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "rocanrol"}

	// cydRussianPhrase represents cyd russian phrase.
	cydRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "рок-н-ролл"}

	// cydSimplifiedChinesePhrase represents cyd simplified chinese phrase.
	cydSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "摇滚"}

	// cydSpanishPhrase represents cyd spanish phrase.
	cydSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "rocanrol"}

	// cydTraditionalChinesePhrase represents cyd traditional chinese phrase.
	cydTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "搖滾"}
)

var (
	// cydPhrase represents cyd phrase.
	cydPhrase = nook.Languages{
		language.AmericanEnglish:      cydAmericanEnglishPhrase,
		language.CanadianFrench:       cydCanadianFrenchPhrase,
		language.Dutch:                cydDutchPhrase,
		language.French:               cydFrenchPhrase,
		language.German:               cydGermanPhrase,
		language.Italian:              cydItalianPhrase,
		language.Japanese:             cydJapanesePhrase,
		language.Korean:               cydKoreanPhrase,
		language.LatinAmericanSpanish: cydLatinAmericanSpanishPhrase,
		language.Russian:              cydRussianPhrase,
		language.SimplifiedChinese:    cydSimplifiedChinesePhrase,
		language.Spanish:              cydSpanishPhrase,
		language.TraditionalChinese:   cydTraditionalChinesePhrase}
)

var (
	// Cyd represents cyd.
	Cyd = nook.Villager{
		Character:   cydCharacter,
		Personality: personality.Cranky,
		Phrase:      cydPhrase}
)
