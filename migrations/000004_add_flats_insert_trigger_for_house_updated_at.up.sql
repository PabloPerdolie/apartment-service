CREATE OR REPLACE FUNCTION update_house_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE houses SET updated_at = NOW() WHERE id = NEW.house_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER flats_after_insert
AFTER INSERT ON flats
FOR EACH ROW
EXECUTE FUNCTION update_house_timestamp();