package model

import (
	"errors"
	"sort"
	"time"

	"github.com/averak/hbaas/app/core/numunit"
)

var (
	ErrMasterDataContentTooLarge = errors.New("master data content is too large")
)

type MasterData struct {
	Revision  int
	Content   []byte
	IsActive  bool
	Comment   string
	CreatedAt time.Time
}

func NewMasterData(revision int, content []byte, isActive bool, comment string, createdAt time.Time) (MasterData, error) {
	if len(content) > numunit.MiB {
		return MasterData{}, ErrMasterDataContentTooLarge
	}
	return MasterData{
		Revision:  revision,
		Content:   content,
		IsActive:  isActive,
		Comment:   comment,
		CreatedAt: createdAt,
	}, nil
}

type MasterDataService struct {
	active          MasterData
	sortedRevisions []int
}

func NewMasterDataService(active MasterData, revisions []int) (*MasterDataService, error) {
	if !active.IsActive {
		return nil, errors.New("active master data is not active")
	}
	if len(revisions) == 0 {
		return nil, errors.New("revisions is empty")
	}
	sort.Ints(revisions)
	return &MasterDataService{
		active:          active,
		sortedRevisions: revisions,
	}, nil
}

func (s *MasterDataService) SwitchActive(target MasterData) (MasterData, MasterData, error) {
	if target.IsActive {
		return MasterData{}, MasterData{}, errors.New("target master data is already active")
	}

	old := s.active
	old.IsActive = false
	target.IsActive = true
	s.active = target
	return old, target, nil
}

func (s *MasterDataService) NewNextRevision(content []byte, comment string, now time.Time) (MasterData, error) {
	// revision が空になることはないので、index out of range は発生しない。
	revision := s.sortedRevisions[len(s.sortedRevisions)-1] + 1
	s.sortedRevisions = append(s.sortedRevisions, revision)
	return NewMasterData(revision, content, false, comment, now)
}
