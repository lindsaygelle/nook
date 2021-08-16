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
		Value:    "Nathan"}

	nateDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Nate"}

	nateFrenchName = nook.Name{
		Language: language.French,
		Value:    "Nathan"}

	nateGermanName = nook.Name{
		Language: language.German,
		Value:    "Nathan"}

	nateItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gianni"}

	nateJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "バッカス"}

	nateKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "박하스"}

	nateLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nachete"}

	nateRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Нэйт"}

	nateSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "巴克思"}

	nateSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Nachete"}

	nateTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "巴克思"}
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
		Animal:   animal.Bear,
		Birthday: nateBirthday,
		Code:     nateCode,
		Key:      character.Nate,
		Gender:   gender.Male,
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
		Personality: personality.Lazy,
		Phrase:      natePhrase}
)
