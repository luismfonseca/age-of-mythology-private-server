package db

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/syndtr/goleveldb/leveldb"
)

var dummyGameBytes = []byte{1, 2, 3}

func TestInstanceDB(t *testing.T) {
	Convey("Creating a new db should always succeed", t, func() {
		path, err := ioutil.TempDir("", "game-list-services-test-db")
		So(err, ShouldBeNil)

		Reset(func() {
			os.Remove(path)
		})

		db, err := New(path)
		So(err, ShouldBeNil)
		So(db, ShouldNotBeNil)

		Convey("InsertGame", func() {
			Convey("should store it successfully", func() {
				err := db.InsertGame("name", dummyGameBytes)
				So(err, ShouldBeNil)

				value, err := db.DB.Get([]byte("name"), nil)
				So(err, ShouldBeNil)
				So(value, ShouldResemble, dummyGameBytes)
			})
		})

		Convey("DeleteGame", func() {
			Convey("should delete successfully a previously stored game", func() {
				err := db.DB.Put([]byte("name"), dummyGameBytes, nil)
				So(err, ShouldBeNil)

				value, err := db.DB.Get([]byte("name"), nil)
				So(err, ShouldBeNil)
				So(value, ShouldResemble, dummyGameBytes)

				err = db.DeleteGame("name")
				So(err, ShouldBeNil)

				value, err = db.DB.Get([]byte("name"), nil)
				So(err, ShouldEqual, leveldb.ErrNotFound)
			})
		})

		Convey("ListGames", func() {
			// TODO
		})
	})
}
