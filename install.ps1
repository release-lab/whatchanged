#!/usr/bin/env pwsh
# inherit from https://deno.land/x/install@v0.1.4/install.ps1
# Copyright 2018 the Deno authors. All rights reserved. MIT license.

$ErrorActionPreference = 'Stop'

if ($v) {
  $Version = "v${v}"
}
if ($args.Length -eq 1) {
  $Version = $args.Get(0)
}

if (Test-Path C:\Windows\SysNative) {
  $arch = "amd64"
} else {
  $arch = "386"
}

$BinDir = "$Home\bin"
$is64Bit = Test-Path C:\Windows\SysNative
$WhatchangedTarGz = "$BinDir\whatchanged.tar.gz"
$WhatchangedExe = "$BinDir\whatchanged.exe"
$Target = "windows_$arch"

# GitHub requires TLS 1.2
[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12

$ResourceUri = if (!$Version) {
  "https://github.com/whatchanged-community/whatchanged/releases/latest/download/whatchanged_${Target}.tar.gz"
} else {
  "https://github.com/whatchanged-community/whatchanged/releases/download/${Version}/whatchanged_${Target}.tar.gz"
}

if (!(Test-Path $BinDir)) {
  New-Item $BinDir -ItemType Directory | Out-Null
}

Invoke-WebRequest $ResourceUri -OutFile $WhatchangedTarGz -UseBasicParsing

if (Get-Command Expand-Archive -ErrorAction SilentlyContinue) {
  Expand-Archive $WhatchangedTarGz -Destination $BinDir -Force
} else {
  if (Test-Path $WhatchangedExe) {
    Remove-Item $WhatchangedExe
  }
  Add-Type -AssemblyName System.IO.Compression.FileSystem
  [IO.Compression.ZipFile]::ExtractToDirectory($WhatchangedTarGz, $BinDir)
}

Remove-Item $WhatchangedTarGz

$User = [EnvironmentVariableTarget]::User
$Path = [Environment]::GetEnvironmentVariable('Path', $User)
if (!(";$Path;".ToLower() -like "*;$BinDir;*".ToLower())) {
  [Environment]::SetEnvironmentVariable('Path', "$Path;$BinDir", $User)
  $Env:Path += ";$BinDir"
}

Write-Output "Whatchanged was installed successfully to $WhatchangedExe"
Write-Output "Run 'whatchanged --help' to get started"