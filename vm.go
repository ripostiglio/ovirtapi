// Copyright (C) 2017 Battelle Memorial Institute
// All rights reserved.
//
// This software may be modified and distributed under the terms
// of the BSD-2 license.  See the LICENSE file for details.

package ovirtapi

import (
	"encoding/json"
	"fmt"
)

// DiskAttachment The underlying storage interface of disks communication with controller.
type DiskAttachment struct {
	Link
	// Defines whether the disk is active in the virtual machine it's attached to.
	Active string `json:"active,omitempty"`
	// Defines whether the disk is bootable.
	Bootable string `json:"bootable,omitempty"`
	// Free text containing comments about this object.
	Comment string `json:"comment,omitempty"`
	// A human-readable description in plain text.
	Description string `json:"description,omitempty"`
	// The type of interface driver used to connect the disk device to the virtual machine.
	Interface string `json:"interface,omitempty"`
	// The logical name of the virtual machine's disk, as seen from inside the virtual machine.
	LogicalName string `json:"logical_name,omitempty"`
	// A human-readable name in plain text.
	Name string `json:"name,omitempty"`
	// Defines whether the virtual machine passes discard commands to the storage.
	PassDiscard string `json:"pass_discard,omitempty"`
	// Indicates whether the disk is connected to the virtual machine as read only.
	ReadOnly string `json:"read_only,omitempty"`
	// Defines whether SCSI reservation is enabled for this disk.
	UsesSCSIReservation string `json:"uses_scsi_reservation,omitempty"`
	// The reference to the disk.
	Disk *Disk `json:"disk,omitempty"`
	// The reference to the template.
	Template *Template `json:"template,omitempty"`
	// The reference to the virtual machine.
	VM *VM `json:"vm,omitempty"`
}

// Bios ...
type Bios struct {
	BootMenu struct {
		Enabled string `json:"enabled"`
	} `json:"boot_menu"`
}

// Console Representation for serial console device.
type Console struct {
	Enabled string `json:"enabled,omitempty"`
}

// Core ...
type Core struct {
	Index  int `json:"index,omitempty,string"`
	Socket int `json:"socket,omitempty,string"`
}

// CustomProperty Custom property representation.
type CustomProperty struct {
	Name   int `json:"name,omitempty"`
	Regexp int `json:"regexp,omitempty"`
	Value  int `json:"value,omitempty"`
}

// VCPUPin ...
type VCPUPin struct {
	CPUSet string `json:"cpu_set,omitempty"`
	VCPU   int    `json:"vcpu,omitempty,string"`
}

// CPUTune ...
type CPUTune struct {
	VCPUPins []VCPUPin `json:"vcpu_pins,omitempty"`
}

// CPUTopology ...
type CPUTopology struct {
	Cores   int `json:"cores,omitempty,string"`
	Sockets int `json:"sockets,omitempty,string"`
	Threads int `json:"threads,omitempty,string"`
}

// CPU ...
type CPU struct {
	Architecture string       `json:"architecture,omitempty"`
	Cores        []Core       `json:"cores,omitempty"`
	CPUTune      *CPUTune     `json:"cpu_tune,omitempty"`
	Level        int          `json:"level,omitempty"`
	CPUMode      string       `json:"cpu_mode,omitempty"`
	Name         string       `json:"name,omitempty"`
	Speed        int          `json:"speed,omitempty"`
	Topology     *CPUTopology `json:"topology,omitempty"`
	Type         string       `json:"type,omitempty"`
}

// Certificate ...
type Certificate struct {
	Comment      string `json:"comment,omitempty"`
	Content      string `json:"content,omitempty"`
	Description  string `json:"description,omitempty"`
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Organization string `json:"organization,omitempty"`
	Subject      string `json:"subject,omitempty"`
}

// Display Represents a graphic console configuration.
type Display struct {
	// The IP address of the guest to connect the graphic console client to.
	Address string `json:"address,omitempty"`
	// Indicates if to override the display address per host.
	AllowOverride string `json:"allow_override,omitempty"`
	// The TLS certificate in case of a TLS connection.
	Certificate *Certificate `json:"certificate,omitempty"`
	// Indicates whether a user is able to copy and paste content from an external host into the graphic console.
	CopyPasteEnabled string `json:"copy_paste_enabled,omitempty"`
	// Returns the action that will take place when the graphic console is disconnected.
	DisconnectAction string `json:"disconnect_action,omitempty"`
	// Indicates if a user is able to drag and drop files from an external host into the graphic console.
	FileTransferEnabled string `json:"file_transfer_enabled,omitempty"`
	// The keyboard layout to use with this graphic console.
	KeyboardLayout string `json:"keyboard_layout,omitempty"`
	// The number of monitors opened for this graphic console.
	Monitors int `json:"monitors,omitempty,string"`
	// The port address on the guest to connect the graphic console client to.
	Port int `json:"port,omitempty,string"`
	// The proxy IP which will be used by the graphic console client to connect to the guest.
	Proxy string `json:"proxy,omitempty"`
	// The secured port address on the guest, in case of using TLS, to connect the graphic console client to.
	SecurePort int `json:"secure_port,omitempty,string"`
	// Indicates if to use one PCI slot for each monitor or to use a single PCI channel for all multiple monitors.
	SingleQxlPci string `json:"single_qxl_pci,omitempty"`
	// Indicates if to use smart card authentication.
	SmartcardEnabled string `json:"smartcard_enabled,omitempty"`
	// The graphic console protocol type.
	Type string `json:"type,omitempty"`
}

// GuestOperatingSystem Represents an operating system installed on the virtual machine.
type GuestOperatingSystem struct {
	// The architecture of the operating system, such as x86_64.
	Architecture string `json:"architecture,omitempty"`
	// Code name of the operating system, such as Maipo.
	Codename string `json:"codename,omitempty"`
	// Full name of operating system distribution.
	Distribution string `json:"distribution,omitempty"`
	// Family of operating system, such as Linux.
	Family string `json:"family,omitempty"`
	// Kernel version of the operating system.
	Kernel *Kernel `json:"kernel,omitempty"`
	// Version of the installed operating system.
	Version *Version `json:"version,omitempty"`
}

// HighAvailability Type representing high availability of a virtual machine.
type HighAvailability struct {
	Enabled  string `json:"enabled,omitempty"`
	Priority int    `json:"priority,omitempty,string"`
}

// Configuration ...
type Configuration struct {
	Data string `json:"data,omitempty"`
	Type string `json:"type,omitempty"`
}

// DNS Represents the DNS resolver configuration.
type DNS struct {
	// Array of hosts serving as search domains.
	SearchDomains []Host `json:"search_domains,omitempty"`
	// Array of hosts serving as DNS servers.
	Servers []Host `json:"servers,omitempty"`
}

// MAC Represents a MAC address of a virtual network interface.
type MAC struct {
	// MAC Address
	Address string `json:"address,omitempty"`
}

// NIC Represents a virtual machine NIC.
type NIC struct {
	// Defines how an IP address is assigned to the NIC.
	BootProtocol string `json:"boot_protocol,omitempty"`
	// Free text containing comments about this object.
	Comment string `json:"comment,omitempty"`
	// A human-readable description in plain text.
	Description string `json:"description,omitempty"`
	// A unique identifier.
	ID string `json:"id,omitempty"`
	// The type of driver used for the NIC.
	Interface string `json:"interface,omitempty"`
	// Defines if the NIC is linked to the virtual machine.
	Linked string `json:"linked,omitempty"`
	// The MAC address of the interface.
	MAC *MAC `json:"mac,omitempty"`
	// A human-readable name in plain text.
	Name string `json:"name,omitempty"`
	// Defines if the network interface should be activated upon operation system startup.
	OnBoot string `json:"on_boot,omitempty"`
	// Defines if the NIC is plugged in to the virtual machine.
	Plugged string `json:"plugged,omitempty"`
}

// NetworkConfiguration ...
type NetworkConfiguration struct {
	DNS  DNS   `json:"dns,omitempty"`
	NICs []NIC `json:"nics,omitempty"`
}

// AuthorizedKey ...
type AuthorizedKey struct {
	// Free text containing comments about this object.
	Comment string `json:"comment,omitempty"`
	// A human-readable description in plain text.
	Description string `json:"description,omitempty"`
	// A unique identifier.
	ID  string `json:"id,omitempty"`
	Key string `json:"key,omitempty"`
	// A human-readable name in plain text.
	Name string `json:"name,omitempty"`
}

// CloudInit ...
type CloudInit struct {
	AuthorizedKeys       []AuthorizedKey       `json:"authorized_keys,omitempty"`
	Files                []File                `json:"files,omitempty"`
	Host                 *Host                 `json:"host,omitempty"`
	NetworkConfiguration *NetworkConfiguration `json:"network_configuration,omitempty"`
	RegenerateSSHKeys    string                `json:"regenerate_ssh_keys,omitempty"`
	Timezone             string                `json:"timezone,omitempty"`
	Users                []User                `json:"users,omitempty"`
}

// File ...
type File struct {
	// Free text containing comments about this object.
	Comment string `json:"comment,omitempty"`
	Content string `json:"content,omitempty"`
	// A human-readable description in plain text.
	Description string `json:"description,omitempty"`
	// A unique identifier.
	ID string `json:"id,omitempty"`
	// A human-readable name in plain text.
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

// IP Represents the IP configuration of a network interface.
type IP struct {
	// The text representation of the IP address.
	Address string `json:"address,omitempty"`
	// The address of the default gateway.
	Gateway string `json:"gateway,omitempty"`
	// The network mask.
	Netmask string `json:"netmask,omitempty"`
	// The version of the IP protocol.
	Version string `json:"version,omitempty"`
}

// NICConfiguration ...
type NICConfiguration struct {
	BootProtocol string `json:"boot_protocol,omitempty"`
	IP           *IP    `json:"ip,omitempty"`
	Name         string `json:"name,omitempty"`
	OnBoot       string `json:"on_boot,omitempty"`
}

type NICConfigurations struct {
	NICConfiguration []NICConfiguration `json:"nic_configuration,omitempty"`
}

// Initialization ...
type Initialization struct {
	ActiveDirectoryOU string `json:"active_directory_ou,omitempty"`
	AuthorizedSSHKeys string `json:"authorized_ssh_keys,omitempty"`
	//TODO Finish structures from here
	CloudInit         *CloudInit         `json:"cloud_init,omitempty"`
	Configuration     *Configuration     `json:"configuration,omitempty"`
	CustomScript      string             `json:"custom_script,omitempty"`
	DNSSearch         string             `json:"dns_search,omitempty"`
	DNSServers        string             `json:"dns_servers,omitempty"`
	Domain            string             `json:"domain,omitempty"`
	HostName          string             `json:"HostName,omitempty"`
	InputLocale       string             `json:"input_locale,omitempty"`
	NICConfigurations *NICConfigurations `json:"nic_configurations,omitempty"`
	OrgName           string             `json:"org_name,omitempty"`
	RegenerateIDs     string             `json:"regenerate_ids,omitempty"`
	RegenerateSSHKeys string             `json:"regenerate_ssh_keys,omitempty"`
	RootPassword      string             `json:"root_password,omitempty"`
	SystemLocale      string             `json:"system_locale,omitempty"`
	Timezone          string             `json:"timezone,omitempty"`
	UILanguage        string             `json:"ui_language,omitempty"`
	UserLocale        string             `json:"user_locale,omitempty"`
	UserName          string             `json:"user_name,omitempty"`
	WindowsLicenseKey string             `json:"windows_license_key,omitempty"`
}

// IO ...
type IO struct {
	Threads int `json:"threads,omitempty,string"`
}

// MemoryOverCommit ...
type MemoryOverCommit struct {
	Percent int `json:"percent,omitempty,string"`
}

// MemoryPolicy Logical grouping of memory related properties of virtual machine-like entities.
type MemoryPolicy struct {
	Ballooning           string                `json:"ballooning,omitempty"`
	Guaranteed           int                   `json:"guaranteed,string,omitempty"`
	Max                  int                   `json:"max,string,omitempty"`
	OverCommit           *MemoryOverCommit     `json:"over_commit,omitempty"`
	TransparentHugePages *TransparentHugePages `json:"transparent_huge_pages,omitempty"`
}

// MigrationBandwidth Defines the bandwidth used by migration.
type MigrationBandwidth struct {
	// The method used to assign the bandwidth.
	AssignmentMethod string `json:"assignment_method,omitempty"`
	// Custom bandwidth in Mbps. Will be applied only if the assignmentMethod attribute is custom.
	CustomValue int `json:"custom_value,omitempty,string"`
}

// MigrationOptions The type for migration options.
type MigrationOptions struct {
	AutoConverge string `json:"auto_converge,omitempty"`
	// The bandwidth that is allowed to be used by the migration.
	Bandwidth  *MigrationBandwidth `json:"bandwidth,omitempty"`
	Compressed string              `json:"compressed,omitempty"`
}

// OperatingSystem Information describing the operating system. This is used for both virtual machines and hosts.
type OperatingSystem struct {
	Boot                  *Boot    `json:"boot,omitempty"`
	Cmdline               string   `json:"cmdline,omitempty"`
	CustomKernelCmdline   string   `json:"custom_kernel_cmdline,omitempty"`
	Initrd                string   `json:"initrd,omitempty"`
	Kernel                string   `json:"kernel,omitempty"`
	ReportedKernelCmdline string   `json:"reported_kernel_cmdline,omitempty"`
	Type                  string   `json:"type,omitempty"`
	Version               *Version `json:"version,omitempty"`
}

// Kernel ...
type Kernel struct {
	Version *Version `json:"version,omitempty"`
}

// Boot Configuration of the boot sequence of a virtual machine.
type Boot struct {
	Devices []string `json:"devices>device,omitempty"`
}

// Version ...
type Version struct {
	Build       int    `json:"build,omitempty,string"`
	Comment     string `json:"comment,omitempty"`
	Description string `json:"description,omitempty"`
	FullVersion string `json:"full_version,omitempty"`
	ID          string `json:"id,omitempty"`
	Major       int    `json:"major,omitempty,string"`
	Minor       int    `json:"minor,omitempty,string"`
	Name        string `json:"name,omitempty"`
	Revision    int    `json:"revision,omitempty,string"`
}

// TimeZone Time zone representation.
type TimeZone struct {
	Name      string `json:"name,omitempty"`
	UTCOffset string `json:"utc_offset,omitempty"`
}

// USB Configuration of the USB device of a virtual machine.
type USB struct {
	Enabled string `json:"enabled,omitempty"`
	Type    string `json:"type,omitempty"`
}

// VMPlacementPolicy ...
type VMPlacementPolicy struct {
	Affinity string `json:"affinity,omitempty"`
	Hosts    []Host `json:"hosts,omitempty"`
}

// VM Represents a virtual machine.
type VM struct {
	OvirtObject
	Comment                    string                `json:"comment,omitempty"`
	Console                    *Console              `json:"console,omitempty"`
	Bios                       *Bios                 `json:"bios,omitempty"`
	CPU                        *CPU                  `json:"cpu,omitempty"`
	CPUShares                  int                   `json:"cpu_shares,omitempty,string"`
	CreationTime               int                   `json:"creation_time,omitempty"`
	CustomCompatibilityVersion *Version              `json:"custom_compatibility_version,omitempty"`
	CustomCPUModel             string                `json:"custom_cpu_model,omitempty"`
	CustomEmulatedMachine      string                `json:"custom_emulated_machine,omitempty"`
	CustomProperties           []CustomProperty      `json:"custom_properties,omitempty"`
	DeleteProtected            string                `json:"delete_protected,omitempty"`
	Display                    *Display              `json:"display,omitempty"`
	FQDN                       string                `json:"fqdn,omitempty"`
	GuestOperatingSystem       *GuestOperatingSystem `json:"guest_operating_system,omitempty"`
	HighAvailability           *HighAvailability     `json:"high_availability,omitempty"`
	Initialization             *Initialization       `json:"initialization,omitempty"`
	Io                         *IO                   `json:"io,omitempty"`
	LargeIcon                  *Link                 `json:"large_icon,omitempty"`
	Memory                     int                   `json:"memory,string,omitempty"`
	MemoryPolicy               *MemoryPolicy         `json:"memory_policy,omitempty"`
	Migration                  *MigrationOptions     `json:"migration,omitempty"`
	MigrationDowntime          int                   `json:"migration_downtime,omitempty,string"`
	Origin                     string                `json:"origin,omitempty"`
	Os                         *OperatingSystem      `json:"os,omitempty"`
	SmallIcon                  *Link                 `json:"small_icon,omitempty"`
	StartPaused                string                `json:"start_paused,omitempty"`
	Stateless                  string                `json:"stateless,omitempty"`
	TimeZone                   *TimeZone             `json:"time_zone,omitempty"`
	Type                       string                `json:"type,omitempty"`
	USB                        *USB                  `json:"usb,omitempty"`
	Cluster                    *Cluster              `json:"cluster,omitempty"`
	CPUProfile                 *Link                 `json:"cpu_profile,omitempty"`
	Quota                      *Link                 `json:"quota,omitempty"`
	NextRunConfigurationExists string                `json:"next_run_configuration_exists,omitempty"`
	NumaTuneMode               string                `json:"numa_tune_mode,omitempty"`
	PlacementPolicy            *VMPlacementPolicy    `json:"placement_policy,omitempty"`
	Runonce                    string                `json:"run_once,omitempty"`
	Starttime                  int                   `json:"start_time,omitempty"`
	StopTime                   int                   `json:"stop_time,omitempty"`
	Status                     string                `json:"status,omitempty"`
	Host                       *Link                 `json:"host,omitempty"`
	InstanceType               *Link                 `json:"instance_type,omitempty"`
	OriginalTemplate           *Link                 `json:"original_template,omitempty"`
	Template                   *Template             `json:"template,omitempty"`
}

// CancelMigration This operation stops any migration of a virtual machine to another physical host.
func (vm *VM) CancelMigration() error {
	return vm.DoAction("cancelmigration", Action{})
}

// Clone Clones to a new VM
func (vm *VM) Clone(async string, newVM *VM) error {
	return vm.DoAction("clone", Action{
		Async: async,
		VM:    newVM,
	})
}

// CommitSnapshot Permanently restores the virtual machine to the state of the previewed snapshot.
func (vm *VM) CommitSnapshot(async string) error {
	return vm.DoAction("commitsnapshot", Action{
		Async: async,
	})
}

// Detach Detaches a virtual machine from a pool.
func (vm *VM) Detach() error {
	return vm.DoAction("detach", Action{})
}

// // Export Exports a virtual machine to an export domain.
// func (vm *VM) Export(async, discardSnapshots, exclusive string, storageDomain *StorageDomain) error {
// 	return vm.DoAction("export", Action {
// 		Async: async,
// 		DiscardSnapshots: discardSnapshots,
// 		Exclusive: exclusive,
// 		StorageDomain: storageDomain,
// 	})
// }

// FreezeFilesystems Freezes virtual machine file systems.
func (vm *VM) FreezeFilesystems(async string) error {
	return vm.DoAction("freezefilesystems", Action{
		Async: async,
	})
}

// Logon Initiates the automatic user logon to access a virtual machine from an external console.
func (vm *VM) Logon(async string) error {
	return vm.DoAction("logon", Action{
		Async: async,
	})
}

// Maintenance Sets the global maintenance mode on the hosted engine virtual machine.
func (vm *VM) Maintenance(async, maintenanceEnabled string) error {
	return vm.DoAction("maintenance", Action{
		Async:              async,
		MaintenanceEnabled: maintenanceEnabled,
	})
}

// Migrate Migrates a virtual machine to another physical host.
func (vm *VM) Migrate(async string, cluster *Cluster, force string, host *Host) error {
	return vm.DoAction("migrate", Action{
		Async:   async,
		Cluster: cluster,
		Force:   force,
		Host:    host,
	})
}

// Reboot Sends a reboot request to a virtual machine.
func (vm *VM) Reboot(async string) error {
	return vm.DoAction("reboot", Action{
		Async: async,
	})
}

// ReorderMACAddresses
func (vm *VM) ReorderMACAddresses(async string) error {
	return vm.DoAction("reordermacaddresses", Action{
		Async: async,
	})
}

// Shutdown This operation sends a shutdown request to a virtual machine.
func (vm *VM) Shutdown(async string) error {
	return vm.DoAction("shutdown", Action{
		Async: async,
	})
}

// Start Starts the virtual machine.
func (vm *VM) Start(async, filter, pause, useCloudInit, useSysprep string, nextBootVM *VM) error {
	return vm.DoAction("start", Action{
		Async:        async,
		Filter:       filter,
		Pause:        pause,
		UseCloudInit: useCloudInit,
		UseSysPrep:   useSysprep,
		VM:           nextBootVM,
	})
}

// Stop This operation forces a virtual machine to power-off.
func (vm *VM) Stop(async string) error {
	return vm.DoAction("stop", Action{
		Async: async,
	})
}

//Suspend This operation saves the virtual machine state to disk and stops it.
func (vm *VM) Suspend(async string) error {
	return vm.DoAction("suspend", Action{
		Async: async,
	})
}

// ThawFilesystems Thaws virtual machine file systems.
func (vm *VM) ThawFilesystems(async string) error {
	return vm.DoAction("thawfilesystems", Action{
		Async: async,
	})
}

// UndoSnapshot Restores the virtual machine to the state it had before previewing the snapshot.
func (vm *VM) UndoSnapshot(async string) error {
	return vm.DoAction("undosnapshot", Action{
		Async: async,
	})
}

// GetVM retrieve a VM from the server
func (con *Connection) GetVM(id string) (*VM, error) {
	body, err := con.GetLinkBody("vms", id)
	if err != nil {
		return nil, err
	}
	object := con.NewVM()
	err = json.Unmarshal(body, object)
	if err != nil {
		return nil, err
	}
	return object, err
}

// Update Synchronize the local VM with a copy from the server
func (vm *VM) Update() error {
	if vm.Href == "" {
		return fmt.Errorf("VM has not been saved to the server")
	}
	newVM, err := vm.Con.GetVM(vm.ID)
	if err != nil {
		return err
	}
	*vm = *newVM
	return nil
}

// GetAllVMs Retrieve all the VMs from the server
func (con *Connection) GetAllVMs() ([]*VM, error) {
	body, err := con.GetLinkBody("vms", "")
	if err != nil {
		return nil, err
	}
	objects := []*VM{}
	err = json.Unmarshal(body, &struct {
		VM *[]*VM
	}{&objects})
	if err != nil {
		return nil, err
	}
	for _, object := range objects {
		object.Con = con
	}
	return objects, err
}

// NewVM Create a new VM structure
func (con *Connection) NewVM() *VM {
	return &VM{OvirtObject: OvirtObject{Con: con}}
}

// Save Updates the server with the local copy of the VM
func (object *VM) Save() error {
	body, err := json.MarshalIndent(object, "", "    ")
	if err != nil {
		return err
	}
	// If there is a link, it is an already saved object, we need to update it
	if object.OvirtObject.Href != "" {
		body, err = object.Con.Request("PUT", object.Con.ResolveLink(object.Href), body)
		if err != nil {
			return err
		}
	} else {
		link, err := object.Con.GetLink("vms")
		if err != nil {
			return err
		}
		body, err = object.Con.Request("POST", link, body)
		if err != nil {
			return err
		}
	}
	tempObject := VM{OvirtObject: OvirtObject{Con: object.Con}}
	err = json.Unmarshal(body, &tempObject)
	if err != nil {
		return err
	}
	*object = tempObject
	return nil
}
