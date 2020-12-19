package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/asdine/storm"
	"github.com/stretchr/testify/require"
	"github.com/waarzitjenu/server/database/types"
)

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	dbForTest   = "./database/" + RandStringRunes(20) + ".db"
	stormDB     *storm.DB
)

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandomInt(min int, max int) int {
	return rand.Intn(max-min) + min
}

func InitRandom() {
	rand.Seed(time.Now().UnixNano())
}

func randomLocation() types.LocationUpdate {
	return types.LocationUpdate{
		Latitude:  RandomFloat64(-180, 180),
		Longitude: RandomFloat64(-90, 90),
		Timestamp: uint64(time.Now().Unix()) - uint64(RandomInt(0, 5)),
		Hdop:      RandomFloat64(1, 20),
		Altitude:  RandomFloat64(0, 680),
		Speed:     RandomFloat64(0, 16),
	}
}

func TestOpen(t *testing.T) {
	t.Run("Opening a database", func(t *testing.T) {
		db, err := Open(dbForTest)
		if err != nil {
			require.Error(t, err)
			require.Contains(t, err.Error(), err.Error())
			return
		}

		require.NoError(t, err)
		require.IsType(t, &storm.DB{}, db)
		stormDB = db
	})
}

func TestUse(t *testing.T) {
	t.Run("Using a database", func(t *testing.T) {
		err := Use(stormDB)
		if err != nil {
			require.Error(t, err)
			require.Contains(t, err.Error(), errors.New("Another database is already in use"))
			return
		}
		require.NoError(t, err)
	})

	t.Run("Using a different database again (should fail)", func(t *testing.T) {
		err := Use(stormDB)
		if err != nil {
			require.Error(t, err)
			require.Contains(t, err.Error(), "Another database is already in use")
			return
		}
		require.Error(t, err)
		require.Nil(t, db)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Saving a single entry to the database", func(t *testing.T) {
		entry := types.LocationUpdate{
			Latitude:  10,
			Longitude: 11,
			Timestamp: 123,
			Hdop:      1,
			Altitude:  300,
			Speed:     13.3333,
		}

		err := Create(entry)
		if err != nil {
			require.Error(t, err)
			return
		}

		require.NoError(t, err)
	})

	t.Run("Adding multiple entries to database", func(t *testing.T) {
		var multipleEntries [5]types.LocationUpdate
		for i := 0; i < 5; i++ {
			multipleEntries[i] = randomLocation()
		}

		err := CreateMultiple(multipleEntries[:]...)
		require.NoError(t, err)
	})

}

func TestRead(t *testing.T) {
	t.Run("Reading entries from database", func(t *testing.T) {
		entries, err := Read(10)
		if err != nil {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)

		data, _ := json.MarshalIndent(entries, "", "  ")
		fmt.Println(string(data))
	})
}

func TestDestroy(t *testing.T) {
	t.Run("Destroying database", func(t *testing.T) {
		err := Destroy(dbForTest)
		if err != nil {
			require.Error(t, err)
			return
		}

		require.NoError(t, err)
	})
}
