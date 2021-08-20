package elephant

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
	chaiBirthday = nook.Birthday{
		Day:   6,
		Month: time.March}
)

var (
	chaiCode = nook.Code{
		Value: "elp11"}
)

var (
	chaiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chai"}

	chaiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Chaï"}

	chaiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chai"}

	chaiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chaï"}

	chaiGermanName = nook.Name{
		Language: language.German,
		Value:    "Chai"}

	chaiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Chai"}

	chaiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フィーカ"}

	chaiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "피카"}

	chaiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Chai"}

	chaiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Чай"}

	chaiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啡卡"}

	chaiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Chai"}

	chaiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啡卡"}
)

var (
	chaiName = nook.Languages{
		language.AmericanEnglish:      chaiAmericanEnglishName,
		language.CanadianFrench:       chaiCanadianFrenchName,
		language.Dutch:                chaiDutchName,
		language.French:               chaiFrenchName,
		language.German:               chaiGermanName,
		language.Italian:              chaiItalianName,
		language.Japanese:             chaiJapaneseName,
		language.Korean:               chaiKoreanName,
		language.LatinAmericanSpanish: chaiLatinAmericanSpanishName,
		language.Russian:              chaiRussianName,
		language.SimplifiedChinese:    chaiSimplifiedChineseName,
		language.Spanish:              chaiSpanishName,
		language.TraditionalChinese:   chaiTraditionalChineseName}
)

var (
	chaiCharacter = nook.Character{
		Animal:   animal.Elephant,
		Birthday: chaiBirthday,
		Code:     chaiCode,
		Key:      character.Chai,
		Gender:   gender.Female,
		Name:     chaiName,
		Special:  false}
)

var (
	chaiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "flap flap"}

	chaiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pachipach"}

	chaiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kaneel"}

	chaiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pachipach"}

	chaiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zimtroll"}

	chaiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ruuu ruuu"}

	chaiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ロルロル"}

	chaiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "롤롤"}

	chaiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "vuelalto"}

	chaiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "машу ушами"}

	chaiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "大耳"}

	chaiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "vuelalto"}

	chaiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "大耳"}
)

var (
	chaiPhrase = nook.Languages{
		language.AmericanEnglish:      chaiAmericanEnglishPhrase,
		language.CanadianFrench:       chaiCanadianFrenchPhrase,
		language.Dutch:                chaiDutchPhrase,
		language.French:               chaiFrenchPhrase,
		language.German:               chaiGermanPhrase,
		language.Italian:              chaiItalianPhrase,
		language.Japanese:             chaiJapanesePhrase,
		language.Korean:               chaiKoreanPhrase,
		language.LatinAmericanSpanish: chaiLatinAmericanSpanishPhrase,
		language.Russian:              chaiRussianPhrase,
		language.SimplifiedChinese:    chaiSimplifiedChinesePhrase,
		language.Spanish:              chaiSpanishPhrase,
		language.TraditionalChinese:   chaiTraditionalChinesePhrase}
)

var (
	Chai = nook.Villager{
		Character:   chaiCharacter,
		Personality: personality.Peppy,
		Phrase:      chaiPhrase}
)
