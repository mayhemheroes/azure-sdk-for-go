// +build emulator
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

import (
	"context"
	"testing"
)

type emulatorTests struct {
	host string
	key  string
}

func newEmulatorTests() *emulatorTests {
	return &emulatorTests{
		host: "https://localhost:8081/",
		key:  "C2y6yDjf5/R+ob0N8A7Cgv30VRDJIWEHLM+4QDU5DE2nQ9nDuVTqobD4b8mGGyPMbIZnqyMsEcaGQy67XIw/Jw==",
	}
}

func (e *emulatorTests) getClient(t *testing.T) *CosmosClient {
	cred, _ := NewSharedKeyCredential(e.key)
	client, err := NewCosmosClient(e.host, cred, nil)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	return client
}

func (e *emulatorTests) createDatabase(
	t *testing.T,
	ctx context.Context,
	client *CosmosClient,
	dbName string) *CosmosDatabase {
	database := CosmosDatabaseProperties{Id: dbName}
	resp, err := client.AddDatabase(ctx, database, nil)
	if err != nil {
		t.Fatalf("Failed to create database: %v", err)
	}

	if resp.DatabaseProperties.Id != database.Id {
		t.Errorf("Unexpected id match: %v", resp.DatabaseProperties)
	}

	return resp.DatabaseProperties.Database
}

func (e *emulatorTests) deleteDatabase(
	t *testing.T,
	ctx context.Context,
	database *CosmosDatabase) {
	_, err := database.Delete(ctx, nil)
	if err != nil {
		t.Fatalf("Failed to delete database: %v", err)
	}
}
