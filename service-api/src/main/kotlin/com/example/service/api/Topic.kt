package com.example.service.api

data class Topic(
    val name: String,
    val partitions: Int,
    val status: TopicStatus
)

enum class TopicStatus {
    CREATED,
    MARK_DELETED
}
