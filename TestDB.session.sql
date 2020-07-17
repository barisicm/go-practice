BEGIN;

DROP TABLE IF EXISTS miyagi.traffic_quotas CASCADE;
DROP TYPE IF EXISTS tdate;
DROP TYPE IF EXISTS towner;

CREATE TYPE tdate AS ENUM ('day', 'month', 'week');
CREATE TYPE towner AS ENUM ('device', 'subscriber_user');

CREATE TABLE IF NOT EXISTS miyagi.traffic_quotas (
    id BIGSERIAL PRIMARY KEY,
    owner_uuid UUID,
    owner_type towner,
    period tdate,
    byte_quota BIGINT
);

COMMIT;