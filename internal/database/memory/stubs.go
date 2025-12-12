package memory

import (
	"context"

	"jonopens/sitemapper/internal/models"
	"jonopens/sitemapper/internal/repositories"
)

// Stub implementations for other repositories

type ReportRepository struct {
	db *Database
}

func (r *ReportRepository) Create(ctx context.Context, report *models.Report) error {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()
	r.db.reports[report.ID] = report
	return nil
}

func (r *ReportRepository) GetByID(ctx context.Context, id string) (*models.Report, error) {
	r.db.mu.RLock()
	defer r.db.mu.RUnlock()
	report, exists := r.db.reports[id]
	if !exists {
		return nil, ErrNotFound
	}
	return report, nil
}

func (r *ReportRepository) GetByUserID(ctx context.Context, userID string) ([]*models.Report, error) {
	r.db.mu.RLock()
	defer r.db.mu.RUnlock()
	var reports []*models.Report
	for _, report := range r.db.reports {
		if report.UserID == userID {
			reports = append(reports, report)
		}
	}
	return reports, nil
}

func (r *ReportRepository) List(ctx context.Context, filters repositories.ReportFilters) ([]*models.Report, error) {
	r.db.mu.RLock()
	defer r.db.mu.RUnlock()
	var reports []*models.Report
	for _, report := range r.db.reports {
		reports = append(reports, report)
	}
	return reports, nil
}

func (r *ReportRepository) Update(ctx context.Context, report *models.Report) error {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()
	r.db.reports[report.ID] = report
	return nil
}

func (r *ReportRepository) Delete(ctx context.Context, id string) error {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()
	delete(r.db.reports, id)
	return nil
}

type UserRepository struct {
	db *Database
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()
	r.db.users[user.ID] = user
	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	r.db.mu.RLock()
	defer r.db.mu.RUnlock()
	user, exists := r.db.users[id]
	if !exists {
		return nil, ErrNotFound
	}
	return user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	r.db.mu.RLock()
	defer r.db.mu.RUnlock()
	for _, user := range r.db.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, ErrNotFound
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()
	r.db.users[user.ID] = user
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()
	delete(r.db.users, id)
	return nil
}

type GroupingRepository struct {
	db *Database
}

func (r *GroupingRepository) Create(ctx context.Context, grouping *models.Group) error {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()
	r.db.groupings[grouping.ID] = grouping
	return nil
}

func (r *GroupingRepository) GetByID(ctx context.Context, id string) (*models.Group, error) {
	r.db.mu.RLock()
	defer r.db.mu.RUnlock()
	grouping, exists := r.db.groupings[id]
	if !exists {
		return nil, ErrNotFound
	}
	return grouping, nil
}

func (r *GroupingRepository) List(ctx context.Context) ([]*models.Group, error) {
	r.db.mu.RLock()
	defer r.db.mu.RUnlock()
	var groupings []*models.Group
	for _, grouping := range r.db.groupings {
		groupings = append(groupings, grouping)
	}
	return groupings, nil
}

func (r *GroupingRepository) Update(ctx context.Context, grouping *models.Group) error {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()
	r.db.groupings[grouping.ID] = grouping
	return nil
}

func (r *GroupingRepository) Delete(ctx context.Context, id string) error {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()
	delete(r.db.groupings, id)
	return nil
}

type ReportJobRepository struct {
	db *Database
}

func (r *ReportJobRepository) Create(ctx context.Context, job *models.ReportJob) error {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()
	r.db.jobs[job.ID] = job
	return nil
}

func (r *ReportJobRepository) GetByID(ctx context.Context, id string) (*models.ReportJob, error) {
	r.db.mu.RLock()
	defer r.db.mu.RUnlock()
	job, exists := r.db.jobs[id]
	if !exists {
		return nil, ErrNotFound
	}
	return job, nil
}

func (r *ReportJobRepository) List(ctx context.Context, filters repositories.JobFilters) ([]*models.ReportJob, error) {
	r.db.mu.RLock()
	defer r.db.mu.RUnlock()
	var jobs []*models.ReportJob
	for _, job := range r.db.jobs {
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (r *ReportJobRepository) Update(ctx context.Context, job *models.ReportJob) error {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()
	r.db.jobs[job.ID] = job
	return nil
}

func (r *ReportJobRepository) Delete(ctx context.Context, id string) error {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()
	delete(r.db.jobs, id)
	return nil
}

type ReleaseRepository struct {
	db *Database
}

func (r *ReleaseRepository) Create(ctx context.Context, release *models.Release) error {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()
	r.db.releases[release.ID] = release
	return nil
}

func (r *ReleaseRepository) GetByID(ctx context.Context, id string) (*models.Release, error) {
	r.db.mu.RLock()
	defer r.db.mu.RUnlock()
	release, exists := r.db.releases[id]
	if !exists {
		return nil, ErrNotFound
	}
	return release, nil
}

func (r *ReleaseRepository) List(ctx context.Context, filters repositories.ReleaseFilters) ([]*models.Release, error) {
	r.db.mu.RLock()
	defer r.db.mu.RUnlock()
	var releases []*models.Release
	for _, release := range r.db.releases {
		releases = append(releases, release)
	}
	return releases, nil
}

func (r *ReleaseRepository) Update(ctx context.Context, release *models.Release) error {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()
	r.db.releases[release.ID] = release
	return nil
}

func (r *ReleaseRepository) Delete(ctx context.Context, id string) error {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()
	delete(r.db.releases, id)
	return nil
}

