package goshoplazza

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const variantsBasePath = "variants"

// VariantService is an interface for interacting with the variant endpoints
// of the Shopify API.
// See https://help.shopify.com/api/reference/product_variant
type VariantService interface {
	List(int64, interface{}) ([]Variant, error)
	Count(int64, interface{}) (int, error)
	Get(int64, interface{}) (*Variant, error)
	Create(int64, Variant) (*Variant, error)
	Update(Variant) (*Variant, error)
	Delete(int64, int64) error
}

// VariantServiceOp handles communication with the variant related methods of
// the Shopify API.
type VariantServiceOp struct {
	client *Client
}

// Variant represents a Shopify variant
type Variant struct {
	ID                   string           `json:"id,omitempty"`
	ProductID            string           `json:"product_id,omitempty"`
	Title                string           `json:"title,omitempty"`
	Sku                  string           `json:"sku,omitempty"`
	Position             int              `json:"position,omitempty"`
	Price                *decimal.Decimal `json:"price,omitempty"`
	CompareAtPrice       *decimal.Decimal `json:"compare_at_price,omitempty"`
	Option1              string           `json:"option1,omitempty"`
	Option2              string           `json:"option2,omitempty"`
	Option3              string           `json:"option3,omitempty"`
	CreatedAt            *time.Time       `json:"created_at,omitempty"`
	UpdatedAt            *time.Time       `json:"updated_at,omitempty"`
	image                Image            `json:"image,omitempty"`
	Barcode              string           `json:"barcode,omitempty"`
	InventoryQuantity    int              `json:"inventory_quantity,omitempty"`
	Weight               *decimal.Decimal `json:"weight,omitempty"`
	WeightUnit           string           `json:"weight_unit,omitempty"`
	Note                 string           `json:"note,omitempty"`
}

// VariantResource represents the result from the variants/X.json endpoint
type VariantResource struct {
	Variant *Variant `json:"variant"`
}

// VariantsResource represents the result from the products/X/variants.json endpoint
type VariantsResource struct {
	Variants []Variant `json:"variants"`
}

// List variants
func (s *VariantServiceOp) List(productID int64, options interface{}) ([]Variant, error) {
	path := fmt.Sprintf("%s/%s/%d/variants.json", globalApiPathPrefix, productsBasePath, productID)
	resource := new(VariantsResource)
	err := s.client.Get(path, resource, options)
	return resource.Variants, err
}

// Count variants
func (s *VariantServiceOp) Count(productID int64, options interface{}) (int, error) {
	path := fmt.Sprintf("%s/%s/%d/variants/count.json", globalApiPathPrefix, productsBasePath, productID)
	return s.client.Count(path, options)
}

// Get individual variant
func (s *VariantServiceOp) Get(variantID int64, options interface{}) (*Variant, error) {
	path := fmt.Sprintf("%s/%s/%d.json", globalApiPathPrefix, variantsBasePath, variantID)
	resource := new(VariantResource)
	err := s.client.Get(path, resource, options)
	return resource.Variant, err
}

// Create a new variant
func (s *VariantServiceOp) Create(productID int64, variant Variant) (*Variant, error) {
	path := fmt.Sprintf("%s/%s/%d/variants.json", globalApiPathPrefix, productsBasePath, productID)
	wrappedData := VariantResource{Variant: &variant}
	resource := new(VariantResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Variant, err
}

// Update existing variant
func (s *VariantServiceOp) Update(variant Variant) (*Variant, error) {
	path := fmt.Sprintf("%s/%s/%d.json", globalApiPathPrefix, variantsBasePath, variant.ID)
	wrappedData := VariantResource{Variant: &variant}
	resource := new(VariantResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Variant, err
}

// Delete an existing product
func (s *VariantServiceOp) Delete(productID int64, variantID int64) error {
	return s.client.Delete(fmt.Sprintf("%s/%s/%d/variants/%d.json", globalApiPathPrefix, productsBasePath, productID, variantID))
}