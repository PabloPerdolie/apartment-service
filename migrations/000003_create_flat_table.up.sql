CREATE TABLE IF NOT EXISTS flats (
    id SERIAL PRIMARY KEY,
    house_id INT NOT NULL REFERENCES houses(id),
    price INT NOT NULL,
    rooms INT NOT NULL,
    status VARCHAR(13),
    moderator_id VARCHAR(36) REFERENCES users(id)
);

CREATE INDEX idx_flats_house_id ON flats(house_id);