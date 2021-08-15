package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	nateBirthday = nook.Birthday{
		Day:   16,
		Month: time.August}
)

var (
	nateCode = nook.Code{
		Value: "bea05"}
)

var (
	nateAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Nate"}

	nateCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Nathanbaaîîîllle"}

	nateDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Nategaaaap"}

	nateFrenchName = nook.Name{
		Language: language.French,
		Value:    "Nathanbaaîîîllle"}

	nateGermanName = nook.Name{
		Language: language.German,
		Value:    "Nathangäähn"}

	nateItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gianniyauuun"}

	nateJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "バッカスんだ"}

	nateKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "박하스힘내"}

	nateLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nacheteuuaaaah"}

	nateRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Нэйтспа-а-ать"}

	nateSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "巴克思嗯是"}

	nateSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Nacheteuuaaaah"}

	nateTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "巴克思嗯是"}
)

var (
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

var (
	nateCharacter = nook.Character{
		Animal:   Bear,
		Birthday: nateBirthday,
		Code:     nateCode,
		Gender:   nook.Male,
		Name:     nateName}
)

var (
	nateAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "yawwwn"}

	nateCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "baaîîîllle"}

	nateDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "gaaaap"}

	nateFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "baaîîîllle"}

	nateGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "gäähn"}

	nateItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "yauuun"}

	nateJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "んだ"}

	nateKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "힘내"}

	nateLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "uuaaaah"}

	nateRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "спа-а-ать"}

	nateSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嗯是"}

	nateSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "uuaaaah"}

	nateTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗯是"}
)

var (
	natePhrase = nook.Languages{
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

var (
	Nate = nook.Villager{
		Character:   nateCharacter,
		Personality: nook.Lazy,
		Phrase:      natePhrase}
)
