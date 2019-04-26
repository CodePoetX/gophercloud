package container

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToContainerListQuery() (string, error)
}

// ToContainerListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToContainerListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if q != nil {
		return q.String(), err
	} else {
		return "", err
	}

}

// List returns a Pager which allows you to iterate over a collection of
// containers. It accepts a ListOpts struct.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToContainerListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return ContainerPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// operate a container, Operation method includes [start,stop,pause,unpause,rebuild,reboot,rename,put_archive,
//add_securtiy_group,commit,network_detach]
// 	Request was accepted for processing, but the processing has not been completed.
// 	A ‘location’ header is included in the response which contains a link to check the progress of the request.
func Operatecontainer(c *gophercloud.ServiceClient, id, operation string) (r OperateResult) {
	url := OperatingURL(c, id, operation)
	_, r.Err = c.Post(url, "", &r.Body, nil)
	return
}

// Get retrieves a specific container based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, id), &r.Body, nil)
	if r.Err == nil {
		containerinfo, extracterr := r.Extract()
		if extracterr == nil {
			if containerinfo.Status == "Created" {
				startResult := Operatecontainer(c, containerinfo.UUID, "start")
				if startResult.Err != nil {
					r.Err = startResult.Err
					return
				}
			}
		}
	}
	return r
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToContainerCreateMap() (map[string]interface{}, error)
}

// ToContainerCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToContainerCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	if opts.RestartPolicy.Name == "" {
		opts.RestartPolicy.Name = "no"
	}

	b["result"] = opts

	return b, nil
	//return gophercloud.BuildRequestBody(opts, "container")
}

// Create accepts a CreateOpts struct and creates a new container using the values
// provided.
//
// The tenant ID that is contained in the URI is the tenant that creates the
// container. An admin user, however, has the option of specifying another tenant
// ID in the CreateOpts struct.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult, err error) {
	b, err := opts.ToContainerCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	//fmt.Printf("%v\n",b["result"])
	_, r.Err = c.Post(createURL(c), b["result"], &r.Body, nil)
	if r.Err != nil {
		err = fmt.Errorf("create container is error : %s", r.Err.Error())
		return
	}
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToContainerUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents options used to update a container.
type UpdateOpts struct {
	Memory string  `json:"memory"`
	Cpu    float64 `json:"cpu"`
	Name   string  `json:"name"`
}

// ToNContainerUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToContainerUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Update accepts a UpdateOpts struct and updates an existing container using the
// values provided. For more information, see the Create function.
func Update(c *gophercloud.ServiceClient, containerID string, opts UpdateOptsBuilder) (r UpdateResult) {
	/*b, err := opts.ToContainerUpdateMap()
	if err != nil {
		r.Err = err
		return
	}*/
	_, r.Err = c.Patch(updateURL(c, containerID), opts, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete accepts a unique ID and deletes the container associated with it.
func Delete(c *gophercloud.ServiceClient, containerID string, force bool) (r DeleteResult) {
	if force {
		_, r.Err = c.Delete(forcedeleteURL(c, containerID), nil)
	} else {
		_, r.Err = c.Delete(deleteURL(c, containerID), nil)
	}
	return
}
