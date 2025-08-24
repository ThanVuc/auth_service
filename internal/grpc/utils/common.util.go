package utils

import (
	"encoding/json"
	"math"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func ToUUID(id string) (pgtype.UUID, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return pgtype.UUID{}, err
	}

	return pgtype.UUID{
		Bytes: uuid,
		Valid: true,
	}, nil
}

func RoundToTwoDecimal(val float64) float64 {
	return math.Round(val*100) / 100
}

func FromPgTypeTimeToUnix(t pgtype.Timestamp) *int64 {
	if !t.Valid {
		return nil
	}
	unixTime := t.Time.Unix()
	return &unixTime
}

func FromPgTypeTimeStamptZToUnix(t pgtype.Timestamptz) *int64 {
	if !t.Valid {
		return nil
	}
	unixTime := t.Time.Unix()
	return &unixTime
}

func Difference[T comparable](a, b []T) []T {
	m := make(map[T]struct{}, len(b))
	for _, item := range b {
		m[item] = struct{}{}
	}

	var diff []T
	for _, item := range a {
		if _, found := m[item]; !found {
			diff = append(diff, item)
		}
	}
	return diff
}

func ToJSON(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(data)
}

func ToBoolPointer(b bool) *bool {
	return &b
}

func ToStringPointer(s string) *string {
	return &s
}
