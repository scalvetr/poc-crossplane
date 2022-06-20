package com.example.service.api

import java.time.ZonedDateTime

data class Topic(
    val name: String,
    val partitions: Int,
    val creationTime: ZonedDateTime?
)

enum class TopicStatus {
    CREATED,
    MARK_DELETED
}
