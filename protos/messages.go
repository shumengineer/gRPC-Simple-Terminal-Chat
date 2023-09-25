package protos

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

const file string = "messages.db"

var messagesTableName string = "messages"

const initTable string = `CREATE TABLE [messages] (
   id INTEGER NOT NULL PRIMARY KEY,
   user VARCHAR(32) NOT NULL,
   body TEXT
   );`

type MessagesTable struct {
	my sync.Mutex
	db *sql.DB
}

func tableCheck(db *sql.DB) error {
	_, table_check := db.Query("select * from " + messagesTableName + ";")

	return table_check
}

func NewMessage() (*MessagesTable, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	if tableCheck(db) != nil {
		if _, err := db.Exec(initTable); err != nil {
			return nil, err
		}
	}

	return &MessagesTable{
		db: db,
	}, nil
}

func (c *MessagesTable) Insert(m *Message) (int, error) {
	res, err := c.db.Exec("INSERT INTO messages VALUES(NULL,?,?);", m.User, m.Body)
	if err != nil {
		return 0, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}

func (c *MessagesTable) List(limit int) ([]*Message, error) {
	rows, err := c.db.Query("SELECT * FROM messages ORDER BY id ASC LIMIT ?", limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	data := []*Message{}

	for rows.Next() {
		i := Message{}
		err = rows.Scan(&i.Id, &i.User, &i.Body)

		if err != nil {
			log.Println("[Messages.List] Error: ", err)
			continue
		}

		data = append(data, &i)
	}

	return data, nil
}
