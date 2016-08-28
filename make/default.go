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

var _makeDefaultYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x56\x6d\x6f\xdb\x36\x17\xfd\xee\x5f\x71\xa0\x47\x80\xdb\x06\x94\xda\x67\xc3\x3e\x68\x70\x01\xc3\xf5\xdc\x62\x5d\x6c\x38\x49\x31\x60\x1b\x5c\x46\xa2\x25\xc2\x12\xa9\x92\x94\x9d\xc2\xf1\x7f\x1f\x48\xc9\x8e\x5f\xe4\xbc\xa0\xc0\xbe\x24\x22\x79\xef\xb9\xe7\x9e\x73\x45\x79\xbd\xf6\x4b\x6a\x32\x44\x3d\x78\xec\xae\x94\xca\x60\xd2\xbf\xfe\xd8\xf3\x47\x63\xfb\x3f\xbc\xe5\x22\xf2\xed\x93\xb7\xd9\x74\xd6\x6b\x3f\xe7\xda\xb8\x60\xff\x55\x2a\xe1\x56\x41\x18\x04\x01\xee\x91\x2a\x56\x82\x2c\xb1\x64\x22\x91\x2a\xdc\xdb\xe9\xce\xc2\x54\x76\x5f\x5b\x08\x26\x96\x3a\xea\x00\x04\xfd\xc9\x24\x82\x92\x62\xbb\x98\xdd\x5c\x7e\x18\x4e\xaf\x06\xe3\xe9\x30\xc2\x05\x8b\x33\x09\xbf\x3f\x99\xe0\x1e\x46\xc1\x23\x1e\xbc\x99\x57\xc7\x4e\x07\x1f\x23\xd0\x22\xf9\xe5\x67\xb7\x1e\xdc\x4c\x3f\x7c\x9a\x46\xb8\x28\x57\x89\xdb\xb8\xb9\xec\xff\x61\x31\x2a\x41\x0b\x56\xe7\x77\xff\x8a\xaa\xb2\x64\x2a\xfa\xa7\x6b\x9f\x73\xb9\x72\xcf\x2e\x7c\xd2\x1f\xfc\xde\x1f\x0d\x67\x5f\x86\xd3\xab\x4f\xe3\xcb\x08\x17\x29\x37\x48\x98\x8e\x15\xbf\x65\x20\x84\xe6\x2b\xfa\x5d\x83\x90\x84\x2b\xf3\x1d\x84\x18\x9a\xea\x06\x97\x74\xd1\x9d\xd5\x38\xd7\xfd\x51\x84\xe5\xdb\xe0\x6d\xf0\xce\xad\x3f\xf7\xaf\x87\x57\xd7\xb3\x9b\xe9\xe7\x08\x5e\x66\x4c\xa9\xa3\x30\x4c\xb9\xc9\xaa\xdb\x20\x96\x45\x58\x95\x9a\xa7\x99\x09\x95\x14\xa1\x62\x39\xa3\x9a\xe9\x30\x91\x2b\x91\x4b\x9a\x84\xfe\x75\x7f\x64\x8f\x88\xbf\x76\xfd\x6c\x48\x4e\x0d\xd3\xc6\xeb\x18\xaa\x52\x66\x9c\x8a\xb7\x15\xcf\x13\xfb\x80\x86\x6f\x69\xb8\x14\x11\x06\xb2\x28\x79\xce\x40\x71\xcb\x05\x55\xdf\x61\x24\x02\xeb\x65\xb8\x45\x9b\xf9\x6b\xab\xe3\xc6\xe5\xde\xb2\xb9\x54\xac\xc6\xb1\xcc\x4b\xc5\x4a\xb7\x88\x8b\x24\xc2\x7d\xb3\x5f\x2c\x12\xae\x40\x4a\x9c\xc5\x01\x46\x63\xbb\xea\xf9\xf6\x2f\x46\xe3\xf1\x55\xcf\x77\x71\x48\x65\x4d\x16\x44\xb6\xe6\x87\xd6\x6b\x62\x9b\x06\xc9\x93\x79\x6e\x15\xf6\xc8\x9f\xf0\xa7\xc3\xc9\x38\x18\x71\x33\x90\x45\xc1\x4d\xcf\x3f\x72\x0b\xbb\x98\x7e\x59\x7e\x61\x4a\x73\x29\x7a\x35\xcc\xde\xc1\x25\x2d\x58\xcf\x56\xf0\x6c\x4b\xae\x56\xf8\x26\x48\xe5\x56\xc2\x19\xcd\xf3\x17\xc9\x98\x73\x51\xdd\x6d\xb9\x83\x8a\xa4\xd9\x4f\xa8\x5a\x71\xf1\x3c\x71\xe9\xdc\x30\xf5\xb0\x3f\xab\x99\x14\x3a\x3d\x95\x7e\x2e\x15\xa4\x06\x17\xf0\x5c\x65\x0f\x5e\x5d\xc9\x83\xb7\xe2\x22\x91\x2b\xed\xfd\x8a\x44\x36\xe1\x40\xa9\xb8\x30\x73\x78\x0e\x93\x8b\xd4\x32\xaf\x65\x97\x7a\xa7\x79\x10\x04\x7f\x0b\x6f\x97\x73\xe4\xef\x5e\xe0\x2e\xa4\xc5\x5e\xa9\x5b\xbc\xdd\xcb\xfd\x4f\x8d\x6d\x3c\x75\x46\x4a\xc1\x3a\x40\x2c\x97\x5b\x91\x0f\xac\x9d\x56\xc2\xf2\x76\xc7\x34\x65\xce\x42\x59\x32\x61\xbd\x56\x72\xa5\x99\xb2\x92\xe9\x4c\xae\x60\x32\x86\x8c\x51\x53\xd0\x32\x78\xc1\xeb\xa2\x0a\x27\xc6\xb6\xc2\x1b\xdc\xdb\x1b\xa3\x62\xcd\xb1\xbb\xe0\xbc\x42\x26\x2c\x42\x2c\x2b\x61\x3c\xbc\x3f\x48\x08\x64\x65\x9a\xd8\xff\x41\x55\x02\xf6\xed\xd7\x8e\x68\xac\x18\x35\xec\x81\x7c\xa9\xe4\x9c\xe7\x6c\x6f\x58\xca\x45\x6a\xa7\xe5\xfc\x25\x1d\x36\xb7\xf4\xeb\x83\xb1\x49\xa5\xab\x02\xe2\xa0\x2d\xb9\x9e\xe3\xd6\x6c\x34\x75\x7a\xc7\x34\x03\x53\x94\xf0\xcb\x45\xba\x03\x32\x94\xe7\x20\x02\x17\xff\x3f\xe9\xc9\x05\xbf\x3f\xed\xf5\x48\x9f\x13\x01\x77\xb9\x87\x71\xce\xe7\x9d\x4e\x71\xc6\xe2\xc5\xd6\xad\xce\x43\x53\x52\xe6\xb5\x5c\x20\x99\x29\xf2\x93\x0e\x3a\x40\xce\x85\x39\x3b\x28\xf6\xf0\x05\xde\xff\x90\x05\xb6\xd6\xbe\x9a\xcd\x24\x5b\x63\xce\x0f\xb2\x3d\x7d\x62\x3a\x8f\xee\x9b\x5d\x4b\x76\xb1\x64\xe6\xb4\x8b\xc3\x69\x40\xf3\xf5\xdf\xd8\xeb\x60\xc9\x1e\xe1\xb2\x64\x06\x52\x20\x08\xb5\x8a\x5f\xf2\xc6\x34\xa9\xfb\x75\x12\x19\x2f\x98\x9a\xc5\x39\xa3\x62\x26\x5b\xbf\x75\x53\x56\xc8\x25\x03\xcd\xad\xc3\xc2\x50\x2e\x98\xaa\x5f\x13\x5e\xd0\x94\x69\x98\x8c\x1a\x50\xc5\xe0\xba\xa0\x58\x31\xb6\x80\xcc\x93\xe0\x94\x41\x5d\x0e\xa5\x06\xa1\x5b\xa7\xba\x36\x5e\x83\xa6\xb2\x8b\x7b\xd0\xd5\x02\xdd\xb5\xbb\x5e\xe1\xbf\xdb\xd8\xad\x3b\xaa\x52\xbd\x4d\x55\xc5\xc9\x7c\xb6\x40\x16\x52\x98\xec\x87\x31\x9b\xfe\x9e\xe6\xf9\x53\x1b\x26\x7f\x16\xe8\x23\x4c\x9f\x44\xdd\x37\xef\x71\xe3\x3e\xd4\xe9\x47\xfe\x25\x54\xa4\xb9\xfd\x78\xd5\x9c\xda\x47\xe9\x60\x42\xea\xc0\xb3\xbe\xaa\x02\x64\x0e\xff\xd5\x81\x27\xe4\xdb\xeb\x33\x9c\x67\xec\x8e\x1b\xf6\xc8\xcc\xd5\xe7\x7b\xb4\x9f\x39\x52\xc3\x3b\x6e\x70\x8f\xb8\x32\x20\x09\xba\xe8\x5a\x5a\xef\x4e\xc5\x7c\x04\x62\xe0\x3e\x01\xc9\x33\x51\x5a\x44\x3a\xdf\x55\xab\xec\xed\x72\xf2\x07\x31\x9b\xb9\x21\xdf\x40\xc8\x9c\xe7\x86\x29\xfb\x1b\xa5\x46\xea\x59\x71\xbd\x16\x99\xb5\xa1\xa6\x8d\xc9\xd5\xf6\xbb\xeb\x02\x20\xe7\xf6\x03\x28\x2c\xa7\xe7\x68\x5d\x27\x7d\x7d\x90\xed\x68\x70\x2f\x7f\x73\x93\xbb\xbd\x87\xed\xef\xd0\xab\xaf\x9d\x7f\x03\x00\x00\xff\xff\xbe\x05\x70\xc3\x0a\x0d\x00\x00")

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

	info := bindataFileInfo{name: "make/default.yaml", size: 3338, mode: os.FileMode(493), modTime: time.Unix(1472346338, 0)}
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

