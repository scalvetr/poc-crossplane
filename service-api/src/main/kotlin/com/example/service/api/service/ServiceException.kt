package com.example.service.api.service

class ServiceException(override val message: String) : RuntimeException(message) {
}