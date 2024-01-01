package feedback

import (
	"fmt"
	"strings"

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
		metadataKey := fmt.Sprintf("request-header-%s", normalizedKey)
		if shouldMaskHeader(normalizedKey) {
			metadataVal = "*****"
		}
		(*m)[metadataKey] = metadataVal
	}
}

// SetMetadataFromIP sets location metadata fields based on the IP address
// of the feedback message using the "request.ip" metadata field.
func (m *FeedbackMeta) SetMetadataFromIP() error {
	maybeIP, ok := (*m)[MetadataRequestIP]
	if !ok {
		return nil
	}
	IP, ok := maybeIP.(string)
	if !ok {
		return nil
	}
	rec, err := metadata.GetLocation(IP)
	if err != nil {
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

// SetMetadataFromUA sets metadata fields based on the user agent string.
func (m *FeedbackMeta) SetMetadataFromUA() error {
	maybeUA, ok := (*m)[MetadataRequestHeaderUA]
	if !ok {
		return nil
	}
	UA, ok := maybeUA.(string)
	if !ok {
		return nil
	}
	ua := metadata.ParseUserAgent(UA)
	m.setIfMissing(MetadataRequestBrowser, ua.Name)
	m.setIfMissing(MetadataRequestBrowserVersion, ua.Version)
	m.setIfMissing(MetadataRequestOS, ua.OS)
	m.setIfMissing(MetadataRequestOSVersion, ua.OSVersion)
	m.setIfMissing(MetadataRequestDevice, ua.Device)
	if ua.Mobile {
		m.setIfMissing(MetadataRequestDeviceType, "mobile")
	} else if ua.Tablet {
		m.setIfMissing(MetadataRequestDeviceType, "tablet")
	} else if ua.Desktop {
		m.setIfMissing(MetadataRequestDeviceType, "desktop")
	} else if ua.Bot {
		m.setIfMissing(MetadataRequestDeviceType, "bot")
	}
	return nil
}

// setIfMissing sets a key/value pair in the metadata if not already set.
func (m *FeedbackMeta) setIfMissing(key string, val interface{}) {
	_, ok := (*m)[key]
	if ok {
		return
	}
	if isEmptyValue(val) {
		return
	}
	(*m)[key] = val
}

func isEmptyValue(val interface{}) bool {
	switch v := val.(type) {
	case int:
		return v == 0
	case string:
		return v == ""
	}
	return false
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
