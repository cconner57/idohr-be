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

// Helpers
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

	// --- SELF-HEALING: Ensure Table Exists ---
	fmt.Println("Checking database schema...")
	setupSQL := `
	CREATE TABLE IF NOT EXISTS pets (
		id INT PRIMARY KEY,
		name TEXT NOT NULL,
		date_of_birth DATE,
		breed TEXT,
		sex TEXT,
		microchip_id TEXT,
		status TEXT,
		litter_name TEXT,
		physical JSONB DEFAULT '{}',
		behavior JSONB DEFAULT '{}',
		medical JSONB DEFAULT '{}',
		adoption JSONB DEFAULT '{}',
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	);`
	if _, err := db.Exec(setupSQL); err != nil {
		log.Fatal("Failed to ensure table exists: ", err)
	}

	// --- IMPORT LOGIC ---
	fileName := "Master Pet List - Cats Master List.csv"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Missing file: backend/%s", fileName)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	records, _ := reader.ReadAll()

	fmt.Printf("Importing records from %s...\n", fileName)

	for i, row := range records {
		if i == 0 || len(row) < 10 || row[0] == "" {
			continue
		}

		id, _ := strconv.Atoi(row[0])

		// Map JSON Blocks
		phys, _ := json.Marshal(map[string]interface{}{"coat": row[6], "color": row[7], "size": row[11]})
		beh, _ := json.Marshal(map[string]interface{}{"energy": row[17], "cats": row[18], "dogs": row[19], "kids": row[20]})
		med, _ := json.Marshal(map[string]interface{}{"chip": row[12], "vax": row[32], "needs": row[25], "meds": row[46]})
		adp, _ := json.Marshal(map[string]interface{}{"fee": parseMoney(row[37]), "date": parseDate(row[38]), "adopted": row[39]})

		query := `
			INSERT INTO pets (id, name, date_of_birth, breed, sex, microchip_id, status, litter_name, physical, behavior, medical, adoption)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
			ON CONFLICT (id) DO UPDATE SET
				name = EXCLUDED.name, status = EXCLUDED.status, physical = EXCLUDED.physical, 
				behavior = EXCLUDED.behavior, medical = EXCLUDED.medical, adoption = EXCLUDED.adoption, 
				updated_at = NOW();`

		_, err = db.Exec(query, id, row[2], parseDate(row[3]), row[5], row[10], row[13], row[36], row[4], phys, beh, med, adp)
		if err != nil {
			log.Printf("Skip ID %d: %v", id, err)
		}
	}
	fmt.Println("Seeding complete!")
}
