package db

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type FileDB struct {
	db          *DB
	networkpath string
	tracepath   string
}

func NewFileDB(path string) (*FileDB, error) {
	var filedb *FileDB
	var newdb *DB

	newdb = NewDB(path)

	networkpath := filepath.Join(path, "/network")
	tracepath := filepath.Join(path, "/trace")

	files, err := ioutil.ReadDir(networkpath)

	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && len(file.Name()) == 21 {
			f, err := ioutil.ReadFile(filepath.Join(networkpath, "/", file.Name()))
			if err != nil {
				continue
			}

			var v Networkconfig

			err = json.Unmarshal(f, &v)
			if err != nil {
				continue
			}

			if v.ID == nil {
				continue
			}

			newdb.Onnetworkchanged(nil, &v, true)

			memberpath := filepath.Join(networkpath, "/", *v.ID, "/member")
			mfiles, merr := ioutil.ReadDir(memberpath)
			if merr != nil {
				continue
			}

			for _, mfile := range mfiles {
				if !mfile.IsDir() && len(mfile.Name()) == 15 {
					mf, mferr := ioutil.ReadFile(filepath.Join(memberpath, "/", mfile.Name()))
					if mferr != nil {
						continue
					}

					var mv Memberconfig

					mferr = json.Unmarshal(mf, &mv)
					if mferr != nil {
						continue
					}

					newdb.Onmemberchanged(nil, &mv, true)

				}
			}

		}

	}

	filedb = &FileDB{
		networkpath: networkpath,
		tracepath:   tracepath,
		db:          newdb,
	}

	return filedb, nil

}
