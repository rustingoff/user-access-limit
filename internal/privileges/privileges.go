package privileges

const (
	CreateDriver = iota
	UpdateDriver
	DeleteDriver
	ViewDrivers
	CreateOrder
	UpdateOrder
	DeleteOrder
	ViewOrders
	CreateCustomer
	UpdateCustomer
	DeleteCustomer
	ViewCustomers
)

var Admin = map[uint]bool{
	CreateDriver:   true,
	UpdateDriver:   true,
	DeleteDriver:   true,
	ViewDrivers:    true,
	CreateOrder:    true,
	UpdateOrder:    true,
	DeleteOrder:    true,
	ViewOrders:     true,
	CreateCustomer: true,
	UpdateCustomer: true,
	DeleteCustomer: true,
	ViewCustomers:  true,
}
