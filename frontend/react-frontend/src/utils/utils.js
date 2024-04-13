import axios from "axios"

import * as AgentKeepAlive from "agentkeepalive";

const keepAliveAgent = new AgentKeepAlive({
    maxSockets: 128, // or 128 / os.cpus().length if running node across multiple CPUs
    maxFreeSockets: 128, // or 128 / os.cpus().length if running node across multiple CPUs
    timeout: 60000, // active socket keepalive for 60 seconds
    freeSocketTimeout: 30000, // free socket keepalive for 30 seconds
    keepAlive: true
})

const httpsKeepAliveAgent = new AgentKeepAlive.HttpsAgent({
    maxSockets: 128, // or 128 / os.cpus().length if running node across multiple CPUs
    maxFreeSockets: 128, // or 128 / os.cpus().length if running node across multiple CPUs
    timeout: 60000, // active socket keepalive for 30 seconds
    freeSocketTimeout: 30000, // free socket keepalive for 30 seconds,
    keepAlive: true
})

export const axiosInstance = axios.create({
    // Create an agent for both HTTP and HTTPS
    httpAgent: keepAliveAgent,
    httpsAgent: httpsKeepAliveAgent,
    timeout: 30000
})