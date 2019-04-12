package container

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a container resource.
func (r commonResult) Extract() (*Container, error) {
	var s Container
	err := r.ExtractInto(&s)
	return &s, err

}

/*func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "container")
}*/

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Network.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Network.
type GetResult struct {
	commonResult
}
type ListResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Network.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

type ContainerPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of containers has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r ContainerPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}
func (c Container) GetIp() (ipv4 string) {
	for _, v := range c.Addresses {
		ipv4 = v[0].Addr
		return
	}
	return
}

// IsEmpty checks whether a NetworkPage struct is empty.
func (r ContainerPage) IsEmpty() (bool, error) {
	is, err := ExtractContainers(r)
	return len(is) == 0, err
}

// ExtractNetworks accepts a Page struct, specifically a NetworkPage struct,
// and extracts the elements into a slice of Network structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractContainers(r pagination.Page) ([]Container, error) {
	var s struct {
		Containers []Container `json:"containers"`
	}
	err := (r.(ContainerPage)).ExtractInto(&s)
	return s.Containers, err
}

func ExtractContainerInto(r pagination.Page, v interface{}) error {
	return r.(ContainerPage).Result.ExtractIntoSlicePtr(v, "container")
}
