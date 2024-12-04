# Accept arguments YEAR and DAY
param([string]$Year, [string]$Day)

if (-not $Year -or -not $Day) {
  Write-Host "Usage: .\script.ps1 <YEAR> <DAY>" -ForegroundColor Red
  exit 1
}

# Functions
function Write-Grey {
  param(
    [string]$Message
  )
  Write-Host $Message -ForegroundColor Gray
}

function Generate-Template {
  @"
package main

import (
    "github.com/jpillora/puzzler/harness/aoc"
)

func main() {
    aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
    // when you're ready to do part 2, remove this "not implemented" block
    if part2 {
        return "not implemented"
    }
    // solve part 1 here
    return 42
}
"@
}

# Pad DAY to 2 digits
$Day = $Day.PadLeft(2, '0')
$Dir = Join-Path -Path "." -ChildPath "$Year\$Day"

# Create missing directories as needed
if (-not (Test-Path -Path $Dir)) {
  New-Item -ItemType Directory -Path $Dir | Out-Null
  Write-Grey "Created directory $Dir"
}

# Create missing files as needed
$FilePath = Join-Path -Path $Dir -ChildPath "code.go"
if (-not (Test-Path -Path $FilePath)) {
  Generate-Template | Out-File -FilePath $FilePath -Encoding UTF8
  Write-Grey "Created file code.go"
}

# Go run
Set-Location -Path $Dir
go run code.go
