package book

import (
	"context"
	customdbmodels "github.com/scul0405/my-shop/server/db/customModels"
	dbmodels "github.com/scul0405/my-shop/server/db/models"
	"github.com/scul0405/my-shop/server/pkg/utils"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Repository interface {
	Create(ctx context.Context, book *dbmodels.Book) (*dbmodels.Book, error)
	GetByID(ctx context.Context, id uint64) (*dbmodels.Book, error)
	Update(ctx context.Context, book *dbmodels.Book, whiteList ...string) error
	Delete(ctx context.Context, id uint64) error
	List(ctx context.Context, pq *utils.PaginationQuery, qms ...qm.QueryMod) (*utils.PaginationList, error)
	ListByCategoryName(ctx context.Context, pq *utils.PaginationQuery, categoryName string) (*utils.PaginationList, error)
	GetByOrderID(ctx context.Context, id uint64) (customdbmodels.BookInOrderSlice, error)
}
