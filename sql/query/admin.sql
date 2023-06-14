-- name: Createadmin :one
INSERT INTO admin (
  id
) VALUES (
  $1
)
RETURNING *;

-- name: Createnormal :one
INSERT INTO normal (
  id,createdby
) VALUES (
  $1,$2
)
RETURNING *;

-- name: CreateUsers :one
INSERT INTO users (
  id,username,password
) VALUES (
  $1,$2,$3
)
RETURNING *;


-- name: CreateDocument :one
INSERT INTO document (
  documentid,name,createdby
) VALUES (
  $1,$2,$3
)
RETURNING *;

/* All Get operations */
-- name: GetAdminByID :one
SELECT * FROM admin
WHERE id = $1;

-- name: GetNormalByID :one
SELECT * FROM normal
WHERE id = $1;

-- name: GetNormalByCreateby :many
SELECT * FROM normal
WHERE id = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetDocumentByCreatebyNormal :many
SELECT * FROM document WHERE documentid = $1 AND createdby = $2;

-- name: GetDocumentByID :many
SELECT * FROM document WHERE documentid = $1;

/* All Update operations */
-- name: UpdateUsers :exec
UPDATE users SET id = $1;

-- name: UpdateNormal :exec
UPDATE normal SET id = $1;

/*All Delete operations*/

-- name: DeleteAdmin :exec
DELETE FROM admin WHERE id = $1;

-- name: DeleteNormal :exec
DELETE FROM normal WHERE id = $1;

-- name: Deleteusers :exec
DELETE FROM users WHERE id = $1;

-- name: DeleteDocumentAdmin :exec
DELETE FROM document WHERE documentid = $1;

-- name: DeleteDocumentNormal :exec
DELETE FROM document WHERE documentid = $1 AND createdby = $2;