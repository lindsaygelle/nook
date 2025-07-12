package cat

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
	// mitziBirthday represents mitzi birthday.
	mitziBirthday = nook.Birthday{
		Day:   25,
		Month: time.September}
)

var (
	// mitziCode represents mitzi code.
	mitziCode = nook.Code{
		Value: "cat01"}
)

var (
	// mitziAmericanEnglishName represents mitzi american english name.
	mitziAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mitzi"}

	// mitziCanadianFrenchName represents mitzi canadian french name.
	mitziCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mistigri"}

	// mitziDutchName represents mitzi dutch name.
	mitziDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Mitzi"}

	// mitziFrenchName represents mitzi french name.
	mitziFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mistigri"}

	// mitziGermanName represents mitzi german name.
	mitziGermanName = nook.Name{
		Language: language.German,
		Value:    "Miezi"}

	// mitziItalianName represents mitzi italian name.
	mitziItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Zampetta"}

	// mitziJapaneseName represents mitzi japanese name.
	mitziJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マール"}

	// mitziKoreanName represents mitzi korean name.
	mitziKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마르"}

	// mitziLatinAmericanSpanishName represents mitzi latin american spanish name.
	mitziLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Misi"}

	// mitziRussianName represents mitzi russian name.
	mitziRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Митци"}

	// mitziSimplifiedChineseName represents mitzi simplified chinese name.
	mitziSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "圆圆"}

	// mitziSpanishName represents mitzi spanish name.
	mitziSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Misi"}

	// mitziTraditionalChineseName represents mitzi traditional chinese name.
	mitziTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "圓圓"}
)

var (
	// mitziName represents mitzi name.
	mitziName = nook.Languages{
		language.AmericanEnglish:      mitziAmericanEnglishName,
		language.CanadianFrench:       mitziCanadianFrenchName,
		language.Dutch:                mitziDutchName,
		language.French:               mitziFrenchName,
		language.German:               mitziGermanName,
		language.Italian:              mitziItalianName,
		language.Japanese:             mitziJapaneseName,
		language.Korean:               mitziKoreanName,
		language.LatinAmericanSpanish: mitziLatinAmericanSpanishName,
		language.Russian:              mitziRussianName,
		language.SimplifiedChinese:    mitziSimplifiedChineseName,
		language.Spanish:              mitziSpanishName,
		language.TraditionalChinese:   mitziTraditionalChineseName}
)

var (
	// mitziCharacter represents mitzi character.
	mitziCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: mitziBirthday,
		Code:     mitziCode,
		Key:      character.Mitzi,
		Gender:   gender.Female,
		Name:     mitziName,
		Special:  false}
)

var (
	// mitziAmericanEnglishPhrase represents mitzi american english phrase.
	mitziAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "mew"}

	// mitziCanadianFrenchPhrase represents mitzi canadian french phrase.
	mitziCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "meeeeeh"}

	// mitziDutchPhrase represents mitzi dutch phrase.
	mitziDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mieuw"}

	// mitziFrenchPhrase represents mitzi french phrase.
	mitziFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "meeeeeh"}

	// mitziGermanPhrase represents mitzi german phrase.
	mitziGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "muh"}

	// mitziItalianPhrase represents mitzi italian phrase.
	mitziItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "miao"}

	// mitziJapanesePhrase represents mitzi japanese phrase.
	mitziJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ニャー"}

	// mitziKoreanPhrase represents mitzi korean phrase.
	mitziKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "야～옹"}

	// mitziLatinAmericanSpanishPhrase represents mitzi latin american spanish phrase.
	mitziLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "misi-misi"}

	// mitziRussianPhrase represents mitzi russian phrase.
	mitziRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мяу-мяу"}

	// mitziSimplifiedChinesePhrase represents mitzi simplified chinese phrase.
	mitziSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喵"}

	// mitziSpanishPhrase represents mitzi spanish phrase.
	mitziSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "misi-misi"}

	// mitziTraditionalChinesePhrase represents mitzi traditional chinese phrase.
	mitziTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喵"}
)

var (
	// mitziPhrase represents mitzi phrase.
	mitziPhrase = nook.Languages{
		language.AmericanEnglish:      mitziAmericanEnglishPhrase,
		language.CanadianFrench:       mitziCanadianFrenchPhrase,
		language.Dutch:                mitziDutchPhrase,
		language.French:               mitziFrenchPhrase,
		language.German:               mitziGermanPhrase,
		language.Italian:              mitziItalianPhrase,
		language.Japanese:             mitziJapanesePhrase,
		language.Korean:               mitziKoreanPhrase,
		language.LatinAmericanSpanish: mitziLatinAmericanSpanishPhrase,
		language.Russian:              mitziRussianPhrase,
		language.SimplifiedChinese:    mitziSimplifiedChinesePhrase,
		language.Spanish:              mitziSpanishPhrase,
		language.TraditionalChinese:   mitziTraditionalChinesePhrase}
)

var (
	// Mitzi represents mitzi.
	Mitzi = nook.Villager{
		Character:   mitziCharacter,
		Personality: personality.Normal,
		Phrase:      mitziPhrase}
)
