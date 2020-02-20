// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/dns/cluster-role-binding.yaml (223B)
// assets/dns/cluster-role.yaml (210B)
// assets/dns/configmap.yaml (434B)
// assets/dns/daemonset.yaml (5.389kB)
// assets/dns/metrics/cluster-role-binding.yaml (279B)
// assets/dns/metrics/cluster-role.yaml (246B)
// assets/dns/metrics/role-binding.yaml (293B)
// assets/dns/metrics/role.yaml (284B)
// assets/dns/namespace.yaml (369B)
// assets/dns/service-account.yaml (85B)
// assets/dns/service.yaml (438B)

package manifests

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsDnsClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xce\x31\x8e\x83\x40\x0c\x05\xd0\x7e\x4e\xe1\x0b\xc0\x6a\xbb\xd5\x74\x9b\xdc\x80\x48\xe9\xcd\x8c\x09\x0e\x60\xa3\xb1\x87\x22\xa7\x8f\x10\x4a\x45\x3a\x17\xfe\xff\xfd\x89\x25\x47\xb8\xce\xd5\x9c\x4a\xa7\x33\x5d\x58\x32\xcb\x23\xe0\xca\x77\x2a\xc6\x2a\x11\x4a\x8f\xa9\xc5\xea\xa3\x16\x7e\xa1\xb3\x4a\x3b\xfd\x59\xcb\xfa\xb3\xfd\x86\x85\x1c\x33\x3a\xc6\x00\x00\x20\xb8\x50\x04\x5d\x49\x6c\xe4\xc1\x9b\x2c\x16\xac\xf6\x4f\x4a\x6e\x31\x34\x70\x78\x37\x2a\x1b\x27\xfa\x4f\x49\xab\x78\xf8\xc4\xf6\xe7\xe3\xb6\x15\xd3\xa9\xa7\xe8\x4c\x1d\x0d\x3b\x74\x9a\x1d\xbe\xd3\xef\x00\x00\x00\xff\xff\xfa\x62\xe7\x50\xdf\x00\x00\x00")

func assetsDnsClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsClusterRoleBindingYaml,
		"assets/dns/cluster-role-binding.yaml",
	)
}

func assetsDnsClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/cluster-role-binding.yaml", size: 223, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd9, 0xf6, 0x2a, 0x3b, 0x84, 0xd7, 0x3e, 0xc4, 0xe1, 0x70, 0x66, 0x31, 0xda, 0xc4, 0x2f, 0x53, 0x27, 0x29, 0x13, 0xfe, 0x80, 0x36, 0xc5, 0xa1, 0x70, 0xdc, 0x2d, 0xef, 0xcf, 0xe0, 0xc4, 0xeb}}
	return a, nil
}

var _assetsDnsClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8d\xbd\x6e\xc5\x30\x08\x46\x77\x3f\x05\xba\x7b\x52\x75\xab\xbc\x76\xe8\xde\xa1\x3b\xd7\xa6\x0a\x8a\x03\x16\xe0\x54\xea\xd3\x5f\xe5\x67\x3b\xe7\x08\xf4\xad\x2c\x35\xc3\x67\x1b\x1e\x64\xdf\xda\x28\x61\xe7\x1f\x32\x67\x95\x0c\xf6\xc4\x32\xe3\x88\x45\x8d\xff\x31\x58\x65\x5e\x3f\x7c\x66\x7d\xdb\xdf\xd3\x46\x81\x15\x03\x73\x02\x10\xdc\x28\x83\x76\x12\x5f\xf8\x37\xa6\x2a\x9e\x6c\x34\xf2\x9c\x26\xc0\xce\x5f\xa6\xa3\xfb\x71\x39\xc1\xe3\x91\x00\x8c\x5c\x87\x15\xba\x1b\x49\xed\xca\x12\x7e\x9a\x93\xed\x5c\xe8\x92\xae\xf5\x82\x63\xc3\x3b\x5e\x7d\x27\x7b\xde\xbf\x8d\x3d\x4e\xf8\xc3\x28\x4b\x7a\x05\x00\x00\xff\xff\xcb\xdd\xd7\x2a\xd2\x00\x00\x00")

func assetsDnsClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsClusterRoleYaml,
		"assets/dns/cluster-role.yaml",
	)
}

func assetsDnsClusterRoleYaml() (*asset, error) {
	bytes, err := assetsDnsClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/cluster-role.yaml", size: 210, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x37, 0xb2, 0x0, 0x7d, 0x4a, 0xd9, 0xf, 0x8, 0x44, 0xe7, 0xab, 0x82, 0xe4, 0x50, 0x94, 0xaa, 0x4e, 0xfd, 0xa0, 0x63, 0xba, 0x18, 0xcf, 0xeb, 0xa6, 0xe4, 0x2d, 0x4, 0x35, 0xd5, 0xc7, 0xd}}
	return a, nil
}

var _assetsDnsConfigmapYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x4e\xcd\x4e\xf3\x30\x10\xbc\xe7\x29\x46\xfa\xce\x5f\x4a\x15\x15\x89\x5c\x7b\xe6\xca\x7d\xb1\x27\x8d\x55\xc7\x36\xeb\x75\x11\x02\xde\x1d\xb5\x95\x82\x2a\x3a\xa7\xf9\xd3\xec\x1e\x43\xf2\x23\xf6\x39\x4d\xe1\xf0\x2c\xa5\x93\x12\x5e\xa8\x35\xe4\x34\xe2\xb4\xed\xfe\x21\xc9\x42\x48\xf2\x17\x52\x8b\x38\x42\x94\xa8\x34\x88\x41\x5b\xb2\xb0\xb0\xf3\x62\x32\x76\xc0\x3e\x2b\xa7\x10\x39\xe2\xab\x03\x80\x7e\xdc\x0d\xbb\x01\x9f\x17\x71\x06\x55\xb3\xd6\x55\xce\x94\x68\xf3\x2a\x8f\xed\x95\x9a\x68\xac\x70\xb1\x55\xa3\xf6\x31\x3b\x89\x08\xe9\xbf\x78\xaf\xbd\x68\x11\x84\xf2\x78\x25\xbf\xb3\x67\x94\xec\x2b\x42\xaa\x74\x4d\x79\x93\xb4\x52\x4d\x29\xcb\x8d\x39\x49\x8c\x36\x6b\x6e\x87\xf9\xfe\xfc\xda\xfe\x5e\x59\xd1\xbc\xd0\x66\xb6\x8a\xf1\x69\xbb\x1b\xd6\x60\xca\xfa\x2e\xea\xd1\x63\x43\x73\x1b\x65\xcd\xf1\xd4\xbb\x9c\xa6\x3f\x4f\xc6\xe0\x3e\x50\xf9\xd6\x98\x2c\x48\xbc\x73\xc5\x89\x9b\x89\xe1\x61\x35\x94\x31\x8b\xef\xae\xad\x9f\x00\x00\x00\xff\xff\x08\x52\xc6\xab\xb2\x01\x00\x00")

func assetsDnsConfigmapYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsConfigmapYaml,
		"assets/dns/configmap.yaml",
	)
}

func assetsDnsConfigmapYaml() (*asset, error) {
	bytes, err := assetsDnsConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/configmap.yaml", size: 434, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8c, 0x76, 0xcf, 0x25, 0x4e, 0xbd, 0x12, 0xda, 0x28, 0x18, 0x55, 0x70, 0xc2, 0x65, 0xe, 0x94, 0x40, 0x52, 0x25, 0x71, 0x45, 0xb8, 0xdf, 0xfe, 0x66, 0x8a, 0xa6, 0x6c, 0xa6, 0x90, 0xc7, 0x29}}
	return a, nil
}

var _assetsDnsDaemonsetYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x6d\x73\x1b\xb9\x0d\xfe\xee\x5f\x81\xac\x3c\x71\xd2\xf3\xfa\xa5\xb9\xa4\xd7\x4d\x7c\x3d\xd7\x96\x1b\xcf\x45\x96\xc6\x52\xee\x3e\x78\x3c\x19\x8a\x0b\x69\x51\x73\x49\x1e\xc9\x5d\x59\x63\xeb\xbf\x77\xb8\xab\x97\x7d\xab\x3b\x6d\xe7\x66\x2e\x1f\x14\x8b\x00\x1f\x02\x20\xf0\x00\xd4\x03\xc9\x38\x82\x4b\x86\xa9\x92\x63\x74\x7b\x4c\xd3\x2f\x68\x2c\x29\x19\x01\xd3\xda\x1e\xe7\xa7\x7b\x3d\x90\x2c\xc5\xc3\xe2\xd3\x6a\xc6\x11\x98\x8c\x41\xb0\x29\x0a\x0b\xcc\x20\x58\x74\xc0\x1c\x98\x4c\x3a\x4a\x71\xcf\x6a\xe4\xd1\x1e\x80\xc3\x54\x0b\xe6\xd0\xff\x0d\xb0\x59\x2d\xfe\x46\x93\x13\xc7\x73\xce\x55\x26\xdd\x0d\x4b\x31\x82\x58\xda\xb5\x54\x1b\x52\x86\xdc\xf2\x42\x30\x6b\x4b\xa1\x5d\x5a\x87\x69\x28\x55\x8c\x21\x37\xe4\x88\x33\xb1\xd6\xe6\x4a\x3a\x46\x12\x8d\xdd\xa0\x87\x85\xa5\x55\x44\x80\x1e\x50\xca\xe6\x08\x64\x9b\xd6\x6e\x34\x0a\xf9\x28\x13\x62\xa4\x04\xf1\x65\x04\xd7\xb3\x1b\xe5\x46\x06\x2d\x4a\xb7\xd5\x72\x68\x52\x92\xcc\x91\x92\x03\xb4\xd6\x6f\x59\xab\x5f\x31\x21\xa6\x8c\x3f\x4c\xd4\x17\x35\xb7\x43\xd9\x37\x46\x99\xed\x3e\xae\xd2\x94\xf9\x50\xdf\x41\xc0\x95\xc1\x58\xda\x00\xee\xb7\x62\x66\xe6\xb6\x90\x85\x5c\xc9\x59\x70\x08\xc1\x31\x3a\x7e\xbc\xd6\x3c\xbe\x50\x06\x67\x24\xb0\xba\x25\x57\x22\x4b\x71\xe0\x03\xb8\xf5\x7c\xe7\xbb\x87\xa1\x79\x58\x2a\x6d\xa5\x00\xa9\xd7\x1f\x31\x97\x44\x50\x3d\xa1\xa2\x61\x90\xc5\x43\x29\x96\x11\x38\x93\xed\xb6\x6a\x65\xea\xe7\x6c\xe3\x3e\x52\xc6\x45\xf0\xfe\xdd\xfb\x77\x15\x94\xf6\x0d\xf8\x7b\x55\x4e\x71\x25\x22\xf8\x7a\x39\xfa\xef\x91\x42\xc7\x75\x27\xda\xe4\xe2\x05\xb4\xbf\x9e\x76\xa0\xa5\xe8\x0c\xf1\x6e\xdb\xaa\x68\x3e\x16\x24\xd1\xda\x91\x51\x53\x8c\x2a\xfa\x89\x73\xfa\x1f\xe8\xaa\x4b\x00\xba\x8c\x6b\x82\x4c\xb8\xa4\x2e\x29\x6c\xf9\xe1\xe4\x87\x93\xda\xb2\xe5\x09\x7a\x7b\x3e\x4f\x26\xa3\x8a\x80\x24\x39\x62\xe2\x12\x05\x5b\x8e\x91\x2b\x19\xdb\x08\x4e\xab\x5b\x35\x1a\x52\x71\xb7\xcc\x66\x9c\xa3\xb5\x93\xc4\xa0\x4d\x94\x88\x23\x38\xad\x48\x67\x8c\x44\x66\xb0\x22\xad\x86\xc7\x57\x84\xca\x5c\x17\xb0\xa0\x1c\xff\x20\xa1\xf8\x70\xf2\x82\xc9\xef\xff\x8f\x50\xbc\xaf\xdc\xbc\x55\x99\xe1\x68\xab\x6e\x09\x4a\xc9\xd9\xba\xa3\x29\xa6\xca\x2c\x23\x78\x7f\xfa\xe7\x01\xd5\xca\xe8\xb7\x0c\x6d\x53\x9b\xeb\xcc\x07\xf5\x24\xed\xc4\xf8\xcb\xc9\x16\xa2\xc2\x61\x25\xeb\x79\x83\x44\x8e\xe6\x0f\xc3\x68\x16\x79\x56\xb0\xb4\x92\x0e\x1f\x6b\xd7\xaf\x0d\xe5\x24\x70\x8e\x71\x83\x44\x5e\xe6\xac\x44\x59\x67\x43\xcf\x73\x2f\x10\x56\xa1\x54\x09\x02\xca\x1c\x6e\xce\x07\xfd\x71\xff\xf6\x97\xfe\x6d\xd1\x99\x2e\xbe\x7c\x1d\x4f\xfa\xb7\xdf\x2e\x87\x83\xf3\xeb\x9b\xae\x0e\xb5\xd9\x8e\x32\x6f\x9b\xe1\x91\xae\x2f\xfa\xe3\x8a\x11\x3d\xb8\xf0\xfc\x0d\xca\x40\xd9\x00\x2d\x6a\x66\x98\xc3\x18\x04\x59\x07\x6a\xb6\x69\x69\xb6\xb6\xeb\x66\x38\xe9\x47\x70\xa5\x0c\x48\xb5\x38\x04\x94\x36\x33\x08\x2e\x41\x8b\x85\x59\x06\x05\x73\x94\x63\xd9\x5a\x3f\xc2\x4c\x19\x40\xc6\x93\xba\xe0\xb0\x86\xc9\x24\x30\x41\xcc\xc2\x82\x5c\xe2\xb1\x9a\xfe\xda\x6c\x36\xa3\x47\x58\x90\x10\xc0\x84\x55\x30\x45\x60\x71\x8c\xf1\x51\x05\x27\x67\x22\xc3\x08\x82\x22\x47\x42\x83\x73\xb2\xce\x2c\x8f\x94\x46\x69\x13\x9a\xb9\xb0\x21\xb0\x39\x0f\x5a\xcd\xac\x12\xba\xe3\x29\xc9\xe3\x29\xb3\x49\x65\x2d\xe4\x95\x2f\xcf\x55\x27\x5e\xb5\xd5\xa1\xb8\xa3\x30\x53\xa0\x49\xa3\x2f\xcd\xbd\x6a\x91\x1b\xa6\xe1\xe0\x9f\x6a\x6a\x21\xd4\xf0\x0c\x8f\xbe\x61\xc2\x83\x77\xf1\xf9\xb9\xc8\xb1\x8f\xb0\x60\xe4\x3e\x02\x3e\x92\x83\x93\x03\x98\xf4\x6f\x07\x55\x84\xe1\xa8\x7f\x33\xfe\x7c\x7d\x35\xf9\x36\x38\xbf\xfd\xb9\x7f\x7b\x16\xec\x7c\x9d\xa3\xc4\xe2\x36\xeb\xa5\x16\x54\xb6\x7f\x1e\x8e\x27\xe3\x6f\x57\xd7\x5f\xfa\x67\xc1\x2e\x0f\xab\x1a\x93\xfe\x60\xd4\x52\x38\x72\xa9\x0e\xaa\x66\x5c\x5f\x8d\xcf\x0e\x0e\xe1\xa0\x68\x2d\x10\x1a\x08\xd9\x36\x75\xe0\xd3\xa7\x4f\x10\xec\x3f\x6d\x12\x70\x55\xdb\xd9\x83\x01\x7b\x40\x60\xc5\x58\xa5\x0c\x33\x4b\xf0\xa5\xb2\x4b\x03\x25\xe2\xb2\x84\x8a\xf5\x03\x0b\xcc\x39\x43\xd3\xcc\xa1\xad\xde\x3c\xd7\x10\xce\x20\x0c\x77\xd2\x50\x49\xb1\xf4\x07\xef\x9c\x5c\x05\xfe\xfb\xd6\xa5\xba\x25\x8b\xc4\x9f\x5b\x06\x3d\x56\x35\x26\x8b\x91\x0b\x9f\xd8\xe1\x39\xd8\x9c\x7f\x23\x6d\x6b\x62\x9f\xdf\x36\xe7\x40\xd2\xc3\x6f\xfc\xbe\xfb\xe9\x7e\x15\xb4\xa0\xbc\xc7\x57\xe8\x78\xb2\x89\x0f\x5c\x8f\x60\x66\x54\x0a\x5c\x64\xd6\xa1\xf1\xdc\x08\x34\x03\x5d\x12\xda\x11\xfc\x8a\x90\xfa\x10\x59\xcc\xd1\x30\x01\xce\x10\xda\x16\xa6\x53\x10\x2b\x20\x17\xc1\xf5\x28\xff\xfe\xd0\x7f\x7e\x28\x3e\xbf\x07\x95\xa3\xf1\xfd\xbf\x60\x11\xbf\xbe\x5d\x39\x82\x49\x82\xe0\x16\x0a\x04\xf3\xf5\x2e\x3b\x80\xbd\xdf\xde\xc1\x18\xb5\x50\xcb\x14\xa5\x5b\xd7\xe8\xcf\x99\x59\x1a\x50\xd2\xdf\x10\x1a\x18\x6a\x94\x63\xc7\xf8\x03\xbc\x19\x8e\x47\xa7\xef\xde\x42\x08\x2e\x51\x16\xbd\x5d\x52\xb9\x16\xb0\xcd\xb4\x6f\x9a\x7e\x6a\x02\xa1\x58\x3c\x65\x82\x49\x8e\xc6\x16\x76\xfa\x46\x43\x05\x97\x30\x9e\x90\x9c\xc3\xe5\xcd\x18\x5c\x62\x54\x36\x4f\x0a\xd3\x1b\x78\x3c\x8d\xed\xd9\x9b\x83\x98\xe6\x10\x3a\x38\x87\x9f\x82\xfd\xa7\x1d\x81\xae\x02\xf8\xce\x26\xfe\x34\x7f\x41\x39\x5f\x1d\xed\x3f\xd5\xf9\x65\x15\x1c\x34\x10\xcb\x7f\x5b\xc4\xf3\xf3\xdf\x01\x14\xbe\x73\x5c\xff\x3e\xb6\xfe\xaf\xc8\x6f\x1b\xd0\xfe\xee\xc9\xa7\xf6\xfe\xd3\x2b\x1f\xe4\xbb\x3f\xdd\xaf\x1a\x2a\xad\x14\x07\x20\x6d\xcf\xde\xec\xbf\xc1\x9c\x09\x7f\x58\xb1\x91\xee\x57\xc1\xdb\x26\x3c\xf8\x5c\xbf\xbb\x83\x60\xff\x6f\x01\x84\xf8\x1b\x9c\xc0\xeb\xd7\x7e\x4b\x8f\x74\x59\x42\x10\x4a\x84\x13\xb8\xbf\xff\xe8\xf9\x40\x76\x78\xbe\xae\xc9\xbb\xb5\x57\xc1\xfd\x59\xb0\xff\xb4\xd9\xde\xa1\x3f\x35\xc8\x1e\x5a\xeb\x33\x6a\xb9\x25\x71\xaf\xb5\x50\x5b\xe9\xc1\x57\x1d\x33\x87\x95\x26\x0e\x05\xed\xd0\x0c\x16\x08\x73\x74\xbe\x25\x51\x5c\x29\x76\xdb\x00\xf8\x15\xcb\x9e\x26\x95\x83\xac\x05\xb6\x48\x50\x7a\xb7\x4d\x31\x11\xad\x9f\x35\x5b\x34\x95\x39\x3f\x2b\x29\x03\x4c\x13\x64\x92\xe5\x8c\x04\x9b\x92\x20\xb7\x6c\x1c\x33\x76\x4c\x20\xa0\x2c\xd8\x03\xb8\xca\x44\xec\x9b\x8a\x75\xfe\x6a\x2b\x07\xd2\xac\x60\xdd\xcd\x09\x64\x21\x46\x81\x0e\xe3\xbd\xae\x3b\x7b\xea\x6d\x62\xff\x9f\x6f\xaa\x07\x7f\xcf\x48\xc4\xc0\x40\xe2\xa2\x42\xe9\x25\xfb\x55\x7d\xf6\xec\xa2\x32\x03\x3c\xb3\x4e\xa5\x5b\xa3\x67\x24\x1c\x1a\x8c\xbd\xdb\x0d\xec\xb9\x41\x0d\x61\x0e\x41\x0f\xf6\x9f\x9a\x3d\xb1\x64\xfd\x5a\x17\xf8\xf1\x85\x3e\x50\xda\x7a\xae\x35\x16\x34\x54\x36\xcd\x9d\x11\x9e\xeb\xdb\x43\x11\xb4\xda\xc0\xab\x6a\x64\x3a\xda\xc0\xba\xb2\x74\x59\x5a\x1b\xe5\x32\x83\xef\x57\x9d\x1b\x00\x90\x27\x0a\x8a\xe4\x5e\x95\x9b\x36\xff\xb5\x2b\x19\xfe\x4d\x28\x7e\x6c\xf9\xde\x3c\xa4\x95\xf7\x5d\x99\xef\x63\x34\x19\x5e\x0e\xa3\x8e\x0a\x60\x4e\xa5\xc4\x99\x10\x4b\xdf\x96\x58\xae\x28\x06\x26\x97\x40\x92\x2b\x69\xc9\x3a\x94\x0e\xa6\x98\xb0\x9c\x2a\x93\xf7\x06\xf5\x16\xb5\xf0\xc3\x68\x57\x46\xa4\x2a\xa6\x19\x61\x0c\x79\xf9\x6b\x8e\x4f\x44\x89\x18\x37\xd2\xd3\xb7\x03\xdd\x70\xb3\x95\x03\xcf\xcf\xeb\xa1\xe1\x65\xbd\xb6\xd7\x1b\x5d\x5f\x1c\xbe\x6a\x0d\xa6\x2a\xc7\x78\xe7\x6b\x91\xd5\xdc\x20\x73\x78\x5c\x56\x4f\xd1\xd2\x76\xa3\x09\x70\xa5\x97\xc0\x93\xcc\xd4\x8b\xa4\x41\x41\x56\x20\x6a\xf8\x70\x02\xaf\x8b\x29\xb0\x26\xcb\xa4\x1f\x2c\xdb\xd3\x48\xed\xf2\x3a\x9f\x7b\x2f\x3e\xe1\x36\x2f\xb8\x58\xda\xcd\x7b\xe9\x12\x67\x2c\x13\x9b\xd3\xfd\x28\x39\x46\x81\xdc\x29\xb3\x43\x78\xc8\xa6\x68\x24\xfa\x99\x8c\xd4\xb1\xb2\x11\x08\x92\xd9\x63\x29\x5c\x6b\x95\xaf\xa4\xd6\xcf\x59\xdd\x3f\xe9\x94\xab\x03\xa6\xa3\xca\xa3\xe8\x86\xa5\x2f\x3d\x0c\x01\xc8\x61\x5a\xf3\x2b\x84\x07\x5c\x46\xb0\xf9\xa1\xa9\xe3\x2d\xdf\x10\xbd\xf0\x68\xf3\x4b\xc5\x8b\x6d\xaf\x89\xd1\xf1\x82\x03\x70\x4b\x8d\x11\x5c\xed\x10\x9c\x12\x7e\x1c\x27\x25\xb7\x26\xf6\x8a\xa9\xc6\x27\xb0\xf5\xd9\x63\x32\x09\x7e\xc8\x5b\x2e\x3c\xe1\x1f\xed\xd5\x7c\x08\xca\x21\x5e\x09\x3c\xaa\x47\x3b\x65\x7e\x6c\xdc\xe5\xa9\xd2\xfe\x18\x65\x22\xe8\x7b\x7e\xb7\x7b\xff\x0a\x00\x00\xff\xff\x90\x46\xec\x71\x0d\x15\x00\x00")

func assetsDnsDaemonsetYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsDaemonsetYaml,
		"assets/dns/daemonset.yaml",
	)
}

func assetsDnsDaemonsetYaml() (*asset, error) {
	bytes, err := assetsDnsDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/daemonset.yaml", size: 5389, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x54, 0xe8, 0x2, 0xef, 0xe7, 0x83, 0x3a, 0xa7, 0xff, 0xcb, 0x33, 0xd8, 0x58, 0x88, 0xe7, 0xa8, 0xba, 0xd, 0x67, 0x80, 0x43, 0xd4, 0xd, 0xa8, 0x6e, 0x48, 0x8f, 0x1c, 0xe9, 0x38, 0xdc, 0x18}}
	return a, nil
}

var _assetsDnsMetricsClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\xb1\x4a\x04\x41\x0c\x86\xfb\x79\x8a\xbc\xc0\xae\xd8\x1d\xd3\xa9\x85\xfd\x09\xf6\xb9\x99\x9c\x1b\x77\x27\x19\x92\xcc\x16\x3e\xbd\x2c\x8a\x08\xe2\xb5\x81\x7c\xdf\xff\xad\x2c\x35\xc3\xd3\x36\x3c\xc8\xce\xba\xd1\x23\x4b\x65\x79\x4b\xd8\xf9\x95\xcc\x59\x25\x83\x5d\xb0\xcc\x38\x62\x51\xe3\x0f\x0c\x56\x99\xd7\x93\xcf\xac\x77\xfb\x7d\x6a\x14\x58\x31\x30\x27\x00\xc1\x46\x19\xaa\xf8\xd4\x54\x38\xd4\x0e\x92\x8f\xcb\x3b\x95\xf0\x9c\x26\xf8\xd2\xbd\x90\xed\x5c\xe8\xa1\x14\x1d\x12\x3f\x7f\xdd\xb4\x51\x2c\x34\x7c\x5a\x4f\xfe\x7d\xf6\x8e\x85\x32\x68\x27\xf1\x85\xaf\xf1\x9b\x6c\xba\xd1\x99\xae\x87\xf9\x4f\xc7\x7f\x6b\x00\xb0\xf3\xb3\xe9\xe8\x37\xba\xd2\x67\x00\x00\x00\xff\xff\x5b\x52\x00\xaa\x17\x01\x00\x00")

func assetsDnsMetricsClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsClusterRoleBindingYaml,
		"assets/dns/metrics/cluster-role-binding.yaml",
	)
}

func assetsDnsMetricsClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/cluster-role-binding.yaml", size: 279, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x79, 0x95, 0x6f, 0xa4, 0xd5, 0xed, 0x48, 0x27, 0x41, 0x56, 0x5c, 0xea, 0x5c, 0x89, 0xdc, 0xc1, 0x44, 0x91, 0xd4, 0xb, 0x18, 0x85, 0x79, 0x75, 0xaa, 0x6e, 0xb5, 0x98, 0xbe, 0xc6, 0x33, 0x43}}
	return a, nil
}

var _assetsDnsMetricsClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcd\x31\x4b\x34\x41\x0c\x87\xf1\x7e\x3e\x45\xe0\xad\x77\x5f\xec\x64\x5a\x05\x3b\x0b\x05\xfb\xec\xce\xdf\xdb\x70\x3b\xc9\x90\x64\x0e\xf4\xd3\x8b\x70\xb6\x0f\x3f\x78\xfe\xd1\xd3\x39\x23\xe1\xe4\x76\x22\x48\x81\x86\x46\xdb\x17\x0d\xb7\x8e\x3c\x30\x83\xd2\x28\x76\xe7\x01\x7a\x7e\x7d\xa7\x8e\x74\xd9\x83\xa0\x6d\x98\x68\x16\x1e\xf2\x01\x0f\x31\xad\xe4\x1b\xef\x2b\xcf\x3c\xcc\xe5\x9b\x53\x4c\xd7\xeb\x63\xac\x62\xff\x6f\x0f\xe5\x2a\xda\xea\xdf\xf0\xcd\x4e\x94\x8e\xe4\xc6\xc9\xb5\x10\x29\x77\x54\x6a\x1a\x4b\x37\x95\x34\x17\xbd\x14\x9f\x27\xa2\x96\x85\x78\xc8\x8b\xdb\x1c\xf1\x4b\x17\xb2\x01\xe7\x34\x5f\x6d\x40\xe3\x90\xcf\x5c\xc5\x0a\x91\x23\x6c\xfa\x8e\x3b\x6b\x1a\x88\x42\x74\x83\x6f\xf7\x74\x41\x96\x9f\x00\x00\x00\xff\xff\x9f\xa8\x4d\x6c\xf6\x00\x00\x00")

func assetsDnsMetricsClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsClusterRoleYaml,
		"assets/dns/metrics/cluster-role.yaml",
	)
}

func assetsDnsMetricsClusterRoleYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/cluster-role.yaml", size: 246, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x64, 0xdb, 0xe0, 0x95, 0x65, 0xae, 0x53, 0x96, 0x3a, 0x5f, 0x5e, 0x8b, 0x69, 0xe2, 0x7d, 0x5, 0xbf, 0x1f, 0x3a, 0xf, 0xff, 0xd0, 0x6b, 0x23, 0x4f, 0xfd, 0x11, 0x7f, 0x57, 0xd4, 0x4a, 0x8b}}
	return a, nil
}

var _assetsDnsMetricsRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xce\xb1\x4e\xc4\x40\x0c\x04\xd0\x7e\xbf\xc2\x3f\x90\x20\xba\xd3\x76\xd0\xd0\x1f\x12\xbd\x6f\xd7\x97\x98\x64\xed\x95\xed\x4d\xc1\xd7\x23\xa4\x48\x54\x20\x5d\x3b\x9a\xd1\x1b\xec\xfc\x41\xe6\xac\x92\xc1\x6e\x58\x66\x1c\xb1\xaa\xf1\x17\x06\xab\xcc\xdb\xc5\x67\xd6\xa7\xe3\x39\x6d\x2c\x35\xc3\x55\x77\x7a\x65\xa9\x2c\x4b\x6a\x14\x58\x31\x30\x27\x00\xc1\x46\x19\xba\x69\xa3\x58\x69\xf8\xb4\x5d\xfc\x8c\xbd\x63\xa1\x0c\xda\x49\x7c\xe5\x7b\x4c\x55\x3c\x99\xee\x74\xa5\xfb\xcf\x14\x3b\xbf\x99\x8e\xfe\x8f\x9f\x00\x7e\xf9\xbf\x34\x1f\xb7\x4f\x2a\xe1\x39\x4d\x67\xfb\x9d\xec\xe0\x42\x2f\xa5\xe8\x90\x78\xf0\x65\x53\xe1\x50\x63\x59\x20\x7d\x07\x00\x00\xff\xff\xb9\xd9\xab\x8d\x25\x01\x00\x00")

func assetsDnsMetricsRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsRoleBindingYaml,
		"assets/dns/metrics/role-binding.yaml",
	)
}

func assetsDnsMetricsRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/role-binding.yaml", size: 293, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc, 0x7d, 0xc7, 0x45, 0x33, 0xc4, 0xd8, 0xf, 0x8d, 0x89, 0x8d, 0x6, 0x47, 0xa7, 0xa, 0x6b, 0x17, 0xf5, 0x5f, 0x5a, 0x2f, 0xd8, 0xf9, 0x6, 0x71, 0xaa, 0x78, 0x8d, 0xb5, 0x7a, 0xf6, 0x99}}
	return a, nil
}

var _assetsDnsMetricsRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xb1\x4e\xec\x40\x0c\x45\xfb\xf9\x0a\x6b\x5f\x9d\x7d\xa2\x5b\x4d\x8d\x44\x47\x01\x12\xbd\x77\xe6\x42\xac\x24\xe3\x91\xed\x04\xc1\xd7\xa3\xcd\x06\x89\xca\xf7\x1e\x59\x3e\xfe\x47\x2f\x3a\xc3\xa9\x01\x15\x95\xae\x5f\xd4\x4d\x17\xc4\x88\xd5\x29\x94\xbc\x18\x77\xd0\xe3\xf3\x2b\x2d\x08\x93\xe2\x84\x56\xbb\x4a\x8b\xc4\x5d\xde\x60\x2e\xda\x32\xd9\x95\xcb\x99\xd7\x18\xd5\xe4\x9b\x43\xb4\x9d\xa7\x8b\x9f\x45\xff\x6f\x0f\x69\x92\x56\xf3\x2e\x4a\x0b\x82\x2b\x07\xe7\x44\xd4\x78\x41\xfe\xe3\x1b\xa6\x8b\x1f\xd8\x3b\x17\x64\xd2\x8e\xe6\xa3\xbc\xc7\x50\x9b\x27\x5b\x67\x78\x4e\x03\x71\x97\x27\xd3\xb5\xfb\xed\xca\x40\xa7\x53\x22\x32\xb8\xae\x56\x70\x30\x87\x6d\x52\xe0\x7b\xf9\xfd\xf8\xde\xba\xd6\x5b\xd8\x60\xd7\x63\xf9\x03\xb1\xcf\x59\xfc\x1e\x3e\x39\xca\x98\x7e\x02\x00\x00\xff\xff\x29\x39\xda\x05\x1c\x01\x00\x00")

func assetsDnsMetricsRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsRoleYaml,
		"assets/dns/metrics/role.yaml",
	)
}

func assetsDnsMetricsRoleYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/role.yaml", size: 284, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8c, 0xf2, 0x4e, 0x40, 0x91, 0xd8, 0x5e, 0x1c, 0x98, 0xb6, 0x2f, 0x11, 0x2a, 0x15, 0x8f, 0xe4, 0x7c, 0xfe, 0xc6, 0x31, 0xf3, 0xb2, 0xa0, 0x38, 0xb2, 0x3f, 0x15, 0x5a, 0x33, 0x12, 0xd2, 0x88}}
	return a, nil
}

var _assetsDnsNamespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xcd\x4e\xc4\x30\x0c\x84\xef\x79\x8a\x51\x38\x2f\x3f\xd7\xbc\x03\x5c\x90\xb8\xbb\x8d\x97\x35\x4d\xed\x2a\x76\xcb\xeb\xa3\xb2\x15\xac\xb4\xc7\x68\x46\xf3\x7d\xf1\x24\x5a\x0b\xde\x68\x66\x5f\x68\xe4\x44\x8b\x7c\x70\x77\x31\x2d\xd8\x5e\xd2\xcc\x41\x95\x82\x4a\x02\x48\xd5\x82\x42\x4c\x7d\x7f\x02\xb6\xb0\xfa\x45\xce\xf1\x28\xf6\xa4\x56\xf9\xe4\xdc\x78\x0c\xeb\x05\x39\x27\x40\x69\xe6\xf2\x5f\x3b\x55\xf5\x04\x34\x1a\xb8\x1d\x13\x0f\x70\x0e\x6c\xd4\x56\x46\x18\x68\x33\xa9\xa8\xbc\xb0\x56\xd1\x4f\x98\x62\x5a\x07\x06\xd5\x59\x7c\x97\x42\x5c\x28\x8e\x82\xef\xf1\xdf\x38\x68\x11\xbf\xd7\xea\xab\x9e\x1a\x6f\xdc\x0a\xf2\x73\x3e\x98\xd4\x9a\x7d\xdf\x78\xcd\xa6\x12\xd6\x77\x62\x18\x9a\xd9\x84\xb3\x75\xbc\x73\xdf\x64\xe4\xd7\x6b\x0a\x1b\xbe\x78\x0c\x87\xec\x16\xe2\xbf\xbf\xbb\x1e\xed\x8e\x3a\xb6\xd5\x83\xfb\xcd\x70\x41\x8e\xbe\x72\x4e\x3f\x01\x00\x00\xff\xff\x82\x6d\x29\x03\x71\x01\x00\x00")

func assetsDnsNamespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsNamespaceYaml,
		"assets/dns/namespace.yaml",
	)
}

func assetsDnsNamespaceYaml() (*asset, error) {
	bytes, err := assetsDnsNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/namespace.yaml", size: 369, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe, 0xab, 0x50, 0x84, 0x61, 0x5f, 0x41, 0xf4, 0x17, 0x3b, 0x6, 0x84, 0xc0, 0x5f, 0x4f, 0xbb, 0xd8, 0x1d, 0xae, 0x26, 0x3e, 0x1f, 0x29, 0x2c, 0x84, 0x6d, 0x5e, 0xc1, 0x87, 0x97, 0x5f, 0xc9}}
	return a, nil
}

var _assetsDnsServiceAccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xc9\xb1\x09\xc4\x30\x0c\x05\xd0\xde\x53\x68\x81\x2b\xae\x55\x77\x33\x1c\xa4\x17\xf2\x0f\x11\xc1\xb2\xb1\x14\xcf\x1f\x02\xe9\x1e\xbc\xd3\xbc\x32\xfd\x31\x97\x29\x7e\xaa\xfd\xf2\x2c\x32\x6c\xc3\x0c\xeb\xce\xb4\xbe\xa5\x21\xa5\x4a\x0a\x17\x22\x97\x06\xa6\xea\xf1\x3a\x86\x28\x98\xfa\x80\xc7\x61\x7b\x7e\x9e\xba\x03\x00\x00\xff\xff\x8e\x2c\xf1\x2e\x55\x00\x00\x00")

func assetsDnsServiceAccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsServiceAccountYaml,
		"assets/dns/service-account.yaml",
	)
}

func assetsDnsServiceAccountYaml() (*asset, error) {
	bytes, err := assetsDnsServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/service-account.yaml", size: 85, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x57, 0x12, 0x50, 0x4d, 0x67, 0x2f, 0x1b, 0x74, 0xa0, 0xa4, 0xbb, 0xa7, 0x59, 0xe9, 0x5a, 0xc6, 0xc1, 0x1a, 0xf8, 0x5f, 0xff, 0x5, 0xdb, 0xc, 0x10, 0x8b, 0xc1, 0x0, 0xcc, 0xf, 0x9f, 0x3a}}
	return a, nil
}

var _assetsDnsServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xcf\x4f\x4b\x03\x31\x10\x05\xf0\x7b\x3e\xc5\xa3\xbd\x89\xad\x14\xf1\x60\xae\x7a\x11\x2f\x05\xff\xdc\xa7\xd9\x61\x0d\xcd\x66\x42\x66\xb6\xb2\xdf\x5e\x36\x62\x6d\x8b\xe0\x25\x90\xcc\xe3\x97\x37\xfb\x98\x3b\x8f\x17\xae\x87\x18\xd8\x51\x89\xef\x5c\x35\x4a\xf6\x38\x6c\xdc\x12\x99\x06\xbe\x6e\xa7\x16\x0a\x0c\xca\x1d\x12\xed\x38\x29\xa8\x32\x94\x0d\x64\xa8\x63\xb6\x38\xb0\xd3\xc2\xc1\x3b\x60\x89\x90\x46\x35\xae\x4f\x5b\x7c\xc6\x94\xb0\x63\xd0\x68\x32\x90\xc5\x40\x29\x4d\x18\x28\x53\xcf\xdd\xba\x85\x95\x13\x07\x93\x8a\xa8\x97\x22\x50\xa4\x9a\xce\xe8\xaa\xd5\xf0\xe8\xb2\x3a\xe0\x7b\xe0\x71\x77\xdb\x2e\x46\xb5\x67\xdb\xb6\xa7\x63\xa0\x8a\x49\x90\xe4\xf1\xf6\xb8\x3d\x07\x56\x16\xca\xbf\xc8\x6f\xe8\x08\xbd\x3e\x9c\x42\x03\x5b\x8d\xe1\xb4\xcd\xfd\xe6\x0f\xea\x2c\x76\x41\x99\x14\x49\xd2\x4f\xcf\x3c\xb5\x25\x67\x7c\xb1\x1f\x77\x5c\x33\x1b\xeb\x3a\xca\xcd\x87\xa8\xcd\xff\x2d\x7e\xc6\x57\x0b\xf7\x15\x00\x00\xff\xff\xf1\x41\x3e\x1b\xb6\x01\x00\x00")

func assetsDnsServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsServiceYaml,
		"assets/dns/service.yaml",
	)
}

func assetsDnsServiceYaml() (*asset, error) {
	bytes, err := assetsDnsServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/service.yaml", size: 438, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf3, 0x70, 0xb6, 0xc7, 0x42, 0xd7, 0x3b, 0x7f, 0x6e, 0x13, 0xfd, 0x73, 0xc4, 0x70, 0x5f, 0x90, 0xa9, 0xef, 0x6c, 0xc8, 0x51, 0xa1, 0x6, 0x39, 0xc6, 0x1f, 0xba, 0x95, 0xd7, 0x32, 0xd5, 0x59}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"assets/dns/cluster-role-binding.yaml": assetsDnsClusterRoleBindingYaml,

	"assets/dns/cluster-role.yaml": assetsDnsClusterRoleYaml,

	"assets/dns/configmap.yaml": assetsDnsConfigmapYaml,

	"assets/dns/daemonset.yaml": assetsDnsDaemonsetYaml,

	"assets/dns/metrics/cluster-role-binding.yaml": assetsDnsMetricsClusterRoleBindingYaml,

	"assets/dns/metrics/cluster-role.yaml": assetsDnsMetricsClusterRoleYaml,

	"assets/dns/metrics/role-binding.yaml": assetsDnsMetricsRoleBindingYaml,

	"assets/dns/metrics/role.yaml": assetsDnsMetricsRoleYaml,

	"assets/dns/namespace.yaml": assetsDnsNamespaceYaml,

	"assets/dns/service-account.yaml": assetsDnsServiceAccountYaml,

	"assets/dns/service.yaml": assetsDnsServiceYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"assets": {nil, map[string]*bintree{
		"dns": {nil, map[string]*bintree{
			"cluster-role-binding.yaml": {assetsDnsClusterRoleBindingYaml, map[string]*bintree{}},
			"cluster-role.yaml":         {assetsDnsClusterRoleYaml, map[string]*bintree{}},
			"configmap.yaml":            {assetsDnsConfigmapYaml, map[string]*bintree{}},
			"daemonset.yaml":            {assetsDnsDaemonsetYaml, map[string]*bintree{}},
			"metrics": {nil, map[string]*bintree{
				"cluster-role-binding.yaml": {assetsDnsMetricsClusterRoleBindingYaml, map[string]*bintree{}},
				"cluster-role.yaml":         {assetsDnsMetricsClusterRoleYaml, map[string]*bintree{}},
				"role-binding.yaml":         {assetsDnsMetricsRoleBindingYaml, map[string]*bintree{}},
				"role.yaml":                 {assetsDnsMetricsRoleYaml, map[string]*bintree{}},
			}},
			"namespace.yaml":       {assetsDnsNamespaceYaml, map[string]*bintree{}},
			"service-account.yaml": {assetsDnsServiceAccountYaml, map[string]*bintree{}},
			"service.yaml":         {assetsDnsServiceYaml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
