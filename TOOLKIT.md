# Personal Daily Weather Brief in Go
**Toolkit Document (Capstone)**

## 1. Title & Objective
**Technology chosen:** Go (Golang) for a lightweight CLI, plus a public weather API (Open-Meteo) and an *optional* LLM enhancement layer.

**Why Go?**
- Great for small, fast utilities used in daily work.
- Simple deployment (single binary).
- Strong standard library for HTTP + JSON.
- Good for clean, maintainable capstones.

**End goal**
Build a tool that converts tomorrow’s forecast for a chosen city into a **short, practical daily weather brief** with simple recommendations for clothing and commuting.

---

## 2. Quick Summary of the Technology
**Go** is a compiled programming language designed for simplicity, performance, and productivity. It is widely used for backend services, DevOps tooling, and command-line utilities.

**Where it’s used**
- Cloud services
- APIs and microservices
- Developer and productivity tooling

**Real-world example**
Many modern infrastructure tools and cloud-native systems are built with Go because of its speed and easy deployment.

---

## 3. System Requirements
**OS**
- Windows 10/11, macOS, or Linux

**Tools**
- Go 1.20+ (1.21+ recommended)
- Git (optional but useful)
- A code editor (VS Code, GoLand, etc.)

**Packages**
- No external packages required for MVP (standard library only)

**Internet**
- Required for live weather API calls
- Optional for LLM enhancement

---

## 4. Installation & Setup Instructions

### 4.1 Install Go
Download and install Go from the official site for your OS.
Verify installation:
```bash
go version
```

### 4.2 Get the project
If using a ZIP:
1. Unzip the folder.
2. Open a terminal inside the project.

If using Git:
```bash
git clone <your-repo-url>
cd weather-brief
```

### 4.3 Install dependencies (none external)
```bash
go mod tidy
```

### 4.4 Run the CLI
```bash
go run ./cmd/brief --city "Nairobi"
```

### 4.5 Offline demo mode
```bash
go run ./cmd/brief --city "Nairobi" --mock
```

### 4.6 Optional LLM enhancement
The project runs without any LLM keys.  
To enable an OpenAI-compatible LLM endpoint:
```bash
export BRIEF_USE_LLM=1
export BRIEF_LLM_PROVIDER=openai_compatible
export BRIEF_LLM_API_KEY="YOUR_KEY"
export BRIEF_LLM_BASE_URL="https://api.openai.com/v1"
export BRIEF_LLM_MODEL="gpt-4o-mini"
```

---

## 5. Minimal Working Example

### 5.1 Code (main entry)
File: `cmd/brief/main.go`

Key idea:
- Parse flags
- Fetch forecast
- Produce a rules-based brief
- Optionally enhance with LLM
- Write `brief.md`

### 5.2 Run
```bash
go run ./cmd/brief --city "Nairobi" --commute motorbike --sensitivity normal --tone friendly
```

### 5.3 Expected console output (example)
```
Daily Weather Brief — Nairobi
- Date: Tue, 09 Dec 2025
- Conditions: Partly cloudy
- Temperature: 19–27°C
- Rain chance (max): 55%
- Wind (max): 18 kph
- Carry an umbrella or light rain jacket.
- Riding note: consider waterproof gear and reduced speed on wet roads.
```

### 5.4 Expected file output
A Markdown file named `brief.md` is created in the current directory.

---

## 6. AI Prompt Journal

### Prompt 1 — Initial brief template
**Prompt**
“Generate a daily weather brief in 5 bullets using tomorrow’s forecast.”

**Response summary**
The AI produced a nice structure but sometimes added extra numbers.

**Evaluation**
Useful for structure, but needed strict guardrails.

---

### Prompt 2 — Add anti-hallucination rules
**Prompt**
“Use ONLY the JSON input. Do not invent numbers or conditions. Return 5–7 bullets.”

**Response summary**
The output became more factual and aligned with the structured data.

**Evaluation**
This prompt style was adopted into the code.

---

### Prompt 3 — Tone testing
**Prompt**
“Produce the same brief in a friendly tone. Keep facts unchanged.”

**Response summary**
Tone changed appropriately while preserving the data.

**Evaluation**
Confirmed that tone could be safely separated from data.

---

### Reflection: How AI helped
- Accelerated the initial output format design.
- Helped refine user-friendly recommendation language.
- Made it easy to generate multiple example outputs for the report.

---

## 7. Common Issues & Fixes

### 7.1 “--city is required”
**Cause**
Flag not provided.

**Fix**
```bash
go run ./cmd/brief --city "Nairobi"
```

---

### 7.2 “no results for city”
**Cause**
City name misspelled or too vague.

**Fix**
Try a more specific name:
```bash
--city "Nairobi, Kenya"
```

---

### 7.3 API network timeouts
**Cause**
Unstable internet.

**Fix**
Use mock mode:
```bash
--mock
```

---

### 7.4 LLM errors
**Cause**
Missing env vars or invalid key.

**Fix**
Either:
- Set the required env vars, or
- Remove:
```bash
export BRIEF_USE_LLM=1
```
The app will fall back to rule-only mode.

---

## 8. References
- Go documentation (official)
- Open-Meteo documentation (forecast + geocoding)
- General prompt engineering best practices for accuracy and non-hallucination
- Workplace AI usage guidelines (human-in-the-loop review)

---

## Appendix: What makes this a good “AI Essentials” capstone?
This project focuses on **AI as a productivity layer**:
- Turning raw information into actionable decisions.
- Using a safe, minimal AI loop.
- Demonstrating rule-first reliability with optional AI polish.
