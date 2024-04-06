package feedback

import (
	"fmt"
	"strings"

	"github.com/mileusna/useragent"
	"provar.se/webapi/lib/metadata"
)

const (
	MetadataRequestIP             = "request-ip"
	MetadataRequestHeaderUA       = "request-header-user-agent"
	MetadataRequestBrowser        = "request-device-browser"
	MetadataRequestBrowserVersion = "request-device-browser-version"
	MetadataRequestOS             = "request-device-os"
	MetadataRequestOSVersion      = "request-device-os-version"
	MetadataRequestDevice         = "request-device-name"
	MetadataRequestDeviceType     = "request-device-type"
	MetadataLocationCity          = "request-location-city"
	MetadataLocationCountry       = "request-location-country"
	MetadataLocationLatitude      = "request-location-latitude"
	MetadataLocationLongitude     = "request-location-longitude"
	MetadataLocationAccuracy      = "request-location-accuracy"
)

// RequestHeadersToMask are the headers that should be masked in the feedback
// metadata stored in the database.
var RequestHeadersToMask = []string{
	"authorization",
}

// SetRequestIP sets the "request.ip" metadata field.
func SetRequestIP(m *map[string]string, value string) {
	(*m)[MetadataRequestIP] = value
}

// SetRequestHeaders sets the "request.header.*" header fields.
func SetRequestHeaders(m *map[string]string, headers map[string][]string) {
	for key, val := range headers {
		normalizedKey := strings.ToLower(key)
		metadataVal := strings.Join(val, ",")
		metadataKey := fmt.Sprintf("request-header-%s", normalizedKey)
		if shouldMaskHeader(normalizedKey) {
			metadataVal = "*****"
		}
		(*m)[metadataKey] = metadataVal
	}
}

// SetMetadataFromIP sets location metadata fields based on the IP address
// of the feedback message using the "request.ip" metadata field.
func SetMetadataFromIP(m *map[string]string) error {
	IP, ok := (*m)[MetadataRequestIP]
	if !ok {
		return nil
	}
	rec, err := metadata.GetLocation(IP)
	if err != nil {
		return err
	}
	setIfMissing(m, MetadataRequestIP, IP)
	setIfMissing(m, MetadataLocationCity, rec.City.Names["en"])
	setIfMissing(m, MetadataLocationCountry, rec.Country.IsoCode)
	setIfMissing(m, MetadataLocationLatitude, fmt.Sprintf("%f", rec.Location.Latitude))
	setIfMissing(m, MetadataLocationLongitude, fmt.Sprintf("%f", rec.Location.Longitude))
	setIfMissing(m, MetadataLocationAccuracy, fmt.Sprintf("%d", rec.Location.AccuracyRadius))
	return nil
}

// SetMetadataFromUA sets metadata fields based on the user agent string.
func SetMetadataFromUA(m *map[string]string) error {
	UA, ok := (*m)[MetadataRequestHeaderUA]
	if !ok {
		return nil
	}
	ua := useragent.Parse(UA)
	setIfMissing(m, MetadataRequestBrowser, ua.Name)
	setIfMissing(m, MetadataRequestBrowserVersion, ua.Version)
	setIfMissing(m, MetadataRequestOS, ua.OS)
	setIfMissing(m, MetadataRequestOSVersion, ua.OSVersion)
	setIfMissing(m, MetadataRequestDevice, ua.Device)
	if ua.Mobile {
		setIfMissing(m, MetadataRequestDeviceType, "mobile")
	} else if ua.Tablet {
		setIfMissing(m, MetadataRequestDeviceType, "tablet")
	} else if ua.Desktop {
		setIfMissing(m, MetadataRequestDeviceType, "desktop")
	} else if ua.Bot {
		setIfMissing(m, MetadataRequestDeviceType, "bot")
	}
	return nil
}

// setIfMissing sets a key/value pair in the metadata if not already set.
func setIfMissing(m *map[string]string, key string, val string) {
	_, ok := (*m)[key]
	if ok {
		return
	}
	(*m)[key] = val
}

// shouldMaskHeader returns true if the header should be masked.
func shouldMaskHeader(key string) bool {
	for _, header := range RequestHeadersToMask {
		if header == key {
			return true
		}
	}
	return false
}
