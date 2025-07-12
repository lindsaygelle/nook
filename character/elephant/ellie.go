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
	// ellieBirthday represents ellie birthday.
	ellieBirthday = nook.Birthday{
		Day:   12,
		Month: time.May}
)

var (
	// ellieCode represents ellie code.
	ellieCode = nook.Code{
		Value: "elp07"}
)

var (
	// ellieAmericanEnglishName represents ellie american english name.
	ellieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ellie"}

	// ellieCanadianFrenchName represents ellie canadian french name.
	ellieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ella"}

	// ellieDutchName represents ellie dutch name.
	ellieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ellie"}

	// ellieFrenchName represents ellie french name.
	ellieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ella"}

	// ellieGermanName represents ellie german name.
	ellieGermanName = nook.Name{
		Language: language.German,
		Value:    "Elfi"}

	// ellieItalianName represents ellie italian name.
	ellieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Elly"}

	// ellieJapaneseName represents ellie japanese name.
	ellieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エクレア"}

	// ellieKoreanName represents ellie korean name.
	ellieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "에끌레르"}

	// ellieLatinAmericanSpanishName represents ellie latin american spanish name.
	ellieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Eli"}

	// ellieRussianName represents ellie russian name.
	ellieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Элла"}

	// ellieSimplifiedChineseName represents ellie simplified chinese name.
	ellieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "泡芙"}

	// ellieSpanishName represents ellie spanish name.
	ellieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Eli"}

	// ellieTraditionalChineseName represents ellie traditional chinese name.
	ellieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "泡芙"}
)

var (
	// ellieName represents ellie name.
	ellieName = nook.Languages{
		language.AmericanEnglish:      ellieAmericanEnglishName,
		language.CanadianFrench:       ellieCanadianFrenchName,
		language.Dutch:                ellieDutchName,
		language.French:               ellieFrenchName,
		language.German:               ellieGermanName,
		language.Italian:              ellieItalianName,
		language.Japanese:             ellieJapaneseName,
		language.Korean:               ellieKoreanName,
		language.LatinAmericanSpanish: ellieLatinAmericanSpanishName,
		language.Russian:              ellieRussianName,
		language.SimplifiedChinese:    ellieSimplifiedChineseName,
		language.Spanish:              ellieSpanishName,
		language.TraditionalChinese:   ellieTraditionalChineseName}
)

var (
	// ellieCharacter represents ellie character.
	ellieCharacter = nook.Character{
		Animal:   animal.Elephant,
		Birthday: ellieBirthday,
		Code:     ellieCode,
		Key:      character.Ellie,
		Gender:   gender.Female,
		Name:     ellieName,
		Special:  false}
)

var (
	// ellieAmericanEnglishPhrase represents ellie american english phrase.
	ellieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "wee one"}

	// ellieCanadianFrenchPhrase represents ellie canadian french phrase.
	ellieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "sacrelotte"}

	// ellieDutchPhrase represents ellie dutch phrase.
	ellieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "dreumes"}

	// ellieFrenchPhrase represents ellie french phrase.
	ellieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "sacrelotte"}

	// ellieGermanPhrase represents ellie german phrase.
	ellieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "polter"}

	// ellieItalianPhrase represents ellie italian phrase.
	ellieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "caromé"}

	// ellieJapanesePhrase represents ellie japanese phrase.
	ellieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ララン"}

	// ellieKoreanPhrase represents ellie korean phrase.
	ellieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "트랄라"}

	// ellieLatinAmericanSpanishPhrase represents ellie latin american spanish phrase.
	ellieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fi-fiú"}

	// ellieRussianPhrase represents ellie russian phrase.
	ellieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кроха"}

	// ellieSimplifiedChinesePhrase represents ellie simplified chinese phrase.
	ellieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啦啷"}

	// ellieSpanishPhrase represents ellie spanish phrase.
	ellieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "trompi"}

	// ellieTraditionalChinesePhrase represents ellie traditional chinese phrase.
	ellieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啦啷"}
)

var (
	// elliePhrase represents ellie phrase.
	elliePhrase = nook.Languages{
		language.AmericanEnglish:      ellieAmericanEnglishPhrase,
		language.CanadianFrench:       ellieCanadianFrenchPhrase,
		language.Dutch:                ellieDutchPhrase,
		language.French:               ellieFrenchPhrase,
		language.German:               ellieGermanPhrase,
		language.Italian:              ellieItalianPhrase,
		language.Japanese:             ellieJapanesePhrase,
		language.Korean:               ellieKoreanPhrase,
		language.LatinAmericanSpanish: ellieLatinAmericanSpanishPhrase,
		language.Russian:              ellieRussianPhrase,
		language.SimplifiedChinese:    ellieSimplifiedChinesePhrase,
		language.Spanish:              ellieSpanishPhrase,
		language.TraditionalChinese:   ellieTraditionalChinesePhrase}
)

var (
	// Ellie represents ellie.
	Ellie = nook.Villager{
		Character:   ellieCharacter,
		Personality: personality.Normal,
		Phrase:      elliePhrase}
)
