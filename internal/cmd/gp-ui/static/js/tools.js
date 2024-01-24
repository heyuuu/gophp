"use strict";

function getQuery() {
    const query = window.location.search.slice(1)
    return Qs.parse(query)
}

function setQuery(q) {
    const query = Qs.stringify(q)
    // window.location.search = query
    history.pushState('', '', '?' + query)
}