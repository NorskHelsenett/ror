#!/bin/bash

function get_byte()
{
    local BYTE=$(head -c 1 /dev/random)

    if [ -z "$BYTE" ]; then
        BYTE=$(get_byte)
    fi

    echo "$BYTE"
}

function encode_password()
{
    BYTE1=$(get_byte)
    BYTE2=$(get_byte)
    BYTE3=$(get_byte)
    BYTE4=$(get_byte)

    SALT="${BYTE1}${BYTE2}${BYTE3}${BYTE4}"
    PASS="$SALT$1"
    TEMP=$(echo -n "$PASS" | openssl sha256 -binary)
    PASS="$SALT$TEMP"
    PASS=$(echo -n "$PASS" | base64)
    echo "$PASS"
}

encode_password $1