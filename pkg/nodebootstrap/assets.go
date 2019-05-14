// Code generated by go-bindata.
// sources:
// assets/10-eksclt.al2.conf
// assets/bootstrap.al2.sh
// assets/bootstrap.ubuntu.sh
// assets/cut_script.sh
// assets/kubelet.yaml
// assets/max_pods_map.txt
// DO NOT EDIT!

package nodebootstrap

import (
	"bytes"
	"compress/gzip"
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
	bytes []byte
	info  os.FileInfo
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

var __10EkscltAl2Conf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\x4f\x6b\xdb\x4e\x10\xbd\xeb\x53\x2c\x24\x87\xdf\x0f\xbc\x52\xe2\xb8\x39\x04\x74\x50\x63\x25\x18\x54\x27\x44\x0e\x2d\xb4\xc5\x8c\x77\xc7\xce\xe0\xd5\xac\xd8\x5d\xd9\x49\x83\xbf\x7b\x91\x65\xb5\x2e\x09\xa5\x37\xed\xbc\x99\xf7\xde\xfc\xd1\x89\xc0\xb5\x57\xc1\x48\x5f\xa3\xa2\x25\x29\xe1\x5f\x7c\xc0\x4a\x0b\xed\x6c\x2d\x89\x45\xc3\x14\xc4\xd2\x3a\xb1\x6e\x16\x68\x30\x0c\xf6\x8f\xac\x82\x1f\x96\x45\x41\xdc\x3c\x8b\xa1\xf8\x2f\x2b\x86\xff\x47\xd1\xd7\x12\xdd\x86\x14\x7e\x8f\x4e\x44\x61\x15\x18\x51\x61\x00\x0d\x01\x44\x0d\x0e\x2a\x0c\xe8\xfc\x95\x78\xc8\x6f\x27\x77\xd3\x81\xc8\x3e\x97\xf3\x71\x7e\x93\x3d\x16\xb3\x79\x17\x8b\x72\xde\x90\xb3\x5c\x21\x87\x1b\x32\x98\x26\x18\x54\xd2\x59\x4c\x7a\xae\x18\x79\x13\x9d\x88\x5b\x63\x17\x60\x04\xb0\x16\x3e\x40\x20\xf5\x87\xc6\x75\xf1\x58\xce\xf2\x87\xf9\x78\x5a\x0e\xc4\xf4\x6e\x9c\xcf\x8b\xec\x63\x5e\xf4\x8f\x59\x36\x99\xce\xca\xbf\xca\x1d\xfa\x3d\xa8\x75\xed\xb0\x65\xf9\x8e\xd8\x9e\x72\x72\x3f\x10\x93\x69\x39\xcb\xa6\xd7\xf9\x7c\x32\xfe\x27\x6e\xd3\xb2\xee\x15\xa2\xfc\x19\x55\x19\xc0\x85\xf4\xe8\x33\x69\xbc\x4b\x16\xc4\x7d\x81\xf8\x16\x09\x21\x25\x5b\x8d\x92\xea\xf4\xf4\xf5\xa0\xbc\x3b\x06\x0c\x2c\xd0\xf8\x1e\xec\xda\xde\x0d\xc0\xd4\x4f\x10\x77\xfa\x31\xd9\x84\xd8\x07\x60\x85\x92\x74\x7a\xfa\x7a\x64\xbc\xe7\xaa\xe0\x59\xd6\x56\xb7\x44\x9f\xb2\x2f\xf3\xfb\xbb\x71\xd9\x43\x0e\x57\xe4\x03\xba\xbd\x5e\x1a\x5c\x83\xc7\xc1\x2d\x85\x27\x19\x80\x38\xfc\x32\xd1\x8d\xbb\x2f\x07\x63\xec\x56\xd6\x8e\x36\x64\x70\x85\xba\x63\xe8\x30\x65\x6c\xa3\x65\xed\xec\x86\x34\xba\x14\xb6\xbe\x07\x2c\xb7\x9c\xe8\xa4\x6b\x38\x50\x85\xa9\xb6\x6a\x8d\xae\xef\x1c\xc3\xd6\xba\xb5\xac\x4d\xb3\x22\x4e\x15\x53\x5f\xc7\x24\x17\xc4\x52\x93\x4b\x13\x5b\x87\x44\x31\xb5\x23\x3d\x82\x95\xe5\x65\x87\xb7\x2b\x6a\x71\xc6\x10\xeb\x43\x46\x6d\xb5\x24\x5e\x3a\x38\xb2\x40\x15\xac\x30\xbd\x3c\x1b\x8e\xce\xce\xcf\x47\x17\xa3\x0f\xc3\x58\xaf\x5d\x8c\xca\xc5\xa7\xaf\x6f\xcf\x7a\x17\xc3\xfe\x7f\x81\xad\x8f\x95\xad\xda\x2b\x48\x6a\x68\x3c\x4a\xa8\xf4\xe5\xe8\xea\x22\x3e\x3f\x88\xb5\x7b\x6e\xed\xd0\xea\xcd\xbd\x74\xe1\xf8\x05\x2a\xf3\x7b\x24\xef\x25\xb6\x87\xd5\x66\x45\x3f\x03\x00\x00\xff\xff\x5c\x6d\xc6\x40\xdd\x03\x00\x00")

func _10EkscltAl2ConfBytes() ([]byte, error) {
	return bindataRead(
		__10EkscltAl2Conf,
		"10-eksclt.al2.conf",
	)
}

func _10EkscltAl2Conf() (*asset, error) {
	bytes, err := _10EkscltAl2ConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "10-eksclt.al2.conf", size: 989, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapAl2Sh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xc1\x6e\xd3\x40\x10\x86\xef\xf3\x14\x3f\x26\x87\x44\x68\xb3\x80\x0a\x52\x1b\x05\xa9\x22\x41\x8a\x04\x69\xa4\xf6\x00\x8a\x82\xb5\xb1\x27\x8d\xd5\xcd\xae\xe5\x1d\x97\x56\xc1\x3c\x3b\x72\xe2\x04\x57\xf4\x50\xe8\xc9\xf6\xef\x99\x6f\xb5\xdf\xcc\xcb\x17\x7a\x99\x39\xbd\x34\x61\x0d\xc5\x44\xab\xd2\x25\x92\x79\x87\x6b\x96\x78\x63\xee\xe2\xdc\xa7\xa1\xdb\xc3\x96\x80\x2f\xe7\x5f\xe3\xd9\xc5\xe8\x32\xfe\x34\xf9\x3c\x1e\x46\x9a\x25\xd1\x7c\x13\x12\xb1\xfa\x50\x19\x6f\x4c\xde\x97\x3b\x89\x08\xb8\x2e\x38\x47\xd4\xd9\x4e\xa6\x97\x57\xe7\xd3\x8f\xe3\xf8\xea\xdb\x6c\x5c\x45\x75\xf4\x80\x54\x45\xf8\x89\x1f\xeb\xcc\x32\x0a\x36\x29\x32\x17\xc4\xb8\x84\x63\xb9\xcf\x19\x35\x75\x80\xd4\x13\x01\x40\xb6\x02\xe6\xf3\x1a\x51\xe7\x55\x84\xe1\x2f\x7c\x9f\xbf\x56\xa7\x8b\x57\x1d\x2c\x16\x18\x40\xd6\xec\x76\xa5\x00\x27\x6b\x8f\xa6\x72\xd0\x64\x05\x4b\x59\xec\x0b\x56\x19\x06\x7b\x6a\xea\x1d\x53\x45\x34\xbd\x18\x8d\xe3\xc9\x6c\xd8\xe9\x26\x65\x61\xa1\x54\xc8\x2c\x3b\xc1\x5a\x24\x3f\xd3\xfa\xcd\xfb\xd3\xfe\xdb\x77\x27\xfd\xe6\xa9\xad\x11\x0e\xa2\x37\x2c\x46\xa5\x46\x8c\xb6\x3e\x31\x56\x65\xf9\xed\x49\x8f\x8e\xb7\x9e\x8c\xfe\x9b\x77\x30\xa1\xb2\xb4\x05\xac\x35\x3e\x1f\x59\xcb\xed\x11\x05\x5f\x16\x09\xa3\x3d\xcb\x9b\x72\xc9\x96\xa5\xcf\xee\x96\xe8\x30\xa9\xe1\x9f\xa1\x9d\xa9\x4e\xb7\xbd\x1d\xbd\x8a\x28\xb0\x40\x79\x38\x5f\xba\xc0\x42\xb4\x53\x1f\x1d\x7d\x6e\x9b\xb7\x2a\xc2\x87\x47\xcf\xda\x99\xdb\x9d\xb8\xef\x7c\x60\x6f\xdb\xfa\xaa\x09\xff\x88\xd8\xfb\xfa\x6b\x0d\x9f\x88\x79\x44\xc0\x53\x9a\x29\xdc\x07\xe1\x4d\x22\x16\xa9\xe1\x8d\x77\xaa\x60\xeb\x4d\xda\xca\xd9\x99\xa5\x65\x34\xbd\xad\x1f\x41\x4c\x21\xc7\xfc\x77\x00\x00\x00\xff\xff\x5e\xca\x96\x59\xa2\x03\x00\x00")

func bootstrapAl2ShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapAl2Sh,
		"bootstrap.al2.sh",
	)
}

func bootstrapAl2Sh() (*asset, error) {
	bytes, err := bootstrapAl2ShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap.al2.sh", size: 930, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapUbuntuSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x96\xe1\x6e\xdb\x36\x10\xc7\xbf\xeb\x29\x6e\xaa\x81\x26\x58\x29\x25\x69\x5a\xa0\x0d\x34\xcc\xab\xdd\xc1\x58\xe6\x14\xb5\x8b\x6e\x48\x33\x97\x26\x4f\x16\x61\x8a\x14\xc8\x93\x9d\x34\xf3\x9e\x7d\xa0\x64\x39\x4e\x80\x6d\xcd\x86\x21\x1f\x62\xdd\xfd\xf8\xbf\xd3\xe9\xee\xa4\x27\xdf\xa4\x73\x65\xd2\x39\xf7\x05\x30\xac\xa3\x28\xaf\x8d\x20\x65\x0d\x2c\x90\x66\x25\xbf\x9e\x55\x56\xfa\x83\x43\xb8\x8d\x00\x7e\xee\xff\x32\x7b\x77\x31\x98\xcc\xde\x8e\xce\x87\x59\x9c\x22\x89\x14\x97\x3e\x45\xa3\x58\xc9\xaf\x59\x40\x13\xba\xa6\x38\x02\x58\x38\xac\x20\xee\xdd\x8e\xc6\x93\x69\x7f\xfc\x66\x38\x9b\xfe\xfa\x6e\xb8\x89\x83\xe9\x9e\xcc\x26\x86\xdf\x61\x5d\x28\x8d\xe0\x90\x4b\x50\xc6\x13\x37\x02\x67\x74\x53\x21\x04\xc5\x33\x90\x36\x8a\x00\x00\x54\x0e\x70\x79\x19\x24\x82\x7d\x13\x43\xf6\x07\xfc\x76\x79\xc4\x5e\x5d\x7d\xdb\x83\xab\x2b\x38\x03\x2a\xd0\x34\x28\x00\x8a\xc2\xc2\x96\x3c\xdb\xda\x1c\x52\xed\x5a\x20\x57\x70\xd6\xaa\x4a\x6b\x30\xda\x44\x51\xd4\x9c\x88\xc7\x17\x83\xe1\x6c\xf4\x2e\xeb\x1d\x88\xda\x69\x60\xcc\x2b\x8d\x86\xa0\x20\xaa\x5e\xa7\xe9\xf1\xcb\x57\xc9\xc9\x8b\xd3\x64\xfb\x3f\xd5\x9c\xd0\x53\x5a\x22\x71\x26\x39\xf1\x54\x5b\xc1\x35\x53\xd5\xea\xf4\x30\x86\xef\xa0\x2b\x92\x20\x9d\x2e\xeb\x39\x6a\xa4\xa4\x41\x12\x34\xab\x6d\xc8\x5d\x8d\x46\x83\x7f\x1d\xb6\xab\x1b\x53\x32\xc4\x7d\x64\xe0\xf0\x70\x76\xa1\xfd\xe3\x83\x86\x87\xf5\xf5\x61\xbb\x06\xc8\x7a\x07\x8f\x6a\x13\x51\x13\xb0\x1c\x4e\x80\x49\x78\x0a\x4f\xff\xff\x80\x1e\x25\x7c\x8a\x7d\xfa\x10\x87\x34\xfd\x14\x7f\x45\x78\x78\x02\xd3\x8b\xc1\x05\x4c\x87\x93\x69\x14\x79\xc3\x2b\xe0\x5a\x71\x0f\x5b\x92\xe1\xd2\x27\xdb\xdf\x9d\xed\x21\x26\x48\xef\x30\x41\xba\xb3\xb5\x98\x27\x5b\xed\x8b\x45\xfe\xc6\x13\x96\x81\x73\xe8\x91\x58\xce\x95\x46\x19\x45\xd1\x41\x04\xdb\x6c\x5e\x87\x29\xf1\x08\xbe\xb0\xb5\x96\x30\x47\xd0\xd6\x2e\x51\x02\x27\xc0\x15\xba\x1b\x20\x55\x62\xa7\x0a\x9e\xb8\x23\x0f\x75\xf5\xac\x51\x58\x17\x4a\x14\xa0\x3c\xac\x0b\x4e\xb0\x46\x90\x16\x94\x81\xfe\xf9\x09\x1c\xec\x7c\x73\x1e\x2a\x67\x0d\x54\x9a\x2b\x03\x6d\x52\xb2\x15\xe0\x46\x42\x89\xdc\x10\x90\x0d\xc1\x2b\xeb\x88\xcf\x35\x86\xcb\xd2\x7a\xea\x68\x90\xca\x93\xb3\xfe\xf0\x19\xcc\x6b\x02\x45\x4f\x7d\x73\xde\x58\x02\xa1\x91\x3b\x28\xec\x3a\x1c\xd2\x96\xcb\xed\x2d\xe5\xce\x96\x77\x89\x87\x02\xad\x15\x15\xb6\x26\x28\xf8\x4a\x99\x45\x23\x40\x16\x44\xed\xc9\x96\xca\x63\x38\xd7\x82\x8a\x3c\xea\x3c\x02\xf0\xb6\x76\x02\xff\xa1\xad\xfe\x16\xfb\x4b\x20\x0c\x4e\x98\x9b\x86\x78\xd4\x3e\x7d\x72\x07\xff\xd7\xa1\xb9\xa7\xb5\x9f\x45\xf6\x79\x7f\xe5\x7f\x8e\x82\x2f\xd7\x7c\xe1\xb3\x83\x66\x4f\xc6\x5c\x4a\x87\xde\x67\x47\x49\xf3\x17\xb7\x56\x63\x25\x32\x55\x65\xbd\xdb\xed\xe6\xdc\x6c\x1d\x42\xd7\x9e\xd0\x31\x69\x7c\xd6\xbb\x7d\x73\xfe\x61\x32\x1d\xbe\x9f\x0d\xc6\x93\x0e\xe8\x6e\x31\xbb\x4b\x7b\xb3\x2f\xaa\xf9\x1c\xb5\xef\x84\xcf\xfb\x3f\x0c\xcf\x27\x9b\x67\x5c\x57\x05\x4f\xda\x8a\x26\xca\xee\x2f\xbe\x6c\xaf\x24\xa3\x41\xa7\xc5\xeb\xf0\x56\x20\x25\x78\x78\xaf\x31\xb2\x4b\x34\x6c\x8d\xf3\xc2\xda\x65\x46\xae\xc6\x3d\xce\x3a\xf5\xa5\xc5\x4a\x2b\x31\xfb\xd8\x52\x1d\xa0\xb5\x5d\xb3\xca\xa9\x95\xd2\xb8\x40\xb9\x7f\xb8\xb2\x92\x29\x93\x3b\xce\x84\x35\xc4\x95\x41\xc7\x54\xc9\x17\x98\xbd\x3c\x3a\x39\x3d\x3a\x3e\x3e\x7d\x7e\xfa\xe2\x24\x91\x4b\x97\xa0\x70\x49\xef\xb6\xff\x71\x32\x1b\x0c\xdf\xf6\x3f\x9c\x4f\x67\xef\x87\x3f\x8e\x2e\xc6\x9b\x84\x97\xfc\x8b\x35\x7c\xed\x13\x61\xcb\xa6\x13\x2a\x5e\x7b\x64\xbc\x94\x2f\x4f\x5f\x3f\x4f\x8e\x77\x95\xb5\xb5\x64\x95\xb3\x2b\x25\xd1\x65\x7c\xed\x1f\x96\xdc\x96\x5c\x99\x6c\x7b\xd9\xf6\x6d\x87\x18\xc5\xe6\xca\x30\xa9\x5c\x96\xda\x8a\x52\x61\x54\xf8\x04\xd8\x73\x0b\x6b\xf2\xd6\x1f\x1a\x32\xf8\x0d\x52\x22\x3b\x62\x77\x7f\xae\x36\x61\x53\x64\xd2\x8a\x25\xba\xee\xc9\x21\xad\xad\x5b\xb2\x4a\xd7\x8b\x90\x82\x51\xdd\xb9\x85\xb3\x75\xc5\xa4\x53\x2b\x74\x59\x7b\x95\x77\x89\x3b\x5c\xa8\x26\xf3\xf0\xe0\xf7\xeb\xba\x73\x84\x41\x66\x21\x30\xed\x3a\x62\xda\x1f\x8d\xa7\xbb\x96\x69\xf6\xa2\x35\xb9\x5a\x64\x0f\x87\xb2\x35\x27\x37\xbc\xec\xaa\x90\x23\xa7\xda\x21\x5b\x84\x37\x5a\xf6\xde\x12\x27\xfc\xa9\x1d\xdf\x09\xba\x15\xba\x37\xe8\x48\xe5\xa1\x67\xee\xa5\xc3\x8d\x35\x37\xa5\xad\x3d\x0b\xdd\x92\xe5\x5c\x7b\xdc\xd5\x5e\xa1\x21\x26\x38\xcb\x95\xc6\x7b\x39\x08\x9e\x08\xd7\x7c\x15\x1d\x86\xa1\x6a\x37\xf7\xdd\xc6\x0f\x8b\x3b\xcc\x6e\x33\x6c\x97\xdf\x5f\x6d\xe2\xe8\x30\xea\xf6\x3b\x77\xf7\xb8\xe8\xcf\x00\x00\x00\xff\xff\xfe\x11\x75\x1e\xb1\x09\x00\x00")

func bootstrapUbuntuShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapUbuntuSh,
		"bootstrap.ubuntu.sh",
	)
}

func bootstrapUbuntuSh() (*asset, error) {
	bytes, err := bootstrapUbuntuShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap.ubuntu.sh", size: 2481, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cut_scriptSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8e\x4d\x4b\xf4\x30\x14\x85\xf7\xe7\x57\x9c\xb7\x6f\x61\x14\x29\x11\x77\x52\xba\x10\xdc\x8a\x82\x1b\xa1\xd4\x4e\xa6\xb9\x9d\x84\xda\xa4\xa4\xa9\x1f\xf8\xf1\xdb\xa5\x23\x03\xee\x5c\xde\xfb\xf0\x3c\x9c\xff\xff\xd4\x32\x47\xb5\x73\x5e\x89\x7f\xe6\x4e\xcf\x16\xa3\x7e\x6d\xa7\x60\xe6\x2a\xcb\x4a\x00\xb8\xb9\x7a\x68\xef\x6e\xaf\xef\xab\xed\x5e\x52\x7b\xa4\x5b\x00\xd2\xd9\xc0\xfc\xc8\x81\x7d\x94\x89\xc9\x46\x11\x2a\x1b\x46\x51\xa3\x8e\xc9\x79\xad\x86\x61\x18\xf8\xc1\x6e\x49\x2c\x7a\x5e\xb0\x30\xdc\x70\x83\x7e\xf1\x5d\x72\xc1\xf3\x77\xf8\xe4\x94\xef\x7f\x96\x5e\xac\x7b\x12\x46\xd1\x86\xe9\x6d\x12\xae\x62\x49\x13\x00\xd2\xf5\x64\x5d\x33\xcb\xd7\x67\xc6\xea\x8b\x8f\xf5\x79\x71\xd9\x9c\xe5\x6c\x1a\x96\x4c\x56\x3c\x48\xf2\x67\xfe\x41\x3d\xdc\x51\xd2\x12\x57\xd4\x3b\x96\x30\xc1\x0b\x3e\x01\x7c\x07\x00\x00\xff\xff\x81\x91\x6a\x6b\x24\x01\x00\x00")

func cut_scriptShBytes() ([]byte, error) {
	return bindataRead(
		_cut_scriptSh,
		"cut_script.sh",
	)
}

func cut_scriptSh() (*asset, error) {
	bytes, err := cut_scriptShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cut_script.sh", size: 292, mode: os.FileMode(493), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _kubeletYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\x4f\x4f\xc3\x30\x0c\xc5\xef\xf9\x14\xfe\x04\x6d\x07\x9a\x04\xb9\x8d\x4d\x70\x60\x27\x36\xe0\xec\xa6\xee\x16\x35\x8d\x27\xc7\x19\x7f\x3e\x3d\x5a\x5a\x90\x26\xa1\x9c\x9e\xde\xb3\xdf\x4f\xce\xe0\x63\x67\xe1\x39\xb7\x14\x48\xd7\x1c\x7b\x7f\xc8\x82\xea\x39\x1a\x3c\xf9\x37\x92\xe4\x39\x5a\x18\xa6\x40\xe5\x4a\xa2\x1a\xee\x52\xe5\xb9\x3e\x2f\x5a\x52\x5c\x18\x83\x5d\x27\x94\x92\x85\xa6\x2a\xcf\xb8\x90\x93\x92\x6c\x78\x44\x1f\x2d\xcc\xb2\x0a\xec\x30\x18\x83\x59\x8f\x14\xd5\xbb\x52\x64\x0d\x00\x46\x8e\x5f\x23\xe7\x74\x11\x00\x14\xb1\x0d\xd4\x59\xe8\x31\x24\x32\x00\x1f\xd4\x1e\x99\x87\xc9\x75\xe8\x8e\xb4\xdf\x6f\x2d\xdc\x8c\x4d\xba\x1e\x50\xc9\x97\xfc\xe7\xb2\xb9\x9f\xc3\xc1\x53\xd4\xf5\xea\xd1\x07\xb2\x50\x93\xba\x9a\x86\xe4\x34\xd4\x0e\x2b\x27\x3a\xd1\xb0\xf8\xef\x3f\x98\x91\x3b\xb2\xf0\x3e\x55\xfe\x5b\xbe\x9a\x47\xa8\x2b\x18\xcb\x5f\x8c\x62\xbe\x46\xbc\xb6\x6f\x9b\x64\x4c\x22\x39\x93\xec\xb7\xbb\x07\x66\x4d\x2a\x78\x9a\x61\x8d\x3b\x08\xe7\xd3\x46\xfc\x99\xc4\xc2\xa4\xfa\x64\x4c\x4f\xa8\x59\xe8\x09\x95\xca\x59\x5e\x58\x51\x69\xfe\xaa\x5d\x59\xb7\x26\x51\xdf\x5f\xee\x48\xf3\xb6\x9f\x00\x00\x00\xff\xff\x1f\x2f\xa9\x0f\xd0\x01\x00\x00")

func kubeletYamlBytes() ([]byte, error) {
	return bindataRead(
		_kubeletYaml,
		"kubelet.yaml",
	)
}

func kubeletYaml() (*asset, error) {
	bytes, err := kubeletYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kubelet.yaml", size: 464, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _max_pods_mapTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x55\x51\xae\xe3\x20\x0c\xfc\xef\x29\x7a\x02\xf4\xb0\x61\x9b\x77\x1c\x44\xfa\xda\x48\xa5\x5b\xd1\xac\x54\xed\xe9\x9f\x48\xd3\x30\xe3\x7c\x5a\x80\x6d\xc6\x33\xe3\x1c\x47\xf7\xfd\xba\xa5\x7a\x39\x1f\x45\xc3\xe1\xa2\xce\xff\x59\xe3\x10\xe5\x90\xbc\x5b\xa3\x38\x1c\xda\xe5\x00\x97\x6b\x70\x6b\xf0\x7d\x78\xf9\xfe\xb0\x9d\x25\xdf\xcf\x26\x93\xb4\x06\x4e\x13\x31\xcd\xd9\x49\xaf\xd8\x42\x15\xb8\x5b\x02\x1e\x3f\xc4\x0d\x94\x28\xf5\x4c\xb3\x6c\x37\x43\x58\x5a\x87\x77\x4a\xef\x32\x7e\xf2\xc7\x63\x85\x12\x9d\xa7\xfa\x71\x74\xf2\x69\xfe\xa4\xa7\x96\x0a\xae\x67\xa5\x9f\xe5\xe8\xfc\x00\x97\x8b\x42\xa1\xd2\x5f\x7a\xbf\x54\xda\x7a\x2f\x31\x19\x84\x46\x6e\x23\x63\xa2\xeb\xd3\xd3\x77\x2e\x8c\xca\xa4\xae\x9c\xe7\x74\x7b\x77\x10\x13\xb6\xb0\xcb\x4b\xf0\x8e\x62\xe7\x84\xe1\x2c\xae\xe3\x3b\x09\x3e\xac\x91\x40\x69\x55\x08\x06\x4e\x54\x70\x6c\x95\x3a\x98\x75\x3d\xd1\x78\xf8\xef\x47\x67\x18\x66\x8a\xf4\x24\x66\x6a\xed\x29\xdc\xbd\x88\x19\x31\x4d\xb4\x65\x82\xe3\xab\x37\x51\x60\x6c\x07\x66\xe7\x1a\x7c\x2d\xbf\x32\x79\x7b\x83\x39\xc2\x14\x66\x71\xf7\x74\xff\x7b\x7c\x13\x18\x7f\xef\x5d\x99\x72\x5d\x4e\x82\x91\x90\x69\x7a\x6e\x43\x1e\xa7\x7f\xe5\xe8\x4f\xad\x49\xec\xea\x21\x2c\xcd\x42\x5a\xa8\x86\x6b\x13\xb3\xa7\xa0\x36\xaa\x5a\x06\x7c\xfa\x9b\xf5\xf3\x87\x9c\x39\xc1\x0f\x23\x56\xac\xf6\x0d\x4a\xd5\x0a\xac\xcd\x6e\xc3\xad\x05\x0a\xcf\xaf\x9e\x61\x29\x82\xec\xf6\xc8\x08\x93\x65\xe3\xc7\x3a\x1a\x21\x07\xe8\x78\x05\x69\x32\x20\x08\xa0\x42\x60\x64\x67\x75\xcf\x92\x6e\xb7\xa3\xf7\x4d\xa2\x5d\xcf\x01\x33\x04\x9e\x7d\x9f\x5b\x73\xbb\x81\xdc\x0d\x2d\x6b\xbc\xef\xf8\x84\xce\x5d\x58\xac\x0f\xdd\xb5\x86\x6a\x01\xe7\xce\x81\xe6\x35\x4b\xff\x43\x63\xd5\x87\x80\x6a\xac\x96\x57\x41\xf6\xdb\x3f\xe4\x4d\x7d\xb2\x11\xc3\xa8\x77\x81\xc5\x02\xc1\xf3\x1c\x6b\x92\x6d\x4a\x68\x04\x42\x70\x0e\xbc\xc0\xa8\x98\x42\x5f\x0d\xdf\x17\xad\x1a\x82\xa8\x92\x28\xed\x9c\xbf\x18\xa3\xcd\x96\x26\xa5\x45\x47\x96\xc4\xcc\x17\xe3\xf7\x89\x4d\xaa\x92\x35\x27\x96\x4d\x0e\x4c\x6c\x04\x7b\xdc\xed\x0a\x5c\xa0\xbc\x3f\xd3\xf6\x74\x59\x56\x68\x8a\xfa\x74\x18\xed\x64\x6a\x9a\x1d\xed\x9e\x0e\x44\x2f\x63\x29\xfd\x6c\x12\xeb\xa0\xc1\x78\x53\x17\x2d\x40\x9b\xcd\xca\xae\xec\x71\x4d\xd2\xc4\x38\x68\x6e\xc2\x51\xfe\x06\x00\x00\xff\xff\x74\xf1\x52\xfe\xf1\x08\x00\x00")

func max_pods_mapTxtBytes() ([]byte, error) {
	return bindataRead(
		_max_pods_mapTxt,
		"max_pods_map.txt",
	)
}

func max_pods_mapTxt() (*asset, error) {
	bytes, err := max_pods_mapTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "max_pods_map.txt", size: 2289, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
	"10-eksclt.al2.conf": _10EkscltAl2Conf,
	"bootstrap.al2.sh": bootstrapAl2Sh,
	"bootstrap.ubuntu.sh": bootstrapUbuntuSh,
	"cut_script.sh": cut_scriptSh,
	"kubelet.yaml": kubeletYaml,
	"max_pods_map.txt": max_pods_mapTxt,
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
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	"10-eksclt.al2.conf": &bintree{_10EkscltAl2Conf, map[string]*bintree{}},
	"bootstrap.al2.sh": &bintree{bootstrapAl2Sh, map[string]*bintree{}},
	"bootstrap.ubuntu.sh": &bintree{bootstrapUbuntuSh, map[string]*bintree{}},
	"cut_script.sh": &bintree{cut_scriptSh, map[string]*bintree{}},
	"kubelet.yaml": &bintree{kubeletYaml, map[string]*bintree{}},
	"max_pods_map.txt": &bintree{max_pods_mapTxt, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

