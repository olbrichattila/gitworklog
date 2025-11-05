# Git Work Log

Ever forgotten to fill in your timesheet, and can't quite remember what you worked on? Or maybe you have multiple local branches across different repositories? This tool is here to help.

**Git Work Log** generates a report of your work activity, including repositories, branches, and commit messages — grouped by date and ordered by time. You don’t even need Git installed locally (though most of us do).

--- 

## Install:

```bash
go install github.com/olbrichattila/gitworklog@latest
```

---

## Features
- Generates a work log from your Git repositories.
- Supports multiple repositories and branches.
- Grouped by date and ordered by commit time.
- Fetch a report for a single day or a date range.
- Output can be redirected to a file for further use.

---

## Configuration

Create a config.yaml file in the same directory as the executable:

```yaml
username: "itsme@gmail.com"
repositories:
  - path: "/home/johndoe/myrepo1"
  - path: "/home/johndoe/myrepo2"
```

---

## Usage
```
gitworklog <fromDate> [toDate]
```
### Examples

**Date range:**
```bash
gitworklog 2025-06-05 2025-06-10
```

**Single date:**
```bash
gitworklog 2025-06-05
```

**Redirect output to a file:**

```bash
gitworklog 2025-06-05 2025-06-10 > myreport.txt
```

---

Enjoy!

Stay on top of your work and never miss logging your activity again.
