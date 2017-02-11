package cookiejar

import (
	"encoding/json"
)

type entrySerialize struct {
	Entry  entry
	SeqNum uint64
}

type jarSerialize struct {
	Entries    map[string]map[string]entrySerialize
	NextSeqNum uint64
}

// Clone returns a new cookie jar from the result of Jar.MarshalJSON
func Clone(o *Options, json []byte) (*Jar, error) {
	jar, err := New(o)
	if len(json) == 0 || err != nil {
		return jar, err
	}

	err = jar.unmarshalJSON(json)
	return jar, err
}

// MarshalJSON implements the MarshalJSON method of the json.Marshaler interface.
func (j *Jar) MarshalJSON() ([]byte, error) {
	js := &jarSerialize{
		Entries: make(map[string]map[string]entrySerialize),
	}

	j.mu.Lock()
	defer j.mu.Unlock()
	// Jar to jarSerialize
	for k, v := range j.entries {
		v1 := make(map[string]entrySerialize)
		for kk, vv := range v {
			v1[kk] = entrySerialize{
				Entry:  vv,
				SeqNum: vv.seqNum,
			}
		}
		js.Entries[k] = v1
	}
	js.NextSeqNum = j.nextSeqNum

	return json.Marshal(js)
}

// implements the unmarshalJSON method
func (j *Jar) unmarshalJSON(data []byte) error {
	js := &jarSerialize{
		Entries: make(map[string]map[string]entrySerialize),
	}
	err := json.Unmarshal(data, js)
	if err != nil {
		return err
	}

	// map[string]map[string]entrySerialize to map[string]map[string]entry
	entries := make(map[string]map[string]entry)
	for k, v := range js.Entries {
		v1 := make(map[string]entry)
		for kk, vv := range v {
			vv.Entry.seqNum = vv.SeqNum
			v1[kk] = vv.Entry
		}
		entries[k] = v1
	}

	j.mu.Lock()
	defer j.mu.Unlock()
	// copy to Jar
	j.entries = entries
	j.nextSeqNum = js.NextSeqNum

	return nil
}
