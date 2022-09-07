package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"x-learn/mq/kafkahttpd/kproducer"
)

func createAccountLog(c *fiber.Ctx) error {
	// Instantiate new Message struct
	cmt := new(Comment)
	if err := c.BodyParser(cmt); err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}

	// convert body into bytes and send it to kafka
	cmtInBytes, err := json.Marshal(cmt)
	producer := kproducer.NewKProducer(brokersUrl)
	err = producer.PushMessageToQueue("comments", cmtInBytes)
	if err != nil {
		return err
	}

	// Return Comment in JSON format
	err = c.JSON(&fiber.Map{
		"success": true,
		"message": "Comment pushed successfully",
		"comment": cmt,
	})
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating product",
		})
		return err
	}
	return err
}