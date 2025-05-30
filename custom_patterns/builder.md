/*
builder pattern is useful for scenarios when the member variables are optional for a particular object
*/
type Product struct {
    FieldA string
    FieldB int
    FieldC bool
}

type ProductBuilder struct {
    product *Product
}

func NewProductBuilder() *ProductBuilder {
    return &ProductBuilder{product: &Product{}}
}

func (b *ProductBuilder) SetFieldA(val string) *ProductBuilder {
    b.product.FieldA = val
    return b
}

func (b *ProductBuilder) SetFieldB(val int) *ProductBuilder {
    b.product.FieldB = val
    return b
}

func (b *ProductBuilder) SetFieldC(val bool) *ProductBuilder {
    b.product.FieldC = val
    return b
}

func (b *ProductBuilder) Build() *Product {
    return b.product
}
