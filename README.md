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

You can configure your command with the following commands:

Add a new git user email:
```bash
gitworklog config set-name <user email address>
```
Add a new local repository path:	
```bash
gitworklog config add-repository <local repository path>
```
Delete from the list of local repository paths:	
```bash
gitworklog config delete-repository <local repository path>
```
List registered repository paths
```bash
	gitworklog config list-repositories
```
display registered git user name
```bash
gitworklog config get-name
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

**Today:**
```bash
gitworklog today
```

**Help:**
```bash
gitworklog
```

**Redirect output to a file:**

```bash
gitworklog 2025-06-05 2025-06-10 > myreport.txt
```

---

Enjoy!

Stay on top of your work and never miss logging your activity again.
