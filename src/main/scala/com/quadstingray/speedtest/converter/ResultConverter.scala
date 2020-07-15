package com.quadstingray.speedtest.converter

import java.util.concurrent.TimeUnit

import com.quadstingray.speedtest.database.InfluxDb._
import com.quadstingray.speedtest.model.{ClientInfos, TestResult}
import com.quadstingray.speedtest.ndt7.lib.SpeedTestResult
import org.influxdb.dto.Point

import scala.collection.mutable

case class ResultConverter() {

  def convertToPoint(measurementResult: SpeedTestResult, clientInfos: ClientInfos, speedtestHost: String): Point = {
    val measurement = Point.measurement("speedtest").time(System.currentTimeMillis, TimeUnit.MILLISECONDS)

    measurement.tag(TagHost, speedtestHost)
    measurement.tag(TagHostServer, measurementResult.connectionInfo.server)

    measurement.addField(FieldDlBandwidth, measurementResult.download.bytePerSeconds)
    measurement.addField(FieldUpBandwidth, measurementResult.upload.bytePerSeconds)
    measurement.addField(FieldLatency, measurementResult.latency)
    measurement.addField(FieldServerKey, measurementResult.server.site)
    measurement.addField(FieldLocation, measurementResult.server.city + ", " + measurementResult.server.country)
    measurement.addField(FieldExternalIp, measurementResult.connectionInfo.client)

    measurement.addField(FieldDistance, clientInfos.distance)
    measurement.addField(FieldClientProvider, clientInfos.provider)
    measurement.addField(FieldClientLocation, clientInfos.location)
    measurement.addField(FieldClientLocation, clientInfos.location)

    measurement.build
  }

  def convertToMap(measurementResult: SpeedTestResult, clientInfos: ClientInfos,speedtestHost: String): Map[String, Any] = {

    val mutableMap = mutable.Map[String, Any]()
    mutableMap.put(TagHost, speedtestHost)
    mutableMap.put(TagHostServer, measurementResult.connectionInfo.server)
    mutableMap.put(FieldDlBandwidth, measurementResult.download.bytePerSeconds)
    mutableMap.put(FieldUpBandwidth, measurementResult.upload.bytePerSeconds)
    mutableMap.put(FieldLatency, measurementResult.latency)
    mutableMap.put(FieldServerKey, measurementResult.server.site)
    mutableMap.put(FieldLocation, measurementResult.server.city + ", " + measurementResult.server.country)
    mutableMap.put(FieldExternalIp, measurementResult.connectionInfo.client)

    mutableMap.put(FieldDistance, clientInfos.distance)
    mutableMap.put(FieldClientProvider, clientInfos.provider)
    mutableMap.put(FieldClientLocation, clientInfos.location)

    mutableMap.toMap
  }

  def convertToResult(measurementResult: SpeedTestResult, clientInfos: ClientInfos): TestResult = {
    TestResult(
      measurementResult.server.site,
      measurementResult.connectionInfo.server,
      measurementResult.server.city + ", " + measurementResult.server.country,
      measurementResult.download.megaBitPerSecond,
      measurementResult.upload.megaBitPerSecond,
      measurementResult.latency,
      clientInfos
    )
  }

}
