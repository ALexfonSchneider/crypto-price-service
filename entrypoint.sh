#!/bin/sh
set -e

echo "Running migrations..."
./migrate

echo "Starting service..."
./crypto-price-service