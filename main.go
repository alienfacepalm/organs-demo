package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const (
    projectId       = "organs-demo" 
    firestoreAPIKey = "./organs-demo-api-key.json" 
    collectionName  = "organs"
)

var firestoreClient *firestore.Client

func initFirestore() {
    ctx := context.Background()
    client, err := firestore.NewClient(ctx, projectId, option.WithCredentialsFile(firestoreAPIKey))
    if err != nil {
        log.Fatalf("Failed to create Firestore client: %v", err)
    }
    firestoreClient = client
}

type Organ struct {
	state string 
	all int 
	kidney int
	liver int 
	pancrease int 
	kidney_pancreas int 
	heart int 
	lung int 
	heart_lung int 
	intestine int 
	abdominal_wall int
	craniofacial int 
	gu_uterus int 
	upper_limb_bilateral int 
	upper_limb_unilateral int
}

func getOrgans(c *gin.Context) {
    ctx := context.Background()

    // Fetch the data from Firestore collection 'organs'
    iter := firestoreClient.Collection(collectionName).Documents(ctx)
    var organs []Organ

    for {
        doc, err := iter.Next()
        if err == iterator.Done {
            break
        }
        if err != nil {
            log.Printf("Error fetching document: %v", err)
            continue
        }

        var organ Organ
        if err := doc.DataTo(&organ); err != nil {
            log.Printf("Error converting document to Organ: %v", err)
            continue
        }
		fmt.Print(organ)
        organs = append(organs, organ)
    }

	fmt.Println("Organs->", organs)
    c.JSON(http.StatusOK, organs)
}

func main() {
    initFirestore()

    r := gin.Default()
    r.GET("/organs", getOrgans)
    r.Run(":8080") // Run the server on port 8080
}
