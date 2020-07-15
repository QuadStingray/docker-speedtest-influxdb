package com.quadstingray.exception

final case class InfluxException(private val message: String, private val cause: Throwable = None.orNull) extends Exception(message, cause)
