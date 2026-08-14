-- Lua tem três laços. O for numérico é o mais usado; while e repeat existem
-- para quando a condição de parada não é uma contagem.
local M = {}

function M.exemplos()
  -- for numérico: variável = início, fim, passo. Os dois extremos ENTRAM na
  -- contagem, ao contrário do que acontece nas linguagens de índice 0.
  for i = 1, 3 do
    io.write(i, " ") -- 1 2 3
  end
  print()

  -- O passo é opcional e vale 1 quando omitido. Negativo conta ao contrário.
  for i = 0, 10, 5 do io.write(i, " ") end -- 0 5 10
  print()
  for i = 3, 1, -1 do io.write(i, " ") end -- 3 2 1
  print()

  -- A variável do for é local ao laço e some depois dele. E o limite é
  -- calculado uma vez só, na entrada: mexer nele lá dentro não muda nada.
  local n = 3
  for i = 1, n do n = 100 end
  print(n) -- 100, e ainda assim o laço rodou 3 vezes

  -- while testa antes de cada volta; pode nunca rodar.
  local saldo = 10
  while saldo > 0 do
    saldo = saldo - 4
  end
  print(saldo) -- -2

  -- repeat testa DEPOIS, então roda pelo menos uma vez. Note que a condição é
  -- de parada ("until"), não de continuação.
  local tentativas = 0
  repeat
    tentativas = tentativas + 1
  until tentativas >= 3
  print(tentativas) -- 3

  -- break sai do laço mais interno na hora.
  local primeiroMultiplo
  for i = 10, 30 do
    if i % 7 == 0 then
      primeiroMultiplo = i
      break
    end
  end
  print(primeiroMultiplo) -- 14
end

-- SUA VEZ
--
-- Devolva a soma dos números pares de 1 até n: somaPares(10) é 2+4+6+8+10.
function M.somaPares(n)
  return -1 -- <- troque isto
end

-- Para ver a saída dos exemplos, abra o shell do container com [s] e rode:
--   lua -e 'require("exercicio").exemplos()'
return M
