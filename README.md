# Personal Daily Weather Brief (Go)

A lightweight CLI that fetches tomorrow's forecast for a city and produces a practical daily brief.
Designed as a simple AI-for-everyday-work capstone.

## What this project demonstrates
- Clean Go project structure
- Calling a public weather API (Open-Meteo)
- Simple rule-based recommendations
- Optional LLM polishing layer (can be disabled)
- Markdown output for easy sharing

## Features
- `--city` required
- Optional preferences:
  - `--commute` (walk|drive|motorbike|public)
  - `--sensitivity` (cold|normal|heat)
  - `--tone` (concise|friendly|professional)
- Outputs:
  - Console text
  - `brief.md` in the current directory (default)

## Requirements
- Go 1.20+ (1.21+ recommended)
- Internet access for live API calls

## Quick start
```bash
git clone <your-repo-url>
cd weather-brief
go mod tidy
go run ./cmd/brief --city "Nairobi"
```

## Example with preferences
```bash
go run ./cmd/brief --city "Nairobi" --commute motorbike --sensitivity normal --tone friendly
```

## LLM optional
By default, the tool uses **rule-based** briefing text to ensure it works without keys.

To enable LLM enhancement, set:
```bash
export BRIEF_USE_LLM=1
export BRIEF_LLM_PROVIDER=openai_compatible
export BRIEF_LLM_API_KEY="YOUR_KEY"
export BRIEF_LLM_BASE_URL="https://api.openai.com/v1"
export BRIEF_LLM_MODEL="gpt-4o-mini"
```

If these are not set, the app will gracefully fall back to rule-only mode.

## Output
The app writes `brief.md`. You can change this with:
```bash
--out mybrief.md
```

## Notes on accuracy
This tool does **not** create new meteorological predictions.
It turns **existing forecasts** into a practical, personalized briefing.

## License
MIT
