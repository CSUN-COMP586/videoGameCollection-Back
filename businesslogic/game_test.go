package businesslogic

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	mocket "github.com/selvatico/go-mocket"
	"github.com/videoGameLibrary/videogamelibrary/config/database"
)

func TestCreateNewGameEntry(t *testing.T) {
	db := SetupTests()
	defer db.Close()

	// commonReply := []map[string]interface{}{{
	// 	"dev_id":         5,
	// 	"pub_id":         10,
	// 	"genre":          "Role-Playing Game",
	// 	"search_created": time.Now(),
	// 	"game_name":      "Witcher",
	// 	"summary":        "You hunt stuff",
	// }}
	mocket.Catcher.Reset().NewMock().WithQuery("INSERT INTO games")
	// .WithReply(commonReply)

	game := Game{
		DevID:         5,
		PubID:         10,
		Genre:         "Role-Playing Game",
		SearchCreated: time.Now(),
		GameName:      "Witcher",
		Summary:       "You hunt stuff",
	}

	handler := GameHandler{Model: &game}

	returnID, err := handler.CreateNewGameEntry(db)
	if err != nil {
		t.Fatalf("Some error")
		t.Fatal(err)
	}
	fmt.Println(returnID)
}

// SetupTests mocks the gorm orm database connection - will move to separate folder in the future
func SetupTests() *gorm.DB {
	mocket.Catcher.Register()

	db, err := gorm.Open(mocket.DriverName, "")
	if err != nil {
		log.Fatal(err)
	}
	database.GormConn = db

	db.LogMode(true)

	return db
}
