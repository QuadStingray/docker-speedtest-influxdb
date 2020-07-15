package com.quadstingray.speedtest.database

import com.quadstingray.exception.InfluxException
import org.influxdb.dto.{ Point, Query }
import org.influxdb.{ BatchOptions, InfluxDB, InfluxDBFactory }

import scala.concurrent.duration._

case class InfluxDb(host: String, userName: String, password: String, database: String) {

  private var influxDB: InfluxDB = _

  def init(): Unit = {
    if (influxDB == null) {
      if (userName.trim != "")
        influxDB = InfluxDBFactory.connect(host, userName, password)
      else
        influxDB = InfluxDBFactory.connect(host)

      if (!influxDB.ping().isGood) {
        influxDB = null
        throw InfluxException("Could not check InfluxDB Connection (Host: %s | UserName: %s | Password: %s)")
      }

      influxDB.enableGzip()
      influxDB.enableBatch(BatchOptions.DEFAULTS)

      influxDB.query(new Query("CREATE DATABASE " + database))
      influxDB.setDatabase(database)

      val retentionPolicyName = "speedtest_ndt7_policy"
      val duration            = "INF"
      var shardDuration       = "30m"

      if (shardDuration.trim.isEmpty) {
        Duration(duration) match {
          case d: Duration if d > (6 * 30.5).days => shardDuration = "7d"
          case d: Duration if d <= 2.days         => shardDuration = "1h"
          case _                                  => shardDuration = "1h"
        }
      }

      val retentionPolicyCreationString = "CREATE RETENTION POLICY \"" + retentionPolicyName + "\" ON \"" + database + "\" DURATION " + duration + " SHARD DURATION " + shardDuration + " DEFAULT"
      val createRetentionPolicyResult   = influxDB.query(new Query(retentionPolicyCreationString))
      createRetentionPolicyResult.hasError
      influxDB.setRetentionPolicy(retentionPolicyName)
    }

  }

  def saveMeasurementToDb(point: Point): Unit = {
    influxDB.write(point)
    influxDB.flush()
    ""
  }

}

object InfluxDb {

  val TagHost             = "host"
  val FieldDlBandwidth    = "download"
  val FieldUpBandwidth    = "upload"
  val FieldLatency        = "ping"
  val FieldDistance       = "distance"
  val FieldServerKey      = "serverId"
  val FieldLocation       = "location"
  val FieldExternalIp     = "externalIp"
  val FieldClientProvider = "clientProvider"
  val FieldClientLocation = "clientLocation"

}
