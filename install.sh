#!/bin/bash

URL="https://github.com/tayron/estagiario/raw/main/dist/linux-64/estagiario"

EXECUTAVEL="estagiario"

echo "Chamando o estagiário..."
curl -L -o "$EXECUTAVEL" "$URL"

# Verifica se o download foi concluído com sucesso
if [ -e "$EXECUTAVEL" ]; then
  echo "Estagiário pronto para te ajudar!"
  chmod +x "$EXECUTAVEL"  # Dá permissão de execução ao arquivo
else
  echo "Não foi possível chamar o estagiário!"
fi
