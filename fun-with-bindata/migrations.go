package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var __20200729095341_init_db_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x72\x75\xf7\xf4\xb3\xe6\xe2\xe2\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\xc8\x49\x2c\x49\xcb\x2f\xca\xb5\x86\x48\x07\x3b\x7b\xb8\xfa\x3a\x22\xc9\x2b\x15\x64\xa7\xa7\x16\xc5\xa7\x95\xe6\x29\x29\x38\x3b\x06\x3b\x3b\xba\xb8\x5a\x73\x71\x39\xfb\xfb\xfa\x7a\x86\x58\x03\x02\x00\x00\xff\xff\xb2\x16\xb3\x83\x5b\x00\x00\x00")

func _20200729095341_init_db_down_sql() ([]byte, error) {
	return bindata_read(
		__20200729095341_init_db_down_sql,
		"20200729095341_init_db.down.sql",
	)
}

var __20200729095341_init_db_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xce\xbd\x0a\xc2\x30\x14\xc5\xf1\x3d\x4f\x71\xc8\xa4\xe0\x1b\x64\x4a\xcb\x55\x2f\x26\xa9\x26\x37\x60\x27\x11\x8c\x22\x62\x91\x7e\x80\x8f\x2f\x2a\x2e\xce\xe7\x77\xe0\x5f\xd1\x8a\x83\x51\xaa\x8e\x64\x85\x90\xea\x35\x79\x0b\x5e\x22\x34\x02\xda\x73\x92\x04\xfd\xb8\x5d\x4a\x7f\x38\x4f\x9d\x36\x3f\x28\xb6\x72\xf4\xe7\xa6\xa1\xf4\x03\x66\x0a\x00\xf4\xf5\xa4\x91\x28\xb2\x75\xd8\x46\xf6\x36\xb6\xd8\x50\xbb\xf8\x8e\xdd\xf1\x5e\x34\xc6\xf2\x1c\x3f\xff\x90\x9d\x43\x0e\xbc\xcb\xa4\xe6\xef\x98\xc6\x7b\x16\xf3\x0a\x00\x00\xff\xff\x6f\x43\xc1\xdb\x9c\x00\x00\x00")

func _20200729095341_init_db_up_sql() ([]byte, error) {
	return bindata_read(
		__20200729095341_init_db_up_sql,
		"20200729095341_init_db.up.sql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"20200729095341_init_db.down.sql": _20200729095341_init_db_down_sql,
	"20200729095341_init_db.up.sql": _20200729095341_init_db_up_sql,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"20200729095341_init_db.down.sql": &_bintree_t{_20200729095341_init_db_down_sql, map[string]*_bintree_t{
	}},
	"20200729095341_init_db.up.sql": &_bintree_t{_20200729095341_init_db_up_sql, map[string]*_bintree_t{
	}},
}}
