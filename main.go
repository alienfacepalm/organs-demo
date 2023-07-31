package main

import (
	"context"
	// "fmt"
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
    State string `firestore:"state"`
    All string `firestore:"all"`
    Kidney string `firestore:"kidney"`
    Liver string `firestore:"liver"`
    Pancrease string `firestore:"pancrease"`
    KidneyPancreas string `firestore:"kidney_pancreas"`
    Heart string `firestore:"heart"`
    Lung string `firestore:"lung"`
    HeartLung string `firestore:"heart_lung"`
    Intestine string `firestore:"intestine"`
    AbdominalWall string `firestore:"abdominal_wall"`
    Craniofacial string `firestore:"craniofacial"`
    GUUterus string `firestore:"gu_uterus"`
    UpperLimbBilateral string `firestore:"upper_limb_bilateral"`
    UpperLimbUnilateral string `firestore:"upper_limb_unilateral"`
}


func getOrgans(c *gin.Context) {
    ctx := context.Background()

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

        organs = append(organs, organ)
    }

    c.JSON(http.StatusOK, organs)
}

func main() {
    initFirestore()

    r := gin.Default()
    r.GET("/organs", getOrgans)
    r.Run(":8080") 
}
