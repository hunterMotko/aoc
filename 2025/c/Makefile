# --- Configuration ---
CC = gcc
CFLAGS = -Wall -Wextra -std=c11
TARGET_DIR = bin
HEADER_DIR = include
LIB=lib

# --- File Paths ---
SOURCE = days/$(DAY)/part$(PART).c
TARGET = $(TARGET_DIR)/days/$(DAY)/part$(PART)
LIB_FILE = $(LIB)/parse.c

# --- PHONY Targets ---
.PHONY: build run clean

# 1. BUILD: make build DAY=1 PART=1
build: $(TARGET_DIR) $(TARGET)
	@echo "Build complete: $(TARGET)"

# 2. RUN: make run DAY=1 PART=1
run: $(TARGET)
	@echo "--- Running $(TARGET) ---"
	@./$(TARGET)

# --- Compilation Rule ---
$(TARGET): $(SOURCE) $(LIB_FILE)
	@echo "--- Compiling $< and $(LIB_FILE) to $@ ---"
	# -I$(HEADER_DIR) includes the 'include' directory for custom helper .h files
	@$(CC) $(CFLAGS) -I$(HEADER_DIR) $< $(LIB_FILE) -o $@

# Create the binary output directory if it doesn't exist
$(TARGET_DIR):
	@mkdir -p $(TARGET_DIR)

# 3. CLEAN: make clean
clean:
	@echo "--- Cleaning up compiled binaries in $(TARGET_DIR) ---"
	@rm -rf $(TARGET_DIR)
