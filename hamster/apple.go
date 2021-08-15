package hamster

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	appleBirthday = nook.Birthday{
		Day:   24,
		Month: time.September}
)

var (
	appleCode = nook.Code{
		Value: "ham01"}
)

var (
	appleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Apple"}

	appleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Estherbibille"}

	appleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Applewangetjes"}

	appleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Estherbibille"}

	appleGermanName = nook.Name{
		Language: language.German,
		Value:    "Jessifiep"}

	appleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ciccitriplo uau"}

	appleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アップルキュルン"}

	appleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "애플큐룽"}

	appleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rosido-re-mi"}

	appleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Эплщечки"}

	appleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "苹果转转"}

	appleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rosido-re-mi"}

	appleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蘋果轉轉"}
)

var (
	appleName = nook.Languages{
		language.AmericanEnglish:      appleAmericanEnglishName,
		language.CanadianFrench:       appleCanadianFrenchName,
		language.Dutch:                appleDutchName,
		language.French:               appleFrenchName,
		language.German:               appleGermanName,
		language.Italian:              appleItalianName,
		language.Japanese:             appleJapaneseName,
		language.Korean:               appleKoreanName,
		language.LatinAmericanSpanish: appleLatinAmericanSpanishName,
		language.Russian:              appleRussianName,
		language.SimplifiedChinese:    appleSimplifiedChineseName,
		language.Spanish:              appleSpanishName,
		language.TraditionalChinese:   appleTraditionalChineseName}
)

var (
	appleCharacter = nook.Character{
		Animal:   Hamster,
		Birthday: appleBirthday,
		Code:     appleCode,
		Gender:   nook.Female,
		Name:     appleName}
)

var (
	appleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cheekers"}

	appleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bibille"}

	appleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "wangetjes"}

	appleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bibille"}

	appleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "fiep"}

	appleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "triplo uau"}

	appleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "キュルン"}

	appleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "큐룽"}

	appleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "do-re-mi"}

	appleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "щечки"}

	appleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "转转"}

	appleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "do-re-mi"}

	appleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "轉轉"}
)

var (
	applePhrase = nook.Languages{
		language.AmericanEnglish:      appleAmericanEnglishName,
		language.CanadianFrench:       appleCanadianFrenchName,
		language.Dutch:                appleDutchName,
		language.French:               appleFrenchName,
		language.German:               appleGermanName,
		language.Italian:              appleItalianName,
		language.Japanese:             appleJapaneseName,
		language.Korean:               appleKoreanName,
		language.LatinAmericanSpanish: appleLatinAmericanSpanishName,
		language.Russian:              appleRussianName,
		language.SimplifiedChinese:    appleSimplifiedChineseName,
		language.Spanish:              appleSpanishName,
		language.TraditionalChinese:   appleTraditionalChineseName}
)

var (
	Apple = nook.Villager{
		Character:   appleCharacter,
		Personality: nook.Peppy,
		Phrase:      applePhrase}
)
