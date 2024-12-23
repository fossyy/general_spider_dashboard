-- TRIGGER: INSERT FIRST VALUE AND UPDATE AUTO INCREMENT FOR TABLE config_histories
CREATE OR REPLACE FUNCTION handle_config_version_trigger()
RETURNS TRIGGER AS $$
DECLARE
dashboard_ver TEXT;
BEGIN
SELECT dashboard_version INTO dashboard_ver
FROM dashboard_configs
WHERE config_name = NEW.name
    LIMIT 1;

-- INIT FOR DASHBOARD VERSION IF VALUE NULL
IF dashboard_ver IS NULL THEN
        INSERT INTO dashboard_configs (config_name, dashboard_version)
        VALUES (NEW.name, 1);

        dashboard_ver := 1;
END IF;

    -- INIT COMMIT FOR FIRST VERSION
    IF TG_OP = 'INSERT' THEN
        INSERT INTO config_histories (id, configs_id, base_url, version, data)
        VALUES (
            gen_random_uuid(),
            NEW.id,
            NEW.domain,
            COALESCE(dashboard_ver, '1') || '.' || COALESCE(NEW.config_version::TEXT, '1'),
            NEW.data
        );

RETURN NEW;

-- COMMIT FOR CONFIG VERSION
ELSIF TG_OP = 'UPDATE' THEN
        INSERT INTO config_histories (id, configs_id, base_url, version, data)
        VALUES (
            gen_random_uuid(),
            NEW.id,
            NEW.domain,
            CONCAT(dashboard_ver, '.', NEW.config_version + 1),
            NEW.data
        );

        NEW.config_version := NEW.config_version + 1;

RETURN NEW;
END IF;

RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- INIT COMMIT FOR FIRST VERSION
CREATE TRIGGER trigger_log_or_increment_config_version_insert
    AFTER INSERT ON configs
    FOR EACH ROW
    EXECUTE FUNCTION handle_config_version_trigger();

-- COMMIT FOR CONFIG VERSION
CREATE TRIGGER trigger_log_or_increment_config_version_update
    BEFORE UPDATE ON configs
    FOR EACH ROW
    EXECUTE FUNCTION handle_config_version_trigger();

-- FUNCTION INSERT TO TABLE TEMP VERSION
CREATE OR REPLACE FUNCTION insert_temp_version(name_config TEXT, dashboard_ver TEXT, configs_ver TEXT)
RETURNS void AS $$
BEGIN
INSERT INTO temp_versions (id, base_url, dashboard_id, configs_id, version)
VALUES (
           gen_random_uuid(),
           (SELECT c.domain FROM configs c WHERE c.name = name_config LIMIT 1),
       (SELECT dc.id FROM dashboard_configs dc WHERE dc.config_name = name_config LIMIT 1),
       (SELECT c.id FROM configs c WHERE c.name = name_config LIMIT 1),
    CONCAT(dashboard_ver, '.', configs_ver)
    )
ON CONFLICT (dashboard_id)
    DO UPDATE
           SET base_url = EXCLUDED.base_url,
           configs_id = EXCLUDED.configs_id,
           version = EXCLUDED.version;
END
$$ LANGUAGE plpgsql;

-- FUNCTION VIEW FOR VIEW RESULT VERSION
CREATE OR REPLACE VIEW combined_version AS
SELECT
    dc.id AS dashboard_id,
    dc.config_name,
    dc.dashboard_version,
    c.config_version,
    CONCAT(dc.dashboard_version, '.', c.config_version) AS full_version
FROM dashboard_configs dc
         LEFT JOIN configs c ON dc.config_name = c.name;

-- FUNCTION TRIGGER FOR RESET CONFIG VERSION IF DASHBOARD VERSION UPDATE
CREATE OR REPLACE FUNCTION reset_config_version_on_dashboard()
RETURNS TRIGGER AS $$
BEGIN
UPDATE configs
SET config_version = 0
WHERE name = NEW.config_name;

RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- CREATE TRIGGER trigger_reset_config_version_on_dashboard
CREATE TRIGGER trigger_reset_config_version_on_dashboard
    AFTER UPDATE OF dashboard_version ON dashboard_configs
    FOR EACH ROW
    EXECUTE FUNCTION reset_config_version_on_dashboard();