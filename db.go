package main

// DB needs to be implemented
type DB interface {
	Set(inst, category, name string, status CheckStatus, hint string, expirySecs int) error
	Get() (inst, category, name string)
}

type Category struct {
	CategoryName string      `json:"CategoryName,omitempty"`
	CheckName    string      `json:"CheckName,omitempty"`
	CheckValue   CheckStatus `json:"CheckValue,omitempty"`
	CheckHint    string      `json:"CheckHint,omitempty"`
}

// InMemoryDB poc
type InMemoryDB struct {
	Instance map[string][]Category `json:"Instance,omitempty"`
}

func (b *InMemoryDB) Set(inst, category, name string, status CheckStatus, hint string, expirySecs int) error {
	//
	return nil
}

func NewInMemoryDB() *InMemoryDB {
	mem := InMemoryDB{}
	mem.Instance = make(map[string][]Category)
	return &mem
}
