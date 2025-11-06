package sitemap

import (
	"fmt"
	"net/url"
)

// Validator validates sitemap structure and content
type Validator struct{}

// NewValidator creates a new sitemap validator
func NewValidator() *Validator {
	return &Validator{}
}

// Validate checks if a sitemap is well-formed according to the sitemap protocol
func (v *Validator) Validate(sitemap *Sitemap) error {
	if sitemap == nil {
		return fmt.Errorf("sitemap is nil")
	}
	
	if len(sitemap.URLs) == 0 {
		return fmt.Errorf("sitemap contains no URLs")
	}
	
	for i, u := range sitemap.URLs {
		if err := v.ValidateURL(&u); err != nil {
			return fmt.Errorf("invalid URL at index %d: %w", i, err)
		}
	}
	
	return nil
}

// ValidateURL validates a single URL entry
func (v *Validator) ValidateURL(u *URL) error {
	if u.Loc == "" {
		return fmt.Errorf("missing <loc> element")
	}
	
	// Validate URL format
	if _, err := url.Parse(u.Loc); err != nil {
		return fmt.Errorf("invalid URL format: %w", err)
	}
	
	// Validate priority range if set
	if u.Priority < 0.0 || u.Priority > 1.0 {
		return fmt.Errorf("priority must be between 0.0 and 1.0, got %f", u.Priority)
	}
	
	// Validate changefreq values if set
	if u.ChangeFreq != "" {
		validFreqs := map[string]bool{
			"always":  true,
			"hourly":  true,
			"daily":   true,
			"weekly":  true,
			"monthly": true,
			"yearly":  true,
			"never":   true,
		}
		if !validFreqs[u.ChangeFreq] {
			return fmt.Errorf("invalid changefreq value: %s", u.ChangeFreq)
		}
	}
	
	return nil
}

