package openstates

import (
	"fmt"
	"github.com/sunlightlabs/go-sunlight/internal"
)

/**
 * Open States Legislator object. Issued for members of the Legislature.
 */
type Legislator struct {
	Timestamps
	Sources

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	FullName  string `json:"full_name"`
	Roles     []struct {
		Term           string `json:"term"`
		District       string `json:"district"`
		Chamber        string `json:"chamber"`
		Committee      string `json:"committee"`
		CommitteeId    string `json:"committee_id"`
		Subcommittee   string `json:"subcommittee"`
		SubcommitteeId string `json:"subcommittee_id"`
		Position       string `json:"position"`
		State          string `json:"state"`
		Party          string `json:"party"`
		Type           string `json:"type"`
	} `json:"roles"`
	Id         string `json:"id"`
	MiddleName string `json:"middle_name"`
	State      string `json:"state"`
	Active     bool   `json:"active"`
	PhotoUrl   string `json:"photo_url"`
	Url        string `json:"url"`
}

/**
 * Get an Open States Legislator by their Open States issued ID, such as
 * MAL000155
 */
func GetLegislator(bigId string) (*Legislator, error) {
	l := &Legislator{}
	err := internal.GetURL(l, openstatesRoot, map[string]string{}, "legislators", bigId)
	if err != nil {
		return nil, err
	}
	return l, nil
}

/**
 * Query the Open States API by params defined in
 * http://sunlightlabs.github.io/openstates-api/legislators.html#methods/legislator-search
 */
func GetLegislators(criteria map[string]string) (*[]Legislator, error) {
	l := []Legislator{}
	err := internal.GetURL(&l, openstatesRoot, criteria, "legislators")
	if err != nil {
		return nil, err
	}
	return &l, nil
}

func GetLegislatorsByLatLon(lat float32, lon float32) (*[]Legislator, error) {
	var params = map[string]string{
		"lat":  fmt.Sprintf("%f", lat),
		"long": fmt.Sprintf("%f", lon),
	}

	l := []Legislator{}
	err := internal.GetURL(&l, openstatesRoot, params, "legislators", "geo")
	if err != nil {
		return nil, err
	}
	return &l, nil
}
