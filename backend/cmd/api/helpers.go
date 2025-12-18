package main

// TODO: Pagination
// Goal: Calculate SQL LIMIT and OFFSET based on query params (page, page_size).
// Logic:
//   1. Read "page" from URL query string (default to 1).
//   2. Read "page_size" from URL query string (default to 20).
//   3. Return a struct with Limit, Offset, and metadata for the response.
// Usage: input := app.readPagination(r); query += fmt.Sprintf("LIMIT %d OFFSET %d", input.limit, input.offset)

// TODO: Slug Generation
// Goal: Create SEO-friendly URLs from names.
// Logic: "Fluffy The Destroyer" -> "fluffy-the-destroyer"
// Usage: slug := app.createSlug(pet.Name); db.Insert(slug)

// TODO: JSON Responses
// Goal: Standardize how we send data back to the client.
// Logic: Set Content-Type header, encode payload, handle errors.
// Usage: app.writeJSON(w, http.StatusOK, data)
