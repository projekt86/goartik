package goartik

import (
	"io/ioutil"
	"regexp"
)

const (
	// A053 : Module A053 model
	A053 ArtikModuleID = iota
	// A020 : 020 model
	A020
	// A030 : 030 model
	A030
	// A520 : 520 model
	A520
	// A530 : 530 model
	A530
	// A710 : 710 model
	A710
	// A1020 : 1020 model
	A1020
)

const (
	artikReleaseTestPath = "/test/device/etc/artik_release"
	artikReleasePath     = "/etc/artik_release"
)

// ArtikModuleID : id of the Artik module
type ArtikModuleID uint8

// String : convert module id object to string object
func (s ArtikModuleID) String() string {
	name := []string{"ARTIK053", "ARTIK052", "ARTIK030", "ARTIK520", "ARTIK530", "ARTIK710", "ARTIK1020"}
	i := uint8(s)
	switch {
	case i <= uint8(A1020):
		return name[i]
	default:
		return "UNKNOWN"
	}
}

// ArtikModule : ArtikModule object
type ArtikModule struct {
	ID     ArtikModuleID
	config config
}

// getModuleID : get model from system
func getModuleID() (ArtikModuleID, error) {
	var path string
	if isVirtualDevice {
		path = artikReleaseTestPath
	} else {
		path = artikReleasePath
	}
	b, err := ioutil.ReadFile(getRelativePath(path))
	if err != nil {
		return 0, err
	}

	var re = regexp.MustCompile(`MODEL=([0-9A-Z-]+)`)
	for _, bv := range re.FindAllStringSubmatch(string(b), -1) {
		if len(bv) == 2 {
			return parseModel(bv[1]), nil
		}
	}

	return 0, nil
}

// parseModel : parse model from string
func parseModel(model string) ArtikModuleID {
	switch model {
	case "ARTIK520":
		return A520
	case "ARTIK710":
		return A710
	}
	panic("not recognized module")
}
