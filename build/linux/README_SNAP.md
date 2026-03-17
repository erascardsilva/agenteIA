# Guia de Publicação Snap - AgenteIA

Este documento descreve como compilar e publicar o AgenteIA na loja Snap Store.

## Estrutura de Arquivos

Os arquivos de configuração estão localizados em `build/linux/`:
- `snapcraft.yaml`: Receita principal de compilação (instruções para o `snapcraft`).

## Build Local (Teste)

Para testar o build via código fonte na sua máquina local:

1. **Instale as ferramentas necessárias:**
   ```bash
   sudo snap install snapcraft --classic
   sudo snap install lxd
   sudo lxd init --auto
   ```

2. **Execute o script de build:**
   Na raiz do projeto, execute o script automatizado:
   ```bash
   ./build-snap.sh
   ```

   *Alternativamente*, se preferir compilar manualmente (via LXD ou Multipass):
   ```bash
   cd build/linux
   snapcraft
   ```

## Como publicar na Snap Store

1.  **Registre o nome do aplicativo** (caso ainda não tenha feito):
    ```bash
    snapcraft register agenteia
    ```

2.  **Envie o arquivo .snap gerado** para a loja:
    Após a compilação, o pacote `.snap` estará na pasta `build/bin/`. Envie para a branch edge:
    ```bash
    snapcraft upload --release=edge build/bin/agenteia_1.0.0_amd64.snap
    ```

3.  Após testes e validações locais, você pode promover seu build para a versão estável através do painel [Snapcraft Developer Dashboard](https://snapcraft.io/snaps).

---
**Erasmo Cardoso - Dev**
