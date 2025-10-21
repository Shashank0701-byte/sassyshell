@echo off
setlocal enabledelayedexpansion

echo 🚀 Installing SassyShell TUI...

REM Check if Go is installed
where go >nul 2>nul
if %ERRORLEVEL% neq 0 (
    echo ❌ Go is not installed. Please install Go 1.21 or higher first.
    echo    Visit: https://golang.org/doc/install
    pause
    exit /b 1
)

echo ✅ Go detected

REM Navigate to TUI directory
cd /d "%~dp0tui"

REM Install dependencies
echo 📦 Installing dependencies...
go mod tidy

REM Build the binary
echo 🔨 Building sassysh-tui.exe...
if not exist build mkdir build
go build -o build\sassysh-tui.exe .

if %ERRORLEVEL% neq 0 (
    echo ❌ Build failed
    pause
    exit /b 1
)

echo ✅ Build complete!
echo.
echo 📋 Next steps:
echo    1. Add the build directory to your PATH, or
echo    2. Copy sassysh-tui.exe to a directory in your PATH
echo.
echo    3. Run the setup wizard:
echo       sassysh-tui.exe setup
echo.
echo    4. Use the regular sassysh command for queries:
echo       sassysh ask "your question"
echo.
echo 🎉 Happy coding with SassyShell TUI!
pause
