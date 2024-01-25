"use strict";

function getQuery() {
    const query = window.location.search.slice(1)
    return Qs.parse(query)
}

function getQueryParam(name) {
    const query = getQuery()
    return query[name]
}

function setQuery(q) {
    const query = Qs.stringify(q)
    // window.location.search = query
    history.pushState('', '', '?' + query)
}

function clipboardWriteText(text) {
    if (navigator.clipboard) {
        const cb = navigator.clipboard
        return cb.writeText(text)
    }
}

function getJson(url, data, success, onError) {
    $.ajax({
        type: "GET",
        url: url,
        data: data,
        success: success,
        error: onError,
        dataType: "json"
    })
}

function postJson(url, data, success, onError) {
    $.ajax({
        type: "POST",
        url: url,
        data: data,
        success: success,
        error: onError,
        dataType: "json"
    })
}