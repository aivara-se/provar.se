package metadata

import (
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
	"provar.se/webapi/lib/config"
)

var (
	cachedClient *geoip2.Reader
)

// Setup connects to a GeoIP database and caches the client.
func Setup() error {
	if cachedClient != nil {
		return nil
	}
	cfg := config.Get()
	db, err := geoip2.Open(cfg.Geolite2)
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

// GetLocation returns the location of an IP address.
func GetLocation(IP string) (*geoip2.City, error) {
	return getLocationClient().City(net.ParseIP(IP))
}

// getLocationClient returns the connected GeoIP database reader.
// NOTE: The setup function must be called before this function.
func getLocationClient() *geoip2.Reader {
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
