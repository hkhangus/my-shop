package usecase

import (
	"context"
	"github.com/scul0405/my-shop/server/config"
	dbconverter "github.com/scul0405/my-shop/server/db/converter"
	dbmodels "github.com/scul0405/my-shop/server/db/models"
	"github.com/scul0405/my-shop/server/internal/book"
	"github.com/scul0405/my-shop/server/internal/dto"
	"github.com/scul0405/my-shop/server/internal/order"
	"github.com/scul0405/my-shop/server/pkg/logger"
	"github.com/scul0405/my-shop/server/pkg/utils"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"time"
)

type orderUseCase struct {
	cfg       *config.Config
	bookRepo  book.Repository
	orderRepo order.Repository
	logger    logger.Logger
}

func NewOrderUseCase(
	cfg *config.Config,
	bookRepo book.Repository,
	orderRepo order.Repository,
	logger logger.Logger) order.UseCase {
	return &orderUseCase{cfg: cfg, bookRepo: bookRepo, orderRepo: orderRepo, logger: logger}
}

func (u *orderUseCase) Create(ctx context.Context, order *dto.CreateOrderDTO) (*dto.OrderDTO, error) {
	createdOrder, err := u.orderRepo.Create(ctx, order.ToModel())
	if err != nil {
		return nil, err
	}

	return dbconverter.OrderModelToDto(createdOrder), nil
}

func (u *orderUseCase) AddBook(ctx context.Context, oid, bid uint64, data *dto.CreateBookOrderDTO) error {
	// check if book is exist
	bookModel, err := u.bookRepo.GetByID(ctx, bid)
	if err != nil {
		return err
	}

	// check if order is exist
	orderModel, err := u.GetByID(ctx, oid)
	if err != nil {
		return err
	}

	bookOrder := &dbmodels.BookOrder{
		BookID:   int64(bid),
		OrderID:  int64(oid),
		Quantity: data.Quantity,
	}

	// Add book to order
	err = u.orderRepo.AddBook(ctx, bookOrder)
	if err != nil {
		return err
	}

	// Update total of order
	orderModel.Total += bookModel.Price
	err = u.Update(ctx, orderModel)
	if err != nil {
		return err
	}

	return nil
}

func (u *orderUseCase) GetByID(ctx context.Context, id uint64) (*dto.OrderDTO, error) {
	orderModel, err := u.orderRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	bookModelSlice, err := u.bookRepo.GetByOrderID(ctx, id)
	if err != nil {
		return nil, err
	}

	orderDTO := dbconverter.OrderModelToDto(orderModel)

	for _, bookModel := range bookModelSlice {
		orderDTO.Books = append(orderDTO.Books, dbconverter.BookModelToDto(bookModel))
	}

	return orderDTO, nil
}

func (u *orderUseCase) Update(ctx context.Context, order *dto.OrderDTO) error {
	whiteList := []string{"total", "status"}
	err := u.orderRepo.Update(ctx, order.ToModel(), whiteList...)
	if err != nil {
		return err
	}

	return nil
}

func (u *orderUseCase) Delete(ctx context.Context, id uint64) error {
	err := u.orderRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *orderUseCase) List(ctx context.Context, pq *utils.PaginationQuery) (*utils.PaginationList, error) {
	var (
		paginationList *utils.PaginationList
		err            error
		ordersDTO      []*dto.OrderDTO
		qms            = make([]qm.QueryMod, 0)
	)

	from, _ := time.Parse(time.DateOnly, ctx.Value("from").(string))
	to, _ := time.Parse(time.DateOnly, ctx.Value("to").(string))

	// check if from and to is not default parse value
	if from != (time.Time{}) {
		qms = append(qms, dbmodels.OrderWhere.CreatedAt.GTE(from))
	}
	if to != (time.Time{}) {
		qms = append(qms, dbmodels.OrderWhere.CreatedAt.LTE(to))
	}

	paginationList, err = u.orderRepo.List(ctx, pq, qms...)

	if err != nil {
		return nil, err
	}

	for _, orderModel := range paginationList.List.(dbmodels.OrderSlice) {
		ordersDTO = append(ordersDTO, dbconverter.OrderModelToDto(orderModel))
	}

	paginationList.List = ordersDTO

	return paginationList, nil
}
