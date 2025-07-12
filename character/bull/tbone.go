package bull

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
	// tboneBirthday represents tbone birthday.
	tboneBirthday = nook.Birthday{
		Day:   20,
		Month: time.May}
)

var (
	// tboneCode represents tbone code.
	tboneCode = nook.Code{
		Value: "bul05"}
)

var (
	// tboneAmericanEnglishName represents tbone american english name.
	tboneAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "T-Bone"}

	// tboneCanadianFrenchName represents tbone canadian french name.
	tboneCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Steakos"}

	// tboneDutchName represents tbone dutch name.
	tboneDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "T-Bone"}

	// tboneFrenchName represents tbone french name.
	tboneFrenchName = nook.Name{
		Language: language.French,
		Value:    "Steakos"}

	// tboneGermanName represents tbone german name.
	tboneGermanName = nook.Name{
		Language: language.German,
		Value:    "Marius"}

	// tboneItalianName represents tbone italian name.
	tboneItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Victor"}

	// tboneJapaneseName represents tbone japanese name.
	tboneJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ボルシチ"}

	// tboneKoreanName represents tbone korean name.
	tboneKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "티본"}

	// tboneLatinAmericanSpanishName represents tbone latin american spanish name.
	tboneLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bovi"}

	// tboneRussianName represents tbone russian name.
	tboneRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Ти-Боун"}

	// tboneSimplifiedChineseName represents tbone simplified chinese name.
	tboneSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗宋"}

	// tboneSpanishName represents tbone spanish name.
	tboneSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bovi"}

	// tboneTraditionalChineseName represents tbone traditional chinese name.
	tboneTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羅宋"}
)

var (
	// tboneName represents tbone name.
	tboneName = nook.Languages{
		language.AmericanEnglish:      tboneAmericanEnglishName,
		language.CanadianFrench:       tboneCanadianFrenchName,
		language.Dutch:                tboneDutchName,
		language.French:               tboneFrenchName,
		language.German:               tboneGermanName,
		language.Italian:              tboneItalianName,
		language.Japanese:             tboneJapaneseName,
		language.Korean:               tboneKoreanName,
		language.LatinAmericanSpanish: tboneLatinAmericanSpanishName,
		language.Russian:              tboneRussianName,
		language.SimplifiedChinese:    tboneSimplifiedChineseName,
		language.Spanish:              tboneSpanishName,
		language.TraditionalChinese:   tboneTraditionalChineseName}
)

var (
	// tboneCharacter represents tbone character.
	tboneCharacter = nook.Character{
		Animal:   animal.Bull,
		Birthday: tboneBirthday,
		Code:     tboneCode,
		Key:      character.TBone,
		Gender:   gender.Male,
		Name:     tboneName,
		Special:  false}
)

var (
	// tboneAmericanEnglishPhrase represents tbone american english phrase.
	tboneAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "moocher"}

	// tboneCanadianFrenchPhrase represents tbone canadian french phrase.
	tboneCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mes cornes"}

	// tboneDutchPhrase represents tbone dutch phrase.
	tboneDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "moehoe"}

	// tboneFrenchPhrase represents tbone french phrase.
	tboneFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "vachte"}

	// tboneGermanPhrase represents tbone german phrase.
	tboneGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hhmmuuuhh"}

	// tboneItalianPhrase represents tbone italian phrase.
	tboneItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "muuusica"}

	// tboneJapanesePhrase represents tbone japanese phrase.
	tboneJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だっぺ"}

	// tboneKoreanPhrase represents tbone korean phrase.
	tboneKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "워메에"}

	// tboneLatinAmericanSpanishPhrase represents tbone latin american spanish phrase.
	tboneLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "muuuchi"}

	// tboneRussianPhrase represents tbone russian phrase.
	tboneRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бычок"}

	// tboneSimplifiedChinesePhrase represents tbone simplified chinese phrase.
	tboneSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是呗"}

	// tboneSpanishPhrase represents tbone spanish phrase.
	tboneSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "muooola"}

	// tboneTraditionalChinesePhrase represents tbone traditional chinese phrase.
	tboneTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是唄"}
)

var (
	// tbonePhrase represents tbone phrase.
	tbonePhrase = nook.Languages{
		language.AmericanEnglish:      tboneAmericanEnglishPhrase,
		language.CanadianFrench:       tboneCanadianFrenchPhrase,
		language.Dutch:                tboneDutchPhrase,
		language.French:               tboneFrenchPhrase,
		language.German:               tboneGermanPhrase,
		language.Italian:              tboneItalianPhrase,
		language.Japanese:             tboneJapanesePhrase,
		language.Korean:               tboneKoreanPhrase,
		language.LatinAmericanSpanish: tboneLatinAmericanSpanishPhrase,
		language.Russian:              tboneRussianPhrase,
		language.SimplifiedChinese:    tboneSimplifiedChinesePhrase,
		language.Spanish:              tboneSpanishPhrase,
		language.TraditionalChinese:   tboneTraditionalChinesePhrase}
)

var (
	// TBone represents t bone.
	TBone = nook.Villager{
		Character:   tboneCharacter,
		Personality: personality.Cranky,
		Phrase:      tbonePhrase}
)
