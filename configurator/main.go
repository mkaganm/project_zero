package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os/exec"
)

func main() {

	// crate elasticsearch indexes
	createUserserviceIndex()
	createMailerserviceIndex()
	createLoggerserviceIndex()
	createCronitorIndex()

	// create postgres tables
	createPostgresDB()
	createPostgreTables()
}

// createUserserviceIndex is a function to create index for userservice in elasticsearch
func createUserserviceIndex() {

	cmd := exec.Command("curl", "-XPUT", "http://elasticsearch:9200/userservice", "-H", "Content-Type: application/json", "-d", `{
	  "settings": {
	    "number_of_shards": 1,
	    "number_of_replicas": 1
	  },
	  "mappings": {
	    "properties": {
	      "timestamp": {
	        "type": "date"
	      },
	      "data": {
	        "type": "object"
	      }
	    }
	  }
	}`)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Default().Println("err", err)
		return
	}

	log.Default().Println("Command output: ", string(output))
}

// createCronitorIndex is a function to create index for cronitor in elasticsearch
func createCronitorIndex() {

	cmd := exec.Command("curl", "-XPUT", "http://elasticsearch:9200/cronitor", "-H", "Content-Type: application/json", "-d", `{
	  "settings": {
	    "number_of_shards": 1,
	    "number_of_replicas": 1
	  },
	  "mappings": {
	    "properties": {
	      "timestamp": {
	        "type": "date"
	      },
	      "data": {
	        "type": "object"
	      }
	    }
	  }
	}`)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Default().Println("err", err)
		return
	}

	log.Default().Println("Command output: ", string(output))
}

// createMailerserviceIndex is a function to create index for mailerservice in elasticsearch
func createMailerserviceIndex() {

	cmd := exec.Command("curl", "-XPUT", "http://elasticsearch:9200/mailerservice", "-H", "Content-Type: application/json", "-d", `{
	  "settings": {
	    "number_of_shards": 1,
	    "number_of_replicas": 1
	  },
	  "mappings": {
	    "properties": {
	      "timestamp": {
	        "type": "date"
	      },
	      "data": {
	        "type": "object"
	      }
	    }
	  }
	}`)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Default().Println("err", err)
		return
	}

	log.Default().Println("Command output: ", string(output))
}

// createMailerserviceIndex is a function to create index for loggerservice in elasticsearch
func createLoggerserviceIndex() {

	cmd := exec.Command("curl", "-XPUT", "http://elasticsearch:9200/loggerservice", "-H", "Content-Type: application/json", "-d", `{
	  "settings": {
	    "number_of_shards": 1,
	    "number_of_replicas": 1
	  },
	  "mappings": {
	    "properties": {
	      "timestamp": {
	        "type": "date"
	      },
	      "data": {
	        "type": "object"
	      }
	    }
	  }
	}`)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Default().Println("err", err)
		return
	}

	log.Default().Println("Command output: ", string(output))
}

// createPostgresDB is function to create db
func createPostgresDB() {

	connStr := fmt.Sprintf("host=postgres_db port=5432 user=admin password=admin sslmode=disable")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Yeni bir veritabanı oluşturun
	_, err = db.Exec("CREATE DATABASE project_zero")
	if err != nil {
		log.Fatal(err)
	}

}

func createPostgreTables() {
	connStr := fmt.Sprintf("host=postgres_db port=5432 user=admin password=admin dbname =project_zero sslmode=disable")

	// Veritabanı bağlantısını açın
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Default().Println("Error when connecting to db :", err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		phone_number VARCHAR(255) NOT NULL,
		is_blocked BOOLEAN DEFAULT false,
		login_attempt_count INTEGER DEFAULT 0,
		is_verified BOOLEAN DEFAULT false,
		created_at TIMESTAMP DEFAULT NOW() NOT NULL,
		updated_at TIMESTAMP DEFAULT NOW() NOT NULL
	);`)
	if err != nil {
		log.Default().Println("Error when creating table :", err)
	}

	_, err = db.Exec(`CREATE OR REPLACE FUNCTION update_updated_at_column()
		RETURNS TRIGGER AS $$
		BEGIN
			NEW.updated_at = NOW();
			RETURN NEW;
		END;
		$$ language 'plpgsql';`)
	if err != nil {
		log.Default().Println("Error when creating table :", err)
	}

	// CREATE TRIGGER komutunu çalıştırın
	_, err = db.Exec(`CREATE TRIGGER update_users_updated_at
		BEFORE UPDATE ON users
		FOR EACH ROW
		EXECUTE FUNCTION update_updated_at_column();`)
	if err != nil {
		log.Default().Println("Error when creating table :", err)
	}

	// CREATE TABLE komutunu çalıştırın
	_, err = db.Exec(`CREATE TABLE verifications (
		id SERIAL PRIMARY KEY,
		user_id INTEGER REFERENCES users(id),
		verification_code_hash VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT NOW() NOT NULL
	);`)
	if err != nil {
		log.Default().Println("Error when creating table", err)
	}

}
