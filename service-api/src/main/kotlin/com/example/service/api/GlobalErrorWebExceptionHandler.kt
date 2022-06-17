package com.example.service.api

import com.example.service.api.service.NotFoundException
import com.example.service.api.service.ServiceException
import org.springframework.beans.factory.ObjectProvider
import org.springframework.boot.autoconfigure.web.ServerProperties
import org.springframework.boot.autoconfigure.web.WebProperties
import org.springframework.boot.autoconfigure.web.reactive.error.AbstractErrorWebExceptionHandler
import org.springframework.boot.autoconfigure.web.reactive.error.DefaultErrorWebExceptionHandler
import org.springframework.boot.web.error.ErrorAttributeOptions
import org.springframework.boot.web.reactive.error.ErrorAttributes
import org.springframework.context.ApplicationContext
import org.springframework.core.annotation.Order
import org.springframework.http.HttpStatus
import org.springframework.http.MediaType
import org.springframework.http.codec.ServerCodecConfigurer
import org.springframework.stereotype.Component
import org.springframework.web.reactive.function.BodyInserters
import org.springframework.web.reactive.function.server.*
import org.springframework.web.reactive.result.view.ViewResolver
import reactor.core.publisher.Mono
import java.util.stream.Collectors

@Component
@Order(-2)
class GlobalErrorWebExceptionHandler(
    errorAttributes: ErrorAttributes,
    webProperties: WebProperties,
    viewResolvers: ObjectProvider<ViewResolver>,
    serverCodecConfigurer: ServerCodecConfigurer,
    applicationContext: ApplicationContext,
    serverProperties: ServerProperties,
    val defaultExceptionHandler: DefaultErrorWebExceptionHandler = DefaultErrorWebExceptionHandler(
        errorAttributes,
        webProperties.resources, serverProperties.error, applicationContext
    )
) : AbstractErrorWebExceptionHandler(
    errorAttributes, webProperties.resources, applicationContext
) {

    init {
        this.setMessageWriters(serverCodecConfigurer.writers)
        this.setMessageReaders(serverCodecConfigurer.readers)
        this.setViewResolvers(viewResolvers.orderedStream().collect(Collectors.toList()))
        defaultExceptionHandler.setViewResolvers(viewResolvers.orderedStream().collect(Collectors.toList()))
        defaultExceptionHandler.setMessageWriters(serverCodecConfigurer.writers)
        defaultExceptionHandler.setMessageReaders(serverCodecConfigurer.readers)
    }

    override fun getRoutingFunction(errorAttributes: ErrorAttributes): RouterFunction<ServerResponse> {
        return RouterFunctions.route(
            RequestPredicates.all()
        ) { request -> renderErrorResponse(request) }
    }

    private fun renderErrorResponse(
        request: ServerRequest
    ): Mono<ServerResponse> {

        var errorPropertiesMap = mutableMapOf<String, Any?>()
        getErrorAttributes(
            request,
            ErrorAttributeOptions.defaults()
        ).forEach { entry -> errorPropertiesMap[entry.key] = entry.value }
        var status = HttpStatus.INTERNAL_SERVER_ERROR

        val exception = super.getError(request)
        if (exception is NotFoundException) {
            status = HttpStatus.NOT_FOUND
            errorPropertiesMap["status"] = status.value()
            errorPropertiesMap["error"] = exception.message
        } else if (exception is ServiceException) {
            status = HttpStatus.UNPROCESSABLE_ENTITY
            errorPropertiesMap["status"] = status.value()
            errorPropertiesMap["error"] = exception.message
        }

        return ServerResponse.status(status)
            .contentType(MediaType.APPLICATION_JSON)
            .body(BodyInserters.fromValue(errorPropertiesMap))
    }

}