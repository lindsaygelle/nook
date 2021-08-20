package dog

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
		Value:    "Olympe"}

	maddieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Maddie"}

	maddieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Olympe"}

	maddieGermanName = nook.Name{
		Language: language.German,
		Value:    "Agnes"}

	maddieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cristina"}

	maddieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マロン"}

	maddieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마롱"}

	maddieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Martina"}

	maddieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мэдди"}

	maddieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "麻蓉"}

	maddieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Martina"}

	maddieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "麻蓉"}
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
		Animal:   animal.Dog,
		Birthday: maddieBirthday,
		Code:     maddieCode,
		Key:      character.Maddie,
		Gender:   gender.Female,
		Name:     maddieName,
		Special:  false}
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
		language.AmericanEnglish:      maddieAmericanEnglishPhrase,
		language.CanadianFrench:       maddieCanadianFrenchPhrase,
		language.Dutch:                maddieDutchPhrase,
		language.French:               maddieFrenchPhrase,
		language.German:               maddieGermanPhrase,
		language.Italian:              maddieItalianPhrase,
		language.Japanese:             maddieJapanesePhrase,
		language.Korean:               maddieKoreanPhrase,
		language.LatinAmericanSpanish: maddieLatinAmericanSpanishPhrase,
		language.Russian:              maddieRussianPhrase,
		language.SimplifiedChinese:    maddieSimplifiedChinesePhrase,
		language.Spanish:              maddieSpanishPhrase,
		language.TraditionalChinese:   maddieTraditionalChinesePhrase}
)

var (
	Maddie = nook.Villager{
		Character:   maddieCharacter,
		Personality: personality.Peppy,
		Phrase:      maddiePhrase}
)
