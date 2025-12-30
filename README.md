# Sitemapper

A powerful CLI tool for tracking and comparing XML sitemaps over time. Built for SEO professionals and developers who need to monitor sitemap changes.

## Features

- 🔍 **Parse & Validate** - Parse sitemaps from files or URLs with validation
- 📊 **Compare Sitemaps** - Diff two sitemaps to see added, removed, and unchanged URLs
- 💾 **Track History** - Save sitemap snapshots to database for historical comparison
- 📈 **Reports** - Generate and view detailed reports with statistics
- 🏷️ **Groupings** - Organize URLs into logical groups
- 🎨 **Multiple Output Formats** - JSON, table, or plain text output
- 🔄 **Interactive Mode** - REPL shell for running multiple commands
- 🗃️ **Database Support** - PostgreSQL, MySQL, SQLite, or in-memory storage

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/yourusername/sitemapper.git
cd sitemapper

# Build the binary
make build

# The binary will be in ./bin/sitemapper
./bin/sitemapper --help
```

### Using Docker

```bash
# Build the image
docker build -t sitemapper .

# Run with docker
docker run -it --rm sitemapper parse https://example.com/sitemap.xml
```

## Quick Start

### 1. Parse a Sitemap

Parse and validate a sitemap from a URL or file:

```bash
# From URL
sitemapper parse https://example.com/sitemap.xml --validate --show-stats

# From local file
sitemapper parse ./sitemap.xml --validate
```

### 2. Track Sitemap Changes

Save a sitemap snapshot to track changes over time:

```bash
# Track a sitemap with a custom name
sitemapper track https://example.com/sitemap.xml --name "example-v1"

# Track from local file
sitemapper track ./sitemap.xml --name "local-snapshot"
```

### 3. Compare Sitemaps

Compare two sitemaps to see differences:

```bash
# Compare two files
sitemapper compare sitemap-old.xml sitemap-new.xml

# Compare tracked snapshots
sitemapper compare report-id-1 report-id-2

# Show unchanged URLs as well
sitemapper compare sitemap1.xml sitemap2.xml --show-unchanged
```

### 4. View Reports

List and view saved reports:

```bash
# List all reports
sitemapper report list

# Get details for a specific report
sitemapper report get <report-id>
```

## Configuration

Create a config file at `configs/config.yaml` or `~/.sitemapper/config.yaml`:

```yaml
# Database configuration
database_type: postgres
database_url: postgresql://localhost:5432/sitemapper?sslmode=disable

# CLI settings
default_user_id: default
output_format: table  # json, table, or text
color_output: true

# Application settings
max_upload_size: 104857600  # 100MB
enable_liveness: false
worker_count: 5
environment: development
```

You can also use environment variables with the `SITEMAPPER_` prefix:

```bash
export SITEMAPPER_DATABASE_TYPE=postgres
export SITEMAPPER_DATABASE_URL=postgresql://localhost:5432/sitemapper
export SITEMAPPER_OUTPUT_FORMAT=json
```

## Usage

### Parse Command

Parse and validate sitemaps:

```bash
# Basic parsing
sitemapper parse <url-or-file>

# With validation
sitemapper parse <url-or-file> --validate

# Show statistics
sitemapper parse <url-or-file> --show-stats

# JSON output
sitemapper parse <url-or-file> --format json
```

### Track Command

Save sitemap snapshots to database:

```bash
# Track a sitemap
sitemapper track <url-or-file> --name <snapshot-name>

# Specify user ID
sitemapper track <url-or-file> --name <name> --user-id <user-id>
```

### Compare Command

Compare two sitemaps:

```bash
# Compare files or URLs
sitemapper compare <source1> <source2>

# Compare tracked snapshots (use report IDs)
sitemapper compare <report-id-1> <report-id-2>

# Show unchanged URLs
sitemapper compare <source1> <source2> --show-unchanged

# JSON output
sitemapper compare <source1> <source2> --format json
```

### Report Commands

Manage reports:

```bash
# List all reports
sitemapper report list

# List reports with limit
sitemapper report list --limit 20

# Get specific report
sitemapper report get <report-id>

# JSON output
sitemapper report get <report-id> --format json
```

### Grouping Commands

Manage URL groupings:

```bash
# List groupings
sitemapper grouping list

# Create a grouping
sitemapper grouping create --name "Blog Posts" --description "All blog URLs"
```

### Interactive Mode

Launch an interactive shell:

```bash
sitemapper interactive
```

In interactive mode:

```
sitemapper> parse https://example.com/sitemap.xml
sitemapper> track ./sitemap.xml --name "snapshot-1"
sitemapper> report list
sitemapper> exit
```

## Examples

### Track sitemap changes over time

```bash
# Day 1: Save initial snapshot
sitemapper track https://example.com/sitemap.xml --name "week-1"

# Day 7: Save another snapshot
sitemapper track https://example.com/sitemap.xml --name "week-2"

# Compare the two snapshots
sitemapper report list
# Note the report IDs from the list output

sitemapper compare <report-id-week-1> <report-id-week-2>
```

### Validate multiple sitemaps

```bash
# Parse and validate each sitemap
for sitemap in sitemaps/*.xml; do
  echo "Validating $sitemap"
  sitemapper parse "$sitemap" --validate
done
```

### Monitor production sitemap

```bash
# Create a cron job to track daily
0 2 * * * /usr/local/bin/sitemapper track https://example.com/sitemap.xml --name "daily-$(date +\%Y\%m\%d)"
```

## Database Setup

### PostgreSQL

```bash
# Create database
createdb sitemapper

# Update config
database_type: postgres
database_url: postgresql://localhost:5432/sitemapper?sslmode=disable
```

### SQLite

```bash
# Config for SQLite
database_type: sqlite
database_url: ./data/sitemapper.db
```

### In-Memory (for testing)

```bash
# Config for in-memory database
database_type: memory
database_url: ""
```

## Development

### Building

```bash
# Build the binary
make build

# Run tests
make test

# Run with coverage
make test-coverage

# Format code
make fmt

# Lint code
make lint

# Tidy dependencies
make tidy
```

### Running Locally

```bash
# Run directly with Go
make run

# Or use the built binary
./bin/sitemapper --help
```

### Docker Development

```bash
# Start database only
docker-compose up -d postgres

# Build and run CLI
docker build -t sitemapper .
docker run -it --rm sitemapper
```

## Project Structure

```
sitemapper/
├── cmd/
│   └── cli/           # CLI entry point
├── configs/           # Configuration files
├── internal/
│   ├── cli/          # CLI commands
│   │   ├── root.go   # Root command
│   │   ├── parse.go  # Parse command
│   │   ├── compare.go
│   │   ├── track.go
│   │   ├── report.go
│   │   ├── grouping.go
│   │   └── interactive.go
│   │   └── output/   # Output formatters
│   ├── config/       # Configuration management
│   ├── database/     # Database implementations
│   ├── models/       # Domain models
│   ├── repositories/ # Data access layer
│   └── services/     # Business logic
├── pkg/
│   ├── http/         # HTTP client utilities
│   └── sitemap/      # Sitemap parsing & validation
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
```

## Architecture

The application follows a clean architecture pattern:

- **CLI Layer**: Command handlers and user interaction
- **Service Layer**: Business logic for sitemap processing
- **Repository Layer**: Database abstraction
- **Package Layer**: Reusable utilities (parsers, validators, HTTP client)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[Your License Here]

## Credits

Built with:
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Viper](https://github.com/spf13/viper) - Configuration management
- [go-prompt](https://github.com/c-bata/go-prompt) - Interactive mode
- [tablewriter](https://github.com/olekukonko/tablewriter) - Table output
- [color](https://github.com/fatih/color) - Colored terminal output

## Roadmap

- [ ] Implement database repository methods
- [ ] Add sitemap index support for tracking
- [ ] URL liveness checking
- [ ] Advanced URL grouping with patterns
- [ ] Export reports to CSV/PDF
- [ ] Webhooks for automated tracking
- [ ] Web UI (optional future enhancement)
