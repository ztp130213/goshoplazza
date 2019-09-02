package goshoplazza

import (
	"fmt"
	"strings"
)

var rootDomain string = "myshoplaza.com"

// var rootDomain string = "preview.shoplazza.com"

// Return the full shop name, including .myshoplaza.com
func ShopFullName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.Trim(name, ".")
	if strings.Contains(name, rootDomain) {
		return name
	}
	return name + "." + rootDomain
}

// Return the short shop name, excluding .myshoplaza.com
func ShopShortName(name string) string {
	// Convert to fullname and remove the myshoplaza part. Perhaps not the most
	// performant solution, but then we don't have to repeat all the trims here
	// :-)
	return strings.Replace(ShopFullName(name), "."+rootDomain, "", -1)
}

// Return the Shop's base url.
func ShopBaseUrl(name string) string {
	name = ShopFullName(name)
	return fmt.Sprintf("https://%s", name)
}

// Return the prefix for a metafield path
func MetafieldPathPrefix(resource string, resourceID int64) string {
	var prefix string
	if resource == "" {
		prefix = fmt.Sprintf("%s/metafields", globalApiPathPrefix)
	} else {
		prefix = fmt.Sprintf("%s/%s/%d/metafields", globalApiPathPrefix, resource, resourceID)
	}
	return prefix
}

// Return the prefix for a fulfillment path
func FulfillmentPathPrefix(resource string, resourceID string) string {
	var prefix string
	if resource == "" {
		prefix = fmt.Sprintf("%s/fulfillments", globalApiPathPrefix)
	} else {
		prefix = fmt.Sprintf("%s/%s/%s/fulfillments", globalApiPathPrefix, resource, resourceID)
	}
	return prefix
}
