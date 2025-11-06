package memory

import (
	"context"

	"jonopens/sitemapper/internal/models"
	"jonopens/sitemapper/internal/repositories"
)

type EntryRepository struct {
	db *Database
}

func (r *EntryRepository) Create(ctx context.Context, entry *models.Entry) error {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()
	r.db.entries[entry.ID] = entry
	return nil
}

func (r *EntryRepository) GetByID(ctx context.Context, id string) (*models.Entry, error) {
	r.db.mu.RLock()
	defer r.db.mu.RUnlock()
	entry, exists := r.db.entries[id]
	if !exists {
		return nil, ErrNotFound
	}
	return entry, nil
}

func (r *EntryRepository) List(ctx context.Context, filters repositories.EntryFilters) ([]*models.Entry, error) {
	r.db.mu.RLock()
	defer r.db.mu.RUnlock()
	
	var entries []*models.Entry
	for _, entry := range r.db.entries {
		// Apply filters if needed
		entries = append(entries, entry)
	}
	return entries, nil
}

func (r *EntryRepository) Update(ctx context.Context, entry *models.Entry) error {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()
	if _, exists := r.db.entries[entry.ID]; !exists {
		return ErrNotFound
	}
	r.db.entries[entry.ID] = entry
	return nil
}

func (r *EntryRepository) Delete(ctx context.Context, id string) error {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()
	delete(r.db.entries, id)
	return nil
}

func (r *EntryRepository) CountByType(ctx context.Context, entryType models.EntryType) (int, error) {
	r.db.mu.RLock()
	defer r.db.mu.RUnlock()
	
	count := 0
	for _, entry := range r.db.entries {
		if entry.Type == entryType {
			count++
		}
	}
	return count, nil
}

