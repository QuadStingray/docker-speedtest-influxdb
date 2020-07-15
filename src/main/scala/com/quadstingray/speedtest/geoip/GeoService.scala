package com.quadstingray.speedtest.geoip

import java.net.URI

import com.github.plokhotnyuk.jsoniter_scala.core.{ readFromString, JsonValueCodec }
import com.github.plokhotnyuk.jsoniter_scala.macros.JsonCodecMaker
import com.quadstingray.speedtest.model.{ GeoIpLocation, Location }
import okhttp3.{ OkHttpClient, Request }

case class GeoService() {
  private val averageRadiusOfEarthKm                             = 6371.0
  implicit private val geoIpCodec: JsonValueCodec[GeoIpLocation] = JsonCodecMaker.make[GeoIpLocation]

  def calculateDistanceInKilometer(userLocation: Location, warehouseLocation: Location): Double = {
    val latDistance: Double = Math.toRadians(userLocation.latitude - warehouseLocation.latitude)
    val lngDistance: Double = Math.toRadians(userLocation.longitude - warehouseLocation.longitude)
    val sinLat: Double      = Math.sin(latDistance / 2)
    val sinLng: Double      = Math.sin(lngDistance / 2)
    val a: Double           = sinLat * sinLat + (Math.cos(Math.toRadians(userLocation.latitude)) * Math.cos(Math.toRadians(warehouseLocation.latitude)) * sinLng * sinLng)
    val c: Double           = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a))
    averageRadiusOfEarthKm * c
  }

  def locateUser: GeoIpLocation = {
    val uri            = new URI("https://ipapi.co/json/")
    val client         = new OkHttpClient.Builder().build()
    val request        = new Request.Builder().url(uri.toString).build()
    val responseString = client.newCall(request).execute().body().string()
    readFromString[GeoIpLocation](responseString)
  }

  def locateIp(ip: String): GeoIpLocation = {
    val uri            = new URI("https://ipapi.co/" + ip + "/json/")
    val client         = new OkHttpClient.Builder().build()
    val request        = new Request.Builder().url(uri.toString).build()
    val responseString = client.newCall(request).execute().body().string()
    readFromString[GeoIpLocation](responseString)
  }

}
