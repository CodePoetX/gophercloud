package container

import "github.com/gophercloud/gophercloud"

func getURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("container", id)
}

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("container")
}

// `listURL` is a pure function. `listURL(c)` is a URL for which a GET
// request will respond with a list of capsules in the service `c`.
func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("container")
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("containers", id)
}

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("containers")
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func forcedeleteURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id) + "?force=True"
}
