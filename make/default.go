// Code generated by go-bindata.
// sources:
// make/default.yaml
// DO NOT EDIT!

package make

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

var _makeDefaultYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x56\x6d\x6f\xdb\x36\x17\xfd\xee\x5f\x71\xa0\x47\x80\xdb\x06\x94\xda\x67\xc3\x3e\x68\x70\x01\xc3\xf5\xdc\x62\x5d\x6c\x38\x49\x31\x60\x1b\x5c\x46\xa2\x25\xc2\x12\xa9\x92\x94\x9d\xc2\xf1\x7f\x1f\x48\xc9\x8e\x5f\xe4\xbc\xa0\xc0\xbe\x24\x22\x79\xef\xb9\xe7\x9e\x73\x45\x79\xbd\xf6\x4b\x6a\x32\x44\x3d\x78\xec\xae\x94\xca\x60\xd2\xbf\xfe\xd8\xf3\x47\x63\xfb\x3f\xbc\xe5\x22\xf2\xed\x93\xb7\xd9\x74\xd6\x6b\x3f\xe7\xda\xb8\x60\xff\x55\x2a\xe1\x56\x41\x18\x04\x01\xee\x91\x2a\x56\x82\x2c\xb1\x64\x22\x91\x2a\xdc\xdb\xe9\xce\xc2\x54\x76\x5f\x5b\x08\x26\x96\x3a\xea\x00\x04\xfd\xc9\x24\x82\x92\x62\xbb\x98\xdd\x5c\x7e\x18\x4e\xaf\x06\xe3\xe9\x30\xc2\x05\x8b\x33\x09\xbf\x3f\x99\xe0\x1e\x46\xc1\x23\x1e\xbc\x99\x57\xc7\x4e\x07\x1f\x23\xd0\x22\xf9\xe5\x67\xb7\x1e\xdc\x4c\x3f\x7c\x9a\x46\xb8\x28\x57\x89\xdb\xb8\xb9\xec\xff\x61\x31\x2a\x41\x0b\x56\xe7\x77\xff\x8a\xaa\xb2\x64\x2a\xfa\xa7\x6b\x9f\x73\xb9\x72\xcf\x2e\x7c\xd2\x1f\xfc\xde\x1f\x0d\x67\x5f\x86\xd3\xab\x4f\xe3\xcb\x08\x17\x29\x37\x48\x98\x8e\x15\xbf\x65\x20\x84\xe6\x2b\xfa\x5d\x83\x90\x84\x2b\xf3\x1d\x84\x18\x9a\xea\x06\x97\x74\xd1\x9d\xd5\x38\xd7\xfd\x51\x84\xe5\xdb\xe0\x6d\xf0\xce\xad\x3f\xf7\xaf\x87\x57\xd7\xb3\x9b\xe9\xe7\x08\x5e\x66\x4c\xa9\xa3\x30\x4c\xb9\xc9\xaa\xdb\x20\x96\x45\x58\x95\x9a\xa7\x99\x09\x95\x14\xa1\x62\x39\xa3\x9a\xe9\x30\x91\x2b\x91\x4b\x9a\x84\xfe\x75\x7f\x64\x8f\x88\xbf\x76\xfd\x6c\x88\xdd\xf1\x3a\x86\xaa\x94\x19\xa7\xe1\x6d\xc5\xf3\xc4\x3e\xa0\x61\x5b\x1a\x2e\x45\x84\x81\x2c\x4a\x9e\x33\x50\xdc\x72\x41\xd5\x77\x18\x89\xc0\x3a\x19\x6e\xb1\x66\xfe\xda\xaa\xb8\x71\xb9\xb7\x6c\x2e\x15\xab\x71\x2c\xef\x52\xb1\xd2\x2d\xe2\x22\x89\x70\xdf\xec\x17\x8b\x84\x2b\x90\x12\x67\x71\x80\xd1\xd8\xae\x7a\xbe\xfd\x8b\xd1\x78\x7c\xd5\xf3\x5d\x1c\x52\x59\x93\x05\x91\xad\xf9\xa1\x75\xda\x35\x08\x92\x27\xf3\xdc\xea\xeb\x91\x3f\xe1\x4f\x87\x93\x71\x30\xe2\x66\x20\x8b\x82\x9b\x9e\x7f\xe4\x15\x76\x31\xfd\xb2\xfc\xc2\x94\xe6\x52\xf4\x6a\x98\xbd\x83\x4b\x5a\xb0\x9e\xad\xe0\xd9\x96\x5c\xad\xf0\x4d\x90\xca\xad\x84\x33\x9a\xe7\x2f\x92\x31\xe7\xa2\xba\xdb\x72\x07\x15\x49\xb3\x9f\x50\xb5\xe2\xe2\x79\xe2\xd2\xb9\x61\xea\x61\x7f\x56\x33\x29\x74\x7a\x2a\xfd\x5c\x2a\x48\x0d\x2e\xe0\xb9\xca\x1e\xbc\xba\x92\x07\x6f\xc5\x45\x22\x57\xda\xfb\x15\x89\x6c\xc2\x81\x52\x71\x61\xe6\xf0\x1c\x26\x17\xa9\x65\x5e\xcb\x2e\xf5\x4e\xf3\x20\x08\xfe\x16\xde\x2e\xe7\xc8\xdf\xbd\xc0\x5d\x48\x8b\xbd\x52\xb7\x78\xbb\x97\xfb\x9f\x1a\xdb\x78\xea\x8c\x94\x82\x75\x80\x58\x2e\xb7\x22\x1f\x58\x3b\xad\x84\xe5\xed\x8e\x69\xca\x9c\x85\xb2\x64\xc2\x7a\xad\xe4\x4a\x33\x65\x25\xd3\x99\x5c\xc1\x64\x0c\x19\xa3\xa6\xa0\x65\xf0\x82\xd7\x45\x15\x4e\x8c\x6d\x85\x37\xb8\xb7\xf7\x45\xc5\x9a\x63\x77\xbd\x79\x85\x4c\x58\x84\x58\x56\xc2\x78\x78\x7f\x90\x10\xc8\xca\x34\xb1\xff\x83\xaa\x04\x0c\xd3\x46\x3b\xa2\xb1\x62\xd4\xb0\x07\xf2\xa5\x92\x73\x9e\xb3\xbd\x61\x29\x17\xa9\x9d\x96\xf3\x57\x74\xd8\xdc\xd1\xaf\x0f\xc6\x26\x95\xae\x0a\x88\x83\xb6\xe4\x7a\x8e\x5b\xb3\xd1\xd4\xe9\x1d\xd3\x0c\x4c\x51\xc2\x2f\x17\xe9\x0e\xc8\x50\x9e\x83\x08\x5c\xfc\xff\xa4\x27\x17\xfc\xfe\xb4\xd7\x23\x7d\x4e\x04\xdc\xe5\x1e\xc6\x39\x9f\x77\x3a\xc5\x19\x8b\x17\x5b\xb7\x3a\x0f\x4d\x49\x99\xd7\x72\x81\x64\xa6\xc8\x4f\x3a\xe8\x00\x39\x17\xe6\xec\xa0\xd8\xc3\x17\x78\xff\x43\x16\xd8\x5a\xfb\x6a\x36\x93\x6c\x8d\x39\x3f\xc8\xf6\xf4\x89\xe9\x3c\xba\x6f\x76\x2d\xd9\xc5\x92\x99\xd3\x2e\x0e\xa7\x01\xcd\xb7\x7f\x63\xaf\x83\x25\x7b\x84\xcb\x92\x19\x48\x81\x20\xd4\x2a\x7e\xc9\x1b\xd3\xa4\xee\xd7\x49\x64\xbc\x60\x6a\x16\xe7\x8c\x8a\x99\x6c\xfd\xd6\x4d\x59\x21\x97\x0c\x34\xb7\x0e\x0b\x43\xb9\x60\xaa\x7e\x4d\x78\x41\x53\xa6\x61\x32\x6a\x40\x15\x83\xeb\x82\x62\xc5\xd8\x02\x32\x4f\x82\x53\x06\x75\x39\x94\x1a\x84\x6e\x9d\xea\xda\x78\x0d\x9a\xca\x2e\xee\x41\x57\x0b\x74\xd7\xee\x7a\x85\xff\x6e\x63\xb7\xee\xa8\x4a\xf5\x36\x55\x15\x27\xf3\xd9\x02\x59\x48\x61\xb2\x1f\xc6\x6c\xfa\x7b\x9a\xe7\x4f\x6d\x98\xfc\x59\xa0\x8f\x30\x7d\x12\x75\xdf\xbc\xc7\x8d\xfb\x50\xa7\x1f\xf9\x97\x50\x91\xe6\xf6\xe3\x55\x73\x6a\x1f\xa5\x83\x09\xa9\x03\xcf\xfa\xaa\x0a\x90\x39\xfc\x57\x07\x9e\x90\x6f\xaf\xcf\x70\x9e\xb1\x3b\x6e\xd8\x23\x33\x57\x9f\xef\xd1\x7e\xe6\x48\x0d\xef\xb8\xc1\x3d\xe2\xca\x80\x24\xe8\xa2\x6b\x69\xbd\x3b\x15\xf3\x11\x88\x81\xfb\x04\x24\xcf\x44\x69\x11\xe9\x7c\x57\xad\xb2\xb7\xcb\xc9\x1f\xc4\x6c\xe6\x86\x7c\x03\x21\x73\x9e\x1b\xa6\xec\x6f\x94\x1a\xa9\x67\xc5\xf5\x5a\x64\xd6\x86\x9a\x36\x26\x57\xdb\xef\xae\x0b\x80\x9c\xdb\x0f\xa0\xb0\x9c\x9e\xa3\x75\x9d\xf4\xf5\x41\xb6\xa3\xc1\xbd\xfc\xcd\x4d\xee\xf6\x1e\xb6\xbf\x43\xaf\xbe\x76\xfe\x0d\x00\x00\xff\xff\xd6\x8b\x0a\x11\x08\x0d\x00\x00")

func makeDefaultYamlBytes() ([]byte, error) {
	return bindataRead(
		_makeDefaultYaml,
		"make/default.yaml",
	)
}

func makeDefaultYaml() (*asset, error) {
	bytes, err := makeDefaultYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "make/default.yaml", size: 3336, mode: os.FileMode(493), modTime: time.Unix(1472348142, 0)}
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
	"make/default.yaml": makeDefaultYaml,
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
	"make": &bintree{nil, map[string]*bintree{
		"default.yaml": &bintree{makeDefaultYaml, map[string]*bintree{}},
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

