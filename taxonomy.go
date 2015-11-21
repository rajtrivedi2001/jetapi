package jetapi

import (
	"net/url"
	"strconv"
)

type TaxonomyNode struct {
	Id               int    `json:"jet_node_id"`
	Name             string `json:"jet_node_name"`
	Path             string `json:"jet_node_path"`
	ParentId         int    `json:"parent_id"`
	Level            int    `json:"jet_level"`
	SuggestedTaxCode string `json:"suggested_tax_code"`
}

type TaxonomyLinks struct {
	Links []string        `json:"node_urls"`
	Nodes []*TaxonomyNode `json:"-"`
}

func (a *JetApi) GetTaxonomyLinks(limit int, offset int) (*TaxonomyLinks, error) {
	links := TaxonomyLinks{}

	if err := a.EnsureValidToken(); err != nil {
		return nil, err
	}

	query := url.Values{}
	query.Set("limit", strconv.Itoa(limit))
	query.Set("offset", strconv.Itoa(offset))

	req := a.CreateGetRequest("/taxonomy/links/v1", query)

	err := a.DoRequest(req, &links)
	if err != nil {
		return nil, err
	}

	links.Nodes = []*TaxonomyNode{}
	for _, link := range links.Links {
		var node TaxonomyNode

		req := a.CreateGetRequest(link, url.Values{})
		err := a.DoRequest(req, &node)
		if err != nil {
			return nil, err
		}

		links.Nodes = append(links.Nodes, &node)
	}

	return &links, nil
}
