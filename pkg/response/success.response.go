package response

import "github.com/gin-gonic/gin"

/*
	@Author: Sinh
	@Date: 2025/6/1
	@Description: This package provides a standardized way to handle success responses in the application.
*/

type SuccessResponse struct {
	StatusCode       int         `json:"statusCode"`
	Message          string      `json:"message"`
	ReasonStatusCode string      `json:"reasonStatusCode"`
	Metadata         interface{} `json:"metadata"`
}

func Ok(c *gin.Context, message string, metadata interface{}) {
	c.JSON(int(OK), SuccessResponse{
		StatusCode:       int(OK),
		Message:          message,
		ReasonStatusCode: MSG[OK],
		Metadata:         metadata,
	})
}

func Created(c *gin.Context, message string, metadata interface{}) {
	c.JSON(int(CREATED), SuccessResponse{
		StatusCode:       int(CREATED),
		Message:          message,
		ReasonStatusCode: MSG[CREATED],
		Metadata:         metadata,
	})
}

func Accepted(c *gin.Context, message string, metadata interface{}) {
	c.JSON(int(ACCEPTED), SuccessResponse{
		StatusCode:       int(ACCEPTED),
		Message:          message,
		ReasonStatusCode: MSG[ACCEPTED],
		Metadata:         metadata,
	})
}

func NoContent(c *gin.Context, message string, metadata interface{}) {
	c.JSON(int(NO_CONTENT), SuccessResponse{
		StatusCode:       int(NO_CONTENT),
		Message:          message,
		ReasonStatusCode: MSG[NO_CONTENT],
		Metadata:         metadata,
	})
}
