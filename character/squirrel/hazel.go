package squirrel

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
	hazelBirthday = nook.Birthday{
		Day:   30,
		Month: time.August}
)

var (
	hazelCode = nook.Code{
		Value: "squ18"}
)

var (
	hazelAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hazel"}

	hazelCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pamela"}

	hazelDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hazel"}

	hazelFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pamela"}

	hazelGermanName = nook.Name{
		Language: language.German,
		Value:    "Nadine"}

	hazelItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cigliola"}

	hazelJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アイリス"}

	hazelKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아이리스"}

	hazelLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Anacarda"}

	hazelRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Хейзел"}

	hazelSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "艾栗丝"}

	hazelSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Anacarda"}

	hazelTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "艾栗絲"}
)

var (
	hazelName = nook.Languages{
		language.AmericanEnglish:      hazelAmericanEnglishName,
		language.CanadianFrench:       hazelCanadianFrenchName,
		language.Dutch:                hazelDutchName,
		language.French:               hazelFrenchName,
		language.German:               hazelGermanName,
		language.Italian:              hazelItalianName,
		language.Japanese:             hazelJapaneseName,
		language.Korean:               hazelKoreanName,
		language.LatinAmericanSpanish: hazelLatinAmericanSpanishName,
		language.Russian:              hazelRussianName,
		language.SimplifiedChinese:    hazelSimplifiedChineseName,
		language.Spanish:              hazelSpanishName,
		language.TraditionalChinese:   hazelTraditionalChineseName}
)

var (
	hazelCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: hazelBirthday,
		Code:     hazelCode,
		Key:      character.Hazel,
		Gender:   gender.Female,
		Name:     hazelName}
)

var (
	hazelAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "uni-wow"}

	hazelCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tac tac"}

	hazelDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "brauw"}

	hazelFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tac tac"}

	hazelGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kekeke"}

	hazelItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "blink"}

	hazelJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だもんね"}

	hazelKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "알아알아"}

	hazelLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tac-tac"}

	hazelRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "фундук"}

	hazelSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "当然哦"}

	hazelSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "castaña"}

	hazelTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "當然囉"}
)

var (
	hazelPhrase = nook.Languages{
		language.AmericanEnglish:      hazelAmericanEnglishPhrase,
		language.CanadianFrench:       hazelCanadianFrenchPhrase,
		language.Dutch:                hazelDutchPhrase,
		language.French:               hazelFrenchPhrase,
		language.German:               hazelGermanPhrase,
		language.Italian:              hazelItalianPhrase,
		language.Japanese:             hazelJapanesePhrase,
		language.Korean:               hazelKoreanPhrase,
		language.LatinAmericanSpanish: hazelLatinAmericanSpanishPhrase,
		language.Russian:              hazelRussianPhrase,
		language.SimplifiedChinese:    hazelSimplifiedChinesePhrase,
		language.Spanish:              hazelSpanishPhrase,
		language.TraditionalChinese:   hazelTraditionalChinesePhrase}
)

var (
	Hazel = nook.Villager{
		Character:   hazelCharacter,
		Personality: personality.BigSister,
		Phrase:      hazelPhrase}
)
