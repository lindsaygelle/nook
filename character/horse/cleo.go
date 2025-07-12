package horse

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
	// cleoBirthday represents cleo birthday.
	cleoBirthday = nook.Birthday{
		Day:   9,
		Month: time.February}
)

var (
	// cleoCode represents cleo code.
	cleoCode = nook.Code{
		Value: "hrs07"}
)

var (
	// cleoAmericanEnglishName represents cleo american english name.
	cleoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cleo"}

	// cleoCanadianFrenchName represents cleo canadian french name.
	cleoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cléa"}

	// cleoDutchName represents cleo dutch name.
	cleoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cleo"}

	// cleoFrenchName represents cleo french name.
	cleoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cléa"}

	// cleoGermanName represents cleo german name.
	cleoGermanName = nook.Name{
		Language: language.German,
		Value:    "Birgit"}

	// cleoItalianName represents cleo italian name.
	cleoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Susanna"}

	// cleoJapaneseName represents cleo japanese name.
	cleoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アイソトープ"}

	// cleoKoreanName represents cleo korean name.
	cleoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아이소토프"}

	// cleoLatinAmericanSpanishName represents cleo latin american spanish name.
	cleoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Clotilde"}

	// cleoRussianName represents cleo russian name.
	cleoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Клео"}

	// cleoSimplifiedChineseName represents cleo simplified chinese name.
	cleoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小铀"}

	// cleoSpanishName represents cleo spanish name.
	cleoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Clotilde"}

	// cleoTraditionalChineseName represents cleo traditional chinese name.
	cleoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小鈾"}
)

var (
	// cleoName represents cleo name.
	cleoName = nook.Languages{
		language.AmericanEnglish:      cleoAmericanEnglishName,
		language.CanadianFrench:       cleoCanadianFrenchName,
		language.Dutch:                cleoDutchName,
		language.French:               cleoFrenchName,
		language.German:               cleoGermanName,
		language.Italian:              cleoItalianName,
		language.Japanese:             cleoJapaneseName,
		language.Korean:               cleoKoreanName,
		language.LatinAmericanSpanish: cleoLatinAmericanSpanishName,
		language.Russian:              cleoRussianName,
		language.SimplifiedChinese:    cleoSimplifiedChineseName,
		language.Spanish:              cleoSpanishName,
		language.TraditionalChinese:   cleoTraditionalChineseName}
)

var (
	// cleoCharacter represents cleo character.
	cleoCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: cleoBirthday,
		Code:     cleoCode,
		Key:      character.Cleo,
		Gender:   gender.Female,
		Name:     cleoName,
		Special:  false}
)

var (
	// cleoAmericanEnglishPhrase represents cleo american english phrase.
	cleoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sugar"}

	// cleoCanadianFrenchPhrase represents cleo canadian french phrase.
	cleoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "beauté"}

	// cleoDutchPhrase represents cleo dutch phrase.
	cleoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snoepje"}

	// cleoFrenchPhrase represents cleo french phrase.
	cleoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "beauté"}

	// cleoGermanPhrase represents cleo german phrase.
	cleoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "plätzchen"}

	// cleoItalianPhrase represents cleo italian phrase.
	cleoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "dolcetto"}

	// cleoJapanesePhrase represents cleo japanese phrase.
	cleoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "あそばせ"}

	// cleoKoreanPhrase represents cleo korean phrase.
	cleoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "훗훗"}

	// cleoLatinAmericanSpanishPhrase represents cleo latin american spanish phrase.
	cleoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "hop-puá"}

	// cleoRussianPhrase represents cleo russian phrase.
	cleoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "рафинад"}

	// cleoSimplifiedChinesePhrase represents cleo simplified chinese phrase.
	cleoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "游游"}

	// cleoSpanishPhrase represents cleo spanish phrase.
	cleoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bomboncito"}

	// cleoTraditionalChinesePhrase represents cleo traditional chinese phrase.
	cleoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "遊遊"}
)

var (
	// cleoPhrase represents cleo phrase.
	cleoPhrase = nook.Languages{
		language.AmericanEnglish:      cleoAmericanEnglishPhrase,
		language.CanadianFrench:       cleoCanadianFrenchPhrase,
		language.Dutch:                cleoDutchPhrase,
		language.French:               cleoFrenchPhrase,
		language.German:               cleoGermanPhrase,
		language.Italian:              cleoItalianPhrase,
		language.Japanese:             cleoJapanesePhrase,
		language.Korean:               cleoKoreanPhrase,
		language.LatinAmericanSpanish: cleoLatinAmericanSpanishPhrase,
		language.Russian:              cleoRussianPhrase,
		language.SimplifiedChinese:    cleoSimplifiedChinesePhrase,
		language.Spanish:              cleoSpanishPhrase,
		language.TraditionalChinese:   cleoTraditionalChinesePhrase}
)

var (
	// Cleo represents cleo.
	Cleo = nook.Villager{
		Character:   cleoCharacter,
		Personality: personality.Snooty,
		Phrase:      cleoPhrase}
)
