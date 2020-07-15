package com.quadstingray.speedtest

import java.util.Date

import com.github.plokhotnyuk.jsoniter_scala.core._
import com.github.plokhotnyuk.jsoniter_scala.macros.JsonCodecMaker
import com.quadstingray.speedtest.converter.ResultConverter
import com.quadstingray.speedtest.geoip.GeoService
import com.quadstingray.speedtest.model.{ ClientInfos, GeoIpLocation, Location, TestResult }
import com.quadstingray.speedtest.ndt7.{ ServerClient, SpeedTest }

import scala.concurrent.duration.Duration

object Application extends App {

  private val geoService: GeoService           = GeoService()
  private val speedTestClient                  = ServerClient()
  private val resultConverter: ResultConverter = ResultConverter()

  val outputFormat: String = System.getProperty("output")
  var intervalString       = System.getProperty("interval")
  var influxUrl            = System.getProperty("influx_url")
  var influxUser           = System.getProperty("influx_user")
  var influxPassword       = System.getProperty("influx_password")
  var influxDB             = System.getProperty("influx_db")
  var speedtestHost             = System.getProperty("host")

  private val influxDb = database.InfluxDb(influxUrl, influxUser, influxPassword, influxDB)

  private var interval: Duration  = Duration("1ms")
  private var repeatable: Boolean = false

  try {
    interval = Duration(intervalString)
    repeatable = true
  } catch {
    case e: Exception => repeatable = false
  }

  val geoIpLocation: GeoIpLocation = geoService.locateUser

  do {
    val nextRun: Long = try {
      val server        = speedTestClient.nextServer
      val serverDetails = speedTestClient.serverDetailsBySite(server.site)

      val speedTestResult = SpeedTest.runTest(Some(server))
      val distance = geoService.calculateDistanceInKilometer(Location(geoIpLocation.latitude, geoIpLocation.longitude),
                                                             Location(serverDetails.get.latitude, serverDetails.get.longitude))

      val clientInfos = ClientInfos(
        geoIpLocation.ip,
        geoIpLocation.org,
        geoIpLocation.city,
        distance
      )

      outputFormat match {
        case "JSON" => {
          implicit val codec: JsonValueCodec[TestResult] = JsonCodecMaker.make[TestResult]
          val result                                     = resultConverter.convertToResult(speedTestResult, clientInfos)
          println(writeToString(result, WriterConfig.withIndentionStep(5)))
        }
        case "INFLUX" =>
          influxDb.init()
          val point = resultConverter.convertToPoint(speedTestResult, clientInfos, speedtestHost)
          influxDb.saveMeasurementToDb(point)
        case _ =>
          println(
            "Download: %f | Upload: %f | Ping: %f | ServerId: %s | Provider: %s | Distance: %s".format(
              speedTestResult.download.megaBitPerSecond,
              speedTestResult.upload.megaBitPerSecond,
              speedTestResult.latency.floatValue(),
              speedTestResult.server.site,
              clientInfos.provider,
              distance
            )
          )
      }

      new Date().toInstant.getEpochSecond + interval.toSeconds
    } catch {
      case _: Exception => new Date().toInstant.getEpochSecond
    }
    while (nextRun < new Date().toInstant.getEpochSecond) {}
  } while (repeatable)

  ""

}
