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
$TemplatePath = Join-Path -Path "." -ChildPath "template.go"
if (-not (Test-Path -Path $FilePath)) {
  Get-Content -Path $TemplatePath | Out-File -FilePath $FilePath -Encoding UTF8
  Write-Grey "Created file code.go"
}

# Go run
Set-Location -Path $Dir
go run code.go
