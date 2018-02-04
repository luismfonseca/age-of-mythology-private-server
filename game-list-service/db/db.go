package db

// Implemented using LevelDB.

import (
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type Database interface {
	InsertGame(id string, gameBytes []byte) (dbError error)
	DeleteGame(id string) (dbError error)
	ListGames(offset int, amount int) (gamesBytes map[string][]byte, dbError error)
	GetNumberOfGames() (count int)
}

type database struct {
	*leveldb.DB
}

// compile time check to verify `database` type respects the `Database` interface
var _ Database = &database{}
var defaultWriteOptions = &opt.WriteOptions{Sync: true}

func New(path string) (*database, error) {
	levelDB, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}

	return &database{levelDB}, err
}

func (d *database) InsertGame(id string, gameBytes []byte) error {
	return errors.Wrap(
		d.DB.Put([]byte(id), gameBytes, defaultWriteOptions),
		"Failed to insert game")
}

func (d *database) DeleteGame(id string) error {
	return errors.Wrap(
		d.DB.Delete([]byte(id), defaultWriteOptions),
		"Failed to delete game")
}

func (d *database) ListGames(offset int, amount int) (map[string][]byte, error) {
	iterator := d.DB.NewIterator(nil, nil)
	defer iterator.Release()

	gamesBytes := make(map[string][]byte)
	for range make([]interface{}, offset) { // very silly, but it works :o
		hasMore := iterator.Next()
		if !hasMore {
			return nil, nil
		}
	}

	for range make([]interface{}, amount) { // yes, still very silly
		hasMore := iterator.Next()
		if !hasMore {
			return gamesBytes, nil
		}
		gamesBytes[string(iterator.Key())] = make([]byte, len(iterator.Value()))
		copy(gamesBytes[string(iterator.Key())], iterator.Value())
	}

	return gamesBytes, nil
}

func (d *database) GetNumberOfGames() int {
	iterator := d.DB.NewIterator(nil, nil)

	count := 0
	for iterator.Next() {
		count++
	}
	return count
}

func (d *database) Close() {
	d.DB.Close()
}
