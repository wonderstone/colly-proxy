package AweiProxy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMonthDayFileName(t *testing.T) {
	// test GetMonthDayFileName
	month := "4"
	date := "23"
	target := "423v2"
	result := IfMonthDayFileName(month, date, target)
	fmt.Println(result)
	assert.Equal(t, true, result)
}

func TestGetFile(t *testing.T) {
	// test GetFile
	month := "4"
	date := "24"

	updated, result := GetFile(URL, month, date)
	fmt.Println(updated, result)
	assert.NotNil(t, result)
}
