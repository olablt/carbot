#!/bin/bash

# Kill any existing Chromium processes
pkill chromium-browser
pkill chromium-browser
sleep 1 # Wait a bit for the browser to start and print the URL
# Start Chromium in the background and redirect the output to a temporary file
chromium-browser --remote-debugging-port=9222 > /home/bob/Projects/autoplius/.tmp/chromium_debug.txt 2>&1 &
sleep 2 # Wait a bit for the browser to start and print the URL

# Extract the URL
DEBUG_URL=$(grep -m 1 'DevTools listening on' /home/bob/Projects/autoplius/chromium_debug_url.txt | awk '{print $NF}')
echo $DEBUG_URL > /home/bob/Projects/autoplius/chromium_debug_url.txt

# Use the URL (e.g., open it in a default browser, print it, etc.)
echo "Debug URL: $DEBUG_URL"
