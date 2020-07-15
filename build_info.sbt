import scala.sys.process._

enablePlugins(BuildInfoPlugin)

buildInfoPackage := "%s.%s".format(organization.value, name.value)

buildInfoOptions += BuildInfoOption.BuildTime

buildInfoKeys += BuildInfoKey.action("gitLastCommitHash") {
  "git rev-parse HEAD".!!.trim
}

buildInfoKeys += BuildInfoKey.action("organization") {
  organization.value
}

buildInfoPackage := "com.quadstingray.speedtest.app"

homepage := Some(url("https://quadstingray.github.io/speedtest/"))

scmInfo := Some(ScmInfo(url("https://github.com/QuadStingray/speedtest"), "https://github.com/QuadStingray/speedtest.git"))

developers := List(Developer("QuadStingray", "QuadStingray", "github@quadstingray.com", url("https://github.com/QuadStingray")))

licenses += ("Apache-2.0", url("https://github.com/QuadStingray/speedtest/blob/master/LICENSE"))