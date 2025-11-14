#!/usr/bin/env bash

REQ_FILE="mkdocs-requirements.txt"

# Colors
GREEN="\033[0;32m"
RED="\033[0;31m"
YELLOW="\033[1;33m"
NC="\033[0m" # No Color


pip_cmd="pip3"

if [[ -z "$pip_cmd" ]]; then
    echo -e "${RED}Error: pip or pip3 not found.${NC}"
    exit 1
fi

echo -e "${GREEN}Using $pip_cmd${NC}"


# -----------------------------
# Check if Python package is installed
# -----------------------------
is_installed() {
    local cmd="$pip_cmd show $1"
    $cmd >/dev/null 2>&1
    return $?
}



# -----------------------------
# Prompt Y/N
# -----------------------------
prompt_yn() {
    while true; do
        read -p "$1 (Y/N): " yn
        case $yn in
            [Yy]* ) return 0 ;;
            [Nn]* ) return 1 ;;
            * ) echo "Please answer Y or N." ;;
        esac
    done
}


# -----------------------------
# Install missing requirements
# -----------------------------
if [[ ! -f "$REQ_FILE" ]]; then
    echo -e "${YELLOW}Warning: $REQ_FILE not found, skipping requirements check.${NC}"
else
    missing=()
    while IFS= read -r pkg; do
        [[ -z "$pkg" || "$pkg" =~ ^# ]] && continue
        if ! is_installed "$pkg"; then
            missing+=("$pkg")
        fi
    done < "$REQ_FILE"

    if (( ${#missing[@]} > 0 )); then
        echo -e "${YELLOW}Missing required packages:${NC}"
        for m in "${missing[@]}"; do
            echo "  - $m"
        done

        if prompt_yn "Install missing packages?"; then
            echo "Installing: ${missing[*]}"
            $pip_cmd install "${missing[@]}"
        else
            echo -e "${RED}Cannot continue without required packages.${NC}"
            exit 1
        fi
    else
        echo -e "${GREEN}All requirement packages are installed.${NC}"
    fi
fi

# -----------------------------
# Start MKDocs
# -----------------------------
echo -e "${GREEN}Starting MKDocs...${NC}"
if command -v mkdocs >/dev/null 2>&1; then
    mkdocs serve
else
    # fallback to python module
    python3 -m mkdocs serve
fi
