package com.quadstingray.speedtest.model

case class GeoIpLocation(
                          ip: String,
                          city: String,
                          region: String,
                          region_code: String,
                          country: String,
                          country_name: String,
                          continent_code: String,
                          in_eu: Boolean,
                          postal: String,
                          latitude: Double,
                          longitude: Double,
                          timezone: String,
                          utc_offset: String,
                          country_calling_code: String,
                          currency: String,
                          languages: String,
                          asn: String,
                          org: String
                        )
