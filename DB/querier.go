// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package DB

import (
	"context"
)

type Querier interface {
	CreateDocument(ctx context.Context, arg CreateDocumentParams) (Document, error)
	CreateUsers(ctx context.Context, arg CreateUsersParams) (User, error)
	Createadmin(ctx context.Context, id int32) (int32, error)
	Createnormal(ctx context.Context, arg CreatenormalParams) (Normal, error)
	//All Delete operations
	DeleteAdmin(ctx context.Context, id int32) error
	DeleteDocumentAdmin(ctx context.Context, documentid int32) error
	DeleteDocumentNormal(ctx context.Context, arg DeleteDocumentNormalParams) error
	DeleteNormal(ctx context.Context, id int32) error
	Deleteusers(ctx context.Context, id int32) error
	// All Get operations
	GetAdminByID(ctx context.Context, id int32) (int32, error)
	GetDocumentByCreatebyNormal(ctx context.Context, arg GetDocumentByCreatebyNormalParams) ([]Document, error)
	GetDocumentByID(ctx context.Context, documentid int32) ([]Document, error)
	GetNormalByCreateby(ctx context.Context, id int32) ([]Normal, error)
	GetNormalByID(ctx context.Context, id int32) (Normal, error)
	GetUserByID(ctx context.Context, id int32) (User, error)
	UpdateNormal(ctx context.Context, id int32) error
	// All Update operations
	UpdateUsers(ctx context.Context, id int32) error
}

var _ Querier = (*Queries)(nil)
