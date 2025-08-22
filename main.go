// main.go
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func main() {
	// CLI flags
	saPath := flag.String("service-account", "", "Path to Firebase service account JSON")
	uid := flag.String("uid", "", "User UID to assign roles")
	role := flag.String("role", "", "Role to assign (e.g., admin)")
	flag.Parse()

	if *saPath == "" || *uid == "" || *role == "" {
		log.Fatal("Usage: go run main.go --service-account=key.json --uid=USER_UID --role=admin")
	}

	ctx := context.Background()
	opt := option.WithCredentialsFile(*saPath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v", err)
	}

	// Fetch existing claims
	user, err := client.GetUser(ctx, *uid)
	if err != nil {
		log.Fatalf("error getting user: %v", err)
	}

	claims := user.CustomClaims
	if claims == nil {
		claims = make(map[string]interface{})
	}

	// Add or update role
	if existingRoles, ok := claims["roles"]; ok {
		// merge if already exists
		roleList := []string{}
		if arr, ok := existingRoles.([]interface{}); ok {
			for _, r := range arr {
				roleList = append(roleList, r.(string))
			}
		}
		roleList = append(roleList, *role)
		claims["roles"] = roleList
	} else {
		claims["roles"] = []string{*role}
	}

	// Assign custom claims
	if err := client.SetCustomUserClaims(ctx, *uid, claims); err != nil {
		log.Fatalf("error setting custom claims: %v", err)
	}

	// Output confirmation
	claimsJSON, _ := json.Marshal(claims)
	fmt.Printf("Roles assigned to user %s: %s\n", *uid, string(claimsJSON))
}
