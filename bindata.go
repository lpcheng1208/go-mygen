// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/tpl/curd.tpl
// assets/tpl/e.tpl
// assets/tpl/entity.tpl
// assets/tpl/example.tpl
// assets/tpl/init.tpl
// assets/tpl/markdown.tpl
// assets/tpl/tables.tpl
package main

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsTplCurdTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x5f\x6b\xe3\xd8\x15\x7f\x96\x40\xdf\xe1\xd4\x2c\x41\x9a\x18\x25\xdd\x2d\x7d\x08\x78\x96\xc4\x76\x76\x4c\x13\x27\x6b\x2b\x2d\x25\x84\x8c\x2c\x5d\x27\x6a\xf4\xc7\xb9\xba\x4e\x1c\x84\x60\x16\x5a\x76\x06\x76\xb6\xfb\xd0\xc9\x2e\xdb\xa1\xed\x2c\xbb\x65\x29\xcc\x66\x1e\x5a\xba\x0f\x85\x7e\x99\xd8\x93\x8f\x51\xee\x95\x64\xc9\xb6\x6c\x4b\xf9\x53\x86\x25\x03\xc3\x8c\xaf\xae\xce\xdf\xdf\xf9\x9d\x7b\x8f\xc8\x59\x07\x81\xe7\xc9\x4d\x82\xbb\x1a\x51\xd4\x96\x89\xea\xaa\x85\x7c\x7f\xd3\xd1\x91\x09\x2e\x5b\x06\x4f\xe0\xab\x02\x5f\x59\x83\x07\xee\xb1\x29\x57\xd6\x04\x5e\xe9\x05\xff\x57\x7a\x02\xef\x0b\xbc\xc0\x2f\x2d\x81\xed\x10\x20\x58\xb5\x5d\x55\x23\x86\x63\x0b\x7c\xbb\x6b\x6b\x50\x47\xa7\x69\xf2\x45\xbd\x05\xb2\x2c\x87\x02\x25\x78\x30\xdd\x08\x4f\xe0\x8d\x36\x98\xc8\x16\xf5\x96\x04\x0f\x61\x99\xae\x00\x00\x60\x44\xba\xd8\x86\x85\xa9\xaf\x86\xfb\xe8\x9f\xca\xda\x0a\xe8\xad\xdd\xe5\xbd\x62\xb0\xe6\x33\xbb\xb3\x4a\xa0\x6f\x5b\xaa\x4b\x10\xae\xac\x15\xd9\x9b\x3e\x73\x39\xe1\x2e\x38\xad\xdf\x21\x8d\xcc\xf6\x5a\xe9\x89\x64\x18\xb9\x79\x4e\x73\x59\xac\xe3\x38\xa5\xb7\x02\xa4\x57\x14\x78\xce\x0f\x73\xc1\x4c\xbb\xfa\xfc\xdf\xfd\x3f\x9e\x0f\x9e\x3d\x19\xbc\x7c\xf6\xf6\xeb\xdf\x5f\xbd\xfa\xbe\xff\xfa\xcb\xc1\x0f\xff\x0a\x0d\x14\xad\x19\xda\x25\x38\x40\xa4\xec\x98\x5d\xcb\x76\x45\x89\xa2\xc0\xb0\x0f\xa8\x49\xa1\x45\x05\x0a\x99\x55\xd3\x5c\x37\x90\xa9\x6f\x18\x2e\xf1\x7d\x28\xc4\x40\x08\x54\xf7\xbf\xfd\xfa\xea\xd5\x67\x83\x17\x6f\x06\xcf\x7f\x90\x33\x6b\x6d\x38\xa7\xae\xe8\x1e\x9b\x4a\x8f\x84\x7a\x8b\xd0\x51\xb1\x6a\xb9\x14\x2e\x86\x4d\x10\x6e\xab\x1a\xf2\x7c\x09\x44\xec\x9c\xba\x0d\xe4\x76\x4d\x02\xbb\x7b\x54\xee\xf6\xd1\x41\xd5\x26\x06\x39\xf3\xfd\x34\x2d\x45\x40\x18\xd3\xbf\x0e\x96\xa8\x37\xc7\x5d\x84\xcf\x82\xc5\x95\x12\x58\x72\x65\x4d\xfe\x98\x2e\x85\xfa\x23\xc5\xb2\x2c\x4b\x0c\x82\x74\xe3\xcf\x4a\x60\x1b\x2c\x3d\xf4\x17\x7d\xab\x2a\x37\x89\xaa\x1d\x89\x08\x63\x29\x8a\x10\x8b\x85\x8e\xda\x08\x03\x53\x22\x97\x4d\xc7\x45\xa2\x24\xf0\x6d\x27\x5a\xaa\xa3\x1e\x11\x99\x1d\xd8\x39\xa5\x06\x4c\x38\x50\xef\x9a\xe6\x84\x13\x9e\x1f\xa9\x0e\xc4\x34\x35\xd5\x16\x05\xde\xf3\xb0\x6a\x1f\x20\x60\x2f\xb1\xbc\xb8\x35\xbb\xed\xf8\xfe\x02\x76\x4e\x65\xcf\x93\x1f\x75\xad\x4e\x18\x86\xa5\x25\xaa\xab\xec\x58\x16\xb2\x89\xef\xd3\x97\x91\xad\xfb\x7e\x9a\x97\x14\xfd\xe9\x9e\xc6\x15\x18\x94\x52\x9c\x8c\x12\xa8\x9d\x0e\xb2\xf5\x44\x82\x8a\x0c\xc5\x73\xf2\xe3\xcd\x70\xc3\xf3\xa8\x69\xc7\x20\x7f\xe4\x28\x94\xb3\x0a\x6d\xd3\x51\xc9\x2f\x7f\x51\x60\xa2\x62\xe7\x56\x26\xdc\x95\xd7\x83\x9d\xa9\x6e\x9b\x2e\x82\x31\xc1\x86\x9d\x4d\x6c\xcd\xce\x21\x94\x18\x16\x92\x15\xc3\x42\x19\x04\xd3\x6d\x39\x8c\xfd\xe0\xfd\x6c\xc6\x7e\xf0\xfe\x34\xa1\xf3\x5f\x6f\x06\x95\x38\x0d\x38\xe1\x3f\x14\x40\xfe\x48\x09\x24\xe9\xe0\xf9\x8b\x21\x1d\xe4\x61\x03\x56\x8c\x28\x1b\x1b\x84\xf8\xbb\x2e\x15\x8c\xb2\x40\x42\xf7\x28\x13\xbc\xcb\xe5\x3a\x97\x94\xe2\x28\x95\xde\x81\x92\x84\xfb\x9a\xbc\xeb\x9a\x9c\xac\xc7\xfd\xc1\x9f\xff\x39\x38\x7f\x93\xa7\x14\x9b\xea\x09\x1a\xef\xca\x27\xaa\xd9\x45\x93\x65\xd8\x82\x96\xe3\x98\xe3\xf5\xe5\x12\x8b\x8c\x76\xda\x6d\x8c\x3a\x2a\x8e\xa4\xde\xa0\xc1\x52\xd1\x71\x7f\xc5\x61\xbb\x09\x55\xb1\x87\xd5\x1e\xd2\x44\x66\xee\x75\x5b\xf9\x89\x8a\x41\x6d\xb7\x91\x46\xca\x4e\xd7\x26\xc0\xf0\x28\xf0\x89\xa5\x62\xd8\x25\x03\xfd\x32\x3d\xc6\xac\xb2\xa7\x48\x17\xaf\xa5\xb2\x45\xdb\x68\x42\xe5\x43\x58\x1e\x4b\xe5\x4d\x92\xa9\xf4\x6e\x33\x9d\x4a\xef\x3e\x9d\x37\x4e\xe7\xd2\x12\x0c\xce\xdf\xf4\xbf\xf9\xcb\xe5\x7f\x5f\x0d\x3e\xb9\xc8\x96\xcb\x32\x46\x2a\x41\x41\x34\xb2\xf4\x3d\x09\x44\x53\x75\x49\x4d\x0f\x7c\x1e\xcf\xab\xe6\xd8\x2e\x81\xa8\xe3\x96\xa0\x50\xab\x37\xab\x0d\x05\x6a\x75\x65\x0b\x0a\xb0\x18\x36\x3d\x26\x90\x29\xd8\xe9\x74\x10\x4e\xc8\x87\x45\x28\x80\xe8\x79\x72\xcd\x76\x11\x26\x89\x0b\x82\x04\xbf\x5e\xdd\xd8\xa9\x36\x13\x4f\x37\x55\x7c\xe4\xfb\x52\x61\x1e\x3d\xa0\xff\x07\xa0\x12\x8d\x2e\xb0\x2e\x68\x72\x01\xce\xee\xb0\x21\x07\xe9\x18\xc3\xdb\x06\x5d\x0c\xac\xb8\x26\xde\x66\x41\x0b\x48\x2f\x0f\xba\x94\xde\x4f\x0d\x5f\x63\x7c\x75\x8f\xaf\x5b\xc0\x57\xee\x36\xb4\xd3\xd1\xf3\x52\x57\x7a\x33\x9a\x00\xd5\xce\x76\x65\x55\xa9\x66\xc7\x53\xb3\xaa\x00\x7b\x46\x0d\x4a\xce\x34\x7e\xf3\xa8\xda\xa8\x32\x19\xd8\xb0\x54\x7c\xf6\x2b\x74\xe6\xfb\x50\x82\x0f\x0b\x02\x1f\x5e\x43\x28\x9e\xd4\x23\x24\xee\xee\x25\xfa\x66\x11\x96\xa5\x38\xdf\xef\x19\x45\x78\xef\x44\x35\xe9\xde\x50\x09\x95\xcf\x14\xf9\x7e\x28\x67\x78\x67\x0e\x7e\x17\xc1\xf3\xe8\x3b\x2c\xe9\x61\xfa\x87\xe3\x17\x4b\x1e\x9e\xc7\x26\x2e\x27\xf1\xc9\x20\x99\x91\xcc\x15\x1f\x98\x97\xb3\xe2\xef\xd3\x92\x48\x4b\x78\xb2\x9a\x95\x98\xbf\x7e\x77\x75\xf1\x4d\x72\x3c\x96\x2d\x37\xeb\x86\xad\x8b\xf4\xa2\x8b\x5c\x6a\xe8\xf5\x66\x5e\x51\x3e\x56\x4a\x50\x68\x56\x37\xaa\x65\x85\xe6\xc3\x92\x47\x06\x7e\x8b\x50\x58\x6f\x6c\x6d\x66\x4a\x15\xa3\x3c\x6a\x4f\x71\x38\x27\x4a\x8e\xf1\x02\x6e\xbd\xad\x91\xc0\xba\x81\x5d\x22\x1a\x61\x77\x09\x62\x71\xcd\x1b\xff\xad\x07\x82\x61\x36\x05\x94\xb0\x51\xdb\xac\x29\xf0\xf3\xc2\x68\x73\x88\x03\x15\xe3\xc5\xd0\xa7\xc6\x6a\xf0\xf2\x49\xff\x8b\xcf\x2f\x7f\x7c\x92\x33\x62\x94\xe9\xc5\x77\x30\x50\x5b\x8d\x4a\xb5\x01\x6b\xbf\x85\x5a\x05\x2a\xd5\x66\x39\x73\x94\x52\x22\xd4\x7f\xfe\xa2\xff\xf4\xcb\x3c\x51\xd9\x36\xbb\xda\x51\x0a\x8e\x2c\xb5\xb3\xcb\x96\x46\x69\x63\x34\x20\xdc\x04\xb3\x05\x61\x81\xc7\x63\xd9\x7f\x5c\x64\x4b\x4d\xa4\x39\xb6\x1e\xf2\xca\x63\xc8\x17\xa7\xd3\x43\x84\xd1\x14\xae\xe3\xb0\x73\xea\x4e\x9b\x5f\xc7\x80\xe2\x26\xba\x3b\xc7\xa5\xf7\x77\x8e\x8b\x42\xcb\xf9\x02\xcf\x05\xe7\x1b\xaa\x24\x3e\xdf\x70\x38\x9a\x1e\x31\x7e\x4d\x0d\x18\xdd\x46\xaf\x58\xa2\xc0\x73\x00\x00\xfb\x51\xa0\xa3\xdf\x94\x75\x13\xfb\x05\x9e\xa3\xaf\xb4\x9d\x50\x59\x3c\x19\x8f\x0c\x65\xcb\x6c\x74\xb6\xb0\x6f\xe8\x45\x58\xa0\x22\x98\xc1\x29\xbe\x4d\x1f\x59\x73\x09\x07\x99\x87\xa1\x37\xbb\xfb\x86\xbe\x07\x25\x66\x58\xe8\xfb\x2c\x8c\x41\xeb\x0c\x06\x7f\xba\x18\x7c\xf6\x49\xff\xe9\xa7\x6f\x9f\x7d\x7a\xf9\x9f\xbf\xf5\xff\xf0\x1d\xdb\xf6\x11\x22\xa0\xb1\xaa\x00\x5d\x25\x6a\x0e\x34\xae\x9d\xad\x62\x2c\x1a\xba\x0b\xac\x67\xe5\x46\x65\x30\x8e\xcf\x92\x1b\xba\x33\xfc\xa0\x66\xe8\xae\x04\xa5\x52\xfc\x49\x2d\x39\xd4\x0f\x3f\x93\x71\x93\x04\x70\x87\x48\x2f\xc0\x22\xcd\x4b\x61\x1c\xf1\x86\x0d\x22\x15\xd5\x40\x1d\xa4\x92\x8f\xbb\xc8\x25\x86\x63\xd3\x13\xbe\x18\x39\xc2\x58\x88\x9e\xf6\xb9\xd9\x87\x80\xe1\xfe\x10\x73\x86\xde\xa3\x95\x42\xb7\x07\x27\x03\x9a\x04\x86\xa4\x40\xce\xae\xa1\xf7\x28\x3a\x0c\x3d\xc2\xc6\xbc\xb2\x4b\xf6\xfd\xdb\xad\xbe\xa8\xac\xb8\x64\x4d\x71\x77\x5c\x50\x53\xed\xbd\x76\x35\x0d\xfb\xff\xe5\x8f\xff\x88\x78\x3b\xac\x1e\xc7\x46\x39\x4a\x67\xcb\x46\x29\x34\x9e\x7a\xf7\xcc\x84\xe3\xdb\x60\xe7\x12\xe3\xe6\x28\x66\xa9\x9f\x13\x0c\x5d\x0a\xa3\x1f\x58\x3c\x89\x93\x85\x85\xe8\x97\x7b\x6c\xca\x55\x8c\xeb\x0e\x3d\x53\x31\x47\x66\x11\xdc\x08\x84\xd2\xc3\x1e\x9c\x20\x86\x01\x37\x0d\x1b\x81\xe6\x74\x6d\x92\x71\x28\x40\xb7\xd2\x23\x85\x16\xcf\xce\xd2\xc8\x28\x25\xdc\xe5\xad\x9d\xba\x22\x3e\x90\xb2\x47\x79\xf4\x03\x64\x4a\x2c\xa3\x38\x32\x63\x62\x72\x9b\x1f\xc9\x04\xe1\x4d\xff\xbe\x99\x42\x87\x63\x9f\x3d\x83\xe6\xf0\xf4\xdb\xc1\xf9\xeb\x00\xc8\x83\xaf\x2e\xfa\x5f\xfc\xbd\xff\xfa\xab\xfe\xcb\xef\xd9\xc3\xf2\x21\xd2\x8e\x80\x1c\x06\xb0\x06\xc3\x85\x43\xf5\x04\x7d\x98\x2d\xd8\x8f\x54\x37\x89\xef\xf4\x9b\x56\x1a\xb2\xf3\x87\x7a\xfe\x71\x83\x92\x8f\x96\x1c\x98\xce\xc9\x4d\x12\xe7\xb9\xf3\x03\x37\x4f\x50\x68\xec\x43\x58\x2e\x52\x4d\x02\xef\xff\x2f\x00\x00\xff\xff\x68\x1d\x8b\x73\xef\x22\x00\x00")

func assetsTplCurdTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplCurdTpl,
		"assets/tpl/curd.tpl",
	)
}

func assetsTplCurdTpl() (*asset, error) {
	bytes, err := assetsTplCurdTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/curd.tpl", size: 8943, mode: os.FileMode(511), modTime: time.Unix(1585914016, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTplETpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x92\xc1\x6b\xd4\x4e\x14\xc7\xcf\x3b\x30\xff\xc3\x63\x4f\x09\x84\xc9\xbd\x3f\x7e\x27\x5d\x6f\xf6\xa0\x82\xe7\x49\x7c\xc9\x86\x66\x33\xdb\x37\x13\xaa\x94\x3d\x28\x4b\x59\x41\xa1\x07\x95\x6e\x4f\x15\x4a\xc1\x83\x8d\xc7\x55\x2a\xfe\x33\x3b\xd9\xf6\xe4\xbf\x20\x93\x6c\xdc\x56\x45\x41\xf0\x14\xde\x37\xdf\x99\xf7\xf9\xc0\x64\xa3\xb1\x22\x03\x1e\x67\xbd\xfe\x23\x69\x64\x24\x35\x86\x7a\x37\xef\x73\xe6\xa2\x34\x33\xc3\x32\x12\xb1\x1a\x85\xe3\x9d\x34\x44\x22\x45\xba\xcf\x99\xef\x7e\x87\x21\xd8\x8b\x67\x76\xb1\xb0\xe7\xc7\xcb\x8f\xcf\x39\x33\x4f\xc6\x08\x03\xd0\x86\xca\xd8\xc0\x3e\x67\x93\x75\xed\xf2\xfc\x83\xfd\xfc\xda\x9e\x1c\xd4\x27\x33\xce\x92\xb2\x88\xc1\x43\x18\xf8\x70\xdf\xc8\x78\xc7\x43\x22\x68\xae\xf6\xdb\x8f\x3b\xda\xcb\x12\x88\x02\x40\x84\xad\xff\x01\x45\xa6\xb7\xcb\x3c\x77\x4d\xff\x3f\x88\x9a\x42\x8f\xd0\x94\x54\x00\x22\x67\xbd\x09\x60\xae\xf1\x66\xde\xc0\x8a\x87\x99\x19\xae\xd7\xa0\xef\x9a\x1b\xac\x16\xe8\x6a\x3e\xb5\x8b\xb3\xe5\x97\xb7\xf5\xd3\xea\x67\xb8\xbb\x3a\xdd\xf0\x05\x30\xd2\xa9\xf3\xcb\x8a\xf4\xef\x58\x89\x7e\x0f\x4b\x72\xec\x21\x36\x7b\xfe\x08\x0b\x89\xa2\x91\x34\xbf\x64\xbe\x73\x1d\xba\xed\xad\xb9\x03\x90\x94\x6a\x10\x42\x64\x85\x41\x4a\x64\x8c\xfb\x93\x7f\x26\x93\x34\x36\x2d\x40\xbb\x59\x08\xf1\xa3\xd9\xec\xb4\x7e\xf3\xbe\x3e\xaa\xec\xe1\x59\x7d\x54\xad\xde\x7d\x5a\x1d\x4f\xaf\x5e\xcd\x2f\xab\xea\xeb\xc5\x0b\x3b\x9b\x2f\x17\x2f\xed\xe9\x74\x75\x78\x70\x43\x75\x43\xd6\x3d\x1e\x2f\x52\x2a\x0f\xba\xc9\x11\xe9\xbd\xcc\xc4\x43\x97\x34\x63\x2c\x35\x82\xde\xcd\xc5\x80\x68\x5b\xdd\x53\x7b\x3a\xe8\xc6\x5b\xaa\x28\x6e\xab\x02\xbf\x07\x0f\x1e\xbb\x71\xeb\x9a\x96\xa1\x12\x83\xce\x9a\xb3\x2e\x4e\x64\xae\xbb\x7c\xc2\xd9\xb7\x00\x00\x00\xff\xff\x75\xdf\xbe\x7f\x52\x03\x00\x00")

func assetsTplETplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplETpl,
		"assets/tpl/e.tpl",
	)
}

func assetsTplETpl() (*asset, error) {
	bytes, err := assetsTplETplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/e.tpl", size: 850, mode: os.FileMode(511), modTime: time.Unix(1585652912, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTplEntityTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8e\x41\x0a\xc2\x40\x0c\x45\xf7\x85\xde\x21\x8b\x2e\xa5\x07\x10\x5c\x09\x5d\x76\x21\xbd\xc0\x68\x83\x54\x67\x3a\x25\x4d\x29\x25\xe4\xee\x12\x75\xa4\x54\xc1\x59\x4d\x5e\xfe\x0f\x2f\xcf\x44\xca\xc6\x9d\x3d\x1e\x63\x08\xd8\xb3\x6a\x9e\xf1\x32\x20\x24\xae\x0a\x23\xd3\x74\x61\x10\x0b\x93\xeb\xaf\x08\xc5\x6d\x07\x45\xc7\x18\x60\x7f\x80\xb2\xea\xd0\xb7\xa3\xaa\xc8\x93\x95\xb5\x0b\x56\x7b\xbd\x04\x9b\x65\x78\xc3\x44\xaa\x48\xc1\x71\x2a\xc3\x26\x7f\xc2\xe0\xe8\x6e\x3a\x22\xd8\xb7\xf6\xd1\x2f\xdb\x7a\xf2\x7e\x6d\x6c\xf3\x1f\x6b\x8a\xf3\x56\x9a\xe2\xbc\x72\x4e\xc0\x4e\x7d\x9c\x61\xb5\xf9\x65\xf6\x08\x00\x00\xff\xff\x6e\x2e\x32\xff\x48\x01\x00\x00")

func assetsTplEntityTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplEntityTpl,
		"assets/tpl/entity.tpl",
	)
}

func assetsTplEntityTpl() (*asset, error) {
	bytes, err := assetsTplEntityTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/entity.tpl", size: 328, mode: os.FileMode(511), modTime: time.Unix(1584588363, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTplExampleTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x54\x6d\x6b\xdc\x46\x10\xfe\x2e\xd0\x7f\xd8\x0a\xee\x22\x99\x8b\xe4\xd4\x49\x09\x82\xa3\xed\xbd\x84\x06\x92\xd4\xc5\x2e\xfd\x50\x4a\xd8\x93\xf6\xee\x16\x4b\xbb\xba\xdd\x55\x1a\xd7\x18\x2e\x50\x9a\x84\x24\x76\x20\x4d\x02\xad\xd3\x77\xf7\x85\x42\xed\x0f\xa5\x31\x97\x94\xfe\x19\x4b\x67\xff\x8b\x32\x92\xee\xcd\x75\x8a\x69\x09\xe4\xdb\xec\xec\xcc\xce\xcc\xf3\x3c\xb3\x11\xf6\x56\x70\x87\xa0\x70\x55\xf6\x02\x5d\xd3\x35\x1a\x46\x5c\x28\x64\xea\x9a\xe1\x63\x85\x5b\x58\x12\x47\xf6\x02\x43\xd7\x8c\x76\xa8\x0c\x5d\xbb\x8a\x8c\x0e\x55\xdd\xb8\x65\x7b\x3c\x74\x3a\xfc\xb4\xec\x05\xa7\x7d\x41\xaf\x11\xe1\x64\xaf\x40\xa8\x22\x52\x51\xd6\x01\x53\x2a\x41\x59\x47\x82\xc9\x88\x72\x62\x01\x01\x16\x94\x52\xab\x11\x41\x8d\x5a\x9d\xb3\x36\xed\x34\x99\xa2\x6a\x15\x49\x25\x62\x4f\xa1\x35\x5d\x7b\x87\x4b\x85\x10\x42\x79\x3e\x72\x9c\x64\x6b\x37\x79\xd2\xd7\xb5\x45\xe8\x0f\x21\x44\x99\x42\x8e\x33\xfc\x75\x27\xd9\xfc\x5e\xd7\xae\xe0\x90\xcc\x84\x0f\x3f\xff\x39\xbd\xf5\x54\xd7\x16\xb1\x94\xb3\xef\xec\x7c\x36\xfc\xe6\x86\xae\x35\x6a\x79\xce\xe4\x62\xf0\x20\xb9\x7f\x4f\xd7\xea\x5d\x2c\x24\x51\x53\x4f\x3d\x7f\x94\x65\x2c\xd3\x90\x7c\xc2\x19\x99\xdc\xa4\x8f\xff\x48\xee\x0e\x74\xed\x32\xbe\x7e\xd1\x0f\x48\xd1\x53\xba\xd5\x4f\x7e\xf8\x69\xf8\xcb\xe0\xf0\xf1\xef\x07\x7f\x7d\x95\x6e\x6c\x67\x11\xef\x46\x84\xcd\x44\xe4\x77\xe9\xc3\x5d\x5d\x5b\x07\x40\x1c\x07\x8d\x5d\xe9\xbd\xdf\x92\xc1\x03\x5d\x6b\xc7\xcc\x43\x17\x19\x55\x8d\x9a\xe9\xb5\x3b\x47\xf0\xb2\xd0\x9c\xec\x05\x76\xa3\x06\x88\xd1\x76\xd1\x98\xb4\x9b\xbd\x18\x07\x17\x78\xe0\x43\x8e\x3d\xea\xbb\x82\x0c\xc3\x82\xc8\x69\x27\xaa\x22\xe3\xd4\xdb\x92\x62\x67\xa9\x8b\x59\xa7\x8b\xe9\x29\x23\xeb\xc7\x97\x0c\xb9\x55\xd4\x0e\x95\xbd\x14\x09\xca\x54\xdb\x34\x4a\xd2\x2d\xc9\xb7\x94\x17\x99\x60\xf9\x96\x53\x92\x6f\x7a\x39\x5c\xd5\x92\x2c\x47\x60\xc1\xc3\x55\x25\x62\x52\x0e\xb8\x57\xbd\xc4\x3d\x1c\x94\x15\x0d\xc9\x55\xa8\x56\x2d\x49\xa3\x92\x37\x00\xe8\x17\x26\x70\x54\x98\x40\xfb\xc8\xcb\xc5\xc8\xcc\xb9\x2a\x0e\x05\x3f\x15\x5d\x8b\x45\x60\xbf\x17\x13\xb1\xda\x94\x1e\x8e\xc8\xcc\xb0\x56\x25\x93\x99\xc7\x19\x23\x9e\xa2\x9c\x55\x10\x11\x02\x26\x02\xc4\x80\x0b\xd3\xc8\xf5\x5a\x41\xbe\x64\x56\x86\x1f\x44\xbc\x56\x45\x8c\x06\x00\x53\x84\x19\xf5\x4c\x22\x84\x95\x01\x32\x79\xca\x5e\x22\xaa\x20\xb4\xce\x19\x93\x59\xe1\xc2\x61\x1d\x13\x08\xda\x98\x09\x04\x87\xa5\x6b\x82\xa8\x58\x30\x34\x89\xcf\xea\x5c\xc3\x02\x35\x6a\x23\x66\xf3\xf3\xda\x5a\x86\xd7\xfa\x7a\xa3\x85\xe6\xc6\x87\xcb\xdc\x27\xd9\xd6\x66\x2a\xa1\x8c\xaa\xfa\xf8\x29\x73\xc4\x34\x8c\x3c\xab\x9a\x62\xbb\x5c\x23\x00\x72\xba\x5c\x2a\xe0\x04\xe0\x76\x17\x16\xe6\xcf\x57\xf2\x6d\x72\x0d\xc1\x79\x7e\x83\xa5\x74\x8d\x33\xaf\x2f\x9c\x3d\xf7\x06\x9c\x73\x36\x5c\x63\x85\x32\x3f\x20\x3e\xb8\x0a\x4e\x5c\x23\x56\xed\xf3\x61\xeb\x2c\xf8\x0a\x40\xdc\x33\xf3\x95\xf1\x86\xb8\xe8\x5c\x25\x1b\xb2\x51\x43\xd5\x29\x5d\x5b\xba\x36\x3d\x62\x15\x5d\x21\x1f\x8f\x1d\x66\xa3\x66\x4d\x56\x24\xfd\x7a\xfb\x60\xe7\xbb\xf4\x76\x3f\xdd\xba\x3d\xfc\xe2\xd3\x7c\x57\x0a\x08\x96\x89\x54\xe3\xb4\x0b\x94\xf9\x1f\x74\x89\x20\xa6\x42\x73\xc5\x8f\x64\x2f\x67\xb0\x1c\x85\x4a\xd7\x5a\x9c\xaf\x00\x52\xe3\xec\x35\x28\xa6\xb8\xcf\xb3\xca\x82\xc8\x38\x50\x63\x09\x4d\xf5\x6a\x43\x19\xb3\x0c\xf9\x16\x7c\x22\x37\xff\x4c\x1f\xee\x1e\xde\xbc\x7b\xb8\xd5\x3f\xf8\xf1\x46\x7a\xe7\x79\xba\xb1\xbd\xbf\xd7\xdf\xdf\xbb\x93\xfd\xb0\xff\x10\x99\xb2\x9b\x42\x70\x31\x91\x19\x2c\xdc\x22\xec\x5b\xc0\xcc\xbc\xec\x87\xf3\x1f\xe5\x57\xf0\x41\x6c\x3c\x4d\x36\x1f\xc1\xff\x71\x7f\x63\x7f\xaf\x9f\x3e\xf9\xf6\xc5\x00\x5c\xc2\x52\x9d\x64\xf6\x7f\x1d\x4e\x48\x65\x32\x1a\x1c\xb7\x1f\x27\x68\xfd\x68\xdf\xfd\x67\x19\x3a\x9b\xc7\xb5\x5b\xe7\x31\x7b\x65\xfa\x4d\x6e\x7d\x99\x3c\x1b\xbc\x18\xdb\xba\x20\x58\x9d\x48\x59\xbe\xff\xdf\x74\x55\x54\x28\xfb\xfe\x4b\x9a\xa5\x79\x1d\x87\x51\x40\xde\x8f\xfc\x97\x3c\x49\x51\xe1\xff\x4d\xf2\x77\x00\x00\x00\xff\xff\x0f\x0d\x50\xf5\xaa\x08\x00\x00")

func assetsTplExampleTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplExampleTpl,
		"assets/tpl/example.tpl",
	)
}

func assetsTplExampleTpl() (*asset, error) {
	bytes, err := assetsTplExampleTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/example.tpl", size: 2218, mode: os.FileMode(511), modTime: time.Unix(1584588363, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTplInitTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x54\x5d\x8b\xe4\x44\x14\x7d\xee\x40\xfe\xc3\x35\xd0\xb3\x89\x84\xce\x7c\xec\x2e\xd2\xd2\xb4\xf6\xcc\x8a\x82\xa3\xee\xce\xbc\x4b\x4d\x52\xe9\x2e\x36\xa9\xca\x54\x55\x86\x19\xa5\x61\x7d\x11\x1f\x5c\x15\x56\xdd\x97\x55\x04\x15\x15\xc1\xdd\x07\x51\x18\x04\x7f\x4d\xf7\xe8\xbf\x90\x5b\x55\xe9\x4e\xda\x1d\xf6\x29\xa9\x73\x4f\xdd\x8f\x53\xa7\x8a\x95\x95\x90\x1a\x42\xdf\xeb\x05\x19\xd1\xe4\x84\x28\x9a\xa8\xd3\x22\x40\x20\x2f\xb5\xf9\x72\xaa\x93\x5a\x5a\x4c\x69\xc9\xf8\x54\x05\xbe\x17\xf9\x9e\xef\x9d\x11\x69\x36\x97\x44\x69\x2a\x0f\x26\xf0\xb2\x3a\x2d\x06\x07\x13\xdf\xeb\xa9\x82\x9c\xd1\x36\x62\x36\xe4\x35\x4f\x81\x71\xa6\xc3\x08\x3e\xf4\xbd\x5e\x9a\x4f\x61\x38\x82\xec\x64\x5f\xf0\x9c\x4d\x11\xea\xbd\x29\x94\x1e\x02\x00\x04\x85\x48\x49\x31\x13\x4a\x07\x31\x06\xde\x13\xd2\x06\xf6\xf6\xb6\x6f\x1b\xe4\x1d\x52\x52\x4b\x95\x42\x34\x2c\xa2\x94\xc5\x76\x76\xf7\x6e\xde\xba\x6d\xd1\x83\x89\xe3\x06\xf7\x19\xcf\x0a\x9a\x59\x78\x7f\x46\xa4\xa2\x7a\x08\x41\xad\xf3\x57\xca\x93\x9b\x16\x3e\x24\xe7\xef\x56\x94\x0f\x61\x67\x7b\xbb\x01\xde\xca\x0a\x3a\x84\x5b\x66\x3d\x6f\xcf\x3c\x32\x03\x1d\x4c\xc2\x34\x9f\x46\xbe\xd7\x4b\x92\x66\xf6\x8d\xc8\x1c\x15\xd0\x17\x15\x5d\xcd\x0b\x4a\xcb\x3a\xd5\x46\x0a\x1c\x1b\xdb\x06\xab\x31\x24\xc9\xe2\xc9\xb3\xc5\x37\x0f\x7c\xcf\x0c\x6e\x42\x8c\x9b\x6f\x92\x5c\xfd\xfa\x74\xf1\xf9\xf7\xbe\x67\x14\xe8\xee\xba\xfa\xf2\xe7\xe5\x27\x7f\xe2\x2e\xa2\xd4\x46\xc2\xa7\x1f\x5f\x7d\xf7\x91\xef\x39\x35\x3a\xa1\xcb\x47\x8b\x2f\x1e\xfa\x5e\xa3\x48\x3b\xe1\x5f\x5f\xdb\x5d\xc7\xac\xa4\x1f\x08\x4e\xd7\xa1\xe5\xe3\x3f\x16\x9f\x5e\xfa\x5e\xa3\x4f\xab\xc3\xe5\x93\x07\x8b\x1f\x7e\xba\xfa\xe5\xf2\xdf\xc7\xbf\xff\xf3\xf7\xb7\xcb\xcf\x7e\xb4\x34\xd4\xf5\x7f\x34\x4b\x58\x7e\xf5\xcc\x89\x94\x24\xb0\x82\x96\x0f\x7f\x5b\x5c\x3e\x6a\x59\xc7\xea\xb9\xd2\x30\x6a\x2c\x66\x54\x64\xb9\xeb\x4e\x0d\xee\x9c\xd6\xa4\x78\x43\x14\x19\xd2\x07\x4d\xf3\x31\x04\x81\xf5\x5e\xaf\x0d\xc3\x08\x82\x1b\xaf\x2b\x46\x92\xa3\x19\xe1\xd3\x19\x61\x37\x02\x77\xd0\x99\xe2\xe8\xd1\xbc\xd4\x83\xa3\x4a\x32\xae\xf3\x30\xe8\xab\x61\x5f\xbd\xa6\xd3\x2a\xc4\xbf\x2c\x4a\xfa\x6a\x9c\x5a\xe5\x46\x7d\xb5\x55\xe1\x1f\xa6\x1e\x69\x59\xd3\xad\x42\xa4\xa3\xb7\xd1\xcb\x5b\x9a\x95\xf4\x7d\xac\x37\xea\x2b\x6b\x35\x6c\x02\x0f\x63\xb5\xc0\x63\x5b\x2d\xd0\x14\xeb\x88\x90\xeb\x85\x3d\xc2\xd5\xd2\x1d\x9b\x59\xd7\xb2\x18\xdc\xad\xa9\xbc\xb8\xa3\x52\x52\xd1\xce\xf8\x11\x32\xd0\xa5\x19\xcd\xa9\x04\x14\xd5\x5d\x45\x94\x8e\x4a\x89\xa3\x4a\x9a\x8a\x33\x2a\xc3\xe8\x55\x83\xbc\x34\x02\xce\x0a\x4b\x72\xac\x1d\xa4\x1d\xba\x1b\xb0\x5f\x08\x45\x1d\x79\xa7\xcb\xee\x55\x84\xb3\x34\xc4\x40\x64\x80\x79\x2b\xc9\x2e\x26\x39\xb2\x77\xa5\x9d\x63\xf7\xba\x1c\xbb\xed\x1c\x2b\xd4\x80\x88\xcd\x43\xfc\x4d\x05\xe7\x34\xd5\x4c\xf0\xb8\x99\x07\xfd\x81\xbe\x0b\x83\xf2\x02\x5f\xb7\x18\x32\xc5\x23\xeb\x96\xcd\x01\x3b\x69\xe7\x9d\x7c\x83\x23\xaa\x9d\x83\xf7\x05\xe7\xca\x08\xeb\x80\xe8\x79\x4c\xbc\x12\x1d\x26\x02\xc8\x94\x54\xd7\x92\xc3\x7a\x83\xf1\x7d\x92\x40\x8a\x32\x40\x76\xe2\xec\xbe\xa1\x30\x36\x2b\x64\xe3\xf3\xd5\x03\xd4\x6e\xdf\x65\x6e\x62\x03\xb7\xd3\xcd\xe2\xa2\x9c\x15\xd7\x14\xec\x9e\x46\xb7\x5e\xf3\xaa\x3d\xa7\x9c\x0b\xbd\xb0\x9a\x96\x84\x2b\x62\x26\x06\xa5\x89\xd4\xae\xec\xf1\xf9\x84\x4e\x19\x0f\x23\x08\xcd\x65\x3e\x3e\x8f\x6d\x6d\x6b\xcd\xcd\xa1\x1c\xb9\xc9\x2a\x69\x45\x89\x06\x49\x55\x25\xb8\xa2\xa0\x05\x8c\xe3\x71\x3c\x76\xc9\xef\x99\xf0\xdd\x9a\x2a\xac\x7b\x48\xe4\xfd\x30\x15\x35\xd7\xf8\x08\x45\xcd\x5b\xd6\x2a\xd3\xbc\x1f\xc7\x92\x95\xf7\xd8\x74\xa6\xc3\x06\xb1\x99\xc2\x60\x1c\x07\x31\x98\x1c\x51\x0c\x41\x1c\x60\x27\xff\x05\x00\x00\xff\xff\xe3\x1c\x96\xb0\x4b\x07\x00\x00")

func assetsTplInitTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplInitTpl,
		"assets/tpl/init.tpl",
	)
}

func assetsTplInitTpl() (*asset, error) {
	bytes, err := assetsTplInitTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/init.tpl", size: 1867, mode: os.FileMode(511), modTime: time.Unix(1585828214, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTplMarkdownTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\xc1\x6a\xea\x40\x14\x86\xf7\x81\xbc\xc3\x81\xb8\x50\xb8\xc9\x03\x08\xf7\x6e\xae\x1b\xb9\x17\xeb\x42\xba\x11\x17\xb1\x9e\x16\x69\x92\x96\x68\xa0\x61\x66\xc0\x45\xa1\x85\xd2\xd6\x45\xc1\xac\x0a\x05\x11\xbb\xa8\x52\xda\x82\x98\x3e\x8e\x19\xf5\x2d\x4a\x66\x9c\xd4\x56\x9a\xc5\x84\xf9\xff\xc3\xf9\x4f\xbe\x13\xc3\x80\xd5\xc3\x98\xf7\x62\x3e\x3a\xd7\x35\x5d\xa3\xc9\xfc\x36\xb9\x1c\xd0\x54\x4d\xfa\xd7\x40\x61\x35\x7d\xe5\xd1\x0d\x50\x48\x86\x17\xfc\x65\x0c\x54\xd7\x68\xd1\x34\x4d\x71\xa8\x13\xe4\x4b\xd7\x08\xf1\x6d\xef\x08\xc1\xaa\xd9\x4d\x07\xff\xb7\x3b\x5d\xc6\x28\x10\x62\x95\xbd\x16\x9e\x31\x46\xeb\xe9\x45\x98\x15\xdb\x45\xc6\x1a\x79\xe3\xd3\xfd\x6a\x15\x40\x3e\x94\x10\xeb\xef\x89\xeb\xa2\x27\x9a\x89\x14\xf4\x5a\x8c\xa5\x03\xab\xc0\xdc\x31\x86\xbf\x20\xd7\xee\xa2\x0b\xc5\xdf\x60\x95\xb0\x73\x20\xe3\x75\xcd\x30\x0c\x20\x44\x78\x2a\xc9\x52\xf7\xad\x3c\x5d\xfb\x93\x95\x65\x79\x69\x86\x29\x68\x3c\x0d\xf8\xe4\x4d\xd2\x90\xa4\x84\x00\x14\x96\xcf\x71\x72\x7f\x05\x14\x78\x34\x4d\xfa\xa3\xc5\x6c\xbe\x7c\x9c\x53\x58\xc7\xd1\x6a\x32\x4c\x7a\xef\x99\xc3\xa3\xe9\x62\x16\xaf\xef\x26\xbb\x58\x15\x55\x73\x97\xed\x37\x61\x0b\xb2\x1c\x75\x03\x59\x40\x72\x02\xd7\x93\x1f\x03\x20\xb0\x4b\xa9\x16\x9e\x62\x25\x70\x9b\xe8\x33\x26\xf5\x72\xa7\x12\x38\xce\x66\x37\x25\x3c\xb4\x03\xa7\xbb\x6f\x3b\x01\xaa\x82\xaa\xdf\x76\x6d\x3f\xfc\x87\xa1\x52\x64\xab\x1f\x16\x51\xaf\xed\x55\x1b\x79\x23\xfb\x97\x0a\x99\xf7\x11\x00\x00\xff\xff\x54\x04\xaf\xad\x62\x02\x00\x00")

func assetsTplMarkdownTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplMarkdownTpl,
		"assets/tpl/markdown.tpl",
	)
}

func assetsTplMarkdownTpl() (*asset, error) {
	bytes, err := assetsTplMarkdownTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/markdown.tpl", size: 610, mode: os.FileMode(511), modTime: time.Unix(1584588363, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTplTablesTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe2\xe5\x4a\xce\xcf\x2b\x2e\x51\xd0\xe0\xe5\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xd4\x51\x50\xc9\x2c\x49\xcd\x55\xb0\xb2\x55\xd0\xab\xad\x05\xc9\x80\xf9\x7a\xa1\x05\x05\xa9\x45\x21\x89\x49\x39\xa9\x7e\x89\xb9\xa9\xb5\xb5\x0a\x0a\x0a\x0a\xb6\x0a\x4a\x30\x69\x24\x19\x25\x05\x7d\x7d\x05\x98\xb8\x73\x7e\x6e\x6e\x6a\x5e\x49\x6d\x6d\x75\x75\x6a\x5e\x0a\xc8\x40\x4d\x40\x00\x00\x00\xff\xff\xf6\x4d\x87\xcf\x77\x00\x00\x00")

func assetsTplTablesTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplTablesTpl,
		"assets/tpl/tables.tpl",
	)
}

func assetsTplTablesTpl() (*asset, error) {
	bytes, err := assetsTplTablesTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/tables.tpl", size: 119, mode: os.FileMode(511), modTime: time.Unix(1584588363, 0)}
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
	"assets/tpl/curd.tpl":     assetsTplCurdTpl,
	"assets/tpl/e.tpl":        assetsTplETpl,
	"assets/tpl/entity.tpl":   assetsTplEntityTpl,
	"assets/tpl/example.tpl":  assetsTplExampleTpl,
	"assets/tpl/init.tpl":     assetsTplInitTpl,
	"assets/tpl/markdown.tpl": assetsTplMarkdownTpl,
	"assets/tpl/tables.tpl":   assetsTplTablesTpl,
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
	"assets": &bintree{nil, map[string]*bintree{
		"tpl": &bintree{nil, map[string]*bintree{
			"curd.tpl":     &bintree{assetsTplCurdTpl, map[string]*bintree{}},
			"e.tpl":        &bintree{assetsTplETpl, map[string]*bintree{}},
			"entity.tpl":   &bintree{assetsTplEntityTpl, map[string]*bintree{}},
			"example.tpl":  &bintree{assetsTplExampleTpl, map[string]*bintree{}},
			"init.tpl":     &bintree{assetsTplInitTpl, map[string]*bintree{}},
			"markdown.tpl": &bintree{assetsTplMarkdownTpl, map[string]*bintree{}},
			"tables.tpl":   &bintree{assetsTplTablesTpl, map[string]*bintree{}},
		}},
	}},
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
