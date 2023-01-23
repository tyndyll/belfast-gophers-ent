package main


import (
	"net/http"
)

type DBType string

const (
	Postgres = "postgres"
	MySQL = "mysql"
	// ent supports sqlite but the preferred sqlite package is a cgo package 
	//   so I'm omitting it to keep the build easy
	// SQLite = "sqlite"

)


func NewRepo(dbtype DBType, user, password, host, db string, port int) (*LotrRepo, error) {
	var client ent.Client
	var err error

	switch(dbtype) {
	case Postgres:
		connectionString = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", host, port, user, db, password)
		client, err = ent.Open("postgres",)
	case MySQL:
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", user, password, host, port, db)
		client, err = ent.Open("mysql", connectionString)	
	}
	if err != nil {
        return nil, err
    }

	repo := &LotrRepo{
		DB: client,
	}

	return repo, nil
}

type LotrRepo struct {
	DB *ent.Client
}

func (repo *LotrRepo) GetAllCharacters() ([]string, error) {
	panic("not implemented")
	return nil, nil
}

func (repo *LotrRepo) Close() error {
	
}

type LotrWeb struct {
	Repository
}

func (web *LotrWeb) GetCharacter(w http.ResponseWriter, r *http.Request) {

}



func main() {
	webService := &LotrWeb{}

	http.HandleFunc("/character", webService.GetCharacter)

	http.ListenAndServe(":8080", nil)
}