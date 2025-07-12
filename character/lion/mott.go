package lion

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
	// mottBirthday represents mott birthday.
	mottBirthday = nook.Birthday{
		Day:   10,
		Month: time.July}
)

var (
	// mottCode represents mott code.
	mottCode = nook.Code{
		Value: "lon06"}
)

var (
	// mottAmericanEnglishName represents mott american english name.
	mottAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mott"}

	// mottCanadianFrenchName represents mott canadian french name.
	mottCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Aimé"}

	// mottDutchName represents mott dutch name.
	mottDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Mott"}

	// mottFrenchName represents mott french name.
	mottFrenchName = nook.Name{
		Language: language.French,
		Value:    "Aimé"}

	// mottGermanName represents mott german name.
	mottGermanName = nook.Name{
		Language: language.German,
		Value:    "Leonhard"}

	// mottItalianName represents mott italian name.
	mottItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Leopold"}

	// mottJapaneseName represents mott japanese name.
	mottJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リック"}

	// mottKoreanName represents mott korean name.
	mottKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "릭"}

	// mottLatinAmericanSpanishName represents mott latin american spanish name.
	mottLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jones"}

	// mottRussianName represents mott russian name.
	mottRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мотт"}

	// mottSimplifiedChineseName represents mott simplified chinese name.
	mottSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "李克"}

	// mottSpanishName represents mott spanish name.
	mottSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jones"}

	// mottTraditionalChineseName represents mott traditional chinese name.
	mottTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "李克"}
)

var (
	// mottName represents mott name.
	mottName = nook.Languages{
		language.AmericanEnglish:      mottAmericanEnglishName,
		language.CanadianFrench:       mottCanadianFrenchName,
		language.Dutch:                mottDutchName,
		language.French:               mottFrenchName,
		language.German:               mottGermanName,
		language.Italian:              mottItalianName,
		language.Japanese:             mottJapaneseName,
		language.Korean:               mottKoreanName,
		language.LatinAmericanSpanish: mottLatinAmericanSpanishName,
		language.Russian:              mottRussianName,
		language.SimplifiedChinese:    mottSimplifiedChineseName,
		language.Spanish:              mottSpanishName,
		language.TraditionalChinese:   mottTraditionalChineseName}
)

var (
	// mottCharacter represents mott character.
	mottCharacter = nook.Character{
		Animal:   animal.Lion,
		Birthday: mottBirthday,
		Code:     mottCode,
		Key:      character.Mott,
		Gender:   gender.Male,
		Name:     mottName,
		Special:  false}
)

var (
	// mottAmericanEnglishPhrase represents mott american english phrase.
	mottAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cagey"}

	// mottCanadianFrenchPhrase represents mott canadian french phrase.
	mottCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grif-grif"}

	// mottDutchPhrase represents mott dutch phrase.
	mottDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "rrrrrauw"}

	// mottFrenchPhrase represents mott french phrase.
	mottFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grif-grif"}

	// mottGermanPhrase represents mott german phrase.
	mottGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "gröhl"}

	// mottItalianPhrase represents mott italian phrase.
	mottItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "groarrivo"}

	// mottJapanesePhrase represents mott japanese phrase.
	mottJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ハハハ"}

	// mottKoreanPhrase represents mott korean phrase.
	mottKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "하하하"}

	// mottLatinAmericanSpanishPhrase represents mott latin american spanish phrase.
	mottLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "sapristi"}

	// mottRussianPhrase represents mott russian phrase.
	mottRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "зоосад"}

	// mottSimplifiedChinesePhrase represents mott simplified chinese phrase.
	mottSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哈哈哈"}

	// mottSpanishPhrase represents mott spanish phrase.
	mottSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "sapristi"}

	// mottTraditionalChinesePhrase represents mott traditional chinese phrase.
	mottTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哈哈哈"}
)

var (
	// mottPhrase represents mott phrase.
	mottPhrase = nook.Languages{
		language.AmericanEnglish:      mottAmericanEnglishPhrase,
		language.CanadianFrench:       mottCanadianFrenchPhrase,
		language.Dutch:                mottDutchPhrase,
		language.French:               mottFrenchPhrase,
		language.German:               mottGermanPhrase,
		language.Italian:              mottItalianPhrase,
		language.Japanese:             mottJapanesePhrase,
		language.Korean:               mottKoreanPhrase,
		language.LatinAmericanSpanish: mottLatinAmericanSpanishPhrase,
		language.Russian:              mottRussianPhrase,
		language.SimplifiedChinese:    mottSimplifiedChinesePhrase,
		language.Spanish:              mottSpanishPhrase,
		language.TraditionalChinese:   mottTraditionalChinesePhrase}
)

var (
	// Mott represents mott.
	Mott = nook.Villager{
		Character:   mottCharacter,
		Personality: personality.Jock,
		Phrase:      mottPhrase}
)
