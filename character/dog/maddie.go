package dog

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
	// maddieBirthday represents maddie birthday.
	maddieBirthday = nook.Birthday{
		Day:   11,
		Month: time.January}
)

var (
	// maddieCode represents maddie code.
	maddieCode = nook.Code{
		Value: "dog09"}
)

var (
	// maddieAmericanEnglishName represents maddie american english name.
	maddieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Maddie"}

	// maddieCanadianFrenchName represents maddie canadian french name.
	maddieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Olympe"}

	// maddieDutchName represents maddie dutch name.
	maddieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Maddie"}

	// maddieFrenchName represents maddie french name.
	maddieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Olympe"}

	// maddieGermanName represents maddie german name.
	maddieGermanName = nook.Name{
		Language: language.German,
		Value:    "Agnes"}

	// maddieItalianName represents maddie italian name.
	maddieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cristina"}

	// maddieJapaneseName represents maddie japanese name.
	maddieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マロン"}

	// maddieKoreanName represents maddie korean name.
	maddieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마롱"}

	// maddieLatinAmericanSpanishName represents maddie latin american spanish name.
	maddieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Martina"}

	// maddieRussianName represents maddie russian name.
	maddieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мэдди"}

	// maddieSimplifiedChineseName represents maddie simplified chinese name.
	maddieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "麻蓉"}

	// maddieSpanishName represents maddie spanish name.
	maddieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Martina"}

	// maddieTraditionalChineseName represents maddie traditional chinese name.
	maddieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "麻蓉"}
)

var (
	// maddieName represents maddie name.
	maddieName = nook.Languages{
		language.AmericanEnglish:      maddieAmericanEnglishName,
		language.CanadianFrench:       maddieCanadianFrenchName,
		language.Dutch:                maddieDutchName,
		language.French:               maddieFrenchName,
		language.German:               maddieGermanName,
		language.Italian:              maddieItalianName,
		language.Japanese:             maddieJapaneseName,
		language.Korean:               maddieKoreanName,
		language.LatinAmericanSpanish: maddieLatinAmericanSpanishName,
		language.Russian:              maddieRussianName,
		language.SimplifiedChinese:    maddieSimplifiedChineseName,
		language.Spanish:              maddieSpanishName,
		language.TraditionalChinese:   maddieTraditionalChineseName}
)

var (
	// maddieCharacter represents maddie character.
	maddieCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: maddieBirthday,
		Code:     maddieCode,
		Key:      character.Maddie,
		Gender:   gender.Female,
		Name:     maddieName,
		Special:  false}
)

var (
	// maddieAmericanEnglishPhrase represents maddie american english phrase.
	maddieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "yippee"}

	// maddieCanadianFrenchPhrase represents maddie canadian french phrase.
	maddieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "youpi"}

	// maddieDutchPhrase represents maddie dutch phrase.
	maddieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "jippie"}

	// maddieFrenchPhrase represents maddie french phrase.
	maddieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "youp"}

	// maddieGermanPhrase represents maddie german phrase.
	maddieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hurra"}

	// maddieItalianPhrase represents maddie italian phrase.
	maddieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "yuuuppii"}

	// maddieJapanesePhrase represents maddie japanese phrase.
	maddieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ワンツー"}

	// maddieKoreanPhrase represents maddie korean phrase.
	maddieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "달코미"}

	// maddieLatinAmericanSpanishPhrase represents maddie latin american spanish phrase.
	maddieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "reguau"}

	// maddieRussianPhrase represents maddie russian phrase.
	maddieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ур-ра"}

	// maddieSimplifiedChinesePhrase represents maddie simplified chinese phrase.
	maddieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "一二"}

	// maddieSpanishPhrase represents maddie spanish phrase.
	maddieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chupi"}

	// maddieTraditionalChinesePhrase represents maddie traditional chinese phrase.
	maddieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "一二"}
)

var (
	// maddiePhrase represents maddie phrase.
	maddiePhrase = nook.Languages{
		language.AmericanEnglish:      maddieAmericanEnglishPhrase,
		language.CanadianFrench:       maddieCanadianFrenchPhrase,
		language.Dutch:                maddieDutchPhrase,
		language.French:               maddieFrenchPhrase,
		language.German:               maddieGermanPhrase,
		language.Italian:              maddieItalianPhrase,
		language.Japanese:             maddieJapanesePhrase,
		language.Korean:               maddieKoreanPhrase,
		language.LatinAmericanSpanish: maddieLatinAmericanSpanishPhrase,
		language.Russian:              maddieRussianPhrase,
		language.SimplifiedChinese:    maddieSimplifiedChinesePhrase,
		language.Spanish:              maddieSpanishPhrase,
		language.TraditionalChinese:   maddieTraditionalChinesePhrase}
)

var (
	// Maddie represents maddie.
	Maddie = nook.Villager{
		Character:   maddieCharacter,
		Personality: personality.Peppy,
		Phrase:      maddiePhrase}
)
