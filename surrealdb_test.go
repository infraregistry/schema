package schema

import (
	"log"
	"math/rand"
	"testing"

	"github.com/infraregistry/schema/models"
	"github.com/jaswdr/faker/v2"
	"github.com/surrealdb/surrealdb.go"
)

func TestConnectSurrealDB(t *testing.T) {
	db, err := ConnectSurrealDB("ws://localhost:9000/rpc", "root", "root", "test", "test")
	if err != nil {
		t.Fatal(err)
	}

	if _, err = db.Use("test", "test"); err != nil {
		t.Fatal(err)
	}

	f := faker.New()

	createdComponents := make([]models.Component, 0)
	for i := 0; i < 10; i++ {
		record, err := db.Create("component", models.Component{
			Name:        f.App().Name(),
			Description: f.Lorem().Sentence(rand.Intn(50)),
			Status:      models.ComponentStatusActive,
		})
		if err != nil {
			t.Fatal(err)
		}

		// Unmarshal data
		var component []models.Component
		err = surrealdb.Unmarshal(record, &component)
		if err != nil {
			t.Fatal(err)
		}

		createdComponents = append(createdComponents, component[0])
	}

	_, err = db.Query("RELATE $left->component_link->$right", map[string]string{
		"left":  createdComponents[0].ID,
		"right": createdComponents[1].ID,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Query("RELATE $left->component_link->$right", map[string]string{
		"left":  createdComponents[0].ID,
		"right": createdComponents[2].ID,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Query("RELATE $left->component_link->$right", map[string]string{
		"left":  createdComponents[1].ID,
		"right": createdComponents[3].ID,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Query("RELATE $left->component_link->$right", map[string]string{
		"left":  createdComponents[3].ID,
		"right": createdComponents[4].ID,
	})
	if err != nil {
		t.Fatal(err)
	}

	// if len(createdUser) == 0 || createdUser[0].ID == "" {
	// 	t.Fatalf("Invalid created user ID")
	// }

	// // Get user by ID
	// data, err = db.Select(createdUser[0].ID)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// // Unmarshal data
	// var selectedUser User
	// err = surrealdb.Unmarshal(data, &selectedUser)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// if selectedUser.ID == "" {
	// 	t.Fatalf("Invalid selected user ID")
	// }

	// // Change part/parts of user
	// changes := map[string]string{"name": "Jane"}
	// if _, err = db.Change(selectedUser.ID, changes); err != nil {
	// 	t.Fatal(err)
	// }

	// // Update user
	// if _, err = db.Update(selectedUser.ID, changes); err != nil {
	// 	t.Fatal(err)
	// }

	// if _, err = db.Query("SELECT * FROM $record", map[string]interface{}{
	// 	"record": createdUser[0].ID,
	// }); err != nil {
	// 	t.Fatal(err)
	// }

	// // Delete user by ID
	// if _, err = db.Delete(selectedUser.ID); err != nil {
	// 	t.Fatal(err)
	// }
}

func TestGetGraph(t *testing.T) {
	res, err := GetGraph[[]models.Component]("component")
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("res: %v", res)
}

func TestSelect(t *testing.T) {
	res, err := Select[models.Component]("component")
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("res: %v", res)
}

func TestSelectOnly(t *testing.T) {
	res, err := SelectOnly[models.Component]("component")
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("res: %v", res)
}
