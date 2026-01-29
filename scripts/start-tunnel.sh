#!/bin/bash

# Configuration
URL="http://localhost:3000"
BIN_NAME="cloudflared"

# Function to check dependencies
check_cloudflared() {
    if ! command -v ./$BIN_NAME &> /dev/null; then
        echo "⬇️  cloudflared not found. Downloading..."
        # Detect OS architecture
        ARCH=$(uname -m)
        if [ "$ARCH" = "x86_64" ]; then
            wget -q https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-linux-amd64 -O $BIN_NAME
        elif [ "$ARCH" = "aarch64" ]; then
            wget -q https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-linux-arm64 -O $BIN_NAME
        else
             echo "❌ Unsupported architecture: $ARCH"
             exit 1
        fi
        
        chmod +x $BIN_NAME
        echo "✅ cloudflared downloaded."
    fi
}

# Main execution
echo "🚀 Starting Cloudflare Tunnel..."
check_cloudflared

echo "🌍 Exposing $URL to the internet..."
echo "⚠️  Copy the URL below ending in .trycloudflare.com to share!"
echo ""

# Run tunnel
./$BIN_NAME tunnel --url $URL
