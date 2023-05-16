#!/bin/bash

URL="https://github.com/tayron/estagiario/raw/main/dist/linux-64/estagiario"

EXECUTAVEL="estagiario"

echo "Baixando o executável..."
curl -L -o "$EXECUTAVEL" "$URL"

# Verifica se o download foi concluído com sucesso
if [ -e "$EXECUTAVEL" ]; then
  echo "Download concluído com sucesso!"
  chmod +x "$EXECUTAVEL"  # Dá permissão de execução ao arquivo
else
  echo "Erro ao baixar o executável."
fi
