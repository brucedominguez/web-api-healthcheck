CREATE TABLE wallet.wallet (
  wallet_id VARCHAR(36),
  tx_id VARCHAR(255),
  credit BIGINT,
  debit BIGINT,
  balance BIGINT,
  version BIGINT,
  seq_num    BIGSERIAL    NOT NULL UNIQUE,
  timestamp  TIMESTAMPTZ  NOT NULL DEFAULT now(),
  PRIMARY KEY (tx_id, version)
);

-- Example transaction
INSERT INTO wallet.wallet (wallet_id, tx_id, credit, debit, balance, version) VALUES
('007','tx101', 100, 0, 100, 1);