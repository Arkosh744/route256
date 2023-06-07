package service

import (
	"context"
	"sync"
	"time"

	"route256/libs/client/pg"
	"route256/loms/internal/models"
)

type service struct {
	repo      Repository
	storage   orderStorage
	txManager pg.TxManager
}

func New(repo Repository, tx pg.TxManager) *service {
	return &service{
		repo:      repo,
		txManager: tx,
		storage: orderStorage{
			storage: make(map[int64]*orderStatus),
		},
	}
}

type Repository interface {
	CreateOrder(ctx context.Context, user int64) (int64, error)
	CreateOrderItems(ctx context.Context, orderID int64, items []models.Item) error
	UpdateOrderStatus(ctx context.Context, orderID int64, status string) error
	CreateReservation(ctx context.Context, orderID, warID int64, sku uint32, count uint64) error
	GetReservations(ctx context.Context, orderID int64) ([]models.ReservationItem, error)
	DeleteReservation(ctx context.Context, orderID int64) error
	InsertStock(ctx context.Context, item models.ReservationItem) error
	UpdateStock(ctx context.Context, warehouseID int64, sku uint32, count uint64) error
	DeleteStock(ctx context.Context, warehouseID int64, sku uint32) error
	GetStocks(ctx context.Context, sku uint32) ([]models.StockItem, error)
}

// orderStorage is a storage for orders to cancel them after timeout.
type orderStorage struct {
	storage map[int64]*orderStatus
	mu      sync.Mutex
}

type orderStatus struct {
	Paid      bool
	Timer     *time.Timer
	ctxCancel context.CancelFunc
}

const orderTimeout = 10 * time.Minute
