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

-- Esta é a asserção que separa a versão ingênua da correta: a chamada com
-- ponto não passa a conta como self, então o depósito nunca acontece do
-- jeito certo — e o programa quebra antes de chegar aqui.
verificar("deposita 50 numa conta de 100", 150, ex.depositar(ex.Conta.nova(100), 50))
verificar("deposita 0 não muda o saldo", 100, ex.depositar(ex.Conta.nova(100), 0))

local conta = ex.Conta.nova(20)
ex.depositar(conta, 10)
verificar("dois depósitos seguidos acumulam", 40, ex.depositar(conta, 10))

if falhas > 0 then
  print(("\n%d verificação(ões) falharam"):format(falhas))
  os.exit(1)
end
print("\ntodas as verificações passaram")
