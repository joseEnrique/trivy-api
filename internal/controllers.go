package internal

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetVulnerability(ctx *fiber.Ctx) error {
	var results Results
	image := ctx.Query("image")
	results, err := trivy(image)
	if err != nil {
		fmt.Printf("Error parsing JSON output: %s\n", err)
		return errors.New("error parsing json output")
	}
	return ctx.JSON(results.VulnerabilityReports[0])

}
