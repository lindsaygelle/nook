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
	// anicottiBirthday represents anicotti birthday.
	anicottiBirthday = nook.Birthday{
		Day:   24,
		Month: time.February}
)

var (
	// anicottiCode represents anicotti code.
	anicottiCode = nook.Code{
		Value: "mus10"}
)

var (
	// anicottiAmericanEnglishName represents anicotti american english name.
	anicottiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Anicotti"}

	// anicottiCanadianFrenchName represents anicotti canadian french name.
	anicottiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Annie"}

	// anicottiDutchName represents anicotti dutch name.
	anicottiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Anicotti"}

	// anicottiFrenchName represents anicotti french name.
	anicottiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Annie"}

	// anicottiGermanName represents anicotti german name.
	anicottiGermanName = nook.Name{
		Language: language.German,
		Value:    "Eva"}

	// anicottiItalianName represents anicotti italian name.
	anicottiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Squitta"}

	// anicottiJapaneseName represents anicotti japanese name.
	anicottiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ラザニア"}

	// anicottiKoreanName represents anicotti korean name.
	anicottiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "라자냐"}

	// anicottiLatinAmericanSpanishName represents anicotti latin american spanish name.
	anicottiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Clorinda"}

	// anicottiRussianName represents anicotti russian name.
	anicottiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Аникотти"}

	// anicottiSimplifiedChineseName represents anicotti simplified chinese name.
	anicottiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗萱儿"}

	// anicottiSpanishName represents anicotti spanish name.
	anicottiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Clorinda"}

	// anicottiTraditionalChineseName represents anicotti traditional chinese name.
	anicottiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羅萱兒"}
)

var (
	// anicottiName represents anicotti name.
	anicottiName = nook.Languages{
		language.AmericanEnglish:      anicottiAmericanEnglishName,
		language.CanadianFrench:       anicottiCanadianFrenchName,
		language.Dutch:                anicottiDutchName,
		language.French:               anicottiFrenchName,
		language.German:               anicottiGermanName,
		language.Italian:              anicottiItalianName,
		language.Japanese:             anicottiJapaneseName,
		language.Korean:               anicottiKoreanName,
		language.LatinAmericanSpanish: anicottiLatinAmericanSpanishName,
		language.Russian:              anicottiRussianName,
		language.SimplifiedChinese:    anicottiSimplifiedChineseName,
		language.Spanish:              anicottiSpanishName,
		language.TraditionalChinese:   anicottiTraditionalChineseName}
)

var (
	// anicottiCharacter represents anicotti character.
	anicottiCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: anicottiBirthday,
		Code:     anicottiCode,
		Key:      character.Anicotti,
		Gender:   gender.Female,
		Name:     anicottiName,
		Special:  false}
)

var (
	// anicottiAmericanEnglishPhrase represents anicotti american english phrase.
	anicottiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cannoli"}

	// anicottiCanadianFrenchPhrase represents anicotti canadian french phrase.
	anicottiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "dis donc"}

	// anicottiDutchPhrase represents anicotti dutch phrase.
	anicottiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "cannoli"}

	// anicottiFrenchPhrase represents anicotti french phrase.
	anicottiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "dis donc"}

	// anicottiGermanPhrase represents anicotti german phrase.
	anicottiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "piepser"}

	// anicottiItalianPhrase represents anicotti italian phrase.
	anicottiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cippoli"}

	// anicottiJapanesePhrase represents anicotti japanese phrase.
	anicottiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ルンルン"}

	// anicottiKoreanPhrase represents anicotti korean phrase.
	anicottiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "룰룰"}

	// anicottiLatinAmericanSpanishPhrase represents anicotti latin american spanish phrase.
	anicottiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cloricló"}

	// anicottiRussianPhrase represents anicotti russian phrase.
	anicottiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "канноли"}

	// anicottiSimplifiedChinesePhrase represents anicotti simplified chinese phrase.
	anicottiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "开心"}

	// anicottiSpanishPhrase represents anicotti spanish phrase.
	anicottiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cloricló"}

	// anicottiTraditionalChinesePhrase represents anicotti traditional chinese phrase.
	anicottiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "開心"}
)

var (
	// anicottiPhrase represents anicotti phrase.
	anicottiPhrase = nook.Languages{
		language.AmericanEnglish:      anicottiAmericanEnglishPhrase,
		language.CanadianFrench:       anicottiCanadianFrenchPhrase,
		language.Dutch:                anicottiDutchPhrase,
		language.French:               anicottiFrenchPhrase,
		language.German:               anicottiGermanPhrase,
		language.Italian:              anicottiItalianPhrase,
		language.Japanese:             anicottiJapanesePhrase,
		language.Korean:               anicottiKoreanPhrase,
		language.LatinAmericanSpanish: anicottiLatinAmericanSpanishPhrase,
		language.Russian:              anicottiRussianPhrase,
		language.SimplifiedChinese:    anicottiSimplifiedChinesePhrase,
		language.Spanish:              anicottiSpanishPhrase,
		language.TraditionalChinese:   anicottiTraditionalChinesePhrase}
)

var (
	// Anicotti represents anicotti.
	Anicotti = nook.Villager{
		Character:   anicottiCharacter,
		Personality: personality.Peppy,
		Phrase:      anicottiPhrase}
)
