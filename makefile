# Detect the operating system
ifeq ($(OS),Windows_NT)
    detected_OS := Windows
else
    detected_OS := $(shell uname -s)
endif

# Source and executable definitions
SRC=socialsent.go
EXEC=socialsent
WIN_EXEC=socialsent.exe

all:
	go build $(SRC)

clean:
# Windows-specific clean command
ifeq ($(detected_OS),Windows)
	del /f /q $(EXEC) $(WIN_EXEC)
else
# Linux/Unix/MacOS-specific clean command
	rm -f $(EXEC) $(WIN_EXEC)
endif
