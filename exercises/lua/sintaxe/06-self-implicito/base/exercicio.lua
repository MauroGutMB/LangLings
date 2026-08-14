local M = {}

local Conta = {}
Conta.__index = Conta

function Conta.nova(saldoInicial)
  return setmetatable({saldo = saldoInicial}, Conta)
end

-- depositar soma valor ao saldo desta conta. self é a própria conta, passada
-- implicitamente por quem chama com :
function Conta:depositar(valor)
  self.saldo = self.saldo + valor
end

-- M.depositar soma valor ao saldo de conta e devolve o saldo resultante.
--
-- TODO: chamar isto quebra na hora. Por quê?
function M.depositar(conta, valor)
  conta.depositar(valor)
  return conta.saldo
end

M.Conta = Conta
return M
