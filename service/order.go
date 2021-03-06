package service

import (
	"github.com/Team-73/backend/domain"
	"github.com/Team-73/backend/domain/contract"
	"github.com/Team-73/backend/domain/entity"
	"github.com/Team-73/backend/utils/priceutils"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

type orderService struct {
	svc *Service
}

//newOrderService return a new instance of the service
func newOrderService(svc *Service) contract.OrderService {
	return &orderService{
		svc: svc,
	}
}

func (s *orderService) CreateOrder(order entity.Order) (int64, *resterrors.RestErr) {

	newOrderID, err := s.svc.db.Order().CreateOrder(order)
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(order.Products); i++ {
		price, err := s.svc.db.Order().CreateOrderProductAndReturnProductPrice(newOrderID, order.Products[i])
		if err != nil {
			return 0, err
		}

		order.TotalPrice += (price * float64(order.Products[i].Quantity))
	}

	if order.AcceptTip {
		order.TotalTip = order.TotalPrice * domain.DefaultTipPercent
	}

	err = s.svc.db.Order().UpdateOrder(newOrderID, order)
	if err != nil {
		return 0, err
	}

	return newOrderID, nil
}

func (s *orderService) GetOrdersByUserID(userID int64) (*[]entity.OrdersByUserID, *resterrors.RestErr) {

	orders, err := s.svc.db.Order().GetOrdersByUserID(userID)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *orderService) GetOrderDetail(orderID int64) (*entity.OrderDetail, *resterrors.RestErr) {

	order, err := s.svc.db.Order().GetOrderDetail(orderID)
	if err != nil {
		return nil, err
	}

	order.SubTotal = order.TotalPrice - order.TotalTip
	order.SubTotal = priceutils.ToFixed(order.SubTotal, 2)

	order.ProductsDetail, err = s.svc.db.Order().GetOrderProducts(orderID)
	if err != nil {
		return nil, err
	}

	return order, nil
}
