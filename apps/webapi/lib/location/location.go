package location

import (
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
)

var (
	cachedClient *geoip2.Reader
)

// Setup connects to a GeoIP database and caches the client.
func Setup(databasePath string) error {
	if cachedClient != nil {
		return nil
	}
	db, err := geoip2.Open(databasePath)
	if err != nil {
		return err
	}
	err = testConnection(db)
	if err != nil {
		return err
	}
	cachedClient = db
	return nil
}

// GetClient returns the connected GeoIP database reader.
// NOTE: The setup function must be called before this function.
func GetClient() *geoip2.Reader {
	if cachedClient == nil {
		log.Fatalln("Cannot to use GeoIP database before connecting")
	}
	return cachedClient
}

// testConnection tests the connection to the GeoIP database.
func testConnection(reader *geoip2.Reader) error {
	_, err := reader.City(net.ParseIP("1.1.1.1"))
	if err != nil {
		return err
	}
	return nil
}
