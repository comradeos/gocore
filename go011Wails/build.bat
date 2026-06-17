@echo off
setlocal

where wails >nul 2>nul
if %errorlevel%==0 (
    set "WAILS_CMD=wails"
) else (
    for /f "delims=" %%i in ('go env GOPATH') do set "GOPATH_VALUE=%%i"
    set "WAILS_CMD=%GOPATH_VALUE%\bin\wails.exe"
)

if not exist "%WAILS_CMD%" (
    echo Wails CLI not found.
    echo Install it with:
    echo   go install github.com/wailsapp/wails/v2/cmd/wails@latest
    exit /b 1
)

echo Building Wails application...
call "%WAILS_CMD%" build -clean
if errorlevel 1 exit /b 1

echo Build completed.
