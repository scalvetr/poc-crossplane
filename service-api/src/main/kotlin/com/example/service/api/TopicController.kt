package com.example.service.api

import com.example.service.api.service.TopicService
import kotlinx.coroutines.flow.Flow
import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.*

@RestController
@RequestMapping("/v1/topics")
class TopicController(
    private val topicService: TopicService
) {
    @PostMapping("")
    @ResponseStatus(HttpStatus.CREATED)
    suspend fun createTopic(@RequestHeader("x-tenant") tenant: String, @RequestBody topic: Topic) {
        topicService.createTopic(tenant, topic)
    }

    @PutMapping("/{topicName}")
    @ResponseStatus(HttpStatus.NO_CONTENT)
    suspend fun updateTopic(
        @RequestHeader("x-tenant") tenant: String,
        @PathVariable("topicName") topicName: String,
        @RequestBody topic: Topic
    ) {
        topicService.updateTopic(tenant, topic)
    }

    @GetMapping("/{topicName}")
    @ResponseStatus(HttpStatus.OK)
    suspend fun getTopic(
        @RequestHeader("x-tenant") tenant: String,
        @PathVariable("topicName") topicName: String
    ): Topic {
        return topicService.getTopic(tenant, topicName)
    }

    @DeleteMapping("/{topicName}")
    @ResponseStatus(HttpStatus.NO_CONTENT)
    suspend fun deleteTopic(@RequestHeader("x-tenant") tenant: String, @PathVariable("topicName") topicName: String) {
        topicService.deleteTopic(tenant, topicName)
    }

    @GetMapping("")
    @ResponseStatus(HttpStatus.OK)
    fun getTopics(@RequestHeader("x-tenant") tenant: String): Flow<Topic> {
        return topicService.getTopics(tenant)
    }
}
