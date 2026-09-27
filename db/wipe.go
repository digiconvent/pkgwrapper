package db

import (
	"database/sql"
	"os"
)

func (s *SqliteDatabase) Wipe() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.DB.Close()

	for _, suffix := range []string{"", "-wal", "-shm"} {
		if err := os.Remove(s.uri + suffix); err != nil && !os.IsNotExist(err) {
			return err
		}
	}

	newDB, err := sql.Open("sqlite3", s.uri)
	if err != nil {
		return err
	}
	s.DB = newDB
	return nil
}
