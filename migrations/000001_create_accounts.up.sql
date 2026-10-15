-- TODO: Document a repeatable migration command and verify it against a fresh database.
-- TODO: Decide whether Go or PostgreSQL generates account UUIDs.
-- TODO: Define supported currencies and account statuses, then enforce them with CHECK constraints.
-- TODO: If this migration has already been shared/applied, put schema changes in a new migration.
CREATE TABLE accounts (
  id UUID PRIMARY KEY,
  currency VARCHAR(3) NOT NULL,
  status VARCHAR(20) NOT NULL DEFAULT 'active',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
