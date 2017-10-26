package gorobdd

import (
	"fmt"
)

type Vocabulary []string

func NewVocabulary(labels ...string) (*Vocabulary, error) {
	if e := detectDuplicates(labels...); e != nil {
		return nil, e
	}
	v := Vocabulary(labels)
	return &v, nil
}

func detectDuplicates(labels ...string) error {
	m := make(map[string]bool)
	d := []string{}
	for _, l := range labels {
		if m[l] {
			d = append(d, l)
		} else {
			m[l] = true
		}
	}
	if len(d) > 0 {
		return fmt.Errorf("Duplicates in vocabulary: %v", d)
	} else {
		return nil
	}
}

func (v Vocabulary) String() string {
	return fmt.Sprintf("Voc: %v", ([]string)(v))
}
