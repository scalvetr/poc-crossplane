package com.example.service.api.service

class NotFoundException(override val message: String) : RuntimeException(message) {
}