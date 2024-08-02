package model

import (
	"errors"
	"slices"

	"github.com/averak/hbaas/app/core/numunit"
	"github.com/google/uuid"
)

var (
	ErrKVSEntryValueTooLarge  = errors.New("value size is too large")
	ErrPrivateKVSETagMismatch = errors.New("private KVS etag mismatch")
)

// GlobalKVSBucket は、全ユーザで共有される KVS を表します。
// 専用 API が提供されていないプロダクト固有の機能は、この機能で代用される想定です。
type GlobalKVSBucket struct {
	KVSBucket
}

func NewGlobalKVSBucket(entries []KVSEntry) GlobalKVSBucket {
	return GlobalKVSBucket{
		KVSBucket: NewKVSBucket(entries),
	}
}

// PrivateKVSBucket はユーザごとに独立した KVS で、他ユーザには公開されません。
type PrivateKVSBucket struct {
	KVSBucket

	UserID uuid.UUID
	etag   uuid.UUID
}

func NewPrivateKVSBucket(userID uuid.UUID, etag uuid.UUID, entries []KVSEntry) PrivateKVSBucket {
	return PrivateKVSBucket{
		KVSBucket: NewKVSBucket(entries),
		UserID:    userID,
		etag:      etag,
	}
}

func (m *PrivateKVSBucket) Set(etag uuid.UUID, entries []KVSEntry) error {
	// 同時更新の競合を楽観ロックで防ぐために、クライアントは etag を取得してから更新を行う必要がある。
	if m.etag != etag {
		return ErrPrivateKVSETagMismatch
	}
	m.KVSBucket.Set(entries)
	m.etag = uuid.New()
	return nil
}

func (m PrivateKVSBucket) ETag() uuid.UUID {
	return m.etag
}

// KVSBucket は、KVS におけるエントリの配列操作を提供します。
type KVSBucket struct {
	raw []KVSEntry
}

func NewKVSBucket(entries []KVSEntry) KVSBucket {
	return KVSBucket{
		raw: entries,
	}
}

func (m KVSBucket) Raw() []KVSEntry {
	return m.raw
}

func (m KVSBucket) HasKey(key string) bool {
	return slices.ContainsFunc(m.raw, func(e KVSEntry) bool {
		return e.Key == key
	})
}

func (m *KVSBucket) Set(entries []KVSEntry) {
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

// KVSEntry は、KVS に読み書きされるデータの最小単位です。
type KVSEntry struct {
	Key   string
	Value []byte
}

func NewKVSEntry(key string, value []byte) (KVSEntry, error) {
	// クライアントは、膨大なデータを1エントリにまとめる or 複数シャードに分割することができる。
	// どちらが有利かはデータの I/O 比率に依存するので、無理に最適化せず 100KiB を上限とする。
	if len(value) > 100*numunit.KiB {
		return KVSEntry{}, ErrKVSEntryValueTooLarge
	}

	return KVSEntry{
		Key:   key,
		Value: value,
	}, nil
}

func (e KVSEntry) IsEmpty() bool {
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
