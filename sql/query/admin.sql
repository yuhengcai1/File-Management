-- name: CreateUsers :one
INSERT INTO users (
  id,username,userhash,admin
) VALUES (
  $1,$2,$3,$4
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
-- name: GetUserByIDAndAdmin :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByNAME :one
SELECT * FROM users WHERE username = $1;

-- name: GetDocumentByCreatebyNormal :many
SELECT * FROM document WHERE documentid = $1 AND createdby = $2;

-- name: GetDocumentByID :one
SELECT * FROM document WHERE documentid = $1;

/* All Update operations */
-- name: UpdateUsers :exec
UPDATE users SET userhash = $1 WHERE id = $2;

/*All Delete operations*/

-- name: Deleteusers :exec
DELETE FROM users WHERE id = $1;

-- name: DeleteDocumentAdmin :exec
DELETE FROM document WHERE documentid = $1;

-- name: DeleteDocumentNormal :exec
DELETE FROM document WHERE documentid = $1 AND createdby = $2;