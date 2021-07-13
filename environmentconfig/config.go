package environmentconfig

import (
	"bufio"
	"os"
	"strings"
)

// EnvironmentConfig represents environment variables in .env file.
type EnvironmentConfig struct {
	// general
	AppName string `envconfig:"APP_NAME" required:"true" default:"voucher pool"`

	// Database
	DBDriver string `envconfig:"DB_DRIVER" required:"true" default:"mysql"`
	DBUser   string `envconfig:"DB_USER" required:"true" default:"root"`
	DBPass   string `envconfig:"DB_PASS" required:"true"`
	DBName   string `envconfig:"DB_NAME" required:"true"`
	DBHost   string `envconfig:"DB_HOST" required:"true" default:"127.0.0.1"`
	DBPort   string `envconfig:"DB_PORT" required:"true" default:"3306"`
}

// Export will export environment variables from given file.
func Export(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "export ") {
			line = line[7:]
		}

		token := strings.SplitN(line, "=", 2)
		// Remove spaces, ' and "
		key := strings.Trim(strings.Trim(strings.TrimSpace(token[0]), "'"), "\"")
		value := strings.Trim(strings.Trim(strings.TrimSpace(token[1]), "'"), "\"")

		os.Setenv(key, value)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
