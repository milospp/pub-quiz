let configData = {
    "GATEWAY_SERVICE_URL": `http://${document.domain}:8000`,
    "AUTH_SERVICE_URL": `http://${document.domain}:8003`,
    "QUIZ_SERVICE_URL": `http://${document.domain}:8000`,
    "QUIZ_SOCKET_SERVICE_URL": `ws://${document.domain}:8002`,
    "STATS_SERVICE_URL": `http://${document.domain}:8004`
}

export default configData