package water

// MacOSDriverProvider enumerates possible MacOS TUN/TAP implementations
type MacOSDriverProvider int

const (
	// MacOSDriverSystem refers to the default P2P driver
	MacOSDriverSystem MacOSDriverProvider = 0
	// MacOSDriverTunTapOSX refers to the third-party tuntaposx driver
	// see https://sourceforge.net/p/tuntaposx
	MacOSDriverTunTapOSX MacOSDriverProvider = 1
)

// Linux
// DevicePermissions determines the owner and group owner for the newly created
// interface.
type DevicePermissions struct {
	// Owner is the ID of the user which will be granted ownership of the
	// device.  If set to a negative value, the owner value will not be
	// changed.  By default, Linux sets the owner to -1, which allows any user.
	Owner uint

	// Group is the ID of the group which will be granted access to the device.
	// If set to a negative value, the group value will not be changed.  By
	// default, Linux sets the group to -1, which allows any group.
	Group uint
}

// PlatformSpecificParams defines parameters in Config that are specific to
// all OSs. A zero-value of such type is valid, yielding an interface
// with OS defined name.
type PlatformSpecificParams struct {
	// Name is the name for the interface to be used.
	//
	// Darwin
	// For TunTapOSXDriver, it should be something like "tap0".
	// For SystemDriver, the name should match `utun[0-9]+`, e.g. utun233
	//
	// Linux
	// This overrides the default name assigned by OS such as tap0 or tun0. A zero-value
	// of this field, i.e. an empty string, indicates that the default name should be used.
	//
	// Windows
	// InterfaceName is a friendly name of the network adapter as set in Control Panel.
	// Of course, you may have multiple tap0901 adapters on the system, in which
	// case we need a friendlier way to identify them.
	Name string

	// Darwin Only
	// Driver should be set if an alternative driver is desired
	// e.g. TunTapOSXDriver
	Driver MacOSDriverProvider

	// Linux
	// Persist specifies whether persistence mode for the interface device
	// should be enabled or disabled.
	Persist bool

	// Linux
	// Permissions, if non-nil, specifies the owner and group owner for the
	// interface.  A zero-value of this field, i.e. nil, indicates that no
	// changes to owner or group will be made.
	Permissions *DevicePermissions

	// Linux
	// MultiQueue specifies whether the multiqueue flag should be set on the
	// interface.  From version 3.8, Linux supports multiqueue tuntap which can
	// uses multiple file descriptors (queues) to parallelize packets sending
	// or receiving.
	MultiQueue bool

	// Windows Only
	// ComponentID associates with the virtual adapter that exists in Windows.
	// This is usually configured when driver for the adapter is installed. A
	// zero-value of this field, i.e., an empty string, causes the interface to
	// use the default ComponentId. The default ComponentId is set to tap0901,
	// the one used by OpenVPN.
	ComponentID string

	// Windows Only
	// Network is required when creating a TUN interface. The library will call
	// net.ParseCIDR() to parse this string into LocalIP, RemoteNetaddr,
	// RemoteNetmask. The underlying driver will need those to generate ARP
	// response to Windows kernel, to emulate an TUN interface.
	// Please note that Network must be same with IP and Mask that configured manually.
	Network string

	// Windows Only
	// Configure IP and DNS by device DHCP
	IsDHCP     bool
	DHCPServer string
	DNS1       string
	DNS2       string
}

var defaultPlatformSpecificParams = PlatformSpecificParams{
	ComponentID: "tap0901",
	Network:     "192.168.56.2/24",
}

func (p PlatformSpecificParams) baseOn(custom PlatformSpecificParams) PlatformSpecificParams {
	res := p
	if custom.Name != "" {
		res.Name = custom.Name
	}
	if custom.Driver != MacOSDriverSystem {
		res.Driver = custom.Driver
	}
	if custom.Persist {
		res.Persist = custom.Persist
	}
	if custom.Permissions != nil {
		res.Permissions = custom.Permissions
	}
	if custom.MultiQueue {
		res.MultiQueue = custom.MultiQueue
	}
	if custom.ComponentID != "" {
		res.ComponentID = custom.ComponentID
	}
	if custom.Network != "" {
		res.Network = custom.Network
	}
	if custom.IsDHCP {
		res.IsDHCP = custom.IsDHCP
	}
	if custom.DHCPServer != "" {
		res.DHCPServer = custom.DHCPServer
	}
	if custom.DNS1 != "" {
		res.DNS1 = custom.DNS1
	}
	if custom.DNS2 != "" {
		res.DNS2 = custom.DNS2
	}
	return res
}
