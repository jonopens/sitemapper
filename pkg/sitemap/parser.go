package sitemap

import (
	"encoding/xml"
	"fmt"
)

// Parser handles XML sitemap parsing
type Parser struct{}

// NewParser creates a new sitemap parser
func NewParser() *Parser {
	return &Parser{}
}

// Sitemap represents a parsed sitemap
type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	URLs    []URL    `xml:"url"`
}

// SitemapIndex represents a sitemap index
type SitemapIndex struct {
	XMLName  xml.Name  `xml:"sitemapindex"`
	Sitemaps []Sitemap `xml:"sitemap"`
}

// URL represents a single URL entry in a sitemap
type URL struct {
	Loc        string  `xml:"loc"`
	LastMod    string  `xml:"lastmod"`
	ChangeFreq string  `xml:"changefreq"`
	Priority   float64 `xml:"priority"`
}

// Parse parses XML data into a Sitemap structure
func (p *Parser) Parse(data []byte) (*Sitemap, error) {
	var sitemap Sitemap
	if err := xml.Unmarshal(data, &sitemap); err != nil {
		return nil, fmt.Errorf("failed to parse sitemap: %w", err)
	}
	return &sitemap, nil
}

// ParseIndex parses XML data into a SitemapIndex structure
func (p *Parser) ParseIndex(data []byte) (*SitemapIndex, error) {
	var index SitemapIndex
	if err := xml.Unmarshal(data, &index); err != nil {
		return nil, fmt.Errorf("failed to parse sitemap index: %w", err)
	}
	return &index, nil
}

// DetectType determines if the XML is a sitemap or sitemap index
func (p *Parser) DetectType(data []byte) (string, error) {
	// Try to detect from root element
	if len(data) == 0 {
		return "", fmt.Errorf("empty data")
	}
	
	// Simple detection based on root element
	var temp struct {
		XMLName xml.Name
	}
	if err := xml.Unmarshal(data, &temp); err != nil {
		return "", err
	}
	
	switch temp.XMLName.Local {
	case "urlset":
		return "sitemap", nil
	case "sitemapindex":
		return "index", nil
	default:
		return "", fmt.Errorf("unknown sitemap type: %s", temp.XMLName.Local)
	}
}

