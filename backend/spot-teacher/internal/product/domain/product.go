package domain

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

/* entity */
type Product struct {
	ID          ProductID
	Name        ProductName  // 商品名
	Description *string      // 商品の説明 (nil許可)
	Price       ProductPrice // 商品の価格 (非負)
}

func NewProduct(id ProductID, name ProductName, description *string, price ProductPrice) (*Product, error) {
	if price < 0 {
		return nil, errors.New("price must be non-negative")
	}
	return &Product{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
	}, nil
}

type ProductID int64 // Goの基本型として int64 を使用

func NewProductID(value int64) (ProductID, error) {
	if value <= 0 {
		// IDが0以下は無効とするバリデーション例
		return 0, errors.New("product ID must be positive")
	}
	return ProductID(value), nil
}

func (p ProductID) Value() int64 {
	return int64(p)
}

type ProductName string

func NewProductName(value string) (ProductName, error) {
	const maxLength = 255
	trimmedValue := strings.TrimSpace(value)
	if trimmedValue == "" {
		return "", errors.New("product name cannot be empty or only whitespace")
	}
	if utf8.RuneCountInString(trimmedValue) > maxLength {
		// 3. エラーメッセージ内で変数を使う場合は fmt.Errorf を使用
		return "", fmt.Errorf("product name cannot exceed %d characters", maxLength)
	}
	return ProductName(value), nil
}

func (p ProductName) Value() string {
	return string(p)
}

type ProductPrice float64

func NewProductPrice(value float64) (ProductPrice, error) {
	if value < 0 {
		return 0, errors.New("product price must be non-negative")
	}
	return ProductPrice(value), nil
}
func (p ProductPrice) Value() float64 {
	return float64(p)
}
