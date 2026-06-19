package common

import (
	"math"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// UUIDPtr parses s into *uuid.UUID. Returns nil if s is empty.
// Panics if s is non-empty but invalid (use MustParse semantics).
func UUIDPtr(s string) *uuid.UUID {
	if s == "" {
		return nil
	}
	u := uuid.MustParse(s)
	return &u
}

// ToPtr returns a pointer to v. Generic utility for any type.
func ToPtr[T any](v T) *T { return &v }

// HashPassword hashes the plain password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword compares plain vs hashed
func CheckPassword(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}

// Rounder to round float number with some intermed and precision decimal
// e.g: 2.558--->2.6
// Rounder(2.558, 0.5, 1)--->2.6 (intermed is a number that determine range of round ceil and round floor
// can accept negatif number
// from:https://play.golang.org/p/KNhgeuU5sT
func Rounder(num float64, intermed float64, precision int) float64 {
	pow := math.Pow(10, float64(precision))
	digit := pow * num
	_, div := math.Modf(digit)

	var round float64
	if num > 0 {
		if div >= intermed {
			round = math.Ceil(digit)
		} else {
			round = math.Floor(digit)
		}
	} else {
		if div >= intermed {
			round = math.Floor(digit)
		} else {
			round = math.Ceil(digit)
		}
	}

	return round / pow
}
