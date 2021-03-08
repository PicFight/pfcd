package deps

import (
	"github.com/jfixby/pin"
	"strings"
)

type Dependency struct {
	Import  string
	Fork    int
	Version string
}

type GoModHandler struct {
	Tag          string
	Dependencies []Dependency
}

type DepsGraph struct {
	Vertices map[string]*GoModHandler
}

func (v DepsGraph) ListChildrenForVertex(vertexID string) []string {
	result := []string{}

	//dps := v.g.Vertices[v.Tag]
	dps := v.Vertices[vertexID]
	deps := dps.Dependencies
	PREF := "github.com/decred/dcrd/"
	for _, dp := range deps {
		im := dp.Import
		if !strings.HasPrefix(im, PREF) {
			continue
		}
		key := im[len(PREF)-1:]
		cv := v.Vertices[key]
		if cv == nil {

			pin.D("missing key", key+" : "+dp.Version)
			pin.D("v.g.Vertices", v.Vertices)
			pin.AssertNotNil(key, cv)
			continue
		}
		//pin.D("v.g.Vertices", v.g.Vertices)
		pin.AssertNotNil(key, cv)
		result = append(result, cv.Tag)
	}
	return result
}

func (d DepsGraph) ListVertices() []string {
	result := []string{}
	for _, v := range d.Vertices {
		result = append(result, v.Tag)
	}
	return result
}
