package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"currency_converter/config"
	"currency_converter/database"
	"currency_converter/models"
)

type ConvertRequest struct {
	Amount float64 `json:"amount"`
}

type ConvertResponse struct {
	ConvertedAmount float64 `json:"converted_amount"`
}

func Convert(c *fiber.Ctx) error {
	baseCurrency := c.Query("base_currency")
	targetCurrency := c.Query("target_currency")
	if baseCurrency == "" || targetCurrency == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "base_currency and target_currency are required"})
	}

	var request ConvertRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	apiURL := fmt.Sprintf("%s%s", config.Cfg.ExternalAPI.URL, baseCurrency)
	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch conversion rate"})
	}
	defer resp.Body.Close()

	var apiResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to decode API response"})
	}

	rates := apiResponse["rates"].(map[string]interface{})
	conversionRate, ok := rates[targetCurrency].(float64)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid target currency"})
	}

	convertedAmount := request.Amount * conversionRate

	conversion := models.Conversion{
		BaseCurrency:    baseCurrency,
		TargetCurrency:  targetCurrency,
		Amount:          request.Amount,
		ConvertedAmount: convertedAmount,
	}

	database.DB.Create(&conversion)

	return c.JSON(ConvertResponse{ConvertedAmount: convertedAmount})
}
