package com.quadstingray.speedtest.model

case class TestResult(
                       serverId: String,
                       host: String,
                       hostLocation: String,
                       downloadMbs: Double,
                       uploadMbs: Double,
                       ping: Double,
                       clientInfos: ClientInfos
)

