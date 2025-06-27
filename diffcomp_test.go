package dkit

import "testing"

func TestShowDiff(t *testing.T) {
	diff := NewDiffCompare()
	s1 := `{
  "appConfig": {
    "appName": "WebApp v1.0",
    "environment": "development",
    "debugMode": true,
    "logLevel": "INFO",
    "features": [
      "userAuth",
      "notifications",
      "analytics"
    ],
    "database": {
      "type": "MongoDB",
      "host": "localhost",
      "port": 27017,
      "user": "devUser",
      "password": "devPassword"
    },
    "timeouts": {
      "connect": 5000,
      "read": 10000
    },
    "details": "{\"buildId\":\"abc-123\",\"deploymentDate\":\"2025-01-15T10:00:00Z\",\"configHash\":\"a1b2c3d4\"}",
    "adminEmails": [
      "admin1@example.com",
      "admin2@example.com"
    ]
  },
  "userInfo": {
    "id": "user123",
    "name": "Alice Wonderland",
    "age": 30,
    "isActive": true,
    "roles": ["user"],
    "profile": {
      "city": "New York",
      "zipCode": "10001"
    },
    "details": "{\"lastActivity\":\"2025-06-20T14:30:00Z\",\"browser\":\"Chrome\",\"os\":\"macOS\"}"
  }
}`
	s2 := `{
  "appConfig": {
    "appName": "WebApp v2.0-PROD",
    "environment": "production",
    "debugMode": false,
    "features": [
      "userAuth",
      "emailService",
      "notifications"
    ],
    "newSetting": "enabled",
    "database": {
      "type": "PostgreSQL",
      "host": "prod.db.example.com",
      "port": 5433,
      "user": "prodUser",
      "dbName": "prod_app_db"
    },
    "timeouts": {
      "connect": 5000,
      "read": 12000,
      "write": 15000
    },
    "details": "{\"buildId\":\"xyz-456\",\"deploymentDate\":\"2025-06-27T10:00:00Z\",\"environment\":\"production\"}",
    "adminEmails": [
      "admin3@example.com",
      "admin1@example.com"
    ],
    "auditLog": "{\"timestamp\":\"2025-06-27T10:30:00Z\",\"action\":\"config_update\",\"user\":\"devops_team\",\"changes\":{\"logLevel\":\"deleted\",\"newSetting\":\"added\"}}"
  },
  "userInfo": {
    "id": "user123",
    "name": "Bob Smith",
    "age": 31,
    "isActive": false,
    "roles": ["user", "admin"],
    "profile": {
      "city": "London",
      "country": "UK"
    },
    "lastLogin": "2025-06-27T10:00:00Z",
    "details": "{\"lastActivity\":\"2025-06-27T10:45:00Z\",\"browser\":\"Firefox\",\"os\":\"Windows\",\"device\":\"laptop\"}"
  }
}`
	diff.ShowDiff(s1, s2)
}
