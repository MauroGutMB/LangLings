local M = {}

-- dividir devolve o quociente e o resto da divisão inteira de a por b.
local function dividir(a, b)
  return a // b, a % b
end

-- empacotar devolve {quociente, resto, "extra"}, com quociente e resto
-- vindos de dividir(a, b).
--
-- TODO: falta um elemento na lista devolvida. Por quê?
function M.empacotar(a, b)
  return {dividir(a, b), "extra"}
end

return M
