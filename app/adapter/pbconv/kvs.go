package pbconv

import (
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/pkg/vector"
	"github.com/averak/hbaas/protobuf/resource"
)

func FromKVSCriteriaPb(pbs []*resource.KVSCriterion) model.KVSCriteria {
	exactMatch := make([]string, 0)
	prefixMatch := make([]string, 0)
	for _, pb := range pbs {
		switch pb.GetMatchingType() {
		case resource.KVSCriterion_MATCHING_TYPE_UNSPECIFIED:
			continue
		case resource.KVSCriterion_MATCHING_TYPE_EXACT_MATCH:
			exactMatch = append(exactMatch, pb.GetKey())
		case resource.KVSCriterion_MATCHING_TYPE_PREFIX_MATCH:
			prefixMatch = append(prefixMatch, pb.GetKey())
		}
	}
	return model.NewKVSCriteria(exactMatch, prefixMatch)
}

func ToKVSEntryPb(entry model.KVSEntry) *resource.KVSEntry {
	return &resource.KVSEntry{
		Key:   entry.Key,
		Value: entry.Value,
	}
}

func ToKVSEntryPbs(entries []model.KVSEntry) []*resource.KVSEntry {
	return vector.Map(entries, ToKVSEntryPb)
}

func FromKVSEntryPb(pb *resource.KVSEntry) (model.KVSEntry, error) {
	return model.NewKVSEntry(pb.GetKey(), pb.GetValue())
}

func FromKVSEntryPbs(pbs []*resource.KVSEntry) ([]model.KVSEntry, error) {
	res := make([]model.KVSEntry, 0)
	for _, pb := range pbs {
		entry, err := FromKVSEntryPb(pb)
		if err != nil {
			return nil, err
		}
		res = append(res, entry)
	}
	return res, nil
}
