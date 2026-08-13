-- Estado persistido de cada exercício.
--
-- A chave primária é o caminho relativo do exercício (ex: go/sintaxe/slices-append).
-- Renomear um diretório cria um registro novo e deixa o antigo órfão — por isso
-- a coluna `orphaned` existe e por isso a reconciliação nunca deleta em silêncio.
CREATE TABLE progress (
    path               TEXT PRIMARY KEY,
    language           TEXT NOT NULL,
    category           TEXT NOT NULL,

    status             TEXT NOT NULL DEFAULT 'not_started'
                       CHECK (status IN ('not_started', 'in_progress', 'completed')),
    attempts           INTEGER NOT NULL DEFAULT 0 CHECK (attempts >= 0),

    -- Datas em texto RFC3339 (UTC). Explícito de propósito: o valor no banco é
    -- legível com qualquer cliente sqlite, sem depender de conversão do driver.
    first_completed_at TEXT,
    last_validated_at  TEXT,

    -- Hash dos arquivos editáveis na última validação. É o que permite
    -- descartar um save que não mudou nada.
    last_content_hash  TEXT NOT NULL DEFAULT '',

    -- Resultado da última validação. Junto com status, dispensa uma coluna
    -- separada para "regrediu": regressão é status completo + última falhou.
    last_passed        INTEGER NOT NULL DEFAULT 0 CHECK (last_passed IN (0, 1)),

    orphaned           INTEGER NOT NULL DEFAULT 0 CHECK (orphaned IN (0, 1))
);

CREATE INDEX idx_progress_language ON progress (language);
