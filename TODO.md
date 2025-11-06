# Sitemapper - TODO List

This document tracks all implementation tasks needed to complete the Sitemapper application.

## 🎯 High Priority

### Database Implementation (PostgreSQL)
Complete the PostgreSQL repository implementations with actual SQL queries.

#### Entry Repository (`internal/database/postgres/entry_repository.go`)
- [ ] Implement Create - Insert entry into database
- [ ] Implement GetByID - Fetch entry by ID
- [ ] Implement List - List entries with filters
- [ ] Implement Update - Update existing entry
- [ ] Implement Delete - Remove entry
- [ ] Implement CountByType - Count entries by type

#### Report Repository (`internal/database/postgres/report_repository.go`)
- [ ] Implement Create - Insert report into database
- [ ] Implement GetByID - Fetch report by ID
- [ ] Implement GetByUserID - Fetch all reports for a user
- [ ] Implement List - List reports with filters
- [ ] Implement Update - Update existing report
- [ ] Implement Delete - Remove report

#### User Repository (`internal/database/postgres/user_repository.go`)
- [ ] Implement Create - Insert user into database
- [ ] Implement GetByID - Fetch user by ID
- [ ] Implement GetByEmail - Fetch user by email address
- [ ] Implement Update - Update existing user
- [ ] Implement Delete - Remove user

#### Category Repository (`internal/database/postgres/category_repository.go`)
- [ ] Implement Create - Insert category into database
- [ ] Implement GetByID - Fetch category by ID
- [ ] Implement List - List all categories
- [ ] Implement Update - Update existing category
- [ ] Implement Delete - Remove category

#### Report Job Repository (`internal/database/postgres/report_job_repository.go`)
- [ ] Implement Create - Insert job into database
- [ ] Implement GetByID - Fetch job by ID
- [ ] Implement List - List jobs with filters
- [ ] Implement Update - Update existing job
- [ ] Implement Delete - Remove job

#### Release Repository (`internal/database/postgres/release_repository.go`)
- [ ] Implement Create - Insert release into database
- [ ] Implement GetByID - Fetch release by ID
- [ ] Implement List - List releases with filters
- [ ] Implement Update - Update existing release
- [ ] Implement Delete - Remove release

---

## 🔧 Services Layer

### Sitemap Service (`internal/services/sitemap_service.go`)
- [ ] Implement ProcessSitemap
  - Parse sitemap using parser
  - Categorize entries by URL segments
  - Store entries in database
  - Handle sitemap index vs regular sitemap
- [ ] Implement CompareSitemaps
  - Fetch old and new sitemaps
  - Compare URLs, lastmod, priority, etc.
  - Generate changeset
  - Store comparison results

### Report Service (`internal/services/report_service.go`)
- [ ] Implement GenerateReport
  - Fetch sitemap entries
  - Calculate metrics (entry count, category count, etc.)
  - Check for invalid entries
  - Generate and store report

### Job Service (`internal/services/job_service.go`)
- [ ] Implement ScheduleJob
  - Validate cron expression
  - Schedule recurring sitemap checks
  - Store job configuration
- [ ] Implement ExecuteJob
  - Fetch job configuration
  - Execute sitemap processing
  - Update job status
  - Handle errors and retries

### Category Service (`internal/services/category_service.go`)
- [ ] Implement CategorizeURLs
  - Parse URL path segments
  - Group URLs by common patterns
  - Create/update categories
  - Return categorization mapping

---

## 🌐 HTTP Handlers

### Sitemap Handler (`internal/handlers/sitemap_handler.go`)
- [ ] Implement Upload
  - Parse multipart form data or JSON
  - Validate sitemap XML
  - Extract user ID from auth context
  - Call sitemap service to process
  - Return sitemap ID and status
- [ ] Implement Get
  - Fetch sitemap by ID
  - Check authorization
  - Return sitemap details
- [ ] Implement List
  - Parse pagination parameters
  - Fetch user's sitemaps
  - Return paginated list

### Report Handler (`internal/handlers/report_handler.go`)
- [ ] Get user ID from auth context instead of query parameter
- [ ] Implement Generate
  - Parse request parameters
  - Trigger async report generation
  - Return job ID

### User Handler (`internal/handlers/user_handler.go`)
- [ ] Implement Create
  - Parse request body
  - Validate user data
  - Hash password (if applicable)
  - Create user in database
  - Return user ID

### Auth Middleware (`internal/handlers/middleware/auth.go`)
- [ ] Implement authentication logic
  - Extract token from Authorization header
  - Validate JWT token
  - Set user context
  - Handle unauthorized requests

---

## 📦 Package Utilities

### HTTP Client (`pkg/http/client.go`)
- [ ] Implement Post method with retry logic
  - Handle request body serialization
  - Implement exponential backoff
  - Return response or error after retries

### Scheduler (`pkg/scheduler/scheduler.go`)
- [ ] Implement actual cron scheduling
  - Integrate with robfig/cron or similar library
  - Parse cron expressions
  - Execute jobs at scheduled times
- [ ] Implement Start method
  - Initialize cron scheduler
  - Start background job runner
- [ ] Implement Stop method
  - Stop all running jobs
  - Cleanup resources
  - Wait for graceful shutdown

---

## 🗄️ Database & Migrations

### Schema Creation
- [ ] Create PostgreSQL migration files
  - users table
  - entries table
  - reports table
  - categories table
  - entry_categories join table
  - report_jobs table
  - releases table
  - changesets table
- [ ] Create MySQL migration files (same schema)
- [ ] Create SQLite migration files (same schema)

### Migration System (`Makefile`)
- [ ] Add migration command to Makefile
  - Use golang-migrate or pressly/goose
  - Support up/down migrations
- [ ] Add migration rollback command

---

## 🔐 Authentication & Authorization

### User Authentication
- [ ] Implement JWT token generation
- [ ] Implement JWT token validation
- [ ] Add password hashing (bcrypt)
- [ ] Create login endpoint
- [ ] Create signup endpoint
- [ ] Add refresh token support

### Authorization
- [ ] Implement resource ownership checks
- [ ] Add role-based access control (if needed)
- [ ] Secure all endpoints with auth middleware

---

## ✅ Testing

### Unit Tests
- [ ] Write tests for all services using memory database
- [ ] Write tests for pkg utilities
- [ ] Write tests for handlers with mock services
- [ ] Achieve >80% code coverage

### Integration Tests
- [ ] Test complete flow: upload → process → report
- [ ] Test database adapter switching
- [ ] Test error scenarios

---

## 🚀 Deployment & DevOps

### Database Adapters
- [ ] Complete MySQL repository implementations
  - Copy PostgreSQL implementations
  - Adjust for MySQL-specific SQL syntax
- [ ] Complete SQLite repository implementations
  - Copy PostgreSQL implementations
  - Adjust for SQLite-specific SQL syntax

### Configuration
- [ ] Add environment variable support
- [ ] Add configuration validation
- [ ] Document all configuration options

### Monitoring & Logging
- [ ] Add structured logging (logrus or zap)
- [ ] Add request ID tracking
- [ ] Add metrics/instrumentation (Prometheus)
- [ ] Add health check endpoint

---

## 📚 Documentation

- [ ] Add API documentation (OpenAPI/Swagger)
- [ ] Add example requests/responses
- [ ] Document authentication flow
- [ ] Add deployment guide
- [ ] Add development setup guide
- [ ] Add troubleshooting guide

---

## 🎨 Nice to Have

### Features
- [ ] Add webhook support for job completion
- [ ] Add email notifications
- [ ] Add Slack integration
- [ ] Add sitemap comparison visualization
- [ ] Add export functionality (CSV, JSON)
- [ ] Add batch sitemap upload

### Performance
- [ ] Add database connection pooling
- [ ] Add Redis caching for reports
- [ ] Add request rate limiting
- [ ] Optimize large sitemap processing

### Developer Experience
- [ ] Add hot-reload for development (air)
- [ ] Add database seeding script
- [ ] Add sample data generator
- [ ] Add API client library

---

## 📝 Notes

- Start with PostgreSQL implementation as the reference implementation
- MySQL and SQLite adapters can largely copy PostgreSQL with syntax adjustments
- Use the memory adapter for all unit tests
- Consider adding integration tests with docker-compose

## 🏁 Getting Started

Recommended order to tackle these TODOs:

1. **Database Schema & Migrations** - Create tables first
2. **PostgreSQL Repositories** - Implement data access layer
3. **Services Layer** - Implement business logic
4. **Handlers** - Implement HTTP endpoints
5. **Authentication** - Secure the application
6. **Testing** - Ensure quality
7. **Documentation** - Help others use it

Good luck! 🚀

