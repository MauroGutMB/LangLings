local ex = require("exercicio")

local falhas = 0

local function verificar(o_que, esperado, obtido)
  if esperado == obtido then
    print("ok    " .. o_que)
    return
  end
  print(("FALHA %s\n      esperado: %s\n      obtido:   %s")
    :format(o_que, tostring(esperado), tostring(obtido)))
  falhas = falhas + 1
end

verificar("cortarAntesDe('nome:valor', ':')", "nome", ex.cortarAntesDe("nome:valor", ":"))
verificar("cortarAntesDe('a/b/c', '/')", "a", ex.cortarAntesDe("a/b/c", "/"))
-- Esta é a asserção que separa a versão ingênua da correta: com separador de
-- mais de um caractere, sub(1, pos) inclui só o primeiro dele, escondendo o
-- erro se você olhar só o primeiro caso de teste acima.
verificar("cortarAntesDe('chave==valor', '==')", "chave", ex.cortarAntesDe("chave==valor", "=="))
verificar("cortarAntesDe sem separador", "semseparador", ex.cortarAntesDe("semseparador", ":"))
verificar("cortarAntesDe separador na primeira posição", "", ex.cortarAntesDe(":resto", ":"))

if falhas > 0 then
  print(("\n%d verificação(ões) falharam"):format(falhas))
  os.exit(1)
end
print("\ntodas as verificações passaram")
