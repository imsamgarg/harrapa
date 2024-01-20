-- +goose Up
CREATE TABLE inscriptions(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    site TEXT,
    cisi_id VARCHAR(20),
    sequence_images TEXT [] CHECK (array_length(sequence_images, 1) > 0),
    sequence_numbers INTEGER [] CHECK (array_length(sequence_numbers, 1) > 0),
    wells_id VARCHAR(20),
    artifact_type VARCHAR(20),
    material_type VARCHAR(50),
    field_symbol VARCHAR(30),
    excavation_number INTEGER,
    area VARCHAR(10)
);

-- +goose Down
DROP Table inscriptions;