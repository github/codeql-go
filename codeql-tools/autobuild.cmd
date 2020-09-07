@echo off
SETLOCAL EnableDelayedExpansion

rem Some legacy environment variables for the autobuilder.
set LGTM_SRC=%CD%

if %CODEQL_GO_TRACE% = "" (
  type NUL && "%CODEQL_EXTRACTOR_GO_ROOT%/tools/%CODEQL_PLATFORM%/go-build"
) else (
  type NUL && "%CODEQL_EXTRACTOR_GO_ROOT%/tools/%CODEQL_PLATFORM%/go-autobuilder.exe"
)
exit /b %ERRORLEVEL%

ENDLOCAL
