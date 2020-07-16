name := "speedtest-app"

organization := "com.quadstingray"

scalaVersion := crossScalaVersions.value.last

crossScalaVersions := List("2.13.3")

enablePlugins(GraalVMNativeImagePlugin)

graalVMNativeImageGraalVersion := Some("20.1.0")

graalVMNativeImageOptions ++= Seq(
  "-J-Xmx 9g",
  "--allow-incomplete-classpath",
  "--initialize-at-run-time=" +
    "com.typesafe.config.impl.ConfigImpl$EnvVariablesHolder," +
    "com.typesafe.config.impl.ConfigImpl$SystemPropertiesHolder",
  "--enable-https",
  "--enable-http",
  "--enable-url-protocols=https",
  "--enable-all-security-services",
  "--report-unsupported-elements-at-runtime",
  "--no-fallback",
  "--verbose",
  "--static"
)

libraryDependencies += "com.quadstingray" %% "speedtest" % "0.7.3"

libraryDependencies += "ch.qos.logback" % "logback-classic" % "1.2.3"

libraryDependencies += "com.typesafe.scala-logging" %% "scala-logging" % "3.9.2"

libraryDependencies += "org.influxdb" % "influxdb-java" % "2.19"

resolvers ++= Seq(
  "snapshots" at "https://oss.sonatype.org/content/repositories/snapshots",
  "releases" at "https://oss.sonatype.org/content/repositories/releases",
  "apache" at "https://repo.maven.apache.org/maven2/",
  "apache-snapshots" at "https://repository.apache.org/content/repositories/snapshots/",
  "maven2" at "https://repo1.maven.org/maven2",
  JCenterRepository
)

//libraryDependencies ~= {
//  _.map(_.exclude("com.squareup.okhttp3", "logging-interceptor"))
//}


// Remove after Influx Update
//dependencyOverrides += "com.squareup.okhttp3" % "logging-interceptor" % "4.7.2"
//
//dependencyOverrides += "com.squareup.retrofit2" % "retrofit" % "2.9.0"
//
//dependencyOverrides += "com.squareup.retrofit2" % "converter-moshi" % "2.9.0"