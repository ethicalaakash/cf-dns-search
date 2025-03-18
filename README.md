# Cloudflare DNS Record Search

A simple command-line tool to search for DNS records across all your Cloudflare zones that contain a specific domain or URL.

## Description

This tool connects to Cloudflare's API and:
- Fetches all zones in your account
- Lists DNS records in each zone
- Filters records containing the specified URL/domain
- Displays matching records

## Prerequisites

- Go 1.16 or higher
- Cloudflare API Token with DNS read permissions
- Access to Cloudflare zones you want to search

## Installation

```bash
git clone git@github.com:ethicalaakash/cf-dns-search.git
cd cf-dns-search
go mod download
```

## Configuration

1. Create a Cloudflare API token at https://dash.cloudflare.com/profile/api-tokens
   - Ensure it has `Zone.DNS` read permissions

2. Set your API token as an environment variable:
```bash
export CLOUDFLARE_API_TOKEN="your-api-token-here"
```

## Build

To compile the binary into the `bin` directory:

```bash
# Create bin directory if it doesn't exist
mkdir -p bin

# Build the binary
go build -o bin/cf-dns-search main.go

# Make it executable
chmod +x bin/cf-dns-search
```

You can then run the compiled binary:

```bash
# Run with default settings
./bin/cf-dns-search

# Search for specific domain
./bin/cf-dns-search -url="example.com"
```

## Flags

- `-url string`: Target URL to search in DNS records.
