package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Cached is set when the preview_picture is already downloaded
type Cached struct {
	Status   string `bson:"status"`
	FileName string `bson:"filename"`
}

// Link struct is for listing many links
type Link struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id"`
	Title          string             `bson:"title" json:"title"`
	URL            string             `bson:"url" json:"url"`
	Tags           []string           `bson:"tags" json:"tags"`
	DomainName     string             `bson:"domain_name" json:"domain_name"`
	PreviewPicture string             `bson:"preview_picture" json:"preview_picture"`
	ReadingTime    float64            `bson:"reading_time" json:"reading_time"`
	CreatedAt      string             `bson:"created_at" json:"created_at"`
	UpdatedAt      string             `bson:"updated_at" json:"updated_at"`
	IsStarred      bool               `bson:"is_starred" json:"is_starred"`
	IsArchived     bool               `bson:"is_archived" json:"is_archived"`
	Content        string             `bson:"content,omitempty" json:"content,omitempty"`
	Cached         Cached             `bson:"cached" json:"cached"`
	Updated        time.Time          `bson:"updatedAt" json:"updatedAt"`
	Marked         string             `bson:"marked,omitempty" json:"marked,omitempty"`
}

// Tag struct for listing distinct tags
type Tag struct {
	ID string `bson:"_id" json:"_id"`
}

// DB struct for managing the mongo database
type DB struct {
	URL    string
	Name   string
	Client *mongo.Client
}

// Connect method initiates the connection to the server and returns the client
func (db *DB) Connect() error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.URL))
	if err != nil {
		log.Fatal(err)
		return err
	}
	db.Client = client
	return nil
}

// GetAll returns all items in the given database of the given collection
func (db *DB) GetAll(collname string) ([]*Link, error) {
	filter := bson.D{}
	results, err := db.GetAllWith(collname, &filter)
	if err != nil {
		log.Fatal(err)
	}
	return results, nil
}

// GetOne method that returns a specific link given the id
func (db *DB) GetOne(collname string, id string) (*Link, error) {
	collection := db.Client.Database(db.Name).Collection(collname)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatalf("ObjectIDFromHex: %v", err)
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	var result *Link
	e := bson.D{}
	err = collection.FindOne(context.TODO(), filter).Decode(&e)
	if err != nil {
		log.Fatalf("FindOne: %v", err)
		return nil, err
	}
	b, _ := bson.Marshal(e)
	err = bson.Unmarshal(b, &result)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return nil, err
	}
	return result, nil
}

// GetTagged returns all items that contain the given tag
func (db *DB) GetTagged(collname string, tag string) ([]*Link, error) {
	filter := bson.D{{
		Key: "tags", Value: tag,
	}}
	results, err := db.GetAllWith(collname, &filter)
	if err != nil {
		log.Fatal(err)
	}
	return results, nil
}

// GetArchived returns all items that is_archived=true
func (db *DB) GetArchived(collname string) ([]*Link, error) {
	filter := bson.D{bson.E{
		Key:   "is_archived",
		Value: true,
	}}
	results, err := db.GetAllWith(collname, &filter)
	if err != nil {
		log.Fatal(err)
	}
	return results, nil
}

// GetArchivedID returns all items that is_archived=true
func (db *DB) GetArchivedID(collname string, id string) (*Link, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatalf("ObjectIDFromHex: %v", err)
	}
	filter := bson.D{
		{
			Key:   "_id",
			Value: objID,
		},
		bson.E{
			Key:   "is_archived",
			Value: true,
		},
	}
	var result *Link
	e := bson.D{}
	collection := db.Client.Database(db.Name).Collection(collname)
	err = collection.FindOne(context.TODO(), filter).Decode(&e)
	if err != nil {
		log.Fatalf("FindOne: %v", err)
		return nil, err
	}
	b, _ := bson.Marshal(e)
	err = bson.Unmarshal(b, &result)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return nil, err
	}
	return result, nil
}

// GetArchivedTag returns all items that is_archived=true and with tag
func (db *DB) GetArchivedTag(collname string, tag string) ([]*Link, error) {
	filter := bson.D{
		bson.E{
			Key:   "is_archived",
			Value: true,
		},
		bson.E{
			Key: "tags", Value: tag,
		},
	}
	results, err := db.GetAllWith(collname, &filter)
	if err != nil {
		log.Fatal(err)
	}
	return results, nil
}

// GetStarred returns all items that is_starred=true
func (db *DB) GetStarred(collname string) ([]*Link, error) {
	filter := bson.D{bson.E{
		Key:   "is_starred",
		Value: true,
	}}
	results, err := db.GetAllWith(collname, &filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return results, nil
}

// GetStarredTag returns all items that is_starred=true and with tag
func (db *DB) GetStarredTag(collname string, tag string) ([]*Link, error) {
	filter := bson.D{
		bson.E{
			Key:   "is_starred",
			Value: true,
		},
		bson.E{
			Key: "tags", Value: tag,
		},
	}
	results, err := db.GetAllWith(collname, &filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return results, nil
}

// GetAllWith gets all items from the given collection with the given filter
func (db *DB) GetAllWith(collname string, filter *bson.D) ([]*Link, error) {
	projection := bson.D{
		bson.E{Key: "title", Value: 1},
		bson.E{Key: "url", Value: 1},
		bson.E{Key: "tags", Value: 1},
		bson.E{Key: "domain_name", Value: 1},
		bson.E{Key: "preview_picture", Value: 1},
		bson.E{Key: "reading_time", Value: 1},
		bson.E{Key: "created_at", Value: 1},
		bson.E{Key: "updated_at", Value: 1},
		bson.E{Key: "is_starred", Value: 1},
		bson.E{Key: "is_archived", Value: 1},
		bson.E{Key: "cached", Value: 1},
		bson.E{Key: "updatedAt", Value: 1},
	}
	collection := db.Client.Database(db.Name).Collection(collname)
	cur, err := collection.Find(context.TODO(), filter, options.Find().SetProjection(projection))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cur.Close(context.TODO())

	var results []*Link

	for cur.Next(context.TODO()) {
		var result Link
		e := bson.D{}
		err := cur.Decode(&e)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		b, _ := bson.Marshal(e)
		err = bson.Unmarshal(b, &result)

		results = append(results, &result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return results, nil
}

// GetAllTags returns a distinct list of all tags
func (db *DB) GetAllTags(collname string) ([]*Tag, error) {
	// db.links.aggregate([{$unwind:"$tags"},{$group:{_id:"$tags"}}])
	pipeline := mongo.Pipeline{
		{{Key: "$unwind", Value: "$tags"}},
		{{Key: "$group", Value: bson.D{bson.E{Key: "_id", Value: "$tags"}}}},
	}
	collection := db.Client.Database(db.Name).Collection(collname)
	cur, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cur.Close(context.TODO())

	var results []*Tag

	for cur.Next(context.TODO()) {
		var result Tag
		e := bson.D{}
		err := cur.Decode(&e)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		b, _ := bson.Marshal(e)
		err = bson.Unmarshal(b, &result)

		results = append(results, &result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return results, nil
}

/* NOT WORKING YET
// SearchFor returns a ranked list of items whose contents/marked contain the search string
func (db *DB) SearchFor(collname string, q string) ([]*Link, error) {
	qstr := strings.Replace(q, "+", " ", -1)
	filter := bson.D{bson.E{"$text", bson.E{"$search", qstr}}}
	results, err := db.GetAllWith(collname, &filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return results, nil
}
*/

// CreateOne inserts/creates a new link from the passed link struct
func (db *DB) CreateOne(collname string, link *Link) error {
	if link.ID.IsZero() {
		link.ID = primitive.NewObjectID()
	}
	collection := db.Client.Database(db.Name).Collection(collname)
	_, err := collection.InsertOne(context.TODO(), link)
	if err != nil {
		log.Fatalf("InsertOne: %v\n", err)
		return err
	}
	return nil
}
