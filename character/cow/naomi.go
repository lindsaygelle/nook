package cow

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
	// naomiBirthday represents naomi birthday.
	naomiBirthday = nook.Birthday{
		Day:   28,
		Month: time.February}
)

var (
	// naomiCode represents naomi code.
	naomiCode = nook.Code{
		Value: "cow07"}
)

var (
	// naomiAmericanEnglishName represents naomi american english name.
	naomiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Naomi"}

	// naomiCanadianFrenchName represents naomi canadian french name.
	naomiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Maiko"}

	// naomiDutchName represents naomi dutch name.
	naomiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Naomi"}

	// naomiFrenchName represents naomi french name.
	naomiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maiko"}

	// naomiGermanName represents naomi german name.
	naomiGermanName = nook.Name{
		Language: language.German,
		Value:    "Jolanda"}

	// naomiItalianName represents naomi italian name.
	naomiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Melina"}

	// naomiJapaneseName represents naomi japanese name.
	naomiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハナコ"}

	// naomiKoreanName represents naomi korean name.
	naomiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "화자"}

	// naomiLatinAmericanSpanishName represents naomi latin american spanish name.
	naomiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Regina"}

	// naomiRussianName represents naomi russian name.
	naomiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Наоми"}

	// naomiSimplifiedChineseName represents naomi simplified chinese name.
	naomiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "玉兰"}

	// naomiSpanishName represents naomi spanish name.
	naomiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Regina"}

	// naomiTraditionalChineseName represents naomi traditional chinese name.
	naomiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "玉蘭"}
)

var (
	// naomiName represents naomi name.
	naomiName = nook.Languages{
		language.AmericanEnglish:      naomiAmericanEnglishName,
		language.CanadianFrench:       naomiCanadianFrenchName,
		language.Dutch:                naomiDutchName,
		language.French:               naomiFrenchName,
		language.German:               naomiGermanName,
		language.Italian:              naomiItalianName,
		language.Japanese:             naomiJapaneseName,
		language.Korean:               naomiKoreanName,
		language.LatinAmericanSpanish: naomiLatinAmericanSpanishName,
		language.Russian:              naomiRussianName,
		language.SimplifiedChinese:    naomiSimplifiedChineseName,
		language.Spanish:              naomiSpanishName,
		language.TraditionalChinese:   naomiTraditionalChineseName}
)

var (
	// naomiCharacter represents naomi character.
	naomiCharacter = nook.Character{
		Animal:   animal.Cow,
		Birthday: naomiBirthday,
		Code:     naomiCode,
		Key:      character.Naomi,
		Gender:   gender.Female,
		Name:     naomiName,
		Special:  false}
)

var (
	// naomiAmericanEnglishPhrase represents naomi american english phrase.
	naomiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "moolah"}

	// naomiCanadianFrenchPhrase represents naomi canadian french phrase.
	naomiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "merengue"}

	// naomiDutchPhrase represents naomi dutch phrase.
	naomiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "loei"}

	// naomiFrenchPhrase represents naomi french phrase.
	naomiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pistouille"}

	// naomiGermanPhrase represents naomi german phrase.
	naomiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "joggi"}

	// naomiItalianPhrase represents naomi italian phrase.
	naomiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "moumou"}

	// naomiJapanesePhrase represents naomi japanese phrase.
	naomiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "セボーン"}

	// naomiKoreanPhrase represents naomi korean phrase.
	naomiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "메르시"}

	// naomiLatinAmericanSpanishPhrase represents naomi latin american spanish phrase.
	naomiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "mumumu"}

	// naomiRussianPhrase represents naomi russian phrase.
	naomiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "муля"}

	// naomiSimplifiedChinesePhrase represents naomi simplified chinese phrase.
	naomiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "这样很好"}

	// naomiSpanishPhrase represents naomi spanish phrase.
	naomiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "merengue"}

	// naomiTraditionalChinesePhrase represents naomi traditional chinese phrase.
	naomiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "這樣很好"}
)

var (
	// naomiPhrase represents naomi phrase.
	naomiPhrase = nook.Languages{
		language.AmericanEnglish:      naomiAmericanEnglishPhrase,
		language.CanadianFrench:       naomiCanadianFrenchPhrase,
		language.Dutch:                naomiDutchPhrase,
		language.French:               naomiFrenchPhrase,
		language.German:               naomiGermanPhrase,
		language.Italian:              naomiItalianPhrase,
		language.Japanese:             naomiJapanesePhrase,
		language.Korean:               naomiKoreanPhrase,
		language.LatinAmericanSpanish: naomiLatinAmericanSpanishPhrase,
		language.Russian:              naomiRussianPhrase,
		language.SimplifiedChinese:    naomiSimplifiedChinesePhrase,
		language.Spanish:              naomiSpanishPhrase,
		language.TraditionalChinese:   naomiTraditionalChinesePhrase}
)

var (
	// Naomi represents naomi.
	Naomi = nook.Villager{
		Character:   naomiCharacter,
		Personality: personality.Snooty,
		Phrase:      naomiPhrase}
)
