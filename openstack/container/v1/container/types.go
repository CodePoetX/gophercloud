package container

type Container struct {
	// The Container IP addresses
	Addresses map[string][]Address `json:"addresses"`

	// UUID for the container
	UUID string `json:"uuid"`

	// User ID for the container
	UserID string `json:"user_id"`

	// Project ID for the container
	ProjectID string `json:"project_id"`

	// cpu for the container
	CPU float64 `json:"cpu"`

	// Memory for the container
	Memory string `json:"memory"`

	// Image for the container
	Image string `json:"image"`

	// The container container
	Labels map[string]string `json:"labels"`

	//// The created time of the container
	//CreatedAt time.Time `json:"-"`
	//
	//// The updated time of the container
	//UpdatedAt time.Time `json:"-"`
	//
	//// The started time of the container
	//StartedAt time.Time `json:"-"`

	// Name for the container
	Name string `json:"name"`

	// Links includes HTTP references to the itself, useful for passing along to
	// other APIs that might want a capsule reference.
	Links []Link `json:"links"`

	// auto remove flag token for the container
	AutoRemove bool `json:"auto_remove"`

	// Host for the container
	Host string `json:"host"`

	// Work directory for the container
	WorkDir string `json:"workdir"`

	// Disk for the container
	Disk int `json:"disk"`

	// Image pull policy for the container
	ImagePullPolicy string `json:"image_pull_policy"`

	// Task state for the container
	TaskState string `json:"task_state"`

	// Host name for the container
	HostName string `json:"hostname"`

	// Environment for the container
	Environment map[string]string `json:"environment"`

	// Status for the container
	Status string `json:"status"`

	// Auto Heal flag for the container
	AutoHeal bool `json:"auto_heal"`

	// Status details for the container
	StatusDetail string `json:"status_detail"`

	// Status reason for the container
	StatusReason string `json:"status_reason"`

	// Image driver for the container
	ImageDriver string `json:"image_driver"`

	// Command for the container
	Command []string `json:"command"`

	// Image for the container
	Runtime string `json:"runtime"`

	// Interactive flag for the container
	Interactive bool `json:"interactive"`

	// Restart Policy for the container
	RestartPolicy map[string]string `json:"restart_policy"`

	// Ports information for the container
	Ports []int `json:"ports"`

	// Security groups for the container
	SecurityGroups []string `json:"security_groups"`

	Privileged bool `json:"privileged"`

	HealthCheck healthcheckdetail `json:"healthcheck"`
}

type ListOpts struct {
	Name        string `q:"name"`
	Image       string `q:"image"`
	ProjectId   string `q:"project_id "`
	UserId      string `q:"user_id "`
	Memory      int    `q:"memory "`
	Host        string `q:"host "`
	TaskState   string `q:"task_state"`
	Status      string `q:"status "`
	Auto_remove string `q:"auto_remove "`
}

type CreateOpts struct {

	//The name of the container.
	Name string `json:"name,omitempty"`

	//The name or ID of the image.
	Image string `json:"image" required:"true"`

	//	Send command to the container.
	Command string `json:"command,omitempty"`

	//	The number of virtual cpus.
	Cpu float64 `json:"cpu,omitempty"`

	//The container memory size in MiB.
	Memory int `json:"memory,omitempty"`

	//	The working directory for commands to run in.
	Workdir string `json:"workdir,omitempty"`

	//A list of networks for the container.
	Nets []Net `json:"nets,omitempty"`

	/*The policy which determines if the image should be pulled prior to starting the container.
	Allowed values are ifnotpresent that means pull the image if it does not already exist on the node,
	always means always pull the image from repository and never mean never pull the image.
	*/
	ImagePullPolicy string `json:"image_pull_policy,omitempty"`

	//Adds a map of labels to a container.
	Labels map[string]string `json:"labels,omitempty"`

	//The environment variables.
	Environment map[string]string `json:"environment,omitempty"`

	/*
		Restart policy to apply when a container exits.
		It must contain a Name key and its allowed values are no, on-failure, always, unless-stopped.
		Optionally, it can contain a MaximumRetryCount key and its value is an integer.
	*/
	RestartPolicy restart_policy `json:"restart_policy,omitempty"`

	//Keep STDIN open even if not attached, allocate a pseudo-TTY.
	Interactive *bool `json:"interactive,omitempty"`

	/*
		The image driver to use to pull container image.
		Allowed values are docker to pull the image from Docker Hub and glance to pull the image from Glance.
	*/
	ImageDriver string `json:"image_driver,omitempty"`

	//Security groups to be added to the container.
	SecurityGroups string `json:"security_groups,omitempty"`

	/*
		The container runtime tool to create container with.
		You can use the default runtime that is runc or any other runtime configured to use with Docker.
	*/
	Runtime string `json:"runtime,omitempty"`

	//The hostname of container.
	Hostname string `json:"hostname,omitempty"`

	//enable auto-removal of the container on daemon side when the container’s process exits.
	Auto_remove *bool `json:"auto_remove,omitempty"`

	//The flag of healing non-existent container in docker.
	Auto_heal *bool `json:"auto_heal,omitempty"`

	/*
		The availability zone from which to run the container.
		Typically, an admin user will use availability zones to arrange container hosts into logical groups.
		An availability zone provides a form of physical isolation and redundancy from other availability zones.
		For instance, if some racks in your data center are on a separate power source,
		you can put containers in those racks in their own availability zone.
		By segregating resources into availability zones,
		you can ensure that your application resources are spread across disparate machines to achieve high availability
		in the event of hardware or other failure.
		You can see the available availability zones by calling the services API.

	*/
	Availability_zone string `json:"availability_zone,omitempty"`

	//The dictionary of data to send to the scheduler.
	Hints map[string]string `json:"hints,omitempty"`

	/*
		A list of dictionary data to specify how volumes are mounted into the container.
	*/
	Mounts []mount `json:"mounts,omitempty"`

	//Give extended privileges to the container.
	Privileged *bool `json:"privileged,omitempty"`

	//A dict of health check for the container.
	//Healthcheck healthcheckdetail `json:"healthcheck,omitempty"`
	//exposed_ports needs to implement

}

type Label struct {
	App string `json:"app"`
}

type restart_policy struct {
	/*Name string `json:"Name" required:"true"`
	MaximumRetryCount int `json:"MaximumRetryCount" required:"true"`*/
	Name              string `json:"Name" `
	MaximumRetryCount int    `json:"MaximumRetryCount" `
}

/*
	When you do not specify the nets parameter, the container attaches to the only network created for the current tenant.
	To provision the container with a NIC for a network, specify the UUID or name of the network in the network attribute.
	To provision the container with a NIC for an already existing port, specify the port id or name in the port attribute.

    If multiple networks are defined, the order in which they appear in the container will not necessarily
	reflect the order in which they are given in the request. Users should therefore not depend on device order
	to deduce any information about their network devices.
*/
type Net struct {
	V4FixedIp *string `json:"v4-fixed-ip,omitempty"`
	Network   string  `json:"network,omitempty"`
	V6FixedIp *string `json:"v6-fixed-ip,omitempty"`
	Port      string  `json:"port,omitempty"`
}

/*
The container will mount the volumes at create time.
	Each item can have an type attribute that specifies the volume type.
	The supported volume types are volume or bind.
	If this attribute is not specified, the default is volume.
	To provision a container with pre-existing Cinder volumes bind-mounted,
	specify the UUID or name of the volume in the source attribute.
	Alternatively, Cinder volumes can be dynamically created.
	In this case, the size of the volume needs to be specified in the size attribute.
	Another option is to mount a user-provided file into the container.
	In this case, the type attribute should be ‘bind’ and the content of the file is contained in the source attribute.
	The volumes will be mounted into the file system of the container and
	the path to mount the volume needs to be specified in the destination attribute.
*/
type mount struct {
	Source      string `json:"source,omitempty"`
	Destination string `json:"destination,omitempty"`
	Type        string `json:"type,omitempty"`
	Size        string `json:"size,omitempty"`
}

/*
Specify a test command to perform to check that the container is healthy. Four parameters are supported:

cmd: Command to run to check health.
interval: Time between running the check in seconds.
retries: Consecutive failures needed to report unhealthy.
timeout: Maximum time to allow one check to run in seconds.
*/
type healthcheckdetail struct {
	Cmd      string `json:"cmd,omitempty"`
	Interval int    `json:"interval,omitempty"`
	Retries  int    `json:"retries,omitempty"`
	Timeout  int    `json:"timeout,omitempty"`
}
