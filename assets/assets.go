// Code generated by go-bindata.
// sources:
// assets/template/error.tmpl
// assets/template/login.tmpl
// assets/template/paste_form.tmpl
// assets/static/css/pure.css
// DO NOT EDIT!

package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _templateErrorTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x90\x3f\x4f\xc3\x30\x10\xc5\xe7\xf2\x29\x8c\x07\xba\x90\x5a\xd9\x90\xb0\xb3\x14\x66\x18\xba\x30\xba\xf6\x15\x5b\x5c\x9c\xca\x77\x69\xa9\xaa\x7c\x77\x9c\x1a\x44\x55\x2f\xf7\xc7\xef\xfd\xa4\x7b\xfa\xfe\xe5\x6d\xbd\xf9\x78\x7f\x15\x81\x7b\xec\xee\x74\x2d\xa2\x3c\x1d\xc0\xfa\xda\x5e\xc6\x1e\xd8\x0a\x17\x6c\x26\x60\xb3\x1c\x79\xd7\x3c\x2d\xaf\xbe\x39\x32\x42\xa7\x55\xad\x37\xb6\x64\x7b\x30\xf2\x10\xe1\xb8\x1f\x32\x4b\xe1\x86\xc4\x90\xd8\xc8\x63\xf4\x1c\x8c\x87\x43\x74\xd0\x5c\x86\x47\x11\x53\xe4\x68\xb1\x21\x67\x11\x4c\x2b\xaf\x60\x18\xd3\x97\xc8\x80\x46\x12\x9f\x10\x28\x00\x14\x5a\xc8\xb0\x33\x52\x11\x5b\x8e\x4e\x39\x22\xb5\x1f\x33\xac\x4a\xf3\xeb\xd5\xea\xff\x16\xbd\x1d\xfc\x49\x38\xb4\x44\x46\xce\xba\xe6\xb3\xa8\x16\x0b\x1d\xda\xee\x7c\x5e\x6d\xe0\x9b\xa7\x49\x3c\x6c\x47\xc4\x67\xa1\xa9\xb7\x88\xf3\x7e\x3d\x78\x98\x26\xad\xea\xa2\x00\xdb\x3f\xf4\xcc\x2b\xc1\xa9\x9a\xdc\x4f\x00\x00\x00\xff\xff\x2d\x7d\xb5\x4d\x51\x01\x00\x00")

func templateErrorTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateErrorTmpl,
		"template/error.tmpl",
	)
}

func templateErrorTmpl() (*asset, error) {
	bytes, err := templateErrorTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/error.tmpl", size: 337, mode: os.FileMode(420), modTime: time.Unix(1434915168, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templateLoginTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x51\x3d\x73\xc3\x20\x0c\xdd\xfb\x2b\x28\x4b\x96\x62\x2e\x5b\x07\x70\x87\xb6\x73\x33\x74\xe9\x48\x41\x31\x5c\x31\xf6\x81\xec\x5c\x8e\xf3\x7f\xaf\xe3\x8f\xd6\x0e\x8b\xd0\x93\xde\x93\x74\x4f\x3c\xbe\x7d\xbc\x7e\x7e\x9d\xde\x89\xc5\xda\x97\x0f\x62\x0e\x64\x7c\xc2\x82\x32\xf3\x77\x4a\x6b\x40\x45\xb4\x55\x31\x01\xca\x43\x87\x67\xf6\x7c\xd8\x94\xd1\xa1\x87\x52\xf0\x39\xde\xd1\x82\xaa\x41\xd2\xde\xc1\xa5\x6d\x22\x52\xa2\x9b\x80\x10\x50\xd2\x8b\x33\x68\xa5\x81\xde\x69\x60\x53\xf2\x44\x5c\x70\xe8\x94\x67\x49\x2b\x0f\xf2\x48\x37\x62\xde\x85\x1f\x12\xc1\x4b\x9a\xf0\xea\x21\x59\x80\x51\xcd\x46\x38\x4b\xca\x13\x2a\x74\x9a\xeb\x94\x78\xdb\x45\x28\xc6\xcf\xc2\x15\xfc\xff\x16\xf1\xdd\x98\x2b\xd1\x5e\xa5\x24\xe9\xad\x8f\x55\x6b\x57\xe7\x77\x78\xc7\x8e\x6c\x3b\x3d\xe7\xa8\x42\x05\xa4\x38\xc5\xa6\x77\x06\x62\x1a\x86\xbf\xe2\xb2\x5e\xb9\x03\x26\x50\xad\xfb\xf9\xa6\x72\xe1\xa5\x5d\xc8\x32\xe7\x62\x18\x68\x39\x05\xc1\xd5\x9e\x29\xf8\x56\x2b\x67\x08\x66\x19\x26\x78\xb7\x1a\xc4\x6f\xa7\x8c\x9e\xf1\xd9\xb4\xdf\x00\x00\x00\xff\xff\xcf\x9d\x1c\x50\xcc\x01\x00\x00")

func templateLoginTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateLoginTmpl,
		"template/login.tmpl",
	)
}

func templateLoginTmpl() (*asset, error) {
	bytes, err := templateLoginTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/login.tmpl", size: 460, mode: os.FileMode(420), modTime: time.Unix(1435084262, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templatePaste_formTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x55\x4f\x6f\xdb\x3a\x0c\xbf\xbf\x4f\xc1\xa7\x4b\x12\xbc\x3a\x6e\xd0\xbe\xa2\xc0\xec\x1c\xb6\x05\xd8\x80\x01\x2d\xd0\x5e\x76\x54\x2c\xa6\xd6\x26\xcb\x9e\x44\x39\xcd\xd6\x7e\xf7\xc9\xb2\x93\xd8\x69\x82\x6e\xd5\xc5\x94\xc8\xdf\x8f\x7f\x44\xca\xc9\xbf\x1f\x6f\x3e\xdc\x7f\xbd\x5d\x40\x4e\x85\x9a\xff\x93\xb4\x1f\xf0\x2b\xc9\x91\x8b\x56\x0c\xdb\x02\x89\x43\x96\x73\x63\x91\xd2\x91\xa3\x55\x74\x3d\xea\xa9\x49\x92\xc2\x79\x12\xb7\xdf\x03\x98\xe6\x05\xa6\xac\x96\xb8\xae\x4a\x43\x0c\xb2\x52\x13\x6a\x4a\xd9\x5a\x0a\xca\x53\x81\xb5\xcc\x30\x0a\x9b\x33\x90\x5a\x92\xe4\x2a\xb2\x19\x57\x98\xce\x58\x8f\x4c\x49\xfd\x1d\x0c\xaa\x94\x59\xda\x28\xb4\x39\xa2\x67\xcb\x0d\xae\x52\x16\x5b\xe2\x24\xb3\x38\xb3\x36\xae\x9c\xc1\xa9\x17\x3a\x6c\x12\xef\x73\x49\x96\xa5\xd8\x40\xa6\xb8\xb5\x29\x6b\xec\xa2\x87\xbe\x87\x55\x69\x0a\xe0\x19\xc9\x52\x7b\x4a\x06\x3e\xfc\xbc\x14\x29\xbb\xbd\xb9\xbb\x67\x03\x98\x8b\x66\xd1\x0c\x82\xd8\x80\x7a\x24\x6d\x3d\xf0\x91\xb8\xc1\x6d\xee\x5d\xc6\x0c\xa4\x27\xab\xb8\x25\x8c\x76\x47\x95\xe2\x19\xe6\xa5\x12\x68\x3a\x1d\xec\x74\x06\x7f\x38\x69\x50\xbc\x74\xcd\x9a\x62\x77\x4e\x0e\x7c\x0b\x59\x1f\x64\xd8\x73\x5b\x56\x4d\x72\xf6\x20\xde\xa3\x38\x17\x5d\x44\x97\x47\x0c\xdb\xbb\xe0\x4b\x54\xe0\x53\x4f\x19\x3e\x56\x3e\xc6\x48\xea\x13\xb6\xcd\x5a\x04\x9b\x93\xea\xc4\xa2\xc2\x8c\xba\x6a\xed\x09\x43\xe0\x7f\xc2\x1f\x48\xda\xdc\xa0\xe6\xca\x79\x96\x73\x06\x02\x57\xdc\x29\x9a\x6b\xac\xd1\x24\x71\xab\xff\x2b\x92\x8b\xf3\x73\x36\xe7\x2b\x42\x03\xff\x43\x21\xb5\x23\xb4\x6f\x23\xba\xda\x33\xcd\x20\x2f\xdd\xdb\xe2\xb9\xbe\xba\xec\xf3\x08\xbe\x79\x9d\x26\x89\xdb\xea\x9e\xb8\xc9\x38\x5c\xe5\x91\x7e\x88\x7d\x43\x1c\x39\x5e\x3a\x22\x1f\x10\x6d\x2a\x1f\x8f\x75\xcb\x42\xd2\x70\x34\x3a\x83\x9e\x1c\x55\x46\x16\xdc\x6c\x60\xd7\xbf\xbe\xaf\xee\x02\x34\x89\x5b\x93\x83\x1e\x1e\xfa\x4e\xe2\x66\xc6\xb6\xd3\xdc\x8c\x70\x27\xdb\xcc\xc8\xca\x27\x36\x5e\x39\x1d\xa6\x16\xc6\x13\xf8\x15\x74\x6b\xa9\x45\xb9\x9e\x72\x21\x16\xb5\x9f\xa5\x2f\xd2\x77\xbf\x46\x33\x1e\xa9\x92\x8b\xd1\x19\xbc\x44\x34\x6b\x77\xca\xc5\x37\x67\xe9\x13\xca\x87\x9c\x06\x16\xcd\xaa\xb9\x81\x1c\xd2\xad\x0f\xa9\x3d\x71\x6b\x0a\x4f\x4f\x20\xca\xcc\x15\xde\xe5\x74\x2b\x2c\x14\x86\x7d\xa6\xa4\xff\xb4\x86\xef\x06\x84\x3b\xc8\x03\x6e\xad\xdf\x6f\x3e\x8b\xf1\xc1\x5b\x31\x99\x86\x87\x6f\x9a\xb7\xbe\x52\x1f\x44\xf4\x1a\x76\x3b\xf0\x93\x81\x7b\x8f\xbb\x84\xff\x60\x54\x3d\x8e\xf6\x91\x3c\xef\xa4\x61\xf6\x7b\x8b\x93\x45\x35\x68\xe5\x4f\xf4\x65\xed\x23\x7d\x91\xb9\xb2\xd8\xe1\x9f\xf7\xdb\xe7\xf1\xa4\x3b\xf4\xbd\xd9\x5d\xa1\x7f\xa5\xc3\xcf\xe7\x77\x00\x00\x00\xff\xff\xd5\xdb\xd0\xef\x94\x06\x00\x00")

func templatePaste_formTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatePaste_formTmpl,
		"template/paste_form.tmpl",
	)
}

func templatePaste_formTmpl() (*asset, error) {
	bytes, err := templatePaste_formTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/paste_form.tmpl", size: 1684, mode: os.FileMode(420), modTime: time.Unix(1434918967, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _staticCssPureCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x3c\x59\x8f\xdc\x36\x93\xef\xfe\x15\x8a\x0d\x23\x33\x03\x49\x23\xa9\x6f\x09\x36\x36\x87\x93\xf5\x22\x09\x16\x48\x1e\x76\xe1\xcc\x02\x6c\x89\xdd\xad\x1d\x5d\x90\xd4\x73\xb8\xb7\xff\xfb\x16\x0f\x49\xbc\xd4\xdd\xe3\x7c\xfe\x90\x87\xcf\x0d\xc7\x22\x59\x2c\x16\xeb\x22\xab\x48\xe6\xf6\xe6\x9b\x57\xff\xb9\xaf\xb1\xf5\xe0\xb9\x73\xd7\x7b\xf5\x43\x59\x3d\xd7\xe9\x76\xd7\x5a\x81\xe7\x4f\xad\xff\x46\xbb\xb2\xfc\xc6\xfa\x58\xc4\xae\xf5\x5d\x96\x59\xb4\xa9\xb1\x6a\xdc\xe0\xfa\x01\x27\xee\xab\x5f\xd2\x18\x17\x0d\x4e\xac\x7d\x91\xe0\xda\x6a\x77\xd8\xfa\xfe\xf7\x1f\x2d\x5e\xed\xbe\xda\xb5\x6d\xd5\x84\xb7\xb7\xdb\xb4\xdd\xed\xd7\x6e\x5c\xe6\xb7\xcf\x04\xe7\x6d\x05\x83\xde\xae\xb3\x72\x7d\x9b\xa3\xa6\xc5\xf5\xed\x2f\x1f\x7f\xf8\xf0\xdb\xef\x1f\xdc\x3c\x79\x75\x73\xfb\x8a\xd0\x55\x94\x75\x8e\xb2\xf4\x33\x76\xe3\xa6\xb1\x1e\xfe\x67\xe2\x7a\xd6\xff\x59\xbf\x7e\xfc\xa3\x43\x0f\x25\xc0\xeb\xa6\xe5\x6d\x0f\x2a\x4c\xe0\x2a\xbe\xb6\x7e\x4b\xe3\x32\x43\x8d\xf5\x33\xca\x32\xb4\xdd\x01\x85\xa8\x48\xac\xff\x28\x0b\xd4\xee\x50\x61\xfd\x86\x51\xc6\x47\xb3\x94\xd1\x60\x30\x37\x38\x3b\x9c\x75\x73\xbb\x6b\xf3\xec\xb0\x29\x8b\xd6\xd9\xa0\x3c\xcd\x9e\xc3\x06\x15\x8d\x03\xfc\x49\x37\x91\x93\x37\x4e\x8b\x9f\x5a\xa7\x01\x58\x07\x25\xff\xbb\x6f\xda\xd0\xf7\xbc\xb7\x91\xf3\x88\xd7\xf7\x69\x6b\x6e\x3d\xae\xcb\xe4\xf9\x90\xa3\x7a\x9b\x16\xa1\x77\x44\x75\x9b\xc6\x19\xb6\x51\x93\x26\xd8\x4e\x70\x8b\xd2\xac\xb1\x37\xe9\x36\x46\x55\x9b\x96\x05\xf9\x04\x6e\xda\x9b\xb2\x04\x46\xda\x3b\x8c\x12\xf2\xcf\xb6\x2e\xf7\x95\x9d\xa3\xb4\xb0\x73\x5c\xec\xed\x02\x3d\xd8\x0d\x8e\x69\x8f\x66\x9f\x03\xfa\xe7\x43\x92\x36\x55\x86\x9e\x43\x10\x44\x7c\x7f\x44\xfb\x24\x2d\xed\x18\x15\x0f\xa8\xb1\xab\xba\xdc\x82\x9c\x1b\xfb\x01\x46\x2d\x7b\xc8\xb4\xc8\xd2\x02\x3b\xb4\x43\xf4\x80\x09\x69\x28\x73\x80\x19\xdb\x22\x5c\xa3\x06\x93\x56\x86\x28\x2c\xca\xf6\xea\x53\x0c\x9c\xa9\xcb\xac\xb9\xbb\xee\x51\x14\x65\x81\xa3\x1d\x26\x42\x82\xd9\x7d\xda\xa5\x49\x82\x8b\x3b\xbb\xc5\x39\x34\xb7\x58\x82\x3b\xa2\xc3\x1a\xc5\xf7\x64\x2e\x45\xe2\x80\x30\xcb\x3a\x6c\x6b\xe0\x70\x85\x6a\x5c\xb4\x47\x14\x22\x98\xd1\x03\x30\x27\xdc\x95\x40\xce\xa1\xdc\xb7\x84\x04\xc2\xb6\xf5\xba\xfe\xd4\xa6\x6d\x86\xef\x0e\xeb\xb2\x06\x9e\x38\xeb\xb2\x6d\xcb\x3c\xf4\xab\x27\x2b\x81\x4f\x9c\x1c\xd7\x76\x03\xe4\x15\x5b\x26\xc1\x47\x46\xd4\xc2\xf3\x8e\xc9\xa6\x60\x75\x4d\xfb\x9c\xe1\x30\x6d\x61\x8a\xf1\x71\xe7\xf3\x4a\x90\x58\x18\xe0\x3c\xe2\x42\x72\xe7\x0b\x9c\x5b\xde\x11\x8a\xf7\x02\xc1\xe1\x9b\xcd\xc6\x8b\x18\xd5\x6f\x3c\xc0\xda\x80\xda\x64\x02\x8a\x25\x08\xbb\xd9\x03\x11\xfb\x4a\xa8\x5d\xcc\xde\x46\x94\xcb\x1d\x93\xa2\xaa\x6c\x52\x22\xb8\xb0\xc6\xc0\x22\x98\xef\x28\xeb\x09\xa6\xb6\xac\x42\xc7\x9d\xe1\x9c\xe0\x3e\xf0\x49\x3b\x6e\x40\x6a\xd2\x7c\xcb\xb9\x01\x2c\x6a\x1e\xb6\x54\x4a\x61\x0d\xaa\x73\x7d\x20\x0c\xdc\x64\xe5\x63\xc8\x44\x72\x64\x7a\xd5\x29\xa2\x0f\x33\x9c\x7a\xd5\xd3\x71\x57\x1f\x9c\xbc\xfc\x0c\xdc\x7c\x22\xf4\xa6\xc5\x36\x24\x52\x06\x71\x90\xaa\x68\xa4\xba\x17\x78\x05\x28\xfb\x91\xd0\xbe\x2d\x8f\x71\x09\x7a\x7d\xbf\x4e\x40\xe7\xb0\xdd\xa0\xbc\x92\xec\x29\x2f\x8b\x12\xc4\x1d\x63\xbb\xff\x8a\x06\x5e\x01\x55\xc7\xf5\x1e\x66\x58\xd8\x69\x51\xed\x5b\xbb\xac\x5a\xa6\xf9\xc0\x10\xd0\x76\x9b\x58\x18\xe8\x0a\x3a\x30\x31\xa4\x05\x78\x81\xb4\xa5\x18\xfa\x42\x6f\x6a\x0c\xd3\x40\xde\x43\xda\xa4\xeb\x0c\x77\x23\x30\x94\x07\x6a\xb4\x54\x0b\x37\xe0\x08\x98\x9e\x72\x08\xe2\x0d\x2c\x4a\xc8\xa7\xf6\xb9\xc2\xef\x58\xf5\x9d\x2d\x54\x11\xe7\xd9\x4a\x35\x20\xa5\x3c\x6d\xef\x0e\x9d\x53\x40\x55\x85\x11\xa0\x8f\x71\xc8\xfa\x47\xf1\xbe\x6e\x80\xf8\xaa\x4c\x81\xa1\x35\x1f\xec\x13\x18\x0a\x02\xea\x92\x3b\x71\xd8\xbe\xf2\xc0\x3b\x25\x78\x83\xf6\x59\xcb\x3b\x85\x21\x95\xdd\xa6\x8c\xf7\x8d\x93\x16\x05\x78\x0a\xda\x4f\xaf\xef\xd5\x24\xaa\x50\x92\x10\x71\x7a\x47\x0a\x7a\x10\x75\x93\xb9\xc2\xa3\x30\x9b\x78\x87\xe3\x7b\x90\xb8\x3c\x69\x04\x1e\x81\xd8\x61\xaf\x1b\xbd\x49\x3e\xa9\xf8\x59\x8f\x62\x9f\xaf\x71\x7d\x07\x74\x71\xae\x50\xa2\x9c\xa6\x4a\x0b\x47\x14\xf8\x08\x34\xf8\x01\x19\xfa\xc0\x09\xa6\x1a\x27\x32\x1f\x58\x1d\xef\x8c\xcc\x27\x72\xde\xa4\x38\x4b\xa2\x53\xfa\xde\x75\x7c\x91\x39\x18\x28\x18\x68\x67\x15\x4e\x4c\x88\xc8\x0c\x93\x1d\xeb\x90\xe0\xb8\xac\x11\xf1\x13\xa6\xd9\x50\x35\xa5\xd3\x01\xfd\xeb\x84\x4b\x5c\x61\x53\x66\x69\x62\x35\x69\x06\x5a\xdf\x9b\x82\x15\x54\x83\x60\xdc\x09\xb8\x0e\xcb\x9d\x07\xf4\x9f\x05\xf1\x23\x19\xde\xe2\x22\x31\xe9\x48\x6f\x70\xb2\x91\x77\x76\xa9\x79\xda\x96\xa8\x6b\xe7\xa1\xc1\x48\x33\x54\x35\x38\xec\x3e\x22\xde\x40\xec\x9e\xe3\x4f\xec\x76\x77\x18\xc6\x73\x99\xb7\xb2\xbb\x85\x44\x5a\x3f\xbe\x49\xf3\xaa\xac\x5b\x04\x4b\x85\x4b\xb6\x1c\x0e\xf1\x7f\x39\x7a\x72\x1e\xd3\xa4\xdd\xb1\xb5\x58\x50\x8c\x48\x5e\x10\x59\x97\xed\x21\xc3\x6d\x2b\xd0\xe0\xb8\x13\x70\x3b\xd1\x8d\x52\xcd\x4c\x21\xba\x79\x04\x8a\x05\xd8\xe9\x04\x60\xa9\xcb\x80\x15\x0b\xa6\x42\x6a\x81\x19\x69\x0e\xee\xab\xa9\x30\x4e\x22\xd1\xdb\xfd\x54\x63\xfc\x3b\x78\x16\xfb\xbb\x3a\xcd\x4b\xfb\xf5\x8f\x75\x09\xc2\x21\x35\xaf\xed\x7f\xc7\x20\x22\xe2\xf5\x49\x23\xca\x6c\x61\xa7\xd1\xd1\xdd\xc9\x7d\x93\xe1\x41\x33\x49\xc1\xa1\x92\xa8\xcb\x47\xeb\xb1\x46\xd5\x00\x0f\x3b\x14\xd2\x4c\x15\x99\x7f\x2b\xa0\xa4\x9a\x2e\x32\x0e\xd7\xdf\x90\x02\x35\x20\xe4\xb6\x1f\x62\x14\x60\xac\xe1\xe8\x96\x15\xae\x91\x53\x16\xd9\xb3\x15\x3a\xa5\x03\xbe\x9f\x3a\x1f\xbb\x63\xba\x81\x8d\x5c\x20\x7b\xf3\x6e\xe4\x46\xae\x8d\x3e\x97\x64\xad\x8f\xcc\x52\x92\xb0\xf3\x3a\x65\x4d\x85\x65\x54\x95\x1b\xd5\x64\x4e\xa0\xf5\x29\x86\xad\x65\x63\xdd\xbc\x7b\xcd\xa8\x7a\x7d\x37\xb2\x0f\xec\xc8\x76\x7c\xbb\xff\x92\xbe\x03\xe1\x7b\xd2\x7f\x07\xc2\xb7\xef\x4c\xfb\xef\x89\xf0\xed\x3b\x33\x01\x7e\x26\xc0\x0c\xdf\x53\xe1\x7b\x26\x7c\xfb\xce\x5c\xa8\x9f\x0b\xf5\x4b\x01\xcf\x52\x80\x19\xbe\x17\xc2\x37\xcc\x25\x10\x80\x84\xc2\x42\x2c\xf8\x12\x1c\x4c\x7a\x2a\x50\x1e\x88\xd3\x13\x0a\x53\xb1\x30\x13\x0b\x73\xb1\xb0\x10\x0b\x4b\xb1\xb0\x12\x0b\xbe\x27\x95\x24\x1a\x7c\x89\x08\x5f\xa2\xc2\x97\xc8\xf0\x25\x3a\x7c\x89\x10\x5f\xa2\xc4\x97\x48\xf1\x25\x5a\x02\x89\x96\x40\xe6\x87\x44\x4b\x20\xd1\x12\x10\x5a\xfe\x0e\x06\xc0\x64\x78\x60\x8e\x74\xea\xfa\xf3\xf9\xe2\x2d\x38\xbf\xae\x38\x99\x2d\xde\x1e\x8d\x2a\x12\x0c\xdd\x96\xee\x04\xfe\xf4\xdd\xa0\xe8\x05\x13\xb1\xdb\x52\x52\x0b\xde\xcb\x0f\xdc\x59\xdf\x07\x0a\xd3\xf9\x4a\xec\x33\x97\xb4\xa7\xeb\x33\x77\xe7\x22\x85\xa4\xac\x90\x38\xe3\xa0\x81\x37\x40\xad\xdc\x95\x88\x7c\x36\x20\x0c\x3c\x77\x29\xd2\x4e\xca\x0a\xf1\xb2\xb2\x76\xfd\x06\xd2\x83\xa9\x8c\x7d\x21\x40\xad\x64\x86\x92\xb2\x42\xee\x44\x52\x78\xde\x6f\x32\x91\x39\x4a\xca\x12\x55\xa2\x41\xaf\x84\x7e\x0b\x81\xa5\x50\x90\x58\x1a\xf4\xac\x99\x0e\xac\x99\x68\xac\x11\xad\xdb\x13\x54\xc3\x97\x39\x4f\xca\xf2\x54\x44\x3d\x9a\xc9\x5c\x25\x65\x85\xab\x81\x6c\xb2\xbc\xe3\x6c\xa0\x6c\xaa\x50\xe6\x0b\xaa\x33\x53\x14\x75\xa6\x6a\xaa\xec\xb2\x04\x05\x9a\x29\xba\x3a\x53\x95\x75\xd2\x73\x69\x3e\xd0\x32\xd3\xb8\xb4\x94\xfd\x48\xd7\x45\x54\xe9\x79\xa0\xf2\x7f\x22\xbb\x9b\xae\x93\xa2\xd3\x73\x4d\xa7\x05\x8d\x5a\x28\xfa\xba\x50\xf5\x55\x5a\x56\x04\x95\x5a\x0c\x84\x2d\x14\x85\xf5\x05\x0d\x5a\x28\x1a\xbb\x50\x35\x76\xda\xf3\x67\xe9\x89\x50\x0a\x7f\xe6\xb2\x87\xec\xba\x28\x7a\xbd\x54\xf5\x5a\x5c\x90\x02\x41\xa1\x96\xa2\x62\x2f\x15\xc5\x96\x97\xa4\x40\x50\xa7\x95\xa2\xb4\x2b\x55\x69\x03\x41\xa9\x56\x8a\xd2\xae\x34\xa5\x35\xae\xfb\xe2\x5a\x1c\x88\xae\x8a\xa4\x80\x58\x03\x0f\x5d\x8c\xde\xbe\xf3\xee\x5a\x30\x16\x3d\xee\xd2\x16\x53\xdf\x4e\x76\xfd\x74\x13\xa7\x78\xf6\x1c\xb6\xca\x19\x66\xce\x9d\xd5\xc4\x98\x44\x96\x4a\xa0\xd9\x6f\xef\xf6\xb0\x8f\x71\x92\x1a\x6d\x59\xf2\x46\xaa\x66\x41\x31\x6f\x20\x71\x92\xa1\xb6\xd1\x2b\x75\x28\x3d\x8e\x12\x62\x44\x35\x02\x13\x9a\x8c\xb5\x12\x03\x0d\xc1\x6d\x1f\x3f\x44\x7d\x36\x44\x62\xb9\xb8\x93\x13\x73\x06\x3c\xeb\x40\xe2\x86\x3e\x3a\x22\x51\x11\x09\x08\x78\xa2\x67\x3a\x9d\xf2\xcf\x7a\xbb\x46\x57\x9e\x4d\x7e\xee\xf2\x3a\xd2\x62\xae\x37\xab\xd5\xaa\x1f\xdf\x12\xa0\x3d\x00\x56\xd3\x5e\x6f\x3e\xcc\xc9\x8f\x09\x6d\x08\xf5\x18\xf3\xf8\xc4\x49\x98\xbd\x6f\x42\x88\xdd\xa4\xd9\x38\x34\x37\x66\x4b\x2c\x31\x54\x51\xfe\x1c\x36\x69\x06\xa2\x0f\x49\xfe\x2f\x4d\xc2\x1f\xff\xeb\x63\x8e\xb6\xf8\x8f\x2e\xdb\xe1\xfe\x9a\xc6\x75\xd9\x94\x9b\xd6\xdd\x92\xd1\x40\x6d\xae\xe8\x86\xfe\x07\x42\x64\xd3\xd6\xef\xbe\x25\x99\x2e\xfa\xe7\x5b\xdb\x82\x8d\x83\xd0\xe0\xa3\xbe\xe1\x67\xde\xf9\x0f\x12\xd0\xca\xd3\x4d\xc9\x80\x7d\x24\xd3\x8f\x42\x14\x1d\xd5\xb6\x67\x01\x7b\x2c\x22\x00\x7b\x53\x97\xf9\x95\x90\x0d\xbc\xb6\x29\xa7\x20\xc0\x28\xab\x2b\x58\xaa\x6c\x91\xff\xde\xec\xfa\xda\x6e\xcb\x2b\xb1\xce\xbf\xbe\x3e\x31\x32\x1b\x70\x20\x40\x18\x49\xc5\x6c\xa9\xa3\xf9\x46\xc4\x44\x09\x35\xac\x65\xa5\x61\xf3\x2e\xc0\x55\xfe\xe3\xe9\xfb\x8b\x08\x8f\x06\x65\x1a\xd2\xb1\x92\x3e\xf2\xa4\xad\xd4\x81\xd5\xb1\x14\xd1\x0e\x25\x10\x82\x82\xa8\xe1\x47\xcc\x45\x1a\x09\x46\x4f\x8b\x06\xb7\x44\x15\xac\xb9\xd2\x1a\xf0\xc6\x68\x48\x28\xf0\xe4\xeb\x9f\x2b\x89\x06\x21\x7b\x26\x91\xd6\x55\x9b\x6b\x0d\x66\x33\xb4\x89\xf1\xab\xda\xd6\xcf\x8e\xe7\x4b\x34\xe6\x53\x33\x7e\xb9\xed\xe1\x82\xe2\x7f\xb7\x41\x59\x83\xaf\x3b\x04\x28\xab\x76\xe8\xaa\x24\xbb\xfb\xf6\xf9\xdd\x14\xec\xcb\xb9\x27\x39\x42\x87\x57\x85\xee\x94\x79\x54\xa1\x2c\x7c\xf2\x55\xa0\x28\xc9\xea\x90\x95\x8f\x38\x89\x04\xb1\xd0\x44\x92\xec\x5e\x68\xe6\x45\x4e\xdc\xff\x35\xff\xeb\x54\x75\x4a\x4e\x28\x64\x6e\xb2\xe5\x02\x64\x83\x8c\xb0\xc8\x08\xac\x1f\x20\x80\x36\x2c\x96\x78\xd1\xf9\xeb\xcd\xa6\x8b\xd0\x09\x93\xc5\xec\x2d\x71\xb5\x9d\x7a\xa8\x6d\x15\xc4\xfe\x24\x86\x1a\x6b\xc7\x39\x4a\xb3\xb1\xc6\x7d\x3d\xda\x94\xa0\x16\x8f\xb5\xe5\xb0\x00\xed\xc6\x1a\xdb\x34\x1f\xed\x48\x90\x5e\xd2\xee\xc0\xce\x02\x8d\x92\xf6\x88\xf1\xfd\x58\x1b\x4f\xc0\x8e\xb4\xf2\x8c\xe5\x18\xe5\x78\x74\x48\x2a\x22\xa9\x91\xa7\xf5\x85\x9a\x3e\xe1\x28\x2d\xc8\xee\x1c\x56\x64\xe3\xc6\x49\x5f\x86\xe3\x38\x16\x15\x9c\x3a\x10\xee\x79\x26\xf0\xf7\x4d\x92\x24\xca\x0a\x3b\xad\x9e\x46\xf6\x54\x5f\x65\x1b\x33\xb0\x85\x9d\xaa\x11\xde\xdc\x5d\xff\xb3\x27\xfc\x95\xa7\x26\x4a\x7c\x98\x5a\x40\xa6\x36\xeb\x73\x7f\x46\x13\x95\x3c\xef\xa8\xa1\x9e\x84\x62\xe6\x7a\x12\x84\x18\xed\x49\x00\x6a\xba\x27\x21\x98\x01\x9f\x04\xa1\x66\x7a\x76\x98\xcb\xa1\xb8\x49\x9f\x84\xa5\x86\x7d\x12\xa2\x3b\x5f\x39\x05\xd3\x1d\x4b\x9c\x9c\x1d\x3e\x43\x0a\x13\xbf\x0e\xc2\x23\x05\xad\xbe\x33\x7e\x75\xbf\xa1\xac\xff\x7e\xb0\xfa\xe9\xc3\x77\x27\xad\xe9\xcb\x30\x30\xaa\x61\xd9\x3d\x23\x0e\x76\x00\x76\x7a\xea\xdd\xc1\x99\x42\x49\xbb\x4b\x8b\xce\x70\x19\x15\x51\xd7\x44\x0c\x96\xa4\xe4\x2c\x03\x79\xec\xb3\x43\x6a\x6b\x2d\x94\xa2\xee\x70\x97\x7a\x10\xcf\x78\xfc\x61\xb4\x38\x6d\xf7\x34\x6a\x75\x67\x21\x99\xe5\x9d\x05\x23\xd6\x77\x16\x88\x5a\xe0\x59\x28\x66\x85\x67\xc1\xa8\x8d\x5d\x34\xe4\xcb\x20\xb9\x45\x9e\x85\xa7\x56\x79\x16\x8a\x5b\xe6\x59\x38\x6e\x9d\xe7\x67\x8d\x2f\x20\x8d\x59\xa9\x19\x8c\x59\xaa\xb9\xad\xb3\x56\xfd\x94\x5a\xda\x6b\x6a\x1b\x36\x8c\x70\x02\x0d\xbc\x14\xa3\x24\x48\x26\x27\x8d\xf9\xab\x0c\xf0\x09\x28\x4f\xc8\x81\x95\x69\xbe\xc6\xb6\x7e\xbe\x7d\xab\x61\x37\x8a\x31\xee\x06\x5e\x2c\x16\x8a\xdb\x81\x95\x5a\x9f\x28\xf5\x0f\xb0\x6a\x3f\xc0\xce\x23\x19\xf7\x86\x26\x08\xd1\x8f\x76\xed\xfc\x6a\xc4\x9b\xf5\x6a\x8a\xa6\x4b\x85\x00\xbc\x9a\x04\x41\x72\xde\xef\x75\xc8\x2e\xf6\x82\x17\x75\x50\x7c\xa2\xdc\xa7\xf3\x90\xe3\xa4\xf2\x0b\x1b\x3c\x3d\x16\xd0\x2b\x30\x63\xdb\x21\x4d\x2e\x4a\x5c\xc0\xe5\x9c\xef\xb3\x36\xad\xc8\xa5\x22\xf1\x42\x81\x00\x97\xa1\x35\xce\x64\xbf\x6a\x91\x5d\x8c\x08\xd3\x1f\xc3\x77\xc7\xed\xca\x51\xbb\xc7\x4e\xd9\xd5\xd8\x88\xe1\x67\x27\xef\x92\xab\x8e\x84\xe3\xec\x01\x13\x75\xe9\x6c\x84\xee\xea\x13\xa9\xec\x94\x6d\x32\x99\x44\xfa\xc5\x28\xce\x12\x3c\x23\x3f\x61\x5c\x72\x6c\x1b\xdf\xe3\xe4\x54\x7c\x64\x82\x31\xc5\x49\x26\x38\x2d\x5e\x32\x01\x29\x71\x93\x09\x44\x8d\x9f\x4c\x30\x5a\x1c\x65\x9c\x9c\x12\x2f\x8d\x0d\xf6\x12\x38\x3d\xbe\x32\x41\xab\x71\x96\x09\x46\x8f\xb7\x4c\x50\x7a\xdc\x65\x16\xe3\x59\x92\xb4\x38\xcc\x04\x44\xbd\x81\x09\x46\x8b\xd9\xfa\x16\x6a\x2c\xa6\x86\x3e\xa8\x93\xf5\xbc\x33\x2b\x7a\x13\xc5\x1b\x55\x4f\x29\x4a\xba\x14\x03\x0d\xe3\x3a\x0c\xb6\xa1\xa1\xa3\xc9\xd4\xa6\x4f\xb1\x6b\x61\x55\x3b\x9c\x55\x0e\x8b\xca\x44\xa0\x1c\x37\x0d\xda\x62\xde\x72\xd9\xb1\xee\x0d\xcf\xfc\x1b\xe3\x4f\xd3\x84\x7a\x5e\xea\xe7\xbb\x26\x70\xbe\x71\x64\x17\x37\x1d\x76\x69\x47\xf1\x22\x72\x3c\x76\xaa\x27\x77\x87\xc2\x49\x03\xbd\x9d\x6b\x0e\x54\xcd\x11\x75\xe7\xdb\x86\xeb\x96\x1e\x49\xb7\x93\xe4\xe0\x59\x22\x9a\xfe\xf6\x22\xf7\xc4\xd0\xd5\x37\x04\x93\xac\x4c\x3f\x1d\xba\x08\xe0\x44\xdf\x31\x4b\xcd\x07\x25\xe5\x0e\xd4\xa9\x07\x02\xfa\x6e\x9c\xb1\x44\xf1\xfe\xbd\xf7\xf5\xaa\xa7\xb1\x2e\xaa\x4e\x4a\x8d\x23\xb6\xd2\x91\x43\xf0\x0e\x9c\xf3\x2c\x07\xdc\xbc\x12\xdd\x9b\x6e\x98\xd2\x5b\xa4\xfe\x19\x92\xf4\xc5\xdb\x44\x18\x5f\xae\x3f\x03\x07\x13\xfc\x14\x4e\xce\xa0\x4c\xeb\xa6\x85\xc0\x25\xcd\x0c\x32\x50\x11\x0f\xb0\xf4\xda\xab\x3e\xb5\x29\xac\x69\xe4\xaf\xd7\xaf\x85\xf2\x72\x7a\x72\xfc\x30\x43\x5f\x42\x8a\xd0\x6d\x9c\xaa\x0b\xa9\x79\x01\x05\xca\xa8\x4e\xa0\xcb\x19\xd8\xc0\x19\x72\x76\x78\x7e\x10\xd6\xf9\xcb\x89\xea\x2f\x25\xa3\xf0\xf5\xb3\x4b\x0d\x26\x70\x26\xfd\xc1\xf5\x28\x90\xef\x04\xc3\x89\xfe\x38\xd0\xa4\xbf\xef\x70\x02\x48\xb8\x72\xa1\x03\xfd\x05\x87\xcc\x0d\xcb\xc9\xf0\xa6\x95\x76\x54\xf3\xf9\x7c\xc4\x87\x0d\x47\x87\xee\x72\xa1\x78\x4f\x3e\xa4\x62\xbd\x02\x4a\xad\xf3\xbf\xe5\x38\x49\x91\x45\xef\xce\x35\x71\x8d\x71\x41\xdf\x35\x5c\xf5\x17\x1b\xad\x70\xba\x04\xab\xbf\x3e\x08\xd3\xe6\xc7\x1e\xd2\x8d\xe3\x4e\xb8\x0b\xcd\x9d\xea\xeb\xe8\x68\xbc\xf8\xaf\x1c\xf9\xdf\x2d\x47\x2e\x86\x20\xd2\xce\xdf\x78\xbd\x55\x70\x36\xba\xb8\x85\x56\x5d\xde\x5a\xa3\x2a\x70\x0d\x40\x92\xb8\xd6\x2a\x88\x5c\x6b\x13\x65\xae\x35\x4a\x42\xd7\x49\x16\xa4\x6a\xc4\x7b\x11\x80\x2c\x77\x0d\x4c\x14\xbc\xd6\x28\x4b\x5e\x6b\x96\x45\x6f\x60\xf9\xf8\xb0\x3c\x5d\x2e\xcb\xfa\xfc\xbe\x48\xda\x9c\x19\x14\x45\xd8\xaf\x11\x2f\x17\x8d\x45\x9b\x5f\xb0\x01\xfb\x42\x57\x6c\x68\x51\x3c\xa6\xe0\x26\x69\xf0\x2c\x1d\x21\x40\x44\xbd\x04\xcf\xc9\x07\x27\x0f\xa2\x0e\x5f\xe5\x38\x83\x60\x76\x36\xe9\x13\x6c\x10\xfb\x2d\x15\x2d\x46\x74\xb5\xf0\xe8\xb6\xca\x8b\xd4\xad\x10\xed\x96\xa5\x4d\xb7\xc7\xa3\xe5\xb4\xc5\xf9\x41\xdb\x98\xa9\x3d\x0e\xe4\x3f\xfc\x99\x12\x3d\x44\xd6\x72\x0a\xde\x51\x43\xda\x9f\xbe\xf6\xc0\x3c\x9b\x21\x08\x95\x8f\x50\xdc\x8b\x34\x91\xd7\x65\xd0\x53\x61\xbd\xf1\x4e\x8a\x7e\x0d\x49\x44\xbc\x2b\xeb\xf4\x33\x48\x0c\x65\xc2\xd6\xe1\xe2\x3e\x96\xca\x04\xd3\x42\x7d\xbe\x2b\x61\x86\x7d\x16\x8a\xcf\xf9\x3c\x60\x83\x2b\x04\x1c\x28\xeb\x17\xdd\xcf\x3d\x15\xc7\xf5\x54\xaa\x54\x1f\x4c\xae\x9c\xb6\xd2\x0d\x60\xad\x1c\xc7\x0f\xfb\x7b\xb4\x6e\xca\x6c\xdf\x62\xa6\x8f\x94\xe9\x4c\x25\x75\xb5\x31\x6b\xa9\x79\xea\xfd\xa8\x82\x9a\xd3\x17\x15\x4c\xb8\xfc\x0a\x95\x88\x87\x26\x63\xd9\xcd\x24\x76\xab\xe2\xbd\x01\x9d\xc8\x73\x76\x85\xc2\x04\xa5\x06\x3e\xea\x5c\x25\xf2\x51\xd3\xf7\x7b\xaf\x28\x7a\x88\x36\xed\x70\x37\x81\xef\xef\x66\x74\x7f\xc7\xde\x2f\xbc\xfe\x33\x98\x7d\xbf\x7c\x2d\xb8\x1a\xfa\xa2\xef\x3c\x7b\x2e\x18\x55\x1a\xe3\xc3\x6b\x11\x27\xec\xf1\xc8\x73\x18\xf2\x4e\xa6\x7b\x54\xe3\x3c\x87\xac\x36\xea\x6b\x9e\xba\x77\x7b\xc6\x9e\xa3\x16\x73\xc2\x54\x5e\x86\xe8\xbc\xe9\x99\xa7\x64\xb8\xae\x28\xcc\x92\xcd\x49\x9c\x25\x55\x2b\x72\xa7\xb0\xaf\x13\x9c\x5f\xff\xfe\xab\x6f\xa3\x03\x11\x7d\x6e\xcb\x7d\xbc\x93\x83\x74\xef\x05\x74\x0a\xcf\xad\x68\xdd\x1a\xd5\xa6\x1b\x2f\x8a\x23\xd0\xd3\xca\x24\xd9\xdc\x39\xdb\x21\x0c\x63\xd9\xda\xf3\x8a\x34\x60\xe6\x5e\xb3\xea\x9f\x56\xfa\x74\xe5\xee\x03\x7d\x77\x78\x30\x23\x79\x6e\xe5\xf5\xe2\xbe\xaa\x70\x1d\xa3\xa6\x3f\x83\x98\xcd\x67\xc9\x7c\xaa\xae\x01\x87\xe1\x84\xc2\xe8\x6c\x4e\xa5\xcf\x8d\x0b\x89\x72\xcf\x4a\x22\x71\x24\x91\x22\xf5\x3b\xf4\xf7\x95\x66\xa6\x66\x4b\x35\x32\xf6\x3c\xf8\xe4\x63\xe2\x93\xbe\x46\xa5\x7f\x40\xaa\xd5\xb2\x6c\x87\xf1\xa0\x47\x56\x12\x76\x41\x49\x25\xd5\xbe\x00\x86\xbe\x4c\x25\x97\x9b\x84\xc7\xc5\x0c\xe4\x85\xcf\xe9\x22\x9c\x57\xed\xb3\x13\xe3\x2c\x6b\xc2\x66\x57\x3e\x9a\x8e\x47\xd6\xe4\x27\xe2\xb7\xf8\xd3\x73\x61\x78\xfe\xa2\x96\xbe\x93\xb6\x96\xb3\xb7\xb7\xbe\x85\xd4\x57\x6a\x7d\x46\x8a\x9e\x4c\x68\xb7\x8f\xa5\x11\xda\x4e\x33\x78\x71\xd7\xcd\x89\xad\x5b\x2a\x75\xdd\xc4\x98\x55\xf4\x37\x07\x05\x4f\xad\x3e\xf6\x8d\xd4\x57\xbe\x63\xf9\xbb\x8e\x1e\x43\x62\xaa\xa3\x4d\xca\x43\x09\x74\x76\xe4\xc8\x98\x88\x9a\x9b\xd4\xc3\x23\x3f\xe1\xb9\xb8\xb6\x0b\x57\x1f\x7d\xd3\x0d\xbb\x42\xe5\x25\x2a\x4e\x61\x9d\x32\x49\x8c\xf0\x6f\x36\x01\xf9\x49\xb0\x4d\x5b\xa7\x15\x49\x25\xd7\x21\x44\x5a\x6c\xa2\x57\x41\xe1\xf8\xd7\x17\xa3\x60\x6c\xc1\x6c\xcc\xb1\x63\x27\x5d\xd5\x84\x7e\xe4\x7f\x8b\xf0\x1e\x28\x18\xf2\x5c\xef\x55\x5c\x26\x86\x8b\x9e\x54\xd6\x2a\xa9\xa5\x57\xb0\x41\x85\x84\xbc\xdd\x45\xa4\x8a\xe8\x5e\x42\xec\xff\x07\x00\x00\xff\xff\x6f\x05\x93\x99\x86\x43\x00\x00")

func staticCssPureCssBytes() ([]byte, error) {
	return bindataRead(
		_staticCssPureCss,
		"static/css/pure.css",
	)
}

func staticCssPureCss() (*asset, error) {
	bytes, err := staticCssPureCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/css/pure.css", size: 17286, mode: os.FileMode(420), modTime: time.Unix(1424723083, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"template/error.tmpl": templateErrorTmpl,
	"template/login.tmpl": templateLoginTmpl,
	"template/paste_form.tmpl": templatePaste_formTmpl,
	"static/css/pure.css": staticCssPureCss,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"static": &bintree{nil, map[string]*bintree{
		"css": &bintree{nil, map[string]*bintree{
			"pure.css": &bintree{staticCssPureCss, map[string]*bintree{
			}},
		}},
	}},
	"template": &bintree{nil, map[string]*bintree{
		"error.tmpl": &bintree{templateErrorTmpl, map[string]*bintree{
		}},
		"login.tmpl": &bintree{templateLoginTmpl, map[string]*bintree{
		}},
		"paste_form.tmpl": &bintree{templatePaste_formTmpl, map[string]*bintree{
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
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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
                err = RestoreAssets(dir, path.Join(name, child))
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

