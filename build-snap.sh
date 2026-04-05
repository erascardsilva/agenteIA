#!/bin/bash

# Script de build do pacote Snap do AgenteIA
# O build real é feito via GitHub Actions (ubuntu-22.04)
# Veja: .github/workflows/snap.yml

echo "O build do Snap é gerado automaticamente via GitHub Actions."
echo ""
echo "Para acionar o build:"
echo "  git push origin main"
echo ""
echo "Após o workflow concluir, baixe o artefato 'agenteia-snap' em:"
echo "  GitHub -> Actions -> Build Snap -> Artifacts"
echo ""
echo "Para instalar localmente após download:"
echo "  sudo snap install --dangerous agenteia_*.snap"
echo ""
echo "Erasmo Cardoso - Software Engineer | Electronics Specialist"
