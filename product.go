package goshoplazza

import (
	"fmt"
	"time"
)

const productsBasePath = "products"
const productsResourceName = "products"

// ProductService is an interface for interfacing with the product endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/product
type ProductService interface {
	List(interface{}) ([]Product, error)
	Count(interface{}) (int, error)
	Get(string, interface{}) (*Product, error)
	Create(Product) (*Product, error)
	Update(Product) (*Product, error)
	Delete(string) error

	// MetafieldsService used for Product resource to communicate with Metafields resource
	// MetafieldsService
}

// ProductServiceOp handles communication with the product related methods of
// the Shopify API.
type ProductServiceOp struct {
	client *Client
}

// Product represents a Shopify product
type Product struct {
	ID                    string          `json:"id,omitempty"`
	Title                 string          `json:"title,omitempty"`
	Brief                 string          `json:"brief,omitempty"`
	Description           string          `json:"description,omitempty"`
	Vendor                string          `json:"vendor,omitempty"`
	VendorURL             string          `json:"vendor_url,omitempty"`
	HasOnlyDefaultVariant bool            `json:"has_only_default_variant"`
	RequiresShipping      bool            `json:"requires_shipping"`
	Taxable               bool            `json:"taxable"`
	InventoryTracking     bool            `json:"inventory_tracking"`
	InventoryPolicy       string          `json:"inventory_policy"`
	InventoryQuantity     int64           `json:"inventory_quantity"`
	Handle                string          `json:"handle,omitempty"`
	CreatedAt             *time.Time      `json:"created_at,omitempty"`
	UpdatedAt             *time.Time      `json:"updated_at,omitempty"`
	PublishedAt           *time.Time      `json:"published_at,omitempty"`
	Published             bool            `json:"published,omitempty"`
	Note                  string          `json:"note,omitempty"`
	MetaTitle             string          `json:"meta_title"`
	MetaDescription       string          `json:"meta_description"`
	MetaKeyword           string          `json:"meta_keyword"`
	NeedVariantImage      bool            `json:"need_variant_image"`
	Options               []ProductOption `json:"options,omitempty"`
	Variants              []Variant       `json:"variants,omitempty"`
	Image                 Image           `json:"image,omitempty"`
	Images                []Image         `json:"images,omitempty"`
}

// The options provided by Shopify
type ProductOption struct {
	ID        string   `json:"id,omitempty"`
	ProductID string   `json:"product_id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Position  int      `json:"position,omitempty"`
	Values    []string `json:"values,omitempty"`
}

// Represents the result from the products/X endpoint
type ProductResource struct {
	Product *Product `json:"product"`
}

// Represents the result from the products endpoint
type ProductsResource struct {
	Products []Product `json:"products"`
}

// List products
func (s *ProductServiceOp) List(options interface{}) ([]Product, error) {
	path := fmt.Sprintf("%s/%s", globalApiPathPrefix, productsBasePath)
	resource := new(ProductsResource)
	err := s.client.Get(path, resource, options)
	return resource.Products, err
}

// Count products
func (s *ProductServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/%s/count", globalApiPathPrefix, productsBasePath)
	return s.client.Count(path, options)
}

// Get individual product
func (s *ProductServiceOp) Get(productID string, options interface{}) (*Product, error) {
	path := fmt.Sprintf("%s/%s/%s.json", globalApiPathPrefix, productsBasePath, productID)
	resource := new(ProductResource)
	err := s.client.Get(path, resource, options)
	return resource.Product, err
}

// Create a new product
func (s *ProductServiceOp) Create(product Product) (*Product, error) {
	path := fmt.Sprintf("%s/%s", globalApiPathPrefix, productsBasePath)
	wrappedData := ProductResource{Product: &product}
	resource := new(ProductResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Product, err
}

// Update an existing product
func (s *ProductServiceOp) Update(product Product) (*Product, error) {
	path := fmt.Sprintf("%s/%s/%d", globalApiPathPrefix, productsBasePath, product.ID)
	wrappedData := ProductResource{Product: &product}
	resource := new(ProductResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Product, err
}

// Delete an existing product
func (s *ProductServiceOp) Delete(productID string) error {
	return s.client.Delete(fmt.Sprintf("%s/%s/%s.json", globalApiPathPrefix, productsBasePath, productID))
}

// List metafields for a product
// func (s *ProductServiceOp) ListMetafields(productID int64, options interface{}) ([]Metafield, error) {
// 	metafieldService := &MetafieldServiceOp{client: s.client, resource: productsResourceName, resourceID: productID}
// 	return metafieldService.List(options)
// }

// // Count metafields for a product
// func (s *ProductServiceOp) CountMetafields(productID int64, options interface{}) (int, error) {
// 	metafieldService := &MetafieldServiceOp{client: s.client, resource: productsResourceName, resourceID: productID}
// 	return metafieldService.Count(options)
// }

// // Get individual metafield for a product
// func (s *ProductServiceOp) GetMetafield(productID int64, metafieldID int64, options interface{}) (*Metafield, error) {
// 	metafieldService := &MetafieldServiceOp{client: s.client, resource: productsResourceName, resourceID: productID}
// 	return metafieldService.Get(metafieldID, options)
// }

// // Create a new metafield for a product
// func (s *ProductServiceOp) CreateMetafield(productID int64, metafield Metafield) (*Metafield, error) {
// 	metafieldService := &MetafieldServiceOp{client: s.client, resource: productsResourceName, resourceID: productID}
// 	return metafieldService.Create(metafield)
// }

// // Update an existing metafield for a product
// func (s *ProductServiceOp) UpdateMetafield(productID int64, metafield Metafield) (*Metafield, error) {
// 	metafieldService := &MetafieldServiceOp{client: s.client, resource: productsResourceName, resourceID: productID}
// 	return metafieldService.Update(metafield)
// }

// // // Delete an existing metafield for a product
// func (s *ProductServiceOp) DeleteMetafield(productID int64, metafieldID int64) error {
// 	metafieldService := &MetafieldServiceOp{client: s.client, resource: productsResourceName, resourceID: productID}
// 	return metafieldService.Delete(metafieldID)
// }
