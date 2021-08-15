package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	maddieBirthday = nook.Birthday{
		Day:   11,
		Month: time.January}
)

var (
	maddieCode = nook.Code{
		Value: "dog09"}
)

var (
	maddieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Maddie"}

	maddieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Olympeyoupi"}

	maddieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Maddiejippie"}

	maddieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Olympeyoup"}

	maddieGermanName = nook.Name{
		Language: language.German,
		Value:    "Agneshurra"}

	maddieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cristinayuuuppii"}

	maddieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マロンワンツー"}

	maddieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마롱달코미"}

	maddieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Martinareguau"}

	maddieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мэддиур-ра"}

	maddieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "麻蓉一二"}

	maddieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Martinachupi"}

	maddieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "麻蓉一二"}
)

var (
	maddieName = nook.Languages{
		language.AmericanEnglish:      maddieAmericanEnglishName,
		language.CanadianFrench:       maddieCanadianFrenchName,
		language.Dutch:                maddieDutchName,
		language.French:               maddieFrenchName,
		language.German:               maddieGermanName,
		language.Italian:              maddieItalianName,
		language.Japanese:             maddieJapaneseName,
		language.Korean:               maddieKoreanName,
		language.LatinAmericanSpanish: maddieLatinAmericanSpanishName,
		language.Russian:              maddieRussianName,
		language.SimplifiedChinese:    maddieSimplifiedChineseName,
		language.Spanish:              maddieSpanishName,
		language.TraditionalChinese:   maddieTraditionalChineseName}
)

var (
	maddieCharacter = nook.Character{
		Animal:   Dog,
		Birthday: maddieBirthday,
		Code:     maddieCode,
		Gender:   nook.Female,
		Name:     maddieName}
)

var (
	maddieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "yippee"}

	maddieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "youpi"}

	maddieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "jippie"}

	maddieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "youp"}

	maddieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hurra"}

	maddieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "yuuuppii"}

	maddieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ワンツー"}

	maddieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "달코미"}

	maddieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "reguau"}

	maddieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ур-ра"}

	maddieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "一二"}

	maddieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chupi"}

	maddieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "一二"}
)

var (
	maddiePhrase = nook.Languages{
		language.AmericanEnglish:      maddieAmericanEnglishName,
		language.CanadianFrench:       maddieCanadianFrenchName,
		language.Dutch:                maddieDutchName,
		language.French:               maddieFrenchName,
		language.German:               maddieGermanName,
		language.Italian:              maddieItalianName,
		language.Japanese:             maddieJapaneseName,
		language.Korean:               maddieKoreanName,
		language.LatinAmericanSpanish: maddieLatinAmericanSpanishName,
		language.Russian:              maddieRussianName,
		language.SimplifiedChinese:    maddieSimplifiedChineseName,
		language.Spanish:              maddieSpanishName,
		language.TraditionalChinese:   maddieTraditionalChineseName}
)

var (
	Maddie = nook.Villager{
		Character:   maddieCharacter,
		Personality: nook.Peppy,
		Phrase:      maddiePhrase}
)
