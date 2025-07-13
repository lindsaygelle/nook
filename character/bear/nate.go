package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// nateBirthday represents Nate's birthday.
var (
	// nateBirthday represents nate birthday.
	nateBirthday = nook.Birthday{
		Day:   16,
		Month: time.August}
)

// nateCode represents Nate's unique code.
var (
	// nateCode represents nate code.
	nateCode = nook.Code{
		Value: "bea05"}
)

// Different names for Nate in various languages.
var (
	// nateAmericanEnglishName represents nate american english name.
	nateAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Nate"}

	// nateCanadianFrenchName represents nate canadian french name.
	nateCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Nathan"}

	// nateDutchName represents nate dutch name.
	nateDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Nate"}

	// nateFrenchName represents nate french name.
	nateFrenchName = nook.Name{
		Language: language.French,
		Value:    "Nathan"}

	// nateGermanName represents nate german name.
	nateGermanName = nook.Name{
		Language: language.German,
		Value:    "Nathan"}

	// nateItalianName represents nate italian name.
	nateItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gianni"}

	// nateJapaneseName represents nate japanese name.
	nateJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "バッカス"}

	// nateKoreanName represents nate korean name.
	nateKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "박하스"}

	// nateLatinAmericanSpanishName represents nate latin american spanish name.
	nateLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nachete"}

	// nateRussianName represents nate russian name.
	nateRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Нэйт"}

	// nateSimplifiedChineseName represents nate simplified chinese name.
	nateSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "巴克思"}

	// nateSpanishName represents nate spanish name.
	nateSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Nachete"}

	// nateTraditionalChineseName represents nate traditional chinese name.
	nateTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "巴克思"}
)

// nateName represents Nate's name in different languages.
var (
	// nateName represents nate name.
	nateName = nook.Languages{
		language.AmericanEnglish:      nateAmericanEnglishName,
		language.CanadianFrench:       nateCanadianFrenchName,
		language.Dutch:                nateDutchName,
		language.French:               nateFrenchName,
		language.German:               nateGermanName,
		language.Italian:              nateItalianName,
		language.Japanese:             nateJapaneseName,
		language.Korean:               nateKoreanName,
		language.LatinAmericanSpanish: nateLatinAmericanSpanishName,
		language.Russian:              nateRussianName,
		language.SimplifiedChinese:    nateSimplifiedChineseName,
		language.Spanish:              nateSpanishName,
		language.TraditionalChinese:   nateTraditionalChineseName}
)

// nateCharacter represents Nate's character information.
var (
	// nateCharacter represents nate character.
	nateCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: nateBirthday,
		Code:     nateCode,
		Key:      character.Nate,
		Gender:   gender.Male,
		Name:     nateName,
		Special:  false}
)

// Different phrases spoken by Nate in various languages.
var (
	// nateAmericanEnglishPhrase represents nate american english phrase.
	nateAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "yawwwn"}

	// nateCanadianFrenchPhrase represents nate canadian french phrase.
	nateCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "baaîîîllle"}

	// nateDutchPhrase represents nate dutch phrase.
	nateDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "gaaaap"}

	// nateFrenchPhrase represents nate french phrase.
	nateFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "baaîîîllle"}

	// nateGermanPhrase represents nate german phrase.
	nateGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "gäähn"}

	// nateItalianPhrase represents nate italian phrase.
	nateItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "yauuun"}

	// nateJapanesePhrase represents nate japanese phrase.
	nateJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "んだ"}

	// nateKoreanPhrase represents nate korean phrase.
	nateKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "힘내"}

	// nateLatinAmericanSpanishPhrase represents nate latin american spanish phrase.
	nateLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "uuaaaah"}

	// nateRussianPhrase represents nate russian phrase.
	nateRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "спа-а-ать"}

	// nateSimplifiedChinesePhrase represents nate simplified chinese phrase.
	nateSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嗯是"}

	// nateSpanishPhrase represents nate spanish phrase.
	nateSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "uuaaaah"}

	// nateTraditionalChinesePhrase represents nate traditional chinese phrase.
	nateTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗯是"}
)

// natePhrase represents Nate's phrases in different languages.
var (
	// natePhrase represents nate phrase.
	natePhrase = nook.Languages{
		language.AmericanEnglish:      nateAmericanEnglishPhrase,
		language.CanadianFrench:       nateCanadianFrenchPhrase,
		language.Dutch:                nateDutchPhrase,
		language.French:               nateFrenchPhrase,
		language.German:               nateGermanPhrase,
		language.Italian:              nateItalianPhrase,
		language.Japanese:             nateJapanesePhrase,
		language.Korean:               nateKoreanPhrase,
		language.LatinAmericanSpanish: nateLatinAmericanSpanishPhrase,
		language.Russian:              nateRussianPhrase,
		language.SimplifiedChinese:    nateSimplifiedChinesePhrase,
		language.Spanish:              nateSpanishPhrase,
		language.TraditionalChinese:   nateTraditionalChinesePhrase}
)

// Nate represents the character Nate with his complete information.
var (
	// Nate represents nate.
	Nate = nook.Villager{
		Character:   nateCharacter,
		Personality: personality.Lazy,
		Phrase:      natePhrase}
)
