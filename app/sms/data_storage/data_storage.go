package smsDataStorage

import (
	"github.com/mediocregopher/radix/v3"
)

type DataStorage struct {
	db *radix.Pool
}