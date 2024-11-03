// BroadcastMessage contains information broadcasted by a device to announce its presence.
// Alias is the name of the device.
// Version is the software version running on the device.
// DeviceModel is the model of the device.
// DeviceType is the type of device (e.g., phone, computer).
// Fingerprint is a unique identifier for the device.
// Port is the network port the device listens on.
// Protocol is the communication protocol version.
// Download indicates if the device can download files.
// Announce indicates if the device is currently announcing its presence.

// DeviceInfo represents information about a device.
// Alias is the name of the device.
// Version is the version of the software.
// DeviceModel is the model of the device.
// DeviceType is the type of the device.
// Fingerprint is a unique identifier for the device.
// Port is the network port used by the device.
// Protocol is the communication protocol version.
// Download indicates if the device supports downloading.

// FileInfo represents metadata about a file.
// ID is a unique identifier for the file.
// FileName is the name of the file.
// Size is the size of the file in bytes.
// FileType is the MIME type or extension of the file.
// SHA256 is the optional SHA256 hash of the file.
// Preview is an optional preview representation of the file.

// Info contains basic information about a device.
// Alias is the name of the device.
// Version is the version of the software.
// DeviceModel is the model of the device (optional).
// DeviceType is the type of the device (optional).
// Fingerprint is a unique identifier for the device.
// Port is the network port used by the device.
// Protocol is the communication protocol version.
// Download indicates if the device supports downloading.

// PrepareReceiveRequest contains information required to prepare for receiving files.
// Info contains the sender's device information.
// Files is a map of file IDs to their corresponding FileInfo.

// PrepareReceiveResponse contains the response after preparing to receive files.
// SessionID is the identifier for the transfer session.
// Files maps file IDs to their corresponding tokens needed for download.

// Comment generated by copilot
package models

type BroadcastMessage struct {
	Alias       string `json:"alias"`
	Version     string `json:"version"`
	DeviceModel string `json:"deviceModel"`
	DeviceType  string `json:"deviceType"`
	Fingerprint string `json:"fingerprint"`
	Port        int    `json:"port"`
	Protocol    string `json:"protocol"`
	Download    bool   `json:"download"`
	Announce    bool   `json:"announce"`
}

type DeviceInfo struct {
	Alias       string `json:"alias"`
	Version     string `json:"version"`
	DeviceModel string `json:"deviceModel"`
	DeviceType  string `json:"deviceType"`
	Fingerprint string `json:"fingerprint"`
	Port        int    `json:"port"`
	Protocol    string `json:"protocol"`
	Download    bool   `json:"download"`
}
type FileInfo struct {
	ID       string `json:"id"`
	FileName string `json:"fileName"`
	Size     int64  `json:"size"`
	FileType string `json:"fileType"`
	SHA256   string `json:"sha256,omitempty"`
	Preview  string `json:"preview,omitempty"`
}
type Info struct {
	Alias       string `json:"alias"`
	Version     string `json:"version"`
	DeviceModel string `json:"deviceModel,omitempty"`
	DeviceType  string `json:"deviceType,omitempty"`
	Fingerprint string `json:"fingerprint"`
	Port        int    `json:"port"`
	Protocol    string `json:"protocol"`
	Download    bool   `json:"download"`
}
type PrepareReceiveRequest struct {
	Info  Info                `json:"info"`
	Files map[string]FileInfo `json:"files"`
}

type PrepareReceiveResponse struct {
	SessionID string            `json:"sessionId"`
	Files     map[string]string `json:"files"` // File ID to Token map
}

// 假设 SendModel 已定义如下
type SendModel struct {
	DeviceName string
	IP         string
}
