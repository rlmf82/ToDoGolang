package model

import "fmt"

func Create(name string, todo string) error {

	insert, err := con.Query(fmt.Sprintf("INSERT INTO test.TODO(name, todo) VALUES('%s', '%s');", name, todo))

	defer insert.Close()

	if err != nil {
		return err
	}

	return nil
}

func Delete(name string) error {

	delete, err := con.Query("DELETE FROM test.TODO where name=?;", name)

	defer delete.Close()

	if err != nil {
		return err
	}

	return nil
}
