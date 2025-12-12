# Sitemapper

A tool for long-running XML sitemap comparison. A tool I wish I had when I was doing SEO fulltime.

## Project Goals

- Keep track of sitemap changes, either on a schedule or ad hoc
- Allow for manual source file upload or fetch from a URL
- Sample large sitemap indexes

## Personal Goals

- Leverage LLM tooling to improve my efficacy

## Models
- Users
- Reports
- Releases (time markers)
- Groupings
- Diffs (maybe?)
- Jobs (long-running, async tasks)
- Schedules

## Flow

- user gives either a url or a sitemap file
- they can name it but by default it comes from the URL+datetime
- ask user if they want to perform liveness check
  - if yes, provide UA for allowlist and fetch headlessly to fetch only status
  - if no, pass an option to not fetch each URL for liveness
- file is parsed
- on job completion, user is notified
- user can review groupings, rename them, etc.
- user can add new sitemap task via a schedule
- user can schedule a reminder via SMS or email (or slack integration)?
- user can review the report
- user can add important dates (releases)
- user can review charts showing counts and filter charts by groupings

## Sitemap Processing

- determine if a file is a sitemap index or a simple sitemap based on `<urlset>` or `<sitemapindex>`
  - a `<url>` or `<sitemap>` with no data is an invalid entry
  - a `<url>` or `<sitemap>` with no `<loc>` is an invalid entry
- if entries, attempt to group by URL segment, if they exist and enqueue as batches
- if sitemaps, attempt to group by sitemap URL segments, if they exist
// more details here breaking down the process