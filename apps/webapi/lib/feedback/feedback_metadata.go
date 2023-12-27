package feedback

import (
	"fmt"
	"net"
	"strings"

	"provar.se/webapi/lib/location"
)

const (
	MetadataRequestIP         = "request.ip"
	MetadataLocationCity      = "location.city"
	MetadataLocationCountry   = "location.country"
	MetadataLocationLatitude  = "location.latitude"
	MetadataLocationLongitude = "location.longitude"
	MetadataLocationAccuracy  = "location.accuracy"
)

// RequestHeadersToMask are the headers that should be masked in the feedback
// metadata stored in the database.
var RequestHeadersToMask = []string{
	"authorization",
}

// FeedbackMeta are additional metadata attached to feedback messages
type FeedbackMeta map[string]interface{}

// SetRequestIP sets the "request.ip" metadata field.
func (m *FeedbackMeta) SetRequestIP(value string) {
	(*m)[MetadataRequestIP] = value
}

// SetRequestHeaders sets the "request.header.*" header fields.
func (m *FeedbackMeta) SetRequestHeaders(headers map[string][]string) {
	for key, val := range headers {
		normalizedKey := strings.ToLower(key)
		metadataVal := strings.Join(val, ",")
		metadataKey := fmt.Sprintf("request.header.%s", normalizedKey)
		if shouldMaskHeader(normalizedKey) {
			metadataVal = "*****"
		}
		(*m)[metadataKey] = metadataVal
	}
}

// SetLocationFromIP sets location metadata fields based on the IP address
// of the feedback message using the "request.ip" metadata field.
func (m *FeedbackMeta) SetLocationFromIP() error {
	maybeIP, ok := (*m)[MetadataRequestIP]
	if !ok {
		return nil
	}
	IP, ok := maybeIP.(string)
	if !ok {
		return nil
	}
	fmt.Println(IP)
	rec, err := location.GetClient().City(net.ParseIP(IP))
	if err != nil {
		fmt.Println(err)
		return err
	}
	m.setIfMissing(MetadataRequestIP, IP)
	m.setIfMissing(MetadataLocationCity, rec.City.Names["en"])
	m.setIfMissing(MetadataLocationCountry, rec.Country.IsoCode)
	m.setIfMissing(MetadataLocationLatitude, rec.Location.Latitude)
	m.setIfMissing(MetadataLocationLongitude, rec.Location.Longitude)
	m.setIfMissing(MetadataLocationAccuracy, float64(rec.Location.AccuracyRadius))
	return nil
}

// setIfMissing sets a key/value pair in the metadata if not already set.
func (m *FeedbackMeta) setIfMissing(key string, val interface{}) {
	_, ok := (*m)[key]
	if !ok {
		(*m)[key] = val
	}
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
