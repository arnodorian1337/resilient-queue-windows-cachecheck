param(
    [Parameter(ValueFromRemainingArguments = $true)]
    [string[]] $RemainingArguments
)

$cacheDirectory = Join-Path $env:LOCALAPPDATA "ResilientQueue"
$cacheFile = Join-Path $cacheDirectory "compatibility.cache"
New-Item -ItemType Directory -Path $cacheDirectory -Force | Out-Null
$content = @(
    "state=current"
    "checked_at=$([DateTime]::UtcNow.ToString('o'))"
    "args=$($RemainingArguments -join ' ')"
) -join "`r`n"
[IO.File]::WriteAllText($cacheFile, $content + "`r`n")
Write-Output "Compatibility cache: current"
