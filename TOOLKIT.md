# Personal Daily Weather Brief in Go (Toolkit Document)

## 1. Title & Objective
**Technology Choice:** Go (Golang) + Open-Meteo weather API (optional LLM tone polishing).  
**Justification:** Go is ideal for lightweight tools with strong HTTP/JSON support and easy deployment.  
**End Goal:** A minimal working example that generates an actionable “tomorrow” weather brief.

## 2. Quick Summary of the Technology
**Definition:** Go is a compiled, statically typed language designed for simplicity and reliability.  
**Use Cases:** CLIs, backend APIs, automation, cloud-native apps.  
**Real-World Example:** Many infrastructure/developer tools are built in Go due to speed and portability.

## 3. System Requirements
- OS: Windows/macOS/Linux
- Go 1.20+ (1.21+ recommended)
- Editor: VS Code + Go extension
- Optional plotting: Python 3 + matplotlib

## 4. Installation & Setup Instructions
CLI:
```bash
go mod tidy
go run ./cmd/brief --city "Nairobi"
```
Web:
```bash
go run ./cmd/web
```

## 5. Minimal Working Example (MWE)
Run the CLI command above; expected output is a bullet brief plus `brief.md`.

## 6. AI Prompt Journal (optional LLM)
Prompt style:
- “Use ONLY the JSON input. Do not invent numbers or conditions. Return 5–7 bullet points plus one short closing line.”
Reflection: strict constraints reduced hallucinations and improved reliability.

## 7. Common Issues & Fixes
- City not found: try “City, Country”
- Network issues: use mock mode
- Missing Go: install and verify `go version`

## 8. References
- Go docs (official)
- Open-Meteo docs (forecast + geocoding)
