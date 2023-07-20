// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health_check": {
            "get": {
                "description": "A health checking api to make sure the server is not dead.",
                "tags": [
                    "Health"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/questionnaires": {
            "get": {
                "description": "Get all questionnaires from the database.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Questionnaire"
                ],
                "summary": "Get Questionnaires",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handlers.Questionnaire"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new questionnaire and optionally its questions.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Questionnaire"
                ],
                "summary": "Create Questionnaire",
                "parameters": [
                    {
                        "description": "The questionnaire to be created.",
                        "name": "questionnaire",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.QuestionnaireBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Questionnaire"
                        }
                    }
                }
            }
        },
        "/questionnaires/{id}": {
            "get": {
                "description": "Get a questionnaire by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Questionnaire"
                ],
                "summary": "Get Questionnaire",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The questionnaire's ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Questionnaire"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a questionnaire by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Questionnaire"
                ],
                "summary": "Delete Questionnaire",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The questionnaire's ID.",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/questionnaires/{id}/new/question": {
            "post": {
                "description": "Create a new question for a specified questionnaire",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Questionnaire"
                ],
                "summary": "Create Question",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The questionnaire's ID.",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The question to be created.",
                        "name": "question",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.QuestionBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Question"
                        }
                    }
                }
            }
        },
        "/questionnaires/{id}/new/response": {
            "post": {
                "description": "Create a new response for a specified questionnaire",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Questionnaire"
                ],
                "summary": "Create Response",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The questionnaire's ID.",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The response to be created.",
                        "name": "response",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.QuestionnaireResponseBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.QuestionnaireResponse"
                        }
                    }
                }
            }
        },
        "/questions": {
            "get": {
                "description": "Get all questions from the database.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "Get Questions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handlers.Question"
                            }
                        }
                    }
                }
            }
        },
        "/questions/{id}": {
            "get": {
                "description": "Get a question by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "Get Question",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The question's ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Question"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a question by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "Delete Question",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The question's ID.",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/responses": {
            "get": {
                "description": "Get all responses from the database.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Response"
                ],
                "summary": "Get Responses",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handlers.QuestionnaireResponse"
                            }
                        }
                    }
                }
            }
        },
        "/responses/{id}": {
            "get": {
                "description": "Get a response by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Response"
                ],
                "summary": "Get Response",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The response's ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.QuestionnaireResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a response by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Response"
                ],
                "summary": "Delete Response",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The response's ID.",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Get all users from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get Users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.User"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create User",
                "parameters": [
                    {
                        "description": "The user to be created",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Get a user by an Auth0 ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The user's Auth0 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    }
                }
            },
            "put": {
                "description": "update a user with specified Auth0 ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The user's Auth0 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "user to be updated",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a user by Auth0 ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The user's Auth0 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "ent.Answer": {
            "type": "object",
            "properties": {
                "body": {
                    "description": "Body holds the value of the \"body\" field.",
                    "type": "string"
                },
                "created_at": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                }
            }
        },
        "ent.Question": {
            "type": "object",
            "properties": {
                "body": {
                    "description": "Body holds the value of the \"body\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "type": {
                    "description": "Type holds the value of the \"type\" field.",
                    "type": "string"
                }
            }
        },
        "ent.Questionnaire": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                }
            }
        },
        "ent.User": {
            "type": "object",
            "properties": {
                "birth_year": {
                    "description": "BirthYear holds the value of the \"birth_year\" field.",
                    "type": "integer"
                },
                "created_at": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "demented_among_direct_relatives": {
                    "description": "DementedAmongDirectRelatives holds the value of the \"demented_among_direct_relatives\" field.",
                    "type": "boolean"
                },
                "ear_condition": {
                    "description": "EarCondition holds the value of the \"ear_condition\" field.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/user.EarCondition"
                        }
                    ]
                },
                "education_level": {
                    "description": "EducationLevel holds the value of the \"education_level\" field.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/user.EducationLevel"
                        }
                    ]
                },
                "eyesight_condition": {
                    "description": "EyesightCondition holds the value of the \"eyesight_condition\" field.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/user.EyesightCondition"
                        }
                    ]
                },
                "first_name": {
                    "description": "FirstName holds the value of the \"first_name\" field.",
                    "type": "string"
                },
                "gender": {
                    "description": "Gender holds the value of the \"gender\" field.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/user.Gender"
                        }
                    ]
                },
                "head_injury_experience": {
                    "description": "HeadInjuryExperience holds the value of the \"head_injury_experience\" field.",
                    "type": "boolean"
                },
                "height": {
                    "description": "Height holds the value of the \"height\" field.",
                    "type": "number"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "last_name": {
                    "description": "LastName holds the value of the \"last_name\" field.",
                    "type": "string"
                },
                "marriage": {
                    "description": "Marriage holds the value of the \"marriage\" field.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/user.Marriage"
                        }
                    ]
                },
                "medical_history": {
                    "description": "MedicalHistory holds the value of the \"medical_history\" field.",
                    "type": "string"
                },
                "medication_status": {
                    "description": "MedicationStatus holds the value of the \"medication_status\" field.",
                    "type": "string"
                },
                "occupation": {
                    "description": "Occupation holds the value of the \"occupation\" field.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/user.Occupation"
                        }
                    ]
                },
                "smoking_habit": {
                    "description": "SmokingHabit holds the value of the \"smoking_habit\" field.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/user.SmokingHabit"
                        }
                    ]
                },
                "updated_at": {
                    "description": "UpdatedAt holds the value of the \"updated_at\" field.",
                    "type": "string"
                },
                "weight": {
                    "description": "Weight holds the value of the \"weight\" field.",
                    "type": "number"
                }
            }
        },
        "handlers.AnswerBody": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "question_id": {
                    "type": "string"
                }
            }
        },
        "handlers.Question": {
            "type": "object",
            "properties": {
                "body": {
                    "description": "Body holds the value of the \"body\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "quesionnaires": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Questionnaire"
                    }
                },
                "type": {
                    "description": "Type holds the value of the \"type\" field.",
                    "type": "string"
                }
            }
        },
        "handlers.QuestionBody": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "handlers.Questionnaire": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                },
                "questions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Question"
                    }
                },
                "responses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handlers.QuestionnaireResponse"
                    }
                }
            }
        },
        "handlers.QuestionnaireBody": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "questions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handlers.QuestionBody"
                    }
                }
            }
        },
        "handlers.QuestionnaireResponse": {
            "type": "object",
            "properties": {
                "answers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Answer"
                    }
                },
                "created_at": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                }
            }
        },
        "handlers.QuestionnaireResponseBody": {
            "type": "object",
            "properties": {
                "answers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handlers.AnswerBody"
                    }
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "user.EarCondition": {
            "type": "string",
            "enum": [
                "normal",
                "slightly_affecting_conversation",
                "need_hearing_aid"
            ],
            "x-enum-varnames": [
                "EarConditionNormal",
                "EarConditionSlightlyAffectingConversation",
                "EarConditionNeedHearingAid"
            ]
        },
        "user.EducationLevel": {
            "type": "string",
            "enum": [
                "elementry_school_or_below",
                "middle_school",
                "high_school",
                "bachelor",
                "master",
                "doctorate"
            ],
            "x-enum-varnames": [
                "EducationLevelElementrySchoolOrBelow",
                "EducationLevelMiddleSchool",
                "EducationLevelHighSchool",
                "EducationLevelBachelor",
                "EducationLevelMaster",
                "EducationLevelDoctorate"
            ]
        },
        "user.EyesightCondition": {
            "type": "string",
            "enum": [
                "normal",
                "slightly_affecting_reading",
                "need_glasses"
            ],
            "x-enum-varnames": [
                "EyesightConditionNormal",
                "EyesightConditionSlightlyAffectingReading",
                "EyesightConditionNeedGlasses"
            ]
        },
        "user.Gender": {
            "type": "string",
            "enum": [
                "male",
                "female",
                "nonbinary"
            ],
            "x-enum-varnames": [
                "GenderMale",
                "GenderFemale",
                "GenderNonbinary"
            ]
        },
        "user.Marriage": {
            "type": "string",
            "enum": [
                "single",
                "married",
                "divorced",
                "widowed"
            ],
            "x-enum-varnames": [
                "MarriageSingle",
                "MarriageMarried",
                "MarriageDivorced",
                "MarriageWidowed"
            ]
        },
        "user.Occupation": {
            "type": "string",
            "enum": [
                "student",
                "government_employee",
                "service_industry",
                "industry_and_commerce",
                "freelancer",
                "domestic"
            ],
            "x-enum-varnames": [
                "OccupationStudent",
                "OccupationGovernmentEmployee",
                "OccupationServiceIndustry",
                "OccupationIndustryAndCommerce",
                "OccupationFreelancer",
                "OccupationDomestic"
            ]
        },
        "user.SmokingHabit": {
            "type": "string",
            "enum": [
                "none",
                "sometimes",
                "everyday"
            ],
            "x-enum-varnames": [
                "SmokingHabitNone",
                "SmokingHabitSometimes",
                "SmokingHabitEveryday"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "health-statistic.dechnology.com.tw",
	BasePath:         "/api/v1",
	Schemes:          []string{"http", "https"},
	Title:            "Health Statistic API",
	Description:      "This is a sample server for Health Statistic API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
