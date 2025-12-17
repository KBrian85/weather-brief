# Personal Daily Weather Brief (Go)

A lightweight Go project that can run as:
- a **CLI** (terminal tool), and
- a **Web App** (browser UI; deployable on Render).

It fetches **tomorrow’s forecast** for a city and generates a **practical daily weather brief** (commute + what to carry/wear).  
Designed as an “AI Essentials for everyday work” capstone: **data → interpretation → actionable summary**.

## What this project does (and does not do)
- ✅ Summarizes an **existing forecast** into an actionable brief.
- ✅ Provides **rule-based recommendations** (reliable by default).
- ✅ Optionally uses an LLM to **polish tone** while keeping facts locked to input.
- ❌ Does **not** attempt true meteorological prediction.

---

## Features
- City-based forecast retrieval via **Open-Meteo** (no API key required)
- Clean bullet-style brief (terminal output + browser view)
- Markdown export to `brief.md` (CLI)
- Offline demo mode (`--mock` or `BRIEF_USE_MOCK=1`)
- Optional CSV logging for visualizations (`--csv`)
- Optional LLM enhancement (OpenAI-compatible) with safe fallback
- Render-ready Web App (binds to `PORT`)

---

## Requirements
- Go 1.20+ (1.21+ recommended)
- Internet access for live forecast (not needed for mock mode)
- Optional plotting: Python 3 + `matplotlib`

---

## Quick start (CLI)
```bash
go mod tidy
go run ./cmd/brief --city "Nairobi"
```

Offline demo:
```bash
go run ./cmd/brief --city "Nairobi" --mock
```

CSV logging for charts:
```bash
go run ./cmd/brief --city "Nairobi" --csv
```

---

## Web App (Browser UI)

### Run locally
```bash
go mod tidy
go run ./cmd/web
```

Open:
- http://localhost:8080

Offline demo:
```bash
BRIEF_USE_MOCK=1 go run ./cmd/web
```

### Deploy to Render
This repo includes a `render.yaml` blueprint.

Render settings (if deploying via dashboard):

**Build Command**
```bash
go build -o app ./cmd/web
```

**Start Command**
```bash
./app
```

Tip: for a guaranteed demo without external calls, set in Render Environment:
- `BRIEF_USE_MOCK=1`

---

## Visualization (Charts / Results)

### A) Google Sheets (no coding)
1. Run several times:
   ```bash
   go run ./cmd/brief --city "Nairobi" --csv
   go run ./cmd/brief --city "Mombasa" --csv
   ```
2. Import `runs.csv` into Google Sheets (File → Import).
3. Insert a chart:
   - X-axis: `date`
   - Series: `temp_max_c` and/or `rain_probability_pct`

### B) Local Python plot
```bash
python3 scripts/plot_runs.py runs.csv
```

Outputs:
- `charts/temperature.png`
- `charts/rain_probability.png`

See: `docs/visualization.md` and `docs/results.md`

---

## Screenshots (examples)

### CLI command example
![CLI command example](docs/assets/code_snippet.png)

### Sample results (brief output)
![Sample brief output](docs/assets/sample_output.png)

### Visualization example (temperature)
![Temperature chart example](docs/assets/temperature_chart.png)

### Visualization example (rain probability)
![Rain probability chart example](docs/assets/rain_chart.png)

---

## Repository structure
```
weather-brief/
  cmd/brief/                 # CLI entrypoint
  cmd/web/                   # Web entrypoint (Render-ready)
  internal/weather/          # weather providers (Open-Meteo + mock)
  internal/brief/            # brief generation + CSV logging + optional enhancer
  internal/llm/              # minimal OpenAI-compatible client (optional)
  docs/visualization.md      # visualization guide (Sheets + Python)
  docs/results.md            # sample results page
  docs/assets/               # screenshot-style images + charts
  scripts/plot_runs.py       # local plotting script for runs.csv
  TOOLKIT.md                 # toolkit document (course submission)
  render.yaml                # Render blueprint (optional)
  README.md
  go.mod
  LICENSE
```

---

## License
MIT — see `LICENSE`.
