package hashmap

import "testing"

func TestHash(t *testing.T) {
	word := "Adam"

	sum := 0
	for char := range word {
		sum += char
	}
	sum %= size

	if sum != hash(word) {
		t.Error("Hash function gives wrong value")
	}
}

func TestInsertInBucket(t *testing.T) {
	var bucket = &Bucket{}
	bucket.insert("Adam")

	if (*bucket.Head).Data != "Adam" {
		t.Error("Data was not inserted into bucket")
	}
}

func TestSearchInBucket(t *testing.T) {
	var bucket = &Bucket{}
	bucket.insert("Robin")

	if bucket.search("Robin") == false {
		t.Error("Data was not inserted into bucket")
	}
}

func TestDeleteInBucket(t *testing.T) {
	var bucket = &Bucket{}
	bucket.insert("Adam")
	bucket.insert("Robin")
	bucket.insert("George")

	// test delete from front
	bucket.detele("Adam")
	if bucket.search("Adam") == true {
		t.Error("Data was not deleted from bucket")
	}

	// test delete from end
	bucket.detele("George")
	if bucket.search("George") == true {
		t.Error("Data was not deleted from bucket")
	}

	// test delete from middle
	bucket.insert("Aakash")
	bucket.insert("Sheeldon")

	bucket.detele("Aakash")
	if bucket.search("Aakash") == true {
		t.Error("Data was not deleted from bucket")
	}

}

func TestInsert(t *testing.T) {
	var hashmap = &HashMap{}

	hashmap.Insert("Adam")
	hashmap.Insert("Aakash")
	hashmap.Insert("Robin")

	var key = hash("Adam")
	var head *Node = hashmap.Items[key].Head
	var found = (*head).Data

	if found != "Adam" {
		t.Errorf("Expected Adam found %s", found)
	}
}

func TestSearch(t *testing.T) {
	var hashmap = &HashMap{}

	hashmap.Insert("Adam")
	hashmap.Insert("Aakash")
	hashmap.Insert("Robin")

	if hashmap.Search("Aakash") == false {
		t.Error("Failed to search data")
	}
}

func TestDelete(t *testing.T) {
	var hashmap = &HashMap{}

	hashmap.Insert("Adam")
	hashmap.Insert("Aakash")
	hashmap.Insert("Robin")

	hashmap.Delete("Aakash")

	if hashmap.Search("Aakash") == true {
		t.Error("Failed to delete data")
	}
}
