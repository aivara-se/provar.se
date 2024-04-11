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

// SetRequestIP sets the "request-ip" metadata field based on the IP address.
func SetRequestIP(m *map[string]string, ip string) {
	if ip == "" {
		return
	}
	setIfMissing(m, MetadataRequestIP, ip)
}

// SetIPLocation sets the location metadata fields based on the IP address.
func SetIPLocation(m *map[string]string, ip string) {
	rec, err := metadata.GetLocation(ip)
	if err != nil {
		return
	}
	setIfMissing(m, MetadataLocationCity, rec.City.Names["en"])
	setIfMissing(m, MetadataLocationCountry, rec.Country.IsoCode)
	setIfMissing(m, MetadataLocationLatitude, fmt.Sprintf("%f", rec.Location.Latitude))
	setIfMissing(m, MetadataLocationLongitude, fmt.Sprintf("%f", rec.Location.Longitude))
	setIfMissing(m, MetadataLocationAccuracy, fmt.Sprintf("%d", rec.Location.AccuracyRadius))
}

// SetRawHeaders sets the request headers as metadata fields on the feedback.
func SetRawHeaders(m *map[string]string, headers map[string][]string) {
	for key, val := range headers {
		normalizedKey := strings.ToLower(key)
		metadataVal := strings.Join(val, ",")
		metadataKey := fmt.Sprintf("request-header-%s", normalizedKey)
		if _, ok := (*m)[metadataKey]; ok {
			continue
		}
		if !shouldMaskHeader(normalizedKey) {
			(*m)[metadataKey] = metadataVal
		}
	}
}

// SetUserAgent sets metadata fields based on the user agent string.
func SetUserAgent(m *map[string]string) {
	UA, ok := (*m)[MetadataRequestHeaderUA]
	if !ok {
		return
	}
	ua := useragent.Parse(UA)
	setIfMissing(m, MetadataRequestBrowser, ua.Name)
	setIfMissing(m, MetadataRequestBrowserVersion, ua.Version)
	setIfMissing(m, MetadataRequestOS, ua.OS)
	setIfMissing(m, MetadataRequestOSVersion, ua.OSVersion)
	setIfMissing(m, MetadataRequestDevice, ua.Device)
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
