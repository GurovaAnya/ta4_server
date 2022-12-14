// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Swagger T4A",
    "version": "1.0.0"
  },
  "host": "localhost",
  "paths": {
    "/achievement": {
      "post": {
        "tags": [
          "achievement"
        ],
        "summary": "Save new achievement",
        "parameters": [
          {
            "description": "Achievement",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/AchievementRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Achievement created",
            "schema": {
              "$ref": "#/definitions/Achievement"
            }
          }
        }
      }
    },
    "/project": {
      "get": {
        "description": "Project id is calculated by the token",
        "produces": [
          "application/json"
        ],
        "tags": [
          "project"
        ],
        "summary": "Get current project",
        "operationId": "getProject",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Project"
            }
          },
          "400": {
            "description": "Invalid status value"
          }
        }
      },
      "post": {
        "security": [
          {
            "api_key": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "project"
        ],
        "summary": "Add a new project of a partner",
        "operationId": "addProject",
        "parameters": [
          {
            "description": "Project",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ProjectRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Project created",
            "schema": {
              "$ref": "#/definitions/ProjectResponse"
            }
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    }
  },
  "definitions": {
    "Achievement": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "type": "string",
          "format": "uuid",
          "example": "f74687fa-2df7-450e-8c31-993695dcebf7"
        },
        "name": {
          "type": "string",
          "example": "GTA 5"
        },
        "webhook_url": {
          "type": "string",
          "example": "example.com"
        }
      }
    },
    "AchievementRequest": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "name": {
          "type": "string",
          "example": "GTA 5"
        },
        "webhook_url": {
          "type": "string",
          "example": "example.com"
        }
      }
    },
    "Event": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string",
          "example": "occurrs if a player killed an enemy"
        },
        "id": {
          "type": "string",
          "example": "f74687fa-2df7-450e-8c31-993695dcebf7"
        },
        "sku": {
          "type": "string",
          "example": "enemy_kill"
        }
      }
    },
    "EventCount": {
      "type": "object",
      "properties": {
        "count": {
          "type": "integer",
          "format": "int64"
        },
        "event_id": {
          "type": "string",
          "example": "f74687fa-2df7-450e-8c31-993695dcebf7"
        }
      }
    },
    "Player": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "example": "5e6890cc-c3a1-4c26-b269-b55e715c09b8"
        },
        "wallet_address": {
          "type": "string",
          "example": "ba4TrH8sNRzQfFhJH8mR"
        }
      }
    },
    "Project": {
      "type": "object",
      "required": [
        "name",
        "webhook_url"
      ],
      "properties": {
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "name": {
          "type": "string",
          "example": "GTA 5"
        },
        "webhook_url": {
          "type": "string",
          "example": "example.com"
        }
      }
    },
    "ProjectRequest": {
      "type": "object",
      "required": [
        "name",
        "webhook_url"
      ],
      "properties": {
        "name": {
          "type": "string",
          "example": "GTA 5"
        },
        "webhook_url": {
          "type": "string",
          "example": "example.com"
        }
      }
    },
    "ProjectResponse": {
      "type": "object",
      "properties": {
        "api_key": {
          "type": "string",
          "example": "abcdef123"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        }
      }
    },
    "Trigger": {
      "type": "object",
      "properties": {
        "achievement_id": {
          "type": "string",
          "example": "3a24bc5c-4c9d-11ed-bdc3-0242ac120002"
        },
        "events": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/EventCount"
          }
        },
        "id": {
          "type": "string",
          "example": "5e6890cc-c3a1-4c26-b269-b55e715c09b8"
        },
        "sku": {
          "type": "string",
          "example": "enemy_kill"
        }
      }
    },
    "WebhookEvent": {
      "type": "object",
      "properties": {
        "event_id": {
          "type": "string",
          "example": "5e6890cc-c3a1-4c26-b269-b55e715c09b8"
        },
        "player_id": {
          "type": "string",
          "example": "5e6890cc-c3a1-4c26-b269-b55e715c09b8"
        }
      }
    },
    "WebhookTrigger": {
      "type": "object",
      "properties": {
        "event_id": {
          "type": "string",
          "example": "5e6890cc-c3a1-4c26-b269-b55e715c09b8"
        },
        "player_id": {
          "type": "string",
          "example": "5e6890cc-c3a1-4c26-b269-b55e715c09b8"
        }
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "api_key",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Partner info",
      "name": "project"
    },
    {
      "name": "achievement"
    },
    {
      "name": "event"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Swagger T4A",
    "version": "1.0.0"
  },
  "host": "localhost",
  "paths": {
    "/achievement": {
      "post": {
        "tags": [
          "achievement"
        ],
        "summary": "Save new achievement",
        "parameters": [
          {
            "description": "Achievement",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/AchievementRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Achievement created",
            "schema": {
              "$ref": "#/definitions/Achievement"
            }
          }
        }
      }
    },
    "/project": {
      "get": {
        "description": "Project id is calculated by the token",
        "produces": [
          "application/json"
        ],
        "tags": [
          "project"
        ],
        "summary": "Get current project",
        "operationId": "getProject",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Project"
            }
          },
          "400": {
            "description": "Invalid status value"
          }
        }
      },
      "post": {
        "security": [
          {
            "api_key": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "project"
        ],
        "summary": "Add a new project of a partner",
        "operationId": "addProject",
        "parameters": [
          {
            "description": "Project",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ProjectRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Project created",
            "schema": {
              "$ref": "#/definitions/ProjectResponse"
            }
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    }
  },
  "definitions": {
    "Achievement": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "type": "string",
          "format": "uuid",
          "example": "f74687fa-2df7-450e-8c31-993695dcebf7"
        },
        "name": {
          "type": "string",
          "example": "GTA 5"
        },
        "webhook_url": {
          "type": "string",
          "example": "example.com"
        }
      }
    },
    "AchievementRequest": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "name": {
          "type": "string",
          "example": "GTA 5"
        },
        "webhook_url": {
          "type": "string",
          "example": "example.com"
        }
      }
    },
    "Event": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string",
          "example": "occurrs if a player killed an enemy"
        },
        "id": {
          "type": "string",
          "example": "f74687fa-2df7-450e-8c31-993695dcebf7"
        },
        "sku": {
          "type": "string",
          "example": "enemy_kill"
        }
      }
    },
    "EventCount": {
      "type": "object",
      "properties": {
        "count": {
          "type": "integer",
          "format": "int64"
        },
        "event_id": {
          "type": "string",
          "example": "f74687fa-2df7-450e-8c31-993695dcebf7"
        }
      }
    },
    "Player": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "example": "5e6890cc-c3a1-4c26-b269-b55e715c09b8"
        },
        "wallet_address": {
          "type": "string",
          "example": "ba4TrH8sNRzQfFhJH8mR"
        }
      }
    },
    "Project": {
      "type": "object",
      "required": [
        "name",
        "webhook_url"
      ],
      "properties": {
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "name": {
          "type": "string",
          "example": "GTA 5"
        },
        "webhook_url": {
          "type": "string",
          "example": "example.com"
        }
      }
    },
    "ProjectRequest": {
      "type": "object",
      "required": [
        "name",
        "webhook_url"
      ],
      "properties": {
        "name": {
          "type": "string",
          "example": "GTA 5"
        },
        "webhook_url": {
          "type": "string",
          "example": "example.com"
        }
      }
    },
    "ProjectResponse": {
      "type": "object",
      "properties": {
        "api_key": {
          "type": "string",
          "example": "abcdef123"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        }
      }
    },
    "Trigger": {
      "type": "object",
      "properties": {
        "achievement_id": {
          "type": "string",
          "example": "3a24bc5c-4c9d-11ed-bdc3-0242ac120002"
        },
        "events": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/EventCount"
          }
        },
        "id": {
          "type": "string",
          "example": "5e6890cc-c3a1-4c26-b269-b55e715c09b8"
        },
        "sku": {
          "type": "string",
          "example": "enemy_kill"
        }
      }
    },
    "WebhookEvent": {
      "type": "object",
      "properties": {
        "event_id": {
          "type": "string",
          "example": "5e6890cc-c3a1-4c26-b269-b55e715c09b8"
        },
        "player_id": {
          "type": "string",
          "example": "5e6890cc-c3a1-4c26-b269-b55e715c09b8"
        }
      }
    },
    "WebhookTrigger": {
      "type": "object",
      "properties": {
        "event_id": {
          "type": "string",
          "example": "5e6890cc-c3a1-4c26-b269-b55e715c09b8"
        },
        "player_id": {
          "type": "string",
          "example": "5e6890cc-c3a1-4c26-b269-b55e715c09b8"
        }
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "api_key",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Partner info",
      "name": "project"
    },
    {
      "name": "achievement"
    },
    {
      "name": "event"
    }
  ]
}`))
}
