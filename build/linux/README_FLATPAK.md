# Guia de Publicação Flatpak - AgenteIA

Este documento descreve como compilar e publicar o AgenteIA nas lojas Flatpak (Flathub).

## Estrutura de Arquivos

Os arquivos de configuração estão localizados em `build/linux/`:
- `org.erascardsilva.agenteIA.yaml`: Manifesto de build (instruções para o `flatpak-builder`).
- `org.erascardsilva.agenteIA.metainfo.xml`: Metadados do AppStream (descrição, screenshots, licença).
- `org.erascardsilva.agenteIA.desktop`: Atalho para o menu de aplicativos.

## Build Local (Teste)

Para testar o build via código fonte na sua máquina local:

1. **Instale as ferramentas necessárias:**
   ```bash
   sudo apt install flatpak-builder
   ```

2. **Instale o SDK e Runtime do GNOME 45:**
   ```bash
   flatpak install flathub org.gnome.Sdk//45 org.gnome.Platform//45
   ```

3. **Execute o build:**
   ```bash
   flatpak-builder --force-clean --user --install build-dir build/linux/org.erascardsilva.agenteIA.yaml
   ```

## Como publicar no Flathub (Processo em 2 Etapas)

A publicação oficial exige que você separe o que é código do que é a "receita" de build.

### Passo 1: Preparar seu GitHub (Repositório `agenteIA`)
Você deve garantir que os seguintes arquivos já estejam no seu repositório GitHub:
1.  `build/linux/org.erascardsilva.agenteIA.metainfo.xml`
2.  `build/linux/org.erascardsilva.agenteIA.desktop`
3.  `build/appicon.png`

**Por que?** Quando o Flathub for compilar seu app, ele vai baixar seu código do GitHub e procurar por esses arquivos lá dentro para criar o ícone e o menu.

### Passo 2: A Submissão (Pull Request no Flathub)
Agora você "avisa" ao Flathub que seu app existe seguindo estes comandos técnicos:

1.  **Faça o Fork** do repositório [flathub/flathub](https://github.com/flathub/flathub) no seu GitHub (Desmarque a opção "Copy the master branch only").
2.  **Clone o seu fork** localmente (substitua `seu-usuario` pelo seu nome no GitHub):
    ```bash
    git clone --branch=new-pr git@github.com:seu-usuario/flathub.git
    cd flathub
    ```
3.  **Crie uma branch** para o AgenteIA a partir da `new-pr`:
    ```bash
    git checkout -b submissao-agenteIA new-pr
    ```
4.  **Adicione apenas o manifesto** `org.erascardsilva.agenteIA.yaml` na raiz dessa pasta, faça o commit e o push:
    ```bash
    cp /home/erasmo/GITHUB/agenteIA/build/linux/org.erascardsilva.agenteIA.yaml .
    git add org.erascardsilva.agenteIA.yaml
    git commit -m "Add org.erascardsilva.agenteIA"
    git push origin submissao-agenteIA
    ```
5.  **Abra o Pull Request** no site do GitHub apontando para a branch `new-pr` do Flathub original.

Os servidores do Flathub lerão seu arquivo YAML, verão o link para seu GitHub, baixarão tudo e farão o build automaticamente. Se houver qualquer problema, eles avisarão no Pull Request.

---
**Erasmo Cardoso - Dev**
