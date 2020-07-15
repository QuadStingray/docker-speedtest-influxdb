import com.typesafe.sbt.packager.docker.Cmd

enablePlugins(JavaAppPackaging)
enablePlugins(DockerPlugin)

dockerBaseImage := "oracle/graalvm-ce:20.1.0-java11-ol8"

//dockerRepository := Some("git-registry.intern.iwelt.de/mtec")

dockerCommands += Cmd("ENV", "INTERVAL 30m")
dockerCommands += Cmd("ENV", "OUTPUT INFLUX")
dockerCommands += Cmd("ENV", "INFLUX_URL", "http://influx-db")
dockerCommands += Cmd("ENV", "INFLUX_USER", "admin")
dockerCommands += Cmd("ENV", "INFLUX_PASSWORD", "password")
dockerCommands += Cmd("ENV", "INFLUX_DB", "ndt-speedtest")

dockerUpdateLatest := false

daemonUserUid in Docker := None

daemonUser in Docker := "root"

dockerCommands += Cmd("ENTRYPOINT", "/opt/docker/bin/speedtest-app -Dinterval=$INTERVAL -Doutput=$OUTPUT -Dinflux_url=$INFLUX_URL -Dinflux_user=$INFLUX_USER -Dinflux_password=$INFLUX_PASSWORD -Dinflux_db=$INFLUX_DB")

