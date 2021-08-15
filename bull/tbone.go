package bull

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	tboneBirthday = nook.Birthday{
		Day:   20,
		Month: time.May}
)

var (
	tboneCode = nook.Code{
		Value: "bul05"}
)

var (
	tboneAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "T-Bone"}

	tboneCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Steakosmes cornes"}

	tboneDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "T-Bonemoehoe"}

	tboneFrenchName = nook.Name{
		Language: language.French,
		Value:    "Steakosvachte"}

	tboneGermanName = nook.Name{
		Language: language.German,
		Value:    "Mariushhmmuuuhh"}

	tboneItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Victormuuusica"}

	tboneJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ボルシチだっぺ"}

	tboneKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "티본워메에"}

	tboneLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bovimuuuchi"}

	tboneRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Ти-Боунбычок"}

	tboneSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗宋是呗"}

	tboneSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bovimuooola"}

	tboneTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羅宋是唄"}
)

var (
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
	tboneCharacter = nook.Character{
		Animal:   Bull,
		Birthday: tboneBirthday,
		Code:     tboneCode,
		Gender:   nook.Male,
		Name:     tboneName}
)

var (
	tboneAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "moocher"}

	tboneCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mes cornes"}

	tboneDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "moehoe"}

	tboneFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "vachte"}

	tboneGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hhmmuuuhh"}

	tboneItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "muuusica"}

	tboneJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だっぺ"}

	tboneKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "워메에"}

	tboneLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "muuuchi"}

	tboneRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бычок"}

	tboneSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是呗"}

	tboneSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "muooola"}

	tboneTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是唄"}
)

var (
	tbonePhrase = nook.Languages{
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
	TBone = nook.Villager{
		Character:   tboneCharacter,
		Personality: nook.Cranky,
		Phrase:      tbonePhrase}
)
