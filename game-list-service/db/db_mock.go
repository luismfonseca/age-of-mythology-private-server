package db

import "github.com/stretchr/testify/mock"

type DatabaseMock struct {
	mock.Mock
}

func (dbMock *DatabaseMock) OnInsertGame(id string, gameBytes []byte) *mock.Call {
	return dbMock.On("InsertGame", id, gameBytes)
}

func (dbMock *DatabaseMock) OnAnyInsertGame() *mock.Call {
	return dbMock.On("InsertGame", mock.Anything, mock.Anything)
}

func (dbMock *DatabaseMock) InsertGame(id string, gameBytes []byte) error {
	args := dbMock.Called(id, gameBytes)
	return args.Error(0)
}

func (dbMock *DatabaseMock) OnDeleteGame(id string) *mock.Call {
	return dbMock.On("DeleteGame", id)
}

func (dbMock *DatabaseMock) OnAnyDeleteGame() *mock.Call {
	return dbMock.On("DeleteGame", mock.Anything)
}

func (dbMock *DatabaseMock) DeleteGame(id string) error {
	args := dbMock.Called(id)
	return args.Error(0)
}

func (dbMock *DatabaseMock) OnListGames(offset int, amount int) *mock.Call {
	return dbMock.On("ListGames", offset, amount)
}

func (dbMock *DatabaseMock) OnAnyListGames() *mock.Call {
	return dbMock.On("ListGames", mock.Anything)
}

func (dbMock *DatabaseMock) ListGames(offset int, amount int) (map[string][]byte, error) {
	args := dbMock.Called(offset, amount)
	return args.Get(0).(map[string][]byte), args.Error(1)
}

func (dbMock *DatabaseMock) OnGetNumberOfGames() *mock.Call {
	return dbMock.On("GetNumberOfGames")
}

func (dbMock *DatabaseMock) GetNumberOfGames() int {
	args := dbMock.Called()
	return args.Get(0).(int)
}
