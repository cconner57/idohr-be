package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func parseDate(dateStr string) *time.Time {
	dateStr = strings.TrimSpace(dateStr)
	if dateStr == "" || strings.Contains(dateStr, "Due") || dateStr == "----" {
		return nil
	}
	layout := "1/2/2006"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return nil
	}
	return &t
}

func parseMoney(moneyStr string) float64 {
	cleaned := strings.ReplaceAll(strings.ReplaceAll(moneyStr, "$", ""), ",", "")
	cleaned = strings.TrimSpace(cleaned)
	if cleaned == "----" || cleaned == "" {
		return 0.0
	}
	val, _ := strconv.ParseFloat(cleaned, 64)
	return val
}

func main() {
	godotenv.Load(".env")
	db, err := sql.Open("postgres", os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fileName := "Master Pet List - Cats Master List.csv"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Missing file: %s", fileName)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	records, _ := reader.ReadAll()

	fmt.Println("Starting Import...")

	for i, row := range records {
		// Skip header and empty rows
		if i == 0 || len(row) < 3 || row[0] == "" || row[0] == "ID" {
			continue
		}

		id, err := strconv.Atoi(row[0])
		if err != nil {
			continue
		}

		// Grouping traits for JSONB columns
		phys, _ := json.Marshal(map[string]interface{}{"coat": row[6], "color": row[7], "size": row[11]})
		beh, _ := json.Marshal(map[string]interface{}{"energy": row[17], "cats": row[18], "dogs": row[19], "kids": row[20]})
		med, _ := json.Marshal(map[string]interface{}{"chip": row[12], "vax": row[32], "needs": row[25], "meds": row[46]})
		adp, _ := json.Marshal(map[string]interface{}{"fee": parseMoney(row[37]), "date": parseDate(row[38]), "adopted": row[39]})

		query := `
			INSERT INTO pets (id, name, date_of_birth, breed, sex, microchip_id, status, litter_name, physical, behavior, medical, adoption)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
			ON CONFLICT (id) DO UPDATE SET
				name = EXCLUDED.name, 
				status = EXCLUDED.status, 
				litter_name = EXCLUDED.litter_name,
				physical = EXCLUDED.physical, 
				behavior = EXCLUDED.behavior, 
				medical = EXCLUDED.medical, 
				adoption = EXCLUDED.adoption, 
				updated_at = NOW();`

		_, err = db.Exec(query, id, row[2], parseDate(row[3]), row[5], row[10], row[13], row[36], row[4], phys, beh, med, adp)
		if err != nil {
			log.Printf("Error ID %d (%s): %v", id, row[2], err)
		}
	}
	fmt.Println("Seeding complete! Database is now synced with CSV.")
}
