package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

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

// ... existing code ...

func initFirestore() {
    ctx := context.Background()
    
    // Try current directory first, then parent directory
    keyFile := "./organs-demo-api-key.json"
    if _, err := os.Stat(keyFile); os.IsNotExist(err) {
        keyFile = "../organs-demo-api-key.json"
    }
    
    client, err := firestore.NewClient(ctx, projectId, option.WithCredentialsFile(keyFile))
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
    Pancreas string `firestore:"pancreas"`
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
    fmt.Println("Getting organ list")
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
    
    // Add CORS middleware
    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        
        c.Next()
    })

    // Group API routes
    api := r.Group("/api")
    {
        api.GET("/organs", getOrgans)
    }

    // Add a health check endpoint
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })

    // Add error handling for 404
    r.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
    })

    fmt.Println("Server starting on :8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
