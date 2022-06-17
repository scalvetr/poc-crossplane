package com.example.service.api.service

import com.example.service.api.Topic
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import org.springframework.stereotype.Service

@Service
class TopicService(
    private val topicLimit: Int = 10,
    private val partitionLimit: Int = 10,
    private val database: MutableMap<String, MutableMap<String, Topic>> = mutableMapOf(),
    private val tenants: List<String> = listOf("tenant1", "tenant2", "tenant3")
) {

    suspend fun createTopic(tenant: String, topic: Topic) {
        checkTenant(tenant)
        checkTopic(topic)
        checkTopicLimit(tenant)
        database[tenant]?.put(topic.name, topic)
    }

    suspend fun updateTopic(tenant: String, topic: Topic) {
        checkTenant(tenant)
        checkTopic(topic)
        if (database[tenant]?.containsKey(topic.name) == false)
            throw NotFoundException("Topic with name ${topic.name} doesn't exist")

        database[tenant]?.put(topic.name, topic)
    }

    suspend fun getTopic(tenant: String, topicName: String): Topic {
        checkTenant(tenant)
        return database[tenant]?.get(topicName)
            ?: throw NotFoundException("Topic with name $topicName doesn't exist")
    }

    suspend fun deleteTopic(tenant: String, topicName: String) {
        checkTenant(tenant)
        if (database[tenant]?.remove(topicName) == null) {
            throw NotFoundException("Topic with name $topicName doesn't exist")
        }
    }

    fun getTopics(tenant: String): Flow<Topic> {
        checkTenant(tenant)
        return flow {
            database[tenant]?.values?.toList()?.onEach {
                emit(it)
            }
        }
    }

    private fun checkTopicLimit(tenant: String) {
        val size = database[tenant]?.values?.size ?: 0
        if (size >= topicLimit) {
            throw ServiceException("Reached the maximum number of topics for this tenant: $topicLimit")
        }
    }

    private fun checkTopic(topic: Topic) {
        if (topic.partitions > partitionLimit) {
            throw ServiceException("Invalid number of partitions ${topic.partitions}, cannot exceed $partitionLimit")
        }
    }

    private fun checkTenant(tenant: String) {
        if (!tenants.contains(tenant)) {
            throw ServiceException("Invalid tenant $tenant")
        }

        if (database[tenant] == null) {
            database[tenant] = mutableMapOf()
        }
    }
}
