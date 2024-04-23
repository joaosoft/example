package domain

type Person struct {
	Id   int    `json:"id" db:"id_person"`
	Name string `json:"name" db:"name"`
	Age  int    `json:"age" db:"age"`
}

type SavePerson struct {
	Name string `json:"name" db:"name"`
	Age  int    `json:"age" db:"age"`
}
