#!/bin/bash

# Script para gerar o Flatpak do AgenteIA
# Requisitos: wails, flatpak, flatpak-builder

set -e

echo "--- Iniciando processo de geração do Flatpak ---"

# 1. Build do Wails
echo "Compilando aplicação com Wails..."
wails build

# 2. Configuração de diretórios
BUILD_DIR="build/linux"
REPO_DIR="build/flatpak-repo"
EXPORT_DIR="build/flatpak-export"

mkdir -p $REPO_DIR
mkdir -p $EXPORT_DIR
mkdir -p build/bin

# 3. Build do Flatpak
echo "Gerando Flatpak bundle..."
cd $BUILD_DIR

flatpak-builder --force-clean --repo=../flatpak-repo build-dir org.erascardsilva.agenteIA.yaml

# 4. Criar o arquivo .flatpak (bundle) para distribuição
flatpak build-bundle ../flatpak-repo ../bin/agenteIA.flatpak org.erascardsilva.agenteIA

cd ../..

echo "--- Processo concluído! ---"
echo "O instalador foi gerado em: build/bin/agenteIA.flatpak"
echo ""
echo "Para instalar localmente:"
echo "flatpak install --user build/bin/agenteIA.flatpak"
echo ""
echo "Erasmo Cardoso - Dev"
