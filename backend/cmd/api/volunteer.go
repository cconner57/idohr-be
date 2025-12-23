package main

import (
	"fmt"
	"net/http"

	"github.com/cconner57/adoption-os/backend/internal/data"
	"github.com/cconner57/adoption-os/backend/internal/validator"
)

func (app *application) submitVolunteerApplication(w http.ResponseWriter, r *http.Request) {
	var input data.VolunteerApplication

	// Read JSON from frontend
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Validate
	v := validator.New()
	data.ValidateVolunteerApplication(v, &input)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Log success (Task requirement: console log validation success)
	fmt.Printf("Volunteer Application Validated Successfully: %+v\n", input)

	// Send Email Notification
	// Runs in a goroutine to avoid blocking the response
	go func() {
		defer func() {
			if err := recover(); err != nil {
				app.logger.Println("Panic in email goroutine:", err)
			}
		}()

		subject := fmt.Sprintf("New Volunteer Application: %s %s", input.FirstName, input.LastName)
		body := fmt.Sprintf(`
New Volunteer Application Received!

Name: %s %s
Address: %s, %s, %s
Phone: %s
Birthday: %s
Age: %d

Emergency Contact:
%s (%s)

Interest Reason:
%s

This application was submitted via the Adoption OS Volunteer Form.
`, input.FirstName, input.LastName, input.Address, input.City, input.Zip, input.PhoneNumber, input.Birthday, safeInt(input.Age), input.EmergencyContactName, input.EmergencyContactPhone, input.InterestReason)

		// Send to the configured sender address (acting as Admin)
		// Or you could read a specific recipient from config.
		recipient := app.config.smtp.sender
		if recipient == "" {
			recipient = "admin@example.com" // Fallback
		}

		err := app.mailer.Send(recipient, subject, body)
		if err != nil {
			app.logger.Println("Failed to send email notification:", err)
		} else {
			app.logger.Println("Email notification sent successfully to", recipient)
		}
	}()

	// TODO: Insert into DB (Reserved for next task)
	// err = app.models.Volunteers.Insert(&input)
	// if err != nil {
	// 	app.serverErrorResponse(w, r, err)
	// 	return
	// }

	// Send back a success response
	err = app.writeJSON(w, http.StatusOK, envelope{"status": "success", "message": "Application received and validated"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
