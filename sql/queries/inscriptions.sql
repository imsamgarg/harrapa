-- Active: 1705748378099@@127.0.0.1@5432@harrapa@public

-- name: AddInscription :one
INSERT INTO
    inscriptions (
        area, artifact_type, cisi_id, excavation_number, field_symbol, material_type, sequence_images, sequence_numbers, site, wells_id
    )
VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
    )
RETURNING
    id;

-- name: GetInsciptionById :one
SELECT * FROM inscriptions WHERE id = $1;

-- name: GetAllInsciptions :many
SELECT * FROM inscriptions;

-- name: DeleteInsciption :exec
DELETE FROM inscriptions WHERE id = $1;