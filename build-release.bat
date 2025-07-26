@echo off
echo Building Aurora - Release Version (Embedded Assets, No Console)...
go build -ldflags "-H windowsgui" -o aurora-release.exe .
echo Done! Created: aurora-release.exe
pause
