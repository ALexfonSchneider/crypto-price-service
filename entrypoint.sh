#!/bin/sh
set -e

echo "🔧 Running migrations..."
./migrate

echo "🚀 Starting service..."
exec ./crypto-price-service