-- name: GetAdminDivision :one
SELECT * FROM admin_division
WHERE division_id = ?;

-- name: ListAdminDivision :many
SELECT * FROM admin_division;

-- name: CreateAdminDivision :execresult
INSERT INTO admin_division (
  division_name
) VALUES (
  ?
);

-- name: UpdateAdminDivision :execresult
UPDATE admin_division 
SET division_name = ?
WHERE division_id = ?;