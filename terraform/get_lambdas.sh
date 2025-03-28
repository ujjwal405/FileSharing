#!/bin/bash

echo "[]" > "${PWD}/lambda_dirs.json"
dirs=$(ls -d ../backend/lambda/*/ 2>/dev/null || true)
if [ -n "$dirs" ]; then
  echo "[" > "${PWD}/lambda_dirs.json"
  first=true
  for dir in $dirs; do
    if [ "$first" = true ]; then
      first=false
    else
      echo "," >> "${PWD}/lambda_dirs.json"
    fi
    echo "\"$(basename "$dir")\"" >> "${PWD}/lambda_dirs.json"
  done
  echo "]" >> "${PWD}/lambda_dirs.json"
fi
