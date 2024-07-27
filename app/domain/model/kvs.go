package model

import (
	"errors"
	"slices"

	"github.com/averak/hbaas/app/core/numunit"
	"github.com/google/uuid"
)

var (
	ErrKVEntryValueTooLarge      = errors.New("value size is too large")
	ErrPrivateKVRevisionMismatch = errors.New("private KVS revision mismatch")
)

// GlobalKVBucket は、全ユーザで共有される KVS を表します。
// 専用 API が提供されていないプロダクト固有の機能は、この機能で代用される想定です。
type GlobalKVBucket struct {
	KVBucket
}

func NewGlobalKVBucket(entries ...KVEntry) GlobalKVBucket {
	return GlobalKVBucket{
		KVBucket: NewKVBucket(entries...),
	}
}

// PrivateKVBucket はユーザごとに独立した KVS で、別ユーザには公開されません。
type PrivateKVBucket struct {
	KVBucket

	UserID   uuid.UUID
	revision uuid.UUID
}

func NewPrivateKVBucket(userID uuid.UUID, revision uuid.UUID, entries ...KVEntry) PrivateKVBucket {
	return PrivateKVBucket{
		KVBucket: NewKVBucket(entries...),
		UserID:   userID,
		revision: revision,
	}
}

func (m *PrivateKVBucket) Set(revision uuid.UUID, entries ...KVEntry) error {
	if m.revision != revision {
		return ErrPrivateKVRevisionMismatch
	}
	m.KVBucket.Set(entries...)
	m.revision = uuid.New()
	return nil
}

func (m PrivateKVBucket) Revision() uuid.UUID {
	return m.revision
}

// KVBucket は、KVS におけるエントリの配列操作を提供します。
type KVBucket struct {
	raw []KVEntry
}

func NewKVBucket(entries ...KVEntry) KVBucket {
	return KVBucket{
		raw: entries,
	}
}

func (m KVBucket) Raw() []KVEntry {
	return m.raw
}

func (m KVBucket) HasKey(key string) bool {
	return slices.ContainsFunc(m.raw, func(e KVEntry) bool {
		return e.Key == key
	})
}

func (m *KVBucket) Set(entries ...KVEntry) {
	for _, e := range entries {
		if m.HasKey(e.Key) {
			for i := range m.raw {
				if m.raw[i].Key == e.Key {
					m.raw[i].Value = e.Value
				}
			}
		} else {
			m.raw = append(m.raw, e)
		}
	}
}

// KVEntry は、KVS に読み書きされるデータの最小単位です。
type KVEntry struct {
	Key   string
	Value []byte
}

func NewKVEntry(key string, value []byte) (KVEntry, error) {
	// クライアントは、膨大なデータを1エントリにまとめる or 複数シャードに分割することができる。
	// どちらが有利かはデータの I/O 比率に依存するので、無理に最適化せず 100KiB を上限とする。
	if len(value) > 100*numunit.KiB {
		return KVEntry{}, ErrKVEntryValueTooLarge
	}

	return KVEntry{
		Key:   key,
		Value: value,
	}, nil
}

func (e KVEntry) IsEmpty() bool {
	return len(e.Value) == 0
}

type KVSCriteria struct {
	ExactMatch  []string
	PrefixMatch []string
}

func NewKVSCriteria(exactMatch, prefixMatch []string) KVSCriteria {
	return KVSCriteria{
		ExactMatch:  exactMatch,
		PrefixMatch: prefixMatch,
	}
}

func (c KVSCriteria) IsEmpty() bool {
	return len(c.ExactMatch) == 0 && len(c.PrefixMatch) == 0
}
