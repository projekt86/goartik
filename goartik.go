package goartik

import (
	"fmt"
)

// NewModule : create new ArtikModule object pointer
func NewModule() (*ArtikModule, error) {
	id, err := getModuleID()
	if err != nil {
		return nil, err
	}
	return createModule(id)
}

// NewModuleWithID : create new ArtikModule object pointer with specific ArtikModuleID
func NewModuleWithID(moduleID ArtikModuleID) (*ArtikModule, error) {
	return createModule(moduleID)
}

func createModule(moduleID ArtikModuleID) (*ArtikModule, error) {
	fmt.Println("New ArtikModule object created")

	m := &ArtikModule{}
	// set module id
	m.ID = moduleID
	// set config
	c, err := newConfig(moduleID)
	if err != nil {
		return nil, err
	}

	m.config = c

	return m, nil
}

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }
