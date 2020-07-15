import scala.sys.process._

version in Paradox := {
  if (isSnapshot.value) {
    val lastVersionString = "git tag -l".!!.split("\r?\n").last
    lastVersionString
  }
  else {
    version.value
  }
}

paradoxProperties += ("app-version" -> {
  if (isSnapshot.value) {
    val lastVersionString = "git tag -l".!!.split("\r?\n").last
    lastVersionString
  }
  else {
    version.value
  }
})

enablePlugins(ParadoxSitePlugin, ParadoxMaterialThemePlugin, GhpagesPlugin)

git.remoteRepo := "git@github.com:QuadStingray/speedtest.git"

sourceDirectory in Paradox := sourceDirectory.value / "main" / "paradox"

enablePlugins(ParadoxMaterialThemePlugin)

Compile / paradoxMaterialTheme := {
  ParadoxMaterialTheme()
    .withLogoIcon("storage")
    .withCopyright("© QuadStingray 2020")
    .withColor("teal", "indigo")
    .withRepository(uri("https://github.com/QuadStingray/speedtest"))
}

enablePlugins(SiteScaladocPlugin)

enablePlugins(GhpagesPlugin)

git.remoteRepo := "git@github.com:QuadStingray/speedtest.git"

ghpagesNoJekyll := true
