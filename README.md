# JsonX | CLI JSON Toolkit

![JSON Tool](https://img.shields.io/badge/Purpose-JSON%20Toolkit-yellow?style=for-the-badge)
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

---

## Command-line JSON Power

A fast, intuitive CLI tool for querying, transforming, and manipulating JSON data directly in your terminal.

**Why JsonX?**
- No more juggling between jq, python, and awk
- Simple, intuitive syntax
- Pipes perfectly with other CLI tools
- Zero learning curve

---

## Features

- 🔍 **Query** - Extract data with JSONPath-like syntax
- 🎨 **Format** - Pretty print or minify JSON
- 🔄 **Transform** - Reshape JSON structures
- 📊 **Validate** - Check JSON syntax
- ⚡ **Fast** - Built for speed in Go

---

## Installation

```bash
git clone https://github.com/simplestar-992/jsonx.git
cd jsonx
go build -o jsonx -ldflags="-s -w"
```

---

## Usage

```bash
# Pretty print
echo '{"name":"test"}' | ./jsonx

# Query data
./jsonx -q '.users[0].name' data.json

# Extract array
./jsonx -q '.items[*].id' data.json

# Filter and transform
./jsonx -q '.data[?@.active==true]' data.json

# Minify
./jsonx --minify data.json

# Validate
./jsonx --validate data.json
```

---

## Examples

```bash
# Extract all user emails
curl -s https://api.example.com/users | ./jsonx -q '.users[*].email'

# Find active items
cat data.json | ./jsonx -q '.items[?status=="active"]'

# Get nested values
./jsonx -q '.config.database.host' config.json

# Combine with other tools
cat data.json | ./jsonx -q '.users' | jq '.[] | .name'
```

---

## License

MIT © 2024 [simplestar-992](https://github.com/simplestar-992)
