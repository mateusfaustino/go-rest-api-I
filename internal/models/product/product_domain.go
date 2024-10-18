package product_domain

type ProductDomainInterface interface {
	GetID() int64
	SetID(int64)
	GetName() string
	SetName(string)
	GetPrice() float64
	SetPrice(float64)
}

type ProductDomain struct {
	Id    int64
	Name  string
	Price float64
}

func NewProductDomain(name string, price float64) ProductDomainInterface {
	return &ProductDomain{
		Name:  name,
		Price: price,
	}
}

func (pd *ProductDomain) GetID() int64 {
	return pd.Id
}

func (pd *ProductDomain) SetID(id int64) {
	pd.Id = id
}

func (pd *ProductDomain) GetName() string {
	return pd.Name
}

func (pd *ProductDomain) SetName(name string) {
	pd.Name = name
}

func (pd *ProductDomain) GetPrice() float64 {
	return pd.Price
}

func (pd *ProductDomain) SetPrice(price float64) {
	pd.Price = price
}
