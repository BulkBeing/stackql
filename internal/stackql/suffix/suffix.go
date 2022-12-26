package suffix

import (
	"github.com/stackql/go-openapistackql/openapistackql"

	"github.com/stackql/go-suffix-map/pkg/suffixmap"
)

var (
	_ ParameterSuffixMap = &standardParameterSuffixMap{}
)

type ParameterSuffixMap interface {
	Delete(k string) bool
	Get(k string) (openapistackql.Addressable, bool)
	GetAll() map[string]openapistackql.Addressable
	Put(k string, v openapistackql.Addressable)
	Size() int
}

type standardParameterSuffixMap struct {
	sm suffixmap.SuffixMap
}

func NewParameterSuffixMap() ParameterSuffixMap {
	return &standardParameterSuffixMap{
		sm: suffixmap.NewSuffixMap(nil),
	}
}

func (psm *standardParameterSuffixMap) Get(k string) (openapistackql.Addressable, bool) {
	rv, ok := psm.sm.Get(k)
	if !ok {
		return nil, false
	}
	crv, ok := rv.(openapistackql.Addressable)
	return crv, ok
}

func (psm *standardParameterSuffixMap) GetAll() map[string]openapistackql.Addressable {
	m := psm.sm.GetAll()
	rv := make(map[string]openapistackql.Addressable)
	for k, v := range m {
		p, ok := v.(openapistackql.Addressable)
		if ok {
			rv[k] = p
		}
	}
	return rv
}

func (psm *standardParameterSuffixMap) Put(k string, v openapistackql.Addressable) {
	psm.sm.Put(k, v)
}

func (psm *standardParameterSuffixMap) Delete(k string) bool {
	return psm.sm.Delete(k)
}

func (psm *standardParameterSuffixMap) Size() int {
	return psm.sm.Size()
}
