#!/bin/bash
set -e

echo "[FlipApp] Instalando dependências..."

if ! command -v pnpm &> /dev/null; then
  echo "pnpm não encontrado, instalando..."
  npm install -g pnpm
fi

pnpm install