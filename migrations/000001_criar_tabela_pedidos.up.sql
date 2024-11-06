CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS pedidos (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    id_cliente VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    valor_total DECIMAL(10,2) NOT NULL,
    criado_em TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    atualizado_em TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_pedidos_id_cliente ON pedidos(id_cliente);
CREATE INDEX idx_pedidos_status ON pedidos(status);
CREATE INDEX idx_pedidos_criado_em ON pedidos(criado_em);

ALTER TABLE pedidos ADD CONSTRAINT check_valor_total_positivo 
    CHECK (valor_total >= 0);

ALTER TABLE pedidos ADD CONSTRAINT check_status_valido 
    CHECK (status IN ('pendente', 'aprovado', 'cancelado', 'entregue'));
