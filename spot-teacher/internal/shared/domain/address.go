package domain

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

type Address struct {
	Prefecture Prefecture
	City       string
	Street     *string
	PostCode   PostCode
}

type Prefecture int64

const (
	PrefectureHokkaido Prefecture = iota
	PrefectureAomori
	PrefectureIwate
	PrefectureMiyagi
	PrefectureAkita
	PrefectureYamagata
	PrefectureFukushima
	PrefectureIbaraki
	PrefectureTochigi
	PrefectureGumma
	PrefectureSaitama
	PrefectureChiba
	PrefectureTokyo
	PrefectureKanagawa
	PrefectureNiigata
	PrefectureToyama
	PrefectureIshikawa
	PrefectureFukui
	PrefectureYamanashi
	PrefectureNagano
	PrefectureGifu
	PrefectureShizuoka
	PrefectureAichi
	PrefectureMie
	PrefectureShiga
	PrefectureKyoto
	PrefectureOsaka
	PrefectureHyogo
	PrefectureNara
	PrefectureWakayama
	PrefectureTottori
	PrefectureShimane
	PrefectureOkayama
	PrefectureHiroshima
	PrefectureYamaguchi
	PrefectureTokushima
	PrefectureKagawa
	PrefectureEhime
	PrefectureKochi
	PrefectureFukuoka
	PrefectureSaga
	PrefectureNagasaki
	PrefectureKumamoto
	PrefectureOita
	PrefectureMiyazaki
	PrefectureKagoshima
	PrefectureOkinawa
)

// (オプション) Stringer インターフェースの実装
var prefectureNames = [...]string{
	"北海道",
	"青森県",
	"岩手県",
	"宮城県",
	"秋田県",
	"山形県",
	"福島県",
	"茨城県",
	"栃木県",
	"群馬県",
	"埼玉県",
	"千葉県",
	"東京都",
	"神奈川県",
	"新潟県",
	"富山県",
	"石川県",
	"福井県",
	"山梨県",
	"長野県",
	"岐阜県",
	"静岡県",
	"愛知県",
	"三重県",
	"滋賀県",
	"京都府",
	"大阪府",
	"兵庫県",
	"奈良県",
	"和歌山県",
	"鳥取県",
	"島根県",
	"岡山県",
	"広島県",
	"山口県",
	"徳島県",
	"香川県",
	"愛媛県",
	"高知県",
	"福岡県",
	"佐賀県",
	"長崎県",
	"熊本県",
	"大分県",
	"宮崎県",
	"鹿児島県",
	"沖縄県",
}

// 　TODO enum にもできるか検討　value objectのメソッドをあてる？
func (p Prefecture) String() string {
	if int(p) >= 0 && int(p) < len(prefectureNames) {
		return prefectureNames[p]
	}
	return fmt.Sprintf("Prefecture(%d)", p)
}

func (p Prefecture) Value() int64 {
	return int64(p)
}

type PostCode string

func NewPostCode(value string) (PostCode, error) {
	const expectedLength = 7
	var postCodeRegex = regexp.MustCompile(`^\d{7}$`)
	trimmedValue := strings.TrimSpace(value)
	if trimmedValue == "" {
		return "", errors.New("post code cannot be empty or only whitespace")
	}
	if utf8.RuneCountInString(trimmedValue) != expectedLength {
		return "", fmt.Errorf("post code must be exactly %d characters long, got %d", expectedLength, utf8.RuneCountInString(trimmedValue))
	}
	if !postCodeRegex.MatchString(trimmedValue) {
		return "", errors.New("invalid post code format: must be 7 digits")
	}
	return PostCode(trimmedValue), nil
}

func (p PostCode) Value() string {
	return string(p)
}
