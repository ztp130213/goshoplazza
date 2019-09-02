package goshoplazza

import (
	"fmt"
	"time"
)

// FulfillmentService is an interface for interfacing with the fulfillment endpoints
// of the Shopify API.
// https://help.shopify.com/api/reference/fulfillment
type FulfillmentService interface {
	List(interface{}) ([]Fulfillment, error)
	Count(interface{}) (int, error)
	Get(string, interface{}) (*Fulfillment, error)
	Create(Fulfillment) (*Fulfillment, error)
	Update(Fulfillment) (*Fulfillment, error)
	Complete(string) (*Fulfillment, error)
	Transition(string) (*Fulfillment, error)
	Cancel(string) (*Fulfillment, error)
}

// FulfillmentsService is an interface for other Shopify resources
// to interface with the fulfillment endpoints of the Shopify API.
// https://help.shopify.com/api/reference/fulfillment
type FulfillmentsService interface {
	ListFulfillments(string, interface{}) ([]Fulfillment, error)
	CountFulfillments(string, interface{}) (int, error)
	GetFulfillment(string, string, interface{}) (*Fulfillment, error)
	CreateFulfillment(string, Fulfillment) (*Fulfillment, error)
	UpdateFulfillment(string, Fulfillment) (*Fulfillment, error)
	CompleteFulfillment(string, string) (*Fulfillment, error)
	TransitionFulfillment(string, string) (*Fulfillment, error)
	CancelFulfillment(string, string) (*Fulfillment, error)
}

// FulfillmentServiceOp handles communication with the fulfillment
// related methods of the Shopify API.
type FulfillmentServiceOp struct {
	client     *Client
	resource   string
	resourceID string
}

// Fulfillment represents a Shopify fulfillment.
type Fulfillment struct {
	ID        string     `json:"id,omitempty"`
	OrderID   string     `json:"order_id,omitempty"`
	Status    string     `json:"status,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Service             string     `json:"service,omitempty"`
	UpdatedAt           *time.Time `json:"updated_at,omitempty"`
	TrackingCompany     string     `json:"tracking_company,omitempty"`
	TrackingCompanyCode string     `json:"tracking_company_code,omitempty"`
	// ShipmentStatus  string     `json:"shipment_status,omitempty"`
	TrackingNumber string `json:"tracking_number,omitempty"`
	// TrackingNumbers []string   `json:"tracking_numbers,omitempty"`
	// TrackingUrl     string     `json:"tracking_url,omitempty"`
	// TrackingUrls    []string   `json:"tracking_urls,omitempty"`
	// Receipt         Receipt    `json:"receipt,omitempty"`
	LineItems []LineItem `json:"line_items,omitempty"`
	// NotifyCustomer  bool       `json:"notify_customer"`
	LineItemIDs []string `json:"line_item_ids,omitempty"`
}

// Receipt represents a Shopify receipt.
type Receipt struct {
	TestCase      bool   `json:"testcase,omitempty"`
	Authorization string `json:"authorization,omitempty"`
}

// FulfillmentResource represents the result from the fulfillments/X.json endpoint
type FulfillmentResource struct {
	Fulfillment *Fulfillment `json:"fulfillment"`
}

// FulfillmentsResource represents the result from the fullfilments.json endpoint
type FulfillmentsResource struct {
	Fulfillments []Fulfillment `json:"fulfillments"`
}

// List fulfillments
func (s *FulfillmentServiceOp) List(options interface{}) ([]Fulfillment, error) {
	prefix := FulfillmentPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s", prefix)
	resource := new(FulfillmentsResource)
	err := s.client.Get(path, resource, options)
	return resource.Fulfillments, err
}

// Count fulfillments
func (s *FulfillmentServiceOp) Count(options interface{}) (int, error) {
	prefix := FulfillmentPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/count", prefix)
	return s.client.Count(path, options)
}

// Get individual fulfillment
func (s *FulfillmentServiceOp) Get(fulfillmentID string, options interface{}) (*Fulfillment, error) {
	prefix := FulfillmentPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/%s", prefix, fulfillmentID)
	resource := new(FulfillmentResource)
	err := s.client.Get(path, resource, options)
	return resource.Fulfillment, err
}

// Create a new fulfillment
func (s *FulfillmentServiceOp) Create(fulfillment Fulfillment) (*Fulfillment, error) {
	prefix := FulfillmentPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s", prefix)
	// wrappedData := FulfillmentResource{Fulfillment: &fulfillment}
	resource := new(FulfillmentResource)
	err := s.client.Post(path, fulfillment, resource)
	return resource.Fulfillment, err
}

// Update an existing fulfillment
func (s *FulfillmentServiceOp) Update(fulfillment Fulfillment) (*Fulfillment, error) {
	prefix := FulfillmentPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/%d", prefix, fulfillment.ID)
	wrappedData := FulfillmentResource{Fulfillment: &fulfillment}
	resource := new(FulfillmentResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Fulfillment, err
}

// Complete an existing fulfillment
func (s *FulfillmentServiceOp) Complete(fulfillmentID string) (*Fulfillment, error) {
	prefix := FulfillmentPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/%d/complete", prefix, fulfillmentID)
	resource := new(FulfillmentResource)
	err := s.client.Post(path, nil, resource)
	return resource.Fulfillment, err
}

// Transition an existing fulfillment
func (s *FulfillmentServiceOp) Transition(fulfillmentID string) (*Fulfillment, error) {
	prefix := FulfillmentPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/%d/open", prefix, fulfillmentID)
	resource := new(FulfillmentResource)
	err := s.client.Post(path, nil, resource)
	return resource.Fulfillment, err
}

// Cancel an existing fulfillment
func (s *FulfillmentServiceOp) Cancel(fulfillmentID string) (*Fulfillment, error) {
	prefix := FulfillmentPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/%d/cancel", prefix, fulfillmentID)
	resource := new(FulfillmentResource)
	err := s.client.Post(path, nil, resource)
	return resource.Fulfillment, err
}
